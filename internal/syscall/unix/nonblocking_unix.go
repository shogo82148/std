// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package unix

func IsNonblock(fd int) (nonblocking bool, err error)

func HasNonblockFlag(flag int) bool
