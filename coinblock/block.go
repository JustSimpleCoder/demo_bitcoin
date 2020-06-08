package coinblock

import (
	"bytes"
	"crypto/sha256"
	"demo_bitcoin/transaction"
	"demo_bitcoin/util"
	"encoding/gob"
	"fmt"
	"log"
	"time"
)

type Block struct {
	PrevHash     []byte
	Hash         []byte // 比特币实际存储在另外的存储中
	Data         []byte
	Transactions []*transaction.Transaction
	Version      uint64
	MerkelRoot   []byte //todo 校验的hash
	TimeStamp    uint64
	Difficulty   uint64
	Nonce        uint64
}

func NewBlock(pHash []byte, data string, ts []*transaction.Transaction) *Block {
	block := Block{
		PrevHash:     pHash,
		Hash:         []byte{},
		Data:         []byte(data),
		Transactions: ts,
		Version:      00,
		TimeStamp:    uint64(time.Now().Unix()),
		Difficulty:   0,
		Nonce:        0,
	}
	block.MakeMerkelRoot()
	//block.SetHash()

	p := NewProofOfWork(&block)
	n, h := p.Run()
	block.Nonce, block.Hash = n, h
	return &block
}

func BlockSerializeToByte(b interface{}) []byte {

	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)
	err := encoder.Encode(b)
	if err != nil {
		log.Panic("BlockSerializeToByte 错误:", err)
	}

	return buffer.Bytes()
}
func UnSerializeToBlock(b []byte) Block {

	var block Block
	decoder := gob.NewDecoder(bytes.NewBuffer(b))
	decoder.Decode(&block)
	return block

}

func (b *Block) MakeMerkelRoot() {

	var root []byte
	for _, tx := range b.Transactions {
		root = append(root, tx.TXID...)
	}

	sum256 := sha256.Sum256(root)
	b.MerkelRoot = sum256[:]
}

func (b *Block) SetHash() {

	tmp := [][]byte{
		util.Uint64ToByte(b.Version),
		util.Uint64ToByte(b.TimeStamp),
		util.Uint64ToByte(b.Difficulty),
		util.Uint64ToByte(b.Nonce),
		b.PrevHash,
		b.MerkelRoot,
		b.Data,
	}
	blockInfo := bytes.Join(tmp, []byte{})
	sum256 := sha256.Sum256(blockInfo)
	b.Hash = sum256[:]
}

func (b *Block) View() {

	fmt.Println("")
	fmt.Printf("Ver:%d\nPrevHash: %x\nHash: %x\nData: %s\n", b.Version, b.PrevHash, b.Hash, b.Data)
	fmt.Printf("Nonce: %d\n", b.Nonce)
	fmt.Printf("Difficulty: %d\n", b.Difficulty)
	timeFormat := time.Unix(int64(b.TimeStamp), 0).Format(util.TIME_F)
	fmt.Printf("TimeStamp: %s\n", timeFormat)

	fmt.Printf("MerkelRoot: %x\n", b.MerkelRoot)
	fmt.Printf("Nonce: %d\n", b.Nonce)
	fmt.Printf("AD:%s\n", b.Transactions[0].TXInputs[0].Sig)

}

func NewGodBlock() *Block {
	tx := transaction.NewCoinBaseTX(util.AD_GOD)
	return NewBlock([]byte{}, "God", []*transaction.Transaction{tx})
}

func NewTransaction(from, to string, amount float64, bc *BlockChain) (*transaction.Transaction, error) {

	utxos, calc := bc.FindNeedUTXOs(from, amount)
	if calc < amount {
		log.Printf("!! Find " + from + " Money Not Enough")
		return &transaction.Transaction{}, fmt.Errorf("%s余额不足! 当前:%f", from, calc)
	}

	var inputs []transaction.TXInput
	var outputs []transaction.TXOutput

	for id, indexs := range utxos {
		for _, i := range indexs {
			input := transaction.TXInput{
				TXid:  []byte(id),
				Index: int64(i),
				Sig:   from,
			}
			inputs = append(inputs, input)

		}

	}

	output := transaction.TXOutput{
		Value:      amount,
		PubKeyHash: []byte(to),
	}

	outputs = append(outputs, output)

	if calc > amount {
		outputs = append(outputs, transaction.TXOutput{
			Value:      calc - amount,
			PubKeyHash: []byte(from),
		})
	}

	tx := transaction.Transaction{
		TXID:      nil,
		TXInputs:  inputs,
		TXOutputs: outputs,
	}
	tx.SetHash()

	return &tx, nil
}
