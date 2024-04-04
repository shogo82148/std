// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect

import (
	"github.com/shogo82148/std/internal/abi"
	"github.com/shogo82148/std/unsafe"
)

// ValueはGo言語の値に対する反射インターフェースです。
//
// すべてのメソッドがすべての種類の値に適用されるわけではありません。制限事項がある場合は、各メソッドのドキュメントに記載されています。
// kind-specificメソッドを呼び出す前に、値の種類を特定するためにKindメソッドを使用してください。型に適切でないメソッドを呼び出すと、ランタイムパニックが発生します。
//
// ゼロ値は何も表していません。
// その [Value.IsValid] メソッドはfalseを返し、Kindメソッドは [Invalid] を返し、Stringメソッドは"<invalid Value>"を返し、他のすべてのメソッドはパニックを発生させます。
// ほとんどの関数やメソッドは無効な値を返さないでしょう。
// もし返す場合は、そのドキュメントに条件が明示されています。
//
// Valueは、基礎となるGoの値が同等の直接操作で並行に使用できる場合、複数のゴルーチンで同時に使用することができます。
//
// 2つのValueを比較するには、Interfaceメソッドの結果を比較してください。
// ==を使用して2つのValueを比較すると、それらが表す基礎の値を比較しません。
type Value struct {

	// typ_はValueによって表される値の型を保持します。
	// vのエスケープを避けるために、typメソッドを使用してアクセスしてください。
	typ_ *abi.Type

	// ポインタ値のデータまたは、flagIndirが設定されている場合はデータへのポインタ。
	// flagIndirが設定されているか、typ.pointers()が true の場合に有効です。
	ptr unsafe.Pointer

	// flagは値に関するメタデータを保持しています。
	//
	// 最下位の5ビットは値の種類（typ.Kind()と同じ）を示します。
	//
	// 次のビットはフラグビットです：
	//	- flagStickyRO：非公開で埋め込まれていないフィールドから取得されたため、読み取り専用
	//	- flagEmbedRO：非公開の埋め込まれたフィールドから取得されたため、読み取り専用
	//	- flagIndir：valはデータへのポインタを保持しています
	//	- flagAddr：v.CanAddrがtrue（flagIndirとptrが非nilであることを意味します）
	//	- flagMethod：vはメソッドの値です。
	// ifaceIndir(typ)の場合、codeはflagIndirが設定されていると想定することができます。
	//
	// 残りの22ビット以上は、メソッドの番号を示します。
	// flag.kind() != Funcの場合、codeはflagMethodが未設定であると想定することができます。
	flag
}

// ValueErrorは、Valueメソッドがサポートされていない [Value] に対してValueメソッドが呼び出された場合に発生します。このような場合は、各メソッドの説明に記載されています。
type ValueError struct {
	Method string
	Kind   Kind
}

func (e *ValueError) Error() string

// Addrはvのアドレスを表すポインタ値を返します。
// [Value.CanAddr] がfalseを返す場合にはパニックを発生させます。
// Addrは通常、構造体のフィールドやスライスの要素に対して
// ポインタレシーバを必要とするメソッドを呼び出すために
// ポインタを取得するために使用されます。
func (v Value) Addr() Value

// Boolはvの基礎となる値を返します。
// vの種類が [Bool] でない場合、パニックを起こします。
func (v Value) Bool() bool

// Bytesはvの基礎となる値を返します。
// vの基礎となる値がバイトのスライスまたはアドレスのつけられたバイトの配列でない場合、パニックを発生させます。
func (v Value) Bytes() []byte

// CanAddr関数は、値のアドレスを [Value.Addr] で取得できるかどうかを報告します。
// このような値は、addressableと呼ばれます。値がaddressableであるとは、
// スライスの要素、addressableな配列の要素、addressableな構造体のフィールド、
// またはポインタの参照結果である場合を指します。
// CanAddrがfalseを返す場合、 [Value.Addr] を呼び出すとパニックが発生します。
func (v Value) CanAddr() bool

// CanSetはvの値を変更できるかどうかを報告します。
// 値はアドレス可能であり、非公開の構造体フィールドの使用では
// 取得されていない場合にのみ変更できます。
// CanSetがfalseを返す場合、Setや任意のタイプ固有のセッター（例： [Value.SetBool] 、 [Value.SetInt] ）を呼び出すとパニックが発生します。
func (v Value) CanSet() bool

// Callは引数inを使って関数vを呼び出します。
// 例えば、len(in) == 3の場合、v.Call(in)はGoの呼び出しv(in[0], in[1], in[2])を表します。
// CallはvのKindが [Func] でない場合にパニックを発生させます。
// Callは結果をValuesとして返します。
// Goと同様に、各入力引数は関数の対応する入力パラメータの型に代入可能でなければなりません。
// もしvが可変引数関数である場合、Callは対応する値をコピーして可変引数スライスパラメータを作成します。
func (v Value) Call(in []Value) []Value

// CallSliceは可変長関数vを入力引数inで呼び出し、スライスin[len(in)-1]をvの最終可変引数に割り当てます。
// 例えば、len(in) == 3の場合、v.CallSlice(in)はGoの呼び出しv(in[0], in[1], in[2]...)を表します。
// CallSliceはvのKindが [Func] でないか、可変引数でない場合にパニックを引き起こします。
// 出力結果はValuesとして返されます。
// Goと同様に、各入力引数は関数の対応する入力パラメータの型に代入可能でなければなりません。
func (v Value) CallSlice(in []Value) []Value

// Cap は v の容量を返します。
// v の Kind が [Array] 、 [Chan] 、 [Slice] 、または [Array] のポインタでない場合、パニックを発生させます。
func (v Value) Cap() int

// Closeはチャネルvを閉じます。
// vの種類が [Chan] でない場合はパニックを起こします。
func (v Value) Close()

// CanComplexは [Value.Complex] をパニックを起こさずに使用できるかどうかを報告します。
func (v Value) CanComplex() bool

// Complexはvの基礎となる値、つまりcomplex128を返します。
// vのKindが [Complex64] または [Complex128] でない場合、パニックを発生させます。
func (v Value) Complex() complex128

// Elemは、インターフェースvに格納されている値またはvのポインタが指す値を返します。
// vのKindが [Interface] または [Pointer] でない場合、パニックが発生します。
// vがnilの場合、ゼロのValueを返します。
func (v Value) Elem() Value

// Fieldはvのi番目のフィールドを返します。
// vのKindが [Struct] でない場合や、iが範囲外の場合にはパニックを起こします。
func (v Value) Field(i int) Value

// FieldByIndexは、インデックスに対応するネストされたフィールドを返します。
// 評価にはnilポインターを通過する必要があるか、または構造体でないフィールドの場合、パニックします。
func (v Value) FieldByIndex(index []int) Value

// FieldByIndexErrは、インデックスに対応するネストしたフィールドを返します。
// 評価によってnilポインターを進める必要がある場合はエラーを返し、
// structではないフィールドを進める必要がある場合はパニックを引き起こします。
func (v Value) FieldByIndexErr(index []int) (Value, error)

// FieldByNameは指定された名前の構造体フィールドを返します。
// フィールドが見つからない場合はゼロ値を返します。
// vのKindが [Struct] でない場合はパニックを起こします。
func (v Value) FieldByName(name string) Value

// FieldByNameFuncは、マッチ関数を満たす名前を持つ構造体フィールドを返します。
// vのKindが [Struct] でない場合、パニックを起こします。
// フィールドが見つからなかった場合、ゼロの値を返します。
func (v Value) FieldByNameFunc(match func(string) bool) Value

// CanFloat は、 [Value.Float] をパニックせずに使用できるかどうかを報告します。
func (v Value) CanFloat() bool

// Floatはvの基礎値をfloat64として返します。
// もしvの種類が [Float32] または [Float64] でない場合、パニックを起こします。
func (v Value) Float() float64

// Indexはvのi番目の要素を返します。
// vのKindが [Array] 、 [Slice] 、または [String] でない場合、またはiが範囲外の場合はパニックが発生します。
func (v Value) Index(i int) Value

// CanIntはIntをパニックすることなく使用できるかどうかを報告します。
func (v Value) CanInt() bool

// Intはvの基になる値をint64として返します。
// vのKindが [Int] 、 [Int8] 、 [Int16] 、 [Int32] 、または [Int64] でない場合、パニックします。
func (v Value) Int() int64

// CanInterfaceは、パニックなしで [Value.Interface] を使用できるかどうかを報告します。
func (v Value) CanInterface() bool

// Interfaceは、vの現在の値をinterface{}として返します。
// これは以下と同等です：
//
//	var i interface{} = (vの基底値)
//
// もしValueが非公開の構造体フィールドにアクセスして取得された場合はパニックを起こします。
func (v Value) Interface() (i any)

// InterfaceDataは未指定のuintptr値のペアを返します。
// vの種類がInterfaceでない場合、エラーが発生します。
//
// Goの以前のバージョンでは、この関数はインターフェースの値をuintptrのペアで返していました。
// Go 1.4以降、インターフェースの値の実装はInterfaceDataの定義された使用を除外しています。
//
// 廃止予定: インターフェースの値のメモリ表現はInterfaceDataと互換性がありません。
func (v Value) InterfaceData() [2]uintptr

// IsNilは引数vがnilであるかどうかを報告します。引数は
// chan、func、interface、map、pointer、またはsliceの値である必要があります。そうでない場合、IsNilはパニックを引き起こします。注意点として、Go言語での通常のnilとの比較とは常に等しくありません。例えば、vが初期化されていないインターフェース変数iを使用して [ValueOf] を呼び出した場合、i==nilはtrueとなりますが、v.IsNilはvがゼロ値であるためパニックを引き起こします。
func (v Value) IsNil() bool

// IsValidは、vが値を表すかどうかを報告します。
// vがゼロ値の場合はfalseを返します。
// [Value.IsValid] がfalseを返す場合、Stringを除いた他のメソッドはすべてパニックします。
// ほとんどの関数やメソッドは無効な値を返しません。
// 無効な値を返す場合、そのドキュメントは明示的に条件を説明します。
func (v Value) IsValid() bool

// IsZeroは、vが自身の型のゼロ値であるかどうかを報告します。
// 引数が無効な場合、パニックを引き起こします。
func (v Value) IsZero() bool

// SetZeroはvをvの型のゼロ値に設定します。
// [Value.CanSet] がfalseを返す場合はパニックを発生させます。
func (v Value) SetZero()

// KindはvのKindを返します。
// もしvがゼロ値である場合（ [Value.IsValid] がfalseを返す場合）、KindはInvalidを返します。
func (v Value) Kind() Kind

// Lenはvの長さを返します。
// vの種類が [Array] 、 [Chan] 、 [Map] 、 [Slice] 、 [String] 、または [Array] のポインタでない場合、パニックを発生させます。
func (v Value) Len() int

// MapIndex は、マップ v に関連付けられたキーの値を返します。
// v の Kind が [Map] でない場合、パニックが発生します。
// key がマップ内に見つからない場合や v が nil マップを表す場合、ゼロ値が返されます。
// Go と同様に、キーの値はマップのキーの型に代入可能でなければなりません。
func (v Value) MapIndex(key Value) Value

// MapKeysは、マップ内に存在するすべてのキーを含むスライスを返します。
// 順序は指定されていません。
// vのKindが [Map] でない場合、パニックを発生させます。
// vがnilのマップを表す場合、空のスライスを返します。
func (v Value) MapKeys() []Value

// MapIterは、マップを範囲指定するためのイテレータです。
// Value.MapRangeを参照してください。
type MapIter struct {
	m     Value
	hiter hiter
}

// Keyはiterの現在のマップエントリのキーを返します。
func (iter *MapIter) Key() Value

// SetIterKeyはiterの現在のマップエントリのキーをvに割り当てます。
// これはv.Set(iter.Key())と同等ですが、新しいValueを割り当てることを回避します。
// Goと同様に、キーはvの型に割り当て可能である必要があり、
// 非公開フィールドから派生していない必要があります。
func (v Value) SetIterKey(iter *MapIter)

// Valueはiterの現在のマップエントリの値を返します。
func (iter *MapIter) Value() Value

// SetIterValue は iter の現在のマップエントリの値を v に割り当てます。
// v.Set(iter.Value()) と同等ですが、新しい Value を割り当てることを避けます。
// Go と同様に、値は v の型に割り当て可能である必要があり、
// 非公開フィールドから派生していない必要があります。
func (v Value) SetIterValue(iter *MapIter)

// Nextはマップイテレータを進め、別のエントリがあるかどうかを報告します。
// iterが終了した場合はfalseを返します。 [MapIter.Key] 、 [MapIter.Value] 、または [MapIter.Next] への後続の呼び出しはパニックを引き起こします。
func (iter *MapIter) Next() bool

// Reset は iter を v を参照するように変更します。
// もし v の Kind が [Map] ではなく、かつ v がゼロ値でない場合にはパニックを起こします。
// Reset(Value{}) は iter がどのマップも参照しないようにし、
// 以前に繰り返し処理されたマップがガベージコレクションされる可能性があります。
func (iter *MapIter) Reset(v Value)

// MapRangeはマップの範囲イテレータを返します。
// vの種類が [Map] でない場合はパニックを起こします。
//
// イテレータを進めるためには [MapIter.Next] を呼び出し、各エントリにアクセスするためには [MapIter.Key]/[MapIter.Value] を使用します。
// イテレータが使い果たされると [MapIter.Next] はfalseを返します。
// MapRangeはrange文と同じイテレーションのセマンティクスに従います。
//
// 例:
//
//	iter := reflect.ValueOf(m).MapRange()
//	for iter.Next() {
//		k := iter.Key()
//		v := iter.Value()
//		...
//	}
func (v Value) MapRange() *MapIter

// メソッドは、vのi番目のメソッドに対応する関数値を返します。
// 返された関数に対するCallの引数には、レシーバを含めないでください。
// 返された関数は常にvをレシーバとして使用します。
// iが範囲外であるか、vがnilインターフェースの値である場合、Methodはパニックを引き起こします。
func (v Value) Method(i int) Value

// NumMethodは値のメソッドセット内のメソッド数を返します。
//
// インターフェース型では、エクスポートされたメソッドと非エクスポートされたメソッドの数を返します。
// インターフェース以外の型では、エクスポートされたメソッドの数を返します。
func (v Value) NumMethod() int

// MethodByNameは、指定された名前のメソッドに対応する関数値を返します。
// 返された関数に対するCallの引数には、レシーバを含めないでください。返された関数は常にvをレシーバとして使用します。
// メソッドが見つからない場合、ゼロ値を返します。
func (v Value) MethodByName(name string) Value

// NumFieldは構造体vのフィールドの数を返します。
// vのKindが [Struct] でない場合は、パニックが発生します。
func (v Value) NumField() int

// OverflowComplex は complex128 型の x が v の型で表現できないかどうかを報告します。
// もし v の Kind が [Complex64] または [Complex128] でない場合は、パニックを起こします。
func (v Value) OverflowComplex(x complex128) bool

// OverflowFloatは、float64のxがvの型で表現できない場合にtrueを返します。
// vの種類が [Float32] または [Float64] でない場合、パニックを発生させます。
func (v Value) OverflowFloat(x float64) bool

// OverflowInt は、int64型のxがvの型で表現できない場合にtrueを返します。
// vのKindが [Int] 、 [Int8] 、 [Int16] 、 [Int32] 、または [Int64] でない場合はパニックを発生させます。
func (v Value) OverflowInt(x int64) bool

// OverflowUintはuint64 xがvの型で表現できないかどうかを報告します。
// vのKindが [Uint] 、 [Uintptr] 、 [Uint8] 、 [Uint16] 、 [Uint32] 、または [Uint64] でない場合は、panicします。
func (v Value) OverflowUint(x uint64) bool

// Pointerはvの値をuintptrとして返します。
// vの種類が [Chan] 、 [Func] 、 [Map] 、 [Pointer] 、 [Slice] 、[String]、または [UnsafePointer] でない場合はパニックが発生します。
//
// vの種類が [Func] の場合、返されるポインタは基礎となるコードポインタですが、
// 単一の関数を一意に識別するために必要なものではありません。
// ただし、結果がゼロであることは、vがnilのfunc Valueである場合に限ります。
//
// vの種類が [Slice] の場合、返されるポインタはスライスの最初の要素へのポインタです。
// スライスがnilの場合、返される値は0です。
// スライスが空で非nilの場合、返される値は0でない値です。
//
// vのKindが [String] である場合、返されるポインタは文字列の基礎となるバイトの最初の要素を指します。
//
// 同等の結果を得るには、uintptr(Value.UnsafePointer())を使用することが推奨されます。
func (v Value) Pointer() uintptr

// Recv はチャネル v から値を受信して返します。
// v の Kind が [Chan] でない場合、パニックが発生します。
// 受信は値が準備されるまでブロックされます。
// boolean 値 ok は、値 x がチャネル上の送信に対応する場合は true、チャネルが閉じられているために受信したゼロ値の場合は false です。
func (v Value) Recv() (x Value, ok bool)

// Sendはチャネルvにxを送信します。
// vの種類が [Chan] でないか、xの型がvの要素の型と異なる場合、パニックを引き起こします。
// Go言語のように、xの値はチャネルの要素の型に代入可能でなければなりません。
func (v Value) Send(x Value)

// Setはxに値vを割り当てます。
// [Value.CanSet] がfalseを返す場合、パニックを発生させます。
// Go言語と同様に、xの値はvの型に割り当て可能であり、
// 非公開フィールドから派生していない必要があります。
func (v Value) Set(x Value)

// SetBoolはvの基になる値を設定します。
// vのKindが [Bool] でない場合、または [Value.CanSet] がfalseを返す場合はパニックを発生させます。
func (v Value) SetBool(x bool)

// SetBytesはvの基本値を設定します。
// vの基本値がバイトのスライスでない場合、パニックを引き起こします。
func (v Value) SetBytes(x []byte)

// SetComplex は v の基礎値を x に設定します。
// もし v の Kind が [Complex64] や [Complex128] ではない場合、または [Value.CanSet] が false の場合はパニックになります。
func (v Value) SetComplex(x complex128)

// SetFloatはvの基底値をxに設定します。
// vのKindが [Float32] または [Float64] でない場合、または [Value.CanSet] がfalseの場合、パニックが発生します。
func (v Value) SetFloat(x float64)

// SetIntはvの基になる値をxに設定します。
// vのKindが [Int], [Int8], [Int16], [Int32], [Int64] でない場合、または [Value.CanSet] がfalseの場合、パニックとなります。
func (v Value) SetInt(x int64)

// SetLenはvの長さをnに設定します。
// vのKindが [Slice] でない場合や、nが負の値であるか
// スライスの容量よりも大きい場合には、パニックを発生させます。
func (v Value) SetLen(n int)

// SetCapはvの容量をnに設定します。
// vの種類が [Slice] でない場合や、nがスライスの長さより小さく、
// スライスの容量よりも大きい場合は、パニックを引き起こします。
func (v Value) SetCap(n int)

// SetMapIndexはマップv内のキーに関連付けられている要素をelemに設定します。
// vのKindが [Map] でない場合はパニックを発生させます。
// elemがゼロ値の場合、SetMapIndexはマップからキーを削除します。
// さらに、vがnilのマップを保持している場合は、SetMapIndexはパニックを発生させます。
// Go言語と同様に、keyのelemはマップのキーの型に割り当て可能でなければならず、
// elemの値はマップのelemの型に割り当て可能でなければなりません。
func (v Value) SetMapIndex(key, elem Value)

// SetUintはvの基になる値をxに設定します。
// vの種類が [Uint] 、[Uintptr] 、 [Uint8] 、 [Uint16] 、 [Uint32] 、または [Uint64] でない場合、または [Value.CanSet] がfalseの場合はパニックを起こします。
func (v Value) SetUint(x uint64)

// SetPointerは、[unsafe.Pointer]の値であるvをxに設定します。
// vの種類が [UnsafePointer] でない場合、パニックを起こします。
func (v Value) SetPointer(x unsafe.Pointer)

// SetStringはvの基礎となる値をxに設定します。
// vのKindが [String] でない場合、またはCanSet()がfalseの場合はパニックを発生させます。
func (v Value) SetString(x string)

// Sliceはv[i:j]を返します。
// vのKindが [Array] 、 [Slice] 、または [String] でない場合、またはvがアドレス指定できない配列、またはインデックスが範囲外の場合はpanicします。
func (v Value) Slice(i, j int) Value

// Slice3はスライス操作の3つのインデックス形式です：v[i:j:k]を返します。
// もしvの種類が [Array] または [Slice] でない場合、またはvがアドレス不可能な配列である場合、
// もしくはインデックスが範囲外の場合、panicを発生させます。
func (v Value) Slice3(i, j, k int) Value

// Stringは、文字列vの基礎となる値を文字列として返します。
// Stringは、GoのStringメソッドの規約による特別なケースです。
// 他のゲッターと異なり、vのKindが [String] でない場合でもエラーにはなりません。
// 代わりに、"<T value>"という形式の文字列を返します。ここで、Tはvの型です。
// fmtパッケージは、Valueを特別扱いします。暗黙的にStringメソッドを呼び出さず、代わりに保持している具体的な値を表示します。
func (v Value) String() string

// TryRecvはチャネルvから値を受信しようとしますが、ブロックしません。
// vのKindが [Chan] でない場合、パニックが発生します。
// 受信が値を配信する場合、xは転送された値であり、okはtrueです。
// ブロックすることなく受信を完了できない場合、xはゼロ値であり、okはfalseです。
// チャネルが閉じられている場合、xはチャネルの要素型のゼロ値であり、okはfalseです。
func (v Value) TryRecv() (x Value, ok bool)

// TrySend はチャネル v に x を送信しようと試みますが、ブロックしません。
// v の種類が [Chan] でない場合は、パニックを発生させます。
// 値が送信されたかどうかを報告します。
// Go のように、x の値はチャネルの要素型に割り当て可能である必要があります。
func (v Value) TrySend(x Value) bool

// Typeはvの型を返します。
func (v Value) Type() Type

// CanUintは、パニックせずに [Value.Uint] を使用できるかどうかを報告します。
func (v Value) CanUint() bool

// Uintはvの基礎値をuint64として返します。
// vのKindが [Uint] 、 [Uintptr] 、 [Uint8] 、 [Uint16] 、 [Uint32] 、または [Uint64] でない場合にはパニックを発生させます。
func (v Value) Uint() uint64

// UnsafeAddrはvのデータへのポインタを、uintptrとして返します。
// vがアドレス可能でない場合、panicします。
//
// 同等の結果を得るためには、uintptr(Value.Addr().UnsafePointer())を使用することが推奨されます。
func (v Value) UnsafeAddr() uintptr

<<<<<<< HEAD
// UnsafePointerはvの値を [unsafe.Pointer] として返します。
// vのKindが [Chan] 、 [Func] 、 [Map] 、 [Pointer] 、 [Slice] 、または [UnsafePointer] でない場合はパニックを発生させます。
=======
// UnsafePointer returns v's value as a [unsafe.Pointer].
// It panics if v's Kind is not [Chan], [Func], [Map], [Pointer], [Slice], [String] or [UnsafePointer].
>>>>>>> upstream/master
//
// vのKindが [Func] の場合、返されるポインタは基礎となるコードポインタですが、必ずしも単一の関数を一意に識別するためのものではありません。
// 唯一の保証は、vがnil func Valueである場合にのみ結果がゼロであることです。
//
// vのKindが [Slice] の場合、返されるポインタはスライスの最初の要素へのポインタです。スライスがnilの場合、返される値もnilです。
// スライスが空であるが非nilの場合、返される値は非nilです。
//
// vのKindが [String] である場合、返されるポインタは文字列の基礎となるバイトの最初の要素を指します。
func (v Value) UnsafePointer() unsafe.Pointer

// StringHeaderは文字列のランタイム表現です。
// 安全かつ可搬性が保証されておらず、将来のリリースで表現が変わる可能性があります。
// さらに、Dataフィールドだけではデータがガベージコレクションされないことは保証できないため、プログラムは基礎データへの正しい型付きポインタを別途保持する必要があります。
//
// Deprecated: 代わりにunsafe.Stringまたはunsafe.StringDataを使用してください。
type StringHeader struct {
	Data uintptr
	Len  int
}

// SliceHeaderはスライスのランタイム表現です。
// これは安全でも可搬性がありませんし、将来のバージョンで変更されるかもしれません。
// さらに、Dataフィールドだけではデータがガベージコレクトされないことを保証できないため、
// プログラムは基礎データへの正しい型のポインタを別に保持する必要があります。
//
// 廃止予定: 代わりにunsafe.Sliceまたはunsafe.SliceDataを使用してください。
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// Grow increases the slice's capacity, if necessary, to guarantee space for
// another n elements. After Grow(n), at least n elements can be appended
// to the slice without another allocation.
//
// It panics if v's Kind is not a [Slice] or if n is negative or too large to
// allocate the memory.
func (v Value) Grow(n int)

// Clearメソッドは、マップの内容をクリアするか、スライスの内容をゼロにします。
//
// もしvのKindが [Map] または [Slice] でない場合、panicが発生します。
func (v Value) Clear()

// Appendは値xをスライスsに追加し、結果のスライスを返します。
// Goのように、各xの値はスライスの要素の型に割り当て可能でなければなりません。
func Append(s Value, x ...Value) Value

// AppendSliceは、スライスtをスライスsに追加し、結果のスライスを返します。
// スライスsとtは同じ要素の型でなければなりません。
func AppendSlice(s, t Value) Value

// Copyは、dstが満たされるか、srcが使い果たされるまで、srcの内容をdstにコピーします。
// コピーされた要素の数を返します。
// Dstとsrcはそれぞれ [Slice] または [Array] の種類でなければならず、
// dstとsrcは同じ要素の型でなければなりません。
//
// 特別な場合として、dstの要素の種類がUint8である場合、srcの種類はStringであることができます。
func Copy(dst, src Value) int

// SelectDirはセレクト文の通信方向を表します。
type SelectDir int

const (
	_ SelectDir = iota
	SelectSend
	SelectRecv
	SelectDefault
)

// SelectCaseは、select操作内の1つのcaseを表します。
// caseの種類は、Dir（通信の方向）に依存します。
//
// もしDirがSelectDefaultである場合、caseはデフォルトのcaseを表します。
// ChanとSendはゼロ値でなければなりません。
//
// もしDirがSelectSendである場合、caseは送信操作を表します。
// 通常、Chanの基礎となる値はチャネルであり、Sendの基礎となる値はチャネルの要素型に代入可能でなければなりません。
// 特別な場合として、もしChanがゼロ値である場合、そのcaseは無視され、フィールドのSendも無視され、ゼロ値またはゼロ値でないどちらでもかまいません。
//
// もしDirが [SelectRecv] である場合、caseは受信操作を表します。
// 通常、Chanの基礎となる値はチャネルであり、Sendはゼロ値でなければなりません。
// もしChanがゼロ値である場合、そのcaseは無視されますが、Sendはゼロ値でなければなりません。
// 受信操作が選択されると、受信された値はSelectによって返されます。
type SelectCase struct {
	Dir  SelectDir
	Chan Value
	Send Value
}

// Selectは、ケースのリストによって説明されるselect操作を実行します。
// Goのselect文と同様に、少なくとも1つのケースが進行できるまでブロックされ、一様な擬似乱数選択を行い、
// その後、選択されたケースを実行します。選択されたケースのインデックスを返し、
// もしケースが受信操作である場合は、受信した値と、その値がチャネルに送信された値と対応するかどうかを示す
// 真偽値を返します（チャネルがクローズされたためにゼロ値が受信された場合とは異なります）。
// Selectは最大65536のケースをサポートしています。
func Select(cases []SelectCase) (chosen int, recv Value, recvOK bool)

// MakeSliceは指定したスライスの型、長さ、容量の新しいゼロ初期化されたスライス値を作成します。
func MakeSlice(typ Type, len, cap int) Value

// SliceAtは、基礎となるデータがpで始まり、長さと容量がnと等しいスライスを表す [Value] を返します。
//
// これは [unsafe.Slice] のようなものです。
func SliceAt(typ Type, p unsafe.Pointer, n int) Value

// MakeChanは指定された型とバッファサイズで新しいチャネルを作成します。
func MakeChan(typ Type, buffer int) Value

// MakeMapは指定された型の新しいマップを作成します。
func MakeMap(typ Type) Value

// MakeMapWithSizeは、指定された型とおおよそのn個の要素のための初期空間を持つ新しいマップを作成します。
func MakeMapWithSize(typ Type, n int) Value

// Indirectは、vが指す値を返します。
// vがnilポインターの場合、Indirectはゼロ値を返します。
// vがポインターでない場合、Indirectはvを返します。
func Indirect(v Value) Value

// ValueOfはインターフェースiに格納された具体的な値で初期化された新しいValueを返します。ValueOf(nil)はゼロのValueを返します。
func ValueOf(i any) Value

// Zeroは指定された型のゼロ値を表すValueを返します。
// 結果はValue構造体のゼロ値とは異なり、値が存在しないことを表します。
// 例えば、Zero(TypeOf(42))はKindが [Int] で値が0のValueを返します。
// 返された値はアドレスを取ることも変更することもできません。
func Zero(typ Type) Value

// Newは指定された型の新しいゼロ値へのポインタを表すValueを返します。
// つまり、返されたValueのTypeは [PointerTo](typ)です。
func New(typ Type) Value

// NewAtは、pを指し示すポインタを使用して、指定された型の値へのポインタを表すValueを返します。
func NewAt(typ Type, p unsafe.Pointer) Value

// Convert は値 v を型 t に変換した値を返します。
// もし通常の Go の変換ルールによって値 v を型 t に変換することができない場合、または、v を型 t に変換する際にパニックが発生する場合、Convert はパニックを発生させます。
func (v Value) Convert(t Type) Value

// CanConvertは、値vが型tに変換可能かどうかを報告します。
// v.CanConvert(t)がtrueを返す場合、v.Convert(t)はパニックしません。
func (v Value) CanConvert(t Type) bool

// Comparableは値vが比較可能かどうかを報告します。
// もしvの型がインターフェースである場合、これは動的な型をチェックします。
// もしこれがtrueを報告する場合、v.Interface() == x はどんなxに対してもパニックを起こしませんし、
// v.Equal(u) もどんなValue uに対してもパニックを起こしません。
func (v Value) Comparable() bool

// Equalは、vがuと等しい場合にtrueを返します。
// 2つの無効な値に対して、Equalはtrueを返します。
// インターフェース値の場合、Equalはインターフェース内の値を比較します。
// それ以外の場合、値の型が異なる場合はfalseを返します。
// また、配列や構造体の場合、Equalは順番に各要素を比較し、
// 等しくない要素が見つかった場合にfalseを返します。
// すべての比較中、同じ型の値が比較され、その型が比較できない場合、Equalはパニックを引き起こします。
func (v Value) Equal(u Value) bool
