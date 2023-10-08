// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package flate_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/compress/flate"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/strings"
	"github.com/shogo82148/std/sync"
)

// パフォーマンスの重要なアプリケーションでは、Reset を使用して現在の圧縮器または伸張器の状態を破棄し、
// 以前に割り当てられたメモリを活用してそれらを迅速に再初期化することができます。
func Example_reset() {
	proverbs := []string{
		"Don't communicate by sharing memory, share memory by communicating.\n",
		"Concurrency is not parallelism.\n",
		"The bigger the interface, the weaker the abstraction.\n",
		"Documentation is for users.\n",
	}

	var r strings.Reader
	var b bytes.Buffer
	buf := make([]byte, 32<<10)

	zw, err := flate.NewWriter(nil, flate.DefaultCompression)
	if err != nil {
		log.Fatal(err)
	}
	zr := flate.NewReader(nil)

	for _, s := range proverbs {
		r.Reset(s)
		b.Reset()

		// コンプレッサーをリセットし、入力ストリームからエンコードします。
		zw.Reset(&b)
		if _, err := io.CopyBuffer(zw, &r, buf); err != nil {
			log.Fatal(err)
		}
		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}

		// デコンプレッサをリセットし、いくつかの出力ストリームにデコードします。
		if err := zr.(flate.Resetter).Reset(&b, nil); err != nil {
			log.Fatal(err)
		}
		if _, err := io.CopyBuffer(os.Stdout, zr, buf); err != nil {
			log.Fatal(err)
		}
		if err := zr.Close(); err != nil {
			log.Fatal(err)
		}
	}

	// Output:
	// Don't communicate by sharing memory, share memory by communicating.
	// Concurrency is not parallelism.
	// The bigger the interface, the weaker the abstraction.
	// Documentation is for users.
}

// あらかじめ設定された辞書を使用すると、圧縮率を改善することができます。
// 辞書を使用する際の難点は、圧縮機と展開機が事前に使用する辞書について合意する必要があるということです。
func Example_dictionary() {

	// 辞書はバイトの連続です。入力データを圧縮する際、圧縮器は辞書内で見つかった一致する部分文字列を代替しようとします。そのため、辞書には実際のデータストリームで見つかることが期待される部分文字列のみを含めるべきです。
	const dict = `<?xml version="1.0"?>` + `<book>` + `<data>` + `<meta name="` + `" content="`

	// 圧縮するデータには、辞書と一致する頻繁な（必要ではありませんが）部分文字列が含まれることが望ましいです。
	const data = `<?xml version="1.0"?>
<book>
	<meta name="title" content="The Go Programming Language"/>
	<meta name="authors" content="Alan Donovan and Brian Kernighan"/>
	<meta name="published" content="2015-10-26"/>
	<meta name="isbn" content="978-0134190440"/>
	<data>...</data>
</book>
`

	var b bytes.Buffer

	// 特殊に作られた辞書を使用してデータを圧縮する。
	zw, err := flate.NewWriterDict(&b, flate.DefaultCompression, []byte(dict))
	if err != nil {
		log.Fatal(err)
	}
	if _, err := io.Copy(zw, strings.NewReader(data)); err != nil {
		log.Fatal(err)
	}
	if err := zw.Close(); err != nil {
		log.Fatal(err)
	}

	// 解凍プログラムは圧縮プログラムと同じ辞書を使用する必要があります。
	// そうでないと、入力が破損しているように見えるかもしれません。
	fmt.Println("Decompressed output using the dictionary:")
	zr := flate.NewReaderDict(bytes.NewReader(b.Bytes()), []byte(dict))
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}
	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	// 辞書のすべてのバイトを '#' に置き換えて、予め設定された辞書の近似効果を視覚的に示します。
	fmt.Println("Substrings matched by the dictionary are marked with #:")
	hashDict := []byte(dict)
	for i := range hashDict {
		hashDict[i] = '#'
	}
	zr = flate.NewReaderDict(&b, hashDict)
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}
	if err := zr.Close(); err != nil {
		log.Fatal(err)
	}

	// Output:
	// Decompressed output using the dictionary:
	// <?xml version="1.0"?>
	// <book>
	// 	<meta name="title" content="The Go Programming Language"/>
	// 	<meta name="authors" content="Alan Donovan and Brian Kernighan"/>
	// 	<meta name="published" content="2015-10-26"/>
	// 	<meta name="isbn" content="978-0134190440"/>
	// 	<data>...</data>
	// </book>
	//
	// Substrings matched by the dictionary are marked with #:
	// #####################
	// ######
	// 	############title###########The Go Programming Language"/#
	// 	############authors###########Alan Donovan and Brian Kernighan"/#
	// 	############published###########2015-10-26"/#
	// 	############isbn###########978-0134190440"/#
	// 	######...</#####
	// </#####
}

// DEFLATEはネットワーク上で圧縮データを送信するのに適しています。
func Example_synchronization() {
	var wg sync.WaitGroup
	defer wg.Wait()

	// io.Pipeを使用してネットワーク接続をシミュレートします。
	// 実際のネットワークアプリケーションでは、基礎となる接続を適切に閉じる必要があります。
	rp, wp := io.Pipe()

	// 送信機能として機能するために、ゴールーチンを開始します。
	wg.Add(1)
	go func() {
		defer wg.Done()

		zw, err := flate.NewWriter(wp, flate.BestSpeed)
		if err != nil {
			log.Fatal(err)
		}

		b := make([]byte, 256)
		for _, m := range strings.Fields("A long time ago in a galaxy far, far away...") {

			// 最初のバイトがメッセージの長さであり、その後にメッセージ自体が続く、単純なフレーム形式を使用しています。
			b[0] = uint8(copy(b[1:], m))

			if _, err := zw.Write(b[:1+len(m)]); err != nil {
				log.Fatal(err)
			}

			// Flushは、受信者がこれまでに送信されたすべてのデータを読み取ることができることを保証します。
			if err := zw.Flush(); err != nil {
				log.Fatal(err)
			}
		}

		if err := zw.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// 受信者として動作するゴルーチンを開始する。
	wg.Add(1)
	go func() {
		defer wg.Done()

		zr := flate.NewReader(rp)

		b := make([]byte, 256)
		for {

			// メッセージの長さを読み取ります。
			// これは送信側のFlushとCloseに対して
			// 必ず返されることが保証されています。
			if _, err := io.ReadFull(zr, b[:1]); err != nil {
				if err == io.EOF {
					break // 送信者がストリームを閉じました
				}
				log.Fatal(err)
			}

			// メッセージの内容を読み取る。
			n := int(b[0])
			if _, err := io.ReadFull(zr, b[:n]); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Received %d bytes: %s\n", n, b[:n])
		}
		fmt.Println()

		if err := zr.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Output:
	// Received 1 bytes: A
	// Received 4 bytes: long
	// Received 4 bytes: time
	// Received 3 bytes: ago
	// Received 2 bytes: in
	// Received 1 bytes: a
	// Received 6 bytes: galaxy
	// Received 4 bytes: far,
	// Received 3 bytes: far
	// Received 7 bytes: away...
}
