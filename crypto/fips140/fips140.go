// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// fips140パッケージは、FIPS 140-3 Go暗号モジュールおよびFIPS 140-3モードに関する
// 情報を提供します。
//
// 詳細については、[FIPS 140-3 documentation]を参照してください。
//
// [FIPS 140-3 documentation]: https://go.dev/doc/security/fips140
package fips140

// Enabledは、暗号ライブラリがFIPS 140-3モードで動作しているかどうかを報告します。
//
// これは実行時にGODEBUG設定"fips140"を使用して制御できます。
// "on"に設定すると、FIPS 140-3モードが有効になります。"only"に設定すると、
// 承認されていない暗号関数はさらにエラーを返すかパニックします。
//
// これはプログラムの開始後に変更することはできません。
func Enabled() bool

// Versionは、GOFIPS140を使用して凍結されたモジュールに対してビルドしている場合、
// モジュールのセキュリティポリシーで参照されているFIPS 140-3 Go暗号モジュールの
// バージョン（"v1.0.0"など）を返します。それ以外の場合は"latest"を返します。
// エイリアス（"inprogress"など）が使用されている場合は、実際に解決された
// バージョンが返されます。
//
// 同じバージョンの凍結モジュールのコピーが複数ある場合、返されるバージョンは
// プログラムのビルドに使用された凍結モジュールを一意に識別しない可能性があります。
// 一意に識別するためのバージョンサフィックスは、
// runtime/debug.BuildInfo.SettingsのGOFIPS140設定の値を確認することで見つけられます。
func Version() string
