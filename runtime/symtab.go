// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

<<<<<<< HEAD
// Framesを使用すると、Callersが返すPC値のスライスのための関数/ファイル/行情報を取得できます。
=======
// Frames may be used to get function/file/line information for a
// slice of PC values returned by [Callers].
>>>>>>> upstream/master
type Frames struct {
	// callersはまだフレームに展開されていないPCのスライスです。
	callers []uintptr

	// frames はまだ返却されていない Frames のスライスです。
	frames     []Frame
	frameStore [2]Frame
}

// Frameは各コールフレームごとにFramesによって返される情報です。
type Frame struct {

	// PCはこのフレームの位置に対するプログラムカウンタです。
	// 別のフレームを呼び出すフレームの場合、これは
	// 呼び出し命令のプログラムカウンタです。インライン展開のため、
	// 複数のフレームは同じPC値を持つことがありますが、異なる
	// シンボリック情報を持ちます。
	PC uintptr

	// Funcはこの呼び出しフレームのFunc値です。これは、非Goコードや完全にインライン化された関数の場合はnilになることがあります。
	Func *Func

	// Functionはこの呼び出しフレームのパッケージパス修飾された関数名です。非空であれば、この文字列はプログラム内の1つの関数を一意に識別します。
	// これは知られていない場合は空の文字列になることがあります。
	// Funcがnilでない場合、Function == Func.Name()です。
	Function string

	// FileとLineは、このフレームのファイル名と行番号です。
	// 非終端フレームの場合、これは呼び出しの位置になります。
	// もし分かっていない場合は、それぞれ空文字列とゼロになります。
	File string
	Line int

	// startLineは、このフレームの関数の開始行番号です。具体的には、Goの関数のfuncキーワードの行番号です。注意点として、//lineディレクティブは、関数内で任意のファイル名や行番号を変更することができ、したがってLine - startLineのオフセットは常に意味を持たないことがあります。
	// もし知られていない場合、これはゼロになる場合があります。
	startLine int

	// 関数のエントリーポイントのプログラムカウンター。不明の場合はゼロ。
	// Funcがnilでない場合、Entry == Func.Entry()。
	Entry uintptr

	// ランタイムの内部ビューでの関数。このフィールドは、Goの関数にのみ設定されます（funcInfo.valid()がtrueを返します）、Cの関数には設定されません。
	funcInfo funcInfo
}

<<<<<<< HEAD
// CallersFramesはCallersによって返されるPC値のスライスを受け取り、
// 関数/ファイル/行情報を返す準備をします。
// Framesで終わるまでスライスを変更しないでください。
=======
// CallersFrames takes a slice of PC values returned by [Callers] and
// prepares to return function/file/line information.
// Do not change the slice until you are done with the [Frames].
>>>>>>> upstream/master
func CallersFrames(callers []uintptr) *Frames

// Nextは、PC値のスライス内で次の呼び出しフレームを表すFrameを返します。
// すべての呼び出しフレームをすでに返した場合、NextはゼロのFrameを返します。
//
// moreの結果は、次のNext呼び出しで有効なFrameが返されるかどうかを示します。
// これが呼び出し元に一つ返されたかどうかを必ずしも示しません。
//
// 典型的な使用法については、Framesの例を参照してください。
func (ci *Frames) Next() (frame Frame, more bool)

// Funcは実行中のバイナリ内のGo関数を表します。
type Func struct {
	opaque struct{}
}

<<<<<<< HEAD
// FuncForPCは、指定されたプログラムカウンターアドレスを含む関数を記述した*Funcを返します。もし複数の関数がインライン展開の影響で存在する場合は、最も内側の関数を示す*Funcを返しますが、最も外側の関数のエントリーも持っています。
=======
// FuncForPC returns a *[Func] describing the function that contains the
// given program counter address, or else nil.
//
// If pc represents multiple functions because of inlining, it returns
// the *Func describing the innermost function, but with an entry of
// the outermost function.
>>>>>>> upstream/master
func FuncForPC(pc uintptr) *Func

// Nameは関数の名前を返します。
func (f *Func) Name() string

// Entryは関数のエントリーアドレスを返します。
func (f *Func) Entry() uintptr

// FileLineは、プログラムカウンターpcに対応するソースコードのファイル名と行番号を返します。
// pcがfのプログラムカウンターでない場合、結果は正確ではありません。
func (f *Func) FileLine(pc uintptr) (file string, line int)
