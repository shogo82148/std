// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package regexp implements regular expression search.
//
// The syntax of the regular expressions accepted is the same
// general syntax used by Perl, Python, and other languages.
// More precisely, it is the syntax accepted by RE2 and described at
// https://golang.org/s/re2syntax, except for \C.
// For an overview of the syntax, see the [regexp/syntax] package.
//
// The regexp implementation provided by this package is
// guaranteed to run in time linear in the size of the input.
// (This is a property not guaranteed by most open source
// implementations of regular expressions.) For more information
// about this property, see https://swtch.com/~rsc/regexp/regexp1.html
// or any book about automata theory.
//
// All characters are UTF-8-encoded code points.
// Following [utf8.DecodeRune], each byte of an invalid UTF-8 sequence
// is treated as if it encoded utf8.RuneError (U+FFFD).
//
// There are 24 methods of [Regexp] that match a regular expression and identify
// the matched text. Their names are matched by this regular expression:
//
//	(All|Find|FindAll)(String)?(Submatch)?(Index)?
//
// The ‘All’ variants return an iterator over successive non-overlapping
// matches of the entire expression. The ‘FindAll’ variants return a slice
// of those matches instead. Empty matches abutting a preceding
// match are ignored. The ‘FindAll’ variants take an extra integer argument, n.
// If n >= 0, the function returns at most n matches/submatches;
// otherwise, it returns all of them.
//
// The ‘Find’ variants return only the first match that All or FindAll would return.
//
// If ‘String’ is present, the argument is a string; otherwise it is a []byte.
//
// By default, each returned match is denoted by the substring matching the
// regular expression, of type string or []byte according to the type of the argument.
// If ‘Submatch’ is present, each match is represented instead by a slice of
// the substrings matching the regular expression's parenthesized subexpressions
// (also known as capturing groups), numbered from left to right in order of opening
// parenthesis. Submatch 0 is the match of the entire expression, submatch 1 is
// the match of the first parenthesized subexpression, and so on.
// If ‘Index’ is present, each substring is instead denoted by a pair of byte indexes
// within the input string. If an index is negative or substring is nil, it means that
// the subexpression did not match any string in the input. For ‘String’ versions,
// an empty string means either no match or an empty match.
//
// There is also a subset of the methods that can be applied to text read from
// an [io.RuneReader]: [Regexp.MatchReader], [Regexp.FindReaderIndex],
// [Regexp.FindReaderSubmatchIndex].
// Note that regular expression matches may need to
// examine text beyond the text returned by a match, so the methods that
// match text from an [io.RuneReader] may read arbitrarily far into the input
// before returning.
//
// (There are a few other methods that do not match this pattern.)
package regexp

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/regexp/syntax"
)

// Regexp is the representation of a compiled regular expression.
// A Regexp is safe for concurrent use by multiple goroutines,
// except for configuration methods, such as [Regexp.Longest].
type Regexp struct {
	expr           string
	prog           *syntax.Prog
	onepass        *onePassProg
	numSubexp      int
	maxBitStateLen int
	subexpNames    []string
	prefix         string
	prefixBytes    []byte
	prefixRune     rune
	prefixEnd      uint32
	mpool          int
	matchcap       int
	prefixComplete bool
	cond           syntax.EmptyOp
	minInputLen    int

	// This field can be modified by the Longest method,
	// but it is otherwise read-only.
	longest bool
}

// String returns the source text used to compile the regular expression.
func (re *Regexp) String() string

// Copy returns a new [Regexp] object copied from re.
// Calling [Regexp.Longest] on one copy does not affect another.
//
// Deprecated: In earlier releases, when using a [Regexp] in multiple goroutines,
// giving each goroutine its own copy helped to avoid lock contention.
// As of Go 1.12, using Copy is no longer necessary to avoid lock contention.
// Copy may still be appropriate if the reason for its use is to make
// two copies with different [Regexp.Longest] settings.
func (re *Regexp) Copy() *Regexp

// Compile parses a regular expression and returns, if successful,
// a [Regexp] object that can be used to match against text.
//
// When matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses the one that a backtracking search would have found first.
// This so-called leftmost-first matching is the same semantics
// that Perl, Python, and other implementations use, although this
// package implements it without the expense of backtracking.
// For POSIX leftmost-longest matching, see [CompilePOSIX].
func Compile(expr string) (*Regexp, error)

// CompilePOSIX is like [Compile] but restricts the regular expression
// to POSIX ERE (egrep) syntax and changes the match semantics to
// leftmost-longest.
//
// That is, when matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses a match that is as long as possible.
// This so-called leftmost-longest matching is the same semantics
// that early regular expression implementations used and that POSIX
// specifies.
//
// However, there can be multiple leftmost-longest matches, with different
// submatch choices, and here this package diverges from POSIX.
// Among the possible leftmost-longest matches, this package chooses
// the one that a backtracking search would have found first, while POSIX
// specifies that the match be chosen to maximize the length of the first
// subexpression, then the second, and so on from left to right.
// The POSIX rule is computationally prohibitive and not even well-defined.
// See https://swtch.com/~rsc/regexp/regexp2.html#posix for details.
func CompilePOSIX(expr string) (*Regexp, error)

// Longest makes future searches prefer the leftmost-longest match.
// That is, when matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses a match that is as long as possible.
// This method modifies the [Regexp] and may not be called concurrently
// with any other methods.
func (re *Regexp) Longest()

// MustCompile is like [Compile] but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular
// expressions.
func MustCompile(str string) *Regexp

// MustCompilePOSIX is like [CompilePOSIX] but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular
// expressions.
func MustCompilePOSIX(str string) *Regexp

// NumSubexp returns the number of parenthesized subexpressions in this [Regexp].
func (re *Regexp) NumSubexp() int

// SubexpNames returns the names of the parenthesized subexpressions
// in this [Regexp]. The name for the first sub-expression is names[1],
// so that if m is a match slice, the name for m[i] is SubexpNames()[i].
// Since the Regexp as a whole cannot be named, names[0] is always
// the empty string. The slice should not be modified.
func (re *Regexp) SubexpNames() []string

// SubexpIndex returns the index of the first subexpression with the given name,
// or -1 if there is no subexpression with that name.
//
// Note that multiple subexpressions can be written using the same name, as in
// (?P<bob>a+)(?P<bob>b+), which declares two subexpressions named "bob".
// In this case, SubexpIndex returns the index of the leftmost such subexpression
// in the regular expression.
func (re *Regexp) SubexpIndex(name string) int

// LiteralPrefix returns a literal string that must begin any match
// of the regular expression re. It returns the boolean true if the
// literal string comprises the entire regular expression.
func (re *Regexp) LiteralPrefix() (prefix string, complete bool)

// MatchReader reports whether the text returned by the [io.RuneReader]
// contains any match of the regular expression re.
func (re *Regexp) MatchReader(r io.RuneReader) bool

// MatchString reports whether the string s
// contains any match of the regular expression re.
func (re *Regexp) MatchString(s string) bool

// Match reports whether the byte slice b
// contains any match of the regular expression re.
func (re *Regexp) Match(b []byte) bool

// MatchReader reports whether the text returned by the [io.RuneReader]
// contains any match of the regular expression pattern.
// More complicated queries need to use [Compile] and the full [Regexp] interface.
func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)

// MatchString reports whether the string s
// contains any match of the regular expression pattern.
// More complicated queries need to use [Compile] and the full [Regexp] interface.
func MatchString(pattern string, s string) (matched bool, err error)

// Match reports whether the byte slice b
// contains any match of the regular expression pattern.
// More complicated queries need to use [Compile] and the full [Regexp] interface.
func Match(pattern string, b []byte) (matched bool, err error)

// ReplaceAllString returns a copy of src, replacing matches of the [Regexp]
// with the replacement string repl.
// Inside repl, $ signs are interpreted as in [Regexp.Expand].
func (re *Regexp) ReplaceAllString(src, repl string) string

// ReplaceAllLiteralString returns a copy of src, replacing matches of the [Regexp]
// with the replacement string repl. The replacement repl is substituted directly,
// without using [Regexp.Expand].
func (re *Regexp) ReplaceAllLiteralString(src, repl string) string

// ReplaceAllStringFunc returns a copy of src in which all matches of the
// [Regexp] have been replaced by the return value of function repl applied
// to the matched substring. The replacement returned by repl is substituted
// directly, without using [Regexp.Expand].
func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string

// ReplaceAll returns a copy of src, replacing matches of the [Regexp]
// with the replacement text repl.
// Inside repl, $ signs are interpreted as in [Regexp.Expand].
func (re *Regexp) ReplaceAll(src, repl []byte) []byte

// ReplaceAllLiteral returns a copy of src, replacing matches of the [Regexp]
// with the replacement bytes repl. The replacement repl is substituted directly,
// without using [Regexp.Expand].
func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte

// ReplaceAllFunc returns a copy of src in which all matches of the
// [Regexp] have been replaced by the return value of function repl applied
// to the matched byte slice. The replacement returned by repl is substituted
// directly, without using [Regexp.Expand].
func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte

// QuoteMeta returns a string that escapes all regular expression metacharacters
// inside the argument text; the returned string is a regular expression matching
// the literal text.
func QuoteMeta(s string) string

// Find returns the text of the leftmost match for re in b.
// The return value is nil for no match.
func (re *Regexp) Find(b []byte) []byte

// FindString returns the text of the leftmost match for re in s.
// The return value is the empty string both for an empty match and for no match.
// To distinguish those two cases, use [Regexp.FindStringIndex] or [Regexp.FindStringSubmatch].
func (re *Regexp) FindString(s string) string

// FindIndex returns the location of the leftmost match for re in b.
// The match itself is at b[m[0]:m[1]].
// The return value is nil for no match.
func (re *Regexp) FindIndex(b []byte) (m []int)

// FindStringIndex returns the location of the leftmost match for re in s.
// The match itself is at s[m[0]:m[1]].
// The return value is nil for no match.
func (re *Regexp) FindStringIndex(s string) (m []int)

// FindReaderIndex returns the location of the leftmost match for re in r.
// The match starts at byte index m[0] and ends just before byte index m[1].
// The return value is nil for no match.
//
// FindReaderIndex may read arbitrarily far from r,
// including reading beyond the returned match.
func (re *Regexp) FindReaderIndex(r io.RuneReader) (m []int)

// FindSubmatch returns the first match for re in b, including submatches.
// The overall match is m[0], the first submatch is m[1], and so on.
// The return value is nil for no match.
func (re *Regexp) FindSubmatch(b []byte) [][]byte

// FindStringSubmatch returns the first match for re in s, including submatches.
// The overall match is s[0], the first submatch is s[1], and so on.
// The return value is nil for no match.
func (re *Regexp) FindStringSubmatch(s string) []string

// FindSubmatchIndex returns the first match for re in b, including submatches.
// The overall match is b[m[0]:m[1]], the first submatch is b[m[2]:m[3]], and so on.
// The return value is nil for no match.
func (re *Regexp) FindSubmatchIndex(b []byte) []int

// FindStringSubmatchIndex returns the first match for re in s, including submatches.
// The overall match is s[m[0]:m[1]], the first submatch is s[m[2]:m[3]], and so on.
// The return value is nil for no match.
func (re *Regexp) FindStringSubmatchIndex(s string) []int

// FindReaderSubmatchIndex returns the first match for re in r, including submatches.
// The overall match is at byte index m[0] up to m[1],
// the first submatch is at byte index m[2] up to m[3], and so on.
// The return value is nil for no match.
//
// FindReaderSubmatchIndex may read arbitrarily far from r,
// including reading beyond the returned match.
func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int

// FindAll returns all the matches for re in b.
// If n >= 0, FindAll returns no more than n matches.
// See [Regexp.All] for the equivalent iterator form.
func (re *Regexp) FindAll(b []byte, n int) [][]byte

// FindAllString returns all the matches for re in s.
// If n >= 0, FindAllString returns no more than n matches.
// See [Regexp.AllString] for the equivalent iterator form.
func (re *Regexp) FindAllString(s string, n int) []string

// FindAllIndex returns the locations of all matches for re in b.
// If n >= 0, FindAllIndex returns no more than n matches.
// See [Regexp.AllIndex] for the equivalent iterator form.
func (re *Regexp) FindAllIndex(b []byte, n int) [][]int

// FindAllStringIndex returns the locations of all matches for re in s.
// If n >= 0, FindAllStringIndex returns no more than n matches.
// See [Regexp.AllStringIndex] for the equivalent iterator form.
func (re *Regexp) FindAllStringIndex(s string, n int) [][]int

// FindAllSubmatch returns the locations of all matches for re in b,
// including submatch locations.
// In each returned match m, the overall match is m[0],
// the first submatch is m[1], and so on.
// If n >= 0, FindAllSubmatch returns no more than n matches.
// See [Regexp.AllSubmatch] for the equivalent iterator form.
func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte

// FindAllStringSubmatch returns the locations of all matches for re in s,
// including submatch locations.
// In each returned match m, m[0] is the overall match,
// m[1] is the first submatch, and so on.
// If n >= 0, FindAllStringSubmatch returns no more than n matches.
// See [Regexp.AllStringSubmatch] for the equivalent iterator form.
func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string

// FindAllSubmatchIndex returns the locations of all matches for re in b,
// including submatch locations.
// In each returned match m, the overall match is b[m[0]:m[1]],
// the first submatch is b[m[2]:m[3]], and so on.
// If n >= 0, FindAllSubmatchIndex returns no more than n matches.
// See [Regexp.AllSubmatchIndex] for the equivalent iterator form.
func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int

// FindAllStringSubmatchIndex returns the locations of all matches for re in s,
// including submatch locations.
// In each returned match m, the overall match is s[m[0]:m[1]],
// the first submatch is s[m[2]:m[3]], and so on.
// If n >= 0, FindAllStringSubmatchIndex returns no more than n matches.
// See [Regexp.AllStringSubmatchIndex] for the equivalent iterator form.
func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int

// Expand appends template to dst and returns the result; during the
// append, Expand replaces variables in the template with corresponding
// matches drawn from src. The match slice should have been returned by
// [Regexp.FindSubmatchIndex].
//
// In the template, a variable is denoted by a substring of the form
// $name or ${name}, where name is a non-empty sequence of letters,
// digits, and underscores. A purely numeric name like $1 refers to
// the submatch with the corresponding index; other names refer to
// capturing parentheses named with the (?P<name>...) syntax. A
// reference to an out of range or unmatched index or a name that is not
// present in the regular expression is replaced with an empty slice.
//
// In the $name form, name is taken to be as long as possible: $1x is
// equivalent to ${1x}, not ${1}x, and, $10 is equivalent to ${10}, not ${1}0.
//
// To insert a literal $ in the output, use $$ in the template.
func (re *Regexp) Expand(dst []byte, template []byte, src []byte, match []int) []byte

// ExpandString is like [Regexp.Expand] but the template and source are strings.
// It appends to and returns a byte slice in order to give the calling
// code control over allocation.
func (re *Regexp) ExpandString(dst []byte, template string, src string, match []int) []byte

// Split slices s into substrings separated by the expression and returns a slice of
// the substrings between those expression matches.
//
// The slice returned by this method consists of all the substrings of s
// not contained in the slice returned by [Regexp.FindAllString]. When called on an expression
// that contains no metacharacters, it is equivalent to [strings.SplitN].
//
// Example:
//
//	s := regexp.MustCompile("a*").Split("abaabaccadaaae", 5)
//	// s: ["", "b", "b", "c", "cadaaae"]
//
// The count determines the number of substrings to return:
//   - n > 0: at most n substrings; the last substring will be the unsplit remainder;
//   - n == 0: the result is nil (zero substrings);
//   - n < 0: all substrings.
func (re *Regexp) Split(s string, n int) []string

// AppendText implements [encoding.TextAppender]. The output
// matches that of calling the [Regexp.String] method.
//
// Note that the output is lossy in some cases: This method does not indicate
// POSIX regular expressions (i.e. those compiled by calling [CompilePOSIX]), or
// those for which the [Regexp.Longest] method has been called.
func (re *Regexp) AppendText(b []byte) ([]byte, error)

// MarshalText implements [encoding.TextMarshaler]. The output
// matches that of calling the [Regexp.AppendText] method.
//
// See [Regexp.AppendText] for more information.
func (re *Regexp) MarshalText() ([]byte, error)

// UnmarshalText implements [encoding.TextUnmarshaler] by calling
// [Compile] on the encoded value.
func (re *Regexp) UnmarshalText(text []byte) error
