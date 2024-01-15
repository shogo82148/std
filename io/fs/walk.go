// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"github.com/shogo82148/std/errors"
)

<<<<<<< HEAD
// SkipDirはWalkDirFuncsからの戻り値として使用され、呼び出しで指定されたディレクトリがスキップされることを示します。これは、どの関数からもエラーとして返されません。
var SkipDir = errors.New("skip this directory")

// SkipAllは、WalkDirFuncsからの返り値として使用され、残りのすべてのファイルとディレクトリをスキップすることを示します。これは、どの関数からもエラーとして返されません。
=======
// SkipDir is used as a return value from [WalkDirFunc] to indicate that
// the directory named in the call is to be skipped. It is not returned
// as an error by any function.
var SkipDir = errors.New("skip this directory")

// SkipAll is used as a return value from [WalkDirFunc] to indicate that
// all remaining files and directories are to be skipped. It is not returned
// as an error by any function.
>>>>>>> upstream/master
var SkipAll = errors.New("skip everything and stop the walk")

// WalkDirFuncは [WalkDir] によって各ファイルやディレクトリを訪れるために呼び出される関数の型です。
//
// path引数には、 [WalkDir] の引数としてのパスが前置されます。
// つまり、root引数が "dir" でWalkDirがそのディレクトリで "a" という名前のファイルを見つけた場合、
// 引数が "dir/a" であるように歩行関数が呼び出されます。
//
// d引数は、指定されたパスの [fs.DirEntry] です。
//
// 関数によって返されるエラー結果は、 [WalkDir] の進行方法を制御します。
// 関数が特別な値 [SkipDir] を返す場合、WalkDirは現在のディレクトリをスキップします（d.IsDir()がtrueであればパス、
// そうでなければパスの親ディレクトリ）。
// 関数が特別な値 [SkipAll] を返す場合、WalkDirは残りのすべてのファイルおよびディレクトリをスキップします。
// それ以外の場合、関数が非nilのエラーを返す場合、WalkDirは完全に停止し、そのエラーを返します。
//
// エラー引数は、パスに関連するエラーを報告し、WalkDirがそのディレクトリに入ろうとしないことを示します。
// 関数はそのエラーを処理する方法を決定することができます。
// エラーを返すと、WalkDirは木全体のツリーをかけるのをやめます。
//
// [WalkDir] は、2つのケースで非nilのerr引数を持って関数を呼び出します。
//
// まず、ルートディレクトリの初期 [Stat] が失敗した場合、WalkDirは関数をpathがrootに設定され、
// dがnilに設定され、errが [fs.Stat] からのエラーに設定された状態で呼び出します。
//
// 2番目に、ディレクトリのReadDirメソッド（ [ReadDirFile] を参照）が失敗した場合、WalkDirは関数をディレクトリのパスがpathに設定され、
// dがディレクトリを記述する [DirEntry] に設定され、errがReadDirからのエラーに設定された状態で呼び出します。
// この2番目の場合、関数はディレクトリのパスで2回呼び出されます。
// 最初の呼び出しは、ディレクトリの読み取りが試みられる前で、errがnilに設定されるため、関数に [SkipDir] または [SkipAll] を返すチャンスがあり、ReadDirを完全に回避します。
// 2回目の呼び出しは、失敗したReadDirからのエラーを報告します。
// （ReadDirが成功すると、2回目の呼び出しがありません。）
//
// WalkDirFuncと [path/filepath.WalkFunc] の違いは次のとおりです：
//
//   - 2番目の引数の型が [FileInfo] ではなく [DirEntry] であること。
//   - ディレクトリを読み取る前に関数が呼び出され、 [SkipDir] または [SkipAll] がディレクトリの読み取りを完全にバイパスしたり、
//     残りのすべてのファイルとディレクトリをスキップしたりするようにすること。
//   - ディレクトリの読み取りが失敗した場合、そのディレクトリについてのエラーを報告するために、関数が2回呼び出されること。
type WalkDirFunc func(path string, d DirEntry, err error) error

// WalkDirはルートにルートされたファイルツリーを走査し、各ファイルまたはディレクトリに対してfnを呼び出します。
//
// ファイルとディレクトリを訪れる中で発生するエラーは、fnによってフィルタリングされます：
// 詳細については、[fs.WalkDirFunc] のドキュメントを参照してください。
//
// ファイルは辞書式順に走査されますが、出力を決定論的にするために、WalkDirはディレクトリ全体をメモリに読み込んでから、そのディレクトリを走査する必要があります。
//
// WalkDirはディレクトリ内で見つかったシンボリックリンクをたどりませんが、ルート自体がシンボリックリンクであれば、その対象が走査されます。
func WalkDir(fsys FS, root string, fn WalkDirFunc) error
