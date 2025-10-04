// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
iterパッケージは、シーケンス上のイテレータに関連する基本的な定義と操作を提供します。

# Iterators

イテレータは、シーケンスの連続する要素をコールバック関数（通常はyieldと名付けられる）に渡す関数です。
この関数は、シーケンスが終了するか、yieldがfalseを返して早期にイテレーションを停止するよう指示するまで動作します。
このパッケージは、シーケンス要素ごとに1つまたは2つの値をyieldに渡すイテレータの省略形として、
[Seq] および [Seq2]（シーケンスの最初の音節のように発音される）を定義します。

	type (
		Seq[V any]     func(yield func(V) bool)
		Seq2[K, V any] func(yield func(K, V) bool)
	)

Seq2は、通常はキーと値のペアやインデックスと値のペアを表すシーケンスです。

yieldは、イテレータがシーケンスの次の要素を続行すべき場合にtrueを返し、
停止すべき場合にfalseを返します。

For instance, [maps.Keys] returns an iterator that produces the sequence
of keys of the map m, implemented as follows:

	func Keys[Map ~map[K]V, K comparable, V any](m Map) iter.Seq[K] {
		return func(yield func(K) bool) {
			for k := range m {
				if !yield(k) {
					return
				}
			}
		}
	}

さらなる例は [The Go Blog: Range Over Function Types] で見つけることができます。

イテレータ関数は最も一般的に [range loop] によって呼び出されます。例えば：

	func PrintAll[V any](seq iter.Seq[V]) {
		for v := range seq {
			fmt.Println(v)
		}
	}

# Naming Conventions

イテレータ関数とメソッドは、処理されるシーケンスにちなんで名付けられています：

	// Allは、s内のすべての要素を反復するイテレータを返します。
	func (s *Set[V]) All() iter.Seq[V]

コレクション型のイテレータメソッドは、慣例的にAllと名付けられています。
これは、コレクション内のすべての値のシーケンスを反復するためです。

複数のシーケンスを含む型の場合、イテレータの名前は提供されるシーケンスを示すことができます：

	// Citiesは、その国の主要都市を反復するイテレータを返します。
	func (c *Country) Cities() iter.Seq[*City]

	// Languagesは、その国の公用語を反復するイテレータを返します。
	func (c *Country) Languages() iter.Seq[string]

イテレータが追加の設定を必要とする場合、コンストラクタ関数は追加の設定引数を取ることができます：

	// Scanは、min ≤ key ≤ maxのキーと値のペアを反復するイテレータを返します。
	func (m *Map[K, V]) Scan(min, max K) iter.Seq2[K, V]

	// Splitは、sepで区切られたsの（空である可能性のある）部分文字列のイテレータを返します。
	func Split(s, sep string) iter.Seq[string]

複数の反復順序が可能な場合、メソッド名はその順序を示すことがあります：

	// Allは、リストの先頭から末尾までのイテレータを返します。
	func (l *List[V]) All() iter.Seq[V]

	// Backwardは、リストの末尾から先頭までのイテレータを返します。
	func (l *List[V]) Backward() iter.Seq[V]

	// Preorderは、指定されたルートを含む構文木のすべてのノードを
	// 深さ優先の前順で反復するイテレータを返します。
	// 親ノードをその子ノードの前に訪問します。
	func Preorder(root Node) iter.Seq[Node]

# Single-Use Iterators

ほとんどのイテレータは、シーケンス全体を歩く機能を提供します：
呼び出されると、イテレータはシーケンスを開始するために必要なセットアップを行い、
次にシーケンスの連続する要素に対してyieldを呼び出し、
最後に戻る前にクリーンアップを行います。
イテレータを再度呼び出すと、シーケンスを再度歩きます。

一部のイテレータはその慣例を破り、シーケンスを一度だけ歩く機能を提供します。
これらの「単一使用イテレータ」は、通常、最初からやり直すことができないデータストリームから値を報告します。
途中で停止した後にイテレータを再度呼び出すと、ストリームが続行される場合がありますが、
シーケンスが終了した後に再度呼び出しても、値は一切返されません。
単一使用イテレータを返す関数やメソッドのドキュメントコメントには、この事実を記載する必要があります：

	// Linesは、rから読み取った行を反復するイテレータを返します。
	// これは単一使用のイテレータを返します。
	func (r *Reader) Lines() iter.Seq[string]

# Pulling Values

イテレータを受け取る、または返す関数やメソッドは、標準の [Seq] または [Seq2] 型を使用して、
rangeループや他のイテレータアダプタとの互換性を確保する必要があります。
標準のイテレータは「プッシュイテレータ」と考えることができ、
値をyield関数にプッシュします。

rangeループがシーケンスの値を消費する最も自然な方法ではない場合があります。
この場合、[Pull] は標準のプッシュイテレータを「プルイテレータ」に変換し、
シーケンスから一度に1つの値をプルするために呼び出すことができます。
[Pull] はイテレータを開始し、イテレータから次の値を返す関数nextと、
イテレータを停止する関数stopのペアを返します。

たとえば:

	// Pairsは、seqから連続する値のペアを反復するイテレータを返します。
	func Pairs[V any](seq iter.Seq[V]) iter.Seq2[V, V] {
		return func(yield func(V, V) bool) {
			next, stop := iter.Pull(seq)
			defer stop()
			for {
				v1, ok1 := next()
				if !ok1 {
					return
				}
				v2, ok2 := next()
				// If ok2 is false, v2 should be the
				// zero value; yield one last pair.
				if !yield(v1, v2) {
					return
				}
				if !ok2 {
					return
				}
			}
		}
	}

クライアントがシーケンスを最後まで消費しない場合、イテレータ関数が終了して戻ることができるように、
stopを呼び出す必要があります。例に示すように、これを確実に行うための一般的な方法はdeferを使用することです。

# Standard Library Usage

標準ライブラリのいくつかのパッケージは、イテレータベースのAPIを提供しています。
特に [maps] および [slices] パッケージがそれに該当します。
例えば、[maps.Keys] はマップのキーを反復するイテレータを返し、
[slices.Sorted] はイテレータの値をスライスに収集し、それをソートしてスライスを返します。
したがって、マップのソートされたキーを反復するには次のようにします：

	for _, key := range slices.Sorted(maps.Keys(m)) {
		...
	}

# Mutation

イテレータはシーケンスの値のみを提供し、それを直接変更する方法は提供しません。
イテレータがイテレーション中にシーケンスを変更するメカニズムを提供したい場合、
通常のアプローチは追加の操作を持つ位置型を定義し、
その位置を反復するイテレータを提供することです。

例えば、ツリーの実装は次のように提供されるかもしれません：

	// Positionsは、シーケンス内の位置のイテレータを返します。
	func (t *Tree[V]) Positions() iter.Seq[*Pos[V]]

	// Posはシーケンス内の位置を表します。
	// これは、それが渡されるyield呼び出しの間のみ有効です。
	type Pos[V any] struct { ... }

	// Posはカーソルの位置にある値を返します。
	func (p *Pos[V]) Value() V

	// Deleteは、イテレーションのこの時点での値を削除します。
	func (p *Pos[V]) Delete()

	// Setは、カーソルの位置にある値をvに変更します。
	func (p *Pos[V]) Set(v V)

そして、クライアントは次のようにしてツリーから退屈な値を削除できます：

	for p := range t.Positions() {
		if boring(p.Value()) {
			p.Delete()
		}
	}

[The Go Blog: Range Over Function Types]: https://go.dev/blog/range-functions
[range loop]: https://go.dev/ref/spec#For_range
*/
package iter

// Seqは個々の値のシーケンスを反復するイテレータです。
// seq(yield)として呼び出されると、seqはシーケンス内の各値vに対してyield(v)を呼び出し、
// yieldがfalseを返した場合は早期に停止します。
// 詳細については、[iter] パッケージのドキュメントを参照してください。
type Seq[V any] func(yield func(V) bool)

// Seq2は、主にキーと値のペアである値のペアのシーケンスを反復するイテレータです。
// seq(yield)として呼び出されると、seqはシーケンス内の各ペア(k, v)に対してyield(k, v)を呼び出し、
// yieldがfalseを返した場合は早期に停止します。
// 詳細については、[iter] パッケージのドキュメントを参照してください。
type Seq2[K, V any] func(yield func(K, V) bool)

// Pullは、「プッシュスタイル」のイテレータシーケンスseqを、
// 2つの関数nextとstopによってアクセスされる「プルスタイル」のイテレータに変換します。
//
// Nextはシーケンス内の次の値と、その値が有効かどうかを示すブール値を返します。
// シーケンスが終了した場合、nextはゼロ値のVとfalseを返します。
// シーケンスの終わりに達した後やstopを呼び出した後にnextを呼び出すことは有効です。
// これらの呼び出しは引き続きゼロ値のVとfalseを返します。
//
// Stopはイテレーションを終了します。呼び出し元が次の値に興味がなくなり、
// nextがまだシーケンスの終了を示していない（falseのブール値を返していない）場合に
// 呼び出す必要があります。stopを複数回呼び出すことや、nextがすでにfalseを返した後に
// 呼び出すことは有効です。通常、呼び出し元は「defer stop()」を使用するべきです。
//
// nextまたはstopを複数のゴルーチンから同時に呼び出すことはエラーです。
//
// イテレータがnext（またはstop）の呼び出し中にパニックを起こした場合、
// next（またはstop）自体も同じ値でパニックを起こします。
func Pull[V any](seq Seq[V]) (next func() (V, bool), stop func())

// Pull2は、「プッシュスタイル」のイテレータシーケンスseqを、
// 2つの関数nextとstopによってアクセスされる「プルスタイル」のイテレータに変換します。
//
// Nextはシーケンス内の次のペアと、そのペアが有効かどうかを示すブール値を返します。
// シーケンスが終了した場合、nextはゼロ値のペアとfalseを返します。
// シーケンスの終わりに達した後やstopを呼び出した後にnextを呼び出すことは有効です。
// これらの呼び出しは引き続きゼロ値のペアとfalseを返します。
//
// Stopはイテレーションを終了します。呼び出し元が次の値に興味がなくなり、
// nextがまだシーケンスの終了を示していない（falseのブール値を返していない）場合に
// 呼び出す必要があります。stopを複数回呼び出すことや、nextがすでにfalseを返した後に
// 呼び出すことは有効です。通常、呼び出し元は「defer stop()」を使用するべきです。
//
// nextまたはstopを複数のゴルーチンから同時に呼び出すことはエラーです。
//
// イテレータがnext（またはstop）の呼び出し中にパニックを起こした場合、
// next（またはstop）自体も同じ値でパニックを起こします。
func Pull2[K, V any](seq Seq2[K, V]) (next func() (K, V, bool), stop func())
