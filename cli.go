package main

import (
	"demo_bitcoin/coinblock"
	"demo_bitcoin/transaction"
	"strconv"

	//"demo_bitcoin/transaction"
	"demo_bitcoin/util"
	"fmt"
	"os"
)

type CLI struct {
	bc *coinblock.BlockChain
}

const Usage = `
	send --from A --to B --m 10 "A 给 B 10 Bit"
	addBlock --ad address "add address"
	getBit --ad address "get bit"
	printChain

`

func (cli *CLI) Run() {

	if len(os.Args) < 2 {
		fmt.Printf(Usage)
		return
	}

	switch os.Args[1] {

	case "send":
		//send --from A --to B --m 10 "A 给 B 10 Bit"
		if len(os.Args) == 8 && os.Args[2] == "--from" && os.Args[4] == "--to" && os.Args[6] == "--m" {

			from := os.Args[3]
			to := os.Args[5]
			amount, err := strconv.ParseFloat(os.Args[7], 64)
			if err != nil {
				fmt.Printf(Usage)
				return
			}

			baseTX := transaction.NewCoinBaseTX(util.AD_GOD)

			tx, err := coinblock.NewTransaction(from, to, amount, cli.bc)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			var txs []*transaction.Transaction
			txs = append(txs, tx)
			txs = append(txs, baseTX)

			b := coinblock.NewBlock(cli.bc.GetLashHast(), "Add Block", txs)
			cli.bc.AddBlock(b)

			fmt.Printf("From %s To %s  M :%s 交易成功!\n", from, to, os.Args[7])

		} else {
			fmt.Printf(Usage)
		}

	case "addBlock":

		if len(os.Args) == 4 && os.Args[2] == "--ad" {
			//tx := transaction.NewCoinBaseTX(os.Args[3])
			//b := NewBlock(cli.bc.lastHash, "cli-add", []*transaction.Transaction{&tx})
			//cli.bc.AddBlock(b)
		}

	case "printChain":
		cli.bc.View()

	case "getBit":
		if len(os.Args) == 4 && os.Args[2] == "--ad" {
			sum := cli.bc.FindUTXOsSumVal(os.Args[3])
			fmt.Printf("address %s: bit = %f\n", os.Args[3], sum)
		}
	default:
		sum := cli.bc.FindUTXOsSumVal(util.AD_GOD)
		fmt.Println(sum)
		fmt.Printf(Usage)

	}
}
