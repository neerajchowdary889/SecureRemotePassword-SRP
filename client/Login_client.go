package client

import (
	"encoding/binary"
	"fmt"
	"math/big"
	"srp/NG_values"
)

func (user *ClientDetails) GenerateA() *ClientTempDetails {

	a := NG_values.Generate64BitNumber()

	A := new(big.Int)
	A = A.Exp(new(big.Int).SetUint64(uint64(user.G)), new(big.Int).SetUint64(a), user.N)

	aBytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(aBytes, a)

	ClientTempDetails := &ClientTempDetails{
		A: A,
		a: a,
	}

	return ClientTempDetails
}

func (user_tempdetails *ClientTempDetails) Client_ComputeU(B *big.Int) string {
	// u = H(A | B)
	user_tempdetails.B = B
	u := NG_values.H(append(user_tempdetails.A.Bytes(), B.Bytes()...))

	return u
}

func (user *ClientDetails) Compute_K_client(user_tempdetails *ClientTempDetails) (string) {
	// S = (B - kg^x) ^ (a + ux) (mod N)
	var Password string
	fmt.Println("Enter Password: ")
	fmt.Scan(&Password)

	x, _ := user.computeX(Password, false)

	fmt.Println(x)

	K := new(big.Int)
	K.SetString(user.K, 16)

	U := new(big.Int)
	U.SetString(user_tempdetails.u, 16)

	X := new(big.Int)
	X.SetString(x, 16)

	k_gx := new(big.Int).Mul(K, user.V)
	b_k_gx := new(big.Int).Sub(user_tempdetails.B, k_gx)

	ux := new(big.Int).Mul(U, X)
	a_ux := new(big.Int).Add(new(big.Int).SetUint64(user_tempdetails.a), ux)

	S := new(big.Int).Exp(b_k_gx, a_ux, user.N)
	K_client := NG_values.H(S.Bytes())

	return K_client
}
