// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

// ReadBuildInfoは実行中のバイナリに埋め込まれたビルド情報を返します。
// この情報はモジュールサポートでビルドされたバイナリでのみ利用可能です。
func ReadBuildInfo() (info *BuildInfo, ok bool)

// BuildInfoはGoバイナリから読み取られるビルド情報を表します。
type BuildInfo struct {
	// GoVersionはバイナリをビルドしたGoツールチェーンのバージョンです
	// （例: "go1.19.2"）。
	GoVersion string `json:",omitempty"`

	// Pathはバイナリのメインパッケージのパッケージパスです
	// （例: "golang.org/x/tools/cmd/stringer"）。
	Path string `json:",omitempty"`

	// Mainはバイナリのメインパッケージを含むモジュールを記述します。
	Main Module `json:""`

	// Depsは、このバイナリのビルドに寄与したすべての依存モジュール（直接・間接両方）を記述します。
	Deps []*Module `json:",omitempty"`

	// Settingsはバイナリのビルドに使用されたビルド設定を記述します。
	Settings []BuildSetting `json:",omitempty"`
}

// Moduleはビルドに含まれる単一のモジュールを記述します。
type Module struct {
	Path    string  `json:",omitempty"`
	Version string  `json:",omitempty"`
	Sum     string  `json:",omitempty"`
	Replace *Module `json:",omitempty"`
}

// BuildSettingはビルドに影響を与える1つの設定を表すキーと値のペアです。
//
// 定義されたキーには以下のものがあります:
//
//   - -buildmode: 使用されたbuildmodeフラグ（通常は "exe"）
//   - -compiler: 使用されたコンパイラツールチェーンフラグ（通常は "gc"）
//   - CGO_ENABLED: 有効なCGO_ENABLED環境変数
//   - CGO_CFLAGS: 有効なCGO_CFLAGS環境変数
//   - CGO_CPPFLAGS: 有効なCGO_CPPFLAGS環境変数
//   - CGO_CXXFLAGS: 有効なCGO_CXXFLAGS環境変数
//   - CGO_LDFLAGS: 有効なCGO_LDFLAGS環境変数
//   - DefaultGODEBUG: 有効なGODEBUG設定
//   - GOARCH: アーキテクチャターゲット
//   - GOAMD64/GOARM/GO386/etc: GOARCHのアーキテクチャ機能レベル
//   - GOOS: オペレーティングシステムターゲット
//   - GOFIPS140: 固定されたFIPS 140-3モジュールバージョン（存在する場合）
//   - vcs: ビルドが実行されたソースツリーのバージョン管理システム
//   - vcs.revision: 現在のコミットまたはチェックアウトのリビジョン識別子
//   - vcs.time: vcs.revisionに関連付けられた変更日時（RFC3339形式）
//   - vcs.modified: ソースツリーにローカル変更があったかどうか（trueまたはfalse）
type BuildSetting struct {
	// KeyとValueはビルド設定を表します。
	// Keyには等号、スペース、タブ、改行を含めてはいけません。
	Key string `json:",omitempty"`
	// Valueには改行（'\n'）を含めてはいけません。
	Value string `json:",omitempty"`
}

// Stringは[BuildInfo]の文字列表現を返します。
func (bi *BuildInfo) String() string

// ParseBuildInfoは[*BuildInfo.String]が返す文字列を解析し、
// 元のBuildInfoを復元します。
// ただしGoVersionフィールドは設定されません。
// 通常プログラムはこの関数を呼び出すべきではなく、
// 代わりに[ReadBuildInfo]、[debug/buildinfo.ReadFile]、
// または [debug/buildinfo.Read] を呼び出してください。
func ParseBuildInfo(data string) (bi *BuildInfo, err error)
