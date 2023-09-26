// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.staticlockranking
// +build goexperiment.staticlockranking

package runtime

// worldIsStopped is accessed atomically to track world-stops. 1 == world
// stopped.

// lockRankStruct is embedded in mutex
