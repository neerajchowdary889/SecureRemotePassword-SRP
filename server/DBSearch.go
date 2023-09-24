package server

import (
	"encoding/csv"
	// "log"
	"fmt"
	"math/big"
	"os"
	"srp/NG_values"
	"strconv"
)
func Searchcsv(username string)(*ServerStoringDetails, bool){
    file, err := os.Open("server/server.csv")
    if err != nil {
        fmt.Println("Error opening file:", err)
        return nil, false
    }
    defer file.Close()

    // Create a new CSV reader
    reader := csv.NewReader(file)

    // Read the CSV records one by one
    records, err := reader.ReadAll()
    if err != nil {
        fmt.Println("Error reading file:", err)
        return nil, false
    }

	temp := make([]string,0,7)

    // Search for the desired value in the CSV records
    searchValue := NG_values.H(username)
    for _, record := range records {
        for _, value := range record {
            if value == searchValue {
                fmt.Println("Found value:", value)
				for _, value := range record {
					temp = append(temp, value)
					// fmt.Println(temp)				
				}
			}else{
				fmt.Println("Error: Username not found")
				return nil, false
			}
        }
    }

	salt, err := strconv.ParseUint(temp[1], 10, 64)
	if err != nil {
		fmt.Println("Error converting Salt to uint64:", err)
		return nil, false
	}
	G, err := strconv.ParseUint(temp[2], 10, 8)
	if err != nil {
		fmt.Println("Error converting G to uint8:", err)
		return nil, false
	}	
	V := new(big.Int)
	V, ok := V.SetString(temp[4], 10)
	if !ok {
		fmt.Println("Error converting string to big.Int")
		return nil, false
	}
	N := new(big.Int)
	N, ok1 := N.SetString(temp[5], 10)
	if !ok1 {
		fmt.Println("Error converting string to big.Int")
		return nil, false
	}

	From_Server := &ServerStoringDetails{
		Username: temp[0],
		Salt: salt,
		G: uint8(G),
		K: temp[3],
		V: V,
		N: N,
	}

	return From_Server, true

}