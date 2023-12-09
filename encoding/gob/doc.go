// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
<<<<<<< HEAD
Package gob manages streams of gobs - binary values exchanged between an
[Encoder] (transmitter) and a [Decoder] (receiver). A typical use is transporting
arguments and results of remote procedure calls (RPCs) such as those provided by
[net/rpc].

The implementation compiles a custom codec for each data type in the stream and
is most efficient when a single [Encoder] is used to transmit a stream of values,
amortizing the cost of compilation.
=======
gobパッケージは、gobのストリームを管理します - これは、
エンコーダ（送信者）とデコーダ（受信者）間で交換されるバイナリ値です。
典型的な使用例は、リモートプロシージャコール（RPC）の引数と結果を転送することです。
これは[net/rpc]によって提供されます。

実装は、ストリーム内の各データ型に対してカスタムコーデックをコンパイルし、
コンパイルのコストを分散させるために、単一のエンコーダが値のストリームを送信するときに最も効率的です。
>>>>>>> release-branch.go1.21

# 基本

gobのストリームは自己記述的です。ストリーム内の各データ項目は、
その型の仕様によって先行します。これは、事前に定義された小さな型セットの観点から表現されます。
ポインタは伝送されませんが、それらが指すものは伝送されます。つまり、値はフラット化されます。
nilポインタは許可されていません、なぜならそれらには値がないからです。
再帰的な型はうまく動作しますが、再帰的な値（サイクルを持つデータ）は問題となります。これは変わるかもしれません。

<<<<<<< HEAD
To use gobs, create an [Encoder] and present it with a series of data items as
values or addresses that can be dereferenced to values. The [Encoder] makes sure
all type information is sent before it is needed. At the receive side, a
[Decoder] retrieves values from the encoded stream and unpacks them into local
variables.
=======
gobを使用するには、エンコーダを作成し、それに一連のデータ項目を
値または値に逆参照できるアドレスとして提示します。エンコーダは
すべての型情報が必要になる前に送信されることを確認します。受信側では、
デコーダがエンコードされたストリームから値を取得し、それらをローカル
変数に展開します。
>>>>>>> release-branch.go1.21

# 型と値

ソースと宛先の値/型は、正確に一致する必要はありません。構造体の場合、
ソースに存在するが受信変数から欠落しているフィールド（名前で識別）は無視されます。
受信変数に存在するが、送信された型または値から欠落しているフィールドは、
宛先では無視されます。同じ名前のフィールドが両方に存在する場合、
その型は互換性がなければなりません。受信者と送信者の両方が、
gobと実際のGoの値との間で必要なすべての間接参照と逆参照を行います。
例えば、スキーマ的には、gobの型は以下のようになります。

	struct { A, B int }

以下のGoの型から送信されるか、または受信することができます:

	struct { A, B int }	// 同じ
	*struct { A, B int }	// 構造体の追加の間接参照
	struct { *A, **B int }	// フィールドの追加の間接参照
	struct { A, B int64 }	// 異なる具体的な値の型; 下記参照

以下のいずれかにも受信することができます:

	struct { A, B int }	// 同じ
	struct { B, A int }	// 順序は関係ありません; 名前でマッチングします
	struct { A, B, C int }	// 追加のフィールド（C）は無視されます
	struct { B int }	// 欠落しているフィールド（A）は無視されます; データは破棄されます
	struct { B, C int }	// 欠落しているフィールド（A）は無視されます; 追加のフィールド（C）も無視されます。

これらの型に受信しようとすると、デコードエラーが発生します:

	struct { A int; B uint }	// Bの符号が変わります
	struct { A int; B float }	// Bの型が変わります
	struct { }			// 共通のフィールド名がありません
	struct { C, D int }		// 共通のフィールド名がありません

整数は2つの方法で伝送されます: 任意の精度を持つ符号付き整数または
任意の精度を持つ符号なし整数。gob形式ではint8、int16などの
区別はありません。符号付きと符号なしの整数のみが存在します。以下で
説明するように、送信者は可変長エンコーディングで値を送信します。
受信者は値を受け入れ、それを宛先変数に格納します。
浮動小数点数は常にIEEE-754 64ビット精度を使用して送信されます（以下参照）。

符号付き整数は任意の符号付き整数変数（int、int16など）に受け取ることができます。
符号なし整数は任意の符号なし整数変数に受け取ることができます。
浮動小数点数は任意の浮動小数点数変数に受け取ることができます。
ただし、宛先の変数は値を表現できる必要があります。そうでない場合、デコード操作は失敗します。

構造体、配列、スライスもサポートされています。構造体はエクスポートされた
フィールドのみをエンコードおよびデコードします。文字列とバイトの配列は、
特別で効率的な表現でサポートされています（下記参照）。スライスがデコードされるとき、
既存のスライスに容量がある場合、スライスはその場で拡張されます。そうでない場合、
新しい配列が割り当てられます。いずれにせよ、結果として得られるスライスの長さは、
デコードされた要素の数を報告します。

一般的に、割り当てが必要な場合、デコーダはメモリを割り当てます。そうでない場合、
ストリームから読み取った値で宛先の変数を更新します。最初にそれらを初期化することはありません。
したがって、宛先がマップ、構造体、またはスライスなどの複合値である場合、
デコードされた値は要素ごとに既存の変数にマージされます。

関数とチャネルはgobで送信されません。そのような値をトップレベルでエンコードしようとすると失敗します。
chan型またはfunc型の構造体フィールドは、エクスポートされていないフィールドと全く同じように扱われ、無視されます。

<<<<<<< HEAD
Gob can encode a value of any type implementing the [GobEncoder] or
[encoding.BinaryMarshaler] interfaces by calling the corresponding method,
in that order of preference.

Gob can decode a value of any type implementing the [GobDecoder] or
[encoding.BinaryUnmarshaler] interfaces by calling the corresponding method,
again in that order of preference.
=======
Gobは、GobEncoderまたはencoding.BinaryMarshalerインターフェースを実装する任意の型の値をエンコードできます。
これは、その順序の優先度で対応するメソッドを呼び出すことによって行われます。

Gobは、GobDecoderまたはencoding.BinaryUnmarshalerインターフェースを実装する任意の型の値をデコードできます。
これは、その順序の優先度で対応するメソッドを呼び出すことによって行われます。
>>>>>>> release-branch.go1.21

# エンコーディングの詳細

このセクションでは、ほとんどのユーザーにとって重要でないエンコーディングの詳細を文書化します。
詳細は下から上に提示されます。

符号なし整数は2つの方法のいずれかで送信されます。それが128未満の場合、その値を持つバイトとして送信されます。
それ以外の場合、それは最小長のビッグエンディアン（高バイト先）バイトストリームとして送信され、
その前にバイト数を保持する1バイトが先行します。このバイト数は否定されます。
したがって、0は（00）として送信され、7は（07）として送信され、256は（FE 01 00）として送信されます。

ブール値は符号なし整数内にエンコードされます: falseの場合は0、trueの場合は1。

符号付き整数iは、符号なし整数u内にエンコードされます。u内で、ビット1以上が値を含み、
ビット0は受信時にそれらを補完するかどうかを示します。エンコードアルゴリズムは次のようになります:

	var u uint
	if i < 0 {
		u = (^uint(i) << 1) | 1 // iを補完し、ビット0を1にします
	} else {
		u = (uint(i) << 1) // iを補完しない、ビット0は0です
	}
	encodeUnsigned(u)

したがって、低ビットは符号ビットに類似していますが、それを補完ビットにすることで、
最大の負の整数が特別なケースにならないことを保証します。例えば、-129=^128=(^256>>1)は(FE 01 01)としてエンコードされます。

<<<<<<< HEAD
Floating-point numbers are always sent as a representation of a float64 value.
That value is converted to a uint64 using [math.Float64bits]. The uint64 is then
byte-reversed and sent as a regular unsigned integer. The byte-reversal means the
exponent and high-precision part of the mantissa go first. Since the low bits are
often zero, this can save encoding bytes. For instance, 17.0 is encoded in only
three bytes (FE 31 40).
=======
浮動小数点数は常にfloat64値の表現として送信されます。
その値はmath.Float64bitsを使用してuint64に変換されます。そのuint64は
バイト反転され、通常の符号なし整数として送信されます。バイト反転により、
指数とマンティッサの高精度部分が最初に来ます。低ビットはしばしばゼロなので、
これによりエンコーディングバイトを節約できます。例えば、17.0は
3バイト（FE 31 40）でエンコードされます。
>>>>>>> release-branch.go1.21

文字列とバイトのスライスは、符号なしのカウントとその後に続くその値の
未解釈のバイトとして送信されます。

その他のすべてのスライスと配列は、符号なしのカウントに続いてその要素数だけが
その型の標準的なgobエンコーディングを使用して再帰的に送信されます。

マップは、符号なしのカウントに続いてその数だけのキー、要素のペアとして送信されます。
空だがnilでないマップは送信されるので、受信者がすでに割り当てていない場合、
送信されたマップがnilでなく、トップレベルでない限り、受信時に常に割り当てられます。

スライスや配列、マップでは、すべての要素、ゼロ値の要素であっても、
すべての要素がゼロであっても、送信されます。

構造体は、(フィールド番号、フィールド値)のペアのシーケンスとして送信されます。フィールド
値はその型の標準的なgobエンコーディングを使用して、再帰的に送信されます。フィールドが
その型のゼロ値を持つ場合（配列を除く; 上記参照）、それは伝送から省略されます。フィールド番号は
エンコードされた構造体の型によって定義されます: エンコードされた型の最初のフィールドはフィールド0、
次のフィールドはフィールド1、等です。値をエンコードするとき、フィールド番号は効率のために
デルタエンコードされ、フィールドは常にフィールド番号の増加順に送信されます; したがって、デルタは
符号なしです。デルタエンコーディングの初期化はフィールド番号を-1に設定するので、値が7の符号なし整数フィールド0は
符号なしデルタ=1、符号なし値=7または(01 07)として送信されます。最後に、すべてのフィールドが
送信された後、終端マークが構造体の終わりを示します。そのマークはデルタ=0の
値で、表現は(00)です。

<<<<<<< HEAD
Interface types are not checked for compatibility; all interface types are
treated, for transmission, as members of a single "interface" type, analogous to
int or []byte - in effect they're all treated as interface{}. Interface values
are transmitted as a string identifying the concrete type being sent (a name
that must be pre-defined by calling [Register]), followed by a byte count of the
length of the following data (so the value can be skipped if it cannot be
stored), followed by the usual encoding of concrete (dynamic) value stored in
the interface value. (A nil interface value is identified by the empty string
and transmits no value.) Upon receipt, the decoder verifies that the unpacked
concrete item satisfies the interface of the receiving variable.

If a value is passed to [Encoder.Encode] and the type is not a struct (or pointer to struct,
etc.), for simplicity of processing it is represented as a struct of one field.
The only visible effect of this is to encode a zero byte after the value, just as
after the last field of an encoded struct, so that the decode algorithm knows when
the top-level value is complete.

The representation of types is described below. When a type is defined on a given
connection between an [Encoder] and [Decoder], it is assigned a signed integer type
id. When [Encoder.Encode](v) is called, it makes sure there is an id assigned for
the type of v and all its elements and then it sends the pair (typeid, encoded-v)
where typeid is the type id of the encoded type of v and encoded-v is the gob
encoding of the value v.
=======
インターフェース型は互換性がチェックされません。すべてのインターフェース型は、
伝送のために、単一の "interface" 型のメンバーとして扱われます。これはintや[]byteに類似しています。
効果的に、すべてが interface{} として扱われます。インターフェース値は、送信される具体的な型を
識別する文字列として送信されます（この名前はRegisterを呼び出すことで事前に定義する必要があります）。
次に、次のデータの長さのバイト数（値が格納できない場合に値をスキップできるように）、
次にインターフェース値に格納されている具体的（動的）値の通常のエンコーディングが続きます。
（nilのインターフェース値は空の文字列によって識別され、値は送信されません。）
受信時に、デコーダは展開された具体的なアイテムが受信変数のインターフェースを満たしていることを確認します。

値がEncodeに渡され、その型が構造体（または構造体へのポインタなど）でない場合、
処理の簡便性のために、それは1つのフィールドを持つ構造体として表現されます。
これによる唯一の可視的な効果は、エンコードされた構造体の最後のフィールドの後と同様に、
値の後にゼロバイトをエンコードすることで、デコードアルゴリズムがトップレベルの値が完了したことを知ることができます。

型の表現については以下に説明します。型がEncoderとDecoderの間の特定の
接続で定義されると、それには符号付き整数型の
idが割り当てられます。Encoder.Encode(v)が呼び出されると、vの型とそのすべての要素に
idが割り当てられていることを確認し、次にペア(typeid, encoded-v)を送信します。
ここで、typeidはvのエンコードされた型の型idであり、encoded-vは値vのgob
エンコーディングです。
>>>>>>> release-branch.go1.21

型を定義するために、エンコーダは未使用の正の型idを選択し、
ペア(-type id, encoded-type)を送信します。ここで、encoded-typeはwireType
記述のgobエンコーディングで、これらの型から構築されます:

	type wireType struct {
		ArrayT           *ArrayType
		SliceT           *SliceType
		StructT          *StructType
		MapT             *MapType
		GobEncoderT      *gobEncoderType
		BinaryMarshalerT *gobEncoderType
		TextMarshalerT   *gobEncoderType

	}
	type arrayType struct {
		CommonType
		Elem typeId
		Len  int
	}
	type CommonType struct {
		Name string // the name of the struct type
		Id  int    // the id of the type, repeated so it's inside the type
	}
	type sliceType struct {
		CommonType
		Elem typeId
	}
	type structType struct {
		CommonType
		Field []*fieldType // the fields of the struct.
	}
	type fieldType struct {
		Name string // the name of the field.
		Id   int    // the type id of the field, which must be already defined
	}
	type mapType struct {
		CommonType
		Key  typeId
		Elem typeId
	}
	type gobEncoderType struct {
		CommonType
	}

ネストした型idがある場合、すべての内部型idの型が定義されている必要があります。
これは、トップレベルの型idがencoded-vを記述するために使用される前に行われます。

設定の簡便性のため、接続はこれらの型をa priori（事前に）理解するように定義されており、
基本的なgob型（int、uintなど）も理解します。それらのidは以下の通りです:

	bool        1
	int         2
	uint        3
	float       4
	[]byte      5
	string      6
	complex     7
	interface   8
	// gap for reserved ids.
	WireType    16
	ArrayType   17
	CommonType  18
	SliceType   19
	StructType  20
	FieldType   21
	// 22 is slice of fieldType.
	MapType     23

最後に、Encodeの呼び出しによって作成された各メッセージは、メッセージ内の残りのバイト数の
符号なし整数カウントによって先行します。初期型名の後、インターフェース値は同じように
ラップされます。効果的に、インターフェース値はEncodeの再帰的な呼び出しのように動作します。

要約すると、gobストリームは次のように見えます

	(byteCount (-type id, encoding of a wireType)* (type id, encoding of a value))*

ここで * はゼロ回以上の繰り返しを示し、値の型idは事前に定義されているか、
ストリーム内で値の前に定義されていなければなりません。

互換性: このパッケージへの将来の変更は、以前のバージョンを使用してエンコードされたストリームとの
互換性を維持するよう努力します。つまり、このパッケージのリリースされたバージョンは、
セキュリティ修正などの問題を除いて、以前にリリースされたバージョンで書かれたデータを
デコードできるはずです。背景についてはGoの互換性ドキュメントを参照してください: https://golang.org/doc/go1compat

gobワイヤーフォーマットの設計についての議論は「Gobs of data」を参照してください:
https://blog.golang.org/gobs-of-data

# セキュリティ

<<<<<<< HEAD
This package is not designed to be hardened against adversarial inputs, and is
outside the scope of https://go.dev/security/policy. In particular, the [Decoder]
does only basic sanity checking on decoded input sizes, and its limits are not
configurable. Care should be taken when decoding gob data from untrusted
sources, which may consume significant resources.
=======
このパッケージは、敵対的な入力に対して強化されるように設計されていませんし、
https://go.dev/security/policy の範囲外です。特に、Decoderはデコードされた入力サイズに対して
基本的な健全性チェックのみを行い、その制限は設定可能ではありません。信頼できない
ソースからのgobデータをデコードする際には注意が必要であり、大量のリソースを消費する可能性があります。
>>>>>>> release-branch.go1.21
*/
package gob
