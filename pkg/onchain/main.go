package readsm

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	readni "ipt-onchain/pkg/onchain/node-ingress"
	readnr "ipt-onchain/pkg/onchain/node-rules"
	"log"
	"math/big"
	"os"
	"strconv"
)

var rpcURL string = os.Getenv("RPC_URL")
var addr string = os.Getenv("NODE_INGRESS_CONTRACT_ADDRESS")

//Readsm ...
func Readsm(c chan<- string) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(0)
	}

	key, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("Error while generating key")
	}

	ctx := context.Background()
	auth := bind.NewKeyedTransactor(key)
	opts := bind.CallOpts{
		Pending: true,
		From:    auth.From,
		Context: ctx,
	}

	fmt.Println("Instantiating Node Ingress contract")
	instContract, err := readni.NewNodeIngressCaller(common.HexToAddress(addr), client)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v", err)
	}

	//GetContractAddress of deployed NodeRules
	fmt.Println("Getting Noderules contract address")
	str := "rules"
	var rules [32]byte
	copy(rules[:], str)

	nodeRulesContractAddress, err := instContract.GetContractAddress(&opts, rules)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v", err)
	}

	fmt.Println("Creating NodeRules instance")
	instNodeRules, err := readnr.NewNodeRulesCaller(nodeRulesContractAddress, client)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v", err)
	}

	nodeRulesCallerSessionInstance := readnr.NodeRulesCallerSession{
		Contract: instNodeRules,
		CallOpts: opts,
	}

	fmt.Println("Interacting with NodeRules contract adddress...")
	len, err := nodeRulesCallerSessionInstance.GetSize()
	currentNodes := int(len.Uint64())
	if err != nil {
		log.Fatalf("could not instantiate contract: %v", err)
	}

	for i := 0; i < currentNodes; i++ {
		fmt.Println("Getting Enode", i+1)
		recenode, err := nodeRulesCallerSessionInstance.GetByIndex(big.NewInt(int64(i)))
		if err != nil {
			log.Fatalf("could not instantiate contract: %v", err)
		}

		ip := getIP(recenode.Ip)
		c <- ip
	}

	close(c)
}

type en struct {
	EnodeHigh [32]byte
	EnodeLow  [32]byte
	Ip        [16]byte
	Port      uint16
}

func getEnode(recenode en) string {
	pk := getPK(recenode.EnodeHigh, recenode.EnodeLow)
	ip := getIP(recenode.Ip)
	port := getPort(recenode.Port)
	return "enode://" + pk + "@" + ip + ":" + port
}

func getPort(p uint16) string {
	return strconv.FormatInt(int64(p), 10)
}

func getPK(high, low [32]byte) (pk string) {
	enodeHigh := fmt.Sprintf("%x", high)
	enodeLow := fmt.Sprintf("%x", low)
	return enodeHigh + enodeLow
}

func getIP(iprec [16]byte) string {
	ip1 := strconv.FormatInt(int64(iprec[12]), 10)
	ip2 := strconv.FormatInt(int64(iprec[13]), 10)
	ip3 := strconv.FormatInt(int64(iprec[14]), 10)
	ip4 := strconv.FormatInt(int64(iprec[15]), 10)
	ip := ip1 + "." + ip2 + "." + ip3 + "." + ip4
	return ip
}
