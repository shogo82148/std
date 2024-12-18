// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modfetch

import (
	"github.com/shogo82148/std/context"

	"golang.org/x/mod/module"
)

func CachePath(ctx context.Context, m module.Version, suffix string) (string, error)

// DownloadDir returns the directory to which m should have been downloaded.
// An error will be returned if the module path or version cannot be escaped.
// An error satisfying errors.Is(err, fs.ErrNotExist) will be returned
// along with the directory if the directory does not exist or if the directory
// is not completely populated.
func DownloadDir(ctx context.Context, m module.Version) (string, error)

// DownloadDirPartialError is returned by DownloadDir if a module directory
// exists but was not completely populated.
//
// DownloadDirPartialError is equivalent to fs.ErrNotExist.
type DownloadDirPartialError struct {
	Dir string
	Err error
}

func (e *DownloadDirPartialError) Error() string
func (e *DownloadDirPartialError) Is(err error) bool

// SideLock locks a file within the module cache that previously guarded
// edits to files outside the cache, such as go.sum and go.mod files in the
// user's working directory.
// If err is nil, the caller MUST eventually call the unlock function.
func SideLock(ctx context.Context) (unlock func(), err error)

// InfoFile is like Lookup(ctx, path).Stat(version) but also returns the name of the file
// containing the cached information.
func InfoFile(ctx context.Context, path, version string) (*RevInfo, string, error)

// GoMod is like Lookup(ctx, path).GoMod(rev) but avoids the
// repository path resolution in Lookup if the result is
// already cached on local disk.
func GoMod(ctx context.Context, path, rev string) ([]byte, error)

// GoModFile is like GoMod but returns the name of the file containing
// the cached information.
func GoModFile(ctx context.Context, path, version string) (string, error)

// GoModSum returns the go.sum entry for the module version's go.mod file.
// (That is, it returns the entry listed in go.sum as "path version/go.mod".)
func GoModSum(ctx context.Context, path, version string) (string, error)
