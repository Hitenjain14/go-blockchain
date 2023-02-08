package core

import (
	"fmt"
	"sync"

	"github.com/sirupsen/logrus"
)

type Blockchain struct {
	store     Storage
	lock      sync.RWMutex
	headers   []*Header
	validator Validator
}

func NewBlockchain(genesis *Block) (*Blockchain, error) {
	bc := &Blockchain{
		headers: []*Header{},
		store:   NewMemoryStore(),
	}
	bc.validator = NewBlockValidator(bc)
	err := bc.addBlockWithoutValidation(genesis)

	return bc, err
}

func (bc *Blockchain) SetStorage(s Storage) {
	bc.store = s
}

func (bc *Blockchain) SetValidator(v Validator) {
	bc.validator = v
}

func (bc *Blockchain) Height() uint32 {
	bc.lock.RLock()
	defer bc.lock.RUnlock()
	return uint32(len(bc.headers) - 1)
}

func (bc *Blockchain) AddBlock(b *Block) error {

	if err := bc.validator.ValidateBlock(b); err != nil {
		return err
	}

	return bc.addBlockWithoutValidation(b)

}

func (bc *Blockchain) HasBlock(height uint32) bool {
	return height <= bc.Height()
}

func (bc *Blockchain) addBlockWithoutValidation(b *Block) error {

	bc.lock.Lock()
	bc.headers = append(bc.headers, b.Header)
	bc.lock.Unlock()
	logrus.WithFields(logrus.Fields{
		"height": b.Height,
		"hash":   b.Hash(&BlockHasher{}),
	}).Infof("adding new block")
	return bc.store.Put(b)
}

func (bc *Blockchain) GetHeader(height uint32) (*Header, error) {
	if height <= bc.Height() {
		return bc.headers[int(height)], nil
	}
	bc.lock.Lock()
	defer bc.lock.Unlock()
	return nil, fmt.Errorf("block with %d height doesn't exist", height)
}
