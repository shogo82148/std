// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/math/rand/v2"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/strings"
	"github.com/shogo82148/std/text/tabwriter"
	"github.com/shogo82148/std/time"
)

func Example() {
	answers := []string{
		"It is certain",
		"It is decidedly so",
		"Without a doubt",
		"Yes definitely",
		"You may rely on it",
		"As I see it yes",
		"Most likely",
		"Outlook good",
		"Yes",
		"Signs point to yes",
		"Reply hazy try again",
		"Ask again later",
		"Better not tell you now",
		"Cannot predict now",
		"Concentrate and ask again",
		"Don't count on it",
		"My reply is no",
		"My sources say no",
		"Outlook not so good",
		"Very doubtful",
	}
	fmt.Println("Magic 8-Ball says:", answers[rand.IntN(len(answers))])
}

// この例は、*Randの各メソッドの使用を示しています。
// グローバル関数の使用は、レシーバなしで同じです。
func Example_rand() {
	// ジェネレータを作成し、シードを設定します。
	// 通常、固定されていないシードを使用するべきです。例えば、Uint64(), Uint64()のようなものです。
	// 固定シードを使用すると、毎回同じ出力が生成されます。
	r := rand.New(rand.NewPCG(1, 2))

	// ここでのtabwriterは、私たちが整列した出力を生成するのを助けてくれます。
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	show := func(name string, v1, v2, v3 any) {
		fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", name, v1, v2, v3)
	}

	// Float32とFloat64の値は[0, 1)の範囲内にあります。
	show("Float32", r.Float32(), r.Float32(), r.Float32())
	show("Float64", r.Float64(), r.Float64(), r.Float64())

	// ExpFloat64の値は平均が1ですが、指数関数的に減衰します。
	show("ExpFloat64", r.ExpFloat64(), r.ExpFloat64(), r.ExpFloat64())

	// NormFloat64の値は平均が0で、標準偏差が1です。
	show("NormFloat64", r.NormFloat64(), r.NormFloat64(), r.NormFloat64())

	// Int32、Int64、およびUint32は指定された幅の値を生成します。
	// Intメソッド（ここでは示されていません）は'int'のサイズに応じてInt32またはInt64のように動作します。
	show("Int32", r.Int32(), r.Int32(), r.Int32())
	show("Int64", r.Int64(), r.Int64(), r.Int64())
	show("Uint32", r.Uint32(), r.Uint32(), r.Uint32())

	// IntN、Int32N、およびInt64Nは、出力がn未満になるように制限します。
	// これらは、r.Int()%nを使用するよりも慎重に行います。
	show("IntN(10)", r.IntN(10), r.IntN(10), r.IntN(10))
	show("Int32N(10)", r.Int32N(10), r.Int32N(10), r.Int32N(10))
	show("Int64N(10)", r.Int64N(10), r.Int64N(10), r.Int64N(10))

	// Permは、数値[0, n)のランダムな順列を生成します。
	show("Perm", r.Perm(5), r.Perm(5), r.Perm(5))
	// Output:
	// Float32     0.95955694          0.8076733            0.8135684
	// Float64     0.4297927436037299  0.797802349388613    0.3883664855410056
	// ExpFloat64  0.43463410545541104 0.5513632046504593   0.7426404617374481
	// NormFloat64 -0.9303318111676635 -0.04750789419852852 0.22248301107582735
	// Int32       2020777787          260808523            851126509
	// Int64       5231057920893523323 4257872588489500903  158397175702351138
	// Uint32      314478343           1418758728           208955345
	// IntN(10)    6                   2                    0
	// Int32N(10)  3                   7                    7
	// Int64N(10)  8                   9                    4
	// Perm        [0 3 1 4 2]         [4 1 2 0 3]          [4 3 2 0 1]
}

func ExamplePerm() {
	for _, value := range rand.Perm(3) {
		fmt.Println(value)
	}

	// Unordered output: 1
	// 2
	// 0
}

func ExampleN() {
	// 半開放区間[0, 100)内のint64を出力します。
	fmt.Println(rand.N(int64(100)))

	// 0から100ミリ秒の間のランダムな期間スリープします。
	time.Sleep(rand.N(100 * time.Millisecond))
}

func ExampleShuffle() {
	words := strings.Fields("ink runs from the corners of my mouth")
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	fmt.Println(words)
}

func ExampleShuffle_slicesInUnison() {
	numbers := []byte("12345")
	letters := []byte("ABCDE")
	// 数字をシャッフルし、同時にlettersの対応するエントリを交換します。
	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
		letters[i], letters[j] = letters[j], letters[i]
	})
	for i := range numbers {
		fmt.Printf("%c: %c\n", letters[i], numbers[i])
	}
}

func ExampleIntN() {
	fmt.Println(rand.IntN(100))
	fmt.Println(rand.IntN(100))
	fmt.Println(rand.IntN(100))
}
