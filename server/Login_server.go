package server

import (
	// "fmt"
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
	server_tempdetails.A = A
	u := NG_values.H(append(A.Bytes(), server_tempdetails.B.Bytes()...))

	return u
}

func (server_tempdetails *TempServerDetails) Compute_K_server( ServerStoringDetails *ServerStoringDetails)(bool){
	// S = (A * v^u) ^ b (mod N)
	U := new(big.Int)
	U.SetString(server_tempdetails.u, 16)

	vu := new(big.Int).Exp(ServerStoringDetails.V, U, ServerStoringDetails.N)
	A_vu := new(big.Int).Mul(server_tempdetails.A, vu)

	A_vu_b := new(big.Int).Exp(A_vu, new(big.Int).SetUint64(server_tempdetails.b), ServerStoringDetails.N)

	K_server := NG_values.H(A_vu_b.Bytes())

	server_tempdetails.K_server = K_server
	return true
}

func Server_ComputeU_test(A *big.Int, B *big.Int) (string){
	// u = H(A | B)
	u := NG_values.H(append(A.Bytes(), B.Bytes()...))
	return u
}

func SetU(server_tempdetails *TempServerDetails, value string) {
    server_tempdetails.u = value
}

func GetU(server_tempdetails *TempServerDetails) string{
	return server_tempdetails.u
}