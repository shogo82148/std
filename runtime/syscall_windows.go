// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// cbs stores all registered Go callbacks.

// winCallback records information about a registered Go callback.

// abiPartKind is the action an abiPart should take.

// abiPart encodes a step in translating between calling ABIs.

// abiDesc specifies how to translate from a C frame to a Go
// frame. This does not specify how to translate back because
// the result is always a uintptr. If the C ABI is fastcall,
// this assumes the four fastcall registers were first spilled
// to the shadow space.
