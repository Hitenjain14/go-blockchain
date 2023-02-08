package network

import (
	"fmt"
	"time"

	"github.com/hitenjain14/go-blockchain/core"
	"github.com/hitenjain14/go-blockchain/crypto"
	"github.com/sirupsen/logrus"
)

const defaultBlockTime = 5 * time.Second

type ServerOpts struct {
	Transports []Transport
	PrivateKey *crypto.PrivateKey
	BlockTime  time.Duration
}

type Server struct {
	ServerOpts
	blockTime   time.Duration
	isValidator bool
	memPool     *TxPool
	rpcCh       chan RPC
	quitCh      chan struct{}
}

func NewServer(opts ServerOpts) *Server {

	if opts.BlockTime == time.Duration(0) {
		opts.BlockTime = defaultBlockTime
	}

	return &Server{
		ServerOpts:  opts,
		blockTime:   opts.BlockTime,
		isValidator: opts.PrivateKey != nil,
		memPool:     NewTxPool(),
		rpcCh:       make(chan RPC),
		quitCh:      make(chan struct{}, 1),
	}
}

func (s *Server) Start() {
	s.initTransport()
	ticker := time.NewTicker(s.blockTime)

free:
	for {
		select {
		case rpc := <-s.rpcCh:
			fmt.Printf("%+v \n", rpc)
		case <-s.quitCh:
			break free
		case <-ticker.C:
			if s.isValidator {
				s.createNewBlock()
			}
		}
	}

	fmt.Println("server stopped")

}

func (s *Server) addTransaction(tx *core.Transaction) error {
	if err := tx.Verify(); err != nil {
		return err
	}

	hash := tx.Hash(core.TxHasher{})

	if res := s.memPool.Add(tx); res {
		logrus.WithFields(logrus.Fields{
			"hash": hash,
		}).Info("new transaction added to mempool")
	} else {
		logrus.WithFields(logrus.Fields{
			"hash": hash,
		}).Info("transaction already exists in mempool")
	}

	return nil
}

func (s *Server) createNewBlock() error {
	fmt.Println("creating a new block")
	return nil
}

func (s *Server) initTransport() {

	for _, tr := range s.Transports {
		go func(tr Transport) {
			for rpc := range tr.Consume() {
				s.rpcCh <- rpc
			}
		}(tr)
	}

}
