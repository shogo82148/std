// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

// InternalFuzzTargetは内部型ですが、異なるパッケージで使われるために公開されています。
// これは"go test"コマンドの実装の一部です。
type InternalFuzzTarget struct {
	Name string
	Fn   func(f *F)
}

// Fはフラズテストに渡される型です。
//
// フラズテストは生成された入力値を提供されたフラズターゲットに対して実行し、
// テストされるコードに潜在的なバグを見つけて報告することができます。
//
// ファズテストはデフォルトでシードコーパスを実行します。これには [F.Add] で追加されたエントリや testdata/fuzz/<FuzzTestName> ディレクトリ内のエントリが含まれます。
// 必要なセットアップと [F.Add] の呼び出しの後、ファズテストは [F.Fuzz] を呼び出してファズターゲットを指定する必要があります。
// 詳細は testing パッケージのドキュメントや [F.Fuzz] および [F.Add] メソッドのドキュメントを参照してください。
//
// *F メソッドは [F.Fuzz] の前にのみ呼び出すことができます。テストがファズターゲットを実行している間は [*T] メソッドのみ使用できます。
// [F.Fuzz] 関数内で許可される *F メソッドは [F.Failed] と [F.Name] のみです。
type F struct {
	common
	fstate *fuzzState
	tstate *testState

	// inFuzzFn is true when the fuzz function is running. Most F methods cannot
	// be called when inFuzzFn is true.
	inFuzzFn bool

	// corpus is a set of seed corpus entries, added with F.Add and loaded
	// from testdata.
	corpus []corpusEntry

	result     fuzzResult
	fuzzCalled bool
}

var _ TB = (*F)(nil)

// Helperは呼び出し元の関数をテストのヘルパー関数としてマークします。
// ファイルと行情報を表示するとき、その関数はスキップされます。
// Helperは複数のゴルーチンから同時に呼び出すことができます。
func (f *F) Helper()

// Failは関数が失敗したことを示しますが、実行を続けます。
func (f *F) Fail()

// Skippedはテストがスキップされたかどうかを報告します。
func (f *F) Skipped() bool

// Addは、引数をfuzzテストのシードコーパスに追加します。これは、fuzzターゲットの後または中で呼び出された場合は無効になり、argsはfuzzターゲットの引数と一致する必要があります。
func (f *F) Add(args ...any)

// Fuzzはfuzzテストのために、関数ffを実行します。もしffが一連の引数で失敗した場合、それらの引数はシードコーパスに追加されます。
//
// ffは戻り値がなく、最初の引数が[*T]であり、残りの引数がファズ対象の型である関数でなければなりません。
// 例:
//
//	f.Fuzz(func(t *testing.T, b []byte, i int) { ... })
//
// 以下の型が許可されます：[]byte, string, bool, byte, rune, float32, float64, int, int8, int16, int32, int64,
// uint, uint8, uint16, uint32, uint64。将来的にはより多くの型がサポートされるかもしれません。
//
// ffは[*F]のメソッド（例: [F.Log], [F.Error], [F.Skip]）を呼び出してはいけません。代わりに対応する[*T]メソッドを使用してください。
// F.Fuzz関数内で許可される [*F] メソッドは [F.Failed] と [F.Name] のみです。
//
// この関数は高速で決定的であるべきであり、その振る舞いは共有状態に依存してはなりません。
// 可変の入力引数、またはそれらへのポインターは、fuzz関数の実行間で保持されるべきではありません。
// なぜなら、それらをバックアップするメモリは、後続の呼び出し中に変更される可能性があるからです。
// ffは、fuzzingエンジンによって提供される引数の基礎となるデータを変更してはなりません。
//
// ファズ実行中は、F.Fuzzは問題が見つかるか、タイムアウト（-fuzztimeで設定）、またはテストプロセスがシグナルで中断されるまで戻りません。
// F.Fuzzは [F.Skip] または [F.Fail] が事前に呼ばれない限り、必ず一度だけ呼び出す必要があります。
func (f *F) Fuzz(ff any)
