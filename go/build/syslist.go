// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package build

// knownOS is the list of past, present, and future known GOOS values.
// Do not remove from this list, as it is used for filename matching.
// If you add an entry to this list, look at unixOS, below.

// unixOS is the set of GOOS values matched by the "unix" build tag.
// This is not used for filename matching.
// This list also appears in cmd/dist/build.go and
// cmd/go/internal/imports/build.go.

// knownArch is the list of past, present, and future known GOARCH values.
// Do not remove from this list, as it is used for filename matching.
