// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/encoding/json"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/reflect"
)

func ExampleKind() {
	for _, v := range []any{"hi", 42, func() {}} {
		switch v := reflect.ValueOf(v); v.Kind() {
		case reflect.String:
			fmt.Println(v.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			fmt.Println(v.Int())
		default:
			fmt.Printf("unhandled kind %s", v.Kind())
		}
	}

	// Output:
	// hi
	// 42
	// unhandled kind func
}

func ExampleMakeFunc() {
	// swapはMakeFuncに渡される実装です。
	// 事前に型を知らなくてもコードを書くことができるように、
	// reflect.Valuesを使用して動作する必要があります。
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}

	// makeSwapは、fptrがnil関数へのポインタであることを期待しています。
	// それは、MakeFuncで作成された新しい関数にそのポインタを設定します。
	// 関数が呼び出されると、reflectは引数をValuesに変換し、swapを呼び出し、swapの結果スライスを新しい関数の返り値として変換します。
	makeSwap := func(fptr any) {
		// fptrは関数へのポインタです。
		// リフレクト値として関数値自体（おそらくnil）を取得し、
		// その型をクエリできるようにするために値を設定します。
		fn := reflect.ValueOf(fptr).Elem()

		// 適切な型の関数を作成します。
		v := reflect.MakeFunc(fn.Type(), swap)

		// fnを表す値に割り当てる。
		fn.Set(v)
	}

	// int型のためのswap関数を作成して呼び出す。
	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	fmt.Println(intSwap(0, 1))

	// float64型のスワップ関数を作成して呼び出す。
	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14))

	// Output:
	// 1 0
	// 3.14 2.72
}

func ExampleStructTag() {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

	// Output:
	// blue gopher
}

func ExampleStructTag_Lookup() {
	type S struct {
		F0 string `alias:"field_0"`
		F1 string `alias:""`
		F2 string
	}

	s := S{}
	st := reflect.TypeOf(s)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if alias, ok := field.Tag.Lookup("alias"); ok {
			if alias == "" {
				fmt.Println("(blank)")
			} else {
				fmt.Println(alias)
			}
		} else {
			fmt.Println("(not specified)")
		}
	}

	// Output:
	// field_0
	// (blank)
	// (not specified)
}

func ExampleTypeOf() {
	// インターフェース型は静的な型付けにしか使用されませんので、
	// インターフェース型Fooの反射Typeを見つけるための一般的な方法は、*Fooの値を使用するというものです。
	writerType := reflect.TypeOf((*io.Writer)(nil)).Elem()

	fileType := reflect.TypeOf((*os.File)(nil))
	fmt.Println(fileType.Implements(writerType))

	// Output:
	// true
}

func ExampleStructOf() {
	typ := reflect.StructOf([]reflect.StructField{
		{
			Name: "Height",
			Type: reflect.TypeOf(float64(0)),
			Tag:  `json:"height"`,
		},
		{
			Name: "Age",
			Type: reflect.TypeOf(int(0)),
			Tag:  `json:"age"`,
		},
	})

	v := reflect.New(typ).Elem()
	v.Field(0).SetFloat(0.4)
	v.Field(1).SetInt(2)
	s := v.Addr().Interface()

	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}

	fmt.Printf("value: %+v\n", s)
	fmt.Printf("json:  %s", w.Bytes())

	r := bytes.NewReader([]byte(`{"height":1.5,"age":10}`))
	if err := json.NewDecoder(r).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)

	// Output:
	// value: &{Height:0.4 Age:2}
	// json:  {"height":0.4,"age":2}
	// value: &{Height:1.5 Age:10}
}

func ExampleValue_FieldByIndex() {
	// この例は、昇格されたフィールドの名前がその他のフィールドによって隠される場合のケースを示しています。FieldByNameは機能しないため、代わりにFieldByIndexを使用する必要があります。
	type user struct {
		firstName string
		lastName  string
	}

	type data struct {
		user
		firstName string
		lastName  string
	}

	u := data{
		user:      user{"Embedded John", "Embedded Doe"},
		firstName: "John",
		lastName:  "Doe",
	}

	s := reflect.ValueOf(u).FieldByIndex([]int{0, 1})
	fmt.Println("embedded last name:", s)

	// Output:
	// embedded last name: Embedded Doe
}

func ExampleValue_FieldByName() {
	type user struct {
		firstName string
		lastName  string
	}
	u := user{firstName: "John", lastName: "Doe"}
	s := reflect.ValueOf(u)

	fmt.Println("Name:", s.FieldByName("firstName"))
	// Output:
	// Name: John
}
