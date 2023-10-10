// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1 || windows

package net

// SetUnlinkOnCloseは、リスナーがクローズされたときに基礎となるソケットファイルを
// ファイルシステムから削除するかどうかを設定します。
//
// デフォルトの動作では、ソケットファイルは、package netによって作成された場合にのみ
// アンリンクされます。つまり、リスナーや基礎となるソケットファイルが
// ListenまたはListenUnixの呼び出しによって作成された場合、デフォルトでは
// リスナーをクローズするとソケットファイルが削除されます。
// ただし、リスナーが既存のソケットファイルを使用するためにFileListenerを呼び出して作成された場合、
// デフォルトではリスナーをクローズしてもソケットファイルは削除されません。
func (l *UnixListener) SetUnlinkOnClose(unlink bool)
