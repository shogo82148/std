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
// フラズテストでは、デフォルトでシードコーパスが実行されます。これには(*F).Addで提供されたエントリや、testdata/fuzz/<FuzzTestName>ディレクトリのエントリが含まれています。
// 必要なセットアップと(*F).Addへの呼び出しが行われた後、フラズテストは(*F).Fuzzを呼び出してフラズターゲットを提供する必要があります。
// 例についてはtestingパッケージのドキュメントを参照し、詳細についてはF.FuzzおよびF.Addメソッドのドキュメントを参照してください。
//
// *Fのメソッドは(*F).Fuzzの前にのみ呼び出すことができます。テストがフラズターゲットを実行している間は、(*T)のメソッドのみを使用することができます。
// (*F).Fuzz関数内で許可されている*Fのメソッドは、(*F).Failedと(*F).Nameのみです。
type F struct {
	common
	fuzzContext *fuzzContext
	testContext *testContext

	// inFuzzFn は、fuzz 関数が実行中である場合に true です。inFuzzFn が true の場合、ほとんどの F メソッドは呼び出すことができません。
	inFuzzFn bool

	// corpusはシードコーパスのエントリのセットで、F.Addで追加され、testdataから読み込まれます。
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
// ffは、戻り値を持たない関数でなければならず、最初の引数は*T型であり、残りの引数はfuzzテストを実施する型です。
// 例：
//
//  f.Fuzz(func(t *testing.T, b []byte, i int) { ... })
//
// 以下の型が許可されます：[]byte, string, bool, byte, rune, float32, float64, int, int8, int16, int32, int64,
// uint, uint8, uint16, uint32, uint64。将来的にはより多くの型がサポートされるかもしれません。
//
// ffは、(*F).Log、(*F).Error、(*F).Skipなどの*Fメソッドを呼び出してはなりません。代わりに、対応する*Tメソッドを使用してください。
// (*F).Fuzz関数で許可される*Fメソッドは、(*F).Failedと(*F).Nameのみです。
//
// この関数は高速かつ決定論的であるべきであり、その動作は共有状態に依存してはいけません。実行のたびに、実行その他の間に保持している
// 可変の入力引数またはそれらのポインタは、後続の呼び出し中に変更される可能性があるため、保持してはいけません。
// ffは、fuzzingエンジンによって提供された引数の基になるデータを変更してはいけません。
//
// fuzzing中、F.Fuzzは問題が見つかるまで、時間切れ（-fuzztimeで設定）またはテストプロセスがシグナルによって中断されるまで、戻りません。
// F.Fuzzは、F.SkipまたはF.Failが先に呼び出されない限り、正確に1回呼び出す必要があります。
func (f *F) Fuzz(ff any)
