// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/math"
	"github.com/shogo82148/std/math/big"
)

func ExampleRat_SetString() {
	r := new(big.Rat)
	r.SetString("355/113")
	fmt.Println(r.FloatString(3))
	// Output: 3.142
}

func ExampleInt_SetString() {
	i := new(big.Int)
	i.SetString("644", 8) // 8進数
	fmt.Println(i)
	// Output: 420
}

func ExampleFloat_SetString() {
	f := new(big.Float)
	f.SetString("3.14159")
	fmt.Println(f)
	// Output: 3.14159
}

func ExampleRat_Scan() {
	// Scan関数は直接使用されることはほとんどありません。
	// fmtパッケージは、fmt.Scannerの実装としてこれを認識します。
	r := new(big.Rat)
	_, err := fmt.Sscan("1.5000", r)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(r)
	}
	// Output: 3/2
}

func ExampleInt_Scan() {
	// Scan関数は直接使用されることはほとんどありません。
	// fmtパッケージは、fmt.Scannerの実装としてこれを認識します。
	i := new(big.Int)
	_, err := fmt.Sscan("18446744073709551617", i)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(i)
	}
	// Output: 18446744073709551617
}

func ExampleFloat_Scan() {
	// Scan関数は直接使用されることはほとんどありません。
	// fmtパッケージは、fmt.Scannerの実装としてこれを認識します。
	f := new(big.Float)
	_, err := fmt.Sscan("1.19282e99", f)
	if err != nil {
		log.Println("error scanning value:", err)
	} else {
		fmt.Println(f)
	}
	// Output: 1.19282e+99
}

// この例では、big.Intを使用して100桁の最小のフィボナッチ数を計算し、
// それが素数であるかどうかをテストする方法を示しています。
func Example_fibonacci() {
	// シーケンスの最初の2つの数で2つのbig intを初期化します。
	a := big.NewInt(0)
	b := big.NewInt(1)

	// limitを10^99（100桁の最小の整数）として初期化します。
	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)

	// aが1e100より小さい間ループします。
	for a.Cmp(&limit) < 0 {
		// 次のフィボナッチ数を計算し、それをaに格納します。
		a.Add(a, b)
		// aとbを交換して、bがシーケンスの次の数になるようにします。
		a, b = b, a
	}
	fmt.Println(a) // 100桁のフィボナッチ数

	// aが素数であるかどうかをテストします。
	// (ProbablyPrimesの引数は、実行するミラー-ラビン
	// ラウンドの数を設定します。20は良い値です。)
	fmt.Println(a.ProbablyPrime(20))

	// Output:
	// 1344719667586153181419716641724567886890850696275767987106294472017884974410332069524504824747437757
	// false
}

// この例では、big.Floatを使用して精度200ビットで2の平方根を計算し、
// 結果を10進数として印刷する方法を示します。
func Example_sqrt2() {
	// 我々は、仮数部に200ビットの精度で計算を行います。
	const prec = 200

	// ニュートンの方法を使用して2の平方根を計算します。我々は
	// sqrt(2)の初期推定値から始め、次に反復します：
	//     x_{n+1} = 1/2 * ( x_n + (2.0 / x_n) )

	// ニュートンの方法は各反復で正確な桁数を2倍にするため、
	// 我々は少なくともlog_2(prec)ステップが必要です。
	steps := int(math.Log2(prec))

	// 計算に必要な値を初期化します。
	two := new(big.Float).SetPrec(prec).SetInt64(2)
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	// 初期推定値として1を使用します。
	x := new(big.Float).SetPrec(prec).SetInt64(1)

	// tを一時的な変数として使用します。big.Floatの値は、精度が設定されていない（== 0）場合、
	// big.Floatの操作の結果（レシーバ）として使用されるときに自動的に引数の最大精度を引き継ぐため、
	// tの精度を設定する必要はありません。
	t := new(big.Float)

	// Iterate.
	for i := 0; i <= steps; i++ {
		t.Quo(two, x)  // t = 2.0 / x_n
		t.Add(x, t)    // t = x_n + (2.0 / x_n)
		x.Mul(half, t) // x_{n+1} = 0.5 * t
	}

	// big.Floatはfmt.Formatterを実装しているので、通常のfmt.Printfの動詞を使用できます
	fmt.Printf("sqrt(2) = %.50f\n", x)

	// 2とx*xの間の誤差を印刷します。
	t.Mul(x, x) // t = x*x
	fmt.Printf("error = %e\n", t.Sub(two, t))

	// Output:
	// sqrt(2) = 1.41421356237309504880168872420969807856967187537695
	// error = 0.000000e+00
}
