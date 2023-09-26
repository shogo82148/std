// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

// Description describes a runtime metric.
type Description struct {
	Name string

	Description string

	Kind ValueKind

	Cumulative bool
}

// The English language descriptions below must be kept in sync with the
// descriptions of each metric in doc.go by running 'go generate'.

// All returns a slice of containing metric descriptions for all supported metrics.
func All() []Description
