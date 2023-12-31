// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package png_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/image"
	"github.com/shogo82148/std/image/color"
	"github.com/shogo82148/std/image/png"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/os"
)

func ExampleDecode() {
	// この例では、PNG画像のみをデコードできるpng.Decodeを使用しています。
	// 任意の登録済み画像フォーマットをスニフし、デコードできる一般的なimage.Decodeの使用を検討してください。
	img, err := png.Decode(gopherPNG())
	if err != nil {
		log.Fatal(err)
	}

	levels := []string{" ", "░", "▒", "▓", "█"}

	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			c := color.GrayModel.Convert(img.At(x, y)).(color.Gray)
			level := c.Y / 51 // 51 * 5 = 255
			if level == 5 {
				level--
			}
			fmt.Print(levels[level])
		}
		fmt.Print("\n")
	}
}

func ExampleEncode() {
	const width, height = 256, 256

	// 与えられた幅と高さのカラー画像を作成します。
	img := image.NewNRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			img.Set(x, y, color.NRGBA{
				R: uint8((x + y) & 255),
				G: uint8((x + y) << 1 & 255),
				B: uint8((x + y) << 2 & 255),
				A: 255,
			})
		}
	}

	f, err := os.Create("image.png")
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, img); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
