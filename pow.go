package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"math"
	"math/big"

	
	
)

const difficulty = 15

type Pow struct {
	Block  *Block
	Target *big.Int
}

func Newpow(b *Block) *Pow {
	target := big.NewInt(1)
	target.Lsh(target,256-difficulty)
	pow := &Pow{b, target}
	return pow
}
func (P *Pow) Initdata(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			P.Block.Data,
			P.Block.Prehash,
			Tohex(int64(nonce)),
			Tohex(int64(difficulty)),
		},
		[]byte{},
	)
	return data
}
func Tohex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
func (P *Pow) Run() (int, []byte) {
	var inthash big.Int
	var hash [32]byte
	nonce := 0
	for nonce < math.MaxInt64 {
		data := P.Initdata(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("%x",hash)
		inthash.SetBytes(hash[:])
		if inthash.Cmp(P.Target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Printf("%d\n%x", nonce, hash)
	return nonce, hash[:]
}
func (P *Pow) Validate() bool {
	var inthash big.Int
	data := P.Initdata(P.Block.Nonce)
	hash := sha256.Sum256(data)
	inthash.SetBytes(hash[:])
	return inthash.Cmp(P.Target)==-1
}
