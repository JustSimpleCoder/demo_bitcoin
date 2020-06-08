package main

import (
	"demo_bitcoin/coinblock"
)

func main() {
	//fmt.Println("~~~~~~~~~~~")
	//tmptry.DBTest()
	//tmptry.SerDemo()
	//tmptry.Break3()
	//tmptry.MapDemo()
	blockChain := coinblock.NewBlockChain()
	c := CLI{
		bc: blockChain,
	}

	c.Run()

	//for i, block := range blockChain.blocks {
	//	fmt.Println("~~~~~~~~~~~")
	//	fmt.Printf("高度:%d\n前Hash: %x\n当前Hash: %x\nData: %s\n", i, block.PrevHash, block.Hash, block.Data)
	//
	//}
}
