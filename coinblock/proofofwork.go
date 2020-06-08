package coinblock

import (
	"bytes"
	"crypto/sha256"
	"demo_bitcoin/util"
	"log"
	"math/big"
)

type ProofOfWork struct {
	b      *Block
	target *big.Int //大的数
}

func NewProofOfWork(b *Block) *ProofOfWork {

	tmpS := "000011000000000000000000000000000000000000000000000000000000000001111"
	bi := big.Int{}
	bi.SetString(tmpS, 16)

	p := ProofOfWork{
		b:      b,
		target: &bi,
	}

	return &p
}

func (p *ProofOfWork) Run() (uint64, []byte) {

	var nonce uint64
	b := p.b
	var hash [32]byte

	log.Println("开始挖矿....")
	for {
		tmp := [][]byte{
			util.Uint64ToByte(b.Version),
			util.Uint64ToByte(b.TimeStamp),
			util.Uint64ToByte(b.Difficulty),
			util.Uint64ToByte(b.Nonce),
			b.PrevHash,
			b.MerkelRoot,
			b.Data,
		}
		hash = sha256.Sum256(bytes.Join(tmp, []byte{}))

		bi := big.Int{}
		bi.SetBytes(hash[:])
		if bi.Cmp(p.target) < 0 {
			break
		}
		nonce++
	}

	return nonce, hash[:]
}

func (p *ProofOfWork) isVoid() bool {

	return false
}
