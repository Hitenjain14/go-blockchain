package core

import (
	"testing"

	"github.com/hitenjain14/go-blockchain/crypto"
	"github.com/stretchr/testify/assert"
)

func TestSignTransaction(t *testing.T) {

	data := []byte("Hello")

	tx := &Transaction{
		Data: data,
	}

	privKey := crypto.GeneratePrivateKey()

	err := tx.Sign(privKey)
	assert.Nil(t, err)
	assert.NotNil(t, tx.Signature)

}

func TestVerifyTransaction(t *testing.T) {

	data := []byte("Hello")

	tx := &Transaction{
		Data: data,
	}

	privKey := crypto.GeneratePrivateKey()

	err := tx.Sign(privKey)
	assert.Nil(t, err)
	assert.NotNil(t, tx.Signature)

	err = tx.Verify()
	assert.Nil(t, err)

	otherPrivKey := crypto.GeneratePrivateKey()
	otherPubKey := otherPrivKey.PublicKey()

	tx.PublicKey = otherPubKey
	assert.NotNil(t, tx.Verify())

}
