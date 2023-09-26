// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

// locker2 represents an API surface of two concurrent goroutines
// locking the same resource, but through different APIs. It's intended
// to abstract over the relationship of two Lock calls or an RLock
// and a Lock call.
