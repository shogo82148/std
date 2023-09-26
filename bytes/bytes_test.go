// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bytes_test

import (
	. "bytes"
)

type BinOpTest struct {
	a string
	b string
	i int
}

type SplitTest struct {
	s   string
	sep string
	n   int
	a   []string
}

type FieldsTest struct {
	s string
	a []string
}

// Test case for any function which accepts and returns a byte slice.
// For ease of creation, we write the input byte slice as a string.
type StringTest struct {
	in  string
	out []byte
}

type RepeatTest struct {
	in, out string
	count   int
}

var RepeatTests = []RepeatTest{
	{"", "", 0},
	{"", "", 1},
	{"", "", 2},
	{"-", "", 0},
	{"-", "-", 1},
	{"-", "----------", 10},
	{"abc ", "abc abc abc ", 3},
}

type RunesTest struct {
	in    string
	out   []rune
	lossy bool
}

var RunesTests = []RunesTest{
	{"", []rune{}, false},
	{" ", []rune{32}, false},
	{"ABC", []rune{65, 66, 67}, false},
	{"abc", []rune{97, 98, 99}, false},
	{"\u65e5\u672c\u8a9e", []rune{26085, 26412, 35486}, false},
	{"ab\x80c", []rune{97, 98, 0xFFFD, 99}, true},
	{"ab\xc0c", []rune{97, 98, 0xFFFD, 99}, true},
}

type TrimTest struct {
	f            string
	in, arg, out string
}

type TrimNilTest struct {
	f   string
	in  []byte
	arg string
	out []byte
}

type TrimFuncTest struct {
	f        predicate
	in       string
	trimOut  []byte
	leftOut  []byte
	rightOut []byte
}

type IndexFuncTest struct {
	in          string
	f           predicate
	first, last int
}

type ReplaceTest struct {
	in       string
	old, new string
	n        int
	out      string
}

var ReplaceTests = []ReplaceTest{
	{"hello", "l", "L", 0, "hello"},
	{"hello", "l", "L", -1, "heLLo"},
	{"hello", "x", "X", -1, "hello"},
	{"", "x", "X", -1, ""},
	{"radar", "r", "<r>", -1, "<r>ada<r>"},
	{"", "", "<>", -1, "<>"},
	{"banana", "a", "<>", -1, "b<>n<>n<>"},
	{"banana", "a", "<>", 1, "b<>nana"},
	{"banana", "a", "<>", 1000, "b<>n<>n<>"},
	{"banana", "an", "<>", -1, "b<><>a"},
	{"banana", "ana", "<>", -1, "b<>na"},
	{"banana", "", "<>", -1, "<>b<>a<>n<>a<>n<>a<>"},
	{"banana", "", "<>", 10, "<>b<>a<>n<>a<>n<>a<>"},
	{"banana", "", "<>", 6, "<>b<>a<>n<>a<>n<>a"},
	{"banana", "", "<>", 5, "<>b<>a<>n<>a<>na"},
	{"banana", "", "<>", 1, "<>banana"},
	{"banana", "a", "a", -1, "banana"},
	{"banana", "a", "a", 1, "banana"},
	{"☺☻☹", "", "<>", -1, "<>☺<>☻<>☹<>"},
}

type TitleTest struct {
	in, out string
}

var TitleTests = []TitleTest{
	{"", ""},
	{"a", "A"},
	{" aaa aaa aaa ", " Aaa Aaa Aaa "},
	{" Aaa Aaa Aaa ", " Aaa Aaa Aaa "},
	{"123a456", "123a456"},
	{"double-blind", "Double-Blind"},
	{"ÿøû", "Ÿøû"},
	{"with_underscore", "With_underscore"},
	{"unicode \xe2\x80\xa8 line separator", "Unicode \xe2\x80\xa8 Line Separator"},
}

var ToTitleTests = []TitleTest{
	{"", ""},
	{"a", "A"},
	{" aaa aaa aaa ", " AAA AAA AAA "},
	{" Aaa Aaa Aaa ", " AAA AAA AAA "},
	{"123a456", "123A456"},
	{"double-blind", "DOUBLE-BLIND"},
	{"ÿøû", "ŸØÛ"},
}

var EqualFoldTests = []struct {
	s, t string
	out  bool
}{
	{"abc", "abc", true},
	{"ABcd", "ABcd", true},
	{"123abc", "123ABC", true},
	{"αβδ", "ΑΒΔ", true},
	{"abc", "xyz", false},
	{"abc", "XYZ", false},
	{"abcdefghijk", "abcdefghijX", false},
	{"abcdefghijk", "abcdefghij\u212A", true},
	{"abcdefghijK", "abcdefghij\u212A", true},
	{"abcdefghijkz", "abcdefghij\u212Ay", false},
	{"abcdefghijKz", "abcdefghij\u212Ay", false},
}

var ContainsAnyTests = []struct {
	b        []byte
	substr   string
	expected bool
}{
	{[]byte(""), "", false},
	{[]byte(""), "a", false},
	{[]byte(""), "abc", false},
	{[]byte("a"), "", false},
	{[]byte("a"), "a", true},
	{[]byte("aaa"), "a", true},
	{[]byte("abc"), "xyz", false},
	{[]byte("abc"), "xcz", true},
	{[]byte("a☺b☻c☹d"), "uvw☻xyz", true},
	{[]byte("aRegExp*"), ".(|)*+?^$[]", true},
	{[]byte(dots + dots + dots), " ", false},
}

var ContainsRuneTests = []struct {
	b        []byte
	r        rune
	expected bool
}{
	{[]byte(""), 'a', false},
	{[]byte("a"), 'a', true},
	{[]byte("aaa"), 'a', true},
	{[]byte("abc"), 'y', false},
	{[]byte("abc"), 'c', true},
	{[]byte("a☺b☻c☹d"), 'x', false},
	{[]byte("a☺b☻c☹d"), '☻', true},
	{[]byte("aRegExp*"), '*', true},
}
