// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package list implements a doubly linked list.
//
// To iterate over a list (where l is a *List):
//
//	for e := l.Front(); e != nil; e = e.Next() {
//		// do something with e.Value
//	}
package list

// Element is an element of a linked list.
type Element struct {
	next, prev *Element

	list *List

	Value interface{}
}

// Next returns the next list element or nil.
func (e *Element) Next() *Element

// Prev returns the previous list element or nil.
func (e *Element) Prev() *Element

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use.
type List struct {
	root Element
	len  int
}

// Init initializes or clears list l.
func (l *List) Init() *List

// New returns an initialized list.
func New() *List

// Len returns the number of elements of list l.
func (l *List) Len() int

// Front returns the first element of list l or nil
func (l *List) Front() *Element

// Back returns the last element of list l or nil.
func (l *List) Back() *Element

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
func (l *List) Remove(e *Element) interface{}

// Pushfront inserts a new element e with value v at the front of list l and returns e.
func (l *List) PushFront(v interface{}) *Element

// PushBack inserts a new element e with value v at the back of list l and returns e.
func (l *List) PushBack(v interface{}) *Element

// InsertBefore inserts a new element e with value v immediately before mark and returns e.
// If mark is not an element of l, the list is not modified.
func (l *List) InsertBefore(v interface{}, mark *Element) *Element

// InsertAfter inserts a new element e with value v immediately after mark and returns e.
// If mark is not an element of l, the list is not modified.
func (l *List) InsertAfter(v interface{}, mark *Element) *Element

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
func (l *List) MoveToFront(e *Element)

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
func (l *List) MoveToBack(e *Element)

// PushBackList inserts a copy of an other list at the back of list l.
// The lists l and other may be the same.
func (l *List) PushBackList(other *List)

// PushFrontList inserts a copy of an other list at the front of list l.
// The lists l and other may be the same.
func (l *List) PushFrontList(other *List)
