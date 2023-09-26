// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build armbe || arm64be || m68k || mips || mips64 || mips64p32 || ppc || ppc64 || s390 || s390x || shbe || sparc || sparc64

package binary

// NativeEndian is the native-endian implementation of [ByteOrder] and [AppendByteOrder].
var NativeEndian nativeEndian
