// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

// サンプルは単一のメトリックサンプルをキャプチャします。
type Sample struct {

	// Nameはサンプリングされたメトリクスの名前です。
	//
	// これは、Allによって返されるメトリクスの説明の1つの名前と対応する必要があります。
	Name string

	// Valueはメトリックサンプルの値です。
	Value Value
}

// Readは与えられたメトリックサンプルのスライスの各Valueフィールドを埋めます。
// 望ましいメトリックは、適切な名前でスライスに存在している必要があります。
// このAPIのユーザーは、効率化のために同じスライスを呼び出しの間で再利用することを推奨されますが、必須ではありません。
// 再利用にはいくつかの注意点があります。特に、Valueが優れている間にその値を読み取ったり操作したりしてはいけません。それはデータ競合です。
// この特性は、Pointer型のValue（たとえば、Float64Histogramなど）にも含まれます。その基礎となるストレージは、可能な限りReadによって再利用されます。
// このような値を並行設定で安全に使用するには、すべてのデータをディープコピーする必要があります。
// 複数のRead呼び出しを同時に実行することは安全ですが、その引数は共有するアンダーライイングメモリを持っていてはいけません。
// 疑わしい場合は、常に安全な新しい[]Sampleを作成してください。ただし、効率は悪いかもしれません。
// Allに表示されない名前のサンプル値は、その名前が不明であることを示すためにKindBadとしてValueが埋められます。
func Read(m []Sample)
