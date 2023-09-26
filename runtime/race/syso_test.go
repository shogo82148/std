// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !android && !js && !ppc64le
// +build !android,!js,!ppc64le

// Note: we don't run on Android or ppc64 because if there is any non-race test
// file in this package, the OS tries to link the .syso file into the
// test (even when we're not in race mode), which fails. I'm not sure
// why, but easiest to just punt - as long as a single builder runs
// this test, we're good.

package race
