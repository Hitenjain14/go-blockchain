package core

import (
	"fmt"
)

type Validator interface {
	ValidateBlock(*Block) error
}

type BlockValidator struct {
	bc *Blockchain
}

func NewBlockValidator(bc *Blockchain) *BlockValidator {
	return &BlockValidator{bc: bc}
}

func (v *BlockValidator) ValidateBlock(b *Block) error {

	if v.bc.HasBlock(b.Height) {
		return fmt.Errorf("block with %d height already exists in chain with hash %s", b.Height, b.Hash(&BlockHasher{}))
	}

	if b.Height != v.bc.Height()+1 {
		return fmt.Errorf("block with %d height can't be added to chain with height %d", b.Height, v.bc.Height())
	}

	prevHeader, err := v.bc.GetHeader(b.Height - 1)

	hash := BlockHasher{}.Hash(prevHeader)

	if hash != b.PrevBlockHash {
		return fmt.Errorf("block with %d height has invalid previous block hash %s", b.Height, b.PrevBlockHash)
	}

	if err != nil {
		return err
	}

	if err = b.Verify(); err != nil {
		return err
	}

	return nil
}

// [0,1,2,3]
// 4 the height
