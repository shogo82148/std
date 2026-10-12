// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflectdata

import (
	"github.com/shogo82148/std/cmd/compile/internal/types"
)

// MapGroupType makes the map slot group type given the type of the map.
func MapGroupType(t *types.Type) *types.Type

// MapType returns a type interchangeable with internal/runtime/maps.Map.
// Make sure this stays in sync with internal/runtime/maps/map.go.
func MapType() *types.Type

// MapIterType returns a type interchangeable with internal/runtime/maps.Iter.
// Make sure this stays in sync with internal/runtime/maps/table.go.
func MapIterType() *types.Type
