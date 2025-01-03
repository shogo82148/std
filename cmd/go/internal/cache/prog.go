// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/cmd/go/internal/cacheprog"
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/encoding/json"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os/exec"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
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
	can map[cacheprog.Cmd]bool

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
	inFlight   map[int64]chan<- *cacheprog.Response
	outputFile map[OutputID]string

	// writeMu serializes writing to the child process.
	// It must never be held at the same time as mu.
	writeMu sync.Mutex
}

func (c *ProgCache) Get(a ActionID) (Entry, error)

func (c *ProgCache) OutputFile(o OutputID) string

func (c *ProgCache) Put(a ActionID, file io.ReadSeeker) (_ OutputID, size int64, _ error)

func (c *ProgCache) Close() error

func (c *ProgCache) FuzzDir() string
