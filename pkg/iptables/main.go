package goipt

import (
	"fmt"
	"github.com/coreos/go-iptables/iptables"
)

var chain string = "ONCHAIN"

//MainIpt ...
func MainIpt(c <-chan string){
	ipt, err := iptables.New()
	//deletes and creates a new one chain into the specified table.
	err = ipt.ClearChain("filter", chain)
	if err != nil {
		fmt.Printf("ClearChain (of missing) failed: %v", err)
	} else {
		fmt.Println("Successful deletion and creation of custom chain '", chain, "'")
	}
	EnableEnodeOnIPTable(c)
	AddDefaultRules()	
}

//AddDefaultRules : Add the default rules for the linux machine
func AddDefaultRules() {
	fmt.Println("Appending default rules...")
	ipt, err := iptables.New()
	err = ipt.Append("filter", chain, "-p", "tcp", "--dport", "22", "-j", "ACCEPT")
	err = ipt.Append("filter", chain, "-p", "tcp", "--dport", "80", "-j", "ACCEPT")
	err = ipt.Append("filter", chain, "-p", "tcp", "--dport", "443", "-j", "ACCEPT")
	err = ipt.Append("filter", chain, "-j", "DROP")
	
	//err = ipt.Append("filter", chain, "-s", "0/0", "-j", "ACCEPT")
	if err != nil {
		fmt.Printf("Append failed: %v", err)
	}
}

//EnableEnodeOnIPTable enables 60606 port on the desired node IP
func EnableEnodeOnIPTable(c <-chan string) {
	ipt, err := iptables.New()	

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
		err = ipt.AppendUnique("filter", chain, "-p", "tcp", "-s", v, "--dport", "60606", "-j", "ACCEPT")
		err = ipt.AppendUnique("filter", chain, "-p", "udp", "-s", v, "--dport", "60606", "-j", "ACCEPT")
		//err = ipt.Append("filter", chain, "-s", "0/0", "-j", "ACCEPT")
		if err != nil {
			fmt.Printf("Append failed: %v", err)
		}
	}
}

func isIncluded(list []string, value string) bool {
	for _, val := range list {
		if val == value {
			return true
		}
	}
	return false
}
