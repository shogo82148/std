// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// NOTE(rsc): _CONTEXT_CONTROL is actually 0x200001 and should include PC, SP, and LR.
// However, empirically, LR doesn't come along on Windows 10
// unless you also set _CONTEXT_INTEGER (0x200002).
// Without LR, we skip over the next-to-bottom function in profiles
// when the bottom function is frameless.
// So we set both here, to make a working _CONTEXT_CONTROL.
