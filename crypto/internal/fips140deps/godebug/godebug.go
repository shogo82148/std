// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package godebug

import (
	"github.com/shogo82148/std/internal/godebug"
)

type Setting godebug.Setting

func New(name string) *Setting

func (s *Setting) Value() string

func Value(name string) string
