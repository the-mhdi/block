package types

import (
	"encoding/hex"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	head := new(Header)
	faker.FakeData(&head)
	h := head.Hash()
	t.Log(hex.EncodeToString(h))
	//t.Log(head)

}

func TestSign(t *testing.T) {
	head := new(Header)
	faker.FakeData(&head)
	h := head.Hash()
	block := new(Block)
	block.header = head
	faker.FakeData(&block)
	s := block.Sign
	assert.Equal(t, block.headHash, h)

}
