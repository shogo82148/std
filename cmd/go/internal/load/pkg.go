// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package load loads packages.
package load

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/go/build"
	"github.com/shogo82148/std/go/token"

	"github.com/shogo82148/std/runtime/debug"

	"github.com/shogo82148/std/cmd/go/internal/modinfo"
	"github.com/shogo82148/std/cmd/go/internal/modload"
)

// A Package describes a single package found in a directory.
type Package struct {
	PackagePublic
	Internal PackageInternal
}

type PackagePublic struct {
	// Note: These fields are part of the go command's public API.
	// See list.go. It is okay to add fields, but not to change or
	// remove existing ones. Keep in sync with ../list/list.go
	Dir           string                `json:",omitempty"`
	ImportPath    string                `json:",omitempty"`
	ImportComment string                `json:",omitempty"`
	Name          string                `json:",omitempty"`
	Doc           string                `json:",omitempty"`
	Target        string                `json:",omitempty"`
	Shlib         string                `json:",omitempty"`
	Root          string                `json:",omitempty"`
	ConflictDir   string                `json:",omitempty"`
	ForTest       string                `json:",omitempty"`
	Export        string                `json:",omitempty"`
	BuildID       string                `json:",omitempty"`
	Module        *modinfo.ModulePublic `json:",omitempty"`
	Match         []string              `json:",omitempty"`
	Goroot        bool                  `json:",omitempty"`
	Standard      bool                  `json:",omitempty"`
	DepOnly       bool                  `json:",omitempty"`
	BinaryOnly    bool                  `json:",omitempty"`
	Incomplete    bool                  `json:",omitempty"`

	DefaultGODEBUG string `json:",omitempty"`

	// Stale and StaleReason remain here *only* for the list command.
	// They are only initialized in preparation for list execution.
	// The regular build determines staleness on the fly during action execution.
	Stale       bool   `json:",omitempty"`
	StaleReason string `json:",omitempty"`

	// Source files
	// If you add to this list you MUST add to p.AllFiles (below) too.
	// Otherwise file name security lists will not apply to any new additions.
	GoFiles           []string `json:",omitempty"`
	CgoFiles          []string `json:",omitempty"`
	CompiledGoFiles   []string `json:",omitempty"`
	IgnoredGoFiles    []string `json:",omitempty"`
	InvalidGoFiles    []string `json:",omitempty"`
	IgnoredOtherFiles []string `json:",omitempty"`
	CFiles            []string `json:",omitempty"`
	CXXFiles          []string `json:",omitempty"`
	MFiles            []string `json:",omitempty"`
	HFiles            []string `json:",omitempty"`
	FFiles            []string `json:",omitempty"`
	SFiles            []string `json:",omitempty"`
	SwigFiles         []string `json:",omitempty"`
	SwigCXXFiles      []string `json:",omitempty"`
	SysoFiles         []string `json:",omitempty"`

	// Embedded files
	EmbedPatterns []string `json:",omitempty"`
	EmbedFiles    []string `json:",omitempty"`

	// Cgo directives
	CgoCFLAGS    []string `json:",omitempty"`
	CgoCPPFLAGS  []string `json:",omitempty"`
	CgoCXXFLAGS  []string `json:",omitempty"`
	CgoFFLAGS    []string `json:",omitempty"`
	CgoLDFLAGS   []string `json:",omitempty"`
	CgoPkgConfig []string `json:",omitempty"`

	// Dependency information
	Imports   []string          `json:",omitempty"`
	ImportMap map[string]string `json:",omitempty"`
	Deps      []string          `json:",omitempty"`

	// Error information
	// Incomplete is above, packed into the other bools
	Error      *PackageError   `json:",omitempty"`
	DepsErrors []*PackageError `json:",omitempty"`

	// Test information
	// If you add to this list you MUST add to p.AllFiles (below) too.
	// Otherwise file name security lists will not apply to any new additions.
	TestGoFiles        []string `json:",omitempty"`
	TestImports        []string `json:",omitempty"`
	TestEmbedPatterns  []string `json:",omitempty"`
	TestEmbedFiles     []string `json:",omitempty"`
	XTestGoFiles       []string `json:",omitempty"`
	XTestImports       []string `json:",omitempty"`
	XTestEmbedPatterns []string `json:",omitempty"`
	XTestEmbedFiles    []string `json:",omitempty"`
}

// AllFiles returns the names of all the files considered for the package.
// This is used for sanity and security checks, so we include all files,
// even IgnoredGoFiles, because some subcommands consider them.
// The go/build package filtered others out (like foo_wrongGOARCH.s)
// and that's OK.
func (p *Package) AllFiles() []string

// Desc returns the package "description", for use in b.showOutput.
func (p *Package) Desc() string

// IsTestOnly reports whether p is a test-only package.
//
// A “test-only” package is one that:
//   - is a test-only variant of an ordinary package, or
//   - is a synthesized "main" package for a test binary, or
//   - contains only _test.go files.
func (p *Package) IsTestOnly() bool

type PackageInternal struct {
	// Unexported fields are not part of the public API.
	Build             *build.Package
	Imports           []*Package
	CompiledImports   []string
	RawImports        []string
	ForceLibrary      bool
	CmdlineFiles      bool
	CmdlinePkg        bool
	CmdlinePkgLiteral bool
	Local             bool
	LocalPrefix       string
	ExeName           string
	FuzzInstrument    bool
	Cover             CoverSetup
	CoverVars         map[string]*CoverVar
	OmitDebug         bool
	GobinSubdir       bool
	BuildInfo         *debug.BuildInfo
	TestmainGo        *[]byte
	Embed             map[string][]string
	OrigImportPath    string
	PGOProfile        string
	ForMain           string

	Asmflags   []string
	Gcflags    []string
	Ldflags    []string
	Gccgoflags []string
}

// A NoGoError indicates that no Go files for the package were applicable to the
// build for that package.
//
// That may be because there were no files whatsoever, or because all files were
// excluded, or because all non-excluded files were test sources.
type NoGoError struct {
	Package *Package
}

func (e *NoGoError) Error() string

// Resolve returns the resolved version of imports,
// which should be p.TestImports or p.XTestImports, NOT p.Imports.
// The imports in p.TestImports and p.XTestImports are not recursively
// loaded during the initial load of p, so they list the imports found in
// the source file, but most processing should be over the vendor-resolved
// import paths. We do this resolution lazily both to avoid file system work
// and because the eventual real load of the test imports (during 'go test')
// can produce better error messages if it starts with the original paths.
// The initial load of p loads all the non-test imports and rewrites
// the vendored paths, so nothing should ever call p.vendored(p.Imports).
func (p *Package) Resolve(imports []string) []string

// CoverVar holds the name of the generated coverage variables targeting the named file.
type CoverVar struct {
	File string
	Var  string
}

// CoverSetup holds parameters related to coverage setup for a given package (covermode, etc).
type CoverSetup struct {
	Mode    string
	Cfg     string
	GenMeta bool
}

// A PackageError describes an error loading information about a package.
type PackageError struct {
	ImportStack      []string
	Pos              string
	Err              error
	IsImportCycle    bool
	Hard             bool
	alwaysPrintStack bool
}

func (p *PackageError) Error() string

func (p *PackageError) Unwrap() error

// PackageError implements MarshalJSON so that Err is marshaled as a string
// and non-essential fields are omitted.
func (p *PackageError) MarshalJSON() ([]byte, error)

// ImportPathError is a type of error that prevents a package from being loaded
// for a given import path. When such a package is loaded, a *Package is
// returned with Err wrapping an ImportPathError: the error is attached to
// the imported package, not the importing package.
//
// The string returned by ImportPath must appear in the string returned by
// Error. Errors that wrap ImportPathError (such as PackageError) may omit
// the import path.
type ImportPathError interface {
	error
	ImportPath() string
}

var (
	_ ImportPathError = (*importError)(nil)
	_ ImportPathError = (*mainPackageError)(nil)
	_ ImportPathError = (*modload.ImportMissingError)(nil)
	_ ImportPathError = (*modload.ImportMissingSumError)(nil)
	_ ImportPathError = (*modload.DirectImportFromImplicitDependencyError)(nil)
)

func ImportErrorf(path, format string, args ...any) ImportPathError

// An ImportStack is a stack of import paths, possibly with the suffix " (test)" appended.
// The import path of a test package is the import path of the corresponding
// non-test package with the suffix "_test" added.
type ImportStack []string

func (s *ImportStack) Push(p string)

func (s *ImportStack) Pop()

func (s *ImportStack) Copy() []string

func (s *ImportStack) Top() string

// Mode flags for loadImport and download (in get.go).
const (
	// ResolveImport means that loadImport should do import path expansion.
	// That is, ResolveImport means that the import path came from
	// a source file and has not been expanded yet to account for
	// vendoring or possible module adjustment.
	// Every import path should be loaded initially with ResolveImport,
	// and then the expanded version (for example with the /vendor/ in it)
	// gets recorded as the canonical import path. At that point, future loads
	// of that package must not pass ResolveImport, because
	// disallowVendor will reject direct use of paths containing /vendor/.
	ResolveImport = 1 << iota

	// ResolveModule is for download (part of "go get") and indicates
	// that the module adjustment should be done, but not vendor adjustment.
	ResolveModule

	// GetTestDeps is for download (part of "go get") and indicates
	// that test dependencies should be fetched too.
	GetTestDeps
)

// LoadImport scans the directory named by path, which must be an import path,
// but possibly a local import path (an absolute file system path or one beginning
// with ./ or ../). A local relative path is interpreted relative to srcDir.
// It returns a *Package describing the package found in that directory.
// LoadImport does not set tool flags and should only be used by
// this package, as part of a bigger load operation, and by GOPATH-based "go get".
// TODO(rsc): When GOPATH-based "go get" is removed, unexport this function.
// The returned PackageError, if any, describes why parent is not allowed
// to import the named package, with the error referring to importPos.
// The PackageError can only be non-nil when parent is not nil.
func LoadImport(ctx context.Context, opts PackageOpts, path, srcDir string, parent *Package, stk *ImportStack, importPos []token.Position, mode int) (*Package, *PackageError)

// LoadPackage does Load import, but without a parent package load contezt
func LoadPackage(ctx context.Context, opts PackageOpts, path, srcDir string, stk *ImportStack, importPos []token.Position, mode int) *Package

// ResolveImportPath returns the true meaning of path when it appears in parent.
// There are two different resolutions applied.
// First, there is Go 1.5 vendoring (golang.org/s/go15vendor).
// If vendor expansion doesn't trigger, then the path is also subject to
// Go 1.11 module legacy conversion (golang.org/issue/25069).
func ResolveImportPath(parent *Package, path string) (found string)

// FindVendor looks for the last non-terminating "vendor" path element in the given import path.
// If there isn't one, FindVendor returns ok=false.
// Otherwise, FindVendor returns ok=true and the index of the "vendor".
//
// Note that terminating "vendor" elements don't count: "x/vendor" is its own package,
// not the vendored copy of an import "" (the empty import path).
// This will allow people to have packages or commands named vendor.
// This may help reduce breakage, or it may just be confusing. We'll see.
func FindVendor(path string) (index int, ok bool)

type TargetDir int

const (
	ToTool TargetDir = iota
	ToBin
	StalePath
)

// InstallTargetDir reports the target directory for installing the command p.
func InstallTargetDir(p *Package) TargetDir

// DefaultExecName returns the default executable name for a package
func (p *Package) DefaultExecName() string

// An EmbedError indicates a problem with a go:embed directive.
type EmbedError struct {
	Pattern string
	Err     error
}

func (e *EmbedError) Error() string

func (e *EmbedError) Unwrap() error

// ResolveEmbed resolves //go:embed patterns and returns only the file list.
// For use by go mod vendor to find embedded files it should copy into the
// vendor directory.
// TODO(#42504): Once go mod vendor uses load.PackagesAndErrors, just
// call (*Package).ResolveEmbed
func ResolveEmbed(dir string, patterns []string) ([]string, error)

// SafeArg reports whether arg is a "safe" command-line argument,
// meaning that when it appears in a command-line, it probably
// doesn't have some special meaning other than its own name.
// Obviously args beginning with - are not safe (they look like flags).
// Less obviously, args beginning with @ are not safe (they look like
// GNU binutils flagfile specifiers, sometimes called "response files").
// To be conservative, we reject almost any arg beginning with non-alphanumeric ASCII.
// We accept leading . _ and / as likely in file system paths.
// There is a copy of this function in cmd/compile/internal/gc/noder.go.
func SafeArg(name string) bool

// LinkerDeps returns the list of linker-induced dependencies for main package p.
func LinkerDeps(p *Package) ([]string, error)

// InternalGoFiles returns the list of Go files being built for the package,
// using absolute paths.
func (p *Package) InternalGoFiles() []string

// InternalXGoFiles returns the list of Go files being built for the XTest package,
// using absolute paths.
func (p *Package) InternalXGoFiles() []string

// InternalAllGoFiles returns the list of all Go files possibly relevant for the package,
// using absolute paths. "Possibly relevant" means that files are not excluded
// due to build tags, but files with names beginning with . or _ are still excluded.
func (p *Package) InternalAllGoFiles() []string

// UsesSwig reports whether the package needs to run SWIG.
func (p *Package) UsesSwig() bool

// UsesCgo reports whether the package needs to run cgo
func (p *Package) UsesCgo() bool

// PackageList returns the list of packages in the dag rooted at roots
// as visited in a depth-first post-order traversal.
func PackageList(roots []*Package) []*Package

// TestPackageList returns the list of packages in the dag rooted at roots
// as visited in a depth-first post-order traversal, including the test
// imports of the roots. This ignores errors in test packages.
func TestPackageList(ctx context.Context, opts PackageOpts, roots []*Package) []*Package

// LoadImportWithFlags loads the package with the given import path and
// sets tool flags on that package. This function is useful loading implicit
// dependencies (like sync/atomic for coverage).
// TODO(jayconrod): delete this function and set flags automatically
// in LoadImport instead.
func LoadImportWithFlags(path, srcDir string, parent *Package, stk *ImportStack, importPos []token.Position, mode int) (*Package, *PackageError)

// LoadPackageWithFlags is the same as LoadImportWithFlags but without a parent.
// It's then guaranteed to not return an error
func LoadPackageWithFlags(path, srcDir string, stk *ImportStack, importPos []token.Position, mode int) *Package

// PackageOpts control the behavior of PackagesAndErrors and other package
// loading functions.
type PackageOpts struct {
	// IgnoreImports controls whether we ignore explicit and implicit imports
	// when loading packages.  Implicit imports are added when supporting Cgo
	// or SWIG and when linking main packages.
	IgnoreImports bool

	// ModResolveTests indicates whether calls to the module loader should also
	// resolve test dependencies of the requested packages.
	//
	// If ModResolveTests is true, then the module loader needs to resolve test
	// dependencies at the same time as packages; otherwise, the test dependencies
	// of those packages could be missing, and resolving those missing dependencies
	// could change the selected versions of modules that provide other packages.
	ModResolveTests bool

	// MainOnly is true if the caller only wants to load main packages.
	// For a literal argument matching a non-main package, a stub may be returned
	// with an error. For a non-literal argument (with "..."), non-main packages
	// are not be matched, and their dependencies may not be loaded. A warning
	// may be printed for non-literal arguments that match no main packages.
	MainOnly bool

	// AutoVCS controls whether we also load version-control metadata for main packages
	// when -buildvcs=auto (the default).
	AutoVCS bool

	// SuppressBuildInfo is true if the caller does not need p.Stale, p.StaleReason, or p.Internal.BuildInfo
	// to be populated on the package.
	SuppressBuildInfo bool

	// SuppressEmbedFiles is true if the caller does not need any embed files to be populated on the
	// package.
	SuppressEmbedFiles bool
}

// PackagesAndErrors returns the packages named by the command line arguments
// 'patterns'. If a named package cannot be loaded, PackagesAndErrors returns
// a *Package with the Error field describing the failure. If errors are found
// loading imported packages, the DepsErrors field is set. The Incomplete field
// may be set as well.
//
// To obtain a flat list of packages, use PackageList.
// To report errors loading packages, use ReportPackageErrors.
func PackagesAndErrors(ctx context.Context, opts PackageOpts, patterns []string) []*Package

// CheckPackageErrors prints errors encountered loading pkgs and their
// dependencies, then exits with a non-zero status if any errors were found.
func CheckPackageErrors(pkgs []*Package)

// GoFilesPackage creates a package for building a collection of Go files
// (typically named on the command line). The target is named p.a for
// package p or named after the first Go file for package main.
func GoFilesPackage(ctx context.Context, opts PackageOpts, gofiles []string) *Package

// PackagesAndErrorsOutsideModule is like PackagesAndErrors but runs in
// module-aware mode and ignores the go.mod file in the current directory or any
// parent directory, if there is one. This is used in the implementation of 'go
// install pkg@version' and other commands that support similar forms.
//
// modload.ForceUseModules must be true, and modload.RootMode must be NoRoot
// before calling this function.
//
// PackagesAndErrorsOutsideModule imposes several constraints to avoid
// ambiguity. All arguments must have the same version suffix (not just a suffix
// that resolves to the same version). They must refer to packages in the same
// module, which must not be std or cmd. That module is not considered the main
// module, but its go.mod file (if it has one) must not contain directives that
// would cause it to be interpreted differently if it were the main module
// (replace, exclude).
func PackagesAndErrorsOutsideModule(ctx context.Context, opts PackageOpts, args []string) ([]*Package, error)

// EnsureImport ensures that package p imports the named package.
func EnsureImport(p *Package, pkg string)

// PrepareForCoverageBuild is a helper invoked for "go install
// -cover", "go run -cover", and "go build -cover" (but not used by
// "go test -cover"). It walks through the packages being built (and
// dependencies) and marks them for coverage instrumentation when
// appropriate, and possibly adding additional deps where needed.
func PrepareForCoverageBuild(pkgs []*Package)

func SelectCoverPackages(roots []*Package, match []func(*Package) bool, op string) []*Package

// DeclareCoverVars attaches the required cover variables names
// to the files, to be used when annotating the files. This
// function only called when using legacy coverage test/build
// (e.g. GOEXPERIMENT=coverageredesign is off).
func DeclareCoverVars(p *Package, files ...string) map[string]*CoverVar
