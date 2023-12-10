// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// expvarパッケージは、サーバー内の操作カウンターなどの公開変数に対する標準化されたインターフェースを提供します。
// これらの変数は、/debug/varsでJSON形式でHTTP経由で公開されます。
//
// これらの公開変数を設定または変更する操作はアトミックです。
//
// このパッケージはHTTPハンドラを追加するだけでなく、以下の変数も登録します：
//
//	cmdline   os.Args
//	memstats  runtime.Memstats
//
// このパッケージは、HTTPハンドラと上記の変数を登録する副作用のためだけに
// インポートされることがあります。このように使用するには、
// このパッケージをプログラムにリンクします：
//
//	import _ "expvar"
package expvar

import (
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
)

// Varは、すべてのエクスポートされた変数のための抽象型です。
type Var interface {
<<<<<<< HEAD
	String() string
}

// Int is a 64-bit integer variable that satisfies the [Var] interface.
=======
	// Stringは、変数の有効なJSON値を返します。
	// 有効なJSONを返さないStringメソッドを持つ型
	// (time.Timeなど)は、Varとして使用してはなりません。
	String() string
}

// Intは、Varインターフェースを満たす64ビット整数変数です。
>>>>>>> release-branch.go1.21
type Int struct {
	i atomic.Int64
}

func (v *Int) Value() int64

func (v *Int) String() string

func (v *Int) Add(delta int64)

func (v *Int) Set(value int64)

<<<<<<< HEAD
// Float is a 64-bit float variable that satisfies the [Var] interface.
=======
// Floatは、Varインターフェースを満たす64ビット浮動小数点数変数です。
>>>>>>> release-branch.go1.21
type Float struct {
	f atomic.Uint64
}

func (v *Float) Value() float64

func (v *Float) String() string

// Addは、vにdeltaを加えます。
func (v *Float) Add(delta float64)

// Setは、vをvalueに設定します。
func (v *Float) Set(value float64)

<<<<<<< HEAD
// Map is a string-to-Var map variable that satisfies the [Var] interface.
=======
// Mapは、Varインターフェースを満たす文字列からVarへのマップ変数です。
>>>>>>> release-branch.go1.21
type Map struct {
	m      sync.Map
	keysMu sync.RWMutex
	keys   []string
}

<<<<<<< HEAD
// KeyValue represents a single entry in a [Map].
=======
// KeyValueは、Map内の単一のエントリを表します。
>>>>>>> release-branch.go1.21
type KeyValue struct {
	Key   string
	Value Var
}

func (v *Map) String() string

// Initは、マップからすべてのキーを削除します。
func (v *Map) Init() *Map

func (v *Map) Get(key string) Var

func (v *Map) Set(key string, av Var)

<<<<<<< HEAD
// Add adds delta to the *[Int] value stored under the given map key.
func (v *Map) Add(key string, delta int64)

// AddFloat adds delta to the *[Float] value stored under the given map key.
=======
// Addは、指定されたマップキーの下に格納された*Int値にdeltaを加えます。
func (v *Map) Add(key string, delta int64)

// AddFloatは、指定されたマップキーの下に格納された*Float値にdeltaを加えます。
>>>>>>> release-branch.go1.21
func (v *Map) AddFloat(key string, delta float64)

// Deleteは、マップから指定されたキーを削除します。
func (v *Map) Delete(key string)

// Doは、マップ内の各エントリに対してfを呼び出します。
// イテレーション中はマップがロックされますが、
// 既存のエントリは並行して更新される可能性があります。
func (v *Map) Do(f func(KeyValue))

<<<<<<< HEAD
// String is a string variable, and satisfies the [Var] interface.
=======
// Stringは文字列変数で、Varインターフェースを満たします。
>>>>>>> release-branch.go1.21
type String struct {
	s atomic.Value
}

func (v *String) Value() string

<<<<<<< HEAD
// String implements the [Var] interface. To get the unquoted string
// use [String.Value].
=======
// StringはVarインターフェースを実装します。引用符なしの文字列を取得するには
// Valueを使用します。
>>>>>>> release-branch.go1.21
func (v *String) String() string

func (v *String) Set(value string)

<<<<<<< HEAD
// Func implements [Var] by calling the function
// and formatting the returned value using JSON.
=======
// Funcは、関数を呼び出し、返された値をJSONを使用してフォーマットすることでVarを実装します。
>>>>>>> release-branch.go1.21
type Func func() any

func (f Func) Value() any

func (f Func) String() string

// Publishは、名前付きのエクスポート変数を宣言します。これは、パッケージが
// Varsを作成するときのinit関数から呼び出されるべきです。もし名前がすでに
// 登録されている場合、これはlog.Panicを引き起こします。
func Publish(name string, v Var)

// Getは、名前付きのエクスポート変数を取得します。名前が
// 登録されていない場合、nilを返します。
func Get(name string) Var

func NewInt(name string) *Int

func NewFloat(name string) *Float

func NewMap(name string) *Map

func NewString(name string) *String

// Doは、各エクスポートされた変数に対してfを呼び出します。
// イテレーション中はグローバル変数マップがロックされますが、
// 既存のエントリは並行して更新される可能性があります。
func Do(f func(KeyValue))

// Handlerは、expvarのHTTPハンドラを返します。
//
// これは、ハンドラを非標準の場所にインストールする必要がある場合のみ必要です。
func Handler() http.Handler
