package core

import (
	"crypto/elliptic"
	"encoding/gob"
	"io"
)

type Decoder[T any] interface {
	Decode(T) error
}

type Encoder[T any] interface {
	Encode(T) error
}

type GobTxEncoder struct {
	w io.Writer
}

type GobTxDecoder struct {
	r io.Reader
}

type GobBlockEncoder struct {
	w io.Writer
}

type GobBlockDecoder struct {
	r io.Reader
}

func NewGobBlockEncoder(w io.Writer) *GobBlockEncoder {

	return &GobBlockEncoder{w: w}
}

func NewGobBlockDecoder(r io.Reader) *GobBlockDecoder {

	return &GobBlockDecoder{r: r}
}

func NewGobTxEncoder(w io.Writer) *GobTxEncoder {

	return &GobTxEncoder{w: w}
}

func NewGobTxDecoder(r io.Reader) *GobTxDecoder {

	return &GobTxDecoder{r: r}
}

func (e *GobTxEncoder) Encode(tx *Transaction) error {

	return gob.NewEncoder(e.w).Encode(tx)
}

func (d *GobTxDecoder) Decode(tx *Transaction) error {

	return gob.NewDecoder(d.r).Decode(tx)
}

func (e *GobBlockEncoder) Encode(b *Block) error {

	return gob.NewEncoder(e.w).Encode(b)
}

func (d *GobBlockDecoder) Decode(b *Block) error {

	return gob.NewDecoder(d.r).Decode(b)
}

func init() {
	gob.Register(elliptic.P256())
}
