// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// ChaCha8 is ChaCha with 8 rounds.
// See https://cr.yp.to/chacha/chacha-20080128.pdf.
//
// ChaCha8 operates on a 4x4 matrix of uint32 values, initially set to:
//
//	const1 const2 const3 const4
//	seed   seed   seed   seed
//	seed   seed   seed   seed
//	counter64     0      0
//
// We use the same constants as ChaCha20 does, a random seed,
// and a counter. Running ChaCha8 on this input produces
// a 4x4 matrix of pseudo-random values with as much entropy
// as the seed.
//
// Given SIMD registers that can hold N uint32s, it is possible
// to run N ChaCha8 block transformations in parallel by filling
// the first register with the N copies of const1, the second
// with N copies of const2, and so on, and then running the operations.
//
// Each iteration of ChaCha8Rand operates over 32 bytes of input and
// produces 992 bytes of RNG output, plus 32 bytes of input for the next
// iteration.
//
// The 32 bytes of input are used as a ChaCha8 key, with a zero nonce, to
// produce 1024 bytes of output (16 blocks, with counters 0 to 15).
// First, for each block, the values 0x61707865, 0x3320646e, 0x79622d32,
// 0x6b206574 are subtracted from the 32-bit little-endian words at
// position 0, 1, 2, and 3 respectively, and an increasing counter
// starting at zero is subtracted from each word at position 12. Then,
// this stream is permuted such that for each sequence of four blocks,
// first we output the first four bytes of each block, then the next four
// bytes of each block, and so on. Finally, the last 32 bytes of output
// are used as the input of the next iteration, and the remaining 992
// bytes are the RNG output.
//
// See https://github.com/C2SP/C2SP/blob/main/chacha8rand/chacha8rand.md
// for additional details.
//
// Normal ChaCha20 implementations for encryption use this same
// parallelism but then have to deinterlace the results so that
// it appears the blocks were generated separately. For the purposes
// of generating random numbers, the interlacing is fine.
// We are simply locked in to preserving the 4-way interlacing
// in any future optimizations.
package chacha8rand
