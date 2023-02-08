package core

import (
	"testing"

	"github.com/hitenjain14/go-blockchain/types"
	"github.com/stretchr/testify/assert"
)

func TestNewBlockchain(t *testing.T) {
	bc, err := NewBlockchain(randomBlock(t, 0, types.Hash{}))
	assert.Nil(t, err)
	assert.NotNil(t, bc.validator)
	assert.Equal(t, bc.Height(), uint32(0))
}

func TestHasBlock(t *testing.T) {
	bc := newBlockchainWithGenesis(t)
	assert.True(t, bc.HasBlock(0))
	assert.False(t, bc.HasBlock(1))
}

func TestAddBlock(t *testing.T) {

	bc := newBlockchainWithGenesis(t)

	b := randomBlockWithSignature(t, uint32(1), getPrevBlockHash(t, bc, uint32(1)))
	assert.Nil(t, bc.AddBlock(b))
	assert.Equal(t, bc.Height(), uint32(1))
	assert.True(t, bc.HasBlock(1))
	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, uint32(1), getPrevBlockHash(t, bc, uint32(1)))))

}

func TestAddBlockTooHigh(t *testing.T) {
	bc := newBlockchainWithGenesis(t)

	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 2, types.Hash{})))
}

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(t, 0, types.Hash{}))
	assert.Nil(t, err)
	assert.NotNil(t, bc.validator)
	return bc
}

func getPrevBlockHash(t *testing.T, bc *Blockchain, height uint32) types.Hash {

	header, err := bc.GetHeader(height - 1)
	assert.Nil(t, err)

	return BlockHasher{}.Hash(header)

}
