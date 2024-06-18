// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package filepathは、ファイル名パスを操作するためのユーティリティ関数を実装しています。
// これは、対象のオペレーティングシステムで定義されたファイルパスと互換性のある方法で行います。
//
// filepathパッケージは、オペレーティングシステムに応じてスラッシュまたはバックスラッシュを使用します。
// オペレーティングシステムに関係なく常にスラッシュを使用するURLのようなパスを処理するには、[path]パッケージを参照してください。
package filepath

import (
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/os"
)

const (
	Separator     = os.PathSeparator
	ListSeparator = os.PathListSeparator
)

// Cleanは、純粋な字句解析によってパスに相当する最短のパス名を返します。
// 次の規則を適用し、処理できなくなるまで反復的に行います。
//
//  1. 複数の [Separator] 要素を単一の要素に置き換える。
//  2. 各「.」パス名要素（カレントディレクトリ）を削除する。
//  3. 各内部の「..」パス名要素（親ディレクトリ）を削除し、それに続く非「..」要素も削除する。
//  4. ルートパスで始まる「..」要素を削除する：
//     つまり、パスの先頭で "/.." を "/" に置き換える（セパレータが '/' であると仮定）。
//
// 返されるパスは、ルートディレクトリを表す場合のみスラッシュで終わります。
// 例：Unixでは "/"、Windowsでは `C:\` などです。
//
// 最後に、スラッシュの出現箇所はセパレータに置き換えられます。
//
// この処理の結果が空の文字列の場合、Cleanは文字列 "." を返します。
//
// Windowsでは、Cleanは"/"を `\` に置き換える以外では、ボリューム名を変更しません。
// 例えば、Clean("//host/share/../x") は `\\host\share\x` を返します。
//
// 参考資料：Rob Pike, "Lexical File Names in Plan 9 or Getting Dot-Dot Right"
// https://9p.io/sys/doc/lexnames.html
func Clean(path string) string

// IsLocalは、パスが次の条件をすべて満たすかどうかを報告します。
// - パスが評価されるディレクトリをルートとするサブツリー内にある
// - 絶対パスではない
// - 空ではない
// - Windowsの場合、"NUL"のような予約済みの名前ではない
// IsLocal(path)がtrueを返す場合、
// Join(base, path)は常にbase内に含まれるパスを生成し、
// Clean(path)は常に".."パス要素を持たないルート付きのパスを生成します。
// IsLocalは純粋に文字解析の操作です。
// 特に、ファイルシステムに存在する可能性のあるシンボリックリンクの影響は考慮されません。
func IsLocal(path string) bool

// Localizeは、スラッシュ区切りのパスをオペレーティングシステムのパスに変換します。
// 入力パスは、[io/fs.ValidPath] によって報告される有効なパスでなければなりません。
//
// Localizeは、パスがオペレーティングシステムによって表現できない場合にエラーを返します。
// 例えば、Windowsではパスa\bは拒否されます。これは、\がセパレータ文字であり、
// ファイル名の一部にはなり得ないからです。
//
// Localizeによって返されるパスは、IsLocalによって報告されるように、常にローカルになります。
func Localize(path string) (string, error)

// ToSlashは、パス内の各区切り文字をスラッシュ('/')文字で置き換えた結果を返します。複数の区切り文字は複数のスラッシュに置き換えられます。
func ToSlash(path string) string

// FromSlashは、パス内の各スラッシュ('/')文字をセパレータ文字に置き換えた結果を返します。
// 複数のスラッシュは複数のセパレータに置き換えられます。
//
// io/fsパッケージで使用されるスラッシュ区切りのパスをオペレーティングシステムのパスに変換する
// Localize関数も参照してください。
func FromSlash(path string) string

// SplitListは、OS固有の [ListSeparator] で結合されたパスのリストを分割します。
// 通常、 PATHまたはGOPATH環境変数で見つかることがあります。
// strings.Splitとは異なり、SplitListは空のスライスを返します。
// 空の文字列が渡された場合。
func SplitList(path string) []string

// Splitは最後の [Separator] の直後にあるパスを分割し、ディレクトリとファイル名の要素に分けます。
// パスにSeparatorがない場合、Splitは空のディレクトリとファイルを返します。
// 返される値は path = dir + file という性質を持ちます。
func Split(path string) (dir, file string)

// Joinは、パスの要素をいくつでも指定して、OS固有の [Separator] で区切って1つのパスに結合します。空の要素は無視されます。結果はクリーニングされます。ただし、引数リストが空であるか、そのすべての要素が空の場合、Joinは空の文字列を返します。
// Windowsでは、最初の非空要素がUNCパスである場合にのみ結果はUNCパスになります。
func Join(elem ...string) string

// Extはpathで使用されるファイル名の拡張子を返します。
// 拡張子とは、pathの最後の要素の最後のドットから始まる接尾辞であり、
// ドットがない場合は空です。
func Ext(path string) string

// EvalSymlinksは、シンボリックリンクの評価後のパス名を返します。
// pathが相対パスの場合、結果は現在のディレクトリを基準にし、
// 絶対シンボリックリンクを持つコンポーネントが存在する場合を除きます。
// EvalSymlinksは結果に対して [Clean] を呼び出します。
func EvalSymlinks(path string) (string, error)

// IsAbsはパスが絶対パスかどうかを報告します。
func IsAbs(path string) bool

// Absはパスの絶対表現を返します。
// パスが絶対でない場合、現在の作業ディレクトリと結合して絶対パスに変換されます。
// 特定のファイルの絶対パス名が一意であることは保証されません。
// Absは結果に [Clean] を呼び出します。
func Abs(path string) (string, error)

// Relは、中間のセパレーターでbasepathと結合したときにtargpathと同等の相対パスを返します。
// つまり、[Join](basepath, Rel(basepath, targpath))はtargpathと同じです。
// 成功した場合、返されるパスは常にbasepathに対して相対的であり、
// basepathとtargpathが要素を共有していなくても同じです。
// targpathがbasepathに相対化できない場合や、現在の作業ディレクトリの情報が必要な場合はエラーが返されます。
// Relは結果に対して [Clean] を呼び出します。
func Rel(basepath, targpath string) (string, error)

// SkipDirは、[WalkFunc] からの返り値として使用され、呼び出し元で指定されたディレクトリをスキップすることを示します。これは、どの関数からもエラーとして返されません。
var SkipDir error = fs.SkipDir

// SkipAllは [WalkFunc] からの戻り値として使用され、残りのすべてのファイルとディレクトリをスキップすることを示します。これはエラーではなく、いかなる関数からも戻されません。
var SkipAll error = fs.SkipAll

// WalkFuncは、[Walk] によって呼び出される関数の型です。この関数は、各ファイルやディレクトリを訪れるために呼び出されます。
//
// path引数には、Walkの引数がプレフィックスとして含まれています。
// つまり、root引数が「dir」として呼び出され、そのディレクトリに「a」という名前のファイルが見つかった場合、
// ウォーク関数は引数「dir/a」で呼び出されます。
//
// ディレクトリとファイルはJoinで結合され、ディレクトリ名がクリーンアップされるかもしれません。
// たとえば、root引数が「x/../dir」として呼び出され、そのディレクトリに「a」という名前のファイルが見つかった場合、
// ウォーク関数は引数「dir/a」で呼び出されます。「x/../dir/a」とはなりません。
//
// info引数は、指定されたパスのfs.FileInfoです。
//
// 関数が返すエラー結果によって、Walkの継続が制御されます。
// 関数が特殊値 [SkipDir] を返すと、Walkは現在のディレクトリ（info.IsDirがtrueの場合はpath、そうでない場合はpathの親ディレクトリ）をスキップします。
// 関数が特殊値 [SkipAll] を返すと、Walkは残りの全てのファイルとディレクトリをスキップします。
// さもなくば、関数が非nilのエラーを返すと、Walkは完全に停止し、そのエラーを返します。
//
// err引数は、pathに関連するエラーを報告し、Walkがそのディレクトリに進まないことを示します。
// 関数はそのエラーをどのように処理するかを決定できます。先述のように、エラーを返すと、
// Walkは木全体を走査するのを停止します。
//
// Walkは、2つの場合に、非nilのerr引数を持つ関数を呼び出します。
//
// 第1に、ルートディレクトリまたはツリー内の任意のディレクトリまたはファイルの [os.Lstat] が失敗した場合、
// Walkは関数を呼び出し、パスをそのディレクトリまたはファイルのパスに設定し、infoをnilに設定し、errをos.Lstatからのエラーに設定します。
//
// 第2に、ディレクトリのReaddirnamesメソッドが失敗した場合、
// Walkは関数を呼び出し、パスをディレクトリのパスに設定し、infoをディレクトリを説明する [fs.FileInfo] に設定し、errをReaddirnamesからのエラーに設定します。
type WalkFunc func(path string, info fs.FileInfo, err error) error

// WalkDirは、ルートにあるファイルツリーを走査し、各ファイルやディレクトリに対してfnを呼び出します。
// ツリー内のルートも含まれます。
//
// ファイルやディレクトリの訪問中に発生するすべてのエラーは、fnによってフィルタリングされます：
// 詳細については、[fs.WalkDirFunc] のドキュメントを参照してください。
//
// ファイルは辞書順で走査されるため、出力が確定論的になりますが、WalkDirは
// そのディレクトリの走査に進む前にディレクトリ全体をメモリに読み込む必要があります。
//
// WalkDirはシンボリックリンクを辿りません。
//
// WalkDirは、オペレーティングシステムに適切な区切り文字を使用するパスを
// fnに渡して呼び出します。これは、[io/fs.WalkDir]とは異なり、常にスラッシュで区切られたパスを使用します。
func WalkDir(root string, fn fs.WalkDirFunc) error

// Walkはルートとなるファイルツリーを辿り、各ファイルまたはディレクトリに対してfnを呼び出します。
// これにはルートも含まれます。
//
// ファイルとディレクトリの訪問時に発生するエラーは、すべてfnによってフィルタリングされます。
// 詳細については [WalkFunc] のドキュメントを参照してください。
//
// ファイルはレキシカルオーダーで走査されますが、これにより出力は決定論的になります。
// ただし、走査するディレクトリの前にディレクトリ全体をメモリに読み込む必要があります。
//
// Walkはシンボリックリンクを辿りません。
//
// WalkはGo 1.16で導入された [WalkDir] よりも効率が低下します。
// WalkDirでは、訪問するファイルまたはディレクトリごとにos.Lstatを呼び出すのを避けています。
func Walk(root string, fn WalkFunc) error

// Baseはパスの最後の要素を返します。
// パスの末尾のセパレーターは、最後の要素を抽出する前に削除されます。
// パスが空の場合、Baseは「.」を返します。
// パスが完全にセパレーターで構成されている場合、Baseは単一のセパレーターを返します。
func Base(path string) string

// Dirはパスの最後の要素以外のすべてを返します。通常はパスのディレクトリです。
// 最後の要素を除いた後、Dirはパスに [Clean] を呼び出し、末尾のスラッシュは除去されます。
// パスが空の場合、Dirは「.」を返します。
// パスがセパレーターだけで構成されている場合、Dirは単一のセパレーターを返します。
// 返されるパスは、ルートディレクトリでない限り、セパレーターで終了しません。
func Dir(path string) string

// VolumeNameは先頭のボリューム名を返します。
// Windowsの場合、"C:\foo\bar"に対しては"C:"を返します。
// "\\host\share\foo"に対しては"\\host\share"を返します。
// 他のプラットフォームでは、空文字列を返します。
func VolumeName(path string) string
