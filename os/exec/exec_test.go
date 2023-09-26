// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Use an external test to avoid os/exec -> net/http -> crypto/x509 -> os/exec
// circular dependency on non-cgo darwin.

package exec_test

// haveUnexpectedFDs is set at init time to report whether any file descriptors
// were open at program start.

// A tickReader reads an unbounded sequence of timestamps at no more than a
// fixed interval.
