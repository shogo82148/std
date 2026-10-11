// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

// RawConnは、生のネットワーク接続です。
type RawConn interface {
	Control(f func(fd uintptr)) error

	Read(f func(fd uintptr) (done bool)) error

	Write(f func(fd uintptr) (done bool)) error
}

// Connは、netパッケージとosパッケージの一部の型によって実装され、
// 基礎となるファイルディスクリプタまたはハンドルへのアクセスを提供します。
type Conn interface {
	SyscallConn() (RawConn, error)
}
