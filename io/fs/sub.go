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
// そうでなく、fsysが [SubFS] を実装している場合、Subはfsys.Sub(dir)を返します。
// それ以外の場合、Subは新しい [FS] の実装subを返します。subは実質的に、
// sub.Open(name)をfsys.Open(path.Join(dir, name))として実装します。
// この実装は、ReadDir、ReadFile、ReadLink、Lstat、Globの呼び出しも適切に変換します。
// Subは、ディレクトリーが現在存在するかどうかをチェックしません。
//
// Sub(os.DirFS("/"), "prefix")はos.DirFS("/prefix")と同等ですが、
// どちらも"/prefix"の外でのオペレーティングシステムアクセスを避けることを
// 保証しません。これは [os.DirFS] の実装が"/prefix"内で他のディレクトリーを
// 指すシンボリックリンクをチェックしないためです。つまり、[os.DirFS] は
// chroot方式のセキュリティーメカニズムの一般的な代替ではなく、Subはその
// 事実を変えません。特定のディレクトリーツリーへのアクセスを制限するには
// [os.Root] を使用してください。
func Sub(fsys FS, dir string) (FS, error)

var _ FS = (*subFS)(nil)
var _ ReadDirFS = (*subFS)(nil)
var _ ReadFileFS = (*subFS)(nil)
var _ ReadLinkFS = (*subFS)(nil)
var _ GlobFS = (*subFS)(nil)
