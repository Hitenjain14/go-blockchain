package core

import "io"

type Decoder[T any] interface {
	Decode(io.Reader, T) error
}

type Encoder[T any] interface {
	Encode(io.Writer, T) error
}
