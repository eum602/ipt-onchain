package main

import (
	"fmt"
	. "ipt-onchain/pkg/iptables"
	. "ipt-onchain/pkg/onchain"
)

func main() {
	c := fanIn(Readsm(), GetAddedNodeLogs())
	//call IPTABLES
	//Iptgreet(c)
	MainIpt(c)
	fmt.Println("All done")
}
