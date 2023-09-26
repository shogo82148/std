// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Suffix array construction by induced sorting (SAIS).
// See Ge Nong, Sen Zhang, and Wai Hong Chen,
// "Two Efficient Algorithms for Linear Time Suffix Array Construction",
// especially section 3 (https://ieeexplore.ieee.org/document/5582081).
// See also http://zork.net/~st/jottings/sais.html.
//
// With optimizations inspired by Yuta Mori's sais-lite
// (https://sites.google.com/site/yuta256/sais).
//
// And with other new optimizations.

// Many of these functions are parameterized by the sizes of
// the types they operate on. The generator gen.go makes
// copies of these functions for use with other sizes.
// Specifically:
//
// - A function with a name ending in _8_32 takes []byte and []int32 arguments
//   and is duplicated into _32_32, _8_64, and _64_64 forms.
//   The _32_32 and _64_64_ suffixes are shortened to plain _32 and _64.
//   Any lines in the function body that contain the text "byte-only" or "256"
//   are stripped when creating _32_32 and _64_64 forms.
//   (Those lines are typically 8-bit-specific optimizations.)
//
// - A function with a name ending only in _32 operates on []int32
//   and is duplicated into a _64 form. (Note that it may still take a []byte,
//   but there is no need for a version of the function in which the []byte
//   is widened to a full integer array.)

// The overall runtime of this code is linear in the input size:
// it runs a sequence of linear passes to reduce the problem to
// a subproblem at most half as big, invokes itself recursively,
// and then runs a sequence of linear passes to turn the answer
// for the subproblem into the answer for the original problem.
// This gives T(N) = O(N) + T(N/2) = O(N) + O(N/2) + O(N/4) + ... = O(N).
//
// The outline of the code, with the forward and backward scans
// through O(N)-sized arrays called out, is:
//
// sais_I_N
//	placeLMS_I_B
//		bucketMax_I_B
//			freq_I_B
//				<scan +text> (1)
//			<scan +freq> (2)
//		<scan -text, random bucket> (3)
//	induceSubL_I_B
//		bucketMin_I_B
//			freq_I_B
//				<scan +text, often optimized away> (4)
//			<scan +freq> (5)
//		<scan +sa, random text, random bucket> (6)
//	induceSubS_I_B
//		bucketMax_I_B
//			freq_I_B
//				<scan +text, often optimized away> (7)
//			<scan +freq> (8)
//		<scan -sa, random text, random bucket> (9)
//	assignID_I_B
//		<scan +sa, random text substrings> (10)
//	map_B
//		<scan -sa> (11)
//	recurse_B
//		(recursive call to sais_B_B for a subproblem of size at most 1/2 input, often much smaller)
//	unmap_I_B
//		<scan -text> (12)
//		<scan +sa> (13)
//	expand_I_B
//		bucketMax_I_B
//			freq_I_B
//				<scan +text, often optimized away> (14)
//			<scan +freq> (15)
//		<scan -sa, random text, random bucket> (16)
//	induceL_I_B
//		bucketMin_I_B
//			freq_I_B
//				<scan +text, often optimized away> (17)
//			<scan +freq> (18)
//		<scan +sa, random text, random bucket> (19)
//	induceS_I_B
//		bucketMax_I_B
//			freq_I_B
//				<scan +text, often optimized away> (20)
//			<scan +freq> (21)
//		<scan -sa, random text, random bucket> (22)
//
// Here, _B indicates the suffix array size (_32 or _64) and _I the input size (_8 or _B).
//
// The outline shows there are in general 22 scans through
// O(N)-sized arrays for a given level of the recursion.
// In the top level, operating on 8-bit input text,
// the six freq scans are fixed size (256) instead of potentially
// input-sized. Also, the frequency is counted once and cached
// whenever there is room to do so (there is nearly always room in general,
// and always room at the top level), which eliminates all but
// the first freq_I_B text scans (that is, 5 of the 6).
// So the top level of the recursion only does 22 - 6 - 5 = 11
// input-sized scans and a typical level does 16 scans.
//
// The linear scans do not cost anywhere near as much as
// the random accesses to the text made during a few of
// the scans (specifically #6, #9, #16, #19, #22 marked above).
// In real texts, there is not much but some locality to
// the accesses, due to the repetitive structure of the text
// (the same reason Burrows-Wheeler compression is so effective).
// For random inputs, there is no locality, which makes those
// accesses even more expensive, especially once the text
// no longer fits in cache.
// For example, running on 50 MB of Go source code, induceSubL_8_32
// (which runs only once, at the top level of the recursion)
// takes 0.44s, while on 50 MB of random input, it takes 2.55s.
// Nearly all the relative slowdown is explained by the text access:
//
//		c0, c1 := text[k-1], text[k]
//
// That line runs for 0.23s on the Go text and 2.02s on random text.

//go:generate go run gen.go

package suffixarray
