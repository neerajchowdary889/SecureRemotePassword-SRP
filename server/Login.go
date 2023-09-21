package server

import (
	"fmt"
	"math/big"
	"srp/NG_values"
)
func (ServerStoringDetails *ServerStoringDetails) Login() (bool){

	a := NG_values.Generate64BitNumber()

	A := new(big.Int)
	A = A.Exp(new(big.Int).SetUint64(uint64(ServerStoringDetails.G)), new(big.Int).SetUint64(a), ServerStoringDetails.N)

	fmt.Println(A)
	return true
}
