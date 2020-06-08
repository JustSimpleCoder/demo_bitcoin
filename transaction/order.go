package transaction

import (
	"bytes"
	"crypto/sha256"
	"demo_bitcoin/util"
	"encoding/gob"
	"log"
)

// 1.定义结果交易结构
// 2.交易方法
// 3.挖矿交易.
// 4. ...调整程序

type Transaction struct {
	TXID      []byte
	TXInputs  []TXInput
	TXOutputs []TXOutput
}

type TXInput struct {
	// 交易id
	// output 的索引
	//解锁的脚本

	TXid  []byte
	Index int64
	Sig   string
}

type TXOutput struct {
	// 金额
	// 锁定脚本

	Value      float64
	PubKeyHash []byte
}

func (tx *Transaction) SetHash() {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(tx)
	if err != nil {
		log.Panic("Transaction Hash Panic", err)
	}

	hash := sha256.Sum256(buffer.Bytes())
	tx.TXID = hash[:]

	//log.Printf("New TX id:  %x", tx.TXID)

}

func NewCoinBaseTX(address string) *Transaction {

	in := TXInput{
		TXid:  []byte{},
		Index: -1,
		Sig:   "God",
	}

	out := TXOutput{
		Value:      util.REWORD,
		PubKeyHash: []byte(address),
	}

	transaction := Transaction{
		TXID:      []byte{},
		TXInputs:  []TXInput{in},
		TXOutputs: []TXOutput{out},
	}

	transaction.SetHash()
	return &transaction
}

func (tx *Transaction) IsCoinBase() bool {

	if len(tx.TXInputs) == 1 && bytes.Equal(tx.TXInputs[0].TXid, []byte{}) && tx.TXInputs[0].Index == -1 {
		return true
	}

	return false
}
