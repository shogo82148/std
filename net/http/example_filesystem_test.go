// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http_test

import (
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/os"
)

// FileServerFS will serve files starting with a dot, which can expose sensitive
// directories such as .git or sensitive files such as .htpassword.
//
// This example demonstrates hiding dot files by wrapping the fs.FS.
func ExampleFileServerFS_dotFileHiding() {
	root, err := os.OpenRoot("doc")
	if err != nil {
		log.Fatal(err)
	}
	fsys := dotFileHidingFileSystem{root.FS()}
	handler := http.FileServerFS(fsys)
	log.Fatal(http.ListenAndServe(":8080", handler))
}
