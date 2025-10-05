// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

// Package jsontest contains functionality to assist in testing JSON.
package jsontest

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/internal/zstd"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/path"
	"github.com/shogo82148/std/slices"
	"github.com/shogo82148/std/strings"
	"github.com/shogo82148/std/sync"
)

type Entry struct {
	Name string
	Data func() []byte
	New  func() any
}

// Data is a list of JSON testdata.
var Data = func() (entries []Entry) {
	fis := mustGet(fs.ReadDir(testdataFS, "testdata"))
	slices.SortFunc(fis, func(x, y fs.DirEntry) int { return strings.Compare(x.Name(), y.Name()) })
	for _, fi := range fis {
		var entry Entry

		words := strings.Split(strings.TrimSuffix(fi.Name(), ".json.zst"), "_")
		for i := range words {
			words[i] = strings.Title(words[i])
		}
		entry.Name = strings.Join(words, "")

		entry.Data = sync.OnceValue(func() []byte {
			filePath := path.Join("testdata", fi.Name())
			b := mustGet(fs.ReadFile(testdataFS, filePath))
			zr := zstd.NewReader(bytes.NewReader(b))
			return mustGet(io.ReadAll(zr))
		})

		switch entry.Name {
		case "CanadaGeometry":
			entry.New = func() any { return new(canadaRoot) }
		case "CitmCatalog":
			entry.New = func() any { return new(citmRoot) }
		case "GolangSource":
			entry.New = func() any { return new(golangRoot) }
		case "StringEscaped":
			entry.New = func() any { return new(stringRoot) }
		case "StringUnicode":
			entry.New = func() any { return new(stringRoot) }
		case "SyntheaFhir":
			entry.New = func() any { return new(syntheaRoot) }
		case "TwitterStatus":
			entry.New = func() any { return new(twitterRoot) }
		}

		entries = append(entries, entry)
	}
	return entries
}()
