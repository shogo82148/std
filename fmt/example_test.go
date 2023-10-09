// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fmt_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/math"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/strings"
	"github.com/shogo82148/std/time"
)

// Errorf関数を使用することで、表示形式の機能を使って
// 詳細なエラーメッセージを作成することができます。
func ExampleErrorf() {
	const name, id = "bueller", 17
	err := fmt.Errorf("user %q (id %d) not found", name, id)
	fmt.Println(err.Error())

	// Output: user "bueller" (id 17) not found
}

func ExampleFscanf() {
	var (
		i int
		b bool
		s string
	)
	r := strings.NewReader("5 true gophers")
	n, err := fmt.Fscanf(r, "%d %t %s", &i, &b, &s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fscanf: %v\n", err)
	}
	fmt.Println(i, b, s)
	fmt.Println(n)
	// Output:
	// 5 true gophers
	// 3
}

func ExampleFscanln() {
	s := `dmr 1771 1.61803398875
	ken 271828 3.14159`
	r := strings.NewReader(s)
	var a string
	var b int
	var c float64
	for {
		n, err := fmt.Fscanln(r, &a, &b, &c)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d: %s, %d, %f\n", n, a, b, c)
	}
	// Output:
	// 3: dmr, 1771, 1.618034
	// 3: ken, 271828, 3.141590
}

func ExampleSscanf() {
	var name string
	var age int
	n, err := fmt.Sscanf("Kim is 22 years old", "%s is %d years old", &name, &age)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%d: %s, %d\n", n, name, age)

	// Output:
	// 2: Kim, 22
}

func ExamplePrint() {
	const name, age = "Kim", 22
	fmt.Print(name, " is ", age, " years old.\n")

	// Print が返すエラーについては心配する必要がないというのは慣習です。

	// Output:
	// Kim is 22 years old.
}

func ExamplePrintln() {
	const name, age = "Kim", 22
	fmt.Println(name, "is", age, "years old.")

	// Printlnが返すエラーについては、心配しないのが一般的です。

	// Output:
	// Kim is 22 years old.
}

func ExamplePrintf() {
	const name, age = "Kim", 22
	fmt.Printf("%s is %d years old.\n", name, age)

	// Printfが返すエラーについては心配しないのが通例です。

	// Output:
	// Kim is 22 years old.
}

func ExampleSprint() {
	const name, age = "Kim", 22
	s := fmt.Sprint(name, " is ", age, " years old.\n")

	io.WriteString(os.Stdout, s) // シンプルさのためにエラーを無視しています。

	// Output:
	// Kim is 22 years old.
}

func ExampleSprintln() {
	const name, age = "Kim", 22
	s := fmt.Sprintln(name, "is", age, "years old.")

	io.WriteString(os.Stdout, s) // シンプルさを考慮してエラーを無視します。

	// Output:
	// Kim is 22 years old.
}

func ExampleSprintf() {
	const name, age = "Kim", 22
	s := fmt.Sprintf("%s is %d years old.\n", name, age)

	io.WriteString(os.Stdout, s) // 単純化のため、エラーを無視しています。

	// Output:
	// Kim is 22 years old.
}

func ExampleFprint() {
	const name, age = "Kim", 22
	n, err := fmt.Fprint(os.Stdout, name, " is ", age, " years old.\n")

	// Fprint の n と err の返り値は、基礎となる io.Writer から返されたものです。
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprint: %v\n", err)
	}
	fmt.Print(n, " bytes written.\n")

	// Output:
	// Kim is 22 years old.
	// 21 bytes written.
}

func ExampleFprintln() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintln(os.Stdout, name, "is", age, "years old.")

	// Fprintlnのnとerrの返り値は、基礎となるio.Writerから返されるものです。
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintln: %v\n", err)
	}
	fmt.Println(n, "bytes written.")

	// Output:
	// Kim is 22 years old.
	// 21 bytes written.
}

func ExampleFprintf() {
	const name, age = "Kim", 22
	n, err := fmt.Fprintf(os.Stdout, "%s is %d years old.\n", name, age)

	// Fprintfからのnとerrの返り値は、基になるio.Writerによって返されたものです。
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fprintf: %v\n", err)
	}
	fmt.Printf("%d bytes written.\n", n)

	// Output:
	// Kim is 22 years old.
	// 21 bytes written.
}

// Print、Println、およびPrintfは、引数を異なるレイアウトで配置します。この例では、それらの振る舞いを比較することができます。Printlnは常に出力する項目の間に空白を追加しますが、Printは非文字列の引数の間のみ空白を追加し、Printfは正確に指示通りの出力を行います。
// Sprint、Sprintln、Sprintf、Fprint、Fprintln、およびFprintfは、ここに示すPrint、Println、およびPrintf関数と同じように動作します。
func Example_printers() {
	a, b := 3.0, 4.0
	h := math.Hypot(a, b)

	// Printは、どちらも文字列でない場合に引数間に空白を挿入します。
	// 出力に改行は追加されませんので、明示的に追加します。
	fmt.Print("The vector (", a, b, ") has length ", h, ".\n")

	// Printlnは常に引数の間にスペースを挿入するため、
	// この場合にはPrintと同じ出力を生成するためには使用できません。
	// 出力には余分なスペースが含まれています。
	// また、Printlnは常に出力に改行を追加します。
	fmt.Println("The vector (", a, b, ") has length", h, ".")

	// Printfは完全な制御を提供しますが、使用する際にはより複雑です。
	// 出力に改行が追加されないため、フォーマット指定文字列の最後に明示的に追加します。
	fmt.Printf("The vector (%g %g) has length %g.\n", a, b, h)

	// Output:
	// The vector (3 4) has length 5.
	// The vector ( 3 4 ) has length 5 .
	// The vector (3 4) has length 5.
}

// これらの例は、フォーマット文字列を使用して印刷する基本的な方法を示しています。Printf、Sprintf、およびFprintfは、次の引数の書式を指定するフォーマット文字列を受け取ります。たとえば、%d（これを「動詞」と呼びます）は、対応する引数を10進数で表示することを意味します。その引数は整数（または整数を含むもの、例えば整数のスライス）でなければなりません。動詞%v（'v'は'value'を意味します）は、引数をそのデフォルトの形式で表示します。PrintまたはPrintlnのように。特別な動詞%T（'T'は'Type'を意味します）は、値ではなく引数の型を表示します。これらの例は網羅的ではありません。詳細については、パッケージのコメントを参照してください。
func Example_formats() {

	// %vはデフォルトの形式であることを示す基本的な例のセットです。この場合、整数に対しては10進数の形式（%d）が明示的に要求されることがあります。出力はPrintlnが生成するものと同じです。
	integer := 23
	// それぞれの出力は「23」です（引用符なし）。
	fmt.Println(integer)
	fmt.Printf("%v\n", integer)
	fmt.Printf("%d\n", integer)

	// 特別な動詞％Tは、値ではなくアイテムの型を示します。
	fmt.Printf("%T %T\n", integer, &integer)
	// 結果: int *int

	// Println(x) は Printf("%v\n", x) と同じなので、以下の例ではPrintfのみを使用します。
	// 各例は、整数や文字列など特定の型の値をどのようにフォーマットするかを示しています。
	// 各フォーマット文字列は %v で始めてデフォルトの出力を示し、それに続けて1つ以上のカスタムフォーマットが続きます。

	// ブール値は、%v や %t とともに "true" または "false" として表示されます。
	truth := true
	fmt.Printf("%v %t\n", truth, truth)
	// 結果: true true

	// 整数は %v や %d を使って10進数で表示されます。
	// %x を使うと16進数で表示され、%o を使うと8進数、%b を使うと2進数で表示されます。
	answer := 42
	fmt.Printf("%v %d %x %o %b\n", answer, answer, answer, answer, answer)
	// 結果: 42 42 2a 52 101010

	// 浮動小数点数は複数のフォーマットを持っています: %v と %g はコンパクトな表現を出力し、
	// %f は小数点を出力し、%e は指数表記を使用します。ここで使用されている %6.2f のフォーマットは、
	// 浮動小数点値の表示方法を制御するために幅と精度を設定する方法を示しています。この場合、6 は
	// 値の出力テキストの総幅です(出力の余分なスペースに注意してください)。2 は表示する小数点以下の桁数です。
	pi := math.Pi
	fmt.Printf("%v %g %.2f (%6.2f) %e\n", pi, pi, pi, pi, pi)
	// 結果：3.141592653589793 3.141592653589793 3.14 (  3.14) 3.141593e+00

	// 複素数は、実部と虚部の浮動小数点を括弧で囲んで、虚部の後に「i」を付けた形式です。
	point := 110.7 + 22.5i
	fmt.Printf("%v %g %.2f %.2e\n", point, point, point, point)
	// 結果: (110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)

	// Runesは整数ですが、%cで印刷するとそのUnicode値に対応する文字が表示されます。
	// %qの動詞はクオートされた文字として表示し、%Uは16進数のUnicodeコードポイントとして表示し、
	// %#Uはコードポイントとクオートされた表示可能な形式の両方として表示します。
	smile := '😀'
	fmt.Printf("%v %d %c %q %U %#U\n", smile, smile, smile, smile, smile, smile)
	// 結果: 128512 128512 😀 '😀' U+1F600 U+1F600 '😀'

	// 文字列は、%vと%sはそのまま、%qは引用符付き文字列、そして%#qはバッククォート付き文字列としてフォーマットされます。
	placeholders := `foo "bar"`
	fmt.Printf("%v %s %q %#q\n", placeholders, placeholders, placeholders, placeholders)
	// 結果：foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`

	// %vでフォーマットされたマップは、キーと値をデフォルトの形式で表示します。
	// %#v形式（#はこのコンテキストで「フラグ」と呼ばれます）では、マップをGoのソース形式で表示します。
	// マップはキーの値に応じて一貫した順序で表示されます。
	isLegume := map[string]bool{
		"peanut":    true,
		"dachshund": false,
	}
	fmt.Printf("%v %#v\n", isLegume, isLegume)
	// 結果: map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}

	// %vでフォーマットされた構造体は、そのデフォルトの形式でフィールドの値を表示します。
	// %+vの形式はフィールドを名前付きで表示しますが、%#vは構造体をGoのソース形式でフォーマットします。
	person := struct {
		Name string
		Age  int
	}{"Kim", 22}
	fmt.Printf("%v %+v %#v\n", person, person, person)
	// 結果: {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}

	// ポインタのデフォルトのフォーマットは、アンパサンドに続く元の値を表示します。
	// %p フォーマット指定子はポインタの値を16進数で表示します。ここでは、
	// %p への引数には型付きの nil を使用しています。なぜなら、非 nil のポインタ
	// の値は実行ごとに変化するためです；コメントアウトされた Printf 呼び出し
	// を自分で実行するとわかります。
	pointer := &person
	fmt.Printf("%v %p\n", pointer, (*int)(nil))

	// 結果: &{Kim 22} 0x0
	// fmt.Printf("%v %p\n", pointer, pointer)
	// 結果: &{Kim 22} 0x010203 // 上のコメントを参照してください。

	// 配列やスライスは、各要素に対してフォーマットを適用して表示されます。
	greats := [5]string{"Kitano", "Kobayashi", "Kurosawa", "Miyazaki", "Ozu"}
	fmt.Printf("%v %q\n", greats, greats)
	// 結果: [北野 小林 黒沢 宮崎 小津] ["北野" "小林" "黒沢" "宮崎" "小津"]

	kGreats := greats[:3]
	fmt.Printf("%v %q %#v\n", kGreats, kGreats, kGreats)
	// 結果: [北野 小林 黒沢] ["北野" "小林" "黒沢"] []string{"北野", "小林", "黒沢"}

	// バイトスライスは特別です。%dのような整数の形式で要素を印字します。%sと%qの形式ではスライスを文字列として扱います。%xの動詞は、スペースフラグを持つ特別な形式で、バイト間にスペースを挿入します。
	cmd := []byte("a⌘")
	fmt.Printf("%v %d %s %q %x % x\n", cmd, cmd, cmd, cmd, cmd, cmd)
	// 結果: [97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98

	// Stringerを実装するタイプは文字列と同じように表示されます。Stringerは文字列を返すため、%qなどの文字列専用のフォーマット指定子を使用して印刷することができます。
	now := time.Unix(123456789, 0).UTC()
	fmt.Printf("%v %q\n", now, now)
	// 結果: 1973年11月29日 21時33分09秒 +0000 UTC "1973年11月29日 21時33分09秒 +0000 UTC"

	// Output:
	// 23
	// 23
	// 23
	// int *int
	// true true
	// 42 42 2a 52 101010
	// 3.141592653589793 3.141592653589793 3.14 (  3.14) 3.141593e+00
	// (110.7+22.5i) (110.7+22.5i) (110.70+22.50i) (1.11e+02+2.25e+01i)
	// 128512 128512 😀 '😀' U+1F600 U+1F600 '😀'
	// foo "bar" foo "bar" "foo \"bar\"" `foo "bar"`
	// map[dachshund:false peanut:true] map[string]bool{"dachshund":false, "peanut":true}
	// {Kim 22} {Name:Kim Age:22} struct { Name string; Age int }{Name:"Kim", Age:22}
	// &{Kim 22} 0x0
	// [Kitano Kobayashi Kurosawa Miyazaki Ozu] ["Kitano" "Kobayashi" "Kurosawa" "Miyazaki" "Ozu"]
	// [Kitano Kobayashi Kurosawa] ["Kitano" "Kobayashi" "Kurosawa"] []string{"Kitano", "Kobayashi", "Kurosawa"}
	// [97 226 140 152] [97 226 140 152] a⌘ "a⌘" 61e28c98 61 e2 8c 98
	// 1973-11-29 21:33:09 +0000 UTC "1973-11-29 21:33:09 +0000 UTC"
}
