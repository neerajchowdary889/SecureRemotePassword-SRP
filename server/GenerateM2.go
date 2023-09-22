package server

import(
	"math/big"
	"srp/NG_values"
)

func (ServerStoringDetails *ServerStoringDetails) GenerateM2(tempdetails *TempServerDetails, M_1 string) string{
	// M2 = H(A | M1 | K_server)

	M1 := new(big.Int)
	M1.SetString(M_1, 16)

	K_server := new(big.Int)
	K_server.SetString(tempdetails.K_server, 16)

	M2 := NG_values.H(append(tempdetails.A.Bytes(), append(M1.Bytes(), K_server.Bytes()...)...))
	return M2
}

// I think this function might return wrong output...