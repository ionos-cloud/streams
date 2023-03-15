package store

import (
	"github.com/ionos-cloud/streams/codec"
	"github.com/ionos-cloud/streams/msg"
)

// Sink is a storage sink.
type Sink[V any] struct {
	Storage
	codec.Encoder[V]
}

// NewSink is a new storage sink.
func NewSink[K string, V any]() *Sink[V] {
	n := new(Sink[V])

	return n
}

// Write is a storage sink writer.
func (n *Sink[V]) Write(messages ...msg.Message[string, V]) error {
	for _, m := range messages {
		v, err := n.Encode(m.Value())
		if err != nil {
			return err
		}

		err = n.Set(m.Key(), v)
		if err != nil {
			return err
		}

		m.Mark()
	}

	return nil
}
