// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package codehost defines the interface implemented by a code hosting source,
// along with support code for use by implementations.
package codehost

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"
)

// Downloaded size limits.
const (
	MaxGoMod   = 16 << 20
	MaxLICENSE = 16 << 20
	MaxZipFile = 500 << 20
)

// A Repo represents a code hosting source.
// Typical implementations include local version control repositories,
// remote version control servers, and code hosting sites.
//
// A Repo must be safe for simultaneous use by multiple goroutines,
// and callers must not modify returned values, which may be cached and shared.
type Repo interface {
	CheckReuse(ctx context.Context, old *Origin, subdir string) error

	Tags(ctx context.Context, prefix string) (*Tags, error)

	Stat(ctx context.Context, rev string) (*RevInfo, error)

	Latest(ctx context.Context) (*RevInfo, error)

	ReadFile(ctx context.Context, rev, file string, maxSize int64) (data []byte, err error)

	ReadZip(ctx context.Context, rev, subdir string, maxSize int64) (zip io.ReadCloser, err error)

	RecentTag(ctx context.Context, rev, prefix string, allowed func(tag string) bool) (tag string, err error)

	DescendsFrom(ctx context.Context, rev, tag string) (bool, error)
}

// An Origin describes the provenance of a given repo method result.
// It can be passed to CheckReuse (usually in a different go command invocation)
// to see whether the result remains up-to-date.
type Origin struct {
	VCS    string `json:",omitempty"`
	URL    string `json:",omitempty"`
	Subdir string `json:",omitempty"`

	Hash string `json:",omitempty"`

	// If TagSum is non-empty, then the resolution of this module version
	// depends on the set of tags present in the repo, specifically the tags
	// of the form TagPrefix + a valid semver version.
	// If the matching repo tags and their commit hashes still hash to TagSum,
	// the Origin is still valid (at least as far as the tags are concerned).
	// The exact checksum is up to the Repo implementation; see (*gitRepo).Tags.
	TagPrefix string `json:",omitempty"`
	TagSum    string `json:",omitempty"`

	// If Ref is non-empty, then the resolution of this module version
	// depends on Ref resolving to the revision identified by Hash.
	// If Ref still resolves to Hash, the Origin is still valid (at least as far as Ref is concerned).
	// For Git, the Ref is a full ref like "refs/heads/main" or "refs/tags/v1.2.3",
	// and the Hash is the Git object hash the ref maps to.
	// Other VCS might choose differently, but the idea is that Ref is the name
	// with a mutable meaning while Hash is a name with an immutable meaning.
	Ref string `json:",omitempty"`

	// If RepoSum is non-empty, then the resolution of this module version
	// depends on the entire state of the repo, which RepoSum summarizes.
	// For Git, this is a hash of all the refs and their hashes, and the RepoSum
	// is only needed for module versions that don't exist.
	// For Mercurial, this is a hash of all the branches and their heads' hashes,
	// since the set of available tags is dervied from .hgtags files in those branches,
	// and the RepoSum is used for all module versions, available and not,
	RepoSum string `json:",omitempty"`
}

// A Tags describes the available tags in a code repository.
type Tags struct {
	Origin *Origin
	List   []Tag
}

// A Tag describes a single tag in a code repository.
type Tag struct {
	Name string
	Hash string
}

// A RevInfo describes a single revision in a source code repository.
type RevInfo struct {
	Origin  *Origin
	Name    string
	Short   string
	Version string
	Time    time.Time
	Tags    []string
}

// UnknownRevisionError is an error equivalent to fs.ErrNotExist, but for a
// revision rather than a file.
type UnknownRevisionError struct {
	Rev string
}

func (e *UnknownRevisionError) Error() string

func (UnknownRevisionError) Is(err error) bool

// ErrNoCommits is an error equivalent to fs.ErrNotExist indicating that a given
// repository or module contains no commits.
var ErrNoCommits error = noCommitsError{}

// AllHex reports whether the revision rev is entirely lower-case hexadecimal digits.
func AllHex(rev string) bool

// ShortenSHA1 shortens a SHA1 hash (40 hex digits) to the canonical length
// used in pseudo-versions (12 hex digits).
func ShortenSHA1(rev string) string

// WorkDir returns the name of the cached work directory to use for the
// given repository type and name.
func WorkDir(ctx context.Context, typ, name string) (dir, lockfile string, err error)

type RunError struct {
	Cmd      string
	Err      error
	Stderr   []byte
	HelpText string
}

func (e *RunError) Error() string

type RunArgs struct {
	cmdline []any
	dir     string
	local   bool
	env     []string
	stdin   io.Reader
}

// Run runs the command line in the given directory
// (an empty dir means the current directory).
// It returns the standard output and, for a non-zero exit,
// a *RunError indicating the command, exit status, and standard error.
// Standard error is unavailable for commands that exit successfully.
func Run(ctx context.Context, dir string, cmdline ...any) ([]byte, error)

// RunWithArgs is the same as Run but it also accepts additional arguments.
func RunWithArgs(ctx context.Context, args RunArgs) ([]byte, error)
