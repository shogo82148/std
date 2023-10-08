// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip_test

import (
	"github.com/shogo82148/std/archive/zip"
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/compress/flate"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/os"
)

func ExampleWriter() {
	// アーカイブを書き込むためのバッファを作成します。
	buf := new(bytes.Buffer)

	// 新しい zip アーカイブを作成します。
	w := zip.NewWriter(buf)

	// アーカイブにいくつかのファイルを追加します。
	var files = []struct {
		Name, Body string
	}{
		{"readme.txt", "このアーカイブにはいくつかのテキストファイルが含まれています。"},
		{"gopher.txt", "Gopher names:\nGeorge\nGeoffrey\nGonzo"},
		{"todo.txt", "動物取扱業免許を取得する。\nもっと例を書く。"},
	}
	for _, file := range files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// Close でエラーを確認することを忘れないでください。
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func ExampleReader() {
	// 読み取り用に zip アーカイブを開きます。
	r, err := zip.OpenReader("testdata/readme.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// アーカイブ内のファイルを反復処理し、その内容の一部を出力します。
	for _, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.CopyN(os.Stdout, rc, 68)
		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		fmt.Println()
	}
	// Output:
	// Contents of README:
	// This is the source code repository for the Go programming language.
}

func ExampleWriter_RegisterCompressor() {
	// デフォルトの Deflate 圧縮プログラムを、より高い圧縮レベルのカスタム圧縮プログラムで上書きします。

	// アーカイブを書き込むためのバッファを作成します。
	buf := new(bytes.Buffer)

	// 新しい zip アーカイブを作成します。
	w := zip.NewWriter(buf)

	// カスタムの Deflate 圧縮プログラムを登録します。
	w.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		return flate.NewWriter(out, flate.BestCompression)
	})

	// ファイルを w に追加します。
}
