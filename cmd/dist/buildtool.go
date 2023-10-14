// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Goブートストラップバージョンを使用してツールチェーンをビルドします。
//
// 一般的な戦略は、必要なソースファイルを新しいGOPATHワークスペースにコピーし、
// 適切にインポートパスを調整して、
// Goブートストラップツールチェーンのコマンドを使用してそれらのソースをビルドし、
// 次にバイナリをコピーすることです。

package main
