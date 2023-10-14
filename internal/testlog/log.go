// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package testlog provides a back-channel communication path
// between tests and package os, so that cmd/go can see which
// environment variables and files a test consults.
package testlog

// Interface is the interface required of test loggers.
// The os package will invoke the interface's methods to indicate that
// it is inspecting the given environment variables or files.
// Multiple goroutines may call these methods simultaneously.
type Interface interface {
	Getenv(key string)
	Stat(file string)
	Open(file string)
	Chdir(dir string)
}

// SetLogger sets the test logger implementation for the current process.
// It must be called only once, at process startup.
func SetLogger(impl Interface)

// Logger returns the current test logger implementation.
// It returns nil if there is no logger.
func Logger() Interface

// Getenv calls Logger().Getenv, if a logger has been set.
func Getenv(name string)

// Open calls Logger().Open, if a logger has been set.
func Open(name string)

// Stat calls Logger().Stat, if a logger has been set.
func Stat(name string)
