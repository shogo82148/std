// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// SubFSはSubメソッドを持つファイルシステムです。
type SubFS interface {
	FS

	Sub(dir string) (FS, error)
}

// Subはfsysのdirでルートされた部分木に対応する [FS] を返します。
//
// If dir is ".", Sub returns fsys unchanged.
// Otherwise, if fsys implements [SubFS], Sub returns fsys.Sub(dir).
// Otherwise, Sub returns a new [FS] implementation sub that,
// in effect, implements sub.Open(name) as fsys.Open(path.Join(dir, name)).
// The implementation also translates calls to ReadDir, ReadFile,
// ReadLink, Lstat, and Glob appropriately. Sub does not check if the
// directory currently exists.
//
// Sub(os.DirFS("/"), "prefix")はos.DirFS("/prefix")と等価ですが、
// どちらも"/prefix"の外でのオペレーティングシステムアクセスを避けることを
// 保証しません。これは [os.DirFS] の実装が"/prefix"内で他のディレクトリを
// 指すシンボリックリンクをチェックしないためです。つまり、[os.DirFS] は
// chroot方式のセキュリティメカニズムの一般的な代替ではなく、Subはその
// 事実を変えません。特定のディレクトリツリーへのアクセスを制限するには
// [os.Root] を使用してください。
func Sub(fsys FS, dir string) (FS, error)

var _ FS = (*subFS)(nil)
var _ ReadDirFS = (*subFS)(nil)
var _ ReadFileFS = (*subFS)(nil)
var _ ReadLinkFS = (*subFS)(nil)
var _ GlobFS = (*subFS)(nil)
