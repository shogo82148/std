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
	// （例: "go1.19.2").
	GoVersion string

	// Pathはバイナリのメインパッケージのパッケージパスです
	// （例：「golang.org/x/tools/cmd/stringer」）。
	Path string

	// Mainはバイナリのmainパッケージを含むモジュールを説明します。
	Main Module

	// Depsは、このバイナリのビルドに寄与したパッケージの直接および間接の依存モジュールをすべて説明します。
	Deps []*Module

	// Settingsはバイナリのビルドに使用されるビルド設定を記述しています。
	Settings []BuildSetting
}

// Moduleはビルドに含まれる単一のモジュールを記述します。
type Module struct {
	Path    string
	Version string
	Sum     string
	Replace *Module
}

// BuildSettingはビルドに影響を与える1つの設定を表すキーと値のペアです。
//
// 定義されたキーには以下のものがあります:
//
<<<<<<< HEAD
//   - -buildmode: the buildmode flag used (typically "exe")
//   - -compiler: the compiler toolchain flag used (typically "gc")
//   - CGO_ENABLED: the effective CGO_ENABLED environment variable
//   - CGO_CFLAGS: the effective CGO_CFLAGS environment variable
//   - CGO_CPPFLAGS: the effective CGO_CPPFLAGS environment variable
//   - CGO_CXXFLAGS:  the effective CGO_CXXFLAGS environment variable
//   - CGO_LDFLAGS: the effective CGO_LDFLAGS environment variable
//   - GOARCH: the architecture target
//   - GOAMD64/GOARM/GO386/etc: the architecture feature level for GOARCH
//   - GOOS: the operating system target
//   - vcs: the version control system for the source tree where the build ran
//   - vcs.revision: the revision identifier for the current commit or checkout
//   - vcs.time: the modification time associated with vcs.revision, in RFC3339 format
//   - vcs.modified: true or false indicating whether the source tree had local modifications
=======
//   - -buildmode: 使用されたビルドモードフラグ（通常は "exe"）
//   - -compiler: 使用されたコンパイラツールチェインフラグ（通常は "gc"）
//   - CGO_ENABLED: 有効なCGO_ENABLED環境変数
//   - CGO_CFLAGS: 有効なCGO_CFLAGS環境変数
//   - CGO_CPPFLAGS: 有効なCGO_CPPFLAGS環境変数
//   - CGO_CXXFLAGS: 有効なCGO_CPPFLAGS環境変数
//   - CGO_LDFLAGS: 有効なCGO_CPPFLAGS環境変数
//   - GOARCH: アーキテクチャのターゲット
//   - GOAMD64/GOARM/GO386/など: GOARCHのアーキテクチャ機能レベル
//   - GOOS: オペレーティングシステムのターゲット
//   - vcs: ビルドが実行されたソースツリーのバージョン管理システム
//   - vcs.revision: 現在のコミットまたはチェックアウトのためのリビジョン識別子
//   - vcs.time: vcs.revisionに関連付けられた変更時刻（RFC3339形式）
//   - vcs.modified: ソースツリーにローカルの変更があるかどうかを示すtrueまたはfalseです
>>>>>>> release-branch.go1.21
type BuildSetting struct {

	// KeyとValueはビルド設定を説明します。
	// Keyには等号、スペース、タブ、改行を含めることはできません。
	// Valueには改行（'\n'）を含めることはできません。
	Key, Value string
}

func (bi *BuildInfo) String() string

func ParseBuildInfo(data string) (bi *BuildInfo, err error)
