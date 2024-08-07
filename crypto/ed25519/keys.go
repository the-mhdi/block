package ed

import (
	"crypto/ed25519"
	"encoding/hex"
	"io"

	"golang.org/x/crypto/sha3"
)

type Key interface {
	Type() bool //0 (false) private key - 1 public
	String() string
	Bytes() []byte
}

// wrapper around ed25519 privatekey
type PrivateKey struct {
	key ed25519.PrivateKey
}

type PublicKey struct {
	key ed25519.PublicKey
}

func GenerateKeys(rand io.Reader) (*PublicKey, *PrivateKey) {
	pub, pri, err := ed25519.GenerateKey(rand)

	if err != nil {
		panic(err)
	}

	return &PublicKey{key: pub}, &PrivateKey{key: pri}

}

// returns hex string of private key
func (p *PrivateKey) String() string {
	return hex.EncodeToString(p.key)
}

func (p *PrivateKey) Bytes() []byte {
	return p.key
}

func (p *PrivateKey) Type() bool {
	return false
}

func (p *PrivateKey) Sign(msg []byte) Signature {
	return ed25519.Sign(p.key, msg)
}

type Signature []byte

func (s *Signature) Verify(pub *PublicKey, msg []byte) bool {
	return ed25519.Verify(pub.key, msg, *s)
}

// //////////////////////////////////////////////////////////////////////////
func (p *PublicKey) String() string {
	return hex.EncodeToString(p.key)
}

func (p *PublicKey) Bytes() []byte {
	return p.key
}

func (p *PublicKey) Type() bool {
	return true
}

func (p *PublicKey) PublicKeyToAddress() []byte {

	address, _ := p.pubToAddr()

	return address
}

func (p *PublicKey) pubToAddr() ([]byte, error) {
	hash := sha3.New256()
	_, err := hash.Write(p.key)
	if err != nil {
		return nil, err
	}

	return hash.Sum([]byte("0x")), nil
}
