// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modfetch

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"

	"github.com/shogo82148/std/cmd/go/internal/modfetch/codehost"
)

// A Repo represents a repository storing all versions of a single module.
// It must be safe for simultaneous use by multiple goroutines.
type Repo interface {
	// ModulePath returns the module path.
	ModulePath() string

	// CheckReuse checks whether the validation criteria in the origin
	// are still satisfied on the server corresponding to this module.
	// If so, the caller can reuse any cached Versions or RevInfo containing
	// this origin rather than redownloading those from the server.
	CheckReuse(ctx context.Context, old *codehost.Origin) error

	// Versions lists all known versions with the given prefix.
	// Pseudo-versions are not included.
	//
	// Versions should be returned sorted in semver order
	// (implementations can use semver.Sort).
	//
	// Versions returns a non-nil error only if there was a problem
	// fetching the list of versions: it may return an empty list
	// along with a nil error if the list of matching versions
	// is known to be empty.
	//
	// If the underlying repository does not exist,
	// Versions returns an error matching errors.Is(_, os.NotExist).
	Versions(ctx context.Context, prefix string) (*Versions, error)

	// Stat returns information about the revision rev.
	// A revision can be any identifier known to the underlying service:
	// commit hash, branch, tag, and so on.
	Stat(ctx context.Context, rev string) (*RevInfo, error)

	// Latest returns the latest revision on the default branch,
	// whatever that means in the underlying source code repository.
	// It is only used when there are no tagged versions.
	Latest(ctx context.Context) (*RevInfo, error)

	// GoMod returns the go.mod file for the given version.
	GoMod(ctx context.Context, version string) (data []byte, err error)

	// Zip writes a zip file for the given version to dst.
	Zip(ctx context.Context, dst io.Writer, version string) error
}

// A Versions describes the available versions in a module repository.
type Versions struct {
	Origin *codehost.Origin `json:",omitempty"`

	List []string
}

// A RevInfo describes a single revision in a module repository.
type RevInfo struct {
	Version string
	Time    time.Time

	// These fields are used for Stat of arbitrary rev,
	// but they are not recorded when talking about module versions.
	Name  string `json:"-"`
	Short string `json:"-"`

	Origin *codehost.Origin `json:",omitempty"`
}

// Lookup returns the module with the given module path,
// fetched through the given proxy.
//
// The distinguished proxy "direct" indicates that the path should be fetched
// from its origin, and "noproxy" indicates that the patch should be fetched
// directly only if GONOPROXY matches the given path.
//
// For the distinguished proxy "off", Lookup always returns a Repo that returns
// a non-nil error for every method call.
//
// A successful return does not guarantee that the module
// has any defined versions.
func Lookup(ctx context.Context, proxy, path string) Repo
