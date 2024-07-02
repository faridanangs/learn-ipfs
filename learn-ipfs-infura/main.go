package main

import (
	"fmt"
	"os"
	"strings"

	ipfsApi "github.com/ipfs/go-ipfs-api" // v0.2.0
)

func main() {

	shell := ipfsApi.NewShell("http://127.0.0.1:8080/ipfs")
	cid, err := shell.Add(strings.NewReader("Infura IPFS - Getting started demo."))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Data successfully stored in IPFS: %v\n", cid)
}
