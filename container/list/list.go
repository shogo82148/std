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

// Element is an element in the linked list.
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
	front, back *Element
	len         int
}

// Init initializes or clears a List.
func (l *List) Init() *List

// New returns an initialized list.
func New() *List

// Front returns the first element in the list.
func (l *List) Front() *Element

// Back returns the last element in the list.
func (l *List) Back() *Element

// Remove removes the element from the list
// and returns its Value.
func (l *List) Remove(e *Element) interface{}

// PushFront inserts the value at the front of the list and returns a new Element containing the value.
func (l *List) PushFront(value interface{}) *Element

// PushBack inserts the value at the back of the list and returns a new Element containing the value.
func (l *List) PushBack(value interface{}) *Element

// InsertBefore inserts the value immediately before mark and returns a new Element containing the value.
func (l *List) InsertBefore(value interface{}, mark *Element) *Element

// InsertAfter inserts the value immediately after mark and returns a new Element containing the value.
func (l *List) InsertAfter(value interface{}, mark *Element) *Element

// MoveToFront moves the element to the front of the list.
func (l *List) MoveToFront(e *Element)

// MoveToBack moves the element to the back of the list.
func (l *List) MoveToBack(e *Element)

// Len returns the number of elements in the list.
func (l *List) Len() int

// PushBackList inserts each element of ol at the back of the list.
func (l *List) PushBackList(ol *List)

// PushFrontList inserts each element of ol at the front of the list. The ordering of the passed list is preserved.
func (l *List) PushFrontList(ol *List)
