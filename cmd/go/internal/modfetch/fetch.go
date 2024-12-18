// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modfetch

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/errors"

	"github.com/shogo82148/std/cmd/go/internal/base"

	"golang.org/x/mod/module"
)

var ErrToolchain = errors.New("internal error: invalid operation on toolchain module")

// Download downloads the specific module version to the
// local download cache and returns the name of the directory
// corresponding to the root of the module's file tree.
func Download(ctx context.Context, mod module.Version) (dir string, err error)

// Unzip is like Download but is given the explicit zip file to use,
// rather than downloading it. This is used for the GOFIPS140 zip files,
// which ship in the Go distribution itself.
func Unzip(ctx context.Context, mod module.Version, zipfile string) (dir string, err error)

// DownloadZip downloads the specific module version to the
// local zip cache and returns the name of the zip file.
func DownloadZip(ctx context.Context, mod module.Version) (zipfile string, err error)

// RemoveAll removes a directory written by Download or Unzip, first applying
// any permission changes needed to do so.
func RemoveAll(dir string) error

var GoSumFile string
var WorkspaceGoSumFiles []string

// Reset resets globals in the modfetch package, so previous loads don't affect
// contents of go.sum files.
func Reset()

// HaveSum returns true if the go.sum file contains an entry for mod.
// The entry's hash must be generated with a known hash algorithm.
// mod.Version may have a "/go.mod" suffix to distinguish sums for
// .mod and .zip files.
func HaveSum(mod module.Version) bool

// RecordedSum returns the sum if the go.sum file contains an entry for mod.
// The boolean reports true if an entry was found or
// false if no entry found or two conflicting sums are found.
// The entry's hash must be generated with a known hash algorithm.
// mod.Version may have a "/go.mod" suffix to distinguish sums for
// .mod and .zip files.
func RecordedSum(mod module.Version) (sum string, ok bool)

// Sum returns the checksum for the downloaded copy of the given module,
// if present in the download cache.
func Sum(ctx context.Context, mod module.Version) string

var ErrGoSumDirty = errors.New("updates to go.sum needed, disabled by -mod=readonly")

// WriteGoSum writes the go.sum file if it needs to be updated.
//
// keep is used to check whether a newly added sum should be saved in go.sum.
// It should have entries for both module content sums and go.mod sums
// (version ends with "/go.mod"). Existing sums will be preserved unless they
// have been marked for deletion with TrimGoSum.
func WriteGoSum(ctx context.Context, keep map[module.Version]bool, readonly bool) error

// TidyGoSum returns a tidy version of the go.sum file.
// A missing go.sum file is treated as if empty.
func TidyGoSum(keep map[module.Version]bool) (before, after []byte)

// TrimGoSum trims go.sum to contain only the modules needed for reproducible
// builds.
//
// keep is used to check whether a sum should be retained in go.mod. It should
// have entries for both module content sums and go.mod sums (version ends
// with "/go.mod").
func TrimGoSum(keep map[module.Version]bool)

var HelpModuleAuth = &base.Command{
	UsageLine: "module-auth",
	Short:     "module authentication using go.sum",
	Long: `
When the go command downloads a module zip file or go.mod file into the
module cache, it computes a cryptographic hash and compares it with a known
value to verify the file hasn't changed since it was first downloaded. Known
hashes are stored in a file in the module root directory named go.sum. Hashes
may also be downloaded from the checksum database depending on the values of
GOSUMDB, GOPRIVATE, and GONOSUMDB.

For details, see https://golang.org/ref/mod#authenticating.
`,
}

var HelpPrivate = &base.Command{
	UsageLine: "private",
	Short:     "configuration for downloading non-public code",
	Long: `
The go command defaults to downloading modules from the public Go module
mirror at proxy.golang.org. It also defaults to validating downloaded modules,
regardless of source, against the public Go checksum database at sum.golang.org.
These defaults work well for publicly available source code.

The GOPRIVATE environment variable controls which modules the go command
considers to be private (not available publicly) and should therefore not use
the proxy or checksum database. The variable is a comma-separated list of
glob patterns (in the syntax of Go's path.Match) of module path prefixes.
For example,

	GOPRIVATE=*.corp.example.com,rsc.io/private

causes the go command to treat as private any module with a path prefix
matching either pattern, including git.corp.example.com/xyzzy, rsc.io/private,
and rsc.io/private/quux.

For fine-grained control over module download and validation, the GONOPROXY
and GONOSUMDB environment variables accept the same kind of glob list
and override GOPRIVATE for the specific decision of whether to use the proxy
and checksum database, respectively.

For example, if a company ran a module proxy serving private modules,
users would configure go using:

	GOPRIVATE=*.corp.example.com
	GOPROXY=proxy.example.com
	GONOPROXY=none

The GOPRIVATE variable is also used to define the "public" and "private"
patterns for the GOVCS variable; see 'go help vcs'. For that usage,
GOPRIVATE applies even in GOPATH mode. In that case, it matches import paths
instead of module paths.

The 'go env -w' command (see 'go help env') can be used to set these variables
for future go command invocations.

For more details, see https://golang.org/ref/mod#private-modules.
`,
}
