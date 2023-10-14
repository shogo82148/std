// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// js and wasip1 do not support inter-process file locking.
//
//go:build !js && !wasip1

package lockedfile_test
