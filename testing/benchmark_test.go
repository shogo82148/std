// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing_test

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/sort"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/testing"
	"github.com/shogo82148/std/text/template"
)

func ExampleB_RunParallel() {
	// 1つのオブジェクトに対してtext/template.Template.Executeの並列ベンチマーク。
	testing.Benchmark(func(b *testing.B) {
		templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))

		// RunParallelはGOMAXPROCSのゴルーチンを作成し、
		// それらに作業を分散させます。
		b.RunParallel(func(pb *testing.PB) {
			// 各goroutineは独自のbytes.Bufferを持っています。
			var buf bytes.Buffer
			for pb.Next() {
				// ループ本体はすべてのゴルーチンを通じて合計で b.N 回実行されます。
				buf.Reset()
				templ.Execute(&buf, "World")
			}
		})
	})
}

func ExampleB_ReportMetric() {
	// これは特定のアルゴリズム（この場合はソート）に関連するカスタムベンチマークメトリックを報告します。
	testing.Benchmark(func(b *testing.B) {
		var compares int64
		for i := 0; i < b.N; i++ {
			s := []int{5, 4, 3, 2, 1}
			sort.Slice(s, func(i, j int) bool {
				compares++
				return s[i] < s[j]
			})
		}
		// このメトリックは操作ごとのものなので、b.Nで割り、"/op"単位で報告してください。
		b.ReportMetric(float64(compares)/float64(b.N), "compares/op")

		// このメトリックは時間当たりの値ですので、b.Elapsed で割り、
		// "/ns" 単位として報告してください。
		b.ReportMetric(float64(compares)/float64(b.Elapsed().Nanoseconds()), "compares/ns")
	})
}

func ExampleB_ReportMetric_parallel() {
	// これは特定のアルゴリズム（この場合はソート）に関連するカスタムベンチマークメトリックを並列で報告します。
	testing.Benchmark(func(b *testing.B) {
		var compares atomic.Int64
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				s := []int{5, 4, 3, 2, 1}
				sort.Slice(s, func(i, j int) bool {
					// RunParallelは関数を並列で多くの回数実行するため、競合する書き込みを避けるためにカウンターを原子的にインクリメントする必要があります。
					compares.Add(1)
					return s[i] < s[j]
				})
			}
		})

		// 注意：すべての並列呼び出しが完了した後に、各メトリックを1回だけ報告してください。

		// このメトリックは操作ごとに計測されるため、b.Nで割り、"/op"単位で報告してください。
		b.ReportMetric(float64(compares.Load())/float64(b.N), "compares/op")

		// このメトリックは時間に対してのものなので、b.Elapsedで割り算して、"/ns"の単位で報告します。
		b.ReportMetric(float64(compares.Load())/float64(b.Elapsed().Nanoseconds()), "compares/ns")
	})
}
