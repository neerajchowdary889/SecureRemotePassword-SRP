package server

import (
	"fmt"
	"math/big"
	"srp/NG_values"
)
func (ServerStoringDetails *ServerStoringDetails) GenerateB() (*TempServerDetails){

	b := NG_values.Generate64BitNumber()

	B := new(big.Int)

	// B = kv + g^b (mod N)
	K := new(big.Int)
    K.SetString(ServerStoringDetails.K, 16)

	kv := new(big.Int).Mul(K, ServerStoringDetails.V)

	pow_gb := new(big.Int).Exp(new(big.Int).SetUint64(uint64(ServerStoringDetails.G)), new(big.Int).SetUint64(b), ServerStoringDetails.N)

	B = new(big.Int).Add(kv, pow_gb)
	B = new(big.Int).Mod(B, ServerStoringDetails.N)

	ServerTempDetails := &TempServerDetails{
		B: B,
		b: b,
	}

	return ServerTempDetails
}

func (server_tempdetails *TempServerDetails) Server_ComputeU(A *big.Int) (string){
	// u = H(A | B)
	u := NG_values.H(append(A.Bytes(), server_tempdetails.B.Bytes()...))
	fmt.Println("Server_u: ",u)
	return u
}