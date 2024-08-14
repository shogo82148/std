// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflectdata

import (
	"github.com/shogo82148/std/cmd/compile/internal/types"
)

// SwissMapBucketType makes the map bucket type given the type of the map.
func SwissMapBucketType(t *types.Type) *types.Type

// SwissMapType returns a type interchangeable with runtime.hmap.
// Make sure this stays in sync with runtime/map.go.
func SwissMapType() *types.Type

// SwissMapIterType returns a type interchangeable with runtime.hiter.
// Make sure this stays in sync with runtime/map.go.
func SwissMapIterType() *types.Type
