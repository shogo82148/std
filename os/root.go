// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/time"
)

// OpenInRootは、ディレクトリーdir内のファイルnameを開きます。
// これは、OpenRoot(dir)の後にルート内でファイルを開くことと同等です。
//
// nameのいずれかの構成要素がdirの外部の場所を参照している場合、
// OpenInRootはエラーを返します。
//
// 詳細と制限については [Root] を参照してください。
func OpenInRoot(dir, name string) (*File, error)

// Rootは、単一のディレクトリーツリー内のファイルにのみアクセスするために使用できます。
//
// Rootのメソッドは、ルートディレクトリー配下のファイルとディレクトリーにのみアクセスできます。
// Rootのメソッドに渡されたファイル名のいずれかの構成要素がルートの外部の場所を
// 参照している場合、メソッドはエラーを返します。
// ファイル名はディレクトリー自体（.）を参照することができます。
//
// Rootのメソッドはシンボリックリンクをたどりますが、シンボリックリンクは
// ルートの外部の場所を参照できません。
// シンボリックリンクは絶対パスであってはなりません。
//
// Rootのメソッドは、ファイルシステムの境界、Linuxのバインドマウント、
// /procの特殊ファイルの横断や、Unixデバイスファイルへのアクセスを禁止しません。
//
// Rootのメソッドは、複数のゴルーチンから同時に使用しても安全です。
//
// ほとんどのプラットフォームでは、Rootを作成すると、そのディレクトリーを参照する
// ファイルディスクリプターまたはハンドルが開かれます。ディレクトリーが移動された場合、
// Rootのメソッドは新しい場所にある元のディレクトリーを参照します。
//
// Rootの動作は一部のプラットフォームで異なります:
//
//   - GOOS=windowsの場合、ファイル名はNULやCOM1などのWindowsの予約デバイス名を
//     参照できません。
//   - Unixでは、[Root.Chmod]、[Root.Chown]、[Root.Chtimes] は競合状態に対して脆弱です。
//     操作の進行中に操作の対象が通常のファイルからシンボリックリンクに変更された場合、
//     操作はリンク先ではなくリンク自体に対して実行される可能性があります。
//   - GOOS=jsの場合、Rootはシンボリックリンクの検証においてTOCTOU（time-of-check-time-of-use）
//     攻撃に対して脆弱であり、操作がルートから抜け出さないことを保証できません。
//   - GOOS=plan9またはGOOS=jsの場合、Rootは名前の変更をまたいでディレクトリーを追跡しません。
//     これらのプラットフォームでは、Rootはファイルディスクリプターではなくディレクトリー名を参照します。
//   - WASI preview 1（GOOS=wasip1）は [Root.Chmod] をサポートしていません。
type Root struct {
	root *root
}

// OpenRootは指定されたディレクトリーを開きます。
// ディレクトリー名のシンボリックリンクをたどります。
// エラーがある場合、それは [*PathError] 型になります。
func OpenRoot(name string) (*Root, error)

// Nameは、OpenRootに渡されたディレクトリーの名前を返します。
//
// [Close] の後にNameを呼び出しても安全です。
func (r *Root) Name() string

// CloseはRootを閉じます。
// Closeが呼び出された後、Rootのメソッドはエラーを返します。
func (r *Root) Close() error

// Openは、ルート内の指定されたファイルを読み取り用に開きます。
// 詳細は [Open] を参照してください。
func (r *Root) Open(name string) (*File, error)

// Createは、ルート内の指定されたファイルを作成または切り詰めます。
// 詳細は [Create] を参照してください。
func (r *Root) Create(name string) (*File, error)

// OpenFileは、ルート内の指定されたファイルを開きます。
// 詳細は [OpenFile] を参照してください。
//
// permに下位9ビット（0o777）以外のビットが含まれている場合、
// OpenFileはエラーを返します。
func (r *Root) OpenFile(name string, flag int, perm FileMode) (*File, error)

// OpenRootは、ルート内の指定されたディレクトリーを開きます。
// エラーがある場合、それは [*PathError] 型になります。
func (r *Root) OpenRoot(name string) (*Root, error)

// Chmodは、ルート内の指定されたファイルのモードをmodeに変更します。
// 詳細は [Chmod] を参照してください。
func (r *Root) Chmod(name string, mode FileMode) error

// Mkdirは、指定された名前とパーミッションビット（umask適用前）で、
// ルート内に新しいディレクトリーを作成します。
// 詳細は [Mkdir] を参照してください。
//
// permに下位9ビット（0o777）以外のビットが含まれている場合、
// Mkdirはエラーを返します。
func (r *Root) Mkdir(name string, perm FileMode) error

// MkdirAllは、必要な親ディレクトリーとともに、ルート内に新しいディレクトリーを作成します。
// 詳細は [MkdirAll] を参照してください。
//
// permに下位9ビット（0o777）以外のビットが含まれている場合、
// MkdirAllはエラーを返します。
func (r *Root) MkdirAll(name string, perm FileMode) error

// Chownは、ルート内の指定されたファイルの数値のuidとgidを変更します。
// 詳細は [Chown] を参照してください。
func (r *Root) Chown(name string, uid, gid int) error

// Lchownは、ルート内の指定されたファイルの数値のuidとgidを変更します。
// 詳細は [Lchown] を参照してください。
func (r *Root) Lchown(name string, uid, gid int) error

// Chtimesは、ルート内の指定されたファイルのアクセス時刻と更新時刻を変更します。
// 詳細は [Chtimes] を参照してください。
func (r *Root) Chtimes(name string, atime time.Time, mtime time.Time) error

// Removeは、ルート内の指定されたファイルまたは（空の）ディレクトリーを削除します。
// 詳細は [Remove] を参照してください。
func (r *Root) Remove(name string) error

// RemoveAllは、指定されたファイルまたはディレクトリーと、それに含まれるすべての子を削除します。
// 詳細は [RemoveAll] を参照してください。
func (r *Root) RemoveAll(name string) error

// Statは、ルート内の指定されたファイルを記述する [FileInfo] を返します。
// 詳細は [Stat] を参照してください。
func (r *Root) Stat(name string) (FileInfo, error)

// Lstatは、ルート内の指定されたファイルを記述する [FileInfo] を返します。
// ファイルがシンボリックリンクの場合、返されるFileInfoは
// シンボリックリンクを記述します。
// 詳細は [Lstat] を参照してください。
func (r *Root) Lstat(name string) (FileInfo, error)

// Readlinkは、ルート内の指定されたシンボリックリンクのリンク先を返します。
// 詳細は [Readlink] を参照してください。
func (r *Root) Readlink(name string) (string, error)

// Renameは、oldnameをnewnameに名前変更（移動）します。
// 両方のパスはルートからの相対パスです。
// 詳細は [Rename] を参照してください。
func (r *Root) Rename(oldname, newname string) error

// Linkは、oldnameファイルへのハードリンクとしてnewnameを作成します。
// 両方のパスはルートからの相対パスです。
// 詳細は [Link] を参照してください。
//
// oldnameがシンボリックリンクの場合、Linkはそのリンク先ではなくoldnameへの新しいリンクを作成します。
// この動作は、一部のプラットフォームでは [Link] の動作と異なる場合があります。
//
// GOOS=jsの場合、oldnameがシンボリックリンクであればLinkはエラーを返します。
func (r *Root) Link(oldname, newname string) error

// Symlinkは、oldnameへのシンボリックリンクとしてnewnameを作成します。
// 詳細は [Symlink] を参照してください。
//
// Symlinkはoldnameを検証しないため、
// oldnameはルートの外部の場所を参照する可能性があります。
//
// Windowsでは、oldnameがルート内のディレクトリーを参照している場合はディレクトリーリンクが
// 作成されます。それ以外の場合はファイルリンクが作成されます。
func (r *Root) Symlink(oldname, newname string) error

// ReadFileは、ルート内の指定されたファイルを読み取り、その内容を返します。
// 詳細は [ReadFile] を参照してください。
func (r *Root) ReadFile(name string) ([]byte, error)

// WriteFileは、ルート内の指定されたファイルにデータを書き込みます。必要に応じてファイルを作成します。
// 詳細は [WriteFile] を参照してください。
func (r *Root) WriteFile(name string, data []byte, perm FileMode) error

// FSは、ルート内のファイルツリーのためのファイルシステム（fs.FS）を返します。
//
// 結果は [io/fs.StatFS]、[io/fs.ReadFileFS]、[io/fs.ReadDirFS]、
// [io/fs.ReadLinkFS] を実装します。
func (r *Root) FS() fs.FS
