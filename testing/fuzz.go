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
<<<<<<< HEAD
// ファズテストはデフォルトでシードコーパスを実行します。これには、(*F).Addによって提供されたエントリと
// testdata/fuzz/<FuzzTestName>ディレクトリ内のエントリが含まれます。必要なセットアップと(*F).Addへの呼び出しの後、
// ファズテストはファズターゲットを提供するために(*F).Fuzzを呼び出さなければなりません。
// 例についてはテストパッケージのドキュメンテーションを参照し、詳細については [F.Fuzz] と [F.Add] メソッドのドキュメンテーションを参照してください。
//
// *Fのメソッドは(*F).Fuzzの前にのみ呼び出すことができます。テストがフラズターゲットを実行している間は、(*T)のメソッドのみを使用することができます。
// (*F).Fuzz関数内で許可されている*Fのメソッドは、(*F).Failedと(*F).Nameのみです。
=======
// A fuzz test runs the seed corpus by default, which includes entries provided
// by [F.Add] and entries in the testdata/fuzz/<FuzzTestName> directory. After
// any necessary setup and calls to [F.Add], the fuzz test must then call
// [F.Fuzz] to provide the fuzz target. See the testing package documentation
// for an example, and see the [F.Fuzz] and [F.Add] method documentation for
// details.
//
// *F methods can only be called before [F.Fuzz]. Once the test is
// executing the fuzz target, only [*T] methods can be used. The only *F methods
// that are allowed in the [F.Fuzz] function are [F.Failed] and [F.Name].
>>>>>>> upstream/release-branch.go1.25
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
<<<<<<< HEAD
// ffは、戻り値を持たない関数でなければならず、最初の引数は*T型であり、残りの引数はfuzzテストを実施する型です。
// 例：
=======
// ff must be a function with no return value whose first argument is [*T] and
// whose remaining arguments are the types to be fuzzed.
// For example:
>>>>>>> upstream/release-branch.go1.25
//
//	f.Fuzz(func(t *testing.T, b []byte, i int) { ... })
//
// 以下の型が許可されます：[]byte, string, bool, byte, rune, float32, float64, int, int8, int16, int32, int64,
// uint, uint8, uint16, uint32, uint64。将来的にはより多くの型がサポートされるかもしれません。
//
<<<<<<< HEAD
// ffは、(*F).Log、(*F).Error、(*F).Skipなどの*Fメソッドを呼び出してはなりません。代わりに、対応する*Tメソッドを使用してください。
// (*F).Fuzz関数で許可される*Fメソッドは、(*F).Failedと(*F).Nameのみです。
=======
// ff must not call any [*F] methods, e.g. [F.Log], [F.Error], [F.Skip]. Use
// the corresponding [*T] method instead. The only [*F] methods that are allowed in
// the F.Fuzz function are [F.Failed] and [F.Name].
>>>>>>> upstream/release-branch.go1.25
//
// この関数は高速で決定的であるべきであり、その振る舞いは共有状態に依存してはなりません。
// 可変の入力引数、またはそれらへのポインターは、fuzz関数の実行間で保持されるべきではありません。
// なぜなら、それらをバックアップするメモリは、後続の呼び出し中に変更される可能性があるからです。
// ffは、fuzzingエンジンによって提供される引数の基礎となるデータを変更してはなりません。
//
<<<<<<< HEAD
// fuzzing中、F.Fuzzは問題が見つかるまで、時間切れ（-fuzztimeで設定）またはテストプロセスがシグナルによって中断されるまで、戻りません。
// F.Fuzzは、F.Skipまたは [F.Fail] が先に呼び出されない限り、正確に1回呼び出す必要があります。
=======
// When fuzzing, F.Fuzz does not return until a problem is found, time runs out
// (set with -fuzztime), or the test process is interrupted by a signal. F.Fuzz
// should be called exactly once, unless [F.Skip] or [F.Fail] is called beforehand.
>>>>>>> upstream/release-branch.go1.25
func (f *F) Fuzz(ff any)
