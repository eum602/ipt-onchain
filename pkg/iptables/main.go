package goipt

import (
	"fmt"
	"os"

	"strconv"

	"github.com/coreos/go-iptables/iptables"
)

var chain string = "ONCHAIN"
var chainToAppendCustomChain = "INPUT"

var defaultSecurityRules = "DEFAULTSECURITYRULES"
var addDefaultSecurityRules = getDefaultSecurityRulesConfig()

//MainIpt ...
func MainIpt(c <-chan string) {
	ipt, err := iptables.New()

	err = ipt.ClearChain("filter", chain)
	if err != nil {
		fmt.Printf("ClearChain (of missing) failed: %v", err)
	} else {
		fmt.Println("Successful deletion and creation of custom chain '", chain, "'")
	}

	insertChain(ipt, chain, chainToAppendCustomChain, 1)

	err = ipt.ClearChain("filter", defaultSecurityRules)
	if err != nil {
		fmt.Printf("ClearChain (of missing) failed: %v", err)
	} else {
		fmt.Println("Successful deletion and creation of custom chain '", defaultSecurityRules, "'")
	}

	if addDefaultSecurityRules {
		insertChain(ipt, defaultSecurityRules, chainToAppendCustomChain, 2)
		//Adding Default rules
		AddDefaultRules(defaultSecurityRules)
	} else { //Delete the default security chain and remove from main chain
		fmt.Println("Making sure no previous default security chain exist")
		err = ipt.Delete("filter", chainToAppendCustomChain, "-s", "0/0", "-j", defaultSecurityRules)
		if err != nil {
			fmt.Println("No default security chain found...")
		}
		err = ipt.DeleteChain("filter", defaultSecurityRules)
		if err != nil {
			fmt.Println("Something went wrong while deleting security rules")
		}

	}

	// AddDefaultRules(chainToAppendCustomChain)
	EnableEnodeOnIPTable(c)
}

func insertChain(ipt *iptables.IPTables, chain, chainToAppendCustomChain string, pos int) {
	isCustomChainAdded := searchChain(chain, chainToAppendCustomChain)
	if isCustomChainAdded {
		fmt.Println("Found chain ", chain, " on table ", chainToAppendCustomChain, " skipping adding process ...")
	} else {
		fmt.Println("Adding chain ", chain)
		//Add custom chain
		err := ipt.Insert("filter", chainToAppendCustomChain, pos, "-s", "0/0", "-j", chain)
		if err != nil {
			fmt.Printf("Error while adding custom chain %v", err)
		}
	}

}

func searchChain(chainToSearch, chainWhereToSearch string) bool {
	ipt, err := iptables.New()
	rules, err := ipt.List("filter", chainWhereToSearch)
	if err != nil {
		fmt.Printf("List failed: %v", err)
	}

	for _, v := range rules {
		if v == "-A INPUT -j "+chainToSearch {
			return true
		}
	}

	return false
}

//AddDefaultRules : Add the default rules for the linux machine
func AddDefaultRules(chain string) {
	fmt.Println("Appending default rules...")
	ipt, err := iptables.New()
	err = ipt.AppendUnique("filter", chain, "-i", "lo", "-j", "ACCEPT")
	err = ipt.AppendUnique("filter", chain, "-m", "conntrack", "--ctstate", "RELATED,ESTABLISHED", "-j", "ACCEPT")
	err = ipt.AppendUnique("filter", chain, "-p", "tcp", "--dport", "22", "-j", "ACCEPT")
	err = ipt.AppendUnique("filter", chain, "-p", "tcp", "--dport", "80", "-j", "ACCEPT")
	err = ipt.AppendUnique("filter", chain, "-p", "tcp", "--dport", "443", "-j", "ACCEPT")
	err = ipt.AppendUnique("filter", chain, "-j", "DROP")

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
		err = ipt.AppendUnique("filter", chain, "-p", "tcp", "-s", v, "--dport", "4040", "-j", "ACCEPT")
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

func getDefaultSecurityRulesConfig() bool {
	v, err := strconv.ParseBool(os.Getenv("ADD_DEFAULT_SECURITY_RULES"))
	if err != nil {
		fmt.Println("Found and error while getting security rule variable (ADD_DEFAULT_SECURITY_RULES) config from environment ... defaulting  to false")
		return false
	}
	return v
}
