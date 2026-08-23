// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gentools provides shared helper utilities for Go code generator tools
// in archsimd.
//
// Basic usage:
//
//	func main() {
//	    gentools.RegisterFlags(nil)
//	    flag.Parse()
//
//	    var files gentools.Files
//	    defer files.FlushOrExit()
//
//	    buf := files.NewGoFile("src/simd/archsimd/ops_amd64.go")
//	    fmt.Fprintln(buf, "package archsimd")
//	    // ... write generated code to buf ...
//	}
//
// By default (when -w is not specified), gentools outputs all generated files
// as a txtar archive to standard output. Pass -w to write files directly into
// the Go source tree.
package gentools

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/flag"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
)

// Options contains standard options and CLI flags for code generators.
type Options struct {
	GOROOT string
	outDir string
	Write  bool
	Diff   bool
	Txtar  bool

	Output    io.Writer
	ErrOutput io.Writer
}

// RegisterFlags registers standard generator flags with the provided FlagSet
// (or [flag.CommandLine] if fs is nil) and returns a pointer to the Options
// struct.
//
// If fs is nil, the returned options are remembered globally as defaults for
// zero-value Files instances.
func RegisterFlags(fs *flag.FlagSet) *Options

// InputPath resolves relPath relative to either o.OutDir/src, if that file
// exists, or o.GOROOT/src. In effect, o.OutDir is treated as an overlay on
// o.GOROOT.
func (o *Options) InputPath(relPath string) string

// ReadFile reads relPath from either o.OutDir/src or o.GOROOT/src.
func (o *Options) ReadFile(relPath string) ([]byte, error)

// OutputPath returns relPath relative to o.OutDir/src.
func (o *Options) OutputPath(relPath string) string

// WritingToInput returns true if Flush will write to the input tree.
func (o *Options) WritingToInput() bool

// Files manages a collection of generated files for a single generator run.
// The zero value of Files is ready for immediate use and automatically honors
// the command-line flags registered via RegisterFlags.
type Files struct {
	// Options optionally overrides the generator options for this Files instance.
	// If nil, the globally registered options from RegisterFlags are used automatically.
	Options *Options

	files []*fileInfo

	// tmpDir is a temporary directory used for communicating with subprocess
	// gentools.
	tmpDirOnce sync.Once
	tmpDir     string
}

// NewGoFile registers a Go source file at relPath (relative to GOROOT/src). It
// returns a *bytes.Buffer for the generator to populate. During Flush(), Go
// files are formatted with go/format.
func (f *Files) NewGoFile(relPath string) *bytes.Buffer

// NewRawFile registers a non-Go file (e.g. .rules, YAML, txtar) at relPath
// (relative to GOROOT/src). It returns a *bytes.Buffer for the generator to
// populate. During Flush(), content is written directly without go/format.
func (f *Files) NewRawFile(relPath string) *bytes.Buffer

// ExecFlags returns a sequence of flags that can be passed to a gentools
// subprocess. This allows several gentools to be tied together by a larger
// gentool, including if later gentools read the outputs of earlier gentools.
//
// Regardless of the output mode of f, this directs subprocesses to write to a
// temporary directory. Flush then reads the contents of this temporary
// directory back as if this process had written all of those files using f and
// applies the configured output mode.
func (f *Files) ExecFlags() []string

// Flush outputs all registered files according to the mode in options.
//
// In default / -txtar mode, it outputs files as a txtar archive to Output. In
// write mode (-w), it writes all files to disk under GOROOT. In diff mode
// (-diff), it compares generated content against disk, prints diffs to Output,
// and returns an error if out of date.
func (f *Files) Flush() error

// FlushOrExit calls Flush(), prints any error to stderr, and exits with code 1 if Flush fails.
//
// It is intended to be deferred at the beginning of main (e.g., `defer files.FlushOrExit()`).
// Hence, if invoked as part of a panic, it skips flushing and instead allows the panic to propagate.
func (f *Files) FlushOrExit()
