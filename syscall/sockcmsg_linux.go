// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Socket control messages

package syscall

// UnixCredentialsは、別のプロセスに送信するために資格情報をソケット制御メッセージに
// エンコードします。これは認証に使用できます。
func UnixCredentials(ucred *Ucred) []byte

// ParseUnixCredentialsは、Ucred構造体の資格情報を含むソケット制御メッセージを
// デコードします。そのようなメッセージを受信するには、ソケットで
// SO_PASSCREDオプションを有効にする必要があります。
func ParseUnixCredentials(m *SocketControlMessage) (*Ucred, error)
