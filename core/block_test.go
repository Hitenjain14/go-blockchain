package core

import (
	"bytes"
	"testing"
	"time"

	"github.com/hitenjain14/go-blockchain/crypto"
	"github.com/hitenjain14/go-blockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestSignBlock(t *testing.T) {

	b := randomBlock(t, 1, types.Hash{})

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

func TestEncodeDecodeBlock(t *testing.T) {
	b := randomBlock(t, 1, types.Hash{})
	buf := &bytes.Buffer{}
	assert.Nil(t, b.Encode(NewGobBlockEncoder(buf)))

	bDecode := new(Block)
	assert.Nil(t, bDecode.Decode(NewGobBlockDecoder(buf)))

	assert.Equal(t, b.Header, bDecode.Header)
}

func randomBlock(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {

	privKey := crypto.GeneratePrivateKey()

	header := &Header{
		Version:       1,
		PrevBlockHash: prevBlockHash,
		Timestamp:     time.Now().UnixNano(),
		Height:        height,
	}

	txx := randomSignedTransaction(t)

	b, err := NewBlock(header, []*Transaction{txx})
	assert.Nil(t, err)
	dataHash, err := CalculateDataHash(b.Transactions)
	assert.Nil(t, err)
	b.DataHash = dataHash
	assert.Nil(t, b.Sign(privKey))
	return b
}

// func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
// 	b := randomBlock(t, height, prevBlockHash)
// 	privKey := crypto.GeneratePrivateKey()
// 	err := b.Sign(privKey)
// 	assert.Nil(t, err)
// 	return b
// }
