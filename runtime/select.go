// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// scase.kind values.
// Known to compiler.
// Changes here must also be made in src/cmd/compile/internal/gc/select.go's walkselect.

// Select case descriptor.
// Known to compiler.
// Changes here must also be made in src/cmd/internal/gc/select.go's scasetype.

// A runtimeSelect is a single case passed to rselect.
// This must match ../reflect/value.go:/runtimeSelect

// These values must match ../reflect/value.go:/SelectDir.

const (
	_ selectDir = iota
)
