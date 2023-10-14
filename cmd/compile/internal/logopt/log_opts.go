// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package logopt

import (
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/src"
)

type VersionHeader struct {
	Version   int    `json:"version"`
	Package   string `json:"package"`
	Goos      string `json:"goos"`
	Goarch    string `json:"goarch"`
	GcVersion string `json:"gc_version"`
	File      string `json:"file,omitempty"`
}

type DocumentURI string

type Position struct {
	Line      uint `json:"line"`
	Character uint `json:"character"`
}

// A Range in a text document expressed as (zero-based) start and end positions.
// A range is comparable to a selection in an editor. Therefore the end position is exclusive.
// If you want to specify a range that contains a line including the line ending character(s)
// then use an end position denoting the start of the next line.
type Range struct {
	/*Start defined:
	 * The range's start position
	 */
	Start Position `json:"start"`

	/*End defined:
	 * The range's end position
	 */
	End Position `json:"end"`
}

// A Location represents a location inside a resource, such as a line inside a text file.
type Location struct {
	// URI is
	URI DocumentURI `json:"uri"`

	// Range is
	Range Range `json:"range"`
}

/* DiagnosticRelatedInformation defined:
 * Represents a related message and source code location for a diagnostic. This should be
 * used to point to code locations that cause or related to a diagnostics, e.g when duplicating
 * a symbol in a scope.
 */
type DiagnosticRelatedInformation struct {

	/*Location defined:
	 * The location of this related diagnostic information.
	 */
	Location Location `json:"location"`

	/*Message defined:
	 * The message of this related diagnostic information.
	 */
	Message string `json:"message"`
}

// DiagnosticSeverity defines constants
type DiagnosticSeverity uint

const (
	/*SeverityInformation defined:
	 * Reports an information.
	 */
	SeverityInformation DiagnosticSeverity = 3
)

// DiagnosticTag defines constants
type DiagnosticTag uint

/*Diagnostic defined:
 * Represents a diagnostic, such as a compiler error or warning. Diagnostic objects
 * are only valid in the scope of a resource.
 */
type Diagnostic struct {

	/*Range defined:
	 * The range at which the message applies
	 */
	Range Range `json:"range"`

	/*Severity defined:
	 * The diagnostic's severity. Can be omitted. If omitted it is up to the
	 * client to interpret diagnostics as error, warning, info or hint.
	 */
	Severity DiagnosticSeverity `json:"severity,omitempty"`

	/*Code defined:
	 * The diagnostic's code, which usually appear in the user interface.
	 */
	Code string `json:"code,omitempty"`

	/*Source defined:
	 * A human-readable string describing the source of this
	 * diagnostic, e.g. 'typescript' or 'super lint'. It usually
	 * appears in the user interface.
	 */
	Source string `json:"source,omitempty"`

	/*Message defined:
	 * The diagnostic's message. It usually appears in the user interface
	 */
	Message string `json:"message"`

	/*Tags defined:
	 * Additional metadata about the diagnostic.
	 */
	Tags []DiagnosticTag `json:"tags,omitempty"`

	/*RelatedInformation defined:
	 * An array of related diagnostic information, e.g. when symbol-names within
	 * a scope collide all definitions can be marked via this property.
	 */
	RelatedInformation []DiagnosticRelatedInformation `json:"relatedInformation,omitempty"`
}

// A LoggedOpt is what the compiler produces and accumulates,
// to be converted to JSON for human or IDE consumption.
type LoggedOpt struct {
	pos          src.XPos
	lastPos      src.XPos
	compilerPass string
	functionName string
	what         string
	target       []interface{}
}

const (
	None logFormat = iota
	Json0
)

var Format = None

// LogJsonOption parses and validates the version,directory value attached to the -json compiler flag.
func LogJsonOption(flagValue string)

// NewLoggedOpt allocates a new LoggedOpt, to later be passed to either NewLoggedOpt or LogOpt as "args".
// Pos is the source position (including inlining), what is the message, pass is which pass created the message,
// funcName is the name of the function
// A typical use for this to accumulate an explanation for a missed optimization, for example, why did something escape?
func NewLoggedOpt(pos, lastPos src.XPos, what, pass, funcName string, args ...interface{}) *LoggedOpt

// LogOpt logs information about a (usually missed) optimization performed by the compiler.
// Pos is the source position (including inlining), what is the message, pass is which pass created the message,
// funcName is the name of the function.
func LogOpt(pos src.XPos, what, pass, funcName string, args ...interface{})

// LogOptRange is the same as LogOpt, but includes the ability to express a range of positions,
// not just a point.
func LogOptRange(pos, lastPos src.XPos, what, pass, funcName string, args ...interface{})

// Enabled returns whether optimization logging is enabled.
func Enabled() bool

// FlushLoggedOpts flushes all the accumulated optimization log entries.
func FlushLoggedOpts(ctxt *obj.Link, slashPkgPath string)
