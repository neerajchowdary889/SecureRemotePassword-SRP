package client

import (
	"encoding/binary"
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

func (user *ClientDetails) Compute_K_client(user_tempdetails *ClientTempDetails, Password string) (bool) {
	// S = (B - kg^x) ^ (a + ux) (mod N)

	x, _ := user.computeX(Password, false)

	K := new(big.Int)
	K.SetString(user.K, 16)

	U := new(big.Int)
	U.SetString(user_tempdetails.u, 16)

	X := new(big.Int)
	X.SetString(x, 16)

	gx := new(big.Int).Exp(new(big.Int).SetUint64(uint64(user.G)), X, user.N)

	k_gx := new(big.Int).Mul(K, gx)
	b_k_gx := new(big.Int).Sub(user_tempdetails.B, k_gx)

	ux := new(big.Int).Mul(U, X)
	a_ux := new(big.Int).Add(new(big.Int).SetUint64(user_tempdetails.a), ux)

	S := new(big.Int).Exp(b_k_gx, a_ux, user.N)
	K_client := NG_values.H(S.Bytes())

	user_tempdetails.K_client = K_client
	return true
}

func (user_tempdetails *ClientTempDetails) Get_a() uint64{
	return user_tempdetails.a
}

func (user_tempdetails *ClientTempDetails) Set_a(a uint64) {
    user_tempdetails.a = a
}

func (user_tempdetails *ClientTempDetails) Set_u(u string) {
    user_tempdetails.u = u
}