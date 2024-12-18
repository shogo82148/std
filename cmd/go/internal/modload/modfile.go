// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modload

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/errors"

	"github.com/shogo82148/std/golang.org/x/mod/modfile"
	"github.com/shogo82148/std/golang.org/x/mod/module"
)

// ReadModFile reads and parses the mod file at gomod. ReadModFile properly applies the
// overlay, locks the file while reading, and applies fix, if applicable.
func ReadModFile(gomod string, fix modfile.VersionFixer) (data []byte, f *modfile.File, err error)

// CheckAllowed returns an error equivalent to ErrDisallowed if m is excluded by
// the main module's go.mod or retracted by its author. Most version queries use
// this to filter out versions that should not be used.
func CheckAllowed(ctx context.Context, m module.Version) error

// ErrDisallowed is returned by version predicates passed to Query and similar
// functions to indicate that a version should not be considered.
var ErrDisallowed = errors.New("disallowed module version")

// CheckExclusions returns an error equivalent to ErrDisallowed if module m is
// excluded by the main module's go.mod file.
func CheckExclusions(ctx context.Context, m module.Version) error

// CheckRetractions returns an error if module m has been retracted by
// its author.
func CheckRetractions(ctx context.Context, m module.Version) (err error)

type ModuleRetractedError struct {
	Rationale []string
}

func (e *ModuleRetractedError) Error() string

func (e *ModuleRetractedError) Is(err error) bool

// ShortMessage returns a string from go.mod (for example, a retraction
// rationale or deprecation message) that is safe to print in a terminal.
//
// If the given string is empty, ShortMessage returns the given default. If the
// given string is too long or contains non-printable characters, ShortMessage
// returns a hard-coded string.
func ShortMessage(message, emptyDefault string) string

// CheckDeprecation returns a deprecation message from the go.mod file of the
// latest version of the given module. Deprecation messages are comments
// before or on the same line as the module directives that start with
// "Deprecated:" and run until the end of the paragraph.
//
// CheckDeprecation returns an error if the message can't be loaded.
// CheckDeprecation returns "", nil if there is no deprecation message.
func CheckDeprecation(ctx context.Context, m module.Version) (deprecation string, err error)

// Replacement returns the replacement for mod, if any. If the path in the
// module.Version is relative it's relative to the single main module outside
// workspace mode, or the workspace's directory in workspace mode.
func Replacement(mod module.Version) module.Version

// ToDirectoryPath adds a prefix if necessary so that path in unambiguously
// an absolute path or a relative path starting with a '.' or '..'
// path component.
func ToDirectoryPath(path string) string
