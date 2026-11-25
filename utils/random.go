package utils

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"

	"github.com/google/uuid"
)

// avoid '-' and '_' in leading char.
func GenRandomString() string {
	for {
		random := uuid.Must(uuid.NewRandom())
		bin, _ := random.MarshalBinary()
		str := base64.URLEncoding.EncodeToString(bin)
		if str[0] == '-' || str[0] == '_' {
			continue
		}
		return str[:22]
	}
}

// GenRandomInt64
// https://stackoverflow.com/questions/32349807/how-can-i-generate-a-random-int-using-the-crypto-rand-package/32350135
//
// Returns a uniform random number in [0, theMax).
// theMax must be positive (rand.Int panics otherwise).
func GenRandomInt64(theMax int64) (int64, error) {
	if theMax <= 0 {
		return 0, ErrInvalidMax
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(theMax))
	if err != nil {
		return 0, err
	}
	return nBig.Int64(), nil
}
