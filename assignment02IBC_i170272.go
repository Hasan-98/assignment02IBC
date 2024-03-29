package assignment02IBC

import (
	"crypto/sha256"
	"fmt"
	"strconv"
)

const miningReward = 100
const rootUser = "Satoshi"

type Block struct {
	Spender     map[string]int
	Receiver    map[string]int
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

func CalculateBalance(userName string, chainHead *Block) int {

	var temp *Block
	var in = 0
	var out = 0
	//var net = 0
	var len = 0

	if userName == "Satoshi" {

		for temp = chainHead; temp != nil; {

			for key, value := range temp.Receiver {
				//	fmt.Println("cpmapring=", userName, key)
				if userName == key {

					in = in + value

				}
				len = len + 100

			}

			for key, value := range temp.Spender {

				if userName == key {

					out = out + value
				}
			}

			temp = temp.PrevPointer
		}
		//	fmt.Println("returning valur=", (in-out)+len)
		return (in - out) + len
	}

	if userName != "Satoshi" {

		for temp = chainHead; temp != nil; {

			for key, value := range temp.Receiver {
				if userName == key {

					in = in + value

				}

			}

			for key, value := range temp.Spender {

				if userName == key {

					out = out + value
				}
			}

			temp = temp.PrevPointer
		}
		//	fmt.Println("returning valur=", (in - out))
		return (in - out)
	}
	return 0
}

func InsertBlock(spendingUser string, receivingUser string, miner string, amount int, chainHead *Block) *Block {
	var temp *Block = new(Block)

	if chainHead == nil {

		if spendingUser == "" && receivingUser == "" && amount == 0 && miner == "Satoshi" {

			fmt.Println(" valid conditions for the transactions -> gensis done")
			temp.PrevHash = ""
			temp.PrevPointer = nil
			temp.Spender = make(map[string]int)

			temp.Spender[""] = amount
			temp.Receiver = make(map[string]int)
			temp.Receiver[""] = amount
			temp.CurrentHash = CalculateHash(temp)
			//fmt.Println("hash=", temp.CurrentHash)
			return temp

		}

		fmt.Println("invalid transaction denied")
		temp = nil
		return temp

	}

	if chainHead != nil {
		var temp *Block = new(Block)
		//fmt.Println(spendingUser, CalculateBalance(spendingUser, chainHead))
		if miner == "Satoshi" && CalculateBalance(spendingUser, chainHead) >= amount {
			//	fmt.Println("valid transaction checks444444")
			temp.PrevHash = chainHead.CurrentHash
			temp.PrevPointer = chainHead
			temp.Spender = make(map[string]int)

			temp.Spender[spendingUser] = amount
			temp.Receiver = make(map[string]int)
			temp.Receiver[receivingUser] = amount
			temp.CurrentHash = CalculateHash(temp)
			//fmt.Println("hash=", temp.CurrentHash)
			fmt.Println(" valid conditions for the transactions -> done")
			return temp

		}

		if miner != "Satoshi" || CalculateBalance(spendingUser, chainHead) < amount {
			fmt.Println("not valid conditions for the transactions")
			temp = nil
			return chainHead
		}
		fmt.Println("invalid transaction denied")
		return chainHead

	}
	fmt.Println("invalid transaction denied")
	return chainHead
}

func CalculateHash(inputBlock *Block) string {
	//	a := len(inputBlock.Spender)
	//	b := len(inputBlock.Receiver)
	var temp string

	for key, value := range inputBlock.Receiver {

		temp = temp + key + strconv.Itoa(value)

	}

	for key, value := range inputBlock.Spender {

		temp = temp + key + strconv.Itoa(value)

	}

	//fmt.Println(temp)
	obj := sha256.New()
	obj.Write([]byte(fmt.Sprintf("%x", temp)))

	return fmt.Sprintf("%x", obj.Sum(nil))
}

func ListBlocks(chainHead *Block) {
	var temp *Block
	for temp = chainHead; temp != nil; {
		fmt.Println("transaction=")
		for key, value := range temp.Receiver {
			fmt.Println("receiving=", key, "money:", value)

		}

		for key, value := range temp.Spender {
			fmt.Println("sender=", key, "money:", value)

		}
		fmt.Println("--------")
		temp = temp.PrevPointer
	}
}

func VerifyChain(chainHead *Block) {

	var temp *Block
	for temp = chainHead; temp != nil; {
		temp2 := temp.PrevPointer
		//fmt.Println("transaction=", temp.transactions)
		if temp2 != nil {
			pHash := CalculateHash(temp.PrevPointer)
			if pHash != temp.PrevHash {
				fmt.Println("change detected")
				return
			}
		}
		temp = temp.PrevPointer
	}
}
