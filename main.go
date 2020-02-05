package main

import (
	"fmt"
	. "ipt-onchain/pkg/iptables"
	. "ipt-onchain/pkg/onchain"
)

func main() {
	//read from ethereum
	c := make(chan string)
	go Readsm(c)

	//call IPTABLES
	EnableEnodeOnIPTable(c)
	fmt.Println("All done")
}
