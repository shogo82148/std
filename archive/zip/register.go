// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

import (
	"github.com/shogo82148/std/io"
)

// Compressor は、w に書き込む新しい圧縮ライターを返します。
// WriteCloser の Close メソッドは、保留中のデータを w にフラッシュするために使用する必要があります。
// Compressor 自体は、複数のゴルーチンから同時に呼び出されることができますが、
// 各返されたライターは一度に1つのゴルーチンによってのみ使用されます。
type Compressor func(w io.Writer) (io.WriteCloser, error)

// Decompressor は、r から読み取る新しい解凍リーダーを返します。
// [io.ReadCloser] の Close メソッドは、関連するリソースを解放するために使用する必要があります。
// Decompressor 自体は、複数のゴルーチンから同時に呼び出されることができますが、
// 各返されたリーダーは一度に1つのゴルーチンによってのみ使用されます。
type Decompressor func(r io.Reader) io.ReadCloser

// RegisterDecompressor は、特定のメソッド ID にカスタムの解凍プログラムを登録または上書きします。
// メソッドの解凍プログラムが見つからない場合、Writer はパッケージレベルで解凍プログラムを検索します。
// 一般的なメソッド [Store] と [Deflate] は組み込みです。
func RegisterDecompressor(method uint16, dcomp Decompressor)

// RegisterCompressor は、特定のメソッド ID にカスタムの圧縮プログラムを登録または上書きします。
// 一般的なメソッド [Store] と [Deflate] は組み込みです。
func RegisterCompressor(method uint16, comp Compressor)
