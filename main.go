package main

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hitenjain14/go-blockchain/core"
	"github.com/hitenjain14/go-blockchain/crypto"
	"github.com/hitenjain14/go-blockchain/network"
	"github.com/sirupsen/logrus"
)

func main() {

	trLocal := network.NewLocalTransport(network.NetAddr("local"))
	trRemoteA := network.NewLocalTransport(network.NetAddr("remote_a"))
	trRemoteB := network.NewLocalTransport(network.NetAddr("remote_b"))
	trRemoteC := network.NewLocalTransport(network.NetAddr("remote_c"))

	trLocal.Connect(trRemoteA)
	trRemoteA.Connect(trRemoteB)
	trRemoteB.Connect(trRemoteC)
	trRemoteA.Connect(trLocal)

	initRemoteServers([]network.Transport{trRemoteA, trRemoteB, trRemoteC})
	go func() {
		for {
			if err := sendTransaction(trRemoteA, trLocal.Addr()); err != nil {
				logrus.Error(err)
			}
			time.Sleep(2 * time.Second)
		}
	}()

	privKey := crypto.GeneratePrivateKey()

	localServer := makeServer("local", trLocal, &privKey)
	localServer.Start()
}

func makeServer(id string, tr network.Transport, pk *crypto.PrivateKey) *network.Server {
	opts := network.ServerOpts{
		PrivateKey: pk,
		ID:         id,
		Transports: []network.Transport{tr},
	}

	s, err := network.NewServer(opts)
	if err != nil {
		log.Fatal(err)
	}
	return s
}

func initRemoteServers(trs []network.Transport) {
	for i := 0; i < len(trs); i++ {
		id := fmt.Sprintf("remote_%d", i)
		s := makeServer(id, trs[i], nil)
		go s.Start()
	}
}

func sendTransaction(tr network.Transport, to network.NetAddr) error {
	privKey := crypto.GeneratePrivateKey()
	data := []byte(strconv.FormatInt(time.Now().UnixNano(), 10))
	tx := core.NewTransaction([]byte(data))
	if err := tx.Sign(privKey); err != nil {
		return err
	}
	buf := &bytes.Buffer{}

	if err := tx.Encode(core.NewGobTxEncoder(buf)); err != nil {
		return err
	}

	msg := network.NewMessage(network.MessageTypeTx, buf.Bytes())

	return tr.SendMessage(to, msg.Bytes())

}
