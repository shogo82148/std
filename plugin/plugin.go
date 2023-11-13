// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージプラグインは、Goプラグインの読み込みとシンボルの解決を実装します。
//
// プラグインは、エクスポートされた関数と変数を持つGo mainパッケージであり、
// 次のようにビルドされたものです：
//
// go build -buildmode=plugin
//
// プラグインが最初に開かれると、既にプログラムの一部ではないすべてのパッケージのinit関数が呼び出されます。main関数は実行されません。
// プラグインは一度だけ初期化され、クローズすることはできません。
//
// # 警告
//
// 実行中にアプリケーションの一部を動的にロードする機能は、
// ユーザー定義の設定に基づいている場合などに、一部の設計において有用なビルディングブロックとなる場合があります。
// 特に、アプリケーションと動的にロードされる関数がデータ構造を直接共有できるため、プラグインは異なるパーツの非常に高性能な統合を可能にする場合があります。
//
// ただし、プラグインの仕組みには多くの重要な制約があるため、設計時に慎重に考慮する必要があります。例えば：
//
//   - プラグインは現在Linux、FreeBSD、macOSのみをサポートしており、
//     ポータブルなアプリケーションには適していません。
//
//   - プラグインを使用するアプリケーションは、
//     プログラムのさまざまなパーツが正しい場所にファイルシステム（またはコンテナイメージ）に配置されるように細心の注意を払う必要があります。
//     その対照に、単一の静的実行可能ファイルからなるアプリケーションのデプロイは簡単です。
//
//   - パッケージの初期化についての推論は、
//     アプリケーションの実行が開始されてから長い時間が経過するまで、
//     一部のパッケージが初期化されない場合に困難になります。
//
//   - プラグインを読み込むアプリケーションのバグは、
//     攻撃者によって危険または信頼できないライブラリを読み込ませるために悪用される可能性があります。
//
//   - すべてのプログラムの部分（アプリケーションとそのプラグインすべて）が、
//     正確に同じバージョンのツールチェイン、同じビルドタグ、および特定のフラグと環境変数の値を使用してコンパイルされなければ、
//     ランタイムクラッシュが発生する可能性が高いです。
//
//   - アプリケーションとそのプラグインの共通の依存関係が、
//     正確に同じソースコードからビルドされない限り、同様のクラッシュ問題が発生する可能性があります。
//
//   - これらの制約のため、実際的には、
//     アプリケーションとそのプラグインはすべて、システムの単一のコンポーネントまたは一人の担当者によって一緒にビルドされる必要があります。
//     その場合、それらの人またはコンポーネントが所望のプラグインのセットをブランクインポートするGoのソースファイルを生成し、
//     通常の方法で静的な実行可能ファイルをコンパイルする方が簡単かもしれません。
//
// これらの理由から、性能のオーバーヘッドにもかかわらず、
// 多くのユーザーは、ソケット、パイプ、リモートプロシージャコール（RPC）、共有メモリマッピング、
// ファイルシステム操作などの従来のプロセス間通信（IPC）メカニズムの方がより適していると判断します。
package plugin

// PluginはロードされたGoプラグインです。
type Plugin struct {
	pluginpath string
	err        string
	loaded     chan struct{}
	syms       map[string]any
}

<<<<<<< HEAD
// OpenはGoプラグインを開きます。
// もし既にパスが開かれている場合は、既存の*Pluginが返されます。
// 複数のゴルーチンによる同時使用も安全です。
=======
// Open opens a Go plugin.
// If a path has already been opened, then the existing *[Plugin] is returned.
// It is safe for concurrent use by multiple goroutines.
>>>>>>> upstream/master
func Open(path string) (*Plugin, error)

// Lookupは、プラグインpの中でsymNameという名前のシンボルを検索します。
// シンボルは、エクスポートされた変数または関数のことです。
// シンボルが見つからない場合はエラーを報告します。
// 複数のゴルーチンによる同時利用でも安全です。
func (p *Plugin) Lookup(symName string) (Symbol, error)

// シンボルは変数や関数へのポインタです。
//
// 例えば、以下のように定義されたプラグインがあります。
//
//	package main
//
//	import "fmt"
//
//	var V int
//
//	func F() { fmt.Printf("Hello, number %d\n", V) }
//
<<<<<<< HEAD
// このプラグインはOpen関数を用いて読み込むことができ、その後エクスポートされたパッケージの
// シンボルVとFにアクセスすることができます。
=======
// may be loaded with the [Open] function and then the exported package
// symbols V and F can be accessed
>>>>>>> upstream/master
//
//	p, err := plugin.Open("plugin_name.so")
//	if err != nil {
//		panic(err)
//	}
//	v, err := p.Lookup("V")
//	if err != nil {
//		panic(err)
//	}
//	f, err := p.Lookup("F")
//	if err != nil {
//		panic(err)
//	}
//	*v.(*int) = 7
//	f.(func())() // "Hello, number 7"と表示されます
type Symbol any
