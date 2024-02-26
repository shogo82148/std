// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 編集しないでください。このファイルはRE2ディストリビューションからmksyntaxgoによって生成されました。

/*
<<<<<<< HEAD
パッケージ構文は正規表現を解析木に解析し、解析木をプログラムにコンパイルします。通常、正規表現のクライアントはこのパッケージではなく、regexpパッケージ（CompileやMatchなど）の機能を使用します。
=======
Package syntax parses regular expressions into parse trees and compiles
parse trees into programs. Most clients of regular expressions will use the
facilities of package [regexp] (such as [regexp.Compile] and [regexp.Match]) instead of this package.
>>>>>>> upstream/master

＃ 構文

<<<<<<< HEAD
Perlフラグを使用して解析する場合、このパッケージが理解する正規表現の構文は次のとおりです。Parseに代替フラグを渡すことで、構文の一部を無効にすることもできます。
=======
The regular expression syntax understood by this package when parsing with the [Perl] flag is as follows.
Parts of the syntax can be disabled by passing alternate flags to [Parse].
>>>>>>> upstream/master

単一の文字：

	.              任意の文字を含む文字（改行も含む）（フラグs=true）
	[xyz]          文字クラス
	[^xyz]         否定文字クラス
	\d             Perl文字クラス
	\D             否定Perl文字クラス
	[[:alpha:]]    ASCII文字クラス
	[[:^alpha:]]   否定ASCII文字クラス
	\pN            Unicode文字クラス（一文字の名前）
	\p{Greek}      Unicode文字クラス
	\PN            否定Unicode文字クラス（一文字の名前）
	\P{Greek}      否定Unicode文字クラス

複合：

	xy             xの後にy
	x|y            xまたはy（xを優先）

繰り返し：

	x*             xを0回以上、できれば多くの回繰り返す
	x+             xを1回以上、できれば多くの回繰り返す
	x?             xを0回または1回、できれば1回繰り返す
	x{n,m}         nまたはn+1または...またはm個のxを、できれば多くの回繰り返す
	x{n,}          n個以上のxを、できれば多くの回繰り返す
	x{n}           正確にn個のx
	x*?            xを0回以上、できれば少ない回繰り返す
	x+?            xを1回以上、できれば少ない回繰り返す
	x??            xを0回または1回、できれば0回繰り返す
	x{n,m}?        nまたはn+1または...またはm個のxを、できれば少ない回繰り返す
	x{n,}?         n個以上のxを、できれば少ない回繰り返す
	x{n}?          正確にn個のx

実装の制約：x{n,m}、x{n,}、およびx{n}の計数形式は、最小または最大の反復回数が1000を超える形式を拒否します。制限は無制限の繰り返しには適用されません。

グループ化：

	(re)           番号付きのキャプチャグループ（サブマッチ）
	(?P<name>re)   名前付き＆番号付きのキャプチャグループ（サブマッチ）
	(?<name>re)    名前付き＆番号付きのキャプチャグループ（サブマッチ）
	(?:re)         キャプチャしないグループ
	(?flags)       現在のグループ内でフラグを設定する；キャプチャしない
	(?flags:re)    re中にフラグを設定する；キャプチャしない

フラグの構文はxyz（設定）または-xyz（解除）またはxyz（設定）-z（解除）です。フラグは次のとおりです：

	i              大文字小文字を区別しない（デフォルトはfalse）
	m              マルチラインモード：～、$はテキストの始まり/終わりに加えて行の始まり/終わりにもマッチする（デフォルトはfalse）
	s              .が\nにもマッチする（デフォルトはfalse）
	U              マッチングの優先度を反転させる：x*とx*？やx+とx+？などの意味を入れ替える（デフォルトはfalse）

空の文字列：

	^              テキストまたは行の先頭（フラグm=true）
	$              テキストの終わり（\zではなく）または行の終わり（フラグm=true）
	\A             テキストの先頭
	\b             ASCIIの単語の境界（片側は\w、他側は\W、\A、または\z）
	\B             ASCIIの単語の境界ではない
	\z             テキストの終わり

エスケープシーケンス：

	\a             ベル（== \007）
	\f             改ページ（== \014）
	\t             水平タブ（== \011）
	\n             改行（== \012）
	\r             キャリッジリターン（== \015）
	\v             垂直タブ文字（== \013）
	\*             リテラルの*（任意の句読点文字用）
	\123           8進数の文字コード（最大3桁まで）
	\x7F           16進数の文字コード（正確に2桁）
	\x{10FFFF}     16進数の文字コード
	\Q...\E        句読点を含む場合でも、リテラルのテキスト...

文字クラス要素：

	x              単一の文字
	A-Z            文字範囲（包括的）
	\d             Perl文字クラス
	[:foo:]        ASCII文字クラスfoo
	\p{Foo}        Unicode文字クラスFoo
	\pF            Unicode文字クラスF（一文字の名前）

文字クラス要素としての名前付き文字クラス：

	[\d]           数字（== \d）
	[^\d]          数字以外（== \D）
	[\D]           数字以外（== \D）
	[^\D]          英数字以外（== \d）
	[[:name:]]     文字クラス[:name:]内の名前付きASCIIクラス（== [:name:]）
	[^[:name:]]    否定文字クラス[:name:]内の名前付きASCIIクラス（== [:^name:]）
	[\p{Name}]     文字クラス内の名前付きUnicodeプロパティ（== \p{Name}）
	[^\p{Name}]    否定文字クラス内の名前付きUnicodeプロパティ（== \P{Name}）

Perl文字クラス（すべてASCIIのみ）：

	\d             数字（== [0-9]）
	\D             数字以外（== [^0-9]）
	\s             空白（== [\t\n\f\r ]）
	\S             空白以外（== [^\t\n\f\r ]）
	\w             単語の文字（== [0-9A-Za-z_]）
	\W             単語の文字以外（== [^0-9A-Za-z_]）

ASCII文字クラス：

	[[:alnum:]]    英数字（== [0-9A-Za-z]）
	[[:alpha:]]    英字（== [A-Za-z]）
	[[:ascii:]]    ASCII（== [\x00-\x7F]）
	[[:blank:]]    空白（== [\t ]）
	[[:cntrl:]]    制御文字（== [\x00-\x1F\x7F]）
	[[:digit:]]    数字（== [0-9]）
	[[:graph:]]    グラフィカル（== [!-~] == [A-Za-z0-9!"#$%&'()*+,\-./:;<=>?@[\\\]^_`{|}~]）
	[[:lower:]]    小文字（== [a-z]）
	[[:print:]]    印刷可能（== [ -~] == [ [:graph:]]）
	[[:punct:]]    句読点（== [!-/:-@[-`{-~]）
	[[:space:]]    空白（== [\t\n\v\f\r ]）
	[[:upper:]]    大文字（== [A-Z]）
	[[:word:]]     単語の文字（== [0-9A-Za-z_]）
	[[:xdigit:]]   16進数の数字（== [0-9A-Fa-f]）

<<<<<<< HEAD
Unicode文字クラスは、unicode.Categoriesおよびunicode.Scriptsのものです。
=======
Unicode character classes are those in [unicode.Categories] and [unicode.Scripts].
>>>>>>> upstream/master
*/
package syntax
