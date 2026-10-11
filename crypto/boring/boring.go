// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto

// boringパッケージは、Go+BoringCryptoでビルドする場合にのみ利用可能な関数を公開します。
// このパッケージは、Go+BoringCryptoツールチェーンが使用されている限り、
// すべてのターゲットで利用可能です。BoringCryptoコアが実際に使用されているかどうかを
// 判断するにはEnabled関数を使用してください。
//
// Go+BoringCryptoツールチェーンが使用される場合は常に"boringcrypto"ビルドタグが
// 満たされるため、アプリケーションはこのパッケージを使用するファイルにタグを付けることができます。
package boring

// Enabledは、サポートされている暗号操作をBoringCryptoが処理するかどうかを報告します。
func Enabled() bool
