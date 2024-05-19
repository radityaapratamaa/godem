package helper

import (
	"crypto/rand"
	"math/big"
)

func GetRandomMinMax(min, max int) int64 {
	random, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return random.Int64() + int64(min)
}
