// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package doc

// Dirは、予想されるインポートパスとファイルシステムディレクトリーを指定して、コードを保持するディレクトリーを記述します。
type Dir struct {
	importPath string
	dir        string
	inModule   bool
}

// Dirsはディレクトリーツリーをスキャンするための構造体です。
// そのNextメソッドは次に見つかったGoソースディレクトリーを返します。
// ツリーを複数回スキャンするために使用できますが、
// ツリーを一度だけ走査して、見つけたデータをキャッシュします。
type Dirs struct {
	scan   chan Dir
	hist   []Dir
	offset int
}

// Resetはスキャンを最初に戻します。
func (d *Dirs) Reset()

// Nextはスキャン中の次のディレクトリーを返します。ブール値がfalseの場合、スキャンが完了しています。
func (d *Dirs) Next() (Dir, bool)
