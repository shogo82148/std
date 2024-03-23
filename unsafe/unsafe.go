// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package unsafeは、Go言語のプログラムの型の安全性を無視する操作を含んでいます。

unsafeをインポートするパッケージは非可搬性があり、Go 1の互換性ガイドラインで保護されていません。
*/package unsafe

// ArbitraryTypeはドキュメントの目的でここにあり、実際にはunsafeパッケージの一部ではありません。
// これは任意のGo式の型を表します。
type ArbitraryType int

// IntegerTypeはドキュメンテーションのために存在しますが、実際にはunsafeパッケージの一部ではありません。
// 任意の整数型を表します。
type IntegerType int

// Pointerは任意の型へのポインタを表します。Pointerのために他の型では利用できない
// 4つの特別な操作が利用できます：
//   - 任意の型のポインタ値をPointerに変換することができます。
//   - Pointerを任意の型のポインタ値に変換することができます。
//   - uintptrをPointerに変換することができます。
//   - Pointerをuintptrに変換することができます。
//
// Pointerはしたがって、プログラムが型システムを無視して任意のメモリを読み書きすることを可能にします。これは非常に注意して使用する必要があります。
//
// 以下のPointerに関連するパターンは有効です。
// これらのパターンを使用しないコードは、今日では無効であり、将来では無効になる可能性があります。
// 以下の有効なパターンにも重要な注意事項が付属しています。
//
// "go vet"を実行すると、これらのパターンに適合しないPointerの使用箇所を見つけるのに役立ちますが、
// "go vet"からのサイレンスはコードが有効であることを保証するものではありません。
//
// (1) *T1をPointerから*T2に変換する。
//
// T2がT1より大きくなく、両者が等価なメモリレイアウトを共有している場合、
// この変換により、ある型のデータを別の型のデータとして再解釈することができます。例えば、
// math.Float64bitsの実装です：
//
//	func Float64bits(f float64) uint64 {
//		return *(*uint64)(unsafe.Pointer(&f))
//	}
//
// (2) Pointerをuintptrに変換するが、Pointerには戻すことはできません。
//
// Pointerをuintptrに変換すると、指された値のメモリアドレスを整数として取得します。一般的な使用法は、
// それを印字することです。
//
// uintptrをPointerに戻す変換は一般的には有効ではありません。
//
// uintptrは整数であり、参照ではないです。
// Pointerをuintptrに変換すると、ポインタのセマンティクスを持たない整数値が作成されます。
// uintptrがあるオブジェクトのアドレスを保持していても、
// ガベージコレクタはそのuintptrの値を更新することはありませんし、そのuintptrはそのオブジェクトの再利用を防ぎません。
//
// 残りのパターンは、uintptrからPointerへの唯一の有効な変換を列挙しています。
//
// (3) Pointerをuintptrに変換して算術演算を実行し、Pointerに戻す。
//
// pが割り当てられたオブジェクトを指している場合、uintptrに変換して、オフセットを加算し、Pointerに戻すことでそのオブジェクト内を進むことができます。
//
//	p = unsafe.Pointer(uintptr(p) + offset)
//
// このパターンの最も一般的な使用法は、structのフィールドにアクセスしたり、配列の要素にアクセスすることです：
//
//	// f := unsafe.Pointer(&s.f)と同等
//	f := unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Offsetof(s.f))
//
//	// e := unsafe.Pointer(&x[i])と同等
//	e := unsafe.Pointer(uintptr(unsafe.Pointer(&x[0])) + i*unsafe.Sizeof(x[0]))
//
// この方法ではポインタにオフセットを追加することと、ポインタを減算することの両方が有効です。
// アライメントのためにポインタを整数に丸めるために、&^演算子を使用することも有効です。
// すべての場合において、結果は元の割り当てられたオブジェクトを指し続けなければなりません。
//
// Cとは異なり、ポインタを元の割り当てられたスペースのすぐ外側に進めるのは無効です：
//
//	// 無効：endは割り当てられたスペースの外側を指しています。
//	var s thing
//	end = unsafe.Pointer(uintptr(unsafe.Pointer(&s)) + unsafe.Sizeof(s))
//
//	// 無効：endは割り当てられたスペースの外側を指しています。
//	b := make([]byte, n)
//	end = unsafe.Pointer(uintptr(unsafe.Pointer(&b[0])) + uintptr(n))
//
// 両方の変換は同じ式内に出現しなければならず、それらの間には算術演算のみが許されます：
//
//	// 無効：uintptrは変数に保存できません
//	// Pointerに戻す前の変換が行われるまで
//	u := uintptr(p)
//	p = unsafe.Pointer(u + offset)
//
// ポインタが割り当てられたオブジェクトを指す必要があるため、nilではないことに注意してください。
//
//	// 無効：nilのポインタの変換
//	u := unsafe.Pointer(nil)
//	p := unsafe.Pointer(uintptr(u) + offset)
//
<<<<<<< HEAD
// (4) syscall.Syscallを呼び出す際にPointerをuintptrに変換する場合。
=======
// (4) Conversion of a Pointer to a uintptr when calling [syscall.Syscall].
>>>>>>> upstream/master
//
// パッケージsyscallのSyscall関数は、uintptrの引数を直接オペレーティングシステムに渡し、
// その後、呼び出しの詳細によっては、一部の引数をポインタとして再解釈する場合があります。
// つまり、システムコールの実装が暗黙的に特定の引数をuintptrからポインタに変換しています。
//
// ポインタ引数をuintptrに変換して使用する必要がある場合は、その変換は呼び出し式自体に表示される必要があります：
//
//	syscall.Syscall(SYS_READ, uintptr(fd), uintptr(unsafe.Pointer(p)), uintptr(n))
//
// コンパイラは、アセンブリで実装された関数への呼び出しの引数リスト内で、Pointerをuintptrに変換するように処理します。
// これにより、参照されている割り当てられたオブジェクトが、呼び出しが完了するまで維持され、移動されないようになりますが、
// 型のみから判断すると、呼び出し中にオブジェクトが必要なくなっているように見えます。
//
// コンパイラがこのパターンを認識するためには、変換は引数リストに表示される必要があります：
//
//	// 無効：uintptrは変数に保存できません
//	// システムコール中にPointerへの暗黙の変換前
//	u := uintptr(unsafe.Pointer(p))
//	syscall.Syscall(SYS_READ, uintptr(fd), u, uintptr(n))
//
<<<<<<< HEAD
// (5) reflect.Value.Pointerやreflect.Value.UnsafeAddrの結果をuintptrからPointerに変換する場合。
=======
// (5) Conversion of the result of [reflect.Value.Pointer] or [reflect.Value.UnsafeAddr]
// from uintptr to Pointer.
>>>>>>> upstream/master
//
// パッケージreflectのValueのPointerとUnsafeAddrという名前のメソッドは、結果をunsafe.Pointerではなくuintptr型として返すため、
// "unsafe"を最初にインポートせずに結果を任意の型に変更することを防いでいます。しかし、これは結果が壊れやすく、
// 呼び出しの直後にすぐにPointerに変換する必要があることを意味します。
//
//	p := (*int)(unsafe.Pointer(reflect.ValueOf(new(int)).Pointer()))
//
// 上記の場合と同様に、変換前に結果を保存することは無効です：
//
//	// 無効：uintptrは変数に保存できません
//	// Pointerに戻す前の変換が行われるまで
//	u := reflect.ValueOf(new(int)).Pointer()
//	p := (*int)(unsafe.Pointer(u))
//
<<<<<<< HEAD
// (6) reflect.SliceHeaderまたはreflect.StringHeaderのDataフィールドをPointerに変換するか、あるいはその逆。
=======
// (6) Conversion of a [reflect.SliceHeader] or [reflect.StringHeader] Data field to or from Pointer.
>>>>>>> upstream/master
//
// 前のケースと同様に、reflectデータ構造のSliceHeaderとStringHeaderは、
// Dataフィールドをuintptrとして宣言していますが、任意の型に結果を変更することを防いでいます。
//
<<<<<<< HEAD
=======
//	var s string
//	hdr := (*reflect.StringHeader)(unsafe.Pointer(&s)) // case 1
//	hdr.Data = uintptr(unsafe.Pointer(p))              // case 6 (this case)
//	hdr.Len = n
//
// In this usage hdr.Data is really an alternate way to refer to the underlying
// pointer in the string header, not a uintptr variable itself.
//
// In general, [reflect.SliceHeader] and [reflect.StringHeader] should be used
// only as *reflect.SliceHeader and *reflect.StringHeader pointing at actual
// slices or strings, never as plain structs.
// A program should not declare or allocate variables of these struct types.
//
//	// INVALID: a directly-declared header will not hold Data as a reference.
//	var hdr reflect.StringHeader
//	hdr.Data = uintptr(unsafe.Pointer(p))
//	hdr.Len = n
//	s := *(*string)(unsafe.Pointer(&hdr)) // p possibly already lost
>>>>>>> upstream/master
type Pointer *ArbitraryType

// Sizeofは、任意の型の式xを取り、仮想的な変数vがvar v = xとして宣言された場合のサイズ（バイト単位）を返します。
// サイズには、xによって参照される可能性のあるメモリは含まれません。
// たとえば、xがスライスである場合、Sizeofはスライスのディスクリプタのサイズを返しますが、スライスが参照するメモリのサイズではありません。
// 構造体の場合、サイズにはフィールドのアライメントによって導入されるパディングも含まれます。
// Sizeofの返り値は、引数xの型が可変サイズではない場合にはGoの定数です。
// （型が可変サイズである場合、それは型パラメータであるか、可変サイズの要素を持つ配列または構造体型です）。
func Sizeof(x ArbitraryType) uintptr

// Offsetofは、xによって表されるフィールドの構造体内のオフセットを返します。
// ここでxは、structValue.fieldの形式である必要があります。
// 言い換えると、構造体の先頭とフィールドの先頭の間のバイト数を返します。
// Offsetofの戻り値は、引数xの型が可変サイズでない場合、Goの定数です。
// （可変サイズの型の定義については、[Sizeof]の説明を参照してください。）
func Offsetof(x ArbitraryType) uintptr

<<<<<<< HEAD
// Alignofは、任意のタイプの式xを取り、仮想の変数vがvar v = xとして宣言された場合の必要なアライメントを返します。
// vのアドレスが常に0 mod mであるような最大の値mです。
// これは、reflect.TypeOf(x).Align()が返す値と同じです。
// 特殊なケースとして、変数sがstruct型であり、fがそのstruct内のフィールドである場合、Alignof(s.f)は、struct内のその型のフィールドの必要なアライメントを返します。
// このケースは、reflect.TypeOf(s.f).FieldAlign()が返す値と同じです。
// Alignofの戻り値は、引数のタイプが可変サイズではない場合、Goの定数です。
// （可変サイズの型の定義については、[Sizeof]の説明を参照してください。)
func Alignof(x ArbitraryType) uintptr

// 関数Addはlenをptrに加算し、更新されたポインタ
// Pointer(uintptr(ptr) + uintptr(len)) を返します。
// len引数は整数型または無型定数である必要があります。
// 定数のlen引数はint型の値で表現可能でなければなりません。
// もし無型定数である場合は、int型として扱われます。
// Pointerの有効な使用法に関するルールは変わりません。
=======
// Alignof takes an expression x of any type and returns the required alignment
// of a hypothetical variable v as if v was declared via var v = x.
// It is the largest value m such that the address of v is always zero mod m.
// It is the same as the value returned by [reflect.TypeOf](x).Align().
// As a special case, if a variable s is of struct type and f is a field
// within that struct, then Alignof(s.f) will return the required alignment
// of a field of that type within a struct. This case is the same as the
// value returned by [reflect.TypeOf](s.f).FieldAlign().
// The return value of Alignof is a Go constant if the type of the argument
// does not have variable size.
// (See the description of [Sizeof] for a definition of variable sized types.)
func Alignof(x ArbitraryType) uintptr

// The function Add adds len to ptr and returns the updated pointer
// [Pointer](uintptr(ptr) + uintptr(len)).
// The len argument must be of integer type or an untyped constant.
// A constant len argument must be representable by a value of type int;
// if it is an untyped constant it is given type int.
// The rules for valid uses of Pointer still apply.
>>>>>>> upstream/master
func Add(ptr Pointer, len IntegerType) Pointer

// 関数Sliceは、ポインタptrで指定された配列の先頭から長さと容量がlenであるスライスを返します。
// Slice(ptr, len)は次のように表されます：
//
//	(*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
//
// ただし、特別なケースとして、ptrがnilであり、かつlenがゼロの場合は、Sliceはnilを返します。
//
// len引数は整数型または未指定の定数である必要があります。
// 定数のlen引数は非負であり、int型の値で表現可能でなければならず、
// 未指定の定数の場合はint型として与えられます。
// 実行時に、lenが負の値であるか、ptrがnilでlenがゼロでない場合、
// 実行時パニックが発生します。
func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType

// SliceDataは引数スライスの基本配列へのポインタを返します。
//   - もしcap(slice) > 0であれば、SliceDataは&slice[:1][0]を返します。
//   - もしslice == nilであれば、SliceDataはnilを返します。
//   - それ以外の場合、SliceDataは特定のメモリアドレスへの非nilなポインタを返します。
func SliceData(slice []ArbitraryType) *ArbitraryType

// Stringは下層バイトがptrで始まり、長さがlenである文字列値を返します。
//
// lenの引数は整数型または未指定の定数でなければなりません。
// 定数のlenの引数は非負であり、int型の値として表現可能でなければなりません;
// もし未指定の定数である場合はint型として指定されます。
// 実行時に、lenが負であるか、またはptrがnilであり、かつlenがゼロでない場合には、
// 実行時パニックが発生します。
//
// Goの文字列は不変であるため、Stringに渡されたバイトはその後変更してはなりません。
func String(ptr *byte, len IntegerType) string

// StringDataは、strの基礎バイトへのポインタを返します。
// 空の文字列の場合、返り値は特定されず、nilである場合があります。
//
// Goの文字列は不変のため、StringDataによって返されたバイトは変更してはいけません。
func StringData(str string) *byte
