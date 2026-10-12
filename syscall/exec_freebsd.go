// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

type SysProcAttr struct {
	Chroot     string
	Credential *Credential
	Ptrace     bool
	Setsid     bool
	// Setpgidは、子のプロセスグループIDをPgidに設定します。
	// Pgid == 0の場合は、新しい子のプロセスIDに設定します。
	Setpgid bool
	// Setcttyは、子の制御端末をファイルディスクリプターCttyに設定します。
	// Cttyは子プロセス内のディスクリプター番号、つまりProcAttr.Filesへの
	// インデックスでなければなりません。
	// これはSetsidがtrueの場合にのみ意味を持ちます。
	Setctty bool
	Noctty  bool
	Ctty    int
	// Foregroundは、子プロセスグループをフォアグラウンドに配置します。
	// これはSetpgidを暗黙的に含みます。Cttyフィールドには
	// 制御TTYのディスクリプターを設定しなければなりません。
	// Setcttyとは異なり、この場合Cttyは親プロセス内の
	// ディスクリプター番号でなければなりません。
	Foreground bool
	Pgid       int
	Pdeathsig  Signal
	Jail       int
}
