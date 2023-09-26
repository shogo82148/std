// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// AT_HWCAP is not available on FreeBSD-11.1-RELEASE or earlier.
// Default to mandatory VFP hardware support for arm being available.
// If AT_HWCAP is available goarmHWCap will be updated in archauxv.
// TODO(moehrmann) remove once all go supported FreeBSD versions support _AT_HWCAP.
