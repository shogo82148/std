// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// backtrack（バックトラック）は、小さな正規表現とテキストに対して、サブマッチの追跡を行う正規表現検索です。それは、(入力の長さ) * (プログラムの長さ)ビットのビットベクトルを割り当てて、同じ（文字位置、命令）の状態を複数回探索しないようにします。これにより、テストの長さに比例して実行時間が線形に制限されます。
//
// backtrackは、onepassを使用できない場合に、小さな正規表現のNFAコードの高速な代替手段です。

package regexp
