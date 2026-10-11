// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

import (
	"github.com/shogo82148/std/sync"
)

// DLLErrorは、DLLのロードに失敗した理由を記述します。
type DLLError struct {
	Err     error
	ObjName string
	Msg     string
}

func (e *DLLError) Error() string

func (e *DLLError) Unwrap() error

// Deprecated: 代わりに[SyscallN]を使用してください。
func Syscall(trap, nargs, a1, a2, a3 uintptr) (r1, r2 uintptr, err Errno)

// Deprecated: 代わりに[SyscallN]を使用してください。
func Syscall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err Errno)

// Deprecated: 代わりに[SyscallN]を使用してください。
func Syscall9(trap, nargs, a1, a2, a3, a4, a5, a6, a7, a8, a9 uintptr) (r1, r2 uintptr, err Errno)

// Deprecated: 代わりに[SyscallN]を使用してください。
func Syscall12(trap, nargs, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12 uintptr) (r1, r2 uintptr, err Errno)

// Deprecated: 代わりに[SyscallN]を使用してください。
func Syscall15(trap, nargs, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15 uintptr) (r1, r2 uintptr, err Errno)

// Deprecated: 代わりに[SyscallN]を使用してください。
func Syscall18(trap, nargs, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12, a13, a14, a15, a16, a17, a18 uintptr) (r1, r2 uintptr, err Errno)

//go:noescape
func SyscallN(trap uintptr, args ...uintptr) (r1, r2 uintptr, err Errno)

// DLLは、単一のDLLへのアクセスを実装します。
type DLL struct {
	Name   string
	Handle Handle
}

// LoadDLLは、指定されたDLLファイルをメモリにロードします。
//
// nameが絶対パスではなく、Goが使用する既知のシステムDLLでもない場合、
// Windowsは多くの場所で指定されたDLLを検索するため、
// DLLプリロード攻撃を受ける可能性があります。
//
// システムDLLを安全にロードするには、golang.org/x/sys/windowsの[LazyDLL]を使用してください。
func LoadDLL(name string) (*DLL, error)

// MustLoadDLLは[LoadDLL]と同様ですが、ロード操作が失敗した場合はパニックします。
func MustLoadDLL(name string) *DLL

// FindProcは、[DLL] dからnameという名前のプロシージャを検索し、
// 見つかった場合は[*Proc]を返します。検索に失敗した場合はエラーを返します。
func (d *DLL) FindProc(name string) (proc *Proc, err error)

// MustFindProcは[DLL.FindProc]と同様ですが、検索に失敗した場合はパニックします。
func (d *DLL) MustFindProc(name string) *Proc

// Releaseは、[DLL] dをメモリからアンロードします。
func (d *DLL) Release() (err error)

// Procは、[DLL]内のプロシージャへのアクセスを実装します。
type Proc struct {
	Dll  *DLL
	Name string
	addr uintptr
}

// Addrは、pが表すプロシージャのアドレスを返します。
// 戻り値は、プロシージャを実行するためにSyscallに渡すことができます。
func (p *Proc) Addr() uintptr

// Callは、引数aでプロシージャpを実行します。
//
// 返されるエラーは常にnilではなく、GetLastErrorの結果から構築されます。
// 呼び出し元は、エラーを参照する前に、主な戻り値を調べて（呼び出される特定の関数の
// セマンティクスに従って）エラーが発生したかどうかを判断しなければなりません。
// エラーは常に[Errno]型です。
//
// amd64では、Callは浮動小数点値を渡したり返したりできます。C型"float"の
// 引数xを渡すには、uintptr(math.Float32bits(x))を使用します。C型"double"の
// 引数を渡すには、uintptr(math.Float64bits(x))を使用します。浮動小数点の戻り値は
// r2で返されます。C型"float"の戻り値は[math.Float32frombits](uint32(r2))です。
// C型"double"の場合は[math.Float64frombits](uint64(r2))です。
//
//go:uintptrescapes
func (p *Proc) Call(a ...uintptr) (uintptr, uintptr, error)

// LazyDLLは、単一の[DLL]へのアクセスを実装します。
// DLLのロードは、その[LazyDLL.Handle]メソッドまたはその[LazyProc]の
// Addrメソッドのいずれかが最初に呼び出されるまで遅延されます。
//
// LazyDLLは、[LoadDLL]で説明されているのと同じDLLプリロード攻撃の対象となります。
//
// システムDLLを安全にロードするには、golang.org/x/sys/windowsのLazyDLLを使用してください。
type LazyDLL struct {
	mu   sync.Mutex
	dll  *DLL
	Name string
}

// Loadは、DLLファイルd.Nameをメモリにロードします。失敗した場合はエラーを返します。
// DLLがすでにメモリにロードされている場合、LoadはDLLをロードしようとしません。
func (d *LazyDLL) Load() error

// Handleは、dのモジュールハンドルを返します。
func (d *LazyDLL) Handle() uintptr

// NewProcは、[DLL] d内の指定されたプロシージャにアクセスするための[LazyProc]を返します。
func (d *LazyDLL) NewProc(name string) *LazyProc

// NewLazyDLLは、[DLL]ファイルに関連付けられた新しい[LazyDLL]を作成します。
func NewLazyDLL(name string) *LazyDLL

// LazyProcは、[LazyDLL]内のプロシージャへのアクセスを実装します。
// 検索は、[LazyProc.Addr]、[LazyProc.Call]、[LazyProc.Find]メソッドが呼び出されるまで遅延されます。
type LazyProc struct {
	mu   sync.Mutex
	Name string
	l    *LazyDLL
	proc *Proc
}

// Findは、[DLL]からp.Nameという名前のプロシージャを検索します。
// 検索に失敗した場合はエラーを返します。プロシージャがすでに見つかっていて
// メモリにロードされている場合、Findはプロシージャを検索しません。
func (p *LazyProc) Find() error

// Addrは、pが表すプロシージャのアドレスを返します。
// 戻り値は、プロシージャを実行するためにSyscallに渡すことができます。
func (p *LazyProc) Addr() uintptr

// Callは、引数aでプロシージャpを実行します。詳細については
// Proc.Callのドキュメントを参照してください。
//
//go:uintptrescapes
func (p *LazyProc) Call(a ...uintptr) (r1, r2 uintptr, lastErr error)
