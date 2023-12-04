// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates decoding a JPEG image and examining its pixels.
package image_test

import (
	"github.com/shogo82148/std/encoding/base64"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/strings"
)

func Example_decodeConfig() {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	config, format, err := image.DecodeConfig(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Width:", config.Width, "Height:", config.Height, "Format:", format)
}

func Example() {
	// JPEGデータをデコードします。ファイルから読み取る場合は、リーダーを作成します
	//
	// reader, err := os.Open("testdata/video-001.q50.420.jpeg")
	// if err != nil {
	//     log.Fatal(err)
	// }
	// defer reader.Close()
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data))
	m, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := m.Bounds()

	// mの赤、緑、青、アルファ成分の16ビンヒストグラムを計算します。
	//
	// 画像の境界は必ずしも(0, 0)から始まるわけではないので、2つのループは
	// bounds.Min.Yとbounds.Min.Xから始まります。Yを最初に、Xを次にループする方が、
	// Xを最初に、Yを次にループするよりも、より良いメモリアクセスパターンになる可能性が高いです。
	var histogram [16][4]int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := m.At(x, y).RGBA()
			// 色のRGBAメソッドは、範囲[0, 65535]の値を返します。
			// これを12ビット右シフトすると、範囲が[0, 15]に縮小されます。
			histogram[r>>12][0]++
			histogram[g>>12][1]++
			histogram[b>>12][2]++
			histogram[a>>12][3]++
		}
	}

	// 結果を表示します。
	fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
	for i, x := range histogram {
		fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
	}
	// Output:
	// bin               red  green   blue  alpha
	// 0x0000-0x0fff:    364    790   7242      0
	// 0x1000-0x1fff:    645   2967   1039      0
	// 0x2000-0x2fff:   1072   2299    979      0
	// 0x3000-0x3fff:    820   2266    980      0
	// 0x4000-0x4fff:    537   1305    541      0
	// 0x5000-0x5fff:    319    962    261      0
	// 0x6000-0x6fff:    322    375    177      0
	// 0x7000-0x7fff:    601    279    214      0
	// 0x8000-0x8fff:   3478    227    273      0
	// 0x9000-0x9fff:   2260    234    329      0
	// 0xa000-0xafff:    921    282    373      0
	// 0xb000-0xbfff:    321    335    397      0
	// 0xc000-0xcfff:    229    388    298      0
	// 0xd000-0xdfff:    260    414    277      0
	// 0xe000-0xefff:    516    428    298      0
	// 0xf000-0xffff:   2785   1899   1772  15450
}
