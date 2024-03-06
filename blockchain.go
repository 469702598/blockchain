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

func (b *block) drivehash() {
	info := bytes.Join([][]byte{b.data, b.prehash}, []byte{})
	hash := sha256.Sum256(info)
	b.hash = hash[:]
}
func createblock(data string, prehash []byte) *block {
	b := block{[]byte{}, []byte(data), prehash}
	b.drivehash()
	return &b
}

type blockchain struct {
	Blocks []*block
}

func (bc *blockchain) addblock(data string) {
	preblock := bc.Blocks[len(bc.Blocks)-1]
	new := createblock(data, preblock.hash)
	bc.Blocks = append(bc.Blocks, new)
}
func genesisblock() *block {
	return createblock("ginesis", []byte{})
}
func Initchain() *blockchain {
	return &blockchain{[]*block{genesisblock()}}
}