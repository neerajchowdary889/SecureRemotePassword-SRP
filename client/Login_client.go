package client

import (
	"fmt"
	"math/big"
	"srp/NG_values"
)

func (user *ClientDetails) Client_Login() bool {

	a := NG_values.Generate64BitNumber()

	A := new(big.Int)
	A = A.Exp(new(big.Int).SetUint64(uint64(user.G)), new(big.Int).SetUint64(a), user.N)

	fmt.Println("A: ",A)
	return true
}
