package core

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/hitenjain14/go-blockchain/crypto"
	"github.com/hitenjain14/go-blockchain/types"
)

type Header struct {
	Version       uint32
	PrevBlockHash types.Hash
	Timestamp     int64
	Height        uint32
	Nonce         uint64
}

type Block struct {
	*Header
	Transactions []Transaction
	Validator    crypto.PublicKey
	Signature    *crypto.Signature
	hash         types.Hash //header hash cached

}

func (b *Block) Sign(privKey crypto.PrivateKey) error {
	sig, err := privKey.Sign(b.Hash(&BlockHasher{}).ToSlice())
	if err != nil {
		return err
	}
	b.Signature = sig
	b.Validator = privKey.PublicKey()

	return nil
}

func (b *Block) Verify() error {
	if b.Signature == nil {
		return fmt.Errorf("block is not signed")
	}
	if !b.Signature.Verify(b.Validator, b.Hash(&BlockHasher{}).ToSlice()) {
		return fmt.Errorf("invalid block signature")
	}

	for _, tx := range b.Transactions {
		if err := tx.Verify(); err != nil {
			return err
		}
	}

	return nil
}

func NewBlock(header *Header, txx []Transaction) *Block {

	return &Block{
		Header:       header,
		Transactions: txx,
	}

}

func (b *Block) Hash(hasher Hasher[*Header]) types.Hash {

	b.hash = hasher.Hash(b.Header)

	return b.hash
}

func (b *Block) Decode(dec Decoder[*Block]) error {
	return dec.Decode(b)
}

func (b *Block) Encode(enc Encoder[*Block]) error {
	return enc.Encode(b)
}

func (h *Header) Bytes() []byte {
	buf := &bytes.Buffer{}
	enc := gob.NewEncoder(buf)
	err := enc.Encode(h)
	if err != nil {
		panic(err)
	}

	return buf.Bytes()
}

func (b *Block) AddTransaction(tx *Transaction) {
	b.Transactions = append(b.Transactions, *tx)
}

// func (h *Header) EncodeBinary(w io.Writer) error {
// 	if err := binary.Write(w, binary.LittleEndian, &h.Version); err != nil {
// 		return err
// 	}
// 	if err := binary.Write(w, binary.LittleEndian, &h.PrevBlock); err != nil {
// 		return err
// 	}
// 	if err := binary.Write(w, binary.LittleEndian, &h.Timestamp); err != nil {
// 		return err
// 	}
// 	if err := binary.Write(w, binary.LittleEndian, &h.Height); err != nil {
// 		return err
// 	}
// 	if err := binary.Write(w, binary.LittleEndian, &h.Nonce); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (h *Header) DecodeBinary(r io.Reader) error {
// 	if err := binary.Read(r, binary.LittleEndian, &h.Version); err != nil {
// 		return err
// 	}
// 	if err := binary.Read(r, binary.LittleEndian, &h.PrevBlock); err != nil {
// 		return err
// 	}
// 	if err := binary.Read(r, binary.LittleEndian, &h.Timestamp); err != nil {
// 		return err
// 	}
// 	if err := binary.Read(r, binary.LittleEndian, &h.Height); err != nil {
// 		return err
// 	}
// 	if err := binary.Read(r, binary.LittleEndian, &h.Nonce); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (b *Block) EncodeBinary(w io.Writer) error {
// 	if err := b.Header.EncodeBinary(w); err != nil {
// 		return err
// 	}

// 	for _, tx := range b.Transactions {
// 		if err := tx.EncodeBinary(w); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (b *Block) DecodeBinary(r io.Reader) error {
// 	if err := b.Header.DecodeBinary(r); err != nil {
// 		return err
// 	}

// 	for _, tx := range b.Transactions {
// 		if err := tx.DecodeBinary(r); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }
