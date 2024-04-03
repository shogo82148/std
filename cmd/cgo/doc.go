// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Cgoを使用すると、Cコードを呼び出すGoパッケージを作成できます。

# goコマンドでcgoを使用する

cgoを使用するには、擬似パッケージ「C」をインポートする通常のGoコードを書きます。
その後、GoコードはC.size_tのような型、C.stdoutのような変数、またはC.putcharのような関数を参照できます。

「C」のインポートが直前にコメントで指定されている場合、そのコメントは前文と呼ばれ、
パッケージのC部分をコンパイルするときにヘッダーとして使用されます。 例えば：

	// #include <stdio.h>
	// #include <errno.h>
	import "C"

プリアンブルには、関数や変数の宣言や定義を含むCコードを含めることができます。
これらは、パッケージ「C」に定義されているかのように、Goコードから参照できます。
プリアンブルで宣言されたすべての名前は使用できますが、小文字で始まる場合でも使用できます。
例外：プリアンブルの静的変数は、Goコードから参照できません。静的関数は許可されています。

$GOROOT/cmd/cgo/internal/teststdioと$GOROOT/misc/cgo/gmpを参照して、例を確認してください。
cgoの使用についての紹介については、「C？Go？Cgo！」を参照してください：
https://golang.org/doc/articles/c_go_cgo.html 。

CFLAGS、CPPFLAGS、CXXFLAGS、FFLAGS、およびLDFLAGSは、これらのコメント内の疑似#cgoディレクティブで定義できます。
これにより、C、C ++、またはFortranコンパイラの動作を調整できます。
複数のディレクティブで定義された値は連結されます。
ディレクティブには、その効果を制限するビルド制約のリストを含めることができます。
（制約構文の詳細については、https://golang.org/pkg/go/build/#hdr-Build_Constraints を参照してください。）
例えば：

	// #cgo CFLAGS: -DPNG_DEBUG=1
	// #cgo amd64 386 CFLAGS: -DX86=1
	// #cgo LDFLAGS: -lpng
	// #include <png.h>
	import "C"

代わりに、'#cgo pkg-config:'ディレクティブに続いてパッケージ名を指定することで、
pkg-configツールを使用してCPPFLAGSとLDFLAGSを取得できます。
例えば：

	// #cgo pkg-config: png cairo
	// #include <png.h>
	import "C"

デフォルトのpkg-configツールは、PKG_CONFIG環境変数を設定することで変更できます。

セキュリティ上の理由から、許可されるフラグは-D、-U、-I、および-lのみです。
追加のフラグを許可するには、CGO_CFLAGS_ALLOWを正規表現に設定し、新しいフラグに一致するようにします。
許可されるフラグを禁止するには、CGO_CFLAGS_DISALLOWを正規表現に設定し、禁止する必要がある引数に一致するようにします。
両方の場合、正規表現は完全な引数に一致する必要があります。-mfoo=barを許可するには、CGO_CFLAGS_ALLOW='-mfoo.*'を使用します。
単にCGO_CFLAGS_ALLOW='-mfoo'ではなく、同様の名前の変数が許可されるCPPFLAGS、CXXFLAGS、FFLAGS、およびLDFLAGSを制御します。

また、セキュリティ上の理由から、英数字文字といくつかの記号（例えば、'.'など）のみが許可されており、
予期しない方法で解釈されないいくつかの記号が許可されています。
禁止された文字を使用すると、「malformed #cgo argument」エラーが発生します。

ビルド時に、CGO_CFLAGS、CGO_CPPFLAGS、CGO_CXXFLAGS、CGO_FFLAGS、およびCGO_LDFLAGS環境変数は、
これらのディレクティブから派生したフラグに追加されます。
パッケージ固有のフラグは、環境変数ではなくディレクティブを使用して設定する必要があります。
これにより、変更されていない環境でビルドが機能するようになります。
環境変数から取得したフラグは、上記で説明したセキュリティ上の制限の対象外です。

パッケージ内のすべてのcgo CPPFLAGSおよびCFLAGSディレクティブは連結され、
そのパッケージのCファイルをコンパイルするために使用されます。
パッケージ内のすべてのCPPFLAGSおよびCXXFLAGSディレクティブは連結され、
そのパッケージのC++ファイルをコンパイルするために使用されます。
パッケージ内のすべてのCPPFLAGSおよびFFLAGSディレクティブは連結され、
そのパッケージのFortranファイルをコンパイルするために使用されます。
プログラム内の任意のパッケージのすべてのLDFLAGSディレクティブは連結され、リンク時に使用されます。
すべてのpkg-configディレクティブは連結され、同時にpkg-configに送信され、
各適切なコマンドラインフラグのセットに追加されます。

cgoディレクティブが解析されると、文字列${SRCDIR}の出現箇所は、ソースファイルを含むディレクトリの絶対パスに置き換えられます。
これにより、事前にコンパイルされた静的ライブラリをパッケージディレクトリに含め、適切にリンクすることができます。
例えば、パッケージfooがディレクトリ/go/src/fooにある場合：

	// #cgo LDFLAGS: -L${SRCDIR}/libs -lfoo

次のように展開されます：

	// #cgo LDFLAGS: -L/go/src/foo/libs -lfoo

Goツールが1つ以上のGoファイルで特別なインポート「C」を使用することを検出すると、
ディレクトリ内の他の非Goファイルを検索して、Goパッケージの一部としてコンパイルします。
.c、.s、.S、または.sxファイルはCコンパイラでコンパイルされます。
.cc、.cpp、または.cxxファイルはC ++コンパイラでコンパイルされます。
.f、.F、.for、または.f90ファイルはFortranコンパイラでコンパイルされます。
.h、.hh、.hpp、または.hxxファイルは別途コンパイルされませんが、
これらのヘッダーファイルが変更された場合、パッケージ（およびその非Goソースファイル）が再コンパイルされます。
他のディレクトリのファイルを変更しても、パッケージが再コンパイルされるわけではないため、
パッケージのすべての非Goソースコードはサブディレクトリではなく、パッケージディレクトリに保存する必要があります。
デフォルトのCおよびC++コンパイラは、それぞれCCおよびCXX環境変数によって変更できます。
これらの環境変数には、コマンドラインオプションを含めることができます。

cgoツールは常に、ソースファイルのディレクトリを含むパスでCコンパイラを呼び出します。
つまり、-I${SRCDIR}が常に暗示されます。これは、ヘッダーファイルfoo/bar.hが、
ソースディレクトリとシステムインクルードディレクトリ（または-Iフラグで指定された他の場所）の両方に存在する場合、
「#include <foo/bar.h>」が常に他のバージョンよりも優先してローカルバージョンを見つけることを意味します。

cgoツールは、動作が期待されるシステムでのネイティブビルドではデフォルトで有効になっています。
クロスコンパイル時や、CC環境変数が設定されていない場合、およびデフォルトのCコンパイラ（通常はgccまたはclang）がシステムPATHに見つからない場合は、デフォルトで無効になります。
goツールを実行する際にCGO_ENABLED環境変数を設定することで、デフォルトを上書きできます。
cgoの使用を有効にするには、1に設定し、無効にするには0に設定します。
goツールは、cgoが有効になっている場合、ビルド制約「cgo」を設定します。
特別なインポート「C」は、「//go:build cgo」と同じように「cgo」ビルド制約を意味するため、
cgoが無効になっている場合、Cをインポートするファイルはgoツールによってビルドされません。
（ビルド制約の詳細については、https://golang.org/pkg/go/build/#hdr-Build_Constraints を参照してください。）

クロスコンパイル時には、cgoが使用するCクロスコンパイラを指定する必要があります。
これは、make.bashを使用してツールチェーンをビルドする際に、一般的なCC_FOR_TARGETまたは
より具体的なCC_FOR_${GOOS}_${GOARCH}（例：CC_FOR_linux_arm）環境変数を設定することで行うことができます。
または、goツールを実行する際にいつでもCC環境変数を設定することができます。

C++コードに対しては、CXX_FOR_TARGET、CXX_FOR_${GOOS}_${GOARCH}、およびCXX環境変数が同様の方法で機能します。

# GoからCへの参照

Goファイル内で、GoのキーワードであるCの構造体フィールド名には、アンダースコアを前置することでアクセスできます。
たとえば、xが「type」という名前のフィールドを持つC構造体を指す場合、x._typeでフィールドにアクセスできます。
ビットフィールドやアラインメントされていないデータなど、Goで表現できないC構造体フィールドは、
次のフィールドまたは構造体の末尾に到達するまで、適切なパディングに置き換えられてGo構造体から省略されます。

標準のC数値型は、以下の名前で利用可能です。
C.char、C.schar（signed char）、C.uchar（unsigned char）、
C.short、C.ushort（unsigned short）、C.int、C.uint（unsigned int）、
C.long、C.ulong（unsigned long）、C.longlong（long long）、
C.ulonglong（unsigned long long）、C.float、C.double、
C.complexfloat（complex float）、およびC.complexdouble（complex double）。
C型void *は、Goのunsafe.Pointerで表されます。
C型__int128_tおよび__uint128_tは、[16]byteで表されます。

Goで通常はポインタ型で表されるいくつかの特別なC型は、代わりにuintptrで表されます。
詳細については、以下の特別なケースのセクションを参照してください。

構造体、共用体、または列挙型に直接アクセスするには、
C.struct_statのように、struct_、union_、またはenum_を接頭辞として付けます。

C型Tのサイズは、C.sizeof_struct_statのように、C.sizeof_Tとして利用可能です。

Goファイルで、特別な名前_GoString_のパラメータ型を持つC関数を宣言することができます。
この関数は、通常のGo文字列値で呼び出すことができます。
文字列の長さと、文字列の内容へのポインタは、C関数を呼び出すことでアクセスできます。

	size_t _GoStringLen(_GoString_ s);
	const char *_GoStringPtr(_GoString_ s);

これらの関数は、他のCファイルではなく、プリアンブルでのみ使用できます。
Cコードは、_GoStringPtrによって返されるポインタの内容を変更してはいけません。
文字列の内容には、末尾のNULバイトがない場合があることに注意してください。

一般的な場合、GoはCの共用体型をサポートしていないため、
Cの共用体型は同じ長さのGoバイト配列として表されます。

Goの構造体には、Cの型を埋め込むことはできません。

Goのコードは、非空のC構造体の末尾にあるサイズゼロのフィールドを参照することはできません。
（サイズゼロのフィールドでできる唯一の操作である）そのようなフィールドのアドレスを取得するには、
構造体のアドレスを取得し、構造体のサイズを加算する必要があります。

Cgoは、Cの型を同等の非公開のGoの型に変換します。
翻訳が非公開であるため、GoパッケージはCの型をエクスポートされたAPIで公開すべきではありません。
1つのGoパッケージで使用されるCの型は、別のGoパッケージで使用される同じCの型とは異なります。

任意のC関数（void関数でも）は、複数の代入コンテキストで呼び出すことができ、
戻り値（ある場合）とC errno変数をエラーとして取得できます（関数がvoidを返す場合は、結果値をスキップするために_を使用します）。
例えば：

	n, err = C.sqrt(-1)
	_, err := C.voidFunc()
	var n, err = C.sqrt(1)

Cの関数ポインタを呼び出すことは現在サポートされていませんが、
Cの関数ポインタを保持するGo変数を宣言し、GoとCの間で相互に渡すことができます。
Cコードは、Goから受け取った関数ポインタを呼び出すことができます。
例えば：

	package main

	// typedef int (*intFunc) ();
	//
	// int
	// bridge_int_func(intFunc f)
	// {
	//		return f();
	// }
	//
	// int fortytwo()
	// {
	//	    return 42;
	// }
	import "C"
	import "fmt"

	func main() {
		f := C.intFunc(C.fortytwo)
		fmt.Println(int(C.bridge_int_func(f)))
		// Output: 42
	}

C言語では、固定サイズの配列として書かれた関数引数は、実際には配列の最初の要素へのポインタを必要とします。
Cコンパイラはこの呼び出し規約を認識して、呼び出しを適切に調整しますが、Goはできません。
Goでは、最初の要素へのポインタを明示的に渡す必要があります：C.f(&C.x[0])。

可変長のC関数を呼び出すことはサポートされていません。C関数ラッパーを使用することで、これを回避することができます。
例えば：

	package main

	// #include <stdio.h>
	// #include <stdlib.h>
	//
	// static void myprint(char* s) {
	//   printf("%s\n", s);
	// }
	import "C"
	import "unsafe"

	func main() {
		cs := C.CString("Hello from stdio")
		C.myprint(cs)
		C.free(unsafe.Pointer(cs))
	}

いくつかの特別な関数は、データのコピーを作成することによって、
GoとCの型の間を変換します。疑似Go定義では次のようになります。

	// Go文字列からC文字列へ
	// C文字列はmallocを使用してCヒープに割り当てられます。
	// 解放する責任は呼び出し側にあります。たとえば、C.freeを呼び出すことで解放できます
	// （C.freeが必要な場合はstdlib.hを含めることを忘れないでください）。
	func C.CString(string) *C.char

	// Goの[]byteスライスからCの配列へ
	// Cの配列はmallocを使用してCヒープに割り当てられます。
	// 解放する責任は呼び出し側にあります。たとえば、C.freeを呼び出すことで解放できます
	// （C.freeが必要な場合はstdlib.hを含めることを忘れないでください）。
	func C.CBytes([]byte) unsafe.Pointer

	// C文字列からGo文字列へ
	func C.GoString(*C.char) string

	// 明示的な長さを持つCデータからGo文字列へ
	func C.GoStringN(*C.char, C.int) string

	// 明示的な長さを持つCデータからGoの[]byteスライスへ
	func C.GoBytes(unsafe.Pointer, C.int) []byte

// 特別な場合として、C.mallocはCライブラリmallocを直接呼び出すのではなく、
// CライブラリmallocをラップするGoヘルパー関数を呼び出しますが、
// nilを返さないことを保証します。Cのmallocがメモリ不足を示す場合、
// ヘルパー関数はプログラムをクラッシュさせます。Go自体がメモリ不足になった場合と同様です。
// C.mallocは失敗しないため、errnoを返す2つの結果形式はありません。

// CからGoへの参照
//
// Go関数は、Cコードで使用するために次の方法でエクスポートできます。

	//export MyFunction
	func MyFunction(arg1, arg2 int, arg3 string) int64 {...}

	//export MyFunction2
	func MyFunction2(arg1, arg2 int, arg3 string) (int64, *C.char) {...}

Cコードでは、次のように使用できます。

	extern GoInt64 MyFunction(int arg1, int arg2, GoString arg3);
	extern struct MyFunction2_return MyFunction2(int arg1, int arg2, GoString arg3);

// _cgo_export.h生成ヘッダーで見つかった、前文からコピーされたプリアンブルの後にある
// 複数の戻り値を持つ関数は、構造体を返す関数にマップされます。

すべてのGoの型が有用な方法でCの型にマップできるわけではありません。
Goの構造体型はサポートされていません。Cの構造体型を使用してください。
Goの配列型はサポートされていません。Cのポインタ型を使用してください。

Goのstring型の引数を取る関数は、上記で説明した _GoString_ 型のCタイプで呼び出すことができます。
_GoString_ 型は自動的にプリアンブルで定義されます。
Cコードからこの型の値を作成する方法はありません。これは、GoからCに文字列値を渡し、
再びGoに戻すためにのみ有用です。

ファイルで//exportを使用すると、プリアンブルに制限が設けられます。
2つの異なるC出力ファイルにコピーされるため、定義ではなく宣言のみを含める必要があります。
ファイルに定義と宣言の両方が含まれている場合、2つの出力ファイルは重複するシンボルを生成し、
リンカーが失敗します。これを回避するには、定義を他のファイルのプリアンブルまたはCソースファイルに配置する必要があります。

# ポインタの渡し方

Goはガベージコレクションされる言語であり、ガベージコレクタはGoメモリへのすべてのポインタの場所を知る必要があります。
そのため、GoとCの間でポインタを渡す際には制限があります。

このセクションでは、Goポインタという用語は、Goによって割り当てられたメモリへのポインタを意味します
（＆演算子を使用するか、定義済みのnew関数を呼び出すことによって）。
Cポインタという用語は、Cによって割り当てられたメモリへのポインタを意味します（C.mallocを呼び出すことによって）。
ポインタがGoポインタであるかCポインタであるかは、メモリがどのように割り当てられたかによって動的に決定されます。
ポインタの型とは何の関係もありません。

いくつかのGoの型の値は、型のゼロ値以外に常にGoポインタを含みます。
これは、string、slice、interface、channel、map、およびfunction型に当てはまります。
ポインタ型はGoポインタまたはCポインタを保持できます。
配列と構造体型は、要素型によってGoポインタを含む場合と含まない場合があります。
以下のすべてのGoポインタに関する議論は、ポインタ型だけでなく、
Goポインタを含む他の型にも適用されます。

Cに渡されるすべてのGoポインタは、固定されたGoメモリを指す必要があります。
C関数に関数引数として渡されるGoポインタは、
呼び出しの期間中暗黙的に固定されたメモリを指します。
これらの関数引数から到達可能なGoメモリは、Cコードがアクセスできる限り固定されている必要があります。
Goメモリが固定されているかどうかは、そのメモリ領域の動的なプロパティであり、
ポインタの型とは何の関係もありません。

newを呼び出すことによって、複合リテラルのアドレスを取得することによって、またはローカル変数のアドレスを取得することによって作成されたGo値は、
[runtime.Pinner] を使用してメモリを固定することもできます。この型は、メモリの固定状態の期間を管理するために使用でき、
C関数呼び出しの期間を超えてメモリを固定することができます。メモリは複数回固定でき、
固定された回数と同じ回数だけ固定を解除する必要があります。

Goのコードは、ポイント先のメモリにGoポインタが含まれていない場合、
GoポインタをCに渡すことができます。構造体のフィールドにポインタを渡す場合、
問題のGoメモリはフィールドによって占有されるメモリであり、構造体全体ではありません。
配列またはスライスの要素にポインタを渡す場合、問題のGoメモリは
配列全体またはスライスのバッキング配列全体です。

Cのコードは、Goポインタが指すメモリが固定されている限り、Goポインタのコピーを保持できます。

Cのコードは、呼び出しが返された後、Goポインタのコピーを保持することはできません。
ただし、ポインタが指すメモリが [runtime.Pinner] で固定され、PinnerがGoポインタがCメモリに格納されている間にアンピンされない場合は、
Cのメモリに格納されたGoポインタを保持できます。
これは、Cコードが文字列、スライス、チャネルなどのコピーを保持できないことを意味します。
これらは [runtime.Pinner] で固定できないためです。

_GoString_ 型も [runtime.Pinner] で固定することはできません。
Goポインタを含むため、指すメモリは呼び出しの期間だけ固定されます。
_GoString_ 値はCコードによって保持されることはできません。

Cコードによって呼び出されるGo関数は、
ピン留めされたメモリへのGoポインタを返すことができます
（これは、文字列、スライス、チャネルなどを返すことはできないことを意味します）。
Cコードによって呼び出されるGo関数は、Cポインタを引数として取ることができ、
それらのポインタを介して非ポインタデータ、Cポインタ、またはピン留めされたGoポインタを格納することができます。
Goポインタを指すメモリにGoポインタを格納することはできません。
（これは、文字列、スライス、チャネルなどを格納することはできないことを意味します）。
Cコードによって呼び出されるGo関数は、Goポインタを取ることができますが、
指すGoメモリ（およびそのメモリが指すGoメモリなど）がピン留めされていることを保証する必要があります。

<<<<<<< HEAD
これらのルールは、実行時に動的にチェックされます。チェックは、GODEBUG環境変数のcgocheck設定によって制御されます。
デフォルトの設定はGODEBUG=cgocheck=1で、比較的安価な動的チェックが実装されています。
これらのチェックは、GODEBUG=cgocheck=0を使用して完全に無効にすることができます。
ポインタの処理の完全なチェックは、実行時間のコストがかかりますが、GODEBUG=cgocheck=2を使用して利用できます。
=======
These rules are checked dynamically at runtime. The checking is
controlled by the cgocheck setting of the GODEBUG environment
variable. The default setting is GODEBUG=cgocheck=1, which implements
reasonably cheap dynamic checks. These checks may be disabled
entirely using GODEBUG=cgocheck=0. Complete checking of pointer
handling, at some cost in run time, is available by setting
GOEXPERIMENT=cgocheck2 at build time.
>>>>>>> upstream/master

unsafeパッケージを使用することで、この強制を無効にすることができます。
もちろん、Cコードが好きなことをすることを防ぐものは何もありません。
ただし、これらのルールを破るプログラムは、予期しない方法で失敗する可能性があります。

runtime/cgo.Handle型は、GoとCの間で安全にGo値を渡すために使用できます。
詳細については、runtime/cgoパッケージのドキュメントを参照してください。

注：現在の実装にはバグがあります。GoコードはCメモリにnilまたはCポインタ（ただしGoポインタではない）を書き込むことが許可されていますが、
現在の実装では、Cメモリの内容がGoポインタであるように見える場合には、ランタイムエラーが発生することがあります。
したがって、Goコードがその中にポインタ値を格納する場合は、初期化されていないCメモリをGoコードに渡すことを避けてください。
Cでメモリをゼロにしてから渡してください。

# 特殊なケース

通常、Goではポインタ型で表されるいくつかの特殊なC型は、代わりにuintptrで表されます。それらには以下が含まれます:

1. Darwinの*Ref型、CoreFoundationのCFTypeRef型に根ざしています。

2. JavaのJNIインターフェースからのオブジェクト型:

	jobject
	jclass
	jthrowable
	jstring
	jarray
	jbooleanArray
	jbyteArray
	jcharArray
	jshortArray
	jintArray
	jlongArray
	jfloatArray
	jdoubleArray
	jobjectArray
	jweak

3. EGL APIからのEGLDisplayとEGLConfigのタイプ。

これらのタイプは、Go側ではuintptrであるため、Goガベージコレクターが混乱する可能性があるためです。これらは、時には本当にポインタではなく、ポインタタイプにエンコードされたデータ構造です。これらのタイプのすべての操作はCで実行する必要があります。空のこのような参照を初期化するための適切な定数は0であり、nilではありません。

これらの特別なケースは、Go 1.10で導入されました。Go 1.9以前からの自動更新コードには、Go fixツールのcftypeまたはjniの書き換えを使用してください。

	go tool fix -r cftype <pkg>
	go tool fix -r jni <pkg>

適切な場所でnilを0に置き換えます。

EGLDisplayの場合は、Go 1.12で導入されました。Go 1.11以前のコードから自動更新するには、eglの書き換えを使用してください。

	go tool fix -r egl <pkg>

EGLConfigの場合は、Go 1.15で導入されました。Go 1.14以前のコードから自動更新するには、eglconfの書き換えを使用してください。

	go tool fix -r eglconf <pkg>

# cgoを直接使用する

使用法:

	go tool cgo [cgoオプション] [-- コンパイラオプション] goファイル...

Cgoは、指定された入力Goソースファイルを複数の出力GoおよびCソースファイルに変換します。

パッケージのC部分をコンパイルするためにCコンパイラを呼び出す際に、コンパイラオプションは解釈されずにそのまま渡されます。

次のオプションは、cgoを直接実行する場合に使用できます。

	-V
		cgoのバージョンを表示して終了します。
	-debug-define
		デバッグオプション。#defineを出力します。
	-debug-gcc
		デバッグオプション。Cコンパイラの実行と出力をトレースします。
	-dynimport file
		fileがインポートするシンボルのリストを書き込みます。-dynout引数または標準出力に書き込みます。cgoパッケージをビルドするときにgo buildによって使用されます。
	-dynlinker
		-dynimport出力の一部として動的リンカーを書き込みます。
	-dynout file
		-dynimport出力をfileに書き込みます。
	-dynpackage package
		-dynimport出力のためのGoパッケージを設定します。
	-exportheader file
		エクスポートされた関数がある場合、生成されたエクスポート宣言をファイルに書き込みます。
		Cコードはこれを#includeして宣言を見ることができます。
	-importpath string
		Goパッケージのインポートパス。オプションです。生成されたファイルのコメントに使用されます。
	-import_runtime_cgo
		設定されている場合（デフォルトで設定されています）、生成された出力でruntime/cgoをインポートします。
	-import_syscall
		設定されている場合（デフォルトで設定されています）、生成された出力でsyscallをインポートします。
	-gccgo
		gcコンパイラではなく、gccgoコンパイラ向けの出力を生成します。
	-gccgoprefix prefix
		gccgoで使用する-fgo-prefixオプション。
	-gccgopkgpath path
		gccgoで使用する-fgo-pkgpathオプション。
	-gccgo_define_cgoincomplete
		古いgccgoバージョン用に、cgo.Incompleteをインポートする代わりにローカルで定義します。
		"runtime/cgo"パッケージから。
	-godefs
		Cパッケージ名を実際の値に置き換えたGo構文で入力ファイルを書き出します。
		新しいターゲットをブートストラップするときにsyscallパッケージ内のファイルを生成するために使用されます。
	-objdir directory
		すべての生成されたファイルをディレクトリに配置します。
	-srcdir directory
*/
package main
