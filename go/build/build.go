// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package build

import (
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/path/filepath"
	"github.com/shogo82148/std/runtime"
	pathpkg "path"
)

// A Context specifies the supporting context for a build.
type Context struct {
	GOARCH      string
	GOOS        string
	GOROOT      string
	GOPATH      string
	CgoEnabled  bool
	UseAllFiles bool
	Compiler    string

	BuildTags   []string
	ReleaseTags []string

	InstallSuffix string

	JoinPath func(elem ...string) string

	SplitPathList func(list string) []string

	IsAbsPath func(path string) bool

	IsDir func(path string) bool

	HasSubdir func(root, dir string) (rel string, ok bool)

	ReadDir func(dir string) (fi []os.FileInfo, err error)

	OpenFile func(path string) (r io.ReadCloser, err error)
}

// SrcDirs returns a list of package source root directories.
// It draws from the current Go root and Go path but omits directories
// that do not exist.
func (ctxt *Context) SrcDirs() []string

// Default is the default Context for builds.
// It uses the GOARCH, GOOS, GOROOT, and GOPATH environment variables
// if set, or else the compiled code's GOARCH, GOOS, and GOROOT.
var Default Context = defaultContext()

// An ImportMode controls the behavior of the Import method.
type ImportMode uint

const (
	// If FindOnly is set, Import stops after locating the directory
	// that should contain the sources for a package.  It does not
	// read any files in the directory.
	FindOnly ImportMode = 1 << iota

	// If AllowBinary is set, Import can be satisfied by a compiled
	// package object without corresponding sources.
	AllowBinary

	// If ImportComment is set, parse import comments on package statements.
	// Import returns an error if it finds a comment it cannot understand
	// or finds conflicting comments in multiple source files.
	// See golang.org/s/go14customimport for more information.
	ImportComment
)

// A Package describes the Go package found in a directory.
type Package struct {
	Dir           string
	Name          string
	ImportComment string
	Doc           string
	ImportPath    string
	Root          string
	SrcRoot       string
	PkgRoot       string
	BinDir        string
	Goroot        bool
	PkgObj        string
	AllTags       []string
	ConflictDir   string

	GoFiles        []string
	CgoFiles       []string
	IgnoredGoFiles []string
	CFiles         []string
	CXXFiles       []string
	MFiles         []string
	HFiles         []string
	SFiles         []string
	SwigFiles      []string
	SwigCXXFiles   []string
	SysoFiles      []string

	CgoCFLAGS    []string
	CgoCPPFLAGS  []string
	CgoCXXFLAGS  []string
	CgoLDFLAGS   []string
	CgoPkgConfig []string

	Imports   []string
	ImportPos map[string][]token.Position

	TestGoFiles    []string
	TestImports    []string
	TestImportPos  map[string][]token.Position
	XTestGoFiles   []string
	XTestImports   []string
	XTestImportPos map[string][]token.Position
}

// IsCommand reports whether the package is considered a
// command to be installed (not just a library).
// Packages named "main" are treated as commands.
func (p *Package) IsCommand() bool

// ImportDir is like Import but processes the Go package found in
// the named directory.
func (ctxt *Context) ImportDir(dir string, mode ImportMode) (*Package, error)

// NoGoError is the error used by Import to describe a directory
// containing no buildable Go source files. (It may still contain
// test files, files hidden by build tags, and so on.)
type NoGoError struct {
	Dir string
}

func (e *NoGoError) Error() string

// MultiplePackageError describes a directory containing
// multiple buildable Go source files for multiple packages.
type MultiplePackageError struct {
	Dir      string
	Packages []string
	Files    []string
}

func (e *MultiplePackageError) Error() string

// Import returns details about the Go package named by the import path,
// interpreting local import paths relative to the srcDir directory.
// If the path is a local import path naming a package that can be imported
// using a standard import path, the returned package will set p.ImportPath
// to that path.
//
// In the directory containing the package, .go, .c, .h, and .s files are
// considered part of the package except for:
//
//   - .go files in package documentation
//   - files starting with _ or . (likely editor temporary files)
//   - files with build constraints not satisfied by the context
//
// If an error occurs, Import returns a non-nil error and a non-nil
// *Package containing partial information.
func (ctxt *Context) Import(path string, srcDir string, mode ImportMode) (*Package, error)

// MatchFile reports whether the file with the given name in the given directory
// matches the context and would be included in a Package created by ImportDir
// of that directory.
//
// MatchFile considers the name of the file and may use ctxt.OpenFile to
// read some or all of the file's content.
func (ctxt *Context) MatchFile(dir, name string) (match bool, err error)

// Import is shorthand for Default.Import.
func Import(path, srcDir string, mode ImportMode) (*Package, error)

// ImportDir is shorthand for Default.ImportDir.
func ImportDir(dir string, mode ImportMode) (*Package, error)

// NOTE: $ is not safe for the shell, but it is allowed here because of linker options like -Wl,$ORIGIN.
// We never pass these arguments to a shell (just to programs we construct argv for), so this should be okay.
// See golang.org/issue/6038.

// ToolDir is the directory containing build tools.
var ToolDir = filepath.Join(runtime.GOROOT(), "pkg/tool/"+runtime.GOOS+"_"+runtime.GOARCH)

// IsLocalImport reports whether the import path is
// a local import path, like ".", "..", "./foo", or "../foo".
func IsLocalImport(path string) bool

// ArchChar returns the architecture character for the given goarch.
// For example, ArchChar("amd64") returns "6".
func ArchChar(goarch string) (string, error)
