// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// the race detector does not like our pointer shenanigans
// while checking the stack.

//go:build goexperiment.runtimesecret && (arm64 || amd64) && linux && !race

package secret
