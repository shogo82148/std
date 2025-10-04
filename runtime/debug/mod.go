// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug

// ReadBuildInfoは実行中のバイナリに埋め込まれたビルド情報を返します。
// この情報はモジュールサポートでビルドされたバイナリでのみ利用可能です。
func ReadBuildInfo() (info *BuildInfo, ok bool)

// BuildInfoはGoバイナリから読み取られるビルド情報を表します。
type BuildInfo struct {
<<<<<<< HEAD

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
=======
	// GoVersion is the version of the Go toolchain that built the binary
	// (for example, "go1.19.2").
	GoVersion string `json:",omitempty"`

	// Path is the package path of the main package for the binary
	// (for example, "golang.org/x/tools/cmd/stringer").
	Path string `json:",omitempty"`

	// Main describes the module that contains the main package for the binary.
	Main Module `json:""`

	// Deps describes all the dependency modules, both direct and indirect,
	// that contributed packages to the build of this binary.
	Deps []*Module `json:",omitempty"`

	// Settings describes the build settings used to build the binary.
	Settings []BuildSetting `json:",omitempty"`
>>>>>>> upstream/release-branch.go1.25
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
<<<<<<< HEAD
//   - -buildmode: 使用されたビルドモードフラグ（通常は "exe"）
//   - -compiler: 使用されたコンパイラツールチェーンフラグ（通常は "gc"）
//   - CGO_ENABLED: 有効なCGO_ENABLED環境変数
//   - CGO_CFLAGS: 有効なCGO_CFLAGS環境変数
//   - CGO_CPPFLAGS: 有効なCGO_CPPFLAGS環境変数
//   - CGO_CXXFLAGS: 有効なCGO_CXXFLAGS環境変数
//   - CGO_LDFLAGS: 有効なCGO_LDFLAGS環境変数
//   - GOARCH: アーキテクチャターゲット
//   - GOAMD64/GOARM/GO386/etc: GOARCHのアーキテクチャ機能レベル
//   - GOOS: オペレーティングシステムターゲット
//   - vcs: ビルドが実行されたソースツリーのバージョン管理システム
//   - vcs.revision: 現在のコミットまたはチェックアウトのリビジョン識別子
//   - vcs.time: vcs.revisionに関連付けられた修正時刻（RFC3339形式）
//   - vcs.modified: ローカルの変更があるかどうかを示すtrueまたはfalse
type BuildSetting struct {

	// KeyとValueはビルド設定を説明します。
	// Keyには等号、スペース、タブ、改行を含めることはできません。
	// Valueには改行（'\n'）を含めることはできません。
	Key, Value string
=======
//   - -buildmode: the buildmode flag used (typically "exe")
//   - -compiler: the compiler toolchain flag used (typically "gc")
//   - CGO_ENABLED: the effective CGO_ENABLED environment variable
//   - CGO_CFLAGS: the effective CGO_CFLAGS environment variable
//   - CGO_CPPFLAGS: the effective CGO_CPPFLAGS environment variable
//   - CGO_CXXFLAGS:  the effective CGO_CXXFLAGS environment variable
//   - CGO_LDFLAGS: the effective CGO_LDFLAGS environment variable
//   - DefaultGODEBUG: the effective GODEBUG settings
//   - GOARCH: the architecture target
//   - GOAMD64/GOARM/GO386/etc: the architecture feature level for GOARCH
//   - GOOS: the operating system target
//   - GOFIPS140: the frozen FIPS 140-3 module version, if any
//   - vcs: the version control system for the source tree where the build ran
//   - vcs.revision: the revision identifier for the current commit or checkout
//   - vcs.time: the modification time associated with vcs.revision, in RFC3339 format
//   - vcs.modified: true or false indicating whether the source tree had local modifications
type BuildSetting struct {
	// Key and Value describe the build setting.
	// Key must not contain an equals sign, space, tab, or newline.
	Key string `json:",omitempty"`
	// Value must not contain newlines ('\n').
	Value string `json:",omitempty"`
>>>>>>> upstream/release-branch.go1.25
}

// String returns a string representation of a [BuildInfo].
func (bi *BuildInfo) String() string

// ParseBuildInfo parses the string returned by [*BuildInfo.String],
// restoring the original BuildInfo,
// except that the GoVersion field is not set.
// Programs should normally not call this function,
// but instead call [ReadBuildInfo], [debug/buildinfo.ReadFile],
// or [debug/buildinfo.Read].
func ParseBuildInfo(data string) (bi *BuildInfo, err error)
