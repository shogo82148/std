// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージパスは、スラッシュで区切られたパスを操作するためのユーティリティルーチンを実装します。
//
// pathパッケージは、URLのパスなど、スラッシュで区切られたパスにのみ使用する必要があります。
// このパッケージは、ドライブレターやバックスラッシュを含むWindowsパスに対処しません。
// オペレーティングシステムのパスを操作するには、 [path/filepath] パッケージを使用してください。
package path

// Clean関数は、パスを純粋に字句処理して、最短のパス名に変換します。
// 以下のルールを反復処理して、処理できる限り適用します:
//
//  1. 連続するスラッシュを単一のスラッシュに置き換えます。
//  2. 各「.」パス名要素（現在のディレクトリ）を削除します。
//  3. 各「..」パス名要素（親ディレクトリ）とそれに先行する「..」以外の要素を削除します。
//  4. ルートパスを始める「..」要素を削除します：
//     つまり、パスの先頭にある「/..」を「/」に置き換えます。
//
// 返されるパスの末尾には、ルート「/」の場合にのみスラッシュがあります。
//
// この処理結果が空の文字列の場合、Clean関数は「.」という文字列を返します。
//
// 参考文献：Rob Pike, “Lexical File Names in Plan 9 or
// Getting Dot-Dot Right,”
// https://9p.io/sys/doc/lexnames.html
func Clean(path string) string

// Splitは最後のスラッシュの直後にパスを分割し、
// ディレクトリとファイル名のコンポーネントに分ける。
// パスにスラッシュがない場合、Splitは空のディレクトリと
// ファイルをpathに設定して返します。
// 返される値は、path = dir + fileという性質を持っています。
func Split(path string) (dir, file string)

// Joinは任意の数のパス要素をスラッシュで区切って1つのパスに結合します。空の要素は無視されます。結果はクリーンになります。ただし、引数リストが空であるか、その要素がすべて空である場合、Joinは空の文字列を返します。
func Join(elem ...string) string

// Extは、pathで使用されるファイル名の拡張子を返します。
// 拡張子は、pathの最後のスラッシュで区切られた要素の最後のドットから始まるサフィックスです。
// ドットが存在しない場合は空です。
func Ext(path string) string

// Baseはパスの最後の要素を返します。
// 最後の要素を抽出する前に、トレーリングスラッシュは削除されます。
// パスが空の場合、Baseは「.」を返します。
// パスがすべてのスラッシュで構成されている場合、Baseは「/」を返します。
func Base(path string) string

// IsAbsはパスが絶対パスかどうかを報告します。
func IsAbs(path string) bool

<<<<<<< HEAD
// Dirはパスの最後の要素以外のすべてを返します。通常はパスのディレクトリです。
// Splitを使用して最後の要素を削除した後、パスはクリーン化され、末尾のスラッシュは削除されます。
// パスが空の場合、Dirは "." を返します。
// パスがスラッシュだけで構成され、スラッシュ以外のバイトが続く場合、Dirは単一のスラッシュを返します。それ以外の場合、返されるパスはスラッシュで終わりません。
=======
// Dir returns all but the last element of path, typically the path's directory.
// After dropping the final element using [Split], the path is Cleaned and trailing
// slashes are removed.
// If the path is empty, Dir returns ".".
// If the path consists entirely of slashes followed by non-slash bytes, Dir
// returns a single slash. In any other case, the returned path does not end in a
// slash.
>>>>>>> upstream/master
func Dir(path string) string
