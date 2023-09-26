// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

/*
mksyscall_windows generates windows system call bodies

It parses all files specified on command line containing function
prototypes (like syscall_windows.go) and prints system call bodies
to standard output.

The prototypes are marked by lines beginning with "//sys" and read
like func declarations if //sys is replaced by func, but:

  - The parameter lists must give a name for each argument. This
    includes return parameters.

  - The parameter lists must give a type for each argument:
    the (x, y, z int) shorthand is not allowed.

* If the return parameter is an error number, it must be named err.

  - If go func name needs to be different from it's winapi dll name,
    the winapi name could be specified at the end, after "=" sign, like
    //sys LoadLibrary(libname string) (handle uint32, err error) = LoadLibraryA

  - Each function that returns err needs to supply a condition, that
    return value of winapi will be tested against to detect failure.
    This would set err to windows "last-error", otherwise it will be nil.
    The value can be provided at end of //sys declaration, like
    //sys LoadLibrary(libname string) (handle uint32, err error) [failretval==-1] = LoadLibraryA
    and is [failretval==0] by default.

Usage:

	mksyscall_windows [flags] [path ...]

The flags are:

	-trace
		Generate print statement after every syscall.
*/
package main

import (
	"github.com/shogo82148/std/flag"
	"github.com/shogo82148/std/io"
)

var PrintTraceFlag = flag.Bool("trace", false, "generate print statement after every syscall")

// Param is function parameter
type Param struct {
	Name      string
	Type      string
	fn        *Fn
	tmpVarIdx int
}

// BoolTmpVarCode returns source code for bool temp variable.
func (p *Param) BoolTmpVarCode() string

// SliceTmpVarCode returns source code for slice temp variable.
func (p *Param) SliceTmpVarCode() string

// StringTmpVarCode returns source code for string temp variable.
func (p *Param) StringTmpVarCode() string

// TmpVarCode returns source code for temp variable.
func (p *Param) TmpVarCode() string

// SyscallArgList returns source code fragments representing p parameter
// in syscall. Slices are translated into 2 syscall parameters: pointer to
// the first element and length.
func (p *Param) SyscallArgList() []string

// IsError determines if p parameter is used to return error.
func (p *Param) IsError() bool

// Rets describes function return parameters.
type Rets struct {
	Name         string
	Type         string
	ReturnsError bool
	FailCond     string
}

// ErrorVarName returns error variable name for r.
func (r *Rets) ErrorVarName() string

// ToParams converts r into slice of *Param.
func (r *Rets) ToParams() []*Param

// List returns source code of syscall return parameters.
func (r *Rets) List() string

// PrintList returns source code of trace printing part correspondent
// to syscall return values.
func (r *Rets) PrintList() string

// SetReturnValuesCode returns source code that accepts syscall return values.
func (r *Rets) SetReturnValuesCode() string

// SetErrorCode returns source code that sets return parameters.
func (r *Rets) SetErrorCode() string

// Fn describes syscall function.
type Fn struct {
	Name        string
	Params      []*Param
	Rets        *Rets
	PrintTrace  bool
	dllname     string
	dllfuncname string
	src         string

	curTmpVarIdx int
}

// DLLName returns DLL name for function f.
func (f *Fn) DLLName() string

// DLLName returns DLL function name for function f.
func (f *Fn) DLLFuncName() string

// ParamList returns source code for function f parameters.
func (f *Fn) ParamList() string

// ParamPrintList returns source code of trace printing part correspondent
// to syscall input parameters.
func (f *Fn) ParamPrintList() string

// ParamCount return number of syscall parameters for function f.
func (f *Fn) ParamCount() int

// SyscallParamCount determines which version of Syscall/Syscall6/Syscall9/...
// to use. It returns parameter count for correspondent SyscallX function.
func (f *Fn) SyscallParamCount() int

// Syscall determines which SyscallX function to use for function f.
func (f *Fn) Syscall() string

// SyscallParamList returns source code for SyscallX parameters for function f.
func (f *Fn) SyscallParamList() string

// IsUTF16 is true, if f is W (utf16) function. It is false
// for all A (ascii) functions.
func (f *Fn) IsUTF16() bool

// StrconvFunc returns name of Go string to OS string function for f.
func (f *Fn) StrconvFunc() string

// StrconvType returns Go type name used for OS string for f.
func (f *Fn) StrconvType() string

// Source files and functions.
type Source struct {
	Funcs []*Fn
	Files []string
}

// ParseFiles parses files listed in fs and extracts all syscall
// functions listed in  sys comments. It returns source files
// and functions collection *Source if successful.
func ParseFiles(fs []string) (*Source, error)

// DLLs return dll names for a source set src.
func (src *Source) DLLs() []string

// ParseFile adds adition file path to a source set src.
func (src *Source) ParseFile(path string) error

// Generate output source file from a source set src.
func (src *Source) Generate(w io.Writer) error

// TODO: use println instead to print in the following template
