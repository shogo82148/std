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
	Open(name string) (File, error)
}

// ValidPathは与えられたパス名がOpenの呼び出しに使用するために有効かどうかを報告します。
//
// Openに渡されるパス名はUTF-8でエンコードされた、ルートなしのスラッシュで区切られたパス要素のシーケンス（例: "x/y/z"）です。
// パス名には、"."または".."または空の文字列を含めることはできませんが、ルートディレクトリが "."という特殊なケースを除いてはです。
// パスはスラッシュで始まることや終わることはできません: "/x"や"x/"は無効です。
//
<<<<<<< HEAD
// Note that paths are slash-separated on all systems, even Windows.
// Paths containing other characters such as backslash and colon
// are accepted as valid, but those characters must never be
// interpreted by an [FS] implementation as path element separators.
func ValidPath(name string) bool

// A File provides access to a single file.
// The File interface is the minimum implementation required of the file.
// Directory files should also implement [ReadDirFile].
// A file may implement [io.ReaderAt] or [io.Seeker] as optimizations.
=======
// なお、パスは全てのシステムでスラッシュで区切られます（Windowsでも）。
// バックスラッシュやコロンなどの他の文字を含むパスも有効ですが、これらの文字は実装によっては絶対にパス要素の区切りとして解釈されるべきではありません。
func ValidPath(name string) bool

// Fileは単一のファイルへのアクセスを提供します。
// Fileインターフェースはファイルに必要な最小限の実装です。
// ディレクトリファイルはReadDirFileも実装する必要があります。
// ファイルは最適化としてio.ReaderAtまたはio.Seekerを実装する場合があります。
>>>>>>> release-branch.go1.21
type File interface {
	Stat() (FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}

<<<<<<< HEAD
// A DirEntry is an entry read from a directory
// (using the ReadDir function or a [ReadDirFile]'s ReadDir method).
=======
// DirEntryはディレクトリから読み取られたエントリです
// (ReadDir関数やReadDirFileのReadDirメソッドを使用して)。
>>>>>>> release-branch.go1.21
type DirEntry interface {
	Name() string

	IsDir() bool

	Type() FileMode

	Info() (FileInfo, error)
}

// ReadDirFileは、ReadDirメソッドを使用してエントリを読み取ることができるディレクトリファイルです。
// すべてのディレクトリファイルは、このインターフェースを実装する必要があります。
// （任意のファイルがこのインターフェースを実装することも許可されていますが、非ディレクトリの場合はReadDirがエラーを返すべきです。）
type ReadDirFile interface {
	File

	ReadDir(n int) ([]DirEntry, error)
}

<<<<<<< HEAD
// Generic file system errors.
// Errors returned by file systems can be tested against these errors
// using [errors.Is].
=======
// 汎用ファイルシステムのエラー。
// ファイルシステムから返されるエラーは、これらのエラーと比較してテストすることができます
// errors.Is を使用して。
>>>>>>> release-branch.go1.21
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

<<<<<<< HEAD
// A FileMode represents a file's mode and permission bits.
// The bits have the same definition on all systems, so that
// information about files can be moved from one system
// to another portably. Not all bits apply to all systems.
// The only required bit is [ModeDir] for directories.
type FileMode uint32

// The defined file mode bits are the most significant bits of the [FileMode].
// The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
// The values of these bits should be considered part of the public API and
// may be used in wire protocols or disk representations: they must not be
// changed, although new bits might be added.
=======
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
>>>>>>> release-branch.go1.21
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
