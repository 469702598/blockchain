package blockchain
import (
	"bytes"
	"crypto/sha256"
	
)
type Block struct {
	hash    []byte
	data    []byte
	prehash []byte
}

func (b *Block) Drivehash() {
	info := bytes.Join([][]byte{b.data, b.prehash}, []byte{})
	hash := sha256.Sum256(info)
	b.hash = hash[:]
}
func Createblock(data string, prehash []byte) *Block {
	b := Block{[]byte{}, []byte(data), prehash}
	b.Drivehash()
	return &b
}

type Blockchain struct {
	blocks []*Block
}

func (bc *Blockchain) Addblock(data string) {
	preblock := bc.blocks[len(bc.blocks)-1]
	new := Createblock(data, preblock.hash)
	bc.blocks = append(bc.blocks, new)
}
func Genesisblock() *Block {
	return Createblock("ginesis", []byte{})
}
func Initchain() *Blockchain {
	return &Blockchain{[]*Block{Genesisblock()}}
}