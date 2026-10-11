// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1

package syscall

// ParseDirentは、buf内の最大max個のディレクトリエントリをパースし、
// 名前をnamesに追加します。bufから消費されたバイト数、
// namesに追加されたエントリの数、および新しいnamesスライスを返します。
func ParseDirent(buf []byte, max int, names []string) (consumed int, count int, newnames []string)
