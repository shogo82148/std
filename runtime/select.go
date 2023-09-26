// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// A runtimeSelect is a single case passed to rselect.
// This must match ../reflect/value.go:/runtimeSelect

// These values must match ../reflect/value.go:/SelectDir.

const (
	_ selectDir = iota
)
