// SPDX-FileCopyrightText: 2023 PK Lab AG <contact@pklab.io>
// SPDX-License-Identifier: MIT

package websocket

import (
	sdkSync "github.com/dipdup-net/indexer-sdk/pkg/sync"

	"github.com/pkg/errors"
)

type processor[I, M any] func(data I) M

type Channel[I, M any] struct {
	clients   *sdkSync.Map[uint64, client]
	processor processor[I, M]
	filters   Filterable[M]
}

func NewChannel[I, M any](processor processor[I, M], filters Filterable[M]) *Channel[I, M] {
	return &Channel[I, M]{
		clients:   sdkSync.NewMap[uint64, client](),
		processor: processor,
		filters:   filters,
	}
}

func (channel *Channel[I, M]) AddClient(c client) {
	channel.clients.Set(c.Id(), c)
}

func (channel *Channel[I, M]) RemoveClient(id uint64) {
	channel.clients.Delete(id)
}

func (channel *Channel[I, M]) processMessage(msg I) error {
	if channel.clients.Len() == 0 {
		return nil
	}

	data := channel.processor(msg)

	if err := channel.clients.Range(func(_ uint64, value client) (error, bool) {
		if channel.filters.Filter(value, data) {
			value.Notify(data)
		}
		return nil, false
	}); err != nil {
		return errors.Wrap(err, "write message to client")
	}

	return nil
}
