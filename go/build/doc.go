// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージ build はGoのパッケージに関する情報を収集します。
//
// # Go Path
//
// Goパスは、Goのソースコードが含まれているディレクトリツリーのリストです。
// これは、標準のGoツリーで見つからないインポートを解決するために参照されます。
// デフォルトのパスは、GOPATH環境変数の値であり、オペレーティングシステムに適したパスリストとして解釈されます
// (Unixでは変数はコロンで区切られた文字列であり、
// Windowsではセミコロンで区切られた文字列、
// Plan 9ではリストです)。
//
// Goパスにリストされている各ディレクトリには、指定された構造が必要です：
//
// src/ディレクトリにはソースコードが格納されます。'src'以下のパスがインポートパスまたは実行可能ファイル名を決定します。
//
// pkg/ディレクトリにはインストールされたパッケージオブジェクトが格納されます。
// Goツリーと同様に、各ターゲットオペレーティングシステムと
// アーキテクチャのペアに対して、pkgのサブディレクトリがあります
// (pkg/GOOS_GOARCH)。
//
// DIRがGoパスにリストされているディレクトリである場合、DIR/src/foo/barにソースがあるパッケージは「foo/bar」としてインポートされ、
// 「DIR/pkg/GOOS_GOARCH/foo/bar.a」(またはgccgoの場合は「DIR/pkg/gccgo/foo/libbar.a」)にコンパイルされた形式でインストールされます。
//
// bin/ディレクトリにはコンパイルされたコマンドが格納されます。
// 各コマンドは、ソースディレクトリを使用して命名されますが、
// パス全体ではなく最終要素のみを使用します。つまり、
// DIR/src/foo/quuxのソースであるコマンドはDIR/bin/quuxにインストールされます。DIR/bin/foo/quuxではなく、foo/が取り除かれます。
// そのため、PATHにDIR/binを追加することで、インストールされたコマンドにアクセスできます。
//
// 以下にディレクトリレイアウトの例を示します：
//
// GOPATH=/home/user/gocode
//
// /home/user/gocode/
//     src/
//         foo/
//             bar/               (パッケージ bar の Goコード)
//                 x.go
//             quux/              (メインパッケージの Goコード)
//                 y.go
//     bin/
//         quux                   (インストールされたコマンド)
//     pkg/
//         linux_amd64/
//             foo/
//                 bar.a          (インストールされたパッケージオブジェクト)
//
// # ビルド制約
//
// ビルド制約、またはビルドタグとも呼ばれるものは、
// パッケージに含めるべきファイルの条件です。ビルド制約は、次の行コメントによって与えられます。
//
//	//go:build
//
// ビルド制約は、ファイル名の一部にもなることがあります
// (例えば、source_windows.go は、対象の
// オペレーティングシステムがwindowsの場合のみ含まれます)。
//
// 詳細は 'go help buildconstraint'
// (https://golang.org/cmd/go/#hdr-Build_constraints) を参照してください。
//
// # バイナリのみのパッケージ
//
// Go 1.12 およびそれ以前では、
// パッケージをソースコードなしでバイナリ形式で配布することが可能でした。
// パッケージには、ビルド制約で除外されないソースファイルと
// 「//go:binary-only-package」というコメントが含まれていました。ビルド制約と同様に、このコメントはファイルの先頭に配置され、空行と他の行コメントだけが前にあること、コメントの後には空行があることで、パッケージのドキュメンテーションと区別されます。
// ビルド制約とは異なり、このコメントは非テストのGoソースファイルでのみ認識されます。
//
// バイナリのみのパッケージの最小のソースコードは以下のようになります：
//
//	//go:binary-only-package
//
//	package mypkg
//
// ソースコードには追加のGoコードが含まれる可能性があります。このコードはコンパイルされないが、godocなどのツールによって処理され、エンドユーザのドキュメンテーションとして役立つかもしれません。
//
<<<<<<< HEAD
// "go build" およびその他のコマンドはバイナリオンリーパッケージをサポートしていません。
// Import と ImportDir は、これらのコメントを含むパッケージのBinaryOnlyフラグを設定し、ツールやエラーメッセージで利用することができます。
=======
// "go build" and other commands no longer support binary-only-packages.
// [Import] and [ImportDir] will still set the BinaryOnly flag in packages
// containing these comments for use in tools and error messages.
>>>>>>> upstream/master
package build
