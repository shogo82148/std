// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bufio_test

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/strconv"
	"github.com/shogo82148/std/strings"
)

func ExampleWriter() {
	w := bufio.NewWriter(os.Stdout)
	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!")
	w.Flush() // フラッシュするのを忘れないで！
	// Output: Hello, world!
}

func ExampleWriter_AvailableBuffer() {
	w := bufio.NewWriter(os.Stdout)
	for _, i := range []int64{1, 2, 3, 4} {
		b := w.AvailableBuffer()
		b = strconv.AppendInt(b, i, 10)
		b = append(b, ' ')
		w.Write(b)
	}
	w.Flush()
	// Output: 1 2 3 4
}

<<<<<<< HEAD
// Scannerの最も単純な使用方法は、標準入力を行のセットとして読み取ることです。
=======
// ExampleWriter_ReadFrom demonstrates how to use the ReadFrom method of Writer.
func ExampleWriter_ReadFrom() {
	var buf bytes.Buffer
	writer := bufio.NewWriter(&buf)

	data := "Hello, world!\nThis is a ReadFrom example."
	reader := strings.NewReader(data)

	n, err := writer.ReadFrom(reader)
	if err != nil {
		fmt.Println("ReadFrom Error:", err)
		return
	}

	if err = writer.Flush(); err != nil {
		fmt.Println("Flush Error:", err)
		return
	}

	fmt.Println("Bytes written:", n)
	fmt.Println("Buffer contents:", buf.String())
	// Output:
	// Bytes written: 41
	// Buffer contents: Hello, world!
	// This is a ReadFrom example.
}

// The simplest use of a Scanner, to read standard input as a set of lines.
>>>>>>> upstream/release-branch.go1.25
func ExampleScanner_lines() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Printlnは最後の'\n'を追加します
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

// 最新のScan呼び出しを[]byteとして返します。
func ExampleScanner_Bytes() {
	scanner := bufio.NewScanner(strings.NewReader("gopher"))
	for scanner.Scan() {
		fmt.Println(len(scanner.Bytes()) == 6)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "shouldn't see an error scanning a string")
	}
	// Output:
	// true
}

// スキャナを使用して、スペースで区切られたトークンのシーケンスとして入力をスキャンすることにより、
// 単純な単語カウントユーティリティを実装します。
func ExampleScanner_words() {
	// 人工的な入力ソース。
	const input = "Now is the winter of our discontent,\nMade glorious summer by this sun of York.\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// スキャン操作のための分割関数を設定します。
	scanner.Split(bufio.ScanWords)
	// 単語を数えます。
	count := 0
	for scanner.Scan() {
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("%d\n", count)
	// Output: 15
}

// ScanWordsをラップして構築されたカスタム分割関数を使用するScannerを使用して、
// 32ビットの10進数入力を検証します。
func ExampleScanner_custom() {
	// 人工的な入力ソース。
	const input = "1234 5678 1234567901234567890"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// 既存のScanWords関数をラップして、カスタム分割関数を作成します。
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}
	// スキャン操作のための分割関数を設定します。
	scanner.Split(split)
	// 入力を検証します。
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid input: %s", err)
	}
	// Output:
	// 1234
	// 5678
	// Invalid input: strconv.ParseInt: parsing "1234567901234567890": value out of range
}

// 空の最終値を持つカンマ区切りリストを解析するために、カスタム分割関数を使用するScannerを使用します。
func ExampleScanner_emptyFinalToken() {
	// カンマ区切りのリスト。最後のエントリは空です。
	const input = "1,2,3,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))
	// コンマで区切る分割関数を定義します。
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		// 最後に配信されるトークンが1つあります。これが空の文字列である場合があります。
		// ここでbufio.ErrFinalTokenを返すと、Scanにこれ以降のトークンがないことを伝えます。
		// ただし、Scan自体からエラーが返されるわけではありません。
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	// Scan.
	for scanner.Scan() {
		fmt.Printf("%q ", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	// Output: "1" "2" "3" "4" ""
}

// Use a Scanner with a custom split function to parse a comma-separated
// list with an empty final value but stops at the token "STOP".
func ExampleScanner_earlyStop() {
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		i := bytes.IndexByte(data, ',')
		if i == -1 {
			if !atEOF {
				return 0, nil, nil
			}
			// If we have reached the end, return the last token.
			return 0, data, bufio.ErrFinalToken
		}
		// If the token is "STOP", stop the scanning and ignore the rest.
		if string(data[:i]) == "STOP" {
			return i + 1, nil, bufio.ErrFinalToken
		}
		// Otherwise, return the token before the comma.
		return i + 1, data[:i], nil
	}
	const input = "1,2,STOP,4,"
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(onComma)
	for scanner.Scan() {
		fmt.Printf("Got a token %q\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	// Output:
	// Got a token "1"
	// Got a token "2"
}
