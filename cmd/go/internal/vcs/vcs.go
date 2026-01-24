// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package vcs

import (
	"github.com/shogo82148/std/time"

	"github.com/shogo82148/std/cmd/go/internal/web"
)

// A Cmd describes how to use a version control system
// like Mercurial, Git, or Subversion.
type Cmd struct {
	Name  string
	Cmd   string
	Env   []string
	Roots []isVCSRoot

	Scheme  []string
	PingCmd string

	Status func(v *Cmd, rootDir string) (Status, error)
}

// Status is the current state of a local repository.
type Status struct {
	Revision    string
	CommitTime  time.Time
	Uncommitted bool
}

var (
	// VCSTestRepoURL is the URL of the HTTP server that serves the repos for
	// vcs-test.golang.org.
	//
	// In tests, this is set to the URL of an httptest.Server hosting a
	// cmd/go/internal/vcweb.Server.
	VCSTestRepoURL string

	// VCSTestHosts is the set of hosts supported by the vcs-test server.
	VCSTestHosts []string

	// VCSTestIsLocalHost reports whether the given URL refers to a local
	// (loopback) host, such as "localhost" or "127.0.0.1:8080".
	VCSTestIsLocalHost func(*urlpkg.URL) bool
)

func (v *Cmd) IsSecure(repo string) bool

func (v *Cmd) String() string

// Ping pings to determine scheme to use.
func (v *Cmd) Ping(scheme, repo string) error

// FromDir inspects dir and its parents to determine the
// version control system and code repository to use.
// If no repository is found, FromDir returns an error
// equivalent to os.ErrNotExist.
func FromDir(dir, srcRoot string) (repoDir string, vcsCmd *Cmd, err error)

// RepoRoot describes the repository root for a tree of source code.
type RepoRoot struct {
	Repo     string
	Root     string
	SubDir   string
	IsCustom bool
	VCS      *Cmd
}

// ModuleMode specifies whether to prefer modules when looking up code sources.
type ModuleMode int

const (
	IgnoreMod ModuleMode = iota
	PreferMod
)

// RepoRootForImportPath analyzes importPath to determine the
// version control system, and code repository to use.
func RepoRootForImportPath(importPath string, mod ModuleMode, security web.SecurityMode) (*RepoRoot, error)

// An ImportMismatchError is returned where metaImport/s are present
// but none match our import path.
type ImportMismatchError struct {
	importPath string
	mismatches []string
}

func (m ImportMismatchError) Error() string
