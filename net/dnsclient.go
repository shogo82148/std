// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// SRVは単一のDNS SRVレコードを表します。
type SRV struct {
	Target   string
	Port     uint16
	Priority uint16
	Weight   uint16
}

// MXは単一のDNS MXレコードを表します。
type MX struct {
	Host string
	Pref uint16
}

// NSは単一のDNS NSレコードを表します。
type NS struct {
	Host string
}
