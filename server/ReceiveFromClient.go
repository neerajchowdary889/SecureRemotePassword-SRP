package server

import (
	// "fmt"
	"encoding/csv"
	"log"
	"math/big"
	"os"
	"strconv"
)

func UserSignUp(user map[string]interface{}) bool {
    var status bool
    defer func() {
        if err := recover(); err != nil {
            log.Println("Error:", err)
            status = false
        }
    }()

    // Open the CSV file for appending
    file, err := os.OpenFile("server/server.csv", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal("Error opening file:", err)
    }
    defer file.Close()

    // Create a new CSV writer
    writer := csv.NewWriter(file)
    defer writer.Flush()
	

    // Write the user data to the CSV file
    err = writer.Write([]string{
        user["Username"].(string),
        strconv.FormatUint(user["Salt"].(uint64), 10),
        strconv.Itoa(int(user["G"].(uint8))),
        user["K"].(string),
        user["V"].(*big.Int).String(),
        user["N"].(*big.Int).String(),
    })
    if err != nil {
        log.Fatal("Error writing to file:", err)
    }

    // Manually flush the buffer to ensure that the data is written correctly
    writer.Flush()

    status = true
    return status
}



