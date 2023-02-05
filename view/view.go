package view

import (
	"context"
	"time"

	"github.com/ionos-cloud/streams"
	"github.com/ionos-cloud/streams/store"

	"github.com/katallaxie/pkg/server"
	"github.com/katallaxie/pkg/utils"
)

const (
	// Pure ...
	Pure = iota

	// Initializing ...
	Initializing

	// Running ...
	Running
)

// NextCursor ...
type NextCursor struct {
	Key   string
	Value []byte
}

// Iterator ...
type Iterator interface {
	// Next ...
	Next() <-chan NextCursor
}

// Table ...
type Table interface {
	// Set ...
	Set(key string, value []byte) error

	// Delete ...
	Delete(key string) error

	Iterator
}

// Unimplemented ...
type Unimplemented struct{}

// Next ...
func (u *Unimplemented) Next() (string, string) {
	return "", ""
}

// View ...
type View[V any] interface {
	// Get ...
	Get(key string) (V, error)

	// Set ...
	Set(key string, value V) error

	// Delete ...
	Delete(key string) error
}

type view[V any] struct {
	store store.Storage
	table Table

	encoder streams.Encoder[V]
	decoder streams.Decoder[V]

	server.Listener
}

// New ..
func New[V any](table Table, encoder streams.Encoder[V], decoder streams.Decoder[V], store store.Storage) View[V] {
	v := new(view[V])
	v.table = table
	v.store = store

	v.encoder = encoder
	v.decoder = decoder

	return v
}

// Get ...
func (v *view[V]) Get(key string) (V, error) {
	b, err := v.store.Get(key)
	if err != nil {
		return utils.Zero[V](), err
	}

	value, err := v.decoder.Decode(b)
	if err != nil {
		return utils.Zero[V](), err
	}

	return value, nil
}

// Set ...
func (v *view[V]) Set(key string, value V) error {
	b, err := v.encoder.Encode(value)
	if err != nil {
		return nil
	}

	// TODO: this is non optimistic, the update is published to the table and then synced to storage
	err = v.table.Set(key, b)
	if err != nil {
		return err
	}

	return nil
}

// Delete ...
func (v *view[V]) Delete(key string) error {
	err := v.table.Delete(key) // Tombstone message
	if err != nil {
		return err
	}

	return nil
}

// Start ...
func (v *view[V]) Start(ctx context.Context, ready server.ReadyFunc, run server.RunFunc) func() error {
	return func() error {
		ready()

		ticker := time.NewTicker(1 * time.Second)

		for {
			select {
			case <-ticker.C:
			case <-ctx.Done():
				return nil
			case c := <-v.table.Next():
				if c.Value == nil {
					ok, err := v.store.Has(c.Key)
					if err != nil {
						return err
					}

					if !ok {
						continue
					}

					err = v.store.Delete(c.Key)
					if err != nil {
						return err
					}

					continue
				}

				err := v.store.Set(c.Key, c.Value)
				if err != nil {
					return err
				}

				continue
			}
		}
	}
}