// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/os"
)

func ExampleWalkDir() {
	root := "/usr/local/go/bin"
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(path)
		return nil
	})
}
