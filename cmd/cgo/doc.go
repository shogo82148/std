// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
<<<<<<< HEAD
Cgo enables the creation of Go packages that call C code.

# Using cgo with the go command

To use cgo write normal Go code that imports a pseudo-package "C".
The Go code can then refer to types such as C.size_t, variables such
as C.stdout, or functions such as C.putchar.

If the import of "C" is immediately preceded by a comment, that
comment, called the preamble, is used as a header when compiling
the C parts of the package. For example:
=======
Cgoを使用すると、Cコードを呼び出すGoパッケージを作成できます。

# goコマンドでcgoを使用する

cgoを使用するには、擬似パッケージ「C」をインポートする通常のGoコードを書きます。
その後、GoコードはC.size_tのような型、C.stdoutのような変数、またはC.putcharのような関数を参照できます。

「C」のインポートが直前にコメントで指定されている場合、そのコメントは前文と呼ばれ、
パッケージのC部分をコンパイルするときにヘッダーとして使用されます。 例えば：
>>>>>>> release-branch.go1.21

	// #include <stdio.h>
	// #include <errno.h>
	import "C"

<<<<<<< HEAD
The preamble may contain any C code, including function and variable
declarations and definitions. These may then be referred to from Go
code as though they were defined in the package "C". All names
declared in the preamble may be used, even if they start with a
lower-case letter. Exception: static variables in the preamble may
not be referenced from Go code; static functions are permitted.

See $GOROOT/cmd/cgo/internal/teststdio and $GOROOT/misc/cgo/gmp for examples. See
"C? Go? Cgo!" for an introduction to using cgo:
https://golang.org/doc/articles/c_go_cgo.html.

CFLAGS, CPPFLAGS, CXXFLAGS, FFLAGS and LDFLAGS may be defined with pseudo
#cgo directives within these comments to tweak the behavior of the C, C++
or Fortran compiler. Values defined in multiple directives are concatenated
together. The directive can include a list of build constraints limiting its
effect to systems satisfying one of the constraints
(see https://golang.org/pkg/go/build/#hdr-Build_Constraints for details about the constraint syntax).
For example:
=======
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
>>>>>>> release-branch.go1.21

	// #cgo CFLAGS: -DPNG_DEBUG=1
	// #cgo amd64 386 CFLAGS: -DX86=1
	// #cgo LDFLAGS: -lpng
	// #include <png.h>
	import "C"

<<<<<<< HEAD
Alternatively, CPPFLAGS and LDFLAGS may be obtained via the pkg-config tool
using a '#cgo pkg-config:' directive followed by the package names.
For example:
=======
代わりに、'#cgo pkg-config:'ディレクティブに続いてパッケージ名を指定することで、
pkg-configツールを使用してCPPFLAGSとLDFLAGSを取得できます。
例えば：
>>>>>>> release-branch.go1.21

	// #cgo pkg-config: png cairo
	// #include <png.h>
	import "C"

<<<<<<< HEAD
The default pkg-config tool may be changed by setting the PKG_CONFIG environment variable.

For security reasons, only a limited set of flags are allowed, notably -D, -U, -I, and -l.
To allow additional flags, set CGO_CFLAGS_ALLOW to a regular expression
matching the new flags. To disallow flags that would otherwise be allowed,
set CGO_CFLAGS_DISALLOW to a regular expression matching arguments
that must be disallowed. In both cases the regular expression must match
a full argument: to allow -mfoo=bar, use CGO_CFLAGS_ALLOW='-mfoo.*',
not just CGO_CFLAGS_ALLOW='-mfoo'. Similarly named variables control
the allowed CPPFLAGS, CXXFLAGS, FFLAGS, and LDFLAGS.

Also for security reasons, only a limited set of characters are
permitted, notably alphanumeric characters and a few symbols, such as
'.', that will not be interpreted in unexpected ways. Attempts to use
forbidden characters will get a "malformed #cgo argument" error.

When building, the CGO_CFLAGS, CGO_CPPFLAGS, CGO_CXXFLAGS, CGO_FFLAGS and
CGO_LDFLAGS environment variables are added to the flags derived from
these directives. Package-specific flags should be set using the
directives, not the environment variables, so that builds work in
unmodified environments. Flags obtained from environment variables
are not subject to the security limitations described above.

All the cgo CPPFLAGS and CFLAGS directives in a package are concatenated and
used to compile C files in that package. All the CPPFLAGS and CXXFLAGS
directives in a package are concatenated and used to compile C++ files in that
package. All the CPPFLAGS and FFLAGS directives in a package are concatenated
and used to compile Fortran files in that package. All the LDFLAGS directives
in any package in the program are concatenated and used at link time. All the
pkg-config directives are concatenated and sent to pkg-config simultaneously
to add to each appropriate set of command-line flags.

When the cgo directives are parsed, any occurrence of the string ${SRCDIR}
will be replaced by the absolute path to the directory containing the source
file. This allows pre-compiled static libraries to be included in the package
directory and linked properly.
For example if package foo is in the directory /go/src/foo:

	// #cgo LDFLAGS: -L${SRCDIR}/libs -lfoo

Will be expanded to:

	// #cgo LDFLAGS: -L/go/src/foo/libs -lfoo

When the Go tool sees that one or more Go files use the special import
"C", it will look for other non-Go files in the directory and compile
them as part of the Go package. Any .c, .s, .S or .sx files will be
compiled with the C compiler. Any .cc, .cpp, or .cxx files will be
compiled with the C++ compiler. Any .f, .F, .for or .f90 files will be
compiled with the fortran compiler. Any .h, .hh, .hpp, or .hxx files will
not be compiled separately, but, if these header files are changed,
the package (including its non-Go source files) will be recompiled.
Note that changes to files in other directories do not cause the package
to be recompiled, so all non-Go source code for the package should be
stored in the package directory, not in subdirectories.
The default C and C++ compilers may be changed by the CC and CXX
environment variables, respectively; those environment variables
may include command line options.

The cgo tool will always invoke the C compiler with the source file's
directory in the include path; i.e. -I${SRCDIR} is always implied. This
means that if a header file foo/bar.h exists both in the source
directory and also in the system include directory (or some other place
specified by a -I flag), then "#include <foo/bar.h>" will always find the
local version in preference to any other version.

The cgo tool is enabled by default for native builds on systems where
it is expected to work. It is disabled by default when cross-compiling
as well as when the CC environment variable is unset and the default
C compiler (typically gcc or clang) cannot be found on the system PATH.
You can override the default by setting the CGO_ENABLED
environment variable when running the go tool: set it to 1 to enable
the use of cgo, and to 0 to disable it. The go tool will set the
build constraint "cgo" if cgo is enabled. The special import "C"
implies the "cgo" build constraint, as though the file also said
"//go:build cgo".  Therefore, if cgo is disabled, files that import
"C" will not be built by the go tool. (For more about build constraints
see https://golang.org/pkg/go/build/#hdr-Build_Constraints).

When cross-compiling, you must specify a C cross-compiler for cgo to
use. You can do this by setting the generic CC_FOR_TARGET or the
more specific CC_FOR_${GOOS}_${GOARCH} (for example, CC_FOR_linux_arm)
environment variable when building the toolchain using make.bash,
or you can set the CC environment variable any time you run the go tool.

The CXX_FOR_TARGET, CXX_FOR_${GOOS}_${GOARCH}, and CXX
environment variables work in a similar way for C++ code.

# Go references to C

Within the Go file, C's struct field names that are keywords in Go
can be accessed by prefixing them with an underscore: if x points at a C
struct with a field named "type", x._type accesses the field.
C struct fields that cannot be expressed in Go, such as bit fields
or misaligned data, are omitted in the Go struct, replaced by
appropriate padding to reach the next field or the end of the struct.

The standard C numeric types are available under the names
C.char, C.schar (signed char), C.uchar (unsigned char),
C.short, C.ushort (unsigned short), C.int, C.uint (unsigned int),
C.long, C.ulong (unsigned long), C.longlong (long long),
C.ulonglong (unsigned long long), C.float, C.double,
C.complexfloat (complex float), and C.complexdouble (complex double).
The C type void* is represented by Go's unsafe.Pointer.
The C types __int128_t and __uint128_t are represented by [16]byte.

A few special C types which would normally be represented by a pointer
type in Go are instead represented by a uintptr.  See the Special
cases section below.

To access a struct, union, or enum type directly, prefix it with
struct_, union_, or enum_, as in C.struct_stat.

The size of any C type T is available as C.sizeof_T, as in
C.sizeof_struct_stat.

A C function may be declared in the Go file with a parameter type of
the special name _GoString_. This function may be called with an
ordinary Go string value. The string length, and a pointer to the
string contents, may be accessed by calling the C functions
=======
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
>>>>>>> release-branch.go1.21

	size_t _GoStringLen(_GoString_ s);
	const char *_GoStringPtr(_GoString_ s);

<<<<<<< HEAD
These functions are only available in the preamble, not in other C
files. The C code must not modify the contents of the pointer returned
by _GoStringPtr. Note that the string contents may not have a trailing
NUL byte.

As Go doesn't have support for C's union type in the general case,
C's union types are represented as a Go byte array with the same length.

Go structs cannot embed fields with C types.

Go code cannot refer to zero-sized fields that occur at the end of
non-empty C structs. To get the address of such a field (which is the
only operation you can do with a zero-sized field) you must take the
address of the struct and add the size of the struct.

Cgo translates C types into equivalent unexported Go types.
Because the translations are unexported, a Go package should not
expose C types in its exported API: a C type used in one Go package
is different from the same C type used in another.

Any C function (even void functions) may be called in a multiple
assignment context to retrieve both the return value (if any) and the
C errno variable as an error (use _ to skip the result value if the
function returns void). For example:
=======
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
>>>>>>> release-branch.go1.21

	n, err = C.sqrt(-1)
	_, err := C.voidFunc()
	var n, err = C.sqrt(1)

<<<<<<< HEAD
Calling C function pointers is currently not supported, however you can
declare Go variables which hold C function pointers and pass them
back and forth between Go and C. C code may call function pointers
received from Go. For example:
=======
Cの関数ポインタを呼び出すことは現在サポートされていませんが、
Cの関数ポインタを保持するGo変数を宣言し、GoとCの間で相互に渡すことができます。
Cコードは、Goから受け取った関数ポインタを呼び出すことができます。
例えば：
>>>>>>> release-branch.go1.21

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

<<<<<<< HEAD
In C, a function argument written as a fixed size array
actually requires a pointer to the first element of the array.
C compilers are aware of this calling convention and adjust
the call accordingly, but Go cannot. In Go, you must pass
the pointer to the first element explicitly: C.f(&C.x[0]).

Calling variadic C functions is not supported. It is possible to
circumvent this by using a C function wrapper. For example:
=======
C言語では、固定サイズの配列として書かれた関数引数は、実際には配列の最初の要素へのポインタを必要とします。
Cコンパイラはこの呼び出し規約を認識して、呼び出しを適切に調整しますが、Goはできません。
Goでは、最初の要素へのポインタを明示的に渡す必要があります：C.f(&C.x[0])。

可変長のC関数を呼び出すことはサポートされていません。C関数ラッパーを使用することで、これを回避することができます。
例えば：
>>>>>>> release-branch.go1.21

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

<<<<<<< HEAD
A few special functions convert between Go and C types
by making copies of the data. In pseudo-Go definitions:

	// Go string to C string
	// The C string is allocated in the C heap using malloc.
	// It is the caller's responsibility to arrange for it to be
	// freed, such as by calling C.free (be sure to include stdlib.h
	// if C.free is needed).
	func C.CString(string) *C.char

	// Go []byte slice to C array
	// The C array is allocated in the C heap using malloc.
	// It is the caller's responsibility to arrange for it to be
	// freed, such as by calling C.free (be sure to include stdlib.h
	// if C.free is needed).
	func C.CBytes([]byte) unsafe.Pointer

	// C string to Go string
	func C.GoString(*C.char) string

	// C data with explicit length to Go string
	func C.GoStringN(*C.char, C.int) string

	// C data with explicit length to Go []byte
	func C.GoBytes(unsafe.Pointer, C.int) []byte

As a special case, C.malloc does not call the C library malloc directly
but instead calls a Go helper function that wraps the C library malloc
but guarantees never to return nil. If C's malloc indicates out of memory,
the helper function crashes the program, like when Go itself runs out
of memory. Because C.malloc cannot fail, it has no two-result form
that returns errno.

# C references to Go

Go functions can be exported for use by C code in the following way:
=======
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
>>>>>>> release-branch.go1.21

	//export MyFunction
	func MyFunction(arg1, arg2 int, arg3 string) int64 {...}

	//export MyFunction2
	func MyFunction2(arg1, arg2 int, arg3 string) (int64, *C.char) {...}

<<<<<<< HEAD
They will be available in the C code as:
=======
Cコードでは、次のように使用できます。
>>>>>>> release-branch.go1.21

	extern GoInt64 MyFunction(int arg1, int arg2, GoString arg3);
	extern struct MyFunction2_return MyFunction2(int arg1, int arg2, GoString arg3);

<<<<<<< HEAD
found in the _cgo_export.h generated header, after any preambles
copied from the cgo input files. Functions with multiple
return values are mapped to functions returning a struct.

Not all Go types can be mapped to C types in a useful way.
Go struct types are not supported; use a C struct type.
Go array types are not supported; use a C pointer.

Go functions that take arguments of type string may be called with the
C type _GoString_, described above. The _GoString_ type will be
automatically defined in the preamble. Note that there is no way for C
code to create a value of this type; this is only useful for passing
string values from Go to C and back to Go.

Using //export in a file places a restriction on the preamble:
since it is copied into two different C output files, it must not
contain any definitions, only declarations. If a file contains both
definitions and declarations, then the two output files will produce
duplicate symbols and the linker will fail. To avoid this, definitions
must be placed in preambles in other files, or in C source files.

# Passing pointers

Go is a garbage collected language, and the garbage collector needs to
know the location of every pointer to Go memory. Because of this,
there are restrictions on passing pointers between Go and C.

In this section the term Go pointer means a pointer to memory
allocated by Go (such as by using the & operator or calling the
predefined new function) and the term C pointer means a pointer to
memory allocated by C (such as by a call to C.malloc). Whether a
pointer is a Go pointer or a C pointer is a dynamic property
determined by how the memory was allocated; it has nothing to do with
the type of the pointer.

Note that values of some Go types, other than the type's zero value,
always include Go pointers. This is true of string, slice, interface,
channel, map, and function types. A pointer type may hold a Go pointer
or a C pointer. Array and struct types may or may not include Go
pointers, depending on the element types. All the discussion below
about Go pointers applies not just to pointer types, but also to other
types that include Go pointers.

All Go pointers passed to C must point to pinned Go memory. Go pointers
passed as function arguments to C functions have the memory they point to
implicitly pinned for the duration of the call. Go memory reachable from
these function arguments must be pinned as long as the C code has access
to it. Whether Go memory is pinned is a dynamic property of that memory
region; it has nothing to do with the type of the pointer.

Go values created by calling new, by taking the address of a composite
literal, or by taking the address of a local variable may also have their
memory pinned using [runtime.Pinner]. This type may be used to manage
the duration of the memory's pinned status, potentially beyond the
duration of a C function call. Memory may be pinned more than once and
must be unpinned exactly the same number of times it has been pinned.

Go code may pass a Go pointer to C provided the memory to which it
points does not contain any Go pointers to memory that is unpinned. When
passing a pointer to a field in a struct, the Go memory in question is
the memory occupied by the field, not the entire struct. When passing a
pointer to an element in an array or slice, the Go memory in question is
the entire array or the entire backing array of the slice.

C code may keep a copy of a Go pointer only as long as the memory it
points to is pinned.

C code may not keep a copy of a Go pointer after the call returns,
unless the memory it points to is pinned with [runtime.Pinner] and the
Pinner is not unpinned while the Go pointer is stored in C memory.
This implies that C code may not keep a copy of a string, slice,
channel, and so forth, because they cannot be pinned with
[runtime.Pinner].

The _GoString_ type also may not be pinned with [runtime.Pinner].
Because it includes a Go pointer, the memory it points to is only pinned
for the duration of the call; _GoString_ values may not be retained by C
code.

A Go function called by C code may return a Go pointer to pinned memory
(which implies that it may not return a string, slice, channel, and so
forth). A Go function called by C code may take C pointers as arguments,
and it may store non-pointer data, C pointers, or Go pointers to pinned
memory through those pointers. It may not store a Go pointer to unpinned
memory in memory pointed to by a C pointer (which again, implies that it
may not store a string, slice, channel, and so forth). A Go function
called by C code may take a Go pointer but it must preserve the property
that the Go memory to which it points (and the Go memory to which that
memory points, and so on) is pinned.

These rules are checked dynamically at runtime. The checking is
controlled by the cgocheck setting of the GODEBUG environment
variable. The default setting is GODEBUG=cgocheck=1, which implements
reasonably cheap dynamic checks. These checks may be disabled
entirely using GODEBUG=cgocheck=0. Complete checking of pointer
handling, at some cost in run time, is available via GODEBUG=cgocheck=2.

It is possible to defeat this enforcement by using the unsafe package,
and of course there is nothing stopping the C code from doing anything
it likes. However, programs that break these rules are likely to fail
in unexpected and unpredictable ways.

The runtime/cgo.Handle type can be used to safely pass Go values
between Go and C. See the runtime/cgo package documentation for details.

Note: the current implementation has a bug. While Go code is permitted
to write nil or a C pointer (but not a Go pointer) to C memory, the
current implementation may sometimes cause a runtime error if the
contents of the C memory appear to be a Go pointer. Therefore, avoid
passing uninitialized C memory to Go code if the Go code is going to
store pointer values in it. Zero out the memory in C before passing it
to Go.

# Optimizing calls of C code

When passing a Go pointer to a C function the compiler normally ensures
that the Go object lives on the heap. If the C function does not keep
a copy of the Go pointer, and never passes the Go pointer back to Go code,
then this is unnecessary. The #cgo noescape directive may be used to tell
the compiler that no Go pointers escape via the named C function.
If the noescape directive is used and the C function does not handle the
pointer safely, the program may crash or see memory corruption.

For example:

	// #cgo noescape cFunctionName

When a Go function calls a C function, it prepares for the C function to
call back to a Go function. the #cgo nocallback directive may be used to
tell the compiler that these preparations are not necessary.
If the nocallback directive is used and the C function does call back into
Go code, the program will panic.

For example:

	// #cgo nocallback cFunctionName

# Special cases

A few special C types which would normally be represented by a pointer
type in Go are instead represented by a uintptr. Those include:

1. The *Ref types on Darwin, rooted at CoreFoundation's CFTypeRef type.

2. The object types from Java's JNI interface:
=======
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

これらのルールは、実行時に動的にチェックされます。チェックは、GODEBUG環境変数のcgocheck設定によって制御されます。
デフォルトの設定はGODEBUG=cgocheck=1で、比較的安価な動的チェックが実装されています。
これらのチェックは、GODEBUG=cgocheck=0を使用して完全に無効にすることができます。
ポインタの処理の完全なチェックは、実行時間のコストがかかりますが、GODEBUG=cgocheck=2を使用して利用できます。

unsafeパッケージを使用することで、この強制を無効にすることができます。
もちろん、Cコードが好きなことをすることを防ぐものは何もありません。
ただし、これらのルールを破るプログラムは、予期しない方法で失敗する可能性があります。

runtime/cgo.Handle型は、GoとCの間で安全にGo値を渡すために使用できます。
詳細については、runtime/cgoパッケージのドキュメントを参照してください。

注：現在の実装にはバグがあります。GoコードはCメモリにnilまたはCポインタ（ただしGoポインタではない）を書き込むことが許可されていますが、
現在の実装では、Cメモリの内容がGoポインタであるように見える場合には、ランタイムエラーが発生することがあります。
したがって、Goコードがその中にポインタ値を格納する場合は、初期化されていないCメモリをGoコードに渡すことを避けてください。
Cでメモリをゼロにしてから渡してください。

# 特別なケース

Goでは通常、ポインタ型で表されるいくつかの特別なC型は、代わりにuintptrで表されます。
これらには、次のものが含まれます。

1. Darwin上の*Ref型は、CoreFoundationのCFTypeRef型をルートとしています。

2. JavaのJNIインターフェースからのオブジェクトタイプ：
>>>>>>> release-branch.go1.21

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

<<<<<<< HEAD
3. The EGLDisplay and EGLConfig types from the EGL API.

These types are uintptr on the Go side because they would otherwise
confuse the Go garbage collector; they are sometimes not really
pointers but data structures encoded in a pointer type. All operations
on these types must happen in C. The proper constant to initialize an
empty such reference is 0, not nil.

These special cases were introduced in Go 1.10. For auto-updating code
from Go 1.9 and earlier, use the cftype or jni rewrites in the Go fix tool:
=======
3. EGL APIからのEGLDisplayとEGLConfigのタイプ。

これらのタイプは、Go側ではuintptrであるため、Goガベージコレクターが混乱する可能性があるためです。これらは、時には本当にポインタではなく、ポインタタイプにエンコードされたデータ構造です。これらのタイプのすべての操作はCで実行する必要があります。空のこのような参照を初期化するための適切な定数は0であり、nilではありません。

これらの特別なケースは、Go 1.10で導入されました。Go 1.9以前からの自動更新コードには、Go fixツールのcftypeまたはjniの書き換えを使用してください。
>>>>>>> release-branch.go1.21

	go tool fix -r cftype <pkg>
	go tool fix -r jni <pkg>

<<<<<<< HEAD
It will replace nil with 0 in the appropriate places.

The EGLDisplay case was introduced in Go 1.12. Use the egl rewrite
to auto-update code from Go 1.11 and earlier:

	go tool fix -r egl <pkg>

The EGLConfig case was introduced in Go 1.15. Use the eglconf rewrite
to auto-update code from Go 1.14 and earlier:

	go tool fix -r eglconf <pkg>

# Using cgo directly

Usage:

	go tool cgo [cgo options] [-- compiler options] gofiles...

Cgo transforms the specified input Go source files into several output
Go and C source files.

The compiler options are passed through uninterpreted when
invoking the C compiler to compile the C parts of the package.

The following options are available when running cgo directly:

	-V
		Print cgo version and exit.
	-debug-define
		Debugging option. Print #defines.
	-debug-gcc
		Debugging option. Trace C compiler execution and output.
	-dynimport file
		Write list of symbols imported by file. Write to
		-dynout argument or to standard output. Used by go
		build when building a cgo package.
	-dynlinker
		Write dynamic linker as part of -dynimport output.
	-dynout file
		Write -dynimport output to file.
	-dynpackage package
		Set Go package for -dynimport output.
	-exportheader file
		If there are any exported functions, write the
		generated export declarations to file.
		C code can #include this to see the declarations.
	-importpath string
		The import path for the Go package. Optional; used for
		nicer comments in the generated files.
	-import_runtime_cgo
		If set (which it is by default) import runtime/cgo in
		generated output.
	-import_syscall
		If set (which it is by default) import syscall in
		generated output.
	-gccgo
		Generate output for the gccgo compiler rather than the
		gc compiler.
	-gccgoprefix prefix
		The -fgo-prefix option to be used with gccgo.
	-gccgopkgpath path
		The -fgo-pkgpath option to be used with gccgo.
	-gccgo_define_cgoincomplete
		Define cgo.Incomplete locally rather than importing it from
		the "runtime/cgo" package. Used for old gccgo versions.
	-godefs
		Write out input file in Go syntax replacing C package
		names with real values. Used to generate files in the
		syscall package when bootstrapping a new target.
	-objdir directory
		Put all generated files in directory.
=======
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
>>>>>>> release-branch.go1.21
	-srcdir directory
*/
package main
