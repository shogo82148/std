// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/runtime/metrics"
)

func ExampleRead_readingOneMetric() {
	// 読み取りたいメトリックの名前。
	const myMetric = "/memory/classes/heap/free:bytes"

	// メトリクスのサンプルを作成します。
	sample := make([]metrics.Sample, 1)
	sample[0].Name = myMetric

	// メトリックをサンプリングする。
	metrics.Read(sample)

	// メトリックが実際にサポートされているか確認します。
	// もしサポートされていなければ、結果の値は常にKindBadになります。
	if sample[0].Value.Kind() == metrics.KindBad {
		panic(fmt.Sprintf("metric %q no longer supported", myMetric))
	}

	// 結果を処理する。
	//
	// メトリックに特定の Kind を想定することは問題ありません。
	// それらは変更されないことが保証されています。
	freeBytes := sample[0].Value.Uint64()

	fmt.Printf("free but not released memory: %d\n", freeBytes)
}

func ExampleRead_readingAllMetrics() {
	// すべてのサポートされているメトリクスの説明を取得する。
	descs := metrics.All()

	// 各メトリックのサンプルを作成します。
	samples := make([]metrics.Sample, len(descs))
	for i := range samples {
		samples[i].Name = descs[i].Name
	}

	// メトリクスをサンプリングします。可能な場合、サンプルスライスを再利用してください！
	metrics.Read(samples)

	// すべての結果を繰り返す。
	for _, sample := range samples {
		// 名前と値を取り出す。
		name, value := sample.Name, sample.Value

		// 各サンプルを処理する。
		switch value.Kind() {
		case metrics.KindUint64:
			fmt.Printf("%s: %d\n", name, value.Uint64())
		case metrics.KindFloat64:
			fmt.Printf("%s: %f\n", name, value.Float64())
		case metrics.KindFloat64Histogram:
			// ヒストグラムはかなり大きくなるかもしれませんので、この例のために
			// 中央値のざっくりした推定値を取り出しましょう。
			fmt.Printf("%s: %f\n", name, medianBucket(value.Float64Histogram()))
		case metrics.KindBad:
			// これは起きてはいけないはずです。なぜなら、すべてのメトリクスは構築されているからです。
			panic("bug in runtime/metrics package!")
		default:

			// 新しいメトリクスが追加されると、これが発生する可能性があります。
			//
			// ここでは、安全策として、単にそれをログに記録しておくことで、
			// 現時点では無視します。
			// 最悪の場合、一時的に新しいメトリクスの情報を見逃すかもしれませんが、
			fmt.Printf("%s: unexpected metric Kind: %v\n", name, value.Kind())
		}
	}
}
