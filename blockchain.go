package blockchain
import (
	"bytes"
	"crypto/sha256"
	
)
type block struct {
	hash    []byte
	data    []byte
	prehash []byte
}

func (b *block) Drivehash() {
	info := bytes.Join([][]byte{b.data, b.prehash}, []byte{})
	hash := sha256.Sum256(info)
	b.hash = hash[:]
}
func Createblock(data string, prehash []byte) *block {
	b := block{[]byte{}, []byte(data), prehash}
	b.Drivehash()
	return &b
}

type Blockchain struct {
	Blocks []*block
}

func (bc *Blockchain) Addblock(data string) {
	preblock := bc.Blocks[len(bc.Blocks)-1]
	new := Createblock(data, preblock.hash)
	bc.Blocks = append(bc.Blocks, new)
}
func Genesisblock() *block {
	return Createblock("ginesis", []byte{})
}
func Initchain() *Blockchain {
	return &Blockchain{[]*block{Genesisblock()}}
}