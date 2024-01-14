package main

import (
	"flag"
	"fmt"
	"srp/Nodes"
)

func main() {
	// Define flags
	serverFlag := flag.Bool("server", false, "Run as server")
	clientFlag := flag.Bool("client", false, "Run as client")

	// Parse the flags
	flag.Parse()

	// Use the flags
	if *serverFlag {
		fmt.Println("Running Server")
		Nodes.Server_Execution()
	} else if *clientFlag {
		fmt.Println("Running Client")
		Nodes.Client_Execution()
	} else {
		fmt.Println("Please specify either -server or -client flag")
	}
}
