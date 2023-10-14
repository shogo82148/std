// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// Pipeは同期的な、インメモリ、フルダプレックスのネットワーク接続を作成します。
// 両端はConnインターフェースを実装しています。
// 一方の端での読み取りは、もう一方の端での書き込みと一致し、データを直接コピーします。
// 内部バッファリングはありません。
func Pipe() (Conn, Conn)
