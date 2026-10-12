// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

// Fork, exec, wait, etc.

package syscall

import (
	"github.com/shogo82148/std/sync"
)

// ForkLockは、新しいファイルディスクリプターの作成をforkと同期させるために使用されます。
//
// fork/execのシーケンスでは、子に意図したファイルディスクリプターのみを継承させたいと考えます。
// そのために、すべてのファイルディスクリプターをclose-on-execとしてマークし、
// 子の中で、exec後のプログラムに保持させたいものだけを明示的にマーク解除します。
// Unixではこれは簡単ではありません。一般に、新しいファイルディスクリプターを
// close-on-execとして割り当てる方法はありません。代わりに、ディスクリプターを割り当ててから
// close-on-execとしてマークする必要があります。
// これら2つのイベントの間にforkが発生すると、子のexecは不要なファイルディスクリプターを
// 継承してしまいます。
//
// このロックはその競合を解決します。新しいfdの作成とclose-on-execのマーク操作は
// ForkLockを読み取り用に保持して行われ、fork自体はForkLockを書き込み用に保持して
// 行われます。少なくとも、それが考え方です。
// いくつかの複雑な点があります。
//
// 新しいファイルディスクリプターを作成する一部のシステムコールは、任意に長い時間
// ブロックする可能性があります。応答しないNFSサーバーや名前付きパイプでのopen、
// ソケットでのacceptなどです。それらの操作の間ずっとロックを取得することは
// 合理的ではありません。
//
// 一部のファイルディスクリプターを継承することは、他のものよりも悪い結果をもたらします。
// 悪意のない子が開いている通常のファイルを誤って継承しても、大きな問題にはなりません。
// 一方、長時間生存する子がパイプの書き込み側を誤って継承した場合、そのパイプの
// 読み取り側はその子が終了するまでEOFを検出できず、親プログラムがハングする
// 可能性があります。これは、popenを使用するスレッド化されたCプログラムでよくある問題です。
//
// 幸いなことに、継承しないことが最も重要なファイルディスクリプターは、作成に任意に長い
// 時間がかかるものではありません。pipeは即座に戻り、netパッケージはリッスンしている
// ソケットでのacceptにノンブロッキングI/Oを使用します。
// ファイルディスクリプターを作成する操作がForkLockを使用するかどうかのルールは
// 次のとおりです:
//
//   - [Pipe]。利用可能であればpipe2を使用します。そうでなければ、ブロックしないため
//     ForkLockを使用します。
//   - [Socket]。利用可能であればSOCK_CLOEXECを使用します。そうでなければ、ブロックしないため
//     ForkLockを使用します。
//   - [Open]。利用可能であれば [O_CLOEXEC] を使用します。そうでなければ、ブロックする可能性があるため
//     競合を受け入れます。
//   - [Dup]。利用可能であれば [F_DUPFD_CLOEXEC] またはdup3を使用します。そうでなければ、
//     ブロックしないためForkLockを使用します。
var ForkLock sync.RWMutex

// StringSlicePtrは、文字列のスライスをNUL終端されたバイト配列への
// ポインターのスライスに変換します。いずれかの文字列にNULバイトが含まれている場合、
// この関数はエラーを返す代わりにパニックを引き起こします。
//
// Deprecated: 代わりに [SlicePtrFromStrings] を使用してください。
func StringSlicePtr(ss []string) []*byte

// SlicePtrFromStringsは、文字列のスライスをNUL終端されたバイト配列への
// ポインターのスライスに変換します。いずれかの文字列にNULバイトが含まれている場合、
// (nil, [EINVAL])を返します。
func SlicePtrFromStrings(ss []string) ([]*byte, error)

func CloseOnExec(fd int)

func SetNonblock(fd int, nonblocking bool) (err error)

// Credentialは、[StartProcess] によって開始された子プロセスが引き受ける
// ユーザーとグループのIDを保持します。
type Credential struct {
	Uid         uint32
	Gid         uint32
	Groups      []uint32
	NoSetGroups bool
}

// ProcAttrは、[StartProcess] によって開始される新しいプロセスに適用される
// 属性を保持します。
type ProcAttr struct {
	Dir   string
	Env   []string
	Files []uintptr
	Sys   *SysProcAttr
}

// forkとexecの組み合わせで、スレッドセーフになるよう注意しています。
func ForkExec(argv0 string, argv []string, attr *ProcAttr) (pid int, err error)

// StartProcessは、osパッケージのために [ForkExec] をラップします。
func StartProcess(argv0 string, argv []string, attr *ProcAttr) (pid int, handle uintptr, err error)

// Execはexecve(2)システムコールを呼び出します。
func Exec(argv0 string, argv []string, envv []string) (err error)
