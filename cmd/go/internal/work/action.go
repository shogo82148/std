// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Action graph creation (planning).

package work

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/sync"

	"github.com/shogo82148/std/cmd/go/internal/cache"
	"github.com/shogo82148/std/cmd/go/internal/load"
	"github.com/shogo82148/std/cmd/go/internal/trace"
)

// A Builder holds global state about a build.
// It does not hold per-package state, because we
// build packages in parallel, and the builder is shared.
type Builder struct {
	WorkDir            string
	actionCache        map[cacheKey]*Action
	flagCache          map[[2]string]bool
	gccCompilerIDCache map[string]cache.ActionID

	IsCmdList           bool
	NeedError           bool
	NeedExport          bool
	NeedCompiledGoFiles bool
	AllowErrors         bool

	objdirSeq int
	pkgSeq    int

	backgroundSh *Shell

	exec      sync.Mutex
	readySema chan bool
	ready     actionQueue

	id           sync.Mutex
	toolIDCache  map[string]string
	buildIDCache map[string]string
}

// An Actor runs an action.
type Actor interface {
	Act(*Builder, context.Context, *Action) error
}

// An ActorFunc is an Actor that calls the function.
type ActorFunc func(*Builder, context.Context, *Action) error

func (f ActorFunc) Act(b *Builder, ctx context.Context, a *Action) error

// An Action represents a single action in the action graph.
type Action struct {
	Mode       string
	Package    *load.Package
	Deps       []*Action
	Actor      Actor
	IgnoreFail bool
	TestOutput *bytes.Buffer
	Args       []string

	triggers []*Action

	buggyInstall bool

	TryCache func(*Builder, *Action) bool

	CacheExecutable bool

	// Generated files, directories.
	Objdir   string
	Target   string
	built    string
	actionID cache.ActionID
	buildID  string

	VetxOnly  bool
	needVet   bool
	needBuild bool
	vetCfg    *vetConfig
	output    []byte

	sh *Shell

	// Execution state.
	pending      int
	priority     int
	Failed       *Action
	json         *actionJSON
	nonGoOverlay map[string]string
	traceSpan    *trace.Span
}

// BuildActionID returns the action ID section of a's build ID.
func (a *Action) BuildActionID() string

// BuildContentID returns the content ID section of a's build ID.
func (a *Action) BuildContentID() string

// BuildID returns a's build ID.
func (a *Action) BuildID() string

// BuiltTarget returns the actual file that was built. This differs
// from Target when the result was cached.
func (a *Action) BuiltTarget() string

// BuildMode specifies the build mode:
// are we just building things or also installing the results?
type BuildMode int

const (
	ModeBuild BuildMode = iota
	ModeInstall
	ModeBuggyInstall

	ModeVetOnly = 1 << 8
)

// NewBuilder returns a new Builder ready for use.
//
// If workDir is the empty string, NewBuilder creates a WorkDir if needed
// and arranges for it to be removed in case of an unclean exit.
// The caller must Close the builder explicitly to clean up the WorkDir
// before a clean exit.
func NewBuilder(workDir string) *Builder

func (b *Builder) Close() error

func CheckGOOSARCHPair(goos, goarch string) error

// NewObjdir returns the name of a fresh object directory under b.WorkDir.
// It is up to the caller to call b.Mkdir on the result at an appropriate time.
// The result ends in a slash, so that file names in that directory
// can be constructed with direct string addition.
//
// NewObjdir must be called only from a single goroutine at a time,
// so it is safe to call during action graph construction, but it must not
// be called during action graph execution.
func (b *Builder) NewObjdir() string

// AutoAction returns the "right" action for go build or go install of p.
func (b *Builder) AutoAction(mode, depMode BuildMode, p *load.Package) *Action

// CompileAction returns the action for compiling and possibly installing
// (according to mode) the given package. The resulting action is only
// for building packages (archives), never for linking executables.
// depMode is the action (build or install) to use when building dependencies.
// To turn package main into an executable, call b.Link instead.
func (b *Builder) CompileAction(mode, depMode BuildMode, p *load.Package) *Action

// VetAction returns the action for running go vet on package p.
// It depends on the action for compiling p.
// If the caller may be causing p to be installed, it is up to the caller
// to make sure that the install depends on (runs after) vet.
func (b *Builder) VetAction(mode, depMode BuildMode, p *load.Package) *Action

// LinkAction returns the action for linking p into an executable
// and possibly installing the result (according to mode).
// depMode is the action (build or install) to use when compiling dependencies.
func (b *Builder) LinkAction(mode, depMode BuildMode, p *load.Package) *Action
