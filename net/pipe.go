// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

<<<<<<< HEAD
// Pipeは同期的な、インメモリ、フルダプレックスのネットワーク接続を作成します。
// 両端はConnインターフェースを実装しています。
// 一方の端での読み取りは、もう一方の端での書き込みと一致し、データを直接コピーします。
// 内部バッファリングはありません。
=======
// Pipe creates a synchronous, in-memory, full duplex
// network connection; both ends implement the [Conn] interface.
// Reads on one end are matched with writes on the other,
// copying data directly between the two; there is no internal
// buffering.
>>>>>>> upstream/release-branch.go1.22
func Pipe() (Conn, Conn)
