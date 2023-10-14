// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin && !ios

package unix

// GetEntropy calls the macOS getentropy system call.
func GetEntropy(p []byte) error
