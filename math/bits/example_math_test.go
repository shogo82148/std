// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bits_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/math/bits"
)

func ExampleAdd32() {
	// 最初の数は 33<<32 + 12
	n1 := []uint32{33, 12}
	// 2番目の数は 21<<32 + 23
	n2 := []uint32{21, 23}
	// それらをキャリーを生成せずに合計します。
	d1, carry := bits.Add32(n1[1], n2[1], 0)
	d0, _ := bits.Add32(n1[0], n2[0], carry)
	nsum := []uint32{d0, d1}
	fmt.Printf("%v + %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)

	// 最初の数は 1<<32 + 2147483648
	n1 = []uint32{1, 0x80000000}
	// 2番目の数は 1<<32 + 2147483648
	n2 = []uint32{1, 0x80000000}
	// それらを合計してキャリーを生成します。
	d1, carry = bits.Add32(n1[1], n2[1], 0)
	d0, _ = bits.Add32(n1[0], n2[0], carry)
	nsum = []uint32{d0, d1}
	fmt.Printf("%v + %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)
	// Output:
	// [33 12] + [21 23] = [54 35] (carry bit was 0)
	// [1 2147483648] + [1 2147483648] = [3 0] (carry bit was 1)
}

func ExampleAdd64() {
	// 最初の数は 33<<64 + 12
	n1 := []uint64{33, 12}
	// 2番目の数は 21<<64 + 23
	n2 := []uint64{21, 23}
	// それらをキャリーを生成せずに合計します。
	d1, carry := bits.Add64(n1[1], n2[1], 0)
	d0, _ := bits.Add64(n1[0], n2[0], carry)
	nsum := []uint64{d0, d1}
	fmt.Printf("%v + %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)

	// 最初の数は 1<<64 + 9223372036854775808
	n1 = []uint64{1, 0x8000000000000000}
	// 2番目の数は 1<<64 + 9223372036854775808
	n2 = []uint64{1, 0x8000000000000000}
	// それらを合計してキャリーを生成します。
	d1, carry = bits.Add64(n1[1], n2[1], 0)
	d0, _ = bits.Add64(n1[0], n2[0], carry)
	nsum = []uint64{d0, d1}
	fmt.Printf("%v + %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)
	// Output:
	// [33 12] + [21 23] = [54 35] (carry bit was 0)
	// [1 9223372036854775808] + [1 9223372036854775808] = [3 0] (carry bit was 1)
}

func ExampleSub32() {
	// 最初の数は 33<<32 + 23
	n1 := []uint32{33, 23}
	// 2番目の数は 21<<32 + 12
	n2 := []uint32{21, 12}
	// キャリーを生成せずにそれらを減算します。
	d1, carry := bits.Sub32(n1[1], n2[1], 0)
	d0, _ := bits.Sub32(n1[0], n2[0], carry)
	nsum := []uint32{d0, d1}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)

	// 最初の数は 3<<32 + 2147483647
	n1 = []uint32{3, 0x7fffffff}
	// 2番目の数は 1<<32 + 2147483648
	n2 = []uint32{1, 0x80000000}
	// キャリーを生成してそれらを減算します。
	d1, carry = bits.Sub32(n1[1], n2[1], 0)
	d0, _ = bits.Sub32(n1[0], n2[0], carry)
	nsum = []uint32{d0, d1}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)
	// Output:
	// [33 23] - [21 12] = [12 11] (carry bit was 0)
	// [3 2147483647] - [1 2147483648] = [1 4294967295] (carry bit was 1)
}

func ExampleSub64() {
	// 最初の数は 33<<64 + 23
	n1 := []uint64{33, 23}
	// 2番目の数は 21<<64 + 12
	n2 := []uint64{21, 12}
	// キャリーを生成せずにそれらを減算します。
	d1, carry := bits.Sub64(n1[1], n2[1], 0)
	d0, _ := bits.Sub64(n1[0], n2[0], carry)
	nsum := []uint64{d0, d1}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)

	// 最初の数は 3<<64 + 9223372036854775807
	n1 = []uint64{3, 0x7fffffffffffffff}
	// 2番目の数は 1<<64 + 9223372036854775808
	n2 = []uint64{1, 0x8000000000000000}
	// キャリーを生成してそれらを減算します。
	d1, carry = bits.Sub64(n1[1], n2[1], 0)
	d0, _ = bits.Sub64(n1[0], n2[0], carry)
	nsum = []uint64{d0, d1}
	fmt.Printf("%v - %v = %v (carry bit was %v)\n", n1, n2, nsum, carry)
	// Output:
	// [33 23] - [21 12] = [12 11] (carry bit was 0)
	// [3 9223372036854775807] - [1 9223372036854775808] = [1 18446744073709551615] (carry bit was 1)
}

func ExampleMul32() {
	// 最初の数は 0<<32 + 12
	n1 := []uint32{0, 12}
	// 2番目の数は 0<<32 + 12
	n2 := []uint32{0, 12}
	// オーバーフローを生成せずにそれらを掛け合わせます。
	hi, lo := bits.Mul32(n1[1], n2[1])
	nsum := []uint32{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)

	// 最初の数は 0<<32 + 2147483648
	n1 = []uint32{0, 0x80000000}
	// 2番目の数は 0<<32 + 2
	n2 = []uint32{0, 2}
	// オーバーフローを生成するようにそれらを掛け合わせます。
	hi, lo = bits.Mul32(n1[1], n2[1])
	nsum = []uint32{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)
	// Output:
	// 12 * 12 = [0 144]
	// 2147483648 * 2 = [1 0]
}

func ExampleMul64() {
	// 最初の数は 0<<64 + 12
	n1 := []uint64{0, 12}
	// 2番目の数は 0<<64 + 12
	n2 := []uint64{0, 12}
	// オーバーフローを生成せずにそれらを掛け合わせます。
	hi, lo := bits.Mul64(n1[1], n2[1])
	nsum := []uint64{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)

	// 最初の数は 0<<64 + 9223372036854775808
	n1 = []uint64{0, 0x8000000000000000}
	// 2番目の数は 0<<64 + 2
	n2 = []uint64{0, 2}
	// オーバーフローを生成するようにそれらを掛け合わせます。
	hi, lo = bits.Mul64(n1[1], n2[1])
	nsum = []uint64{hi, lo}
	fmt.Printf("%v * %v = %v\n", n1[1], n2[1], nsum)
	// Output:
	// 12 * 12 = [0 144]
	// 9223372036854775808 * 2 = [1 0]
}

func ExampleDiv32() {
	// 最初の数は 0<<32 + 6
	n1 := []uint32{0, 6}
	// 2番目の数は 0<<32 + 3
	n2 := []uint32{0, 3}
	// それらを割ります。
	quo, rem := bits.Div32(n1[0], n1[1], n2[1])
	nsum := []uint32{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)

	// 最初の数は 2<<32 + 2147483648
	n1 = []uint32{2, 0x80000000}
	// 2番目の数は 0<<32 + 2147483648
	n2 = []uint32{0, 0x80000000}
	// それらを割ります。
	quo, rem = bits.Div32(n1[0], n1[1], n2[1])
	nsum = []uint32{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)
	// Output:
	// [0 6] / 3 = [2 0]
	// [2 2147483648] / 2147483648 = [5 0]
}

func ExampleDiv64() {
	// 最初の数は 0<<64 + 6
	n1 := []uint64{0, 6}
	// 2番目の数は 0<<64 + 3
	n2 := []uint64{0, 3}
	// それらを割ります。
	quo, rem := bits.Div64(n1[0], n1[1], n2[1])
	nsum := []uint64{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)

	// 最初の数は 2<<64 + 9223372036854775808
	n1 = []uint64{2, 0x8000000000000000}
	// 2番目の数は 0<<64 + 9223372036854775808
	n2 = []uint64{0, 0x8000000000000000}
	// それらを割ります。
	quo, rem = bits.Div64(n1[0], n1[1], n2[1])
	nsum = []uint64{quo, rem}
	fmt.Printf("[%v %v] / %v = %v\n", n1[0], n1[1], n2[1], nsum)
	// Output:
	// [0 6] / 3 = [2 0]
	// [2 9223372036854775808] / 9223372036854775808 = [5 0]
}
