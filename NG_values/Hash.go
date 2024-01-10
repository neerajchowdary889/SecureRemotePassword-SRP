package NG_values

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func H(v interface{}) string{
	s := fmt.Sprintf("%v",v)
	hash := sha256.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))	
}