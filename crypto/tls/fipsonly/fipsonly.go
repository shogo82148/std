// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto

// fipsonlyパッケージは、すべてのTLS設定をFIPS承認済みの設定に制限します。
//
// この効果は、次のようにプログラム内のどこかでパッケージをインポートすることで発動します:
//
//	import _ "crypto/tls/fipsonly"
//
// このパッケージは、GOEXPERIMENT=boringcryptoでコンパイルされたGoを使用する場合にのみ存在します。
package fipsonly
