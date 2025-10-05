// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slogtest_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/encoding/json"
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/log/slog"
	"github.com/shogo82148/std/testing/slogtest"
)

// この例は、このパッケージを使用してハンドラをテストする一つの手法を示しています。
// ハンドラには [bytes.Buffer] が書き込み用に与えられ、結果の出力の各行が解析されます。
// JSON出力の場合、[encoding/json.Unmarshal] はmap[string]anyへのポインタを与えると、
// 望ましい形式の結果を生成します。
func Example_parsing() {
	var buf bytes.Buffer
	h := slog.NewJSONHandler(&buf, nil)

	results := func() []map[string]any {
		var ms []map[string]any
		for line := range bytes.SplitSeq(buf.Bytes(), []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}
			var m map[string]any
			if err := json.Unmarshal(line, &m); err != nil {
				panic(err) // 実際のテストでは、t.Fatalを使用します。
			}
			ms = append(ms, m)
		}
		return ms
	}
	err := slogtest.TestHandler(h, results)
	if err != nil {
		log.Fatal(err)
	}

	// Output:
}
