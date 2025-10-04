// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test2jsonは、go testの出力を機械可読のJSONストリームに変換します。
//
// 使用方法:
//
//	go tool test2json [-p pkg] [-t] [./pkg.test -test.v=test2json]
//
// Test2jsonは、指定されたテストコマンドを実行し、その出力をJSONに変換します。
// コマンドが指定されていない場合、test2jsonは標準入力からテストの出力を予期します。
// 対応するJSONイベントのストリームを標準出力に書き込みます。
// 入出力の不要なバッファリングは行われないため、テストの状態の "ライブ更新" のために
// JSONストリームを読み取ることができます。
//
// -pフラグは、各テストイベントで報告されるパッケージを設定します。
//
// -tフラグは、各テストイベントにタイムスタンプを追加するようリクエストします。
//
// テストは、-test.v=test2jsonで呼び出す必要があります。-test.vのみを使用することもできます
// （または -test.v=true）、しかし、より低い信頼性の結果となります。
//
// "go test -json"コマンドはtest2jsonを正しく呼び出すことに対応しているため、
// "go tool test2json"は、テストバイナリが"go test"とは別に実行される場合にのみ必要です。
// 可能な限り"go test -json"を使用してください。
//
// また、test2jsonは単一のテストバイナリの出力を変換するためのものであることに注意してください。
// 複数のパッケージを実行する"go test"コマンドの出力を変換するには、再び"go test -json"を使用してください。
//
// # 出力フォーマット
//
// JSONストリームは、改行で区切られたTestEventオブジェクトのシーケンスで、
// Goの構造体に対応します：
//
//	type TestEvent struct {
<<<<<<< HEAD
//		Time    time.Time // RFC3339形式の文字列としてエンコードされます
//		Action  string
//		Package string
//		Test    string
//		Elapsed float64 // 秒単位
//		Output  string
=======
//		Time        time.Time // encodes as an RFC3339-format string
//		Action      string
//		Package     string
//		Test        string
//		Elapsed     float64 // seconds
//		Output      string
//		FailedBuild string
>>>>>>> upstream/release-branch.go1.25
//	}
//
// Timeフィールドはイベントが発生した時刻を保持しています。
// キャッシュされたテスト結果には、通常は省略されます。
//
// Actionフィールドは、固定のアクションの説明の1つです：
//
//	start  - テストバイナリが実行される直前
//	run    - テストが実行開始される
//	pause  - テストが一時停止される
//	cont   - テストが実行再開される
//	pass   - テストにパスする
//	bench  - ベンチマークがログ出力を行うが、失敗はしない
//	fail   - テストまたはベンチマークが失敗する
//	output - テストが出力を行う
//	skip   - テストがスキップされるか、パッケージにテストが含まれていない
//
// JSONストリームは常に "start" イベントで始まります。
//
// Packageフィールドが存在する場合、テストされているパッケージを指定します。
// goコマンドが- jsonモードで並列テストを実行する場合、異なるテストのイベントが交互に現れます。
// Packageフィールドにより、読み手はそれらを区別できます。
//
// Testフィールドが存在する場合、イベントを引き起こしたテスト、例、またはベンチマーク関数を指定します。
// パッケージ全体のテストの場合、Testは設定されません。
//
// Elapsedフィールドは、"pass"と"fail"のイベントに設定されます。
// パスまたは失敗した特定のテストまたはパッケージ全体のテストの経過時間を示します。
//
// OutputフィールドはAction == "output"の場合に設定され、テストの出力の一部です
// （標準出力と標準エラーを結合したもの）。出力は変更されず、テストからの無効なUTF-8出力は、
// 置換文字を使用して有効なUTF-8に変換されます。この例外を除いて、
// Outputフィールドのすべての出力イベントの連結がテストの実行の正確な出力です。
//
<<<<<<< HEAD
// ベンチマークが実行されると、通常はタイミング結果を示す1行の出力が生成されます。
// その行は、Action == "output"かつTestフィールドが存在しないイベントで報告されます。
// ベンチマークが出力を記録したり失敗を報告した場合
// （たとえば、b.Logやb.Errorを使用することによって）、その追加の出力は
// ベンチマーク名が設定されたイベントのシーケンスとして報告され、最後のイベントは
// Action == "bench"または"fail"です。ベンチマークにはAction == "pause"のイベントはありません。
=======
// The FailedBuild field is set for Action == "fail" if the test failure was
// caused by a build failure. It contains the package ID of the package that
// failed to build. This matches the ImportPath field of the "go list" output,
// as well as the BuildEvent.ImportPath field as emitted by "go build -json".
//
// When a benchmark runs, it typically produces a single line of output
// giving timing results. That line is reported in an event with Action == "output"
// and no Test field. If a benchmark logs output or reports a failure
// (for example, by using b.Log or b.Error), that extra output is reported
// as a sequence of events with Test set to the benchmark name, terminated
// by a final event with Action == "bench" or "fail".
// Benchmarks have no events with Action == "pause".
>>>>>>> upstream/release-branch.go1.25
package main
