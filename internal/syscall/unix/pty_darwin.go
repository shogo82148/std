// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

func Grantpt(fd int) error

func Unlockpt(fd int) error

func Ptsname(fd int) (string, error)

func PosixOpenpt(flag int) (fd int, err error)
