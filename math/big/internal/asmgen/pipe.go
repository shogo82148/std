// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asmgen

// A Pipe manages the input and output data pipelines for a function's
// memory operations.
//
// The input is one or more equal-length slices of words, so collectively
// it can be viewed as a matrix, in which each slice is a row and each column
// is a set of corresponding words from the different slices.
// The output can be viewed the same way, although it is often just one row.
type Pipe struct {
	f               *Func
	label           string
	backward        bool
	started         bool
	loaded          bool
	inPtr           []RegPtr
	hints           []Hint
	outPtr          []RegPtr
	index           Reg
	useIndexCounter bool
	indexCounter    int
	readOff         int
	writeOff        int
	factors         []int
	counts          []Reg
	needWrite       bool
	maxColumns      int
	unrollStart     func()
	unrollEnd       func()
}

// Pipe creates and returns a new pipe for use in the function f.
func (f *Func) Pipe() *Pipe

// SetBackward sets the pipe to process the input and output columns in reverse order.
// This is needed for left shifts, which might otherwise overwrite data they will read later.
func (p *Pipe) SetBackward()

// SetUseIndexCounter sets the pipe to use an index counter if possible,
// meaning the loop counter is also used as an index for accessing the slice data.
// This clever trick is slower on modern processors, but it is still necessary on 386.
// On non-386 systems, SetUseIndexCounter is a no-op.
func (p *Pipe) SetUseIndexCounter()

// SetLabel sets the label prefix for the loops emitted by the pipe.
// The default prefix is "loop".
func (p *Pipe) SetLabel(label string)

// SetMaxColumns sets the maximum number of
// columns processed in a single loop body call.
func (p *Pipe) SetMaxColumns(m int)

// SetHint records that the inputs from the named vector
// should be allocated with the given register hint.
//
// If the hint indicates a single register on the target architecture,
// then SetHint calls SetMaxColumns(1), since the hinted register
// can only be used for one value at a time.
func (p *Pipe) SetHint(name string, hint Hint)

// LoadPtrs loads the slice pointer arguments into registers,
// assuming that the slice length n has already been loaded
// into the register n.
//
// Start will call LoadPtrs if it has not been called already.
// LoadPtrs only needs to be called explicitly when code needs
// to use LoadN before Start, like when the shift.go generators
// read an initial word before the loop.
func (p *Pipe) LoadPtrs(n Reg)

// LoadN returns the next n columns of input words as a slice of rows.
// Regs for inputs that have been marked using p.SetMemOK will be direct memory references.
// Regs for other inputs will be newly allocated registers and must be freed.
func (p *Pipe) LoadN(n int) [][]Reg

// StoreN writes regs (a slice of rows) to the next n columns of output, where n = len(regs[0]).
func (p *Pipe) StoreN(regs [][]Reg)

// DropInput deletes the named input from the pipe,
// usually because it has been exhausted.
// (This is not used yet but will be used in a future generator.)
func (p *Pipe) DropInput(name string)

// Start prepares to loop over n columns.
// The factors give a sequence of unrolling factors to use,
// which must be either strictly increasing or strictly decreasing
// and must include 1.
// For example, 4, 1 means to process 4 elements at a time
// and then 1 at a time for the final 0-3; specifying 1,4 instead
// handles 0-3 elements first and then 4 at a time.
// Similarly, 32, 4, 1 means to process 32 at a time,
// then 4 at a time, then 1 at a time.
//
// One benefit of using 1, 4 instead of 4, 1 is that the body
// processing 4 at a time needs more registers, and if it is
// the final body, the register holding the fragment count (0-3)
// has been freed and is available for use.
//
// Start may modify the carry flag.
//
// Start must be followed by a call to Loop1 or LoopN,
// but it is permitted to emit other instructions first,
// for example to set an initial carry flag.
func (p *Pipe) Start(n Reg, factors ...int)

// Restart prepares to loop over an additional n columns,
// beyond a previous loop run by p.Start/p.Loop.
func (p *Pipe) Restart(n Reg, factors ...int)

// Done frees all the registers allocated by the pipe.
func (p *Pipe) Done()

// Loop emits code for the loop, calling block repeatedly to emit code that
// handles a block of N input columns (for arbitrary N = len(in[0]) chosen by p).
// block must call p.StoreN(out) to write N output columns.
// The out slice is a pre-allocated matrix of uninitialized Reg values.
// block is expected to set each entry to the Reg that should be written
// before calling p.StoreN(out).
//
// For example, if the loop is to be unrolled 4x in blocks of 2 columns each,
// the sequence of calls to emit the unrolled loop body is:
//
//	start()  // set by pAtUnrollStart
//	... reads for 2 columns ...
//	block()
//	... writes for 2 columns ...
//	... reads for 2 columns ...
//	block()
//	... writes for 2 columns ...
//	end()  // set by p.AtUnrollEnd
//
// Any registers allocated during block are freed automatically when block returns.
func (p *Pipe) Loop(block func(in, out [][]Reg))

// AtUnrollStart sets a function to call at the start of an unrolled sequence.
// See [Pipe.Loop] for details.
func (p *Pipe) AtUnrollStart(start func())

// AtUnrollEnd sets a function to call at the end of an unrolled sequence.
// See [Pipe.Loop] for details.
func (p *Pipe) AtUnrollEnd(end func())
