// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cache implements a build artifact cache.
package cache

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"
)

// An ActionID is a cache action key, the hash of a complete description of a
// repeatable computation (command line, environment variables,
// input file contents, executable contents).
type ActionID [HashSize]byte

// An OutputID is a cache output key, the hash of an output of a computation.
type OutputID [HashSize]byte

// Cache is the interface as used by the cmd/go.
type Cache interface {
	Get(ActionID) (Entry, error)

	Put(ActionID, io.ReadSeeker) (_ OutputID, size int64, _ error)

	Close() error

	OutputFile(OutputID) string

	FuzzDir() string
}

// A Cache is a package cache, backed by a file system directory tree.
type DiskCache struct {
	dir string
	now func() time.Time
}

// Open opens and returns the cache in the given directory.
//
// It is safe for multiple processes on a single machine to use the
// same cache directory in a local file system simultaneously.
// They will coordinate using operating system file locks and may
// duplicate effort but will not corrupt the cache.
//
// However, it is NOT safe for multiple processes on different machines
// to share a cache directory (for example, if the directory were stored
// in a network file system). File locking is notoriously unreliable in
// network file systems and may not suffice to protect the cache.
func Open(dir string) (*DiskCache, error)

// DebugTest is set when GODEBUG=gocachetest=1 is in the environment.
var DebugTest = false

// Get looks up the action ID in the cache,
// returning the corresponding output ID and file size, if any.
// Note that finding an output ID does not guarantee that the
// saved file for that output ID is still available.
func (c *DiskCache) Get(id ActionID) (Entry, error)

type Entry struct {
	OutputID OutputID
	Size     int64
	Time     time.Time
}

// GetFile looks up the action ID in the cache and returns
// the name of the corresponding data file.
func GetFile(c Cache, id ActionID) (file string, entry Entry, err error)

// GetBytes looks up the action ID in the cache and returns
// the corresponding output bytes.
// GetBytes should only be used for data that can be expected to fit in memory.
func GetBytes(c Cache, id ActionID) ([]byte, Entry, error)

// GetMmap looks up the action ID in the cache and returns
// the corresponding output bytes.
// GetMmap should only be used for data that can be expected to fit in memory.
func GetMmap(c Cache, id ActionID) ([]byte, Entry, error)

// OutputFile returns the name of the cache file storing output with the given OutputID.
func (c *DiskCache) OutputFile(out OutputID) string

func (c *DiskCache) Close() error

// Trim removes old cache entries that are likely not to be reused.
func (c *DiskCache) Trim() error

// Put stores the given output in the cache as the output for the action ID.
// It may read file twice. The content of file must not change between the two passes.
func (c *DiskCache) Put(id ActionID, file io.ReadSeeker) (OutputID, int64, error)

// PutExecutable is used to store the output as the output for the action ID into a
// file with the given base name, with the executable mode bit set.
// It may read file twice. The content of file must not change between the two passes.
func (c *DiskCache) PutExecutable(id ActionID, name string, file io.ReadSeeker) (OutputID, int64, error)

// PutNoVerify is like Put but disables the verify check
// when GODEBUG=goverifycache=1 is set.
// It is meant for data that is OK to cache but that we expect to vary slightly from run to run,
// like test output containing times and the like.
func PutNoVerify(c Cache, id ActionID, file io.ReadSeeker) (OutputID, int64, error)

// PutBytes stores the given bytes in the cache as the output for the action ID.
func PutBytes(c Cache, id ActionID, data []byte) error

// FuzzDir returns a subdirectory within the cache for storing fuzzing data.
// The subdirectory may not exist.
//
// This directory is managed by the internal/fuzz package. Files in this
// directory aren't removed by the 'go clean -cache' command or by Trim.
// They may be removed with 'go clean -fuzzcache'.
//
// TODO(#48526): make Trim remove unused files from this directory.
func (c *DiskCache) FuzzDir() string
