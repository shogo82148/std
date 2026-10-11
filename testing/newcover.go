// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Support for test coverage with redesigned coverage implementation.

package testing

// Coverageは、現在のコードカバレッジを[0, 1]の範囲の割合として報告します。
// カバレッジが有効になっていない場合、Coverageは0を返します。
//
// 一連の多数のテストケースを順番に実行する場合、各テストケースの後でCoverageを確認すると、
// どのテストケースが新しいコードパスを実行しているかを特定するのに役立ちます。
//
// これは'go test -cover'や'go tool cover'によって生成されるレポートの代わりにはなりません。
func Coverage() float64
