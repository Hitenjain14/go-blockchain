package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSignSuccess(t *testing.T) {

	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	// address := pubKey.Address()

	msg := []byte("Sign Message")

	sig, err := privKey.Sign(msg)

	assert.Nil(t, err)
	assert.True(t, sig.Verify(pubKey, msg))

	// fmt.Printf("%+v", sig)

}

func TestSignFail(t *testing.T) {

	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()

	// address := pubKey.Address()

	msg := []byte("Sign Message")

	othPrivKey := GeneratePrivateKey()
	othPubKey := othPrivKey.PublicKey()

	sig, err := privKey.Sign(msg)

	assert.Nil(t, err)
	assert.False(t, sig.Verify(othPubKey, msg))
	assert.False(t, sig.Verify(pubKey, []byte("Other Message")))

	// fmt.Printf("%+v", sig)

}
