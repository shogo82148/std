// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package net

// nssConf represents the state of the machine's /etc/nsswitch.conf file.

// nssCriterion is the parsed structure of one of the criteria in brackets
// after an NSS source name.
