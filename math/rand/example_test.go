// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package rand_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/math/rand"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/strings"
	"github.com/shogo82148/std/text/tabwriter"
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
	fmt.Println("Magic 8-Ball says:", answers[rand.Intn(len(answers))])
}

// この例では、*Randの各メソッドの使用を示しています。
// グローバル関数の使用は、レシーバーなしで同じです。
func Example_rand() {
	// ジェネレータを作成し、シードを設定します。
	// 通常、固定されていないシードを使用するべきで、例えばtime.Now().UnixNano()などです。
	// 固定されたシードを使用すると、毎回の実行で同じ出力が生成されます。
	r := rand.New(rand.NewSource(99))

	// ここでのtabwriterは、整列した出力を生成するのに役立ちます。
	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	defer w.Flush()
	show := func(name string, v1, v2, v3 any) {
		fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", name, v1, v2, v3)
	}

	// Float32とFloat64の値は[0, 1)の範囲内です。
	show("Float32", r.Float32(), r.Float32(), r.Float32())
	show("Float64", r.Float64(), r.Float64(), r.Float64())

	// ExpFloat64の値は平均が1ですが、指数関数的に減衰します。
	show("ExpFloat64", r.ExpFloat64(), r.ExpFloat64(), r.ExpFloat64())

	// NormFloat64の値は平均が0で、標準偏差が1です。
	show("NormFloat64", r.NormFloat64(), r.NormFloat64(), r.NormFloat64())

	// Int31、Int63、およびUint32は、指定された幅の値を生成します。
	// Intメソッド（ここでは示されていません）は、'int'のサイズに応じてInt31またはInt63のどちらかと同様です。
	show("Int31", r.Int31(), r.Int31(), r.Int31())
	show("Int63", r.Int63(), r.Int63(), r.Int63())
	show("Uint32", r.Uint32(), r.Uint32(), r.Uint32())

	// Intn、Int31n、およびInt63nは、出力をn未満に制限します。
	// これらは、r.Int()%nを使用するよりも慎重に行います。
	show("Intn(10)", r.Intn(10), r.Intn(10), r.Intn(10))
	show("Int31n(10)", r.Int31n(10), r.Int31n(10), r.Int31n(10))
	show("Int63n(10)", r.Int63n(10), r.Int63n(10), r.Int63n(10))

	// Permは、数値[0, n)のランダムな順列を生成します。
	show("Perm", r.Perm(5), r.Perm(5), r.Perm(5))
	// Output:
	// Float32     0.2635776           0.6358173           0.6718283
	// Float64     0.628605430454327   0.4504798828572669  0.9562755949377957
	// ExpFloat64  0.3362240648200941  1.4256072328483647  0.24354758816173044
	// NormFloat64 0.17233959114940064 1.577014951434847   0.04259129641113857
	// Int31       1501292890          1486668269          182840835
	// Int63       3546343826724305832 5724354148158589552 5239846799706671610
	// Uint32      2760229429          296659907           1922395059
	// Intn(10)    1                   2                   5
	// Int31n(10)  4                   7                   8
	// Int63n(10)  7                   6                   3
	// Perm        [1 4 2 3 0]         [4 2 1 3 0]         [1 2 4 0 3]
}

func ExamplePerm() {
	for _, value := range rand.Perm(3) {
		fmt.Println(value)
	}

	// Unordered output: 1
	// 2
	// 0
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

func ExampleIntn() {
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
}
