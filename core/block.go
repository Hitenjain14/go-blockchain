package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"time"

	"github.com/hitenjain14/go-blockchain/crypto"
	"github.com/hitenjain14/go-blockchain/types"
)

type Header struct {
	Version       uint32
	PrevBlockHash types.Hash
	DataHash      types.Hash
	Timestamp     int64
	Height        uint32
}

type Block struct {
	*Header
	Transactions []*Transaction
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

	dataHash, err := CalculateDataHash(b.Transactions)
	if err != nil {
		return err
	}
	if dataHash != b.DataHash {
		return fmt.Errorf("block (%s) has invalid data hash", b.Hash(BlockHasher{}))
	}

	return nil
}

func NewBlock(header *Header, txx []*Transaction) (*Block, error) {

	return &Block{
		Header:       header,
		Transactions: txx,
	}, nil

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
	b.Transactions = append(b.Transactions, tx)
}

func CalculateDataHash(txx []*Transaction) (hash types.Hash, err error) {

	buf := &bytes.Buffer{}

	for _, tx := range txx {
		err = tx.Encode(NewGobTxEncoder(buf))
		if err != nil {
			return
		}
	}

	hash = sha256.Sum256(buf.Bytes())

	return
}

func NewBlockFromPrevHeader(prevHeader *Header, txx []*Transaction) (*Block, error) {
	dataHash, err := CalculateDataHash(txx)
	if err != nil {
		return nil, err
	}

	header := &Header{
		Version:       1,
		DataHash:      dataHash,
		PrevBlockHash: BlockHasher{}.Hash(prevHeader),
		Timestamp:     time.Now().UnixNano(),
		Height:        prevHeader.Height + 1,
	}

	return NewBlock(header, txx)

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
