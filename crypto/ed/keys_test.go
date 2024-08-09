package ed

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateKeys(t *testing.T) {

	pub, pri := GenerateKeys(nil)

	assert.Equal(t, []byte(pri.key), pri.Bytes())

	t.Log("pub key ", pub.String(), len(pub.Bytes()))
	t.Log("pri key ", pri.String(), len(pri.Bytes()))
	t.Log("add ", pub.publicKeyToAddress())
	t.Log(pub.sha3Addr())

}
