// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package big_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/math"
	"github.com/shogo82148/std/math/big"
)

func ExampleFloat_Add() {
	// 異なる精度の数値で操作します。
	var x, y, z big.Float
	x.SetInt64(1000)          // xは自動的に64ビット精度に設定されます
	y.SetFloat64(2.718281828) // yは自動的に53ビット精度に設定されます
	z.SetPrec(32)
	z.Add(&x, &y)
	fmt.Printf("x = %.10g (%s, prec = %d, acc = %s)\n", &x, x.Text('p', 0), x.Prec(), x.Acc())
	fmt.Printf("y = %.10g (%s, prec = %d, acc = %s)\n", &y, y.Text('p', 0), y.Prec(), y.Acc())
	fmt.Printf("z = %.10g (%s, prec = %d, acc = %s)\n", &z, z.Text('p', 0), z.Prec(), z.Acc())
	// Output:
	// x = 1000 (0x.fap+10, prec = 64, acc = Exact)
	// y = 2.718281828 (0x.adf85458248cd8p+2, prec = 53, acc = Exact)
	// z = 1002.718282 (0x.faadf854p+10, prec = 32, acc = Below)
}

func ExampleFloat_shift() {
	// (二進数の)指数を直接修正することでFloat "shift"を実装します。
	for s := -5; s <= 5; s++ {
		x := big.NewFloat(0.5)
		x.SetMantExp(x, x.MantExp(nil)+s) // xをsだけシフトします
		fmt.Println(x)
	}
	// Output:
	// 0.015625
	// 0.03125
	// 0.0625
	// 0.125
	// 0.25
	// 0.5
	// 1
	// 2
	// 4
	// 8
	// 16
}

func ExampleFloat_Cmp() {
	inf := math.Inf(1)
	zero := 0.0

	operands := []float64{-inf, -1.2, -zero, 0, +1.2, +inf}

	fmt.Println("   x     y  cmp")
	fmt.Println("---------------")
	for _, x64 := range operands {
		x := big.NewFloat(x64)
		for _, y64 := range operands {
			y := big.NewFloat(y64)
			fmt.Printf("%4g  %4g  %3d\n", x, y, x.Cmp(y))
		}
		fmt.Println()
	}

	// Output:
	//    x     y  cmp
	// ---------------
	// -Inf  -Inf    0
	// -Inf  -1.2   -1
	// -Inf    -0   -1
	// -Inf     0   -1
	// -Inf   1.2   -1
	// -Inf  +Inf   -1
	//
	// -1.2  -Inf    1
	// -1.2  -1.2    0
	// -1.2    -0   -1
	// -1.2     0   -1
	// -1.2   1.2   -1
	// -1.2  +Inf   -1
	//
	//   -0  -Inf    1
	//   -0  -1.2    1
	//   -0    -0    0
	//   -0     0    0
	//   -0   1.2   -1
	//   -0  +Inf   -1
	//
	//    0  -Inf    1
	//    0  -1.2    1
	//    0    -0    0
	//    0     0    0
	//    0   1.2   -1
	//    0  +Inf   -1
	//
	//  1.2  -Inf    1
	//  1.2  -1.2    1
	//  1.2    -0    1
	//  1.2     0    1
	//  1.2   1.2    0
	//  1.2  +Inf   -1
	//
	// +Inf  -Inf    1
	// +Inf  -1.2    1
	// +Inf    -0    1
	// +Inf     0    1
	// +Inf   1.2    1
	// +Inf  +Inf    0
}

func ExampleRoundingMode() {
	operands := []float64{2.6, 2.5, 2.1, -2.1, -2.5, -2.6}

	fmt.Print("   x")
	for mode := big.ToNearestEven; mode <= big.ToPositiveInf; mode++ {
		fmt.Printf("  %s", mode)
	}
	fmt.Println()

	for _, f64 := range operands {
		fmt.Printf("%4g", f64)
		for mode := big.ToNearestEven; mode <= big.ToPositiveInf; mode++ {
			// 上記のサンプルオペランドは、仮数を表現するために2ビットを必要とします
			// それらを整数値に丸めるために、二進数の精度を2に設定します
			f := new(big.Float).SetPrec(2).SetMode(mode).SetFloat64(f64)
			fmt.Printf("  %*g", len(mode.String()), f)
		}
		fmt.Println()
	}

	// Output:
	//    x  ToNearestEven  ToNearestAway  ToZero  AwayFromZero  ToNegativeInf  ToPositiveInf
	//  2.6              3              3       2             3              2              3
	//  2.5              2              3       2             3              2              3
	//  2.1              2              2       2             3              2              3
	// -2.1             -2             -2      -2            -3             -3             -2
	// -2.5             -2             -3      -2            -3             -3             -2
	// -2.6             -3             -3      -2            -3             -3             -2
}
