// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
fmtパッケージは、Cのprintfおよびscanfに類似した関数を使用してフォーマットされたI/Oを実装します。
フォーマット'verbs'はCから派生していますが、よりシンプルです。

# 印刷

動詞：

一般：

	%v	デフォルトのフォーマットでの値
		構造体を印刷する場合、プラスフラグ（%+v）はフィールド名を追加します
	%#v	値のGo構文表現
		（浮動小数点の無限大とNaNは±InfとNaNとして印刷されます）
	%T	値の型のGo構文表現
	%%	リテラルのパーセント記号；値を消費しません

ブール値：

	%t	trueまたはfalseの単語

整数：

	%b	2進数
	%c	対応するUnicodeコードポイントによって表される文字
	%d	10進数
	%o	8進数
	%O	0o接頭辞を付けた8進数
	%q	Go構文で安全にエスケープされたシングルクォート文字リテラル。
	%x	16進数（小文字のa-f）
	%X	16進数（大文字のA-F）
	%U	Unicode形式：U+1234； "U+%04X"と同じ

浮動小数点数と複素数の構成要素：

	%b	2の累乗の指数で表される10進数の科学的表記法。
	    'b'フォーマットを使用したstrconv.FormatFloatの方法で、例：-123456p-78
	%e	科学的表記法、例：-1.234456e+78
	%E	科学的表記法、例：-1.234456E+78
	%f	小数点を含むが指数を含まない表記法、例：123.456
	%F	%fの同義語
	%g	大きな指数の場合は%e、それ以外の場合は%f。精度については後述します。
	%G	大きな指数の場合は%E、それ以外の場合は%F
	%x	16進数表記法（2の累乗の10進数指数を含む）、例：-0x1.23abcp+20
	%X	大文字の16進数表記法、例：-0X1.23ABCP+20

文字列とバイトスライス（これらの動詞では同等に扱われます）：

	%s	文字列またはスライスの解釈されていないバイト
	%q	Go構文で安全にエスケープされたダブルクォート文字列
	%x	16進数、小文字、1バイトあたり2文字
	%X	16進数、大文字、1バイトあたり2文字

スライス：

	%p	16進表記法で表された0番目の要素のアドレス。先頭に0xが付きます。

ポインタ：

	%p	先頭に0xが付いた16進表記法
	%b、%d、%o、%x、%Xの動詞は、ポインタでも整数と同じように機能し、
	整数であるかのように値をフォーマットします。

%vのデフォルトフォーマットは次のとおりです。

	bool:                    %t
	int, int8 etc.:          %d
	uint, uint8 etc.:        %d, %#vで出力された場合は%#x
	float32, complex64, etc: %g
	string:                  %s
	chan:                    %p
	pointer:                 %p

複合オブジェクトの場合、要素は再帰的にこれらのルールを使用して印刷され、次のようにレイアウトされます。

	struct:             {field0 field1 ...}
	array, slice:       [elem0 elem1 ...]
	maps:               map[key1:value1 key2:value2 ...]
	上記へのポインタ：     &{}, &[], &map[]

幅は、動詞の直前にオプションの10進数で指定されます。
省略された場合、幅は値を表すために必要なものです。
精度は、（オプションの）幅の後に、ピリオドに続く10進数で指定されます。
ピリオドが存在しない場合、デフォルトの精度が使用されます。
後続の数字のないピリオドは、精度0を指定します。
例：

	%f     デフォルトの幅、デフォルトの精度
	%9f    幅9、デフォルトの精度
	%.2f   デフォルトの幅、精度2
	%9.2f  幅9、精度2
	%9.f   幅9、精度0

幅と精度は、Unicodeコードポイントの単位で測定されます。
つまり、ルーンです。（これは、Cのprintfと異なり、常にバイトで測定される単位です。）
'*'文字でどちらかまたは両方のフラグを置き換えることができます。
これにより、次のオペランド（フォーマットする前のオペランドの前）から値を取得できます。
このオペランドはint型でなければなりません。

ほとんどの値について、幅は出力する最小のルーン数であり、必要に応じてフォーマットされた形式をスペースでパディングします。

ただし、文字列、バイトスライス、およびバイト配列の場合、精度は入力をフォーマットするための長さを制限します（出力のサイズではありません）。
必要に応じて切り捨てます。通常、ルーンで測定されますが、これらの型を%xまたは%Xフォーマットでフォーマットする場合はバイトで測定されます。

浮動小数点数の場合、幅はフィールドの最小幅を設定し、
精度は必要に応じて小数点以下の桁数を設定します。
ただし、%g /%Gの場合、精度は最大有効桁数を設定します（末尾のゼロは削除されます）。
たとえば、12.345が与えられた場合、フォーマット%6.3fは12.345を印刷し、
%.3gは12.3を印刷します。%e、%f、および%#gのデフォルトの精度は6です。
％gの場合、値を一意に識別するために必要な最小桁数です。

複素数の場合、幅と精度は2つの要素に独立に適用され、結果は括弧で囲まれます。
したがって、1.2 + 3.4iに適用された％fは（1.200000 + 3.400000i）を生成します。

整数コードポイントまたはルーン文字列（[]rune型）を%qでフォーマットする場合、
無効なUnicodeコードポイントは、[strconv.QuoteRune] のようにUnicode置換文字U+FFFDに変更されます。

その他のフラグ：

	'+'	数値の値に対して常に符号を表示します。
		%q (%+q)の出力をASCIIのみに保証します。
	'-'	左側ではなく右側にスペースでパディングします（フィールドを左揃えにします）。
	'#'	代替フォーマット: 二進数(%#b)の場合は先頭に0bを追加、8進数(%#o)の場合は0を追加、
		16進数(%#x または %#X)の場合は0xまたは0Xを追加します。%p (%#p)の0xを抑制します。
		%qの場合、[strconv.CanBackquote] がtrueを返す場合は生の（バッククォートされた）文字列を表示します。
		%e、%E、%f、%F、%g、%Gに対して常に小数点を表示します。
		%gと%Gの末尾のゼロを削除しません。
		文字が印刷可能な場合は、例えばU+0078 'x'と表示します（%U (%#U)）。
	' '	（スペース）数値の省略された符号のスペースを残します（% d）。
		16進数で文字列やスライスを印刷するときにバイト間にスペースを入れます（% x、% X）。
	'0'	スペースではなく先頭ゼロでパディングします。
		数値の場合、これは符号の後にパディングを移動します。

動詞がそれらを期待していない場合、フラグは無視されます。
たとえば、代替の10進数フォーマットがないため、%#dと%dは同じように動作します。

Printfのような各関数に対して、フォーマットを取らないPrint関数もあります。
これは、すべてのオペランドに対して%vと言うのと同等です。
別のバリアントであるPrintlnは、オペランド間に空白を挿入し、改行を追加します。

動詞に関係なく、オペランドがインターフェース値である場合、
インターフェース自体ではなく、内部の具体的な値が使用されます。
したがって、次のようになります。

	var i interface{} = 23
	fmt.Printf("%v\n", i)

とすると、23が出力されます。

％Tおよび％p動詞を使用して印刷される場合を除き、
特定のインターフェースを実装するオペランドには特別なフォーマットが適用されます。
適用順序は次のとおりです。

1. オペランドが [reflect.Value] である場合、オペランドは保持する具体的な値に置き換えられ、
次のルールで印刷が続行されます。

2. オペランドが [Formatter] インターフェースを実装している場合、
それが呼び出されます。この場合、動詞とフラグの解釈はその実装によって制御されます。

3. オペランドが [GoStringer] インターフェースを実装している場合、
%v動詞が#フラグとともに使用（%#v）され、それが呼び出されます。

フォーマット（[Println] などの暗黙的な%v）が文字列（%s %q %x %X）に対して有効である場合、または%vであり、%#vではない場合、次の2つのルールが適用されます。

4. オペランドがerrorインターフェースを実装している場合、Errorメソッドが呼び出され、
オブジェクトが文字列に変換され、動詞に必要な形式でフォーマットされます（ある場合）。

5. オペランドがString() stringメソッドを実装している場合、
そのメソッドが呼び出され、オブジェクトが文字列に変換され、
動詞に必要な形式でフォーマットされます（ある場合）。

スライスや構造体などの複合オペランドの場合、フォーマットは各オペランドの要素に再帰的に適用され、
オペランド全体には適用されません。したがって、%qは文字列のスライスの各要素を引用し、
%6.2fは浮動小数点数の配列の各要素のフォーマットを制御します。

ただし、文字列のような動詞（％s％q％x％X）を使用してバイトスライスを印刷する場合、
バイトスライスは文字列と同じように、単一のアイテムとして扱われます。

次のような再帰を避けるために

	type X string
	func (x X) String() string { return Sprintf("<%s>", x) }

再帰する前に値を変換してください:

	func (x X) String() string { return Sprintf("<%s>", string(x)) }

また、自己参照するデータ構造（スライスなど）が、その型にStringメソッドがある場合、
無限再帰がトリガーされることがあります。しかし、そのような病理はまれであり、
パッケージはそれらに対して保護しません。

構造体を出力する場合、fmtはエクスポートされていないフィールドに対してErrorやStringなどの
フォーマットメソッドを呼び出すことができないため、呼び出しません。

# 明示的な引数インデックス

[Printf]、[Sprintf]、および [Fprintf] では、各フォーマット指定子が呼び出し時に渡された
引数を順番にフォーマットすることがデフォルトの動作です。
ただし、動詞の直前に[n]という表記がある場合、n番目の1から始まる引数が代わりに
フォーマットされることを示します。 幅または精度の'*'の前に同じ表記がある場合、
値を保持する引数インデックスが選択されます。 [n]の括弧式を処理した後、
後続の動詞は、別の指示がない限り、引数n + 1、n + 2などを使用します。

例えば、

	fmt.Sprintf("%[2]d %[1]d\n", 11, 22)

は "22 11" を生成します。一方、

	fmt.Sprintf("%[3]*.[2]*[1]f", 12.0, 2, 6)

は

	fmt.Sprintf("%6.2f", 12.0)

と同等であり、 " 12.00" を生成します。明示的なインデックスは後続の動詞に影響を与えるため、
最初の引数のインデックスをリセットして同じ値を複数回出力するために使用できます。

	fmt.Sprintf("%d %d %#[1]x %#x", 16, 17)

は "16 17 0x10 0x11" を生成します。

# Format errors

動詞に対して無効な引数が指定された場合、例えば%dに対して文字列が提供された場合、
生成された文字列には問題の説明が含まれます。次の例のように:

	Wrong type or unknown verb: %!verb(type=value)
		Printf("%d", "hi"):        %!d(string=hi)
	Too many arguments: %!(EXTRA type=value)
		Printf("hi", "guys"):      hi%!(EXTRA string=guys)
	Too few arguments: %!verb(MISSING)
		Printf("hi%d"):            hi%!d(MISSING)
	Non-int for width or precision: %!(BADWIDTH) or %!(BADPREC)
		Printf("%*s", 4.5, "hi"):  %!(BADWIDTH)hi
		Printf("%.*s", 4.5, "hi"): %!(BADPREC)hi
	Invalid or invalid use of argument index: %!(BADINDEX)
		Printf("%*[2]d", 7):       %!d(BADINDEX)
		Printf("%.[2]d", 7):       %!d(BADINDEX)

すべてのエラーは、文字列「%!」で始まり、時には1文字（動詞）が続き、
かっこで囲まれた説明で終わります。

printルーチンによって呼び出されたときにErrorまたはStringメソッドがパニックを引き起こす場合、
fmtパッケージはパニックからエラーメッセージを再フォーマットし、fmtパッケージを介して
渡されたことを示す装飾を付けます。 たとえば、Stringメソッドがpanic("bad")を呼び出す場合、
生成されたフォーマットされたメッセージは次のようになります。

	%!s(PANIC=bad)

%!sは、失敗が発生したときに使用されるプリント動詞を示すだけです。
ただし、パニックがErrorまたはStringメソッドに対するnilレシーバによって引き起こされる場合、
出力は装飾されていない文字列「<nil>」です。

# スキャン

フォーマットされたテキストをスキャンして値を生成する、同様の関数群があります。
[Scan]、[Scanf]、および [Scanln] は [os.Stdin] から読み取ります。
[Fscan]、[Fscanf]、および [Fscanln] は指定された [io.Reader] から読み取ります。
[Sscan]、[Sscanf]、および [Sscanln] は引数文字列から読み取ります。

[Scan]、[Fscan]、[Sscan] は、入力の改行をスペースとして扱います。

[Scanln]、[Fscanln]、および [Sscanln] は、改行でスキャンを停止し、
アイテムの後に改行またはEOFが続くことを要求します。

[Scanf]、[Fscanf]、および [Sscanf] は、[Printf] と同様のフォーマット文字列に従って引数を解析します。
以下のテキストでは、「スペース」とは、改行以外の任意のUnicode空白文字を意味します。

フォーマット文字列では、%文字で導入された動詞が入力を消費して解析されます。
これらの動詞については、以下で詳しく説明します。フォーマット以外の文字（%、スペース、
改行以外）は、正確にその入力文字を消費し、存在する必要があります。フォーマット文字列
内の改行の前に0個以上のスペースがある場合、0個以上のスペースを入力で消費して、
単一の改行または入力の終わりに続く改行を消費します。フォーマット文字列内の改行の後に
スペースが続く場合、入力で0個以上のスペースを消費します。それ以外の場合、フォーマット
文字列内の1つ以上のスペースの実行は、入力で可能な限り多くのスペースを消費します。
フォーマット文字列内のスペースの実行が改行に隣接していない場合、実行は入力から少なくとも
1つのスペースを消費するか、入力の終わりを見つける必要があります。

スペースと改行の処理は、Cのscanfファミリーとは異なります。
Cでは、改行は他のスペースと同様に扱われ、フォーマット文字列内のスペースの実行が
入力で消費するスペースが見つからない場合でもエラーは発生しません。

動詞は [Printf] と同様に動作します。
たとえば、%xは整数を16進数としてスキャンし、%vは値のデフォルト表現形式をスキャンします。
[Printf] の動詞%pと%T、フラグ#と+は実装されていません。
浮動小数点数と複素数の場合、すべての有効なフォーマット動詞
（%b %e %E %f %F %g %G %x %Xおよび%v）は同等であり、
10進数と16進数の表記法の両方を受け入れます（たとえば、「2.3e + 7」、「0x4.5p-8」）
および数字を区切るアンダースコア（たとえば、「3.14159_26535_89793」）。

動詞によって処理される入力は、暗黙的にスペースで区切られます。
%cを除くすべての動詞の実装は、残りの入力から先頭のスペースを破棄して開始し、
%s動詞（および文字列に読み込む%v）は、最初のスペースまたは改行文字で入力の消費を停止します。

整数をフォーマット指定子なしまたは%v動詞でスキャンする場合、
0b（バイナリ）、0oおよび0（8進数）、0x（16進数）のよく知られた基本設定の接頭辞が受け入れられます。
数字を区切るアンダースコアも受け入れられます。

幅は入力テキストで解釈されますが、精度を指定する構文はありません（%5.2fではなく、%5fのみ）。
幅が指定された場合、先頭のスペースがトリムされた後に適用され、動詞を満たすために読み取る最大ルーン数を指定します。
例えば、

	Sscanf(" 1234567 ", "%5s%d", &s, &i)

は、sを「12345」に、iを67に設定しますが、

	Sscanf(" 12 34 567 ", "%5s%d", &s, &i)

は、sを「12」に、iを34に設定します。

すべてのスキャン関数において、改行文字の直後にすぐにキャリッジリターンがある場合、
それは通常の改行文字として扱われます。
(\r\n は \n と同じ意味を持ちます。)

すべてのスキャン関数において、オペランドが [Scan] メソッドを実装している場合、
つまり [Scanner] インターフェースを実装している場合、そのメソッドがそのオペランドのテキストをスキャンするために使用されます。
また、スキャンされた引数の数が提供された引数の数よりも少ない場合、エラーが返されます。

スキャンするすべての引数は、基本型のポインタまたは [Scanner] インターフェースの実装である必要があります。

[Scanf] や [Fscanf] のように、[Sscanf] は入力全体を消費する必要はありません。
[Sscanf] が使用した入力文字列の量を回復する方法はありません。

注意: [Fscan] などは、返された入力の1文字（rune）を読み取ることができます。
これは、スキャンルーチンを呼び出すループが入力の一部をスキップする可能性があることを意味します。
これは通常、入力値の間にスペースがない場合にのみ問題になります。
[Fscan] に提供されたリーダーがReadRuneを実装している場合、そのメソッドが文字を読み取るために使用されます。
また、リーダーがUnreadRuneも実装している場合、そのメソッドが文字を保存し、連続した呼び出しでデータが失われないようにします。
ReadRuneとUnreadRuneのメソッドを持たないリーダーにReadRuneとUnreadRuneのメソッドをアタッチするには、[bufio.NewReader] を使用します。
*/
package fmt
