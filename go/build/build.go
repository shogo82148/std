// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package build

import (
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
	pathpkg "path"
)

// A Context specifies the supporting context for a build.
type Context struct {
	GOARCH string
	GOOS   string
	GOROOT string
	GOPATH string

	Dir string

	CgoEnabled  bool
	UseAllFiles bool
	Compiler    string

	BuildTags   []string
	ToolTags    []string
	ReleaseTags []string

	InstallSuffix string

	JoinPath func(elem ...string) string

	SplitPathList func(list string) []string

	IsAbsPath func(path string) bool

	IsDir func(path string) bool

	HasSubdir func(root, dir string) (rel string, ok bool)

	ReadDir func(dir string) ([]fs.FileInfo, error)

	OpenFile func(path string) (io.ReadCloser, error)
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
	// that should contain the sources for a package. It does not
	// read any files in the directory.
	FindOnly ImportMode = 1 << iota

	// If AllowBinary is set, Import can be satisfied by a compiled
	// package object without corresponding sources.
	//
	// Deprecated:
	// The supported way to create a compiled-only package is to
	// write source code containing a //go:binary-only-package comment at
	// the top of the file. Such a package will be recognized
	// regardless of this flag setting (because it has source code)
	// and will have BinaryOnly set to true in the returned Package.
	AllowBinary

	// If ImportComment is set, parse import comments on package statements.
	// Import returns an error if it finds a comment it cannot understand
	// or finds conflicting comments in multiple source files.
	// See golang.org/s/go14customimport for more information.
	ImportComment

	// By default, Import searches vendor directories
	// that apply in the given source directory before searching
	// the GOROOT and GOPATH roots.
	// If an Import finds and returns a package using a vendor
	// directory, the resulting ImportPath is the complete path
	// to the package, including the path elements leading up
	// to and including "vendor".
	// For example, if Import("y", "x/subdir", 0) finds
	// "x/vendor/y", the returned package's ImportPath is "x/vendor/y",
	// not plain "y".
	// See golang.org/s/go15vendor for more information.
	//
	// Setting IgnoreVendor ignores vendor directories.
	//
	// In contrast to the package's ImportPath,
	// the returned package's Imports, TestImports, and XTestImports
	// are always the exact import paths from the source files:
	// Import makes no attempt to resolve or check those paths.
	IgnoreVendor
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
	PkgTargetRoot string
	BinDir        string
	Goroot        bool
	PkgObj        string
	AllTags       []string
	ConflictDir   string
	BinaryOnly    bool

	GoFiles           []string
	CgoFiles          []string
	IgnoredGoFiles    []string
	InvalidGoFiles    []string
	IgnoredOtherFiles []string
	CFiles            []string
	CXXFiles          []string
	MFiles            []string
	HFiles            []string
	FFiles            []string
	SFiles            []string
	SwigFiles         []string
	SwigCXXFiles      []string
	SysoFiles         []string

	CgoCFLAGS    []string
	CgoCPPFLAGS  []string
	CgoCXXFLAGS  []string
	CgoFFLAGS    []string
	CgoLDFLAGS   []string
	CgoPkgConfig []string

	TestGoFiles  []string
	XTestGoFiles []string

	Imports        []string
	ImportPos      map[string][]token.Position
	TestImports    []string
	TestImportPos  map[string][]token.Position
	XTestImports   []string
	XTestImportPos map[string][]token.Position

	EmbedPatterns        []string
	EmbedPatternPos      map[string][]token.Position
	TestEmbedPatterns    []string
	TestEmbedPatternPos  map[string][]token.Position
	XTestEmbedPatterns   []string
	XTestEmbedPatternPos map[string][]token.Position
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

// fileInfo records information learned about a file included in a build.

// Import is shorthand for Default.Import.
func Import(path, srcDir string, mode ImportMode) (*Package, error)

// ImportDir is shorthand for Default.ImportDir.
func ImportDir(dir string, mode ImportMode) (*Package, error)

// Special comment denoting a binary-only package.
// See https://golang.org/design/2775-binary-only-packages
// for more about the design of binary-only packages.

// NOTE: $ is not safe for the shell, but it is allowed here because of linker options like -Wl,$ORIGIN.
// We never pass these arguments to a shell (just to programs we construct argv for), so this should be okay.
// See golang.org/issue/6038.
// The @ is for OS X. See golang.org/issue/13720.
// The % is for Jenkins. See golang.org/issue/16959.
// The ! is because module paths may use them. See golang.org/issue/26716.
// The ~ and ^ are for sr.ht. See golang.org/issue/32260.

// ToolDir is the directory containing build tools.
var ToolDir = getToolDir()

// IsLocalImport reports whether the import path is
// a local import path, like ".", "..", "./foo", or "../foo".
func IsLocalImport(path string) bool

// ArchChar returns "?" and an error.
// In earlier versions of Go, the returned string was used to derive
// the compiler and linker tool names, the default object file suffix,
// and the default linker output name. As of Go 1.5, those strings
// no longer vary by architecture; they are compile, link, .o, and a.out, respectively.
func ArchChar(goarch string) (string, error)
