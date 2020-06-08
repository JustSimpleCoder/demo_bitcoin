package coinblock

import (
	"demo_bitcoin/transaction"
	"demo_bitcoin/util"
	"github.com/boltdb/bolt"
	"log"
)

type BlockChain struct {
	//blocks []*Block
	db       *bolt.DB
	lastHash []byte
}

func NewBlockChain() *BlockChain {

	var lastHash []byte
	db, err := bolt.Open(util.BC_DB, 0600, nil)
	//defer db.Close()
	if err != nil {
		log.Panic(err)
	}

	db.Update(func(tx *bolt.Tx) error {

		bucket := tx.Bucket([]byte(util.DB_BUCKET))
		if bucket == nil {
			if bucket, err = tx.CreateBucket([]byte(util.DB_BUCKET)); err != nil {
				log.Panic(err)
			}
		}

		lastHash = bucket.Get([]byte(util.LAST_KEY))

		return nil
	})

	bc := BlockChain{db: db, lastHash: lastHash}

	log.Printf("lasthash: %x", lastHash)

	if len(lastHash) == 0 {
		log.Println("add God")
		block := NewGodBlock()
		bc.AddBlock(block)
	}
	return &bc
}

func (bc *BlockChain) GetLashHast() []byte {

	return bc.lastHash
}

func (bc *BlockChain) AddBlock(block *Block) {

	bc.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(util.DB_BUCKET))
		if bucket == nil {
			log.Panic("Bucket Nil")
		}

		bucket.Put(block.Hash, BlockSerializeToByte(block))
		bucket.Put([]byte(util.LAST_KEY), block.Hash)

		//log.Printf("Add block\nPH:%x\nH:%x\nD:%s", block.PrevHash, block.Hash, block.Data)
		bc.lastHash = block.Hash
		log.Printf("Lasthash :%x", bc.lastHash)

		return nil
	})

	//b := NewBlock(bc.blocks[len(bc.blocks)-1].Hash, data)
	//bc.blocks = append(bc.blocks, b)

}

func (bc *BlockChain) View() {
	log.Printf("bc last: %x", bc.lastHash)
	it := bc.NewIterate()
	for !it.IsOver() {
		block := it.Next()
		block.View()
	}

}

func (bc *BlockChain) FindUTXOs(address string) []transaction.TXOutput {
	var utxos []transaction.TXOutput
	transactions := bc.FindTransactions(address)
	for _, tx := range transactions {
		for _, output := range tx.TXOutputs {
			if address == string(output.PubKeyHash) {
				utxos = append(utxos, output)
			}
		}
	}

	return utxos

}

func (bc *BlockChain) FindUTXOsSumVal(address string) float64 {
	var sum float64
	for _, output := range bc.FindUTXOs(address) {
		sum += output.Value
	}
	return sum
}

func (bc *BlockChain) FindNeedUTXOs(from string, amount float64) (map[string][]uint64, float64) {

	utxos := make(map[string][]uint64)
	var calc float64
	transactions := bc.FindTransactions(from)

FINDOVER:
	for _, tx := range transactions {
		for txIndex, output := range tx.TXOutputs {
			if from == string(output.PubKeyHash) {
				calc += output.Value
				utxos[string(tx.TXID)] = append(utxos[string(tx.TXID)], uint64(txIndex))
				if calc >= amount {
					break FINDOVER
				}
			}
		}
	}

	return utxos, calc
}

func (bc *BlockChain) FindTransactions(address string) []*transaction.Transaction {

	var txs []*transaction.Transaction
	spentOutput := make(map[string][]int64)

	it := bc.NewIterate()
	for !it.IsOver() {
		block := it.Next()
		//block.View()
		for _, tx := range block.Transactions {
		NextOutput:
			for i, output := range tx.TXOutputs {
				if txinputs, ok := spentOutput[string(tx.TXID)]; ok {
					for _, index := range txinputs {
						if index == int64(i) {
							break NextOutput
						}
					}
				}

				if address == string(output.PubKeyHash) {
					txs = append(txs, tx)
				}

			}

			if !tx.IsCoinBase() {
				for _, input := range tx.TXInputs {

					if address == input.Sig {

						spentOutput[string(input.TXid)] = append(spentOutput[string(input.TXid)], input.Index)

					}
				}
			}
		}

	}

	return txs
}
