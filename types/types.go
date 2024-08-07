package types

import (
	"crypto/sha256"
	"math/big"
	"time"

	"github.com/the-mhdi/block/crypto/ed25519/ed"
)

const (
	// HashLength is the expected length of the hash
	HashLength = 32
	// AddressLength is the expected length of the address
	AddressLength = 64
)

type Address [AddressLength]byte
type Hash [HashLength]byte

type Block struct {
	header   *Header
	headHash Hash

	transactions Transaction

	ReceivedAt time.Time

	ReceivedFrom interface{}
}

type Header struct {
	ParentHash Hash
	Coinbase   Address
	Root       Hash //state trie root
	TxHash     Hash //merkle tree of txhash of all the transactions in the block
	Time       uint64
	Number     *big.Int //block number
	GasLimit   uint64
	GasUsed    uint64
}

type Transaction struct {
	data []byte
	time time.Time
}

func (b *Block) Sign(prv *ed.PrivateKey) ed.Signature {
	b.headHash = b.header.Hash()
	return prv.Sign(b.headHash)
}

func (h *Header) Hash() []byte {
	hash := sha256.New()
	hash.Write(h.ParentHash[:])
	hash.Write(h.Coinbase[:])
	hash.Write(h.Root[:])
	hash.Write(h.TxHash[:])

	return hash.Sum(nil)
}
