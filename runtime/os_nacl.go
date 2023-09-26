// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// gsignalStack is unused on nacl.

// lastfaketime stores the last faketime value written to fd 1 or 2.

// lastfaketimefd stores the fd to which lastfaketime was written.
//
// Subsequent writes to the same fd may use the same timestamp,
// but the timestamp must increase if the fd changes.
