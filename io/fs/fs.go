// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fsはファイルシステムへの基本的なインターフェースを定義します。
// ファイルシステムはホストオペレーティングシステムだけでなく、他のパッケージによっても提供されることがあります。
package fs

import (
	"github.com/shogo82148/std/time"
)

// FSは階層的なファイルシステムへのアクセスを提供します。
//
// FSインターフェースはファイルシステムに必要な最小限の実装です。
// ファイルシステムは追加のインターフェース、例えばReadFileFSを実装することができます。
// 追加の機能や最適化された機能を提供することができます。
type FS interface {
	// Open opens the named file.
	//
	// When Open returns an error, it should be of type *PathError
	// with the Op field set to "open", the Path field set to name,
	// and the Err field describing the problem.
	//
	// Open should reject attempts to open names that do not satisfy
	// ValidPath(name), returning a *PathError with Err set to
	// ErrInvalid or ErrNotExist.
	Open(name string) (File, error)
}

// ValidPathは与えられたパス名がOpenの呼び出しに使用するために有効かどうかを報告します。
//
// Openに渡されるパス名はUTF-8でエンコードされた、ルートなしのスラッシュで区切られたパス要素のシーケンス（例: "x/y/z"）です。
// パス名には、"."または".."または空の文字列を含めることはできませんが、ルートディレクトリが "."という特殊なケースを除いてはです。
// パスはスラッシュで始まることや終わることはできません: "/x"や"x/"は無効です。
//
// なお、パスは全てのシステムでスラッシュで区切られます（Windowsでも）。
// バックスラッシュやコロンなどの他の文字を含むパスも有効ですが、これらの文字は実装によっては絶対にパス要素の区切りとして解釈されるべきではありません。
func ValidPath(name string) bool

// Fileは単一のファイルへのアクセスを提供します。
// Fileインターフェースはファイルに必要な最小限の実装です。
// ディレクトリファイルはReadDirFileも実装する必要があります。
// ファイルは最適化としてio.ReaderAtまたはio.Seekerを実装する場合があります。
type File interface {
	Stat() (FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}

// DirEntryはディレクトリから読み取られたエントリです
// (ReadDir関数やReadDirFileのReadDirメソッドを使用して)。
type DirEntry interface {
	// Name returns the name of the file (or subdirectory) described by the entry.
	// This name is only the final element of the path (the base name), not the entire path.
	// For example, Name would return "hello.go" not "home/gopher/hello.go".
	Name() string

	// IsDir reports whether the entry describes a directory.
	IsDir() bool

	// Type returns the type bits for the entry.
	// The type bits are a subset of the usual FileMode bits, those returned by the FileMode.Type method.
	Type() FileMode

	// Info returns the FileInfo for the file or subdirectory described by the entry.
	// The returned FileInfo may be from the time of the original directory read
	// or from the time of the call to Info. If the file has been removed or renamed
	// since the directory read, Info may return an error satisfying errors.Is(err, ErrNotExist).
	// If the entry denotes a symbolic link, Info reports the information about the link itself,
	// not the link's target.
	Info() (FileInfo, error)
}

// ReadDirFileは、ReadDirメソッドを使用してエントリを読み取ることができるディレクトリファイルです。
// すべてのディレクトリファイルは、このインターフェースを実装する必要があります。
// （任意のファイルがこのインターフェースを実装することも許可されていますが、非ディレクトリの場合はReadDirがエラーを返すべきです。）
type ReadDirFile interface {
	File

	// ReadDir reads the contents of the directory and returns
	// a slice of up to n DirEntry values in directory order.
	// Subsequent calls on the same file will yield further DirEntry values.
	//
	// If n > 0, ReadDir returns at most n DirEntry structures.
	// In this case, if ReadDir returns an empty slice, it will return
	// a non-nil error explaining why.
	// At the end of a directory, the error is io.EOF.
	// (ReadDir must return io.EOF itself, not an error wrapping io.EOF.)
	//
	// If n <= 0, ReadDir returns all the DirEntry values from the directory
	// in a single slice. In this case, if ReadDir succeeds (reads all the way
	// to the end of the directory), it returns the slice and a nil error.
	// If it encounters an error before the end of the directory,
	// ReadDir returns the DirEntry list read until that point and a non-nil error.
	ReadDir(n int) ([]DirEntry, error)
}

// 汎用ファイルシステムのエラー。
// ファイルシステムから返されるエラーは、これらのエラーと比較してテストすることができます
// errors.Is を使用して。
var (
	ErrInvalid    = errInvalid()
	ErrPermission = errPermission()
	ErrExist      = errExist()
	ErrNotExist   = errNotExist()
	ErrClosed     = errClosed()
)

// FileInfoはファイルを説明し、Statによって返されます。
type FileInfo interface {
	Name() string
	Size() int64
	Mode() FileMode
	ModTime() time.Time
	IsDir() bool
	Sys() any
}

// FileModeはファイルのモードとパーミッションビットを表します。
// ビットの定義はすべてのシステムで同じであるため、
// ファイルに関する情報をポータブルに他のシステムに移動することができます。
// すべてのビットがすべてのシステムに適用されるわけではありません。
// ディレクトリに対してはModeDirのみが必須です。
type FileMode uint32

// 定義されたファイルモードビットは、FileModeの最も重要なビットです。
// 9つの最も下位のビットは、標準のUnixのrwxrwxrwx権限です。
// これらのビットの値は、パブリックAPIの一部と見なされ、
// ワイヤープロトコルやディスク表現で使用される可能性があります。
// これらのビットは変更しないでくださいが、新しいビットが追加されることはあります。
const (

	// 単一の文字は、Stringメソッドのフォーマットで使用される省略形です。
	ModeDir FileMode = 1 << (32 - 1 - iota)
	ModeAppend
	ModeExclusive
	ModeTemporary
	ModeSymlink
	ModeDevice
	ModeNamedPipe
	ModeSocket
	ModeSetuid
	ModeSetgid
	ModeCharDevice
	ModeSticky
	ModeIrregular

	// タイプビットのマスク。通常のファイルでは、全く設定されません。
	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice | ModeCharDevice | ModeIrregular

	ModePerm FileMode = 0777
)

func (m FileMode) String() string

// IsDirはmがディレクトリを記述しているかどうかを報告します。
// つまり、m内のModeDirビットがセットされているかどうかをテストします。
func (m FileMode) IsDir() bool

// IsRegularはmが正規のファイルを記述しているかどうかを報告します。
// つまり、モードのタイプビットが設定されていないかどうかをテストします。
func (m FileMode) IsRegular() bool

// Permは、m（m＆ModePerm）のUnixパーミッションビットを返します。
func (m FileMode) Perm() FileMode

// Typeはm（m＆ModeType）のタイプビットを返します。
func (m FileMode) Type() FileMode

// PathErrorはエラーとそれを引き起こした操作とファイルパスを記録します。
type PathError struct {
	Op   string
	Path string
	Err  error
}

func (e *PathError) Error() string

func (e *PathError) Unwrap() error

// Timeoutは、このエラーがタイムアウトを示すかどうかを報告します。
func (e *PathError) Timeout() bool
