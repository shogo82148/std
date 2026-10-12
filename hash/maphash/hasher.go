// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package maphash

// Hasherは、ハッシュベースのコンテナーとその要素の間のインターフェースを定義します。
// これは型Tの値に対するハッシュ関数と同値関係を提供し、それらの値を
// ハッシュテーブルや同様のデータ構造に挿入できるようにします。
//
// もちろん、比較可能な型はすでにGoの組み込みのmap型のキーとして使用できますが、
// Hasherを使用すると、比較できない型も適切なハッシュテーブルのキーとして使用できます。
// Hasherは比較可能な型に対しても、通常の同値関係（==）とは異なる同値関係を
// 定義するのに役立つ場合があります。例えば、構造体へのポインター型に対するフィールドベースの比較や、
// 次の例のような文字列の大文字と小文字を区別しない比較などです:
//
//	// CaseInsensitive is a Hasher[string] whose
//	// equivalence relation ignores letter case.
//	type CaseInsensitive struct{}
//
//	func (CaseInsensitive) Hash(h *Hash, s string) {
//		h.WriteString(strings.ToLower(s))
//	}
//
//	func (CaseInsensitive) Equal(x, y string) bool {
//		// (We avoid strings.EqualFold as it is not
//		// consistent with ToLower for all values.)
//		return strings.ToLower(x) == strings.ToLower(y)
//	}
//
// Hasherを使用すると、ブルームフィルターなどの他のハッシュベースのデータ構造でも
// 値を使用できます。
// [ComparableHasher] 型を使用すると、比較可能な型を通常の（==）同値関係のもとで
// そのようなデータ構造で簡単に使用できるようになります。
//
// # Hash invariants
//
// 2つの値がEqual(x, y)で定義されるとおりに等しい場合、それらは
// Hash(h, x)がhに及ぼす効果で定義されるとおりに同じハッシュを持たなければなりません。
//
// Hasherは論理的にステートレスでなければなりません。HashメソッドとEqualメソッドの
// 動作は引数のみに依存します。
//
// # Writing a good function
//
// データ型に対するハッシュ関数と同値関係を定義する場合、まずその型の値を
// 要素の並びとして表す正規のエンコーディングを定義すると役立つ場合があります。
// 各要素は数値、文字列、真偽値、またはポインターです。
// 論理的に等しい2つの値が、異なる方法で表現されていても同じエンコーディングを持つ場合、
// そのエンコーディングは正規です。
// 例えば、文字列の大文字と小文字を区別しない正規のエンコーディングは [strings.ToLower] です。
//
// エンコーディングを定義したら、HasherのHashメソッドは、バイトスライスには [Hash.Write]、
// 文字列には [Hash.WriteString]、バイトには [Hash.WriteByte]、その他の型の要素には
// [WriteComparable] の呼び出しの並びを使用して、値を [Hash] にエンコードするべきです。
// HasherのEqualメソッドは、2つの値のエンコーディングを計算し、対応する要素を比較して、
// 最初に不一致があった時点でfalseを返すべきです。
//
// Hashメソッドは、上で定義したEqualメソッドとの一貫性が保たれている限り、
// 情報を破棄してもかまいません。
// 例えば、CaseInsensitive.Hashの有効な実装は、文字列の最初の文字だけを調べたり、
// 定数値を使用したりすることさえできます。
// ただし、ハッシュ関数の情報の損失が大きいほど、ハッシュの衝突が頻繁になり、
// ハッシュテーブルは遅くなります。
//
// 集合などの一部のデータ型は本質的に順序を持ちません。集合{a, b, c}は
// 集合{c, b, a}と等しいです。
// 場合によっては、要素を何らかの順序でソートすることで、集合の正規のエンコーディングを
// 定義することができます。
// それ以外の場合、メモリーの割り当てが必要になるため非効率的であったり、
// 便利な順序がない場合のように実現不可能であったりします。
// 順序のない集合をハッシュするもう1つの方法は、各要素のハッシュを個別に計算し、
// +や^などの可換（順序に依存しない）演算子を使用してすべての要素のハッシュを
// 結合することです。
//
// 次の仮想的なSet型のHashメソッドは、このアプローチを示しています:
//
//	type Set[T comparable] struct{ ... }
//
//	type setHasher[T comparable] struct{}
//
//	func (setHasher[T]) Hash(hash *maphash.Hash, set *Set[T]) {
//		var accum uint64
//		for elem := range set.Elements() {
//			// Initialize a hasher for the element,
//			// using same seed as the outer hash.
//			var sub maphash.Hash
//			sub.SetSeed(hash.Seed())
//
//			// Hash the element.
//			maphash.WriteComparable(&sub, elem)
//
//			// Mix the element's hash into the set's hash.
//			accum ^= sub.Sum64()
//		}
//		maphash.WriteComparable(hash, accum)
//	}
//
// 多くの言語では、データ型のハッシュ操作は単に整数値を返します。
// しかし、それでは攻撃者が同じハッシュを持つ多数の値を体系的に構築し、
// 「ハッシュフラッディング」として知られるサービス拒否攻撃でハッシュテーブルの
// 漸近的な性能を低下させることが可能になります。
// 対照的に、ハッシュテーブルごとに異なる予測不可能な [Seed] を持つ [Hash] に
// 値の並びを出力してハッシュを計算することで、この攻撃を緩和できます。
//
// 実質的に、Seedは2⁶⁴個の異なるハッシュ関数の中から1つを選択します。
// 上記のコード例では、要素のサブHasherでSetSeedを呼び出し、ランダムなものではなく
// Set自体と同じハッシュ関数を使用するようにしています。
type Hasher[T any] interface {
	Hash(*Hash, T)
	Equal(x, y T) bool
}

// ComparableHasherは、Equal(x, y)メソッドがx == yと一貫性を持つ
// [Hasher] の実装です。
//
// ComparableHasherは比較可能な型に対してのみ定義されます。
// 型システムはComparableHasher[any] のような型のインスタンス化を妨げませんが、
// それでもHashメソッドやEqualメソッドに比較できない引数の値を渡してはいけません。
type ComparableHasher[T comparable] struct {
	_ [0]func(T)
}

func (ComparableHasher[T]) Hash(h *Hash, v T)
func (ComparableHasher[T]) Equal(x, y T) bool
