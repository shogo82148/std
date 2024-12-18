// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// An SRV represents a single DNS SRV record.
type SRV struct {
	Target   string
	Port     uint16
	Priority uint16
	Weight   uint16
}

// An MX represents a single DNS MX record.
type MX struct {
	Host string
	Pref uint16
}

// An NS represents a single DNS NS record.
type NS struct {
	Host string
}
