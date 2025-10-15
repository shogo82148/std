// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package entropy

func SHA384(p *[1024]byte) [48]byte

func TestingOnlySHA384(p []byte) [48]byte
