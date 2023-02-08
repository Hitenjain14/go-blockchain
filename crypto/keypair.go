package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/hitenjain14/go-blockchain/types"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

type PublicKey struct {
	key *ecdsa.PublicKey
}

type Signature struct {
	r, s *big.Int
}

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	if err != nil {
		panic(err)
	}

	return PrivateKey{key: key}
}

func (k PrivateKey) PublicKey() PublicKey {
	return PublicKey{key: &k.key.PublicKey}
}

func (k PrivateKey) Sign(data []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, k.key, data)

	if err != nil {
		return nil, err
	}
	return &Signature{r: r, s: s}, nil
}

func (k PublicKey) ToSlice() []byte {
	return elliptic.Marshal(k.key.Curve, k.key.X, k.key.Y)
}

func (k PublicKey) Address() types.Address {

	h := sha256.Sum256(k.ToSlice())

	return types.AddressFromBytes(h[len(h)-20:])
}

func (sig *Signature) Verify(pub PublicKey, data []byte) bool {
	return ecdsa.Verify(pub.key, data, sig.r, sig.s)

}