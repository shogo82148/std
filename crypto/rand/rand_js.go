// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm

package rand

// The maximum buffer size for crypto.getRandomValues is 65536 bytes.
// https://developer.mozilla.org/en-US/docs/Web/API/Crypto/getRandomValues#exceptions

// reader implements a pseudorandom generator
// using JavaScript crypto.getRandomValues method.
// See https://developer.mozilla.org/en-US/docs/Web/API/Crypto/getRandomValues.
