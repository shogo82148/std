// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"github.com/shogo82148/std/errors"
)

// SkipDirはWalkDirFuncsからの戻り値として使用され、呼び出しで指定されたディレクトリがスキップされることを示します。これは、どの関数からもエラーとして返されません。
var SkipDir = errors.New("skip this directory")

// SkipAllは、WalkDirFuncsからの返り値として使用され、残りのすべてのファイルとディレクトリをスキップすることを示します。これは、どの関数からもエラーとして返されません。
var SkipAll = errors.New("skip everything and stop the walk")

<<<<<<< HEAD
// WalkDirFunc is the type of the function called by [WalkDir] to visit
// each file or directory.
//
// The path argument contains the argument to [WalkDir] as a prefix.
// That is, if WalkDir is called with root argument "dir" and finds a file
// named "a" in that directory, the walk function will be called with
// argument "dir/a".
//
// The d argument is the [DirEntry] for the named path.
//
// The error result returned by the function controls how [WalkDir]
// continues. If the function returns the special value [SkipDir], WalkDir
// skips the current directory (path if d.IsDir() is true, otherwise
// path's parent directory). If the function returns the special value
// [SkipAll], WalkDir skips all remaining files and directories. Otherwise,
// if the function returns a non-nil error, WalkDir stops entirely and
// returns that error.
//
// The err argument reports an error related to path, signaling that
// [WalkDir] will not walk into that directory. The function can decide how
// to handle that error; as described earlier, returning the error will
// cause WalkDir to stop walking the entire tree.
//
// [WalkDir] calls the function with a non-nil err argument in two cases.
//
// First, if the initial [Stat] on the root directory fails, WalkDir
// calls the function with path set to root, d set to nil, and err set to
// the error from fs.Stat.
//
// Second, if a directory's ReadDir method (see [ReadDirFile]) fails, WalkDir calls the
// function with path set to the directory's path, d set to an
// [DirEntry] describing the directory, and err set to the error from
// ReadDir. In this second case, the function is called twice with the
// path of the directory: the first call is before the directory read is
// attempted and has err set to nil, giving the function a chance to
// return [SkipDir] or [SkipAll] and avoid the ReadDir entirely. The second call
// is after a failed ReadDir and reports the error from ReadDir.
// (If ReadDir succeeds, there is no second call.)
//
// The differences between WalkDirFunc compared to [path/filepath.WalkFunc] are:
//
//   - The second argument has type [DirEntry] instead of [FileInfo].
//   - The function is called before reading a directory, to allow [SkipDir]
//     or [SkipAll] to bypass the directory read entirely or skip all remaining
//     files and directories respectively.
//   - If a directory read fails, the function is called a second time
//     for that directory to report the error.
=======
// WalkDirFuncはWalkDirによって各ファイルやディレクトリを訪れるために呼び出される関数の型です。
//
// path引数には、WalkDirの引数としてのパスが前置されます。
// つまり、root引数が "dir" でWalkDirがそのディレクトリで "a" という名前のファイルを見つけた場合、
// 引数が "dir/a" であるように歩行関数が呼び出されます。
//
// d引数は、指定されたパスのfs.DirEntryです。
//
// 関数によって返されるエラー結果は、WalkDirの進行方法を制御します。
// 関数が特別な値SkipDirを返す場合、WalkDirは現在のディレクトリをスキップします（d.IsDir()がtrueであればパス、
// そうでなければパスの親ディレクトリ）。
// 関数が特別な値SkipAllを返す場合、WalkDirは残りのすべてのファイルおよびディレクトリをスキップします。
// それ以外の場合、関数が非nilのエラーを返す場合、WalkDirは完全に停止し、そのエラーを返します。
//
// エラー引数は、パスに関連するエラーを報告し、WalkDirがそのディレクトリに入ろうとしないことを示します。
// 関数はそのエラーを処理する方法を決定することができます。
// エラーを返すと、WalkDirは木全体のツリーをかけるのをやめます。
//
// WalkDirは、2つのケースで非nilのerr引数を持って関数を呼び出します。
//
// まず、ルートディレクトリの初期fs.Statが失敗した場合、WalkDirは関数をpathがrootに設定され、
// dがnilに設定され、errがfs.Statからのエラーに設定された状態で呼び出します。
//
// 2番目に、ディレクトリのReadDirメソッドが失敗した場合、WalkDirは関数をディレクトリのパスがpathに設定され、
// dがディレクトリを記述するfs.DirEntryに設定され、errがReadDirからのエラーに設定された状態で呼び出します。
// この2番目の場合、関数はディレクトリのパスで2回呼び出されます。
// 最初の呼び出しは、ディレクトリの読み取りが試みられる前で、errがnilに設定されるため、関数にSkipDirまたはSkipAllを返すチャンスがあり、ReadDirを完全に回避します。
// 2回目の呼び出しは、失敗したReadDirからのエラーを報告します。
// （ReadDirが成功すると、2回目の呼び出しがありません。）
//
// WalkDirFuncとfilepath.WalkFuncの違いは次のとおりです：
//
//   - 2番目の引数の型がfs.DirEntryであること。
//   - ディレクトリを読み取る前に関数が呼び出され、SkipDirまたはSkipAllがディレクトリの読み取りを完全にバイパスしたり、
//     残りのすべてのファイルとディレクトリをスキップしたりするようにすること。
//   - ディレクトリの読み取りが失敗した場合、そのディレクトリについてのエラーを報告するために、関数が2回呼び出されること。
>>>>>>> release-branch.go1.21
type WalkDirFunc func(path string, d DirEntry, err error) error

// WalkDirはルートにルートされたファイルツリーを走査し、各ファイルまたはディレクトリに対してfnを呼び出します。
//
// ファイルとディレクトリを訪れる中で発生するエラーは、fnによってフィルタリングされます：
// 詳細については、fs.WalkDirFuncのドキュメントを参照してください。
//
// ファイルは辞書式順に走査されますが、出力を決定論的にするために、WalkDirはディレクトリ全体をメモリに読み込んでから、そのディレクトリを走査する必要があります。
//
// WalkDirはディレクトリ内で見つかったシンボリックリンクをたどりませんが、ルート自体がシンボリックリンクであれば、その対象が走査されます。
func WalkDir(fsys FS, root string, fn WalkDirFunc) error
