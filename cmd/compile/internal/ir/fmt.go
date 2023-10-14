// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ir

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io"
)

var OpNames = []string{
	OADDR:             "&",
	OADD:              "+",
	OADDSTR:           "+",
	OALIGNOF:          "unsafe.Alignof",
	OANDAND:           "&&",
	OANDNOT:           "&^",
	OAND:              "&",
	OAPPEND:           "append",
	OAS:               "=",
	OAS2:              "=",
	OBREAK:            "break",
	OCALL:             "function call",
	OCAP:              "cap",
	OCASE:             "case",
	OCLEAR:            "clear",
	OCLOSE:            "close",
	OCOMPLEX:          "complex",
	OBITNOT:           "^",
	OCONTINUE:         "continue",
	OCOPY:             "copy",
	ODELETE:           "delete",
	ODEFER:            "defer",
	ODIV:              "/",
	OEQ:               "==",
	OFALL:             "fallthrough",
	OFOR:              "for",
	OGE:               ">=",
	OGOTO:             "goto",
	OGT:               ">",
	OIF:               "if",
	OIMAG:             "imag",
	OINLMARK:          "inlmark",
	ODEREF:            "*",
	OLEN:              "len",
	OLE:               "<=",
	OLSH:              "<<",
	OLT:               "<",
	OMAKE:             "make",
	ONEG:              "-",
	OMAX:              "max",
	OMIN:              "min",
	OMOD:              "%",
	OMUL:              "*",
	ONEW:              "new",
	ONE:               "!=",
	ONOT:              "!",
	OOFFSETOF:         "unsafe.Offsetof",
	OOROR:             "||",
	OOR:               "|",
	OPANIC:            "panic",
	OPLUS:             "+",
	OPRINTN:           "println",
	OPRINT:            "print",
	ORANGE:            "range",
	OREAL:             "real",
	ORECV:             "<-",
	ORECOVER:          "recover",
	ORETURN:           "return",
	ORSH:              ">>",
	OSELECT:           "select",
	OSEND:             "<-",
	OSIZEOF:           "unsafe.Sizeof",
	OSUB:              "-",
	OSWITCH:           "switch",
	OUNSAFEADD:        "unsafe.Add",
	OUNSAFESLICE:      "unsafe.Slice",
	OUNSAFESLICEDATA:  "unsafe.SliceData",
	OUNSAFESTRING:     "unsafe.String",
	OUNSAFESTRINGDATA: "unsafe.StringData",
	OXOR:              "^",
}

// GoString returns the Go syntax for the Op, or else its name.
func (o Op) GoString() string

// Format implements formatting for an Op.
// The valid formats are:
//
//	%v	Go syntax ("+", "<-", "print")
//	%+v	Debug syntax ("ADD", "RECV", "PRINT")
func (o Op) Format(s fmt.State, verb rune)

var OpPrec = []int{
	OALIGNOF:          8,
	OAPPEND:           8,
	OBYTES2STR:        8,
	OARRAYLIT:         8,
	OSLICELIT:         8,
	ORUNES2STR:        8,
	OCALLFUNC:         8,
	OCALLINTER:        8,
	OCALLMETH:         8,
	OCALL:             8,
	OCAP:              8,
	OCLEAR:            8,
	OCLOSE:            8,
	OCOMPLIT:          8,
	OCONVIFACE:        8,
	OCONVIDATA:        8,
	OCONVNOP:          8,
	OCONV:             8,
	OCOPY:             8,
	ODELETE:           8,
	OGETG:             8,
	OLEN:              8,
	OLITERAL:          8,
	OMAKESLICE:        8,
	OMAKESLICECOPY:    8,
	OMAKE:             8,
	OMAPLIT:           8,
	OMAX:              8,
	OMIN:              8,
	ONAME:             8,
	ONEW:              8,
	ONIL:              8,
	ONONAME:           8,
	OOFFSETOF:         8,
	OPANIC:            8,
	OPAREN:            8,
	OPRINTN:           8,
	OPRINT:            8,
	ORUNESTR:          8,
	OSIZEOF:           8,
	OSLICE2ARR:        8,
	OSLICE2ARRPTR:     8,
	OSTR2BYTES:        8,
	OSTR2RUNES:        8,
	OSTRUCTLIT:        8,
	OTYPE:             8,
	OUNSAFEADD:        8,
	OUNSAFESLICE:      8,
	OUNSAFESLICEDATA:  8,
	OUNSAFESTRING:     8,
	OUNSAFESTRINGDATA: 8,
	OINDEXMAP:         8,
	OINDEX:            8,
	OSLICE:            8,
	OSLICESTR:         8,
	OSLICEARR:         8,
	OSLICE3:           8,
	OSLICE3ARR:        8,
	OSLICEHEADER:      8,
	OSTRINGHEADER:     8,
	ODOTINTER:         8,
	ODOTMETH:          8,
	ODOTPTR:           8,
	ODOTTYPE2:         8,
	ODOTTYPE:          8,
	ODOT:              8,
	OXDOT:             8,
	OMETHVALUE:        8,
	OMETHEXPR:         8,
	OPLUS:             7,
	ONOT:              7,
	OBITNOT:           7,
	ONEG:              7,
	OADDR:             7,
	ODEREF:            7,
	ORECV:             7,
	OMUL:              6,
	ODIV:              6,
	OMOD:              6,
	OLSH:              6,
	ORSH:              6,
	OAND:              6,
	OANDNOT:           6,
	OADD:              5,
	OSUB:              5,
	OOR:               5,
	OXOR:              5,
	OEQ:               4,
	OLT:               4,
	OLE:               4,
	OGE:               4,
	OGT:               4,
	ONE:               4,
	OSEND:             3,
	OANDAND:           2,
	OOROR:             1,

	OAS:         -1,
	OAS2:        -1,
	OAS2DOTTYPE: -1,
	OAS2FUNC:    -1,
	OAS2MAPR:    -1,
	OAS2RECV:    -1,
	OASOP:       -1,
	OBLOCK:      -1,
	OBREAK:      -1,
	OCASE:       -1,
	OCONTINUE:   -1,
	ODCL:        -1,
	ODEFER:      -1,
	OFALL:       -1,
	OFOR:        -1,
	OGOTO:       -1,
	OIF:         -1,
	OLABEL:      -1,
	OGO:         -1,
	ORANGE:      -1,
	ORETURN:     -1,
	OSELECT:     -1,
	OSWITCH:     -1,

	OEND: 0,
}

// StmtWithInit reports whether op is a statement with an explicit init list.
func StmtWithInit(op Op) bool

// Format implements formatting for a Nodes.
// The valid formats are:
//
//	%v	Go syntax, semicolon-separated
//	%.v	Go syntax, comma-separated
//	%+v	Debug syntax, as in DumpList.
func (l Nodes) Format(s fmt.State, verb rune)

// Dump prints the message s followed by a debug dump of n.
func Dump(s string, n Node)

// DumpList prints the message s followed by a debug dump of each node in the list.
func DumpList(s string, list Nodes)

// FDumpList prints to w the message s followed by a debug dump of each node in the list.
func FDumpList(w io.Writer, s string, list Nodes)

// EscFmt is set by the escape analysis code to add escape analysis details to the node print.
var EscFmt func(n Node) string
