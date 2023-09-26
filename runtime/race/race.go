// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (race && linux && amd64) || (race && darwin && amd64) || (race && windows && amd64)
// +build race,linux,amd64 race,darwin,amd64 race,windows,amd64

package race

// void __race_unused_func(void);
