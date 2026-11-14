// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//specgen:category Numeric conversion

package spec

// ConvertFloatToFloat converts each element of x to {{.zE}}.
//
//	z[i] = {{.zE}}(x[i])
//
// {{if lt .zN .xN}}
// The results are rounded to the nearest representable value, or ±inf if they
// cannot be represented.
//
// The results are stored in the low {{lanes .x "x"}} elements of z.
// The remaining elements of z are zero.
// {{else}}
// Only the low {{lanes .z "z"}} elements of x are converted.
// The remaining elements of x are ignored.
// {{end}}
//
//specgen:name ConvertTo{{.zE | title}}
//specgen:require zE!=xE
func ConvertFloatToFloat[xE Floats, W Width, zE Floats](x Vec[xE, W]) (z Vec[zE, W])

// ConvertFloatToInt converts each element of x to {{.zE}}.
//
//	z[i] = {{.zE}}(x[i])
//
// When a conversion is inexact, this truncates the result (rounds toward zero).
// If the converted result would be outside the representable range, the result
// is architecture-dependent.
//
// {{if lt .zN .xN}}
// The results are stored in the low {{lanes .x "x"}} elements of z.
// The remaining elements of z are zero.
// {{else if gt .zN .xN}}
// Only the low {{lanes .z "z"}} elements of x are converted.
// The remaining elements of x are ignored.
// {{end}}
//
//specgen:name ConvertTo{{.zE | title}}
func ConvertFloatToInt[xE Floats, W Width, zE Ints | Uints](x Vec[xE, W]) (z Vec[Variation[zE], W])

// ConvertIntToFloat converts each element of x to {{.zE}}.
//
//	z[i] = {{.zE}}(x[i])
//
// {{if lt .zN .xN}}
// The results are stored in the low {{lanes .x "x"}} elements of z.
// The remaining elements of z are zero.
// {{else if gt .zN .xN}}
// Only the low {{lanes .z "z"}} elements of x are converted.
// The remaining elements of x are ignored.
// {{end}}
//
//specgen:name ConvertTo{{.zE | title}}
func ConvertIntToFloat[xE Ints | Uints, W Width, zE Floats](x Vec[xE, W]) (z Vec[zE, W])

// ConvertIntToUint converts each element of x to {{.zE}}.
//
//	z[i] = {{.zE}}(x[i])
//
// This wraps unrepresentable values.
//
//specgen:name ConvertTo{{.zE | title}}
//specgen:require zB!=xB zN=xN
func ConvertIntToUint[xE Ints | Uints, W Width, zE Ints | Uints](x Vec[xE, W]) (z Vec[zE, W])

// ExtendToZ {{if eq .xB "int"}}sign{{else}}zero{{end}}-extends each element of x to {{.zE}}.
//
//	z[i] = {{.zE}}(x[i])
//
// Only the low {{lanes .z "z"}} elements of x are converted.
// The remaining elements of x are ignored.
//
//specgen:name ExtendTo{{.zE | title}}
//specgen:require zB=xB zN>xN
func ExtendToZ[xE Ints | Uints, W Width, zE Ints | Uints](x Vec[xE, W]) (z Vec[zE, W])

// TruncToZ converts each element of x to {{.zE}} by discarding its high
// bits.
//
//	z[i] = {{.zE}}(x[i])
//
// The results are stored in the low {{lanes .x "x"}} elements of z.
// The remaining elements of z are zero.
//
//specgen:name TruncTo{{.zE | title}}
//specgen:require zB=xB zN<xN
func TruncToZ[xE Ints | Uints, W Width, zE Ints | Uints](x Vec[xE, W]) (z Vec[zE, W])

// SaturateToZ converts each element of x to {{.zE}}, saturating values
// that are outside the range of {{.zE}}.
// {{- if and (eq .xB "int") (eq .zB "uint")}} Negative values become 0.{{end}}
//
//	z[i] = sat(x[i])
//
// The results are stored in the low {{lanes .x "x"}} elements of z.
// The remaining elements of z are zero.
//
//specgen:name SaturateTo{{.zE | title}}
//specgen:require zN<xN
func SaturateToZ[xE Ints | Uints, W Width, zE Ints | Uints](x Vec[xE, W]) (z Vec[zE, W])

// ConcatSaturateToZ converts each element of x and y to {{.zE}}, saturating
// values that are outside the range of {{.zE}}, and returns the concatenated
// result.
// {{- if and (eq .xB "int") (eq .zB "uint")}} Negative values become 0.{{end}}
//
//	z = concat(sat(x[...]), sat(y[...]))
//
//specgen:name ConcatSaturateTo{{.zE | title}}
//specgen:require zN=xN/2
func ConcatSaturateToZ[xE Ints | Uints, W Width, zE Ints | Uints](x, y Vec[xE, W]) (z Vec[zE, W])

// ConcatSaturateToZGrouped divides x, y, and z into groups of 128 bits and
// performs [ConcatSaturateTo{{.zE | title}}] on each group.
//
//specgen:name ConcatSaturateTo{{.zE | title}}Grouped
//specgen:require zN=xN/2
func ConcatSaturateToZGrouped[xE Ints | Uints, W Width, zE Ints | Uints](x, y Vec[xE, W]) (z Vec[zE, W])

// ExtendLoLToZ {{if eq .xB "int"}}sign{{else}}zero{{end}}-extends the low {{.zL}} elements of x to {{.zE}}.
//
//	z[i] = {{.zE}}(x[i])
//
//specgen:name ExtendLo{{.zL}}To{{.zE | title}}
//specgen:require zB=xB zN>xN
//specgen:experiment
func ExtendLoLToZ[E Ints | Uints, zE Ints | Uints](x Vec[E, Width128]) (z Vec[zE, Width128])

// ExtendLoLToZWiden {{if eq .xB "int"}}sign{{else}}zero{{end}}-extends the low {{.zL}} elements of x to {{.zE}}.
//
//	z[i] = {{.zE}}(x[i])
//
//specgen:name ExtendLo{{.zL}}To{{.zE | title}}
//specgen:require zB=xB zL<xL zW>xW
//specgen:experiment
func ExtendLoLToZWiden[E Ints | Uints, zE Ints | Uints, zW FixedWidth](x Vec[E, Width128]) (z Vec[zE, zW])

// ConvertLoLToZ converts the low-indexed {{.zL}} elements of x to {{.zE}}.
//
//	z[i] = {{.zE}}(x[i])
//
//specgen:name ConvertLo{{.zL}}To{{.zE | title}}
//specgen:require zN=xN*2
//specgen:experiment
func ConvertLoLToZ[E Floats, zE Floats](x Vec[E, Width128]) (z Vec[zE, Width128])
