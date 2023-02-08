package network

import (
	"sort"

	"github.com/hitenjain14/go-blockchain/core"
	"github.com/hitenjain14/go-blockchain/types"
)

type TxMapSorter struct {
}

type Sorter[T []*core.Transaction] interface {
	FinalSort([]*core.Transaction)
}

func (s *TxMapSorter) FinalSort(txx []*core.Transaction) {
	// sort.Sort(ByFirstSeen(txx))
	sort.Slice(txx, func(i, j int) bool {
		return txx[i].FirstSeen() < txx[j].FirstSeen()
	})
}

// func NewTxMapSorter(txMap map[types.Hash]*core.Transaction) *TxMapSorter {

// 	txx := make([]*core.Transaction, len(txMap))

// 	i := 0

// 	for _, val := range txMap {
// 		txx[i] = val
// 		i++
// 	}

// 	s := &TxMapSorter{
// 		transactions: txx,
// 	}

// 	sort.Sort(s)

// 	return s
// }

type TxPool struct {
	transactions map[types.Hash]*core.Transaction
}

func NewTxPool() *TxPool {
	return &TxPool{
		transactions: make(map[types.Hash]*core.Transaction),
	}
}

func (p *TxPool) Transactions(sorter Sorter[[]*core.Transaction]) []*core.Transaction {

	txMap := p.transactions

	txx := make([]*core.Transaction, len(txMap))

	i := 0

	for _, val := range txMap {
		txx[i] = val
		i++
	}

	sorter.FinalSort(txx)
	return txx

}

func (p *TxPool) Len() int {
	return len(p.transactions)
}

func (p *TxPool) Flush() {
	p.transactions = make(map[types.Hash]*core.Transaction)
}

func (p *TxPool) Add(tx *core.Transaction) bool {
	hash := tx.Hash(core.TxHasher{})
	if p.Has(hash) {
		return false
	}
	p.transactions[hash] = tx
	return true
}

func (p *TxPool) Has(hash types.Hash) bool {
	_, ok := p.transactions[hash]
	return ok
}
