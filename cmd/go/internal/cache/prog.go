// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/encoding/json"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os/exec"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/time"
)

// ProgCache implements Cache via JSON messages over stdin/stdout to a child
// helper process which can then implement whatever caching policy/mechanism it
// wants.
//
// See https://github.com/golang/go/issues/59719
type ProgCache struct {
	cmd    *exec.Cmd
	stdout io.ReadCloser
	stdin  io.WriteCloser
	bw     *bufio.Writer
	jenc   *json.Encoder

	// can are the commands that the child process declared that it supports.
	// This is effectively the versioning mechanism.
	can map[ProgCmd]bool

	// fuzzDirCache is another Cache implementation to use for the FuzzDir
	// method. In practice this is the default GOCACHE disk-based
	// implementation.
	//
	// TODO(bradfitz): maybe this isn't ideal. But we'd need to extend the Cache
	// interface and the fuzzing callers to be less disk-y to do more here.
	fuzzDirCache Cache

	closing      atomic.Bool
	ctx          context.Context
	ctxCancel    context.CancelFunc
	readLoopDone chan struct{}

	mu         sync.Mutex
	nextID     int64
	inFlight   map[int64]chan<- *ProgResponse
	outputFile map[OutputID]string

	// writeMu serializes writing to the child process.
	// It must never be held at the same time as mu.
	writeMu sync.Mutex
}

// ProgCmd is a command that can be issued to a child process.
//
// If the interface needs to grow, we can add new commands or new versioned
// commands like "get2".
type ProgCmd string

// ProgRequest is the JSON-encoded message that's sent from cmd/go to
// the GOCACHEPROG child process over stdin. Each JSON object is on its
// own line. A ProgRequest of Type "put" with BodySize > 0 will be followed
// by a line containing a base64-encoded JSON string literal of the body.
type ProgRequest struct {
	// ID is a unique number per process across all requests.
	// It must be echoed in the ProgResponse from the child.
	ID int64

	// Command is the type of request.
	// The cmd/go tool will only send commands that were declared
	// as supported by the child.
	Command ProgCmd

	// ActionID is non-nil for get and puts.
	ActionID []byte `json:",omitempty"`

	// OutputID is set for Type "put".
	//
	// Prior to Go 1.24, when GOCACHEPROG was still an experiment, this was
	// accidentally named ObjectID. It was renamed to OutputID in Go 1.24.
	OutputID []byte `json:",omitempty"`

	// Body is the body for "put" requests. It's sent after the JSON object
	// as a base64-encoded JSON string when BodySize is non-zero.
	// It's sent as a separate JSON value instead of being a struct field
	// send in this JSON object so large values can be streamed in both directions.
	// The base64 string body of a ProgRequest will always be written
	// immediately after the JSON object and a newline.
	Body io.Reader `json:"-"`

	// BodySize is the number of bytes of Body. If zero, the body isn't written.
	BodySize int64 `json:",omitempty"`

	// ObjectID is the accidental spelling of OutputID that was used prior to Go
	// 1.24.
	//
	// Deprecated: use OutputID. This field is only populated temporarily for
	// backwards compatibility with Go 1.23 and earlier when
	// GOEXPERIMENT=gocacheprog is set. It will be removed in Go 1.25.
	ObjectID []byte `json:",omitempty"`
}

// ProgResponse is the JSON response from the child process to cmd/go.
//
// With the exception of the first protocol message that the child writes to its
// stdout with ID==0 and KnownCommands populated, these are only sent in
// response to a ProgRequest from cmd/go.
//
// ProgResponses can be sent in any order. The ID must match the request they're
// replying to.
type ProgResponse struct {
	ID  int64
	Err string `json:",omitempty"`

	// KnownCommands is included in the first message that cache helper program
	// writes to stdout on startup (with ID==0). It includes the
	// ProgRequest.Command types that are supported by the program.
	//
	// This lets us extend the protocol gracefully over time (adding "get2",
	// etc), or fail gracefully when needed. It also lets us verify the program
	// wants to be a cache helper.
	KnownCommands []ProgCmd `json:",omitempty"`

	Miss     bool       `json:",omitempty"`
	OutputID []byte     `json:",omitempty"`
	Size     int64      `json:",omitempty"`
	Time     *time.Time `json:",omitempty"`

	// DiskPath is the absolute path on disk of the ObjectID corresponding
	// a "get" request's ActionID (on cache hit) or a "put" request's
	// provided ObjectID.
	DiskPath string `json:",omitempty"`
}

func (c *ProgCache) Get(a ActionID) (Entry, error)

func (c *ProgCache) OutputFile(o OutputID) string

func (c *ProgCache) Put(a ActionID, file io.ReadSeeker) (_ OutputID, size int64, _ error)

func (c *ProgCache) Close() error

func (c *ProgCache) FuzzDir() string
