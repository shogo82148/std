// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modfetch

// LegacyGoMod generates a fake go.mod file for a module that doesn't have one.
// The go.mod file contains a module directive and nothing else: no go version,
// no requirements.
//
// We used to try to build a go.mod reflecting pre-existing
// package management metadata files, but the conversion
// was inherently imperfect (because those files don't have
// exactly the same semantics as go.mod) and, when done
// for dependencies in the middle of a build, impossible to
// correct. So we stopped.
func LegacyGoMod(modPath string) []byte
