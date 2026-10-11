// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

// LookPathは、ホストオペレーティングシステムの規約に従って、
// 現在のパスでfileという名前の実行可能ファイルを検索します。
// fileにスラッシュが含まれている場合は、直接試行され、デフォルトのパスは参照されません。
// それ以外の場合、成功すると結果は絶対パスになります。
//
// 解決されたパスがカレントディレクトリからの相対パスである場合、LookPathは
// [errors.Is](err, [ErrDot])を満たすエラーを返します。
// 詳細はパッケージのドキュメントを参照してください。
//
// LookPathは、以下で説明する場合を除き、PATH環境変数で指定された
// ディレクトリ内でfileという名前の実行可能ファイルを探します。
//
//   - Windowsでは、ファイルはPATHEXT環境変数で指定された拡張子を持たなければなりません。
//     PATHEXTが設定されていない場合、ファイルは
//     ".com"、".exe"、".bat"、".cmd"のいずれかの拡張子を持たなければなりません。
//   - Plan 9では、LookPathはpath環境変数を参照します。
//     fileが"/"、"#"、"./"、"../"で始まる場合は、直接試行され、
//     pathは参照されません。
//   - Wasmでは、LookPathは常にエラーを返します。
func LookPath(file string) (string, error)
