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

func NewGobTxEncoder(w io.Writer) *GobTxEncoder {
	gob.Register(elliptic.P256())
	return &GobTxEncoder{w: w}
}

func NewGobTxDecoder(r io.Reader) *GobTxDecoder {
	gob.Register(elliptic.P256())
	return &GobTxDecoder{r: r}
}

func (e *GobTxEncoder) Encode(tx *Transaction) error {

	return gob.NewEncoder(e.w).Encode(tx)
}

func (d *GobTxDecoder) Decode(tx *Transaction) error {

	return gob.NewDecoder(d.r).Decode(tx)
}
