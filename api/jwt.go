package api

import "github.com/golang-jwt/jwt/v4"

func ParseJwt(raw string, secret []byte) (tok *jwt.Token, err error) {
	tok, err = jwt.Parse(raw, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	return tok, err
}

func ParseClaimString(claim jwt.MapClaims, idx string) (ret string, err error) {
	ret_i, ok := claim[idx]
	if !ok {
		return "", nil
	}
	ret, ok = ret_i.(string)
	if !ok {
		return "", ErrInvalidToken
	}

	return ret, nil
}

func ParseClaimInt(claim jwt.MapClaims, idx string) (ret int, err error) {
	ret_i, ok := claim[idx]
	if !ok {
		return 0, nil
	}
	// XXX it's float64 in go-jwt, but it's ok to have second(time)-level inaccuracy for expire-ts.
	ret_f64, ok := ret_i.(float64)
	if !ok {
		return 0, ErrInvalidToken
	}

	return int(ret_f64), nil
}
