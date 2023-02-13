package core

import (
	"fmt"
)

type Instruction byte

const (
	InstrPush Instruction = 0x0a
	InstrAdd  Instruction = 0x0b
)

type VM struct {
	data  []byte
	ip    int // instruction pointer
	stack []byte
	sp    int // stack pointer
}

func NewVM(data []byte) *VM {
	return &VM{
		data:  data,
		ip:    0,
		stack: make([]byte, 1024),
		sp:    -1,
	}
}

func (vm *VM) Run() error {

	for {
		instr := vm.data[vm.ip]
		fmt.Println(instr)
		vm.ip++
		if vm.ip >= len(vm.data) {
			break
		}
	}
	return nil
}
