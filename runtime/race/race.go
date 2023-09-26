// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (race && linux && amd64) || (race && freebsd && amd64) || (race && netbsd && amd64) || (race && darwin && amd64) || (race && windows && amd64) || (race && linux && ppc64le)
// +build race,linux,amd64 race,freebsd,amd64 race,netbsd,amd64 race,darwin,amd64 race,windows,amd64 race,linux,ppc64le

package race

// void __race_unused_func(void);
