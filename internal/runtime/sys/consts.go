// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sys

import (
	"github.com/shogo82148/std/internal/goarch"
	"github.com/shogo82148/std/internal/goos"
)

// AIX and OpenBSD require a larger stack for syscalls.
// The race build also needs more stack. See issue 54291.
// This arithmetic must match that in cmd/internal/objabi/stack.go:stackGuardMultiplier.
const StackGuardMultiplier = 1 + goos.IsAix + goos.IsOpenbsd + isRace

// DefaultPhysPageSize is the default physical page size.
const DefaultPhysPageSize = goarch.DefaultPhysPageSize

// PCQuantum is the minimal unit for a program counter (1 on x86, 4 on most other systems).
// The various PC tables record PC deltas pre-divided by PCQuantum.
const PCQuantum = goarch.PCQuantum

// Int64Align is the required alignment for a 64-bit integer.
const Int64Align = goarch.Int64Align

// MinFrameSize is the size of the system-reserved words at the bottom
// of a frame (just above the architectural stack pointer).
// It is zero on x86 and PtrSize on most non-x86 (LR-based) systems.
// On PowerPC it is larger, to cover three more reserved words:
// the compiler word, the link editor word, and the TOC save word.
const MinFrameSize = goarch.MinFrameSize

// StackAlign is the required alignment of the SP register.
// The stack must be at least word aligned, but some architectures require more.
const StackAlign = goarch.StackAlign

// HeapAddrBits はヒープアドレスのビット数です。amd64 では、
// アドレスは HeapAddrBits を超える部分が符号拡張されます。
// その他のアーキテクチャーでは、ゼロ拡張されます。
//
// ほとんどの 64 ビットプラットフォームでは、ハードウェアと OS の制限の
// 組み合わせに基づいて、これを 48 ビットに制限しています。
//
// amd64 ハードウェアはアドレスを 48 ビットに制限しており、64 ビットに符号拡張されます。
// 上位 16 ビットがすべて 0 でもすべて 1 でもないアドレスは
// 「非正規 (non-canonical)」であり、無効です。
// これらの「負の」アドレスがあるため、amd64 ではヒープアリーナのインデックスを
// 計算する前に、アドレスを 1<<47 (arenaBaseOffset) だけオフセットします。
// 2017 年に amd64 ハードウェアは 57 ビットアドレスのサポートを追加しましたが、
// 現在この拡張をサポートしているのは Linux だけであり、
// カーネルは、1<<47 を超えるヒントアドレスを指定して mmap が呼び出されない限り
// （Go はそのようなことをしません）、1<<47 を超えるアドレスを選択することはありません。
//
// arm64 ハードウェア (ARMv8 時点) は、ユーザーアドレスを 48 ビット、
// つまり [0, 1<<48) の範囲に制限しています。
//
// ppc64、mips64、s390x は、ハードウェアとして任意の 64 ビットアドレスをサポートしています。
// Linux では、Go はより厳しい OS の制限に依存しています。
// Linux の processor.h に基づくと、64 ビットアーキテクチャーでの
// ユーザーアドレス空間は次のように制限されています。
//
// Architecture  Name              Maximum Value (exclusive)
// ---------------------------------------------------------------------
// amd64         TASK_SIZE_MAX     0x007ffffffff000 (47 bit addresses)
// arm64         TASK_SIZE_64      0x01000000000000 (48 bit addresses)
// ppc64{,le}    TASK_SIZE_USER64  0x00400000000000 (46 bit addresses)
// mips64{,le}   TASK_SIZE64       0x00010000000000 (40 bit addresses)
// s390x         TASK_SIZE         1<<64 (64 bit addresses)
//
// これらの制限は時間とともに増える可能性がありますが、現在は s390x を除いて
// 最大 48 ビットです。すべてのアーキテクチャーで、Linux は mmap された領域を
// 48 ビットよりかなり低いアドレスに配置し始めるため、
// Go の 48 ビット制限を超えることが可能だとしても、実際にはまず起こりません。
//
// 32 ビットプラットフォームでは、コストがかからないため、
// 32 ビットのアドレス空間全体を受け入れます。
// mips32 は仮想メモリの下位 2GB にしかアクセスできないため、
// さらに 31 ビットに制限します。
//
// ios/arm64 では、64 ビットポインターは利用可能と思われますが、
// iOS 14 未満ではポインターが 33 ビットに切り詰められます。
// さらに、アプリケーションが実際に利用できるのはアドレス空間の上位 4 GiB だけです。
// iOS 14 以降では、より多くのアドレス空間が利用可能で、
// OS はこの 33 ビットの外側のアドレスを提供できます。
// ページアロケーターによるアドレス空間の使用量と、ヒープとして受け入れる
// mmap 領域の柔軟性との合理的なバランスとして、40 ビットを選んでいます。
// 古い iOS バージョンではアドレス空間を使いすぎるため、
// 単純に 48 ビット全体へ移行することはできません。
// TODO(mknyszek): iOS 14 未満が非推奨になったら、ios/arm64 を
// 他のすべての arm64 プラットフォームと同様に 48 ビットのアドレス空間に引き上げる。
//
// WebAssembly には現在、4GB のリニアメモリという制限があります。
//
// OpenBSD/arm64 は、次によると 39 ビットに制限されています。
// https://github.com/openbsd/src/blob/65a9769f762528944cb8c228592381427b188551/sys/arch/arm64/include/vmparam.h#L93
const HeapAddrBits = (is64bit*(1-goarch.IsWasm)*(1-goos.IsIos*goarch.IsArm64)*(1-goos.IsOpenbsd*goarch.IsArm64))*48 +
	(1-is64bit+goarch.IsWasm)*(32-(goarch.IsMips+goarch.IsMipsle)) +
	40*goos.IsIos*goarch.IsArm64 +
	39*goos.IsOpenbsd*goarch.IsArm64

// MaxAlloc は割り当ての最大サイズです。64 ビットでは、
// 理論上 1<<HeapAddrBits バイトを割り当てることが可能です。
// 一方 32 ビットでは、アドレス空間のバイト数が uintptr に収まらないため、
// 1<<32 より 1 小さい値になります。
const MaxAlloc = (1 << HeapAddrBits) - (1-is64bit)*1
