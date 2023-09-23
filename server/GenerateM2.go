package server

import(
	"math/big"
	"srp/NG_values"
)

func (ServerStoringDetails *ServerStoringDetails) GenerateM2(tempdetails *TempServerDetails, M_1 string) string{
	// M2 = H(A | M1 | K_server)

    M1 := new(big.Int)
    M1.SetString(M_1, 16)

    A_bytes := tempdetails.A.Bytes()
    M1_bytes := M1.Bytes()
    K_server_bigint, _ := new(big.Int).SetString(tempdetails.K_server, 16)
    K_server_bytes := K_server_bigint.Bytes()

    M2 := NG_values.H(append(A_bytes, append(M1_bytes, K_server_bytes...)...))

    return M2
}
