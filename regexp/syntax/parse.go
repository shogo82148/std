// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

// An Error describes a failure to parse a regular expression
// and gives the offending expression.
type Error struct {
	Code ErrorCode
	Expr string
}

func (e *Error) Error() string

// An ErrorCode describes a failure to parse a regular expression.
type ErrorCode string

const (
	// Unexpected error
	ErrInternalError ErrorCode = "regexp/syntax: internal error"

	// Parse errors
	ErrInvalidCharClass      ErrorCode = "invalid character class"
	ErrInvalidCharRange      ErrorCode = "invalid character class range"
	ErrInvalidEscape         ErrorCode = "invalid escape sequence"
	ErrInvalidNamedCapture   ErrorCode = "invalid named capture"
	ErrInvalidPerlOp         ErrorCode = "invalid or unsupported Perl syntax"
	ErrInvalidRepeatOp       ErrorCode = "invalid nested repetition operator"
	ErrInvalidRepeatSize     ErrorCode = "invalid repeat count"
	ErrInvalidUTF8           ErrorCode = "invalid UTF-8"
	ErrMissingBracket        ErrorCode = "missing closing ]"
	ErrMissingParen          ErrorCode = "missing closing )"
	ErrMissingRepeatArgument ErrorCode = "missing argument to repetition operator"
	ErrTrailingBackslash     ErrorCode = "trailing backslash at end of expression"
	ErrUnexpectedParen       ErrorCode = "unexpected )"
	ErrNestingDepth          ErrorCode = "expression nests too deeply"
)

func (e ErrorCode) String() string

// Flags control the behavior of the parser and record information about regexp context.
type Flags uint16

const (
	FoldCase      Flags = 1 << iota
	Literal
	ClassNL
	DotNL
	OneLine
	NonGreedy
	PerlX
	UnicodeGroups
	WasDollar
	Simple

	MatchNL = ClassNL | DotNL

	Perl        = ClassNL | OneLine | PerlX | UnicodeGroups
	POSIX Flags = 0
)

// Pseudo-ops for parsing stack.

// maxHeight is the maximum height of a regexp parse tree.
// It is somewhat arbitrarily chosen, but the idea is to be large enough
// that no one will actually hit in real use but at the same time small enough
// that recursion on the Regexp tree will not hit the 1GB Go stack limit.
// The maximum amount of stack for a single recursive frame is probably
// closer to 1kB, so this could potentially be raised, but it seems unlikely
// that people have regexps nested even this deeply.
// We ran a test on Google's C++ code base and turned up only
// a single use case with depth > 100; it had depth 128.
// Using depth 1000 should be plenty of margin.
// As an optimization, we don't even bother calculating heights
// until we've allocated at least maxHeight Regexp structures.

// maxSize is the maximum size of a compiled regexp in Insts.
// It too is somewhat arbitrarily chosen, but the idea is to be large enough
// to allow significant regexps while at the same time small enough that
// the compiled form will not take up too much memory.
// 128 MB is enough for a 3.3 million Inst structures, which roughly
// corresponds to a 3.3 MB regexp.

// maxRunes is the maximum number of runes allowed in a regexp tree
// counting the runes in all the nodes.
// Ignoring character classes p.numRunes is always less than the length of the regexp.
// Character classes can make it much larger: each \pL adds 1292 runes.
// 128 MB is enough for 32M runes, which is over 26k \pL instances.
// Note that repetitions do not make copies of the rune slices,
// so \pL{1000} is only one rune slice, not 1000.
// We could keep a cache of character classes we've seen,
// so that all the \pL we see use the same rune list,
// but that doesn't remove the problem entirely:
// consider something like [\pL01234][\pL01235][\pL01236]...[\pL^&*()].
// And because the Rune slice is exposed directly in the Regexp,
// there is not an opportunity to change the representation to allow
// partial sharing between different character classes.
// So the limit is the best we can do.

// Parse parses a regular expression string s, controlled by the specified
// Flags, and returns a regular expression parse tree. The syntax is
// described in the top-level comment.
func Parse(s string, flags Flags) (*Regexp, error)

// ranges implements sort.Interface on a []rune.
// The choice of receiver type definition is strange
// but avoids an allocation since we already have
// a *[]rune.
