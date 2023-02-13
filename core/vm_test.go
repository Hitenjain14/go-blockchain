package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVM(t *testing.T) {

	data := []byte{0x01, 0x0a, 0x02, 0x0a, 0x0b}

	vm := NewVM(data)
	err := vm.Run()
	assert.Nil(t, err)
}
