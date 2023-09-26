// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ((darwin && arm64) || (darwin && amd64 && ios)) && x509omitbundledroots
// +build darwin,arm64 darwin,amd64,ios
// +build x509omitbundledroots

package x509
