// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

import (
	"time"
)

// hget http://codereview.appspot.com/rss/mine/rsc

type Feed struct {
	XMLName Name      `xml:"http://www.w3.org/2005/Atom feed"`
	Title   string    `xml:"title"`
	Id      string    `xml:"id"`
	Link    []Link    `xml:"link"`
	Updated time.Time `xml:"updated,attr"`
	Author  Person    `xml:"author"`
	Entry   []Entry   `xml:"entry"`
}

type Entry struct {
	Title   string    `xml:"title"`
	Id      string    `xml:"id"`
	Link    []Link    `xml:"link"`
	Updated time.Time `xml:"updated"`
	Author  Person    `xml:"author"`
	Summary Text      `xml:"summary"`
}

type Link struct {
	Rel  string `xml:"rel,attr,omitempty"`
	Href string `xml:"href,attr"`
}

type Person struct {
	Name     string `xml:"name"`
	URI      string `xml:"uri"`
	Email    string `xml:"email"`
	InnerXML string `xml:",innerxml"`
}

type Text struct {
	Type string `xml:"type,attr,omitempty"`
	Body string `xml:",chardata"`
}

type PathTestItem struct {
	Value string
}

type PathTestA struct {
	Items         []PathTestItem `xml:">Item1"`
	Before, After string
}

type PathTestB struct {
	Other         []PathTestItem `xml:"Items>Item1"`
	Before, After string
}

type PathTestC struct {
	Values1       []string `xml:"Items>Item1>Value"`
	Values2       []string `xml:"Items>Item2>Value"`
	Before, After string
}

type PathTestSet struct {
	Item1 []PathTestItem
}

type PathTestD struct {
	Other         PathTestSet `xml:"Items"`
	Before, After string
}

type PathTestE struct {
	Underline     string `xml:"Items>_>Value"`
	Before, After string
}

type BadPathTestA struct {
	First  string `xml:"items>item1"`
	Other  string `xml:"items>item2"`
	Second string `xml:"items"`
}

type BadPathTestB struct {
	Other  string `xml:"items>item2>value"`
	First  string `xml:"items>item1"`
	Second string `xml:"items>item1>value"`
}

type BadPathTestC struct {
	First  string
	Second string `xml:"First"`
}

type BadPathTestD struct {
	BadPathEmbeddedA
	BadPathEmbeddedB
}

type BadPathEmbeddedA struct {
	First string
}

type BadPathEmbeddedB struct {
	Second string `xml:"First"`
}

const OK = "OK"

type TestThree struct {
	XMLName Name   `xml:"Test3"`
	Attr    string `xml:",attr"`
}

type Tables struct {
	HTable string `xml:"http://www.w3.org/TR/html4/ table"`
	FTable string `xml:"http://www.w3schools.com/furniture table"`
}

type TableAttrs struct {
	TAttr TAttr
}

type TAttr struct {
	HTable string `xml:"http://www.w3.org/TR/html4/ table,attr"`
	FTable string `xml:"http://www.w3schools.com/furniture table,attr"`
	Lang   string `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
	Other1 string `xml:"http://golang.org/xml/ other,attr,omitempty"`
	Other2 string `xml:"http://golang.org/xmlfoo/ other,attr,omitempty"`
	Other3 string `xml:"http://golang.org/json/ other,attr,omitempty"`
	Other4 string `xml:"http://golang.org/2/json/ other,attr,omitempty"`
}

type MyCharData struct {
	body string
}

var _ Unmarshaler = (*MyCharData)(nil)

type MyAttr struct {
	attr string
}

var _ UnmarshalerAttr = (*MyAttr)(nil)

type MyStruct struct {
	Data *MyCharData
	Attr *MyAttr `xml:",attr"`

	Data2 MyCharData
	Attr2 MyAttr `xml:",attr"`
}

type Pea struct {
	Cotelydon string
}

type Pod struct {
	Pea interface{} `xml:"Pea"`
}

type X struct {
	D string `xml:",comment"`
}

type IXField struct {
	Five        int      `xml:"five"`
	NotInnerXML []string `xml:",innerxml"`
}

type Child struct {
	G struct {
		I int
	}
}

type ChildToEmbed struct {
	X bool
}

type Parent struct {
	I        int
	IPtr     *int
	Is       []int
	IPtrs    []*int
	F        float32
	FPtr     *float32
	Fs       []float32
	FPtrs    []*float32
	B        bool
	BPtr     *bool
	Bs       []bool
	BPtrs    []*bool
	Bytes    []byte
	BytesPtr *[]byte
	S        string
	SPtr     *string
	Ss       []string
	SPtrs    []*string
	MyI      MyInt
	Child    Child
	Children []Child
	ChildPtr *Child
	ChildToEmbed
}
