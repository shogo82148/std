// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package unify implements unification of structured values.
//
// A [Value] represents a possibly infinite set of concrete values, where a
// value is either a string ([String]), a tuple of values ([Tuple]), or a
// string-keyed map of values called a "def" ([Def]). These sets can be further
// constrained by variables ([Var]). A [Value] combined with bindings of
// variables is a [Closure].
//
// [Unify] finds a [Closure] that satisfies two or more other [Closure]s. This
// can be thought of as intersecting the sets represented by these Closures'
// values, or as the greatest lower bound/infimum of these Closures. If no such
// Closure exists, the result of unification is "bottom", or the empty set.
//
// # Examples
//
// The regular expression "a*" is the infinite set of strings of zero or more
// "a"s. "a*" can be unified with "a" or "aa" or "aaa", and the result is just
// "a", "aa", or "aaa", respectively. However, unifying "a*" with "b" fails
// because there are no values that satisfy both.
//
// Sums express sets directly. For example, !sum [a, b] is the set consisting of
// "a" and "b". Unifying this with !sum [b, c] results in just "b". This also
// makes it easy to demonstrate that unification isn't necessarily a single
// concrete value. For example, unifying !sum [a, b, c] with !sum [b, c, d]
// results in two concrete values: "b" and "c".
//
// The special value _ or "top" represents all possible values. Unifying _ with
// any value x results in x.
//
// Unifying composite values—tuples and defs—unifies their elements.
//
// The value [a*, aa] is an infinite set of tuples. If we unify that with the
// value [aaa, a*], the only possible value that satisfies both is [aaa, aa].
// Likewise, this is the intersection of the sets described by these two values.
//
// Defs are similar to tuples, but they are indexed by strings and don't have a
// fixed length. For example, {x: a, y: b} is a def with two fields. Any field
// not mentioned in a def is implicitly top. Thus, unifying this with {y: b, z:
// c} results in {x: a, y: b, z: c}.
//
// Variables constrain values. For example, the value [$x, $x] represents all
// tuples whose first and second values are the same, but doesn't otherwise
// constrain that value. Thus, this set includes [a, a] as well as [[b, c, d],
// [b, c, d]], but it doesn't include [a, b].
//
// Sums are internally implemented as fresh variables that are simultaneously
// bound to all values of the sum. That is !sum [a, b] is actually $var (where
// var is some fresh name), closed under the environment $var=a | $var=b.
package unify

// Unify computes a Closure that satisfies each input Closure. If no such
// Closure exists, it returns bottom.
func Unify(closures ...Closure) (Closure, error)
