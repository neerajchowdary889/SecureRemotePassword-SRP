// salt length should of 64bit long

package NG_values

import(
	"encoding/binary"
	"crypto/rand"
)
const Salt_BitSize int = 64

func Generate64BitNumber() uint64{
	salt := make([]byte,Salt_BitSize)
	_,err := rand.Read(salt)
	if err != nil{
		panic(err)
	}
	return binary.BigEndian.Uint64(salt)
}