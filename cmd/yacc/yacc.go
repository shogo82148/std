/*
Derived from Inferno's utils/iyacc/yacc.c
http://code.google.com/p/inferno-os/source/browse/utils/iyacc/yacc.c

This copyright NOTICE applies to all files in this directory and
subdirectories, unless another copyright notice appears in a given
file or subdirectory.  If you take substantial code from this software to use in
other programs, you must somehow include with it an appropriate
copyright notice that includes the copyright notice and the other
notices below.  It is fine (and often tidier) to do that in a separate
file such as NOTICE, LICENCE or COPYING.

	Copyright © 1994-1999 Lucent Technologies Inc.  All rights reserved.
	Portions Copyright © 1995-1997 C H Forsyth (forsyth@terzarima.net)
	Portions Copyright © 1997-1999 Vita Nuova Limited
	Portions Copyright © 2000-2007 Vita Nuova Holdings Limited (www.vitanuova.com)
	Portions Copyright © 2004,2006 Bruce Ellis
	Portions Copyright © 2005-2007 C H Forsyth (forsyth@terzarima.net)
	Revisions Copyright © 2000-2007 Lucent Technologies Inc. and others
	Portions Copyright © 2009 The Go Authors.  All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

// the following are adjustable
// according to memory size
const (
	ACTSIZE  = 30000
	NSTATES  = 2000
	TEMPSIZE = 2000

	SYMINC   = 50
	RULEINC  = 50
	PRODINC  = 100
	WSETINC  = 50
	STATEINC = 200

	NAMESIZE = 50
	NTYPES   = 63
	ISIZE    = 400

	PRIVATE = 0xE000

	NTBASE     = 010000
	ERRCODE    = 8190
	ACCEPTCODE = 8191
	YYLEXUNK   = 3
	TOKSTART   = 4
)

// no, left, right, binary assoc.
const (
	NOASC = iota
	LASC
	RASC
	BASC
)

// flags for state generation
const (
	DONE = iota
	MUSTDO
	MUSTLOOKAHEAD
)

// flags for a rule having an action, and being reduced
const (
	ACTFLAG = 1 << (iota + 2)
	REDFLAG
)

// output parser flags

// parse tokens
const (
	IDENTIFIER = PRIVATE + iota
	MARK
	TERM
	LEFT
	RIGHT
	BINARY
	PREC
	LCURLY
	IDENTCOLON
	NUMBER
	START
	TYPEDEF
	TYPENAME
	UNION
	ERROR
)

const ENDFILE = 0
const EMPTY = 1
const WHOKNOWS = 0
const OK = 1
const NOMORE = -1000

// macros for getting associativity and precedence levels
func ASSOC(i int) int

func PLEVEL(i int) int

func TYPE(i int) int

// macros for setting associativity and precedence levels
func SETASC(i, j int) int

func SETPLEV(i, j int) int

func SETTYPE(i, j int) int

// I/O descriptors

// communication variables between various I/O routines

// structure declarations
type Lkset []int

type Pitem struct {
	prod   []int
	off    int
	first  int
	prodno int
}

type Item struct {
	pitem Pitem
	look  Lkset
}

type Symb struct {
	name    string
	noconst bool
	value   int
}

type Wset struct {
	pitem Pitem
	flag  int
	ws    Lkset
}

// storage of types

type Resrv struct {
	name  string
	value int
}

type Error struct {
	lineno int
	tokens []string
	msg    string
}

type Row struct {
	actions       []int
	defaultAction int
}

const EOF = -1

//
// utility routines
//
