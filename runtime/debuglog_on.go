// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build debuglog
// +build debuglog

package runtime

// dlogPerM is the per-M debug log data. This is embedded in the m
// struct.
