package blockchain

type Block struct {
	Hash    []byte
	Data    []byte
	Prehash []byte
	Nonce	int
}


func Createblock(data string, prehash []byte) *Block {
	b := &Block{[]byte{}, []byte(data), prehash,0}
	pow:=Newpow(b)
	nonce,hash:=pow.Run()
	b.Nonce=nonce
	b.Hash=hash[:]
	return b
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