package coinblock

import (
	"demo_bitcoin/util"
	"github.com/boltdb/bolt"
	"log"
)

type BlockIterate struct {
	db  *bolt.DB
	cur []byte
}

func (it *BlockIterate) IsOver() bool {
	if len(it.cur) == 0 {
		return true
	}
	return false
}

func (it *BlockIterate) Next() *Block {

	var block Block

	it.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(util.DB_BUCKET))

		if !it.IsOver() {
			log.Printf("Find block :%x", it.cur)
			val := bucket.Get(it.cur)
			block = UnSerializeToBlock(val)
			it.cur = block.PrevHash
		}
		return nil
	})

	return &block
}

func (bc *BlockChain) NewIterate() *BlockIterate {
	return &BlockIterate{
		db:  bc.db,
		cur: bc.lastHash,
	}
}
