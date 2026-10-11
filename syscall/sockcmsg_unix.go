// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

// Socket control messages

package syscall

// CmsgLenは、必要なアライメントを考慮して、[Cmsghdr]構造体のLenフィールドに
// 格納する値を返します。
func CmsgLen(datalen int) int

// CmsgSpaceは、渡されたデータ長のペイロードを持つ補助要素が
// 占めるバイト数を返します。
func CmsgSpace(datalen int) int

// SocketControlMessageは、ソケット制御メッセージを表します。
type SocketControlMessage struct {
	Header Cmsghdr
	Data   []byte
}

// ParseSocketControlMessageは、bをソケット制御メッセージの配列としてパースします。
func ParseSocketControlMessage(b []byte) ([]SocketControlMessage, error)

// UnixRightsは、開いているファイルディスクリプタの集合を、
// 別のプロセスに送信するためのソケット制御メッセージにエンコードします。
func UnixRights(fds ...int) []byte

// ParseUnixRightsは、別のプロセスからの開いているファイルディスクリプタの
// 整数配列を含むソケット制御メッセージをデコードします。
func ParseUnixRights(m *SocketControlMessage) ([]int, error)
