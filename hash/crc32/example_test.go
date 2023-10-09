// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package crc32_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/hash/crc32"
)

func ExampleMakeTable() {
	// このパッケージでは、CRCポリノミアルは逆順記法、またはLSB-firstの表現で表されます。
	//
	// LSB-first表現は、nビットの16進数であり、最上位ビットはx⁰の係数を表し、最下位ビットはxⁿ⁻¹（xⁿの係数は暗黙的に表される）の係数を表します。
	//
	// たとえば、以下のポリノミアルによって定義されるCRC32-Qは、次のような逆順記法を持ちます。
	//	x³²+ x³¹+ x²⁴+ x²²+ x¹⁶+ x¹⁴+ x⁸+ x⁷+ x⁵+ x³+ x¹+ x⁰
	// したがって、MakeTableに渡すべき値は0xD5828281です。
	crc32q := crc32.MakeTable(0xD5828281)
	fmt.Printf("%08x\n", crc32.Checksum([]byte("Hello world"), crc32q))
	// Output:
	// 2964d064
}
