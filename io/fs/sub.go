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
// dirが"."の場合、Subはfsysをそのまま返します。
// それ以外で、fsが [SubFS] を実装している場合、SubはfsysのSub(dir)を返します。
// それ以外の場合、Subは新しい [FS] 実装subを返します。これは実質的に
// sub.Open(name)をfsys.Open(path.Join(dir, name))として実装します。
// この実装はReadDir、ReadFile、ReadLink、Lstat、およびGlobの呼び出しも
// 適切に変換します。Subはディレクトリが現在存在するかどうかを確認しません。
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
