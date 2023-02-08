package core

import (
	"bytes"
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

	tx.From = otherPubKey
	assert.NotNil(t, tx.Verify())

}

func TestEncodeDecode(t *testing.T) {

	tx := randomSignedTransaction(t)

	buf := &bytes.Buffer{}

	assert.Nil(t, tx.Encode(NewGobTxEncoder(buf)))

	txDecoded := new(Transaction)

	assert.Nil(t, txDecoded.Decode(NewGobTxDecoder(buf)))
	assert.Equal(t, tx, txDecoded)

}

func randomSignedTransaction(t *testing.T) *Transaction {
	tx := &Transaction{
		Data: []byte("Hello"),
	}
	privKey := crypto.GeneratePrivateKey()
	err := tx.Sign(privKey)
	assert.Nil(t, err)
	return tx
}
