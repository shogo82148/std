// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ioutil_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io/ioutil"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/path/filepath"
	"github.com/shogo82148/std/strings"
)

func ExampleReadAll() {
	r := strings.NewReader("Go is a general-purpose language designed with systems programming in mind.")

	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)

	// Output:
	// Go is a general-purpose language designed with systems programming in mind.
}

func ExampleReadDir() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func ExampleTempDir() {
	content := []byte("temporary file's content")
	dir, err := ioutil.TempDir("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.RemoveAll(dir) // clean up

	tmpfn := filepath.Join(dir, "tmpfile")
	if err := ioutil.WriteFile(tmpfn, content, 0666); err != nil {
		log.Fatal(err)
	}
}

func ExampleTempDir_suffix() {
	parentDir := os.TempDir()
	logsDir, err := ioutil.TempDir(parentDir, "*-logs")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(logsDir) // クリーンアップ

	// ログは、*-logsで終わるサフィックスを持つすべてのディレクトリを検索することで、必要に応じて早期にクリーンアップできます。
	globPattern := filepath.Join(parentDir, "*-logs")
	matches, err := filepath.Glob(globPattern)
	if err != nil {
		log.Fatalf("%qのマッチングに失敗しました: %v", globPattern, err)
	}

	for _, match := range matches {
		if err := os.RemoveAll(match); err != nil {
			log.Printf("%qの削除に失敗しました: %v", match, err)
		}
	}
}

func ExampleTempFile() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // クリーンアップ

	if _, err := tmpfile.Write(content); err != nil {
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func ExampleTempFile_suffix() {
	content := []byte("temporary file's content")
	tmpfile, err := ioutil.TempFile("", "example.*.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer os.Remove(tmpfile.Name()) // クリーンアップ

	if _, err := tmpfile.Write(content); err != nil {
		tmpfile.Close()
		log.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		log.Fatal(err)
	}
}

func ExampleReadFile() {
	content, err := ioutil.ReadFile("testdata/hello")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File contents: %s", content)

	// Output:
	// File contents: Hello, Gophers!
}

func ExampleWriteFile() {
	message := []byte("Hello, Gophers!")
	err := ioutil.WriteFile("hello", message, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
