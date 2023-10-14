// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/io/fs"
)

// 一部の一般的なシステムコールエラーのポータブルな代替です。
//
// このパッケージから返されるエラーは、errors.Is によってこれらのエラーと比較されることがあります。
var (

	// ErrInvalidは無効な引数を示します。
	// Fileのメソッドは、レシーバーがnilの場合にこのエラーを返します。
	ErrInvalid = fs.ErrInvalid

	ErrPermission = fs.ErrPermission
	ErrExist      = fs.ErrExist
	ErrNotExist   = fs.ErrNotExist
	ErrClosed     = fs.ErrClosed

	ErrNoDeadline       = errNoDeadline()
	ErrDeadlineExceeded = errDeadlineExceeded()
)

// PathErrorはエラーメッセージとそれを引き起こした操作とファイルパスを記録します。
type PathError = fs.PathError

// SyscallErrorは特定のシステムコールからのエラーを記録します。
type SyscallError struct {
	Syscall string
	Err     error
}

func (e *SyscallError) Error() string

func (e *SyscallError) Unwrap() error

// Timeoutは、このエラーがタイムアウトを表すかどうかを報告します。
func (e *SyscallError) Timeout() bool

// NewSyscallErrorは、指定されたシステムコール名とエラーの詳細を持つ新しいSyscallErrorをエラーとして返します。
// 便利な機能として、errがnilの場合、NewSyscallErrorはnilを返します。
func NewSyscallError(syscall string, err error) error

// IsExistは、ファイルまたはディレクトリが既に存在することを報告する既知のエラーを示すブール値を返します。ErrExistで満たされるだけでなく、一部のシスコールエラーでも満たされます。
// この関数はerrors.Isより前に存在しています。それはosパッケージによって返されるエラーのみをサポートしています。新しく書かれるコードでは、errors.Is(err, fs.ErrExist)を使用するべきです。
func IsExist(err error) bool

// IsNotExistは、ファイルやディレクトリが存在しないことを報告することがわかっているエラーかどうかを示すブール値を返します。これは、ErrNotExistと一部のsyscallエラーに合致します。
// この関数は、errors.Isよりも前に存在していました。これは、osパッケージによって返されるエラーのみをサポートしています。新しいコードでは、errors.Is(err, fs.ErrNotExist)を使用する必要があります。
func IsNotExist(err error) bool

// IsPermissionは、エラーがパーミッションが拒否されたことを報告する既知のものであるかどうかを示すブール値を返します。
// ErrPermissionだけでなく、一部のsyscallエラーも対応しています。
//
// この関数はerrors.Isより前に存在しています。この関数はosパッケージが返すエラーのみをサポートしています。
// 新しいコードではerrors.Is(err、fs.ErrPermission)を使用するべきです。
func IsPermission(err error) bool

// IsTimeoutは、エラーがタイムアウトが発生したことを報告することを示すかどうかを示すブール値を返します。
//
// この関数は、errors.Isやエラーがタイムアウトを示すかどうかの概念よりも前から存在しています。たとえば、UnixのエラーコードEWOULDBLOCKは、
// タイムアウトを示す場合と示さない場合があります。新しいコードでは、os.ErrDeadlineExceededなど、エラーが発生した呼び出しに適切な値を使用して
// errors.Isを使用するべきです。
func IsTimeout(err error) bool
