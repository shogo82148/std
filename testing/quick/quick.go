// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// quickパッケージは、ブラックボックステストを支援するためのユーティリティ関数を実装します。
//
// testing/quickパッケージは凍結されており、新たな機能は受け入れられません。
package quick

import (
	"github.com/shogo82148/std/math/rand"
	"github.com/shogo82148/std/reflect"
)

// Generatorは、自身の型のランダムな値を生成することができます。
type Generator interface {
<<<<<<< HEAD
	Generate(rand *rand.Rand, size int) reflect.Value
}

// Value returns an arbitrary value of the given type.
// If the type implements the [Generator] interface, that will be used.
// Note: To create arbitrary values for structs, all the fields must be exported.
=======
	// Generateは、サイズをサイズヒントとして使用して、
	// メソッドの型のランダムなインスタンスを返します。
	Generate(rand *rand.Rand, size int) reflect.Value
}

// Valueは、指定された型の任意の値を返します。
// もし型がGeneratorインターフェースを実装しているなら、それが使用されます。
// 注意：構造体の任意の値を作成するためには、全てのフィールドがエクスポートされていなければなりません。
>>>>>>> release-branch.go1.21
func Value(t reflect.Type, rand *rand.Rand) (value reflect.Value, ok bool)

// Config構造体は、テストの実行オプションを含みます。
type Config struct {
	// MaxCountは、反復の最大回数を設定します。
	// もしゼロなら、MaxCountScaleが使用されます。
	MaxCount int
	// MaxCountScaleは、デフォルトの最大値に適用される非負のスケールファクターです。
	// カウントがゼロの場合、デフォルト（通常は100）が適用されますが、
	// -quickchecksフラグによって設定することもできます。
	MaxCountScale float64
	// Randは、乱数のソースを指定します。
	// もしnilなら、デフォルトの疑似乱数ソースが使用されます。
	Rand *rand.Rand
	// Valuesは、テスト対象の関数の引数と一致する任意のreflect.Valuesのスライスを生成する関数を指定します。
	// もしnilなら、トップレベルのValue関数がそれらを生成するために使用されます。
	Values func([]reflect.Value, *rand.Rand)
}

// SetupErrorは、テストされる関数に関係なく、checkの使用方法に関するエラーの結果です。
type SetupError string

func (s SetupError) Error() string

// CheckErrorは、Checkがエラーを見つけた結果です。
type CheckError struct {
	Count int
	In    []any
}

func (s *CheckError) Error() string

<<<<<<< HEAD
// A CheckEqualError is the result [CheckEqual] finding an error.
=======
// CheckEqualErrorは、CheckEqualがエラーを見つけた結果です。
>>>>>>> release-branch.go1.21
type CheckEqualError struct {
	CheckError
	Out1 []any
	Out2 []any
}

func (s *CheckEqualError) Error() string

<<<<<<< HEAD
// Check looks for an input to f, any function that returns bool,
// such that f returns false. It calls f repeatedly, with arbitrary
// values for each argument. If f returns false on a given input,
// Check returns that input as a *[CheckError].
// For example:
=======
// Checkは、fがfalseを返すような入力を探します。fは、boolを返す任意の関数です。
// fは、各引数に対して任意の値を使用して繰り返し呼び出されます。
// もしfが特定の入力でfalseを返した場合、Checkはその入力を*CheckErrorとして返します。
// 例えば：
>>>>>>> release-branch.go1.21
//
//	func TestOddMultipleOfThree(t *testing.T) {
//		f := func(x int) bool {
//			y := OddMultipleOfThree(x)
//			return y%2 == 1 && y%3 == 0
//		}
//		if err := quick.Check(f, nil); err != nil {
//			t.Error(err)
//		}
//	}
func Check(f any, config *Config) error

<<<<<<< HEAD
// CheckEqual looks for an input on which f and g return different results.
// It calls f and g repeatedly with arbitrary values for each argument.
// If f and g return different answers, CheckEqual returns a *[CheckEqualError]
// describing the input and the outputs.
=======
// CheckEqualは、fとgが異なる結果を返す入力を探します。
// fとgは、各引数に対して任意の値を使用して繰り返し呼び出されます。
// もしfとgが異なる答えを返した場合、CheckEqualはその入力と出力を記述する*CheckEqualErrorを返します。
>>>>>>> release-branch.go1.21
func CheckEqual(f, g any, config *Config) error
