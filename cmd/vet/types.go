// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the pieces of the tool that use typechecking from the go/types package.

package main

// stdImporter is the importer we use to import packages.
// It is shared so that all packages are imported by the same importer.
