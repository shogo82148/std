// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package synctest provides support for testing concurrent code.
//
// See the testing/synctest package for function documentation.
package synctest

//go:linkname Run
func Run(f func())

//go:linkname Wait
func Wait()

// IsInBubble reports whether the current goroutine is in a bubble.
//
//go:linkname IsInBubble
func IsInBubble() bool

// Association is the state of a pointer's bubble association.
type Association int

const (
	Unbubbled     = Association(iota)
	CurrentBubble
	OtherBubble
)

// Associate attempts to associate p with the current bubble.
// It returns the new association status of p.
func Associate[T any](p *T) Association

// Disassociate disassociates p from any bubble.
func Disassociate[T any](p *T)

// IsAssociated reports whether p is associated with the current bubble.
func IsAssociated[T any](p *T) bool

// A Bubble is a synctest bubble.
//
// Not a public API. Used by syscall/js to propagate bubble membership through syscalls.
type Bubble struct {
	b any
}

// Acquire returns a reference to the current goroutine's bubble.
// The bubble will not become idle until Release is called.
func Acquire() *Bubble

// Release releases the reference to the bubble,
// allowing it to become idle again.
func (b *Bubble) Release()

// Run executes f in the bubble.
// The current goroutine must not be part of a bubble.
func (b *Bubble) Run(f func())
