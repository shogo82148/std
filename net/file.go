// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import "github.com/shogo82148/std/os"

// FileConnは、開いているファイルfに対応するネットワーク接続のコピーを返します。
// 使用が完了したら、呼び出し元の責任でfを閉じる必要があります。
// cを閉じてもfには影響しませんし、fを閉じてもcには影響しません。
func FileConn(f *os.File) (c Conn, err error)

// FileListenerは、開いたファイルfに対応するネットワークリスナーのコピーを返します。
// lnを使用後に閉じる責任は呼び出し元にあります。
// lnを閉じるとfには影響しませんし、fを閉じるとlnにも影響しません。
func FileListener(f *os.File) (ln Listener, err error)

// FilePacketConn は、開いているファイル f に対応するパケットネットワーク接続のコピーを返します。
// 使用が終わったら f を閉じるのは呼び出し元の責任です。
// c を閉じても f には影響しませんし、f を閉じても c には影響しません。
func FilePacketConn(f *os.File) (c PacketConn, err error)
