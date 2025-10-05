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
// dirが"."の場合、Subはfsysを変更せずに返します。
// そうでなければ、fsが [SubFS] を実装している場合、Subはfsys.Sub(dir)を返します。
// そうでなければ、Subは新しい [FS] 実装subを返します。これは
// 実質的にsub.Open(name)をfsys.Open(path.Join(dir, name))として実装します。
// この実装はまた、ReadDir、ReadFile、
// ReadLink、Lstat、およびGlobの呼び出しも適切に変換します。
//
// os.DirFS("/prefix")とSub(os.DirFS("/"), "prefix")は同等であることに注意してください。
// どちらも、"/prefix"の外部にあるオペレーティングシステムのアクセスを回避することを保証するものではありません。
// なぜなら、 [os.DirFS] の実装は、"/prefix"内部の他のディレクトリを指すシンボリックリンクをチェックしないためです。
// つまり、os.DirFSはchrootスタイルのセキュリティメカニズムの一般的な代替手段ではなく、Subもその事実を変えません。
func Sub(fsys FS, dir string) (FS, error)

var _ FS = (*subFS)(nil)
var _ ReadDirFS = (*subFS)(nil)
var _ ReadFileFS = (*subFS)(nil)
var _ ReadLinkFS = (*subFS)(nil)
var _ GlobFS = (*subFS)(nil)
