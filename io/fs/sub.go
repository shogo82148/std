// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// SubFSはSubメソッドを持つファイルシステムです。
type SubFS interface {
	FS

	Sub(dir string) (FS, error)
}

<<<<<<< HEAD
// Sub returns an [FS] corresponding to the subtree rooted at fsys's dir.
//
// If dir is ".", Sub returns fsys unchanged.
// Otherwise, if fs implements [SubFS], Sub returns fsys.Sub(dir).
// Otherwise, Sub returns a new [FS] implementation sub that,
// in effect, implements sub.Open(name) as fsys.Open(path.Join(dir, name)).
// The implementation also translates calls to ReadDir, ReadFile, and Glob appropriately.
//
// Note that Sub(os.DirFS("/"), "prefix") is equivalent to os.DirFS("/prefix")
// and that neither of them guarantees to avoid operating system
// accesses outside "/prefix", because the implementation of [os.DirFS]
// does not check for symbolic links inside "/prefix" that point to
// other directories. That is, [os.DirFS] is not a general substitute for a
// chroot-style security mechanism, and Sub does not change that fact.
=======
// Subはfsysのdirでルートされた部分木に対応するFSを返します。
//
// dirが"."の場合、Subはfsysを変更せずに返します。
// そうでない場合、fsがSubFSを実装している場合、Subはfsys.Sub(dir)を返します。
// そうでない場合、Subは効果的に、sub.Open(name)をfsys.Open(path.Join(dir, name))として実装する新しいFS実装subを返します。
// 実装では、ReadDir、ReadFile、そしてGlobへの呼び出しも適切に変換されます。
//
// os.DirFS("/prefix")とSub(os.DirFS("/"), "prefix")は同等であることに注意してください。
// どちらも、"/prefix"の外部にあるオペレーティングシステムのアクセスを回避することを保証するものではありません。
// なぜなら、os.DirFSの実装は、"/prefix"内部の他のディレクトリを指すシンボリックリンクをチェックしないためです。
// つまり、os.DirFSはchrootスタイルのセキュリティメカニズムの一般的な代替手段ではなく、Subもその事実を変えません。
>>>>>>> release-branch.go1.21
func Sub(fsys FS, dir string) (FS, error)
