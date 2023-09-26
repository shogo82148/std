// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http_test

import (
	"log"
	"net/http"
)

// dotFileHidingFile is the http.File use in dotFileHidingFileSystem.
// It is used to wrap the Readdir method of http.File so that we can
// remove files and directories that start with a period from its output.

// dotFileHidingFileSystem is an http.FileSystem that hides
// hidden "dot files" from being served.

func ExampleFileServer_dotFileHiding() {
	fs := dotFileHidingFileSystem{http.Dir(".")}
	http.Handle("/", http.FileServer(fs))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
