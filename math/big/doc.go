// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
// bigパッケージは任意精度算術（大きな数）を実装します。
// 以下の数値型がサポートされています:

//	Int    符号付き整数
//	Rat    有理数
//	Float  浮動小数点数

<<<<<<< HEAD
The zero value for an [Int], [Rat], or [Float] correspond to 0. Thus, new
values can be declared in the usual ways and denote 0 without further
initialization:
=======
Int、Rat、またはFloatのゼロ値は0に対応します。したがって、新しい
値は通常の方法で宣言でき、さらなる初期化なしで0を示します：
>>>>>>> release-branch.go1.21

	var x Int        // &xは値0の*Intです
	var r = &Rat{}   // rは値0の*Ratです
	y := new(Float)  // yは値0の*Floatです

あるいは、新しい値は以下の形式のファクトリ関数で割り当てて初期化することができます：

	func NewT(v V) *T

<<<<<<< HEAD
For instance, [NewInt](x) returns an *[Int] set to the value of the int64
argument x, [NewRat](a, b) returns a *[Rat] set to the fraction a/b where
a and b are int64 values, and [NewFloat](f) returns a *[Float] initialized
to the float64 argument f. More flexibility is provided with explicit
setters, for instance:
=======
例えば、NewInt(x)はint64引数xの値に設定された*Intを返し、
NewRat(a, b)はaとbがint64値である分数a/bに設定された*Ratを返し、
NewFloat(f)はfloat64引数fに初期化された*Floatを返します。
より柔軟性を提供するために、明示的なセッターが提供されています。例えば：
>>>>>>> release-branch.go1.21

	var z1 Int
	z1.SetUint64(123)                 // z1 := 123
	z2 := new(Rat).SetFloat64(1.25)   // z2 := 5/4
	z3 := new(Float).SetInt(z1)       // z3 := 123.0

セッター、数値演算、および述語は、以下の形式のメソッドとして表現されます：

	func (z *T) SetV(v V) *T          // z = v
	func (z *T) Unary(x *T) *T        // z = unary x
	func (z *T) Binary(x, y *T) *T    // z = x binary y
	func (x *T) Pred() P              // p = pred(x)

<<<<<<< HEAD
with T one of [Int], [Rat], or [Float]. For unary and binary operations, the
result is the receiver (usually named z in that case; see below); if it
is one of the operands x or y it may be safely overwritten (and its memory
reused).
=======
TはInt、Rat、またはFloatのいずれかです。単項および二項演算の場合、
結果はレシーバ（通常その場合はzと名付けられます。以下参照）であり、
それがオペランドxまたはyのいずれかであれば、安全に上書き（およびそのメモリの再利用）が可能です。
>>>>>>> release-branch.go1.21

算術式は通常、個々のメソッド呼び出しのシーケンスとして書かれ、
各呼び出しが操作に対応します。レシーバは結果を示し、
メソッドの引数は操作のオペランドです。
例えば、*Int値a、b、cが与えられた場合、次の呼び出し

	c.Add(a, b)

これは、a + bの和を計算し、結果をcに格納します。これにより、
以前にcに格納されていた値は上書きされます。特に指定がない限り、
操作はパラメータのエイリアシングを許可するため、次のように書いても問題ありません。

	sum.Add(sum, x)

これにより、値xがsumに累積されます。

(常にレシーバ経由で結果値を渡すことにより、メモリの使用を
はるかによく制御できます。各結果に新たにメモリを割り当てる代わりに、
操作は結果値に割り当てられたスペースを再利用し、
そのプロセスで新しい結果でその値を上書きすることができます。)

表記法の規則：APIで一貫して名前が付けられている入力メソッドパラメータ（レシーバを含む）
は、その使用法を明確にするためです。入力オペランドは通常、x、y、a、bなどと名付けられますが、
zとは名付けられません。結果を指定するパラメータはzと名付けられます（通常はレシーバ）。

例えば、(*Int).Addの引数はxとyと名付けられています。
そして、レシーバが結果の格納先を指定するため、それはzと呼ばれます：

	func (z *Int) Add(x, y *Int) *Int

この形式のメソッドは、単純な呼び出しチェーンを可能にするため、通常、受け取ったレシーバも返します。

<<<<<<< HEAD
Methods which don't require a result value to be passed in (for instance,
[Int.Sign]), simply return the result. In this case, the receiver is typically
the first operand, named x:

	func (x *Int) Sign() int

Various methods support conversions between strings and corresponding
numeric values, and vice versa: *[Int], *[Rat], and *[Float] values implement
the Stringer interface for a (default) string representation of the value,
but also provide SetString methods to initialize a value from a string in
a variety of supported formats (see the respective SetString documentation).

Finally, *[Int], *[Rat], and *[Float] satisfy [fmt.Scanner] for scanning
and (except for *[Rat]) the Formatter interface for formatted printing.
=======
結果値を渡す必要がないメソッド（例えば、Int.Sign）は、単に結果を返します。
この場合、レシーバは通常、最初のオペランドで、xと名付けられます：

	func (x *Int) Sign() int

さまざまなメソッドが文字列と対応する数値との間の変換をサポートしており、その逆も可能です：
*Int、*Rat、および*Floatの値は、値の（デフォルトの）文字列表現のためのStringerインターフェースを実装しますが、
また、さまざまなサポートされている形式で文字列から値を初期化するためのSetStringメソッドも提供します
（それぞれのSetStringのドキュメンテーションを参照してください）。

最後に、*Int、*Rat、および*Floatは、スキャンのための[fmt.Scanner]を満たし、
（*Ratを除いて）フォーマットされた印刷のためのFormatterインターフェースを満たします。
>>>>>>> release-branch.go1.21
*/
package big
