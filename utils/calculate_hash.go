package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func CalculateHash(bytes []byte) string {
	hasher := md5.New()
	hasher.Write(bytes)
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash[:32]
}
