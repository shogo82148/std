// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/io/fs"
)

// 一部の一般的なシステムコールエラーのポータブルな代替です。
//
<<<<<<< HEAD
// このパッケージから返されるエラーは、errors.Is によってこれらのエラーと比較されることがあります。
=======
// Errors returned from this package may be tested against these errors
// with [errors.Is].
>>>>>>> upstream/master
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

<<<<<<< HEAD
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
=======
// NewSyscallError returns, as an error, a new [SyscallError]
// with the given system call name and error details.
// As a convenience, if err is nil, NewSyscallError returns nil.
func NewSyscallError(syscall string, err error) error

// IsExist returns a boolean indicating whether the error is known to report
// that a file or directory already exists. It is satisfied by [ErrExist] as
// well as some syscall errors.
//
// This function predates [errors.Is]. It only supports errors returned by
// the os package. New code should use errors.Is(err, fs.ErrExist).
func IsExist(err error) bool

// IsNotExist returns a boolean indicating whether the error is known to
// report that a file or directory does not exist. It is satisfied by
// [ErrNotExist] as well as some syscall errors.
//
// This function predates [errors.Is]. It only supports errors returned by
// the os package. New code should use errors.Is(err, fs.ErrNotExist).
func IsNotExist(err error) bool

// IsPermission returns a boolean indicating whether the error is known to
// report that permission is denied. It is satisfied by [ErrPermission] as well
// as some syscall errors.
//
// This function predates [errors.Is]. It only supports errors returned by
// the os package. New code should use errors.Is(err, fs.ErrPermission).
>>>>>>> upstream/master
func IsPermission(err error) bool

// IsTimeoutは、エラーがタイムアウトが発生したことを報告することを示すかどうかを示すブール値を返します。
//
<<<<<<< HEAD
// この関数は、errors.Isやエラーがタイムアウトを示すかどうかの概念よりも前から存在しています。たとえば、UnixのエラーコードEWOULDBLOCKは、
// タイムアウトを示す場合と示さない場合があります。新しいコードでは、os.ErrDeadlineExceededなど、エラーが発生した呼び出しに適切な値を使用して
// errors.Isを使用するべきです。
=======
// This function predates [errors.Is], and the notion of whether an
// error indicates a timeout can be ambiguous. For example, the Unix
// error EWOULDBLOCK sometimes indicates a timeout and sometimes does not.
// New code should use errors.Is with a value appropriate to the call
// returning the error, such as [os.ErrDeadlineExceeded].
>>>>>>> upstream/master
func IsTimeout(err error) bool
