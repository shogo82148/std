// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imports

import (
	"github.com/shogo82148/std/fmt"
)

func ScanDir(path string, tags map[string]bool) ([]string, []string, error)

func ScanFiles(files []string, tags map[string]bool) ([]string, []string, error)

var ErrNoGo = fmt.Errorf("no Go source files")
