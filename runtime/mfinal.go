// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Garbage collector: finalizers and block profiling.

package runtime

// SetFinalizerは、objに関連付けられたファイナライザを提供されたファイナライザ関数に設定します。
// ガベージコレクタが関連付けられたファイナライザを持つ到達不能なブロックを見つけると、
// 関連付けをクリアし、別のゴルーチンでfinalizer(obj)を実行します。
// これにより、objは再び到達可能になりますが、関連付けられたファイナライザはなくなります。
// SetFinalizerが再度呼び出されない限り、次にガベージコレクタがobjが到達不能であることを検出した場合、objは解放されます。
//
// SetFinalizer(obj, nil)は、objに関連付けられたファイナライザをクリアします。
//
// 引数objは、newを呼び出すことによって割り当てられたオブジェクトへのポインタ、
// 複合リテラルのアドレスを取得することによって、またはローカル変数のアドレスを取得することによって割り当てられたオブジェクトへのポインタである必要があります。
// 引数finalizerは、objの型に割り当てることができる単一の引数を取る関数であり、任意の無視される戻り値を持つことができます。
// これらのいずれかがtrueでない場合、SetFinalizerはプログラムを中止する可能性があります。
//
// ファイナライザは依存関係の順序で実行されます。
// AがBを指し示し、両方にファイナライザがあり、それらが到達不能である場合、Aのファイナライザのみが実行されます。
// Aが解放された後、Bのファイナライザが実行されます。
// ファイナライザを持つブロックを含む循環構造がある場合、その循環はガベージコレクトされることは保証されず、
// 依存関係を尊重する順序がないため、ファイナライザが実行されることも保証されません。
//
// ファイナライザは、プログラムがobjが指し示すオブジェクトに到達できなくなった後、
// 任意の時点で実行されるようにスケジュールされます。
// プログラムが終了する前にファイナライザが実行されることは保証されていないため、
// 通常、長時間実行されるプログラム中にオブジェクトに関連付けられた非メモリリソースを解放するためにのみ有用です。
// たとえば、[os.File] オブジェクトは、プログラムがCloseを呼び出さずにos.Fileを破棄するときに、
// 関連するオペレーティングシステムのファイルディスクリプタを閉じるためにファイナライザを使用できますが、
// bufio.WriterのようなインメモリI/Oバッファをフラッシュするためにファイナライザに依存することは誤りです。
// なぜなら、プログラムが終了するときにバッファがフラッシュされないためです。
//
// *objのサイズがゼロバイトの場合、ファイナライザが実行されることは保証されません。
// なぜなら、メモリ内の他のゼロサイズのオブジェクトと同じアドレスを共有する可能性があるためです。
// 詳細については、https://go.dev/ref/spec#Size_and_alignment_guarantees を参照してください。
//
// パッケージレベルの変数の初期化子で割り当てられたオブジェクトに対して、
// ファイナライザが実行されることは保証されません。
// このようなオブジェクトはヒープに割り当てられるのではなく、リンカによって割り当てられる可能性があります。
//
// ファイナライザがオブジェクトが参照されなくなってから任意の時間が経過した後に実行される可能性があるため、
// ランタイムは、オブジェクトを単一の割り当てスロットにまとめるスペース節約の最適化を実行できます。
// そのような割り当て内の参照されなくなったオブジェクトのファイナライザは、常に参照されたオブジェクトと同じバッチに存在する場合、実行されない可能性があります。
// 通常、このバッチ処理は、小さな（16バイト以下の）ポインタフリーオブジェクトに対してのみ行われます。
//
<<<<<<< HEAD
// オブジェクトが到達不能になるとすぐにファイナライザが実行される場合があります。
// ファイナライザを正しく使用するには、プログラムはオブジェクトが不要になるまで到達可能であることを保証する必要があります。
// グローバル変数に格納されたオブジェクト、またはグローバル変数からポインタをトレースできるオブジェクトは到達可能です。
// その他のオブジェクトについては、[KeepAlive] 関数の呼び出しにオブジェクトを渡して、
// オブジェクトが到達可能である必要がある関数内の最後のポイントをマークする必要があります。
=======
// A finalizer may run as soon as an object becomes unreachable.
// In order to use finalizers correctly, the program must ensure that
// the object is reachable until it is no longer required.
// Objects stored in global variables, or that can be found by tracing
// pointers from a global variable, are reachable. A function argument or
// receiver may become unreachable at the last point where the function
// mentions it. To make an unreachable object reachable, pass the object
// to a call of the [KeepAlive] function to mark the last point in the
// function where the object must be reachable.
>>>>>>> 41b4a7d0008e48dd077e189fd86911de2b36d90d
//
// たとえば、pがファイルディスクリプタdを含むos.Fileのような構造体を指し示す場合、
// pにはdを閉じるファイナライザがあり、pの最後の使用がsyscall.Write(p.d、buf、size)の呼び出しである場合、
// プログラムがsyscall.Writeに入るとすぐにpが到達不能になる可能性があります。
// その瞬間にファイナライザが実行され、p.dを閉じ、syscall.Writeが閉じられたファイルディスクリプタ（または、
// より悪い場合、別のgoroutineによって開かれた完全に異なるファイルディスクリプタ）に書き込もうとして失敗する可能性があります。
// この問題を回避するには、syscall.Writeの呼び出し後にKeepAlive(p)を呼び出します。
//
// プログラムのすべてのファイナライザを、1つのgoroutineが順次実行します。
// ファイナライザが長時間実行する必要がある場合は、新しいgoroutineを開始することで実行する必要があります。
//
// Goのメモリモデルの用語で、SetFinalizer(x、f)の呼び出しは、
// ファイナライザ呼び出しf(x)の前に「同期」します。
// ただし、KeepAlive(x)またはxの他の使用がf(x)の前に「同期」されることは保証されていないため、
// 一般的には、ファイナライザがxの可変状態にアクセスする必要がある場合は、ミューテックスまたは他の同期メカニズムを使用する必要があります。
// たとえば、x内の時折変更される可変フィールドを検査するファイナライザを考えてみましょう。
// xが到達不能になり、ファイナライザが呼び出される前に、メインプログラムで時折変更される場合。
// メインプログラムでの変更とファイナライザでの検査は、読み書き競合を回避するために、ミューテックスやアトミック更新などの適切な同期を使用する必要があります。
func SetFinalizer(obj any, finalizer any)

// KeepAliveは、引数を現在到達可能なものとしてマークします。
// これにより、オブジェクトが解放されず、そのファイナライザが実行されないようになります。
// KeepAliveが呼び出されたプログラムのポイントより前に。
//
// KeepAliveが必要な場所を示す非常に簡単な例：
//
//	type File struct { d int }
//	d, err := syscall.Open("/file/path", syscall.O_RDONLY, 0)
//	// ... errがnilでない場合は何かを実行します ...
//	p := &File{d}
//	runtime.SetFinalizer(p, func(p *File) { syscall.Close(p.d) })
//	var buf [10]byte
//	n, err := syscall.Read(p.d, buf[:])
//	// Readが返るまで、pがファイナライズされないようにします。
//	runtime.KeepAlive(p)
//	// このポイント以降、pを使用しないでください。
//
// KeepAlive呼び出しがない場合、ファイナライザは [syscall.Read] の開始時に実行され、
// 実際のシステムコールを行う前にファイルディスクリプタを閉じます。
//
// 注意：KeepAliveは、ファイナライザが予期せず実行されるのを防止するためにのみ使用する必要があります。
// 特に、[unsafe.Pointer] と一緒に使用する場合は、unsafe.Pointerの有効な使用方法のルールが適用されます。
func KeepAlive(x any)
