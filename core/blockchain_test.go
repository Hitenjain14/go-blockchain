package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newBlockchainWithGenesis(t *testing.T) *Blockchain {
	bc, err := NewBlockchain(randomBlock(0))
	assert.Nil(t, err)
	assert.NotNil(t, bc.validator)
	return bc
}

func TestNewBlockchain(t *testing.T) {
	bc, err := NewBlockchain(randomBlock(0))
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

	b := randomBlockWithSignature(t, 1)
	assert.Nil(t, bc.AddBlock(b))
	assert.Equal(t, bc.Height(), uint32(1))
	assert.True(t, bc.HasBlock(1))
	assert.NotNil(t, bc.AddBlock(randomBlockWithSignature(t, 1)))

}
