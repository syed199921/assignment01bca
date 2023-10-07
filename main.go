package main

import (
	"fmt"

	"github.com/syed199921/assignment01bca/assignment01bca"
)

func main() {
	BlockChain := assignment01bca.BlockChain{}
	BlockchainCopy := assignment01bca.LinkedList{}
	BlockChain.AddBlock("Send 1 BTC to Syed", 1)
	BlockChain.AddBlock("Send 2 BTC to Syed", 2)
	BlockChain.AddBlock("Send 3 BTC to Syed", 3)

	BlockChain.DisplayBlockchain()

	BlockChain.ChangeBlock("108498b290f79e67082ace7344260494d3b840812864e302a44e09ebf2879513", "Send 30 BTC to Syed", &BlockchainCopy)

	fmt.Println("The blockchain is modified: ", BlockChain.VerifyChain(&BlockchainCopy))
}
