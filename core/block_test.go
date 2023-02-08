package core

import (
	"testing"
	"time"

	"github.com/hitenjain14/go-blockchain/crypto"
	"github.com/hitenjain14/go-blockchain/types"
	"github.com/stretchr/testify/assert"
)

func randomBlock(height uint32) *Block {
	header := &Header{
		Version:   1,
		PrevBlock: types.RandomHash(),
		Timestamp: time.Now().UnixNano(),
		Height:    height,
	}

	txx := Transaction{
		Data: []byte("Hello"),
	}

	return NewBlock(header, []Transaction{txx})

}

func randomBlockWithSignature(t *testing.T, height uint32) *Block {
	b := randomBlock(height)
	privKey := crypto.GeneratePrivateKey()
	err := b.Sign(privKey)
	assert.Nil(t, err)
	return b
}

func TestSignBlock(t *testing.T) {

	b := randomBlock(1)

	privKey := crypto.GeneratePrivateKey()

	assert.Nil(t, b.Sign(privKey))

	assert.Nil(t, b.Verify())

	b.Height = 2

	assert.NotNil(t, b.Verify())

	b.Height = 1
	otherPrivKey := crypto.GeneratePrivateKey()
	otherPubKey := otherPrivKey.PublicKey()

	b.Validator = otherPubKey
	assert.NotNil(t, b.Verify())

}
