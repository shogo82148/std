// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || openbsd
// +build aix openbsd

package os

// We query the working directory at init, to use it later to search for the
// executable file
// errWd will be checked later, if we need to use initWd
