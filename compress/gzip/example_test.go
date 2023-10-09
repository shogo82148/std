// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gzip_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/compress/gzip"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/net/http/httptest"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/strings"
	"github.com/shogo82148/std/time"
)

func Example_writerReader() {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	// ヘッダーフィールドの設定はオプションです。
	zw.Name = "a-new-hope.txt"
	zw.Comment = "an epic space opera by George Lucas"
	zw.ModTime = time.Date(1977, time.May, 25, 0, 0, 0, 0, time.UTC)

	_, err := zw.Write([]byte("A long time ago in a galaxy far, far away..."))
	if err != nil {
		log.Fatal(err)
	}

	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	zr, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

	// Output:
	// Name: a-new-hope.txt
	// Comment: an epic space opera by George Lucas
	// ModTime: 1977-05-25 00:00:00 +0000 UTC
	//
	// A long time ago in a galaxy far, far away...
}

func ExampleReader_Multistream() {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)

	var files = []struct {
		name    string
		comment string
		modTime time.Time
		data    string
	}{
		{"file-1.txt", "file-header-1", time.Date(2006, time.February, 1, 3, 4, 5, 0, time.UTC), "Hello Gophers - 1"},
		{"file-2.txt", "file-header-2", time.Date(2007, time.March, 2, 4, 5, 6, 1, time.UTC), "Hello Gophers - 2"},
	}

	for _, file := range files {
		zw.Name = file.name
		zw.Comment = file.comment
		zw.ModTime = file.modTime

		if _, err := zw.Write([]byte(file.data)); err != nil {
			log.Fatal(err)
		}

		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}

		zw.Reset(&buf)
	}

	zr, err := gzip.NewReader(&buf)
	if err != nil {
		log.Fatal(err)
	}

	for {
		zr.Multistream(false)
		fmt.Printf("Name: %s\nComment: %s\nModTime: %s\n\n", zr.Name, zr.Comment, zr.ModTime.UTC())

		if _, err := io.Copy(os.Stdout, zr); err != nil {
			log.Fatal(err)
		}

		fmt.Print("\n\n")

		err = zr.Reset(&buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

	// Output:
	// Name: file-1.txt
	// Comment: file-header-1
	// ModTime: 2006-02-01 03:04:05 +0000 UTC
	//
	// Hello Gophers - 1
	//
	// Name: file-2.txt
	// Comment: file-header-2
	// ModTime: 2007-03-02 04:05:06 +0000 UTC
	//
	// Hello Gophers - 2
}

func Example_compressingReader() {

	// これは圧縮リーダーを書く例です。
	// これは、HTTPクライアントのリクエストボディに使用することができます。

	const testdata = "the data to be compressed"

	// このHTTPハンドラはテスト用です。
	handler := http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		zr, err := gzip.NewReader(req.Body)
		if err != nil {
			log.Fatal(err)
		}

		// 例のデータを出力するだけ。
		if _, err := io.Copy(os.Stdout, zr); err != nil {
			log.Fatal(err)
		}
	})
	ts := httptest.NewServer(handler)
	defer ts.Close()

	// 残りは例示コードです。

	// 圧縮したいデータは、io.Readerとして表現されます。
	dataReader := strings.NewReader(testdata)

	// bodyReaderはHTTPリクエストの本文をio.Readerとして表します。
	// httpWriterはHTTPリクエストの本文をio.Writerとして表します。
	bodyReader, httpWriter := io.Pipe()

	// bodyReaderが常に閉じられるようにすることで、以下のゴルーチンが常に終了するようにします。
	defer bodyReader.Close()

	// gzipWriterはデータをhttpWriterに圧縮します。
	gzipWriter := gzip.NewWriter(httpWriter)

	// errchは書き込みゴルーチンからのエラーを集めます。
	errch := make(chan error, 1)

	go func() {
		defer close(errch)
		sentErr := false
		sendErr := func(err error) {
			if !sentErr {
				errch <- err
				sentErr = true
			}
		}

		//  データをgzipWriterにコピーし、それを圧縮して
		//  gzipWriterからbodyReaderに供給します。
		if _, err := io.Copy(gzipWriter, dataReader); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
		if err := gzipWriter.Close(); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
		if err := httpWriter.Close(); err != nil && err != io.ErrClosedPipe {
			sendErr(err)
		}
	}()

	// テストサーバーにHTTPリクエストを送信します。
	req, err := http.NewRequest("PUT", ts.URL, bodyReader)
	if err != nil {
		log.Fatal(err)
	}

	// reqをhttp.Client.Doに渡すと、bodyReaderが閉じることが保証されます。
	resp, err := ts.Client().Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// データの圧縮中にエラーが発生したかどうかを確認する。
	if err := <-errch; err != nil {
		log.Fatal(err)
	}

	// この例では、レスポンスは気にしません。
	resp.Body.Close()

	// Output: the data to be compressed
}
