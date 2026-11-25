package dbcs

import (
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/Ptt-official-app/pttbbs-backend/types"
	"github.com/sirupsen/logrus"
)

func dbcsToUtf8(contentDBCS string) (content [][]*types.Rune) {
	if len(contentDBCS) == 0 {
		return nil
	}

	lines := strings.Split(contentDBCS, "\n")

	// remove last empty line (ending with "\n")
	if lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}

	content = make([][]*types.Rune, len(lines))
	color0 := types.DefaultColor
	for idx, each := range lines {
		content[idx], color0 = dbcsToUtf8PerLine(each, color0)
	}
	return content
}

// dbcsToUtf8PerLine
func dbcsToUtf8PerLine(contentDBCS string, color0 types.Color) (line []*types.Rune, newColor types.Color) {
	color0.IsReset = false
	if len(contentDBCS) == 0 {
		return []*types.Rune{}, color0
	}

	// 2. estimate len utf8
	lenDBCS := len(contentDBCS)
	expectedLenUtf8 := lenDBCS / 2

	// init
	line = make([]*types.Rune, 0, expectedLenUtf8)
	fullStartIdx := 0
	startIdx := 0
	escIdx := 0
	color1 := color0

	// 3. for-loop
	for escIdx < len(contentDBCS) {
		idx := strings.Index(contentDBCS[escIdx:], "\x1b[")
		if idx == -1 {
			break
		}
		idx += escIdx // add offset

		idxM := dbcsToUtf8PerLineIndexM(contentDBCS, idx)
		if idxM == -1 {
			escIdx = idx + 1
			continue
		}

		// 1. 1st parse: get new color
		newColor, isColor1, err := dbcsToUtf8PerLineParseColor(contentDBCS[idx:idxM+1], color0)
		if err != nil {
			escIdx = idx + 1
			continue
		}

		// 2. setup existing sentence
		if idx > startIdx {
			r := &types.Rune{
				Utf8:    contentDBCS[startIdx:idx],
				Color0:  color0,
				Color1:  color1,
				DBCSStr: contentDBCS[fullStartIdx:idx],
			}
			line = append(line, r)
			fullStartIdx = idx
		}

		// 3. 1st parse: hit color1
		if isColor1 {
			color1 = newColor
			theUtf8 := dbcsToUtf8PerLineNextUtf8(contentDBCS[idxM+1:])
			r := &types.Rune{
				Utf8:    theUtf8,
				Color0:  color0,
				Color1:  color1,
				DBCSStr: contentDBCS[fullStartIdx : idxM+len(theUtf8)+1],
			}
			line = append(line, r)

			color0 = color1
			startIdx = idxM + len(theUtf8) + 1
			escIdx = idxM + len(theUtf8) + 1
			fullStartIdx = idxM + len(theUtf8) + 1

			continue
		}

		// 1st parse: color0
		color0 = newColor
		color1 = color0
		escIdx = idxM + 1
		startIdx = idxM + 1

		// 2nd parse:
		if !strings.HasPrefix(contentDBCS[idxM+1:], "\x1b[") { // no color1: loop
			continue
		}

		// possible with color1
		idx1 := idxM + 1
		idxM1 := dbcsToUtf8PerLineIndexM(contentDBCS, idx1)
		if idxM1 == -1 {
			escIdx = idx1 + 1
			continue
		}

		newColor, isColor1, err = dbcsToUtf8PerLineParseColor(contentDBCS[idx1:idxM1+1], color1)
		if err != nil {
			escIdx = idx1 + 1
			continue
		}

		// 2nd parse: color0
		if !isColor1 {
			color0 = newColor
			color1 = newColor
			escIdx = idxM1 + 1
			startIdx = idxM1 + 1

			continue
		}

		// 2nd parse: color1
		color1 = newColor
		theUtf8 := dbcsToUtf8PerLineNextUtf8(contentDBCS[idxM1+1:])
		r := &types.Rune{
			Utf8:    theUtf8,
			Color0:  color0,
			Color1:  color1,
			DBCSStr: contentDBCS[fullStartIdx : idxM1+len(theUtf8)+1],
		}
		line = append(line, r)

		color0 = color1
		startIdx = idxM1 + len(theUtf8) + 1
		escIdx = idxM1 + len(theUtf8) + 1
		fullStartIdx = idxM1 + len(theUtf8) + 1
	}

	if len(contentDBCS[fullStartIdx:]) != 0 {
		r := &types.Rune{
			Utf8:    contentDBCS[startIdx:],
			Color0:  color0,
			Color1:  color0,
			DBCSStr: contentDBCS[fullStartIdx:],
		}
		line = append(line, r)

	}

	return line, color0
}

// dbcsToUtf8PerLineParseColor
//
// colorDBCS: includes "\x1b[" and "m"
func dbcsToUtf8PerLineParseColor(colorDBCS string, origColor types.Color) (color types.Color, isColor1 bool, err error) {
	colorDBCS = colorDBCS[2 : len(colorDBCS)-1]

	if len(colorDBCS) == 0 {
		// \x1b[m
		color = types.DefaultColor
		// color.IsReset = true
		return color, false, nil
	}

	color = origColor
	color.IsReset = false

	theList := strings.Split(colorDBCS, ";")
	isAlreadyForeground := false
	for idx, each := range theList {
		if idx == 0 && len(each) == 0 { // reset
			// no need to do color.IsReset
			// because dbcsIntegrateColor always set IsReset as false
			color = types.DefaultColor
			continue
		}

		eachNum, err := strconv.Atoi(each)
		if err != nil {
			continue
		}

		// https://github.com/ptt/pttbbs/blob/master/daemon/boardd/convert.c#L63
		if eachNum > 100 {
			isColor1 = true
			if eachNum < 110 { // XXX hack that sometimes fg/bg is presented only with 0~7
				if !isAlreadyForeground {
					isAlreadyForeground = true
					eachNum += 30
				} else {
					eachNum += 40
				}
			}

			if eachNum >= 110 && eachNum < 120 {
				eachNum -= 110
			} else {
				eachNum -= 100
			}
		}

		switch {
		case eachNum == 0:
			color = types.DefaultColor
		case eachNum == 1:
			color.Highlight = true
		case eachNum == 5:
			color.Blink = true
		case eachNum >= int(types.COLOR_FOREGROUND_BLACK) && eachNum <= int(types.COLOR_FOREGROUND_WHITE):
			color.Foreground = types.ColorMap(eachNum)
			isAlreadyForeground = true
		case eachNum >= int(types.COLOR_BACKGROUND_BLACK) && eachNum <= int(types.COLOR_BACKGROUND_WHITE):
			color.Background = types.ColorMap(eachNum)
		default:
			logrus.Warnf("(%v/%v) invalid color-code: each: %v eachNum: %v", idx, len(theList), each, eachNum)
		}
	}

	return color, isColor1, nil
}

func dbcsToUtf8PerLineIndexM(contentDBCS string, idx int) (idxM int) {
	// \x1b[1;33;35m
	for idx2, each := range contentDBCS[idx+2:] {
		if each == 'm' {
			return idx + 2 + idx2
		}
		if each != ';' && (each < '0' || each > '9') {
			return -1
		}
	}

	return -1
}

func dbcsToUtf8PerLineNextUtf8(contentDBCS string) (theUtf8 string) {
	if len(contentDBCS) == 0 {
		return ""
	}

	if strings.HasPrefix(contentDBCS, "\x1b[") { // the next is a new color-encoding
		idxM := dbcsToUtf8PerLineIndexM(contentDBCS, 0)
		if idxM != -1 {
			return ""
		}
	}

	// slice by the bytes actually consumed: string(rune) turns an invalid byte
	// into 3-byte U+FFFD, and callers use len(theUtf8) as the offset into contentDBCS.
	_, size := utf8.DecodeRuneInString(contentDBCS)

	return contentDBCS[:size]
}

func dbcsToUtf8PurifyColor(contentDBCS string) (purifiedDBCS string) {
	purified := make([]byte, 0, len(contentDBCS))
	for idx := strings.Index(contentDBCS, "\x1b"); len(contentDBCS) > 0 && idx >= 0; idx = strings.Index(contentDBCS, "\x1b") {
		if idx != 0 {
			purified = append(purified, []byte(contentDBCS[:idx])...)
			contentDBCS = contentDBCS[idx:]
		}

		if len(contentDBCS) < 3 || contentDBCS[1] != '[' {
			contentDBCS = contentDBCS[1:]
			continue
		}

		idxM := strings.Index(contentDBCS, "m")
		if idxM == -1 { // unable to find "m"
			contentDBCS = contentDBCS[2:]
			continue
		}

		// skip to idxM (excluded)
		contentDBCS = contentDBCS[idxM+1:]
	}

	if len(contentDBCS) > 0 {
		purified = append(purified, []byte(contentDBCS)...)
	}

	if len(purified) == 0 {
		return ""
	}

	return string(purified)
}
