package readsm

import (
	"context"
	"fmt"
	readni "ipt-onchain/pkg/onchain/node-ingress"
	readnr "ipt-onchain/pkg/onchain/node-rules"
	"log"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var rpcHTTPURL string = os.Getenv("RPC_HTTP_URL")
var rpcWSURL string = os.Getenv("RPC_WS_URL")
var addr string = os.Getenv("NODE_INGRESS_CONTRACT_ADDRESS")
var from *big.Int = getHead()

var delayToReadSM time.Duration = 60
var delayToCheckNewConnections time.Duration = 30

type en struct {
	EnodeHigh [32]byte
	EnodeLow  [32]byte
	Ip        [16]byte
	Port      uint16
}

func getHead() *big.Int {
	client, _ := getClientAndOpts(rpcHTTPURL)
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return header.Number
}

func getClientAndOpts(rpc string) (*ethclient.Client, bind.CallOpts) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		log.Fatal("Error:", err)
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
	return client, opts
}

func getNodeRulesContractAddress() common.Address {
	client, opts := getClientAndOpts(rpcHTTPURL)
	fmt.Println("Instantiating Node Ingress contract")
	instContract, err := readni.NewNodeIngressCaller(common.HexToAddress(addr), client)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v", err)
	}

	fmt.Println("Getting Noderules contract address")
	str := "rules"
	var rules [32]byte
	copy(rules[:], str)

	nodeRulesContractAddress, err := instContract.GetContractAddress(&opts, rules)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v", err)
	}

	return nodeRulesContractAddress
}

func createNodeRulesCallerSessionInstance() readnr.NodeRulesCallerSession {
	client, opts := getClientAndOpts(rpcHTTPURL)
	nodeRulesContractAddress := getNodeRulesContractAddress()

	fmt.Println("Creating NodeRules instance")
	instNodeRules, err := readnr.NewNodeRulesCaller(nodeRulesContractAddress, client)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v", err)
	}

	nodeRulesCallerSessionInstance := readnr.NodeRulesCallerSession{
		Contract: instNodeRules,
		CallOpts: opts,
	}

	return nodeRulesCallerSessionInstance
}

//Readsm ...
func Readsm() <-chan string {
	c := make(chan string)
	nodeRulesCallerSessionInstance := createNodeRulesCallerSessionInstance()

	fmt.Println("Interacting with NodeRules contract adddress...")
	len, err := nodeRulesCallerSessionInstance.GetSize()
	currentNodes := int(len.Uint64())
	if err != nil {
		log.Fatalf("could not instantiate contract: %v", err)
	}

	go func() {
		for {
			for i := 0; i < currentNodes; i++ {
				fmt.Println("Getting Enode", i+1)
				recenode, err := nodeRulesCallerSessionInstance.GetByIndex(big.NewInt(int64(i)))
				if err != nil {
					log.Fatalf("could not instantiate contract: %v", err)
				}

				ip := getIP(recenode.Ip)
				c <- ip
			}
			//close(c)
			fmt.Println("Waiting for: ", delayToReadSM, "minutes")
			time.Sleep(delayToReadSM * time.Minute)
		}
	}()

	return c
}

//GetAddedNodeLogs ...
func GetAddedNodeLogs() <-chan string {
	c := make(chan string)

	nodeRulesAbi, err := abi.JSON(strings.NewReader(string(readnr.NodeRulesABI)))
	if err != nil {
		log.Fatal("Error while getting the ABI for nodeRules", err)
	}

	client, _ := getClientAndOpts(rpcWSURL)
	nodeRulesContractAddress := getNodeRulesContractAddress()

	ctx := context.Background()

	go func() {
		for {
			// fmt.Println("Waiting for: ",delayToCheckNewConnections,"minutes")
			// time.Sleep(delayToCheckNewConnections * time.Second)
			header := getHead()

			fmt.Println("header is", header.String())

			query := ethereum.FilterQuery{
				FromBlock: from,
				ToBlock:   header,
				Addresses: []common.Address{
					nodeRulesContractAddress,
				},
			}
			logs, err := client.FilterLogs(ctx, query)
			if err != nil {
				log.Fatal("Error while getting logs", err)
			}

			from = header //updating for next reading

			fmt.Println("logs are", logs)

			for _, vLog := range logs {
				nodeAddedEvent := struct {
					NodeAdded bool
					EnodeHigh [32]byte
					EnodeLow  [32]byte
					EnodeIp   [16]byte
					EnodePort uint16
					Raw       types.Log // Blockchain specific contextual infos
				}{}
				err := nodeRulesAbi.Unpack(&nodeAddedEvent, "NodeAdded", vLog.Data)
				if err != nil {
					log.Fatal("Error while decoding", err)
				}

				///////////Identifying the log has the right signature////////////
				eventSignature := []byte("NodeAdded(bool,bytes32,bytes32,bytes16,uint16)")
				hash := crypto.Keccak256Hash(eventSignature)
				fmt.Println("hash.hex", hash.Hex())
				var topics [4]string
				for i := range vLog.Topics {
					topics[i] = vLog.Topics[i].Hex()
				}

				fmt.Println("topic received", topics[0])

				ip := getIP(nodeAddedEvent.EnodeIp)
				fmt.Println("received ip is", ip, nodeAddedEvent.NodeAdded)
				if topics[0] == hash.Hex() {
					fmt.Println("Node will be added")
					c <- ip
				}
			}
			// close(c)
			fmt.Println("Waiting .... ")
			time.Sleep(delayToCheckNewConnections * time.Second)
		}
	}()

	return c
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
