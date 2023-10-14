// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Metadata prints basic system metadata to include in test logs. This is
// separate from cmd/dist so it does not need to build with the bootstrap
// toolchain.

// This program is only used by cmd/dist. Add an "ignore" build tag so it
// is not installed. cmd/dist does "go run main.go" directly.

//go:build ignore

package main
