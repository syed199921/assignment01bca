package assignment01bca

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

type Block struct {
	transaction  string
	nonce        int
	previousHash string
	hash         string
	previous     *Block
}

type Node struct {
	transaction  string
	nonce        int
	previousHash string
	hash         string
	next         *Node
	previous     *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

type BlockChain struct {
	tail *Block
}

func (list *LinkedList) AddNode(transaction string, nonce int, previousHash string, hash string) {
	newNode := &Node{transaction, nonce, previousHash, hash, nil, nil}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.previous = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
}

func (bc *BlockChain) AddBlock(transaction string, nonce int) {

	newBlock := &Block{transaction, nonce, "", "", bc.tail}
	if bc.tail != nil {
		newBlock.previousHash = bc.tail.hash
	} else {

		newBlock.previousHash = ""
	}
	str := newBlock.transaction + strconv.Itoa(newBlock.nonce) + newBlock.previousHash
	newBlock.hash = CalculateHash(str)

	bc.tail = newBlock
}

func (bc *BlockChain) DisplayBlockchain() {
	fmt.Println("Blockchain:")
	for block := bc.tail; block != nil; block = block.previous {
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
		fmt.Printf("Transaction: %s \n Nonce: %d \n Hash: %x \n Previous Hash: %x", block.transaction, block.nonce, block.hash, block.previousHash)
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
		fmt.Println()
	}
}

func (bc *BlockChain) ChangeBlock(blockHash string, newTransaction string, blockchainCopy *LinkedList) {
	for block := bc.tail; block != nil; block = block.previous {

		hash := fmt.Sprintf("%x", []byte(block.hash))

		if hash == blockHash {

			block.transaction = newTransaction

		}
	}

	for block := bc.tail; block != nil; block = block.previous {
		newNode := &Node{block.transaction, block.nonce, block.previousHash, block.hash, nil, nil}

		if blockchainCopy.head == nil {
			blockchainCopy.head = newNode
			blockchainCopy.tail = newNode
		} else {
			newNode.previous = blockchainCopy.tail
			blockchainCopy.tail.next = newNode
			blockchainCopy.tail = newNode
		}
	}

	var startNode *Node = nil

	for node := blockchainCopy.tail; node != nil; node = node.previous {

		nodeHash := fmt.Sprintf("%x", []byte(node.hash))

		if nodeHash == blockHash {

			startNode = node
		}

	}

	for node := startNode; node != nil; node = node.previous {

		if node.previous != nil {
			node.hash = CalculateHash(node.transaction + strconv.Itoa(node.nonce) + node.previousHash)
			node.previous.previousHash = node.hash
		} else {
			node.hash = CalculateHash(node.transaction + strconv.Itoa(node.nonce) + node.previousHash)
		}

	}
	fmt.Println("Modified Blockchain:")
	for node := blockchainCopy.head; node != nil; node = node.next {
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
		fmt.Printf("Transaction: %s \n Nonce: %d \n Hash: %x \n Previous Hash: %x", node.transaction, node.nonce, node.hash, node.previousHash)
		fmt.Println("------------------------------------------------------------------------------------------------------------------")
		fmt.Println()
	}
}

// The VerifyChain function returns false if the blockchain is not modified and true if it is modified
func (bc *BlockChain) VerifyChain(list *LinkedList) bool {
	if list.head != nil {
		for block, node := bc.tail, list.head; block != nil || node != nil; block, node = block.previous, node.next {
			if block.hash != node.hash {
				return true
			} else {
				continue
			}
		}
	}
	return false
}
func CalculateHash(stringToHash string) string {
	h := sha256.New()
	h.Write([]byte(stringToHash))
	return string(h.Sum(nil))
}
