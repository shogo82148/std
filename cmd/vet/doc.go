// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
VetはGoのソースコードを検査し、Printfのような呼び出しがフォーマット文字列と一致しない場合に疑わしい構造を報告します。Vetは完全な報告を保証するわけではないヒューリスティックを使用しているため、すべての報告が本当の問題ではないかもしれませんが、コンパイラでは見つからないエラーを見つけることができます。

Vetは通常、goコマンドを通じて起動されます。
現在のディレクトリのパッケージを検査するためには、以下のコマンドを使用します:

	go vet

パスが指定されたパッケージを検査するためには、以下のコマンドを使用します:

	go vet my/project/...

パッケージを指定する他の方法については「go help packages」を参照してください。

Vetの終了コードは、ツールの誤った呼び出しまたは問題が報告された場合に非ゼロであり、それ以外の場合は0です。ツールはすべての可能な問題をチェックせず、信頼性の低いヒューリスティックに依存しているため、プログラムの正確性の厳密な指標ではなく、ガイダンスとして使用する必要があります。

利用可能なチェックをリストするには、「go tool vet help」と実行します:

<<<<<<< HEAD
	appends          append後の値の欠落をチェックします
	asmdecl          アセンブリファイルとGo宣言の間の不一致を報告します
	assign           無意味な代入をチェックします
	atomic           sync/atomicパッケージを使用する際の一般的なミスをチェックします
	bools            論理演算子に関する一般的なミスをチェックします
	buildtag         //go:buildと// +buildディレクティブをチェックします
	cgocall          cgoポインタ受け渡しルールの一部の違反を検出します
	composites       キーのない複合リテラルをチェックします
	copylocks        値によって誤って渡されたロックをチェックします
	defers           defer文の一般的なミスを報告します
	directive        //go:debugなどのGoツールチェインディレクティブをチェックします
	errorsas         errors.Asに非ポインタまたは非エラー値を渡すことを報告します
	framepointer     フレームポインタを保存する前に破壊するアセンブリを報告します
	httpresponse     HTTPレスポンスを使用する際のミスをチェックします
	ifaceassert      不可能なインターフェース間型アサーションを検出します
	loopclosure      ネストした関数内からのループ変数への参照をチェックします
	lostcancel       context.WithCancelから返されるcancel関数が呼び出されているかをチェックします
	nilfunc          関数とnilの間の無意味な比較をチェックします
	printf           Printfフォーマット文字列と引数の一貫性をチェックします
	shift            整数の幅と等しいかそれを超えるシフトをチェックします
	sigchanyzer      os.Signalのバッファなしチャネルをチェックします
	slog             無効な構造化ログ呼び出しをチェックします
	stdmethods       よく知られたインターフェースのメソッドのシグネチャをチェックします
	stringintconv    string(int)変換をチェックします
	structtag        構造体フィールドタグがreflect.StructTag.Getに準拠しているかをチェックします
	testinggoroutine テストによって開始されたゴルーチンからの(*testing.T).Fatal呼び出しを報告します
	tests            テストと例の一般的な誤用をチェックします
	timeformat       2006-02-01を使用した(time.Time).Formatまたはtime.Parseの呼び出しをチェックします
	unmarshal        unmarshalに非ポインタまたは非インターフェース値を渡すことを報告します
	unreachable      到達不可能なコードをチェックします
	unsafeptr        uintptrからunsafe.Pointerへの無効な変換をチェックします
	unusedresult     一部の関数の呼び出しの未使用結果をチェックします
	waitgroup        sync.WaitGroupの誤用をチェックします
=======
	appends          check for missing values after append
	asmdecl          report mismatches between assembly files and Go declarations
	assign           check for useless assignments
	atomic           check for common mistakes using the sync/atomic package
	bools            check for common mistakes involving boolean operators
	buildtag         check //go:build and // +build directives
	cgocall          detect some violations of the cgo pointer passing rules
	composites       check for unkeyed composite literals
	copylocks        check for locks erroneously passed by value
	defers           report common mistakes in defer statements
	directive        check Go toolchain directives such as //go:debug
	errorsas         report passing non-pointer or non-error values to errors.As
	framepointer     report assembly that clobbers the frame pointer before saving it
	hostport         check format of addresses passed to net.Dial
	httpresponse     check for mistakes using HTTP responses
	ifaceassert      detect impossible interface-to-interface type assertions
	loopclosure      check references to loop variables from within nested functions
	lostcancel       check cancel func returned by context.WithCancel is called
	nilfunc          check for useless comparisons between functions and nil
	printf           check consistency of Printf format strings and arguments
	shift            check for shifts that equal or exceed the width of the integer
	sigchanyzer      check for unbuffered channel of os.Signal
	slog             check for invalid structured logging calls
	stdmethods       check signature of methods of well-known interfaces
	stdversion       report uses of too-new standard library symbols
	stringintconv    check for string(int) conversions
	structtag        check that struct field tags conform to reflect.StructTag.Get
	testinggoroutine report calls to (*testing.T).Fatal from goroutines started by a test
	tests            check for common mistaken usages of tests and examples
	timeformat       check for calls of (time.Time).Format or time.Parse with 2006-02-01
	unmarshal        report passing non-pointer or non-interface values to unmarshal
	unreachable      check for unreachable code
	unsafeptr        check for invalid conversions of uintptr to unsafe.Pointer
	unusedresult     check for unused results of calls to some functions
	waitgroup        check for misuses of sync.WaitGroup
>>>>>>> upstream/release-branch.go1.26

printfなどの特定のチェックの詳細とフラグについての情報は、「go tool vet help printf」と実行してください。

デフォルトでは、すべてのチェックが実行されます。
フラグがtrueに明示的に設定されている場合、それらのテストのみが実行されます。
逆に、フラグが明示的にfalseに設定されている場合、それらのテストは無効になります。
したがって、-printf=trueはprintfチェックを実行し、
-printf=falseはprintfチェック以外のすべてのチェックを実行します。

新しいチェックの作成方法については、golang.org/x/tools/go/analysisを参照してください。

コアフラグ:

	-c=N
	  	エラーのある行とその周囲のN行を表示する
	-json
	  	分析診断（およびエラー）をJSON形式で出力する
*/
package main
