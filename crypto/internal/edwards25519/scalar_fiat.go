// Code generated by Fiat Cryptography. DO NOT EDIT.
//
// Autogenerated: word_by_word_montgomery --lang Go --cmovznz-by-mul --relax-primitive-carry-to-bitwidth 32,64 --public-function-case camelCase --public-type-case camelCase --private-function-case camelCase --private-type-case camelCase --doc-text-before-function-name '' --doc-newline-before-package-declaration --doc-prepend-header 'Code generated by Fiat Cryptography. DO NOT EDIT.' --package-name edwards25519 Scalar 64 '2^252 + 27742317777372353535851937790883648493' mul add sub opp nonzero from_montgomery to_montgomery to_bytes from_bytes
//
// curve description: Scalar
//
// machine_wordsize = 64 (from "64")
//
// requested operations: mul, add, sub, opp, nonzero, from_montgomery, to_montgomery, to_bytes, from_bytes
//
// m = 0x1000000000000000000000000000000014def9dea2f79cd65812631a5cf5d3ed (from "2^252 + 27742317777372353535851937790883648493")
//
//
//
// NOTE: In addition to the bounds specified above each function, all
//
//   functions synthesized for this Montgomery arithmetic require the
//
//   input to be strictly less than the prime modulus (m), and also
//
//   require the input to be in the unique saturated representation.
//
//   All functions also ensure that these two properties are true of
//
//   return values.
//
//
//
// Computed values:
//
//   eval z = z[0] + (z[1] << 64) + (z[2] << 128) + (z[3] << 192)
//
//   bytes_eval z = z[0] + (z[1] << 8) + (z[2] << 16) + (z[3] << 24) + (z[4] << 32) + (z[5] << 40) + (z[6] << 48) + (z[7] << 56) + (z[8] << 64) + (z[9] << 72) + (z[10] << 80) + (z[11] << 88) + (z[12] << 96) + (z[13] << 104) + (z[14] << 112) + (z[15] << 120) + (z[16] << 128) + (z[17] << 136) + (z[18] << 144) + (z[19] << 152) + (z[20] << 160) + (z[21] << 168) + (z[22] << 176) + (z[23] << 184) + (z[24] << 192) + (z[25] << 200) + (z[26] << 208) + (z[27] << 216) + (z[28] << 224) + (z[29] << 232) + (z[30] << 240) + (z[31] << 248)
//
//   twos_complement_eval z = let x1 := z[0] + (z[1] << 64) + (z[2] << 128) + (z[3] << 192) in
//
//                            if x1 & (2^256-1) < 2^255 then x1 & (2^256-1) else (x1 & (2^256-1)) - 2^256

package edwards25519
