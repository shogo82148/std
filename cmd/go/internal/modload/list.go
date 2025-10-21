// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modload

import (
	"github.com/shogo82148/std/context"

	"github.com/shogo82148/std/cmd/go/internal/modinfo"
)

type ListMode int

const (
	ListU ListMode = 1 << iota
	ListRetracted
	ListDeprecated
	ListVersions
	ListRetractedVersions
)

// ListModules returns a description of the modules matching args, if known,
// along with any error preventing additional matches from being identified.
//
// The returned slice can be nonempty even if the error is non-nil.
func ListModules(loaderstate *State, ctx context.Context, args []string, mode ListMode, reuseFile string) ([]*modinfo.ModulePublic, error)
