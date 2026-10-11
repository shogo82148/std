// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// syscallパッケージは、低レベルのオペレーティングシステムのプリミティブへの
// インターフェースを含みます。詳細は基礎となるシステムによって異なり、
// デフォルトでは、godocは現在のシステムのsyscallのドキュメントを表示します。
// godocに別のシステムのsyscallのドキュメントを表示させたい場合は、
// $GOOSと$GOARCHを目的のシステムに設定してください。例えば、
// linux/amd64でfreebsd/armのドキュメントを表示したい場合は、$GOOSを
// freebsdに、$GOARCHをarmに設定してください。
// syscallの主な用途は、"os"、"time"、"net"など、システムへのより移植性の高い
// インターフェースを提供する他のパッケージの内部です。可能であれば、
// このパッケージではなくそれらのパッケージを使用してください。
// このパッケージの関数とデータ型の詳細については、適切なオペレーティングシステムの
// マニュアルを参照してください。
// これらの呼び出しは成功を示すためにerr == nilを返します。そうでない場合、
// errは失敗を説明するオペレーティングシステムのエラーです。
// ほとんどのシステムでは、そのエラーは[Errno]型です。
//
// 注意: このパッケージで定義されている関数、型、定数のほとんどは
// [golang.org/x/sys]パッケージでも利用できます。
// そのパッケージはこのパッケージよりも多くのシステムコールをサポートしており、
// ほとんどの新しいコードは可能な限りそのパッケージを優先するべきです。
// 詳細については https://golang.org/s/go1.4-syscall を参照してください。
package syscall

// StringByteSliceは、文字列をNUL終端された[]byteに変換します。
// sにNULバイトが含まれている場合、この関数はエラーを返す代わりにパニックします。
//
// Deprecated: 代わりにByteSliceFromStringを使用してください。
func StringByteSlice(s string) []byte

// ByteSliceFromStringは、sのテキストを含むNUL終端されたバイトのスライスを返します。
// sのいずれかの位置にNULバイトが含まれている場合、(nil, [EINVAL])を返します。
func ByteSliceFromString(s string) ([]byte, error)

// StringBytePtrは、NUL終端されたバイト配列へのポインタを返します。
// sにNULバイトが含まれている場合、この関数はエラーを返す代わりにパニックします。
//
// Deprecated: 代わりに[BytePtrFromString]を使用してください。
func StringBytePtr(s string) *byte

// BytePtrFromStringは、sのテキストを含むNUL終端されたバイト配列へのポインタを返します。
// sのいずれかの位置にNULバイトが含まれている場合、(nil, [EINVAL])を返します。
func BytePtrFromString(s string) (*byte, error)

// Unixは、tsに格納されている時間を秒とナノ秒で返します。
func (ts *Timespec) Unix() (sec int64, nsec int64)

// Unixは、tvに格納されている時間を秒とナノ秒で返します。
func (tv *Timeval) Unix() (sec int64, nsec int64)

// Nanoは、tsに格納されている時間をナノ秒で返します。
func (ts *Timespec) Nano() int64

// Nanoは、tvに格納されている時間をナノ秒で返します。
func (tv *Timeval) Nano() int64

func Getpagesize() int
func Exit(code int)
