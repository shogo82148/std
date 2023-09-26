// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Fake network poller for NaCl and wasm/js.
// Should never be used, because NaCl and wasm/js network connections do not honor "SetNonblock".

//go:build nacl || (js && wasm)
// +build nacl js,wasm

package runtime
