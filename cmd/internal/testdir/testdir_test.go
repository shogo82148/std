// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// GOROOT/test 内の個々のテストケースは、必要に応じて -race フラグなどのフラグを
// 自分で設定する責任を負っており、そのフラグは個々のテストケースには
// そもそも自動的に伝播しません。
//
// そのため、重複した作業を避けるために、これらのテストをレースモードで繰り返さないようにします。
// 特に、レースビルダーでの時間を節約できます。
//
//go:build !race

// Package testdir_test runs tests in the GOROOT/test directory.
package testdir_test
