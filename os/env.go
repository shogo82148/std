// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// General environment variables.

package os

// Expand replaces ${var} or $var in the string based on the mapping function.
// Invocations of undefined variables are replaced with the empty string.
func Expand(s string, mapping func(string) string) string

// ExpandEnv replaces ${var} or $var in the string according to the values
// of the current environment variables.  References to undefined
// variables are replaced by the empty string.
func ExpandEnv(s string) string

// Getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
func Getenv(key string) string

// Setenv sets the value of the environment variable named by the key.
// It returns an error, if any.
func Setenv(key, value string) error

// Clearenv deletes all environment variables.
func Clearenv()

// Environ returns a copy of strings representing the environment,
// in the form "key=value".
func Environ() []string
