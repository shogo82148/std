// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// It's hard to tease out exactly how big a Sigset is, but
// rt_sigprocmask crashes if we get it wrong, so if binaries
// are running, this is right.
