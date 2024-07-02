package main

import (
	"context"
	"fmt"

	"github.com/ipfs/boxo/coreiface/path"
	"github.com/ipfs/kubo/client/rpc"
)

func main() {
	// "Connect" to local node
	node, err := rpc.NewLocalApi()
	if err != nil {
		fmt.Print(err)
		return
	}
	// Pin a given file by its CID
	ctx := context.Background()
	cid := "bafkreidtuosuw37f5xmn65b3ksdiikajy7pwjjslzj2lxxz2vc4wdy3zku"
	p, _ := path.NewPath(cid)
	err = node.Pin().Add(ctx, p)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(p)
}
