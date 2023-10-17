// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// SubFSはSubメソッドを持つファイルシステムです。
type SubFS interface {
	FS

	// Sub returns an FS corresponding to the subtree rooted at dir.
	Sub(dir string) (FS, error)
}

// Subはfsysのdirでルートされた部分木に対応する [FS] を返します。
//
// dirが"."の場合、Subはfsysを変更せずに返します。
// そうでない場合、fsが [SubFS] を実装している場合、Subはfsys.Sub(dir)を返します。
// そうでない場合、Subは効果的に、sub.Open(name)をfsys.Open(path.Join(dir, name))として実装する新しい [FS] 実装subを返します。
// 実装では、ReadDir、ReadFile、そしてGlobへの呼び出しも適切に変換されます。
//
// os.DirFS("/prefix")とSub(os.DirFS("/"), "prefix")は同等であることに注意してください。
// どちらも、"/prefix"の外部にあるオペレーティングシステムのアクセスを回避することを保証するものではありません。
// なぜなら、 [os.DirFS] の実装は、"/prefix"内部の他のディレクトリを指すシンボリックリンクをチェックしないためです。
// つまり、os.DirFSはchrootスタイルのセキュリティメカニズムの一般的な代替手段ではなく、Subもその事実を変えません。
func Sub(fsys FS, dir string) (FS, error)
