package dbcs

import (
	"bytes"
	"strconv"

	"github.com/Ptt-official-app/pttbbs-backend/types"
)

func Utf8ToDBCS(utf8 [][]*types.Rune) (dbcsBytes [][]byte) {
	dbcsBytes = make([][]byte, len(utf8))

	color := types.DefaultColor
	for idx, eachLine := range utf8 {
		dbcsBytes[idx], color = utf8ToDBCSByLine(eachLine, color)
	}

	return dbcsBytes
}

// utf8ToDBCSByLine
//
// We can directly concat the DBCS within line
func utf8ToDBCSByLine(line []*types.Rune, color types.Color) (lineDBCS []byte, newColor types.Color) {
	lineDBCS = make([]byte, 0, DEFAULT_LINE_BYTES)
	newColor = color
	var eachDBCS []byte
	for idx, each := range line {
		isAddCR := types.IS_CARRIAGE_RETURN && (idx == len(line)-1)
		eachDBCS, newColor = utf8ToDBCSByRune(each, newColor, isAddCR)
		lineDBCS = append(lineDBCS, eachDBCS...)
	}

	return lineDBCS, newColor
}

// utf8ToDBCSByRune
//
// 1. check whether we've already have the DBCS (already parsed)
func utf8ToDBCSByRune(theRune *types.Rune, origColor types.Color, isAddCR bool) (theDBCS []byte, newColor types.Color) {
	theRune.Utf8ToDBCS(origColor, isAddCR)
	return theRune.DBCS, theRune.Color1
}

// dbcsToBig5
//
// dbcs: with '\n' ending.
//
// 1. split by '\n'
// 2. parse dbcs to big5 per line.
func dbcsToBig5(dbcs []byte) (big5 [][]*types.Rune) {
	if len(dbcs) == 0 {
		return nil
	}
	if dbcs[len(dbcs)-1] == '\n' { // remove last '\n'
		dbcs = dbcs[:len(dbcs)-1]
	}
	if len(dbcs) == 0 {
		return nil
	}

	lines := bytes.Split(dbcs, []byte{'\n'})
	big5 = make([][]*types.Rune, len(lines))
	color0 := types.DefaultColor
	for idx, each := range lines {
		big5[idx], color0 = dbcsToBig5PerLine(each, color0)
	}

	return big5
}

func big5ToUtf8(a [][]*types.Rune) [][]*types.Rune {
	for _, each := range a {
		for _, eachEach := range each {
			eachEach.Big5ToUtf8()
		}
	}
	return a
}

// dbcsToBig5PerLine
//
// Assuming dbcs is within same line.
// Goal: able to:
//  1. separate big5 and color.
//  2. big5 with same color belongs to same rune
//
// Inspired from https://github.com/ptt/pttbbs/blob/master/mbbsd/pfterm.c#L838
// (Setting dbcs-state)
//
// have startIdx and endIdx as the indicator of the start and end of the big5-bytes. (endIdx excluded.)
//
// The conditions:
// if not encountered the color: assuming same color and put to same big5str (endIdx++)
// if encountered the color(continuous-colors as parsed-1-color), we need to determine separator and the color:
//
//	if in DBCS_STAT_LEAD (the previous char is the 0th Big5):
//	   ended with previous-char,
//	   start a new rune with startIdx = previous-char, parse color as color1.
//	if in DBCS_STAT_NONE (the previous char is not Big5, color is set as color0):
//	   ended with currentIdx, set previous color1 as color0,
//	   start a new rune with currentIdx.
//	if in DBCS_STAT_TAIL (the previous char is the 1st Big5):
//	   if color1 is not set, set color1 as color0.
//	   ended with endIdx, start a new rune with currentIdx.
//	if in DBCS_COLOR (the previous char is color):
//	   continue parse the color.
//
// Implementation:
// 1. check isCarriage.
// 2. estimate len big5
// 3. for-loop
// 4. check the last rune.
// 5. defensive programming for the last rune.
// 6. add carriage back if isCarriage.
func dbcsToBig5PerLine(dbcs []byte, color0 types.Color) ([]*types.Rune, types.Color) {
	if len(dbcs) == 0 {
		return []*types.Rune{}, color0
	}

	// 1. check isCarriage
	isCarriage := false
	if dbcs[len(dbcs)-1] == '\r' { // remove '\r' in the end
		isCarriage = true
		dbcs = dbcs[:len(dbcs)-1]
	}

	// 2. estimate len big5
	lenDBCS := len(dbcs)
	expectedLenBig5 := lenDBCS / 2

	// init
	big5 := make([]*types.Rune, 0, expectedLenBig5)
	dbcsStat := DBCS_STATE_NONE
	startIdx := 0
	color1 := color0
	dbcs0Pos := -1
	// 3. for-loop
	for idx := 0; idx < len(dbcs); {
		ch := dbcs[idx]
		if ch != '\x1b' { // not the color-code.
			switch dbcsStat {
			case DBCS_STATE_LEAD:
				dbcsStat = DBCS_STATE_TAIL
			case DBCS_STATE_NONE:
				if ch >= 0x80 { // is dbcs0
					dbcsStat = DBCS_STATE_LEAD
					dbcs0Pos = idx
					color1 = types.InvalidColor
				}
			case DBCS_STATE_TAIL: // previous ch is tail.
				if color1.Foreground != types.COLOR_INVALID { // dbcs. need to set a new rune.
					eachDBCS := dbcs[startIdx:idx]
					r := &types.Rune{
						Big5:   dbcsToBig5PurifyColor(eachDBCS),
						Color0: color0,
						Color1: color1,
						DBCS:   eachDBCS,
					}
					big5 = append(big5, r)

					startIdx = idx
					color0 = color1 // color0 becomes color1 now.

				}
				if ch >= 0x80 { // is dbcs0 again
					dbcsStat = DBCS_STATE_LEAD
					dbcs0Pos = idx
					color1 = types.InvalidColor

				} else { // done with tail. reset as NONE
					dbcsStat = DBCS_STATE_NONE
					color1 = types.InvalidColor
				}
			}
			idx++
		} else { // is the start of the color
			switch dbcsStat {
			case DBCS_STATE_LEAD: // the previous char is the 0th big5
				if color1.Foreground == types.COLOR_INVALID {
					color1 = color0
				}

				eachDBCS := dbcs[startIdx:dbcs0Pos]
				if len(eachDBCS) > 0 {
					r := &types.Rune{
						Big5:   dbcsToBig5PurifyColor(eachDBCS),
						Color0: color0,
						Color1: color1,
						DBCS:   eachDBCS,
					}
					big5 = append(big5, r)
					color0 = color1 // color0 becomes color1 now.
				}
				startIdx = dbcs0Pos
			case DBCS_STATE_NONE:
				eachDBCS := dbcs[startIdx:idx]
				if len(eachDBCS) > 0 {
					r := &types.Rune{
						Big5:   dbcsToBig5PurifyColor(eachDBCS),
						Color0: color0,
						Color1: color0,
						DBCS:   eachDBCS,
					}
					big5 = append(big5, r)
				}
				startIdx = idx
			case DBCS_STATE_TAIL:
				if color1.Foreground == types.COLOR_INVALID {
					color1 = color0
				}
				eachDBCS := dbcs[startIdx:idx]
				r := &types.Rune{
					Big5:   dbcsToBig5PurifyColor(eachDBCS),
					Color0: color0,
					Color1: color1,
					DBCS:   eachDBCS,
				}
				big5 = append(big5, r)
				startIdx = idx
				color0 = color1 // color0 becomes color1 now.

				dbcsStat = DBCS_STATE_NONE // dealt with tail, re-starting from none.
				color1 = types.InvalidColor
			}

			color, nBytes := dbcsParseColor(dbcs[idx:])
			if dbcsStat != DBCS_STATE_LEAD {
				color0 = dbcsIntegrateColor(color0, color)
				color1 = color0 // color1 should be the same as color0 once color0 is set.
			} else {
				color1 = dbcsIntegrateColor(color1, color)
			}
			idx += nBytes
		}
	}

	// the last rune
	switch dbcsStat {
	case DBCS_STATE_NONE:
		eachDBCS := dbcs[startIdx:]
		if len(eachDBCS) > 0 {
			r := &types.Rune{
				Big5:   dbcsToBig5PurifyColor(eachDBCS),
				Color0: color0,
				Color1: color0,
				DBCS:   eachDBCS,
			}
			big5 = append(big5, r)
		}
		startIdx = len(dbcs)
	case DBCS_STATE_TAIL: // previous ch is tail.
		if color1.Foreground != types.COLOR_INVALID { // dbcs. need to set a new rune.
			eachDBCS := dbcs[startIdx:]
			if len(eachDBCS) > 0 {
				r := &types.Rune{
					Big5:   dbcsToBig5PurifyColor(eachDBCS),
					Color0: color0,
					Color1: color1,
					DBCS:   eachDBCS,
				}
				big5 = append(big5, r)
			}

			color0 = color1 // color0 becomes color1 now.
			startIdx = len(dbcs)
		} else {
			eachDBCS := dbcs[startIdx:]
			if len(eachDBCS) > 0 {
				r := &types.Rune{
					Big5:   dbcsToBig5PurifyColor(eachDBCS),
					Color0: color0,
					Color1: color0,
					DBCS:   eachDBCS,
				}
				big5 = append(big5, r)
			}
			startIdx = len(dbcs)
		}
	}
	// defensive programming
	if startIdx < len(dbcs) {
		eachDBCS := dbcs[startIdx:]
		r := &types.Rune{
			Big5:   dbcsToBig5PurifyColor(eachDBCS),
			Color0: color0,
			Color1: color0,
			DBCS:   eachDBCS,
		}
		big5 = append(big5, r)
		startIdx = len(dbcs) //nolint // consistent with programming pattern
	}

	if isCarriage {
		if len(big5) == 0 {
			r := &types.Rune{
				Big5:   nil,
				Color0: color0,
				Color1: color0,
				DBCS:   []byte{},
			}
			big5 = append(big5, r)
		}
		theDBCS := big5[len(big5)-1].DBCS
		big5[len(big5)-1].DBCS = append(theDBCS, '\r')
	}

	return big5, color0
}

func dbcsToBig5PurifyColor(dbcs []byte) []byte {
	purified := make([]byte, 0, len(dbcs))

	p_dbcs := dbcs
	for idx := bytes.Index(p_dbcs, []byte{'\x1b'}); len(p_dbcs) > 0 && idx >= 0; idx = bytes.Index(p_dbcs, []byte{'\x1b'}) {
		if idx != 0 {
			purified = append(purified, p_dbcs[:idx]...)
			p_dbcs = p_dbcs[idx:]
		}

		if len(p_dbcs) < 3 || p_dbcs[1] != '[' { // not valid control char.
			p_dbcs = p_dbcs[1:]
			continue
		}

		idxM := bytes.Index(p_dbcs, []byte{'m'})
		if idxM == -1 { // unable to find 'm', set nBytes to include '\x1b[', p_dbcs as next.
			p_dbcs = p_dbcs[2:]
			continue
		}
		// skip to idxM (excluded)
		p_dbcs = p_dbcs[idxM+1:]
	}

	if len(p_dbcs) > 0 {
		purified = append(purified, p_dbcs...)
	}

	if len(purified) == 0 {
		return nil
	}

	return purified
}

func dbcsIntegrateColor(origColor types.Color, newColor types.Color) (color types.Color) {
	if newColor.IsReset {
		newColor.IsReset = false
		return newColor
	}

	color = origColor
	if newColor.Foreground != types.COLOR_INVALID {
		color.Foreground = newColor.Foreground
	}
	if newColor.Background != types.COLOR_INVALID {
		color.Background = newColor.Background
	}
	if newColor.Highlight {
		color.Highlight = true
	}
	if newColor.Blink {
		color.Blink = true
	}

	return color
}

func dbcsParseColor(dbcs []byte) (color types.Color, nBytes int) {
	color = types.InvalidColor

	p_dbcs := dbcs
	nBytes = 0

	for len(p_dbcs) > 0 && bytes.HasPrefix(p_dbcs, []byte{'\x1b'}) { // p_dbcs always start with control-code.
		if len(p_dbcs) < 3 || p_dbcs[1] != '[' { // 0th is control-code, but 1st is not '[', set nBytes to include '\x1b', p_dbcs as next, and continue find the next color
			nBytes++
			p_dbcs = p_dbcs[1:]
			continue
		}
		idxM := bytes.Index(p_dbcs, []byte{'m'})
		if idxM == -1 { // unable to find 'm', set nBytes to include "\x1b[", p_dbcs as next, and continue find the next color.
			nBytes += 2
			p_dbcs = p_dbcs[2:]
			continue
		}

		// found complete color-code
		if idxM == 2 { //'\x1b[m': reset
			color = types.ResetColor
			nBytes += idxM + 1
			p_dbcs = p_dbcs[idxM+1:]
			continue
		}

		colorCodeList := bytes.Split(p_dbcs[2:idxM], []byte{';'})

		for idx, each := range colorCodeList {
			if idx == 0 && len(each) == 0 { // reset
				color = types.ResetColor
				continue
			}

			intColor, err := strconv.Atoi(string(each))
			if err != nil { // not numbers
				continue
			}
			switch {
			case intColor == 0:
				color = types.ResetColor
			case intColor == 1:
				color.Highlight = true
			case intColor == 5:
				color.Blink = true
			case intColor >= 30 && intColor <= 37:
				color.Foreground = types.ColorMap(intColor)
			case intColor >= 40 && intColor <= 47:
				color.Background = types.ColorMap(intColor)
			}
		}

		// set nBytes to include m, p_dbs as next, and continue find the next color
		nBytes += idxM + 1
		p_dbcs = p_dbcs[idxM+1:]
	}

	return color, nBytes
}
