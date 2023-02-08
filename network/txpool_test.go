package network

import (
	"testing"

	"github.com/hitenjain14/go-blockchain/core"
	"github.com/stretchr/testify/assert"
)

func TestTxPool(t *testing.T) {
	p := NewTxPool()
	assert.Equal(t, p.Len(), 0)
}

func TestTxPoolAddTransaction(t *testing.T) {
	p := NewTxPool()
	tx := core.NewTransaction([]byte("hello"))
	p.Add(tx)
	assert.Equal(t, p.Len(), 1)
	p.Flush()
	assert.Equal(t, p.Len(), 0)
}

func TestSort(t *testing.T) {

	p := NewTxPool()
	tx := core.NewTransaction([]byte("hello"))
	tx.SetFirstSeen(2)
	txx := core.NewTransaction([]byte("good"))
	txx.SetFirstSeen(1)
	p.Add(tx)
	p.Add(txx)

	assert.Equal(t, p.Transactions(&TxMapSorter{})[0].FirstSeen(), int64(1))

}
