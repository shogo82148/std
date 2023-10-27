// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package slogは、メッセージ、重大度レベル、およびキー-値ペアとして表されるさまざまなその他の属性を含むログレコードを提供する構造化されたログを提供します。

[Logger] という型を定義し、
[Logger.Info] や [Logger.Error] などのいくつかのメソッドを提供して、
興味深いイベントを報告するための構造化されたログを提供します。

各Loggerは [Handler] に関連付けられています。
Loggerの出力メソッドは、メソッド引数から [Record] を作成し、
それを処理する方法を決定するHandlerに渡します。
対応するLoggerメソッドを呼び出す [Info] や [Error] などのトップレベル関数を介してアクセス可能なデフォルトのLoggerがあります。

ログレコードは、時刻、レベル、メッセージ、およびキー-値ペアのセットで構成されます。
キーは文字列で、値は任意の型である場合があります。
例として、

	slog.Info("hello", "count", 3)

呼び出しの時間、Infoレベル、メッセージ"hello"、および単一のキー"count"と値3を持つレコードを作成します。

[Info] トップレベル関数は、デフォルトのLogger上の [Logger.Info] メソッドを呼び出します。
[Logger.Info] に加えて、Debug、Warn、Errorレベルのメソッドがあります。
これらの一般的なレベルのための便利なメソッドに加えて、
[Logger.Log] メソッドがあり、レベルを引数として受け取ります。
これらのメソッドのそれぞれに対応するトップレベル関数があり、
デフォルトのロガーを使用します。

デフォルトのハンドラは、ログレコードのメッセージ、時刻、レベル、および属性を
文字列としてフォーマットし、 [log] パッケージに渡します。

	2022/11/08 15:28:26 INFO hello count=3

出力フォーマットをより細かく制御するには、別のハンドラを持つロガーを作成します。
このステートメントでは、 [New] を使用して、 [TextHandler] で構造化されたレコードをテキスト形式で標準エラーに書き込む新しいロガーを作成しています。

	logger := slog.New(slog.NewTextHandler(os.Stderr, nil))

[TextHandler] の出力は、キー=値のペアのシーケンスであり、機械によって簡単かつ曖昧に解析できます。
この文:

	logger.Info("hello", "count", 3)

は、次の出力を生成します。

	time=2022-11-08T15:28:26.000-05:00 level=INFO msg=hello count=3

パッケージはまた、行区切りJSONで出力される [JSONHandler] を提供します。

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello", "count", 3)

次の出力を生成します:

	{"time":"2022-11-08T15:28:26.000000000-05:00","level":"INFO","msg":"hello","count":3}

[TextHandler] と [JSONHandler] の両方は、 [HandlerOptions] で構成できます。
最小レベルの設定(以下の [Levels] を参照)、
ログ呼び出しのソースファイルと行の表示、
およびログに記録される前に属性を変更するためのオプションがあります。

デフォルトのロガーは以下のようにして変更できます。

	slog.SetDefault(logger)

[Info] のようなトップレベルの関数がloggerを使用するようになります。
[SetDefault] は、 [log] パッケージが使用するデフォルトのロガーも更新します。
これにより [log.Printf] などを使用する既存のアプリケーションが、
書き換える必要なくログレコードをロガーのハンドラに送信できます。

多くのログ呼び出しで共通の属性があります。
たとえば、サーバーリクエストから生じるすべてのログイベントにURLやトレース識別子を含めたい場合があります。
ログ呼び出しごとに属性を繰り返す代わりに、 [Logger.With] を使用して属性を含む新しいLoggerを構築できます。

	logger2 := logger.With("url", r.URL)

Withの引数は、 [Logger.Info] で使用されるキー-値ペアと同じです。
結果は、元のハンドラと同じハンドラを持つ新しいLoggerですが、
すべての呼び出しの出力に表示される追加の属性が含まれています。

# Levels

[Level] は、ログイベントの重要度または深刻度を表す整数です。
レベルが高いほど、イベントはより深刻です。
このパッケージは、最も一般的なレベルの定数を定義していますが、
任意のintをレベルとして使用できます。

アプリケーションでは、特定のレベル以上のメッセージのみをログに記録することが望ましい場合があります。
一般的な構成の1つは、Infoレベル以上のメッセージをログに記録し、
デバッグログを必要になるまで抑制することです。
組み込みのハンドラは、 [HandlerOptions.Level] を設定することで、
出力する最小レベルを構成できます。
通常、プログラムの`main`関数がこれを行います。
デフォルト値はLevelInfoです。

[HandlerOptions.Level] フィールドを [Level] 値に設定すると、
ハンドラの最小レベルがその寿命全体で固定されます。
[LevelVar] に設定すると、レベルを動的に変化させることができます。
LevelVarはLevelを保持し、複数のゴルーチンから読み書きすることができます。
プログラム全体でレベルを動的に変化させるには、まずグローバルなLevelVarを初期化します。

	var programLevel = new(slog.LevelVar) // Info by default

次に、LevelVarを使用してハンドラを構築し、デフォルトにします。

	h := slog.NewJSONHandler(os.Stderr, &slog.HandlerOptions{Level: programLevel})
	slog.SetDefault(slog.New(h))

プログラムは、単一のステートメントでログレベルを変更できるようになりました。

	programLevel.Set(slog.LevelDebug)

# Groups

属性はグループに集めることができます。
グループには、属性の名前に修飾子として使用される名前があります。
この修飾子がどのように表示されるかは、ハンドラによって異なります。
[TextHandler] は、グループと属性名をドットで区切ります。
[JSONHandler] は、各グループを別々のJSONオブジェクトとして扱い、グループ名をキーとして扱います。

[Group] を使用して、名前とキー値のリストからグループ属性を作成します。

	slog.Group("request",
	    "method", r.Method,
	    "url", r.URL)

[TextHandler] は、このグループを次のように表示します。

	request.method=GET request.url=http://example.com

[JSONHandler] は、次のように表示します。

	"request":{"method":"GET","url":"http://example.com"}

[Logger.WithGroup] を使用して、Loggerのすべての出力にグループ名を付けます。
LoggerでWithGroupを呼び出すと、元のハンドラと同じハンドラを持つ新しいLoggerが生成されますが、すべての属性がグループ名で修飾されます。

これにより、大規模なシステムで重複した属性キーを防止できます。
サブシステムが同じキーを使用する可能性がある場合、異なるグループ名を持つ独自のLoggerを各サブシステムに渡して、潜在的な重複を修飾します。

	logger := slog.Default().With("id", systemID)
	parserLogger := logger.WithGroup("parser")
	parseInput(input, parserLogger)

parseInputがparserLoggerでログを記録する場合、そのキーは "parser"で修飾されるため、共通のキー "id"を使用していても、ログ行には異なるキーがあります。

# Contexts

一部のハンドラは、呼び出し元で利用可能な [context.Context] から情報を取得することを望む場合があります。
トレースが有効になっている場合、現在のスパンの識別子などの情報が含まれます。

[Logger.Log] と [Logger.LogAttrs] メソッドは、対応するトップレベル関数と同様に、最初の引数としてコンテキストを取ります。

Loggerの便利なメソッド(Infoなど)と対応するトップレベル関数は、コンテキストを取りませんが、"Context"で終わる代替メソッドはコンテキストを取ります。例えば、

	slog.InfoContext(ctx, "message")

出力メソッドにコンテキストが利用可能な場合は、コンテキストを渡すことをお勧めします。

# Attrs and Values

[Attr] は、キーと値のペアです。Loggerの出力メソッドは、Attrsと交互にキーと値を受け入れます。以下の文を参照してください。

	slog.Info("hello", slog.Int("count", 3))

以下のように動作します。

	slog.Info("hello", "count", 3)

[Attr] には [Int] 、 [String] 、 [Bool] などの便利なコンストラクタがあり、一般的な型に対して、 [Any] 関数を使用して任意の型の [Attr] を構築することもできます。

[Attr] の値部分は [Value] と呼ばれる型です。
[any] のように、 [Value] は任意のGo値を保持できますが、
すべての数値と文字列を含む一般的な値を、割り当てなしで表現できます。

最も効率的なログ出力には、 [Logger.LogAttrs] を使用してください。
これは [Logger.Log] に似ていますが、交互にキーと値を受け入れるのではなく、Attrsのみを受け入れるため、これも割り当てを回避できます。

logger.LogAttrs(ctx, slog.LevelInfo, "hello", slog.Int("count", 3))

は、以下と同じ出力を生成する最も効率的な方法です。

slog.Info("hello", "count", 3)

	slog.InfoContext(ctx, "hello", "count", 3)

# タイプのログ出力のカスタマイズ

タイプが [LogValuer] インターフェースを実装している場合、その [LogValue] メソッドから返される [Value] がログ出力に使用されます。
これを使用して、タイプの値がログにどのように表示されるかを制御できます。
例えば、パスワードのような秘密情報を伏せたり、構造体のフィールドをグループ化したりすることができます。
詳細については、 [LogValuer] の例を参照してください。

<<<<<<< HEAD
LogValueメソッドは、 [LogValuer] を実装している [Value] を返すことができます。
[Value.Resolve] メソッドは、これらの場合に無限ループや無制限の再帰を回避するように注意して処理します。
ハンドラの作者やその他の人々は、LogValueを直接呼び出す代わりに、Value.Resolveを使用したい場合があります。
=======
A LogValue method may return a Value that itself implements [LogValuer]. The [Value.Resolve]
method handles these cases carefully, avoiding infinite loops and unbounded recursion.
Handler authors and others may wish to use [Value.Resolve] instead of calling LogValue directly.
>>>>>>> upstream/master

# 出力メソッドのラッピング

ロガー関数は、呼び出し元のコールスタック上でリフレクションを使用して、アプリケーション内のログ呼び出しのファイル名と行番号を検索します。
これは、slogをラップする関数に対して誤ったソース情報を生成する可能性があります。
たとえば、mylog.goファイルでこの関数を定義する場合、以下のようになります。

	func Infof(logger *slog.Logger, format string, args ...any) {
	    logger.Info(fmt.Sprintf(format, args...))


	}

そして、main.goで次のように呼び出す場合、

	Infof(slog.Default(), "hello, %s", "world")

slogは、ソースファイルをmylog.goではなくmain.goとして報告しません。

Infofの正しい実装は、ソースの場所(pc)を取得し、NewRecordに渡す必要があります。
パッケージレベルの例である "wrapping" で示されているように、Infof関数の実装方法を示します。

# レコードの操作

ハンドラが別のハンドラやバックエンドに渡す前に、レコードを変更する必要がある場合があります。
レコードには、単純な公開フィールド(例: Time、Level、Message)と、状態(属性など)を間接的に参照する非公開フィールドが混在しています。
これは、レコードの単純なコピーを変更する(例えば、属性を追加するために [Record.Add] または [Record.AddAttrs] を呼び出す)と、元のレコードに予期しない影響を与える可能性があることを意味します。
レコードを変更する前に、 [Record.Clone] を使用して、元のレコードと状態を共有しないコピーを作成するか、 [NewRecord] で新しいレコードを作成し、 [Record.Attrs] を使用して古いレコードをトラバースしてそのAttrsを構築してください。

# パフォーマンスに関する考慮事項

アプリケーションのプロファイリングによって、ログの記録にかかる時間がかなりあることが示された場合、以下の提案が役立つ場合があります。

多くのログ行に共通の属性がある場合は、 [Logger.With] を使用して、その属性を持つLoggerを作成します。
組み込みのハンドラは、 [Logger.With] の呼び出し時にその属性を1回だけフォーマットします。
[Handler] インターフェースは、その最適化を許容するように設計されており、適切に書かれたHandlerはそれを活用するはずです。

ログ呼び出しの引数は常に評価されます。たとえログイベントが破棄された場合でもです。
可能であれば、値が実際にログに記録される場合にのみ計算が行われるように遅延させてください。
たとえば、次の呼び出しを考えてみてください。

	slog.Info("starting request", "url", r.URL.String())  // may compute String unnecessarily

URL.Stringメソッドは、ロガーがInfoレベルのイベントを破棄する場合でも呼び出されます。
代わりに、URLを直接渡してください。

	slog.Info("starting request", "url", &r.URL) // calls URL.String only if needed

組み込みの [TextHandler] は、そのStringメソッドを呼び出しますが、
ログイベントが有効になっている場合にのみ呼び出します。
Stringの呼び出しを回避することは、基礎となる値の構造を保持することもできます。
例えば、 [JSONHandler] は解析されたURLのコンポーネントをJSONオブジェクトとして出力します。
String呼び出しのコストを支払うことを避けたい場合、
値の構造を検査する可能性のあるハンドラを引き起こすことなく、
その値を隠すfmt.Stringer実装でラップしてください。

[LogValuer] インターフェースを使用すると、無効なログ呼び出しで不必要な作業を回避できます。
例えば、高価な値をログに記録する必要がある場合を考えてみましょう。

	slog.Debug("frobbing", "value", computeExpensiveValue(arg))

この行が無効になっていても、computeExpensiveValueが呼び出されます。
これを回避するには、LogValuerを実装する型を定義します。

	type expensive struct { arg int }

	func (e expensive) LogValue() slog.Value {
		return slog.AnyValue(computeExpensiveValue(e.arg))
	}

そして、ログ呼び出しでその型の値を使用します。

	slog.Debug("frobbing", "value", expensive{arg})

これで、行が有効になっている場合にのみcomputeExpensiveValueが呼び出されます。

組み込みのハンドラは、各レコードが1つの塊で書き込まれることを保証するために、 [io.Writer.Write] を呼び出す前にロックを取得します。
ユーザー定義のハンドラは、自分自身のロックを管理する責任があります。

# ハンドラの作成

カスタムハンドラの作成方法についてのガイドについては、https://golang.org/s/slog-handler-guide を参照してください。
*/
package slog
