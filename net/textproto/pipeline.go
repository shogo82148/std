// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package textproto

import (
	"github.com/shogo82148/std/sync"
)

// A Pipeline manages a pipelined in-order request/response sequence.
//
// To use a Pipeline p to manage multiple clients on a connection,
// each client should run:
//
//	id := p.Next()	// take a number
//
//	p.StartRequest(id)	// wait for turn to send request
//	«send request»
//	p.EndRequest(id)	// notify Pipeline that request is sent
//
//	p.StartResponse(id)	// wait for turn to read response
//	«read response»
//	p.EndResponse(id)	// notify Pipeline that response is read
//
// A pipelined server can use the same calls to ensure that
// responses computed in parallel are written in the correct order.
type Pipeline struct {
	mu       sync.Mutex
	id       uint
	request  sequencer
	response sequencer
}

// Next returns the next id for a request/response pair.
func (p *Pipeline) Next() uint

// StartRequest blocks until it is time to send (or, if this is a server, receive)
// the request with the given id.
func (p *Pipeline) StartRequest(id uint)

// EndRequest notifies p that the request with the given id has been sent
// (or, if this is a server, received).
func (p *Pipeline) EndRequest(id uint)

// StartResponse blocks until it is time to receive (or, if this is a server, send)
// the request with the given id.
func (p *Pipeline) StartResponse(id uint)

// EndResponse notifies p that the response with the given id has been received
// (or, if this is a server, sent).
func (p *Pipeline) EndResponse(id uint)

// A sequencer schedules a sequence of numbered events that must
// happen in order, one after the other. The event numbering must start
// at 0 and increment without skipping. The event number wraps around
// safely as long as there are not 2^32 simultaneous events pending.
