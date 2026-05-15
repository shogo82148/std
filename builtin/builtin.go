// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
builtinパッケージは、Goの事前定義された識別子のドキュメンテーションを提供します。
ここに文書化されている項目は実際にはbuiltinパッケージには存在しませんが、
ここでの説明によりgodocは言語の特別な識別子のドキュメンテーションを提示することができます。
*/
package builtin

// Typeは、ドキュメンテーションの目的のみでここに存在します。それは任意のGo型の代わりで、
// しかし、任意の関数呼び出しに対して同じ型を表します。
type Type int

// Type1は、ドキュメンテーションの目的のみでここに存在します。それは任意のGo型の代わりで、
// しかし、任意の関数呼び出しに対して同じ型を表します。
type Type1 int

// TypeOrExprはドキュメンテーションの目的のためだけにここに存在します。Go型または式のいずれかの代わりです。
type TypeOrExpr int

// IntegerTypeはドキュメンテーションの目的のためだけにここに存在します。任意の整数型（int、uint、int8など）の代わりです。
type IntegerType int

// FloatTypeは、ドキュメンテーションの目的のみでここに存在します。それは任意の浮動小数点型の代わりで、
// 例えば、float32またはfloat64を表します。
type FloatType float32

// ComplexTypeは、ドキュメンテーションの目的のみでここに存在します。それは任意の複素数型の代わりで、
// 例えば、complex64またはcomplex128を表します。
type ComplexType complex64
