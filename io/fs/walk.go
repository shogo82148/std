// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

import (
	"github.com/shogo82148/std/errors"
)

// SkipDirは、[WalkDirFunc] からの戻り値として使用され、
// 呼び出しで名前が付けられたディレクトリーをスキップすることを示します。
// これはどの関数からもエラーとして返されません。
var SkipDir = errors.New("skip this directory")

// SkipAllは、[WalkDirFunc] からの戻り値として使用され、
// 残りのすべてのファイルとディレクトリーをスキップすることを示します。
// これはどの関数からもエラーとして返されません。
var SkipAll = errors.New("skip everything and stop the walk")

// WalkDirFuncは [WalkDir] によって各ファイルやディレクトリーを訪れるために呼び出される関数の型です。
//
// path引数には、 [WalkDir] の引数としてのパスが前置されます。
// つまり、root引数が "dir" でWalkDirがそのディレクトリーで "a" という名前のファイルを見つけた場合、
// 引数が "dir/a" であるように歩行関数が呼び出されます。
//
// d引数は、指定されたパスの [fs.DirEntry] です。
//
// 関数によって返されるエラー結果は、 [WalkDir] の進行方法を制御します。
// 関数が特別な値 [SkipDir] を返す場合、WalkDirは現在のディレクトリーをスキップします（d.IsDir()がtrueであればパス、
// そうでなければパスの親ディレクトリー）。
// 関数が特別な値 [SkipAll] を返す場合、WalkDirは残りのすべてのファイルおよびディレクトリーをスキップします。
// それ以外の場合、関数が非nilのエラーを返す場合、WalkDirは完全に停止し、そのエラーを返します。
//
// エラー引数は、パスに関連するエラーを報告し、WalkDirがそのディレクトリーに入ろうとしないことを示します。
// 関数はそのエラーを処理する方法を決定することができます。
// エラーを返すと、WalkDirは木全体のツリーをかけるのをやめます。
//
// [WalkDir] は、2つのケースで非nilのerr引数を持って関数を呼び出します。
//
// まず、ルートディレクトリーの初期 [Stat] が失敗した場合、WalkDirは関数をpathがrootに設定され、
// dがnilに設定され、errが [fs.Stat] からのエラーに設定された状態で呼び出します。
//
// 2番目に、ディレクトリーのReadDirメソッド（ [ReadDirFile] を参照）が失敗した場合、WalkDirは関数をディレクトリーのパスがpathに設定され、
// dがディレクトリーを記述する [DirEntry] に設定され、errがReadDirからのエラーに設定された状態で呼び出します。
// この2番目の場合、関数はディレクトリーのパスで2回呼び出されます。
// 最初の呼び出しは、ディレクトリーの読み取りが試みられる前で、errがnilに設定されるため、関数に [SkipDir] または [SkipAll] を返すチャンスがあり、ReadDirを完全に回避します。
// 2回目の呼び出しは、失敗したReadDirからのエラーを報告します。
// （ReadDirが成功すると、2回目の呼び出しがありません。）
//
// WalkDirFuncと [path/filepath.WalkFunc] の違いは次のとおりです：
//
//   - 2番目の引数の型が [FileInfo] ではなく [DirEntry] であること。
//   - ディレクトリーを読み取る前に関数が呼び出され、 [SkipDir] または [SkipAll] がディレクトリーの読み取りを完全にバイパスしたり、
//     残りのすべてのファイルとディレクトリーをスキップしたりするようにすること。
//   - ディレクトリーの読み取りが失敗した場合、そのディレクトリーについてのエラーを報告するために、関数が2回呼び出されること。
type WalkDirFunc func(path string, d DirEntry, err error) error

// WalkDirはルートにルートされたファイルツリーを走査し、各ファイルまたはディレクトリーに対してfnを呼び出します。
//
// ファイルとディレクトリーを訪れる中で発生するエラーは、fnによってフィルタリングされます：
// 詳細については、[fs.WalkDirFunc] のドキュメントを参照してください。
//
// ファイルは辞書式順に走査されますが、出力を決定論的にするために、WalkDirはディレクトリー全体をメモリーに読み込んでから、そのディレクトリーを走査する必要があります。
//
// WalkDirはディレクトリー内で見つかったシンボリックリンクをたどりませんが、ルート自体がシンボリックリンクであれば、その対象が走査されます。
func WalkDir(fsys FS, root string, fn WalkDirFunc) error
