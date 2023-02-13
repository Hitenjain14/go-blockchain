package network

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	tra := NewLocalTransport("A")
	trb := NewLocalTransport("B")

	tra.Connect(trb)
	trb.Connect(tra)
	assert.Equal(t, tra.(*LocalTransport).peers[trb.Addr()], trb)
	assert.Equal(t, trb.(*LocalTransport).peers[tra.Addr()], tra)

	msg := []byte("hello blockchain")

	assert.Nil(t, tra.SendMessage(trb.Addr(), msg))

	rpc := <-trb.Consume()
	assert.Equal(t, rpc.From, tra.Addr())
	buf := make([]byte, len(msg))
	n, err := rpc.Payload.Read(buf)
	assert.Nil(t, err)
	assert.Equal(t, n, len(msg))
	assert.Equal(t, buf, msg)
}
