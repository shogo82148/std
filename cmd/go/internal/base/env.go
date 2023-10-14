// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package base

// AppendPWD returns the result of appending PWD=dir to the environment base.
//
// The resulting environment makes os.Getwd more efficient for a subprocess
// running in dir, and also improves the accuracy of paths relative to dir
// if one or more elements of dir is a symlink.
func AppendPWD(base []string, dir string) []string

// AppendPATH returns the result of appending PATH=$GOROOT/bin:$PATH
// (or the platform equivalent) to the environment base.
func AppendPATH(base []string) []string
