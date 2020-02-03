package goipt

import (	
	"fmt"	
	"github.com/coreos/go-iptables/iptables"
)

var chain string = "onchain"

//Iptgreet ...
func Iptgreet(c <-chan string) {
	for v := range c {
		fmt.Println("Receiving values from Ethereum network", v)
	}
}

//AddDefaultRules : Add the default rules for the linux machine
func AddDefaultRules() {
	fmt.Println("Adding some default rules...")
}

//PrepareTable : ...
func PrepareTable() {

}
//SaveInitialTable : ...
func SaveInitialTable() {
	ipt, err := iptables.New()
	originaListChain, err := ipt.ListChains("filter") //there are three tables by default ==>
	//nat, filter mangle
	if err != nil {
		fmt.Printf("ListChains of Initial failed: %v", err)
	}
	fmt.Println(originaListChain)
}

//EnableEnodeOnIPTable enables 60606 port on the desired node IP
func EnableEnodeOnIPTable(c <-chan string) {
	ipt, err := iptables.New()	
	// Saving the list of chains before executing tests

	//deletes and creates a new one chain into the specified table.
	err = ipt.ClearChain("filter", chain)
	if err != nil {
		fmt.Printf("ClearChain (of missing) failed: %v", err)
	} else {
		fmt.Println("Successful deletion and creation of custom chain '", chain, "'")
	}

	// chain should be in listChain
	listChain, err := ipt.ListChains("filter")
	if err != nil {
		fmt.Printf("ListChains failed: %v", err)
	}
	if !isIncluded(listChain, chain) {
		fmt.Printf("ListChains doesn't contain the new chain %v", chain)
	}

	for v := range c {
		fmt.Println("Adding new node to the ip tables", v)
		//inserting a new rule
		err = ipt.Append("filter", chain, "-s", v, "-j", "ACCEPT")
		//err = ipt.Append("filter", chain, "-s", "0/0", "-j", "ACCEPT")
		if err != nil {
			fmt.Printf("Append failed: %v", err)
		}
	}

	// err = ipt.ClearChain("filter", chain)
	// if err != nil {
	// 	fmt.Printf("ClearChain (of missing) failed: %v", err)
	// }

	//Deleting a chain ONLY of it is empty
	// err = ipt.DeleteChain("filter", chain)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("Successfull deletion of empty chain '", chain, "'")

}

func isIncluded(list []string, value string) bool {
	for _, val := range list {
		if val == value {
			return true
		}
	}
	return false
}
