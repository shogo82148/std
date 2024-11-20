// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cfg holds configuration shared by multiple parts
// of the go command.
package cfg

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/go/build"
	"github.com/shogo82148/std/internal/buildcfg"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/runtime"
)

// Global build parameters (used during package load)
var (
	Goos   = envOr("GOOS", build.Default.GOOS)
	Goarch = envOr("GOARCH", build.Default.GOARCH)

	ExeSuffix = exeSuffix()

	// ModulesEnabled specifies whether the go command is running
	// in module-aware mode (as opposed to GOPATH mode).
	// It is equal to modload.Enabled, but not all packages can import modload.
	ModulesEnabled bool
)

// ToolExeSuffix returns the suffix for executables installed
// in build.ToolDir.
func ToolExeSuffix() string

// These are general "build flags" used by build and other commands.
var (
	BuildA                 bool
	BuildBuildmode         string
	BuildBuildvcs          = "auto"
	BuildContext           = defaultContext()
	BuildMod               string
	BuildModExplicit       bool
	BuildModReason         string
	BuildLinkshared        bool
	BuildMSan              bool
	BuildASan              bool
	BuildCover             bool
	BuildCoverMode         string
	BuildCoverPkg          []string
	BuildJSON              bool
	BuildN                 bool
	BuildO                 string
	BuildP                 = runtime.GOMAXPROCS(0)
	BuildPGO               string
	BuildPkgdir            string
	BuildRace              bool
	BuildToolexec          []string
	BuildToolchainName     string
	BuildToolchainCompiler func() string
	BuildToolchainLinker   func() string
	BuildTrimpath          bool
	BuildV                 bool
	BuildWork              bool
	BuildX                 bool

	ModCacheRW bool
	ModFile    string

	CmdName string

	DebugActiongraph  string
	DebugTrace        string
	DebugRuntimeTrace string

	// GoPathError is set when GOPATH is not set. it contains an
	// explanation why GOPATH is unset.
	GoPathError   string
	GOPATHChanged bool
	CGOChanged    bool
)

// SetGOROOT sets GOROOT and associated variables to the given values.
//
// If isTestGo is true, build.ToolDir is set based on the TESTGO_GOHOSTOS and
// TESTGO_GOHOSTARCH environment variables instead of runtime.GOOS and
// runtime.GOARCH.
func SetGOROOT(goroot string, isTestGo bool)

// Experiment configuration.
var (
	// RawGOEXPERIMENT is the GOEXPERIMENT value set by the user.
	RawGOEXPERIMENT = envOr("GOEXPERIMENT", buildcfg.DefaultGOEXPERIMENT)
	// CleanGOEXPERIMENT is the minimal GOEXPERIMENT value needed to reproduce the
	// experiments enabled by RawGOEXPERIMENT.
	CleanGOEXPERIMENT = RawGOEXPERIMENT

	Experiment    *buildcfg.ExperimentFlags
	ExperimentErr error
)

// An EnvVar is an environment variable Name=Value.
type EnvVar struct {
	Name    string
	Value   string
	Changed bool
}

// OrigEnv is the original environment of the program at startup.
var OrigEnv []string

// CmdEnv is the new environment for running go tool commands.
// User binaries (during go test or go run) are run with OrigEnv,
// not CmdEnv.
var CmdEnv []EnvVar

// EnvFile returns the name of the Go environment configuration file,
// and reports whether the effective value differs from the default.
func EnvFile() (string, bool, error)

// Getenv gets the value for the configuration key.
// It consults the operating system environment
// and then the go/env file.
// If Getenv is called for a key that cannot be set
// in the go/env file (for example GODEBUG), it panics.
// This ensures that CanGetenv is accurate, so that
// 'go env -w' stays in sync with what Getenv can retrieve.
func Getenv(key string) string

// CanGetenv reports whether key is a valid go/env configuration key.
func CanGetenv(key string) bool

var (
	GOROOT string

	// Either empty or produced by filepath.Join(GOROOT, …).
	GOROOTbin string
	GOROOTpkg string
	GOROOTsrc string

	GOBIN                         = Getenv("GOBIN")
	GOMODCACHE, GOMODCACHEChanged = EnvOrAndChanged("GOMODCACHE", gopathDir("pkg/mod"))

	// Used in envcmd.MkEnv and build ID computations.
	GOARM64   = EnvOrAndChanged("GOARM64", buildcfg.DefaultGOARM64)
	GOARM     = EnvOrAndChanged("GOARM", buildcfg.DefaultGOARM)
	GO386     = EnvOrAndChanged("GO386", buildcfg.DefaultGO386)
	GOAMD64   = EnvOrAndChanged("GOAMD64", buildcfg.DefaultGOAMD64)
	GOMIPS    = EnvOrAndChanged("GOMIPS", buildcfg.DefaultGOMIPS)
	GOMIPS64  = EnvOrAndChanged("GOMIPS64", buildcfg.DefaultGOMIPS64)
	GOPPC64   = EnvOrAndChanged("GOPPC64", buildcfg.DefaultGOPPC64)
	GORISCV64 = EnvOrAndChanged("GORISCV64", buildcfg.DefaultGORISCV64)
	GOWASM    = EnvOrAndChanged("GOWASM", fmt.Sprint(buildcfg.GOWASM))

	GOFIPS140, GOFIPS140Changed = EnvOrAndChanged("GOFIPS140", buildcfg.GOFIPS140)
	GOPROXY, GOPROXYChanged     = EnvOrAndChanged("GOPROXY", "")
	GOSUMDB, GOSUMDBChanged     = EnvOrAndChanged("GOSUMDB", "")
	GOPRIVATE                   = Getenv("GOPRIVATE")
	GONOPROXY, GONOPROXYChanged = EnvOrAndChanged("GONOPROXY", GOPRIVATE)
	GONOSUMDB, GONOSUMDBChanged = EnvOrAndChanged("GONOSUMDB", GOPRIVATE)
	GOINSECURE                  = Getenv("GOINSECURE")
	GOVCS                       = Getenv("GOVCS")
	GOAUTH, GOAUTHChanged       = EnvOrAndChanged("GOAUTH", "netrc")
)

// EnvOrAndChanged returns the environment variable value
// and reports whether it differs from the default value.
func EnvOrAndChanged(name, def string) (v string, changed bool)

var SumdbDir = gopathDir("pkg/sumdb")

// GetArchEnv returns the name and setting of the
// GOARCH-specific architecture environment variable.
// If the current architecture has no GOARCH-specific variable,
// GetArchEnv returns empty key and value.
func GetArchEnv() (key, val string, changed bool)

// WithBuildXWriter returns a Context in which BuildX output is written
// to given io.Writer.
func WithBuildXWriter(ctx context.Context, xLog io.Writer) context.Context

// BuildXWriter returns nil if BuildX is false, or
// the writer to which BuildX output should be written otherwise.
func BuildXWriter(ctx context.Context) (io.Writer, bool)
