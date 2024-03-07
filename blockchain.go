package blockchain
import (
	"bytes"
	"crypto/sha256"
	
)
type Block struct {
	Hash    []byte
	Data    []byte
	Prehash []byte
}

func (b *Block) Drivehash() {
	info := bytes.Join([][]byte{b.Data, b.Prehash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}
func Createblock(data string, prehash []byte) *Block {
	b := Block{[]byte{}, []byte(data), prehash}
	b.Drivehash()
	return &b
}

type Blockchain struct {
	Blocks []*Block
}

func (bc *Blockchain) Addblock(data string) {
	preblock := bc.Blocks[len(bc.Blocks)-1]
	new := Createblock(data, preblock.Hash)
	bc.Blocks = append(bc.Blocks, new)
}
func Genesisblock() *Block {
	return Createblock("ginesis", []byte{})
}
func Initchain() *Blockchain {
	return &Blockchain{[]*Block{Genesisblock()}}
}