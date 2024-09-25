// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modload

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/sync"

	"golang.org/x/mod/modfile"
	"golang.org/x/mod/module"
)

// Variables set by other packages.
//
// TODO(#40775): See if these can be plumbed as explicit parameters.
var (
	// RootMode determines whether a module root is needed.
	RootMode Root

	// ForceUseModules may be set to force modules to be enabled when
	// GO111MODULE=auto or to report an error when GO111MODULE=off.
	ForceUseModules bool

	// ExplicitWriteGoMod prevents LoadPackages, ListModules, and other functions
	// from updating go.mod and go.sum or reporting errors when updates are
	// needed. A package should set this if it would cause go.mod to be written
	// multiple times (for example, 'go get' calls LoadPackages multiple times) or
	// if it needs some other operation to be successful before go.mod and go.sum
	// can be written (for example, 'go mod download' must download modules before
	// adding sums to go.sum). Packages that set this are responsible for calling
	// WriteGoMod explicitly.
	ExplicitWriteGoMod bool
)

// EnterModule resets MainModules and requirements to refer to just this one module.
func EnterModule(ctx context.Context, enterModroot string)

type MainModuleSet struct {
	// versions are the module.Version values of each of the main modules.
	// For each of them, the Path fields are ordinary module paths and the Version
	// fields are empty strings.
	// versions is clipped (len=cap).
	versions []module.Version

	// modRoot maps each module in versions to its absolute filesystem path.
	modRoot map[module.Version]string

	// pathPrefix is the path prefix for packages in the module, without a trailing
	// slash. For most modules, pathPrefix is just version.Path, but the
	// standard-library module "std" has an empty prefix.
	pathPrefix map[module.Version]string

	// inGorootSrc caches whether modRoot is within GOROOT/src.
	// The "std" module is special within GOROOT/src, but not otherwise.
	inGorootSrc map[module.Version]bool

	modFiles map[module.Version]*modfile.File

	tools map[string]bool

	modContainingCWD module.Version

	workFile *modfile.WorkFile

	workFileReplaceMap map[module.Version]module.Version
	// highest replaced version of each module path; empty string for wildcard-only replacements
	highestReplaced map[string]string

	indexMu sync.Mutex
	indices map[module.Version]*modFileIndex
}

func (mms *MainModuleSet) PathPrefix(m module.Version) string

// Versions returns the module.Version values of each of the main modules.
// For each of them, the Path fields are ordinary module paths and the Version
// fields are empty strings.
// Callers should not modify the returned slice.
func (mms *MainModuleSet) Versions() []module.Version

// Tools returns the tools defined by all the main modules.
// The key is the absolute package path of the tool.
func (mms *MainModuleSet) Tools() map[string]bool

func (mms *MainModuleSet) Contains(path string) bool

func (mms *MainModuleSet) ModRoot(m module.Version) string

func (mms *MainModuleSet) InGorootSrc(m module.Version) bool

func (mms *MainModuleSet) GetSingleIndexOrNil() *modFileIndex

func (mms *MainModuleSet) Index(m module.Version) *modFileIndex

func (mms *MainModuleSet) SetIndex(m module.Version, index *modFileIndex)

func (mms *MainModuleSet) ModFile(m module.Version) *modfile.File

func (mms *MainModuleSet) WorkFile() *modfile.WorkFile

func (mms *MainModuleSet) Len() int

// ModContainingCWD returns the main module containing the working directory,
// or module.Version{} if none of the main modules contain the working
// directory.
func (mms *MainModuleSet) ModContainingCWD() module.Version

func (mms *MainModuleSet) HighestReplaced() map[string]string

// GoVersion returns the go version set on the single module, in module mode,
// or the go.work file in workspace mode.
func (mms *MainModuleSet) GoVersion() string

// Godebugs returns the godebug lines set on the single module, in module mode,
// or on the go.work file in workspace mode.
// The caller must not modify the result.
func (mms *MainModuleSet) Godebugs() []*modfile.Godebug

// Toolchain returns the toolchain set on the single module, in module mode,
// or the go.work file in workspace mode.
func (mms *MainModuleSet) Toolchain() string

func (mms *MainModuleSet) WorkFileReplaceMap() map[module.Version]module.Version

var MainModules *MainModuleSet

type Root int

const (
	// AutoRoot is the default for most commands. modload.Init will look for
	// a go.mod file in the current directory or any parent. If none is found,
	// modules may be disabled (GO111MODULE=auto) or commands may run in a
	// limited module mode.
	AutoRoot Root = iota

	// NoRoot is used for commands that run in module mode and ignore any go.mod
	// file the current directory or in parent directories.
	NoRoot

	// NeedRoot is used for commands that must run in module mode and don't
	// make sense without a main module.
	NeedRoot
)

// ModFile returns the parsed go.mod file.
//
// Note that after calling LoadPackages or LoadModGraph,
// the require statements in the modfile.File are no longer
// the source of truth and will be ignored: edits made directly
// will be lost at the next call to WriteGoMod.
// To make permanent changes to the require statements
// in go.mod, edit it before loading.
func ModFile() *modfile.File

func BinDir() string

// InitWorkfile initializes the workFilePath variable for commands that
// operate in workspace mode. It should not be called by other commands,
// for example 'go mod tidy', that don't operate in workspace mode.
func InitWorkfile()

// FindGoWork returns the name of the go.work file for this command,
// or the empty string if there isn't one.
// Most code should use Init and Enabled rather than use this directly.
// It is exported mainly for Go toolchain switching, which must process
// the go.work very early at startup.
func FindGoWork(wd string) string

// WorkFilePath returns the absolute path of the go.work file, or "" if not in
// workspace mode. WorkFilePath must be called after InitWorkfile.
func WorkFilePath() string

// Reset clears all the initialized, cached state about the use of modules,
// so that we can start over.
func Reset()

// Init determines whether module mode is enabled, locates the root of the
// current module (if any), sets environment variables for Git subprocesses, and
// configures the cfg, codehost, load, modfetch, and search packages for use
// with modules.
func Init()

// WillBeEnabled checks whether modules should be enabled but does not
// initialize modules by installing hooks. If Init has already been called,
// WillBeEnabled returns the same result as Enabled.
//
// This function is needed to break a cycle. The main package needs to know
// whether modules are enabled in order to install the module or GOPATH version
// of 'go get', but Init reads the -modfile flag in 'go get', so it shouldn't
// be called until the command is installed and flags are parsed. Instead of
// calling Init and Enabled, the main package can call this function.
func WillBeEnabled() bool

// FindGoMod returns the name of the go.mod file for this command,
// or the empty string if there isn't one.
// Most code should use Init and Enabled rather than use this directly.
// It is exported mainly for Go toolchain switching, which must process
// the go.mod very early at startup.
func FindGoMod(wd string) string

// Enabled reports whether modules are (or must be) enabled.
// If modules are enabled but there is no main module, Enabled returns true
// and then the first use of module information will call die
// (usually through MustModRoot).
func Enabled() bool

func VendorDir() string

// HasModRoot reports whether a main module is present.
// HasModRoot may return false even if Enabled returns true: for example, 'get'
// does not require a main module.
func HasModRoot() bool

// MustHaveModRoot checks that a main module or main modules are present,
// and calls base.Fatalf if there are no main modules.
func MustHaveModRoot()

// ModFilePath returns the path that would be used for the go.mod
// file, if in module mode. ModFilePath calls base.Fatalf if there is no main
// module, even if -modfile is set.
func ModFilePath() string

var ErrNoModRoot = errors.New("go.mod file not found in current directory or any parent directory; see 'go help modules'")

// ReadWorkFile reads and parses the go.work file at the given path.
func ReadWorkFile(path string) (*modfile.WorkFile, error)

// WriteWorkFile cleans and writes out the go.work file to the given path.
func WriteWorkFile(path string, wf *modfile.WorkFile) error

// UpdateWorkGoVersion updates the go line in wf to be at least goVers,
// reporting whether it changed the file.
func UpdateWorkGoVersion(wf *modfile.WorkFile, goVers string) (changed bool)

// UpdateWorkFile updates comments on directory directives in the go.work
// file to include the associated module path.
func UpdateWorkFile(wf *modfile.WorkFile)

// LoadModFile sets Target and, if there is a main module, parses the initial
// build list from its go.mod file.
//
// LoadModFile may make changes in memory, like adding a go directive and
// ensuring requirements are consistent. The caller is responsible for ensuring
// those changes are written to disk by calling LoadPackages or ListModules
// (unless ExplicitWriteGoMod is set) or by calling WriteGoMod directly.
//
// As a side-effect, LoadModFile may change cfg.BuildMod to "vendor" if
// -mod wasn't set explicitly and automatic vendoring should be enabled.
//
// If LoadModFile or CreateModFile has already been called, LoadModFile returns
// the existing in-memory requirements (rather than re-reading them from disk).
//
// LoadModFile checks the roots of the module graph for consistency with each
// other, but unlike LoadModGraph does not load the full module graph or check
// it for global consistency. Most callers outside of the modload package should
// use LoadModGraph instead.
func LoadModFile(ctx context.Context) *Requirements

// CreateModFile initializes a new module by creating a go.mod file.
//
// If modPath is empty, CreateModFile will attempt to infer the path from the
// directory location within GOPATH.
//
// If a vendoring configuration file is present, CreateModFile will attempt to
// translate it to go.mod directives. The resulting build list may not be
// exactly the same as in the legacy configuration (for example, we can't get
// packages at multiple versions from the same module).
func CreateModFile(ctx context.Context, modPath string)

// AllowMissingModuleImports allows import paths to be resolved to modules
// when there is no module root. Normally, this is forbidden because it's slow
// and there's no way to make the result reproducible, but some commands
// like 'go get' are expected to do this.
//
// This function affects the default cfg.BuildMod when outside of a module,
// so it can only be called prior to Init.
func AllowMissingModuleImports()

// WriteOpts control the behavior of WriteGoMod.
type WriteOpts struct {
	DropToolchain     bool
	ExplicitToolchain bool

	AddTools  []string
	DropTools []string

	// TODO(bcmills): Make 'go mod tidy' update the go version in the Requirements
	// instead of writing directly to the modfile.File
	TidyWroteGo bool
}

// WriteGoMod writes the current build list back to go.mod.
func WriteGoMod(ctx context.Context, opts WriteOpts) error

// UpdateGoModFromReqs returns a modified go.mod file using the current
// requirements. It does not commit these changes to disk.
func UpdateGoModFromReqs(ctx context.Context, opts WriteOpts) (before, after []byte, modFile *modfile.File, err error)

func CheckGodebug(verb, k, v string) error
