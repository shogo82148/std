// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Fake network poller for js/wasm.
// Should never be used, because js/wasm network connections do not honor "SetNonblock".

//go:build js && wasm

package runtime
