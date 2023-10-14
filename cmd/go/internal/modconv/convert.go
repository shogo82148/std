// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modconv

import (
	"github.com/shogo82148/std/golang.org/x/mod/modfile"
	"github.com/shogo82148/std/golang.org/x/mod/module"
)

// ConvertLegacyConfig converts legacy config to modfile.
// The file argument is slash-delimited.
func ConvertLegacyConfig(f *modfile.File, file string, data []byte, queryPackage func(path, rev string) (module.Version, error)) error
