// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// facilities is padded to avoid false sharing.

// cpu indicates the availability of s390x facilities that can be used in
// Go assembly but are optional on models supported by Go.
