// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package main

import (
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/sync"
)

func HelloServer(w http.ResponseWriter, req *http.Request)

// Simple counter server. POSTing to it will set the value.
type Counter struct {
	mu sync.Mutex
	n  int
}

// This makes Counter satisfy the expvar.Var interface, so we can export
// it directly.
func (ctr *Counter) String() string

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request)

func FlagServer(w http.ResponseWriter, req *http.Request)

// simple argument server
func ArgServer(w http.ResponseWriter, req *http.Request)

// a channel (just for the fun of it)
type Chan chan int

func ChanCreate() Chan

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request)

// exec a program, redirecting output.
func DateServer(rw http.ResponseWriter, req *http.Request)

func Logger(w http.ResponseWriter, req *http.Request)
