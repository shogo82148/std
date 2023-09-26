// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cgo && !osusergo
// +build cgo,!osusergo

package user

/*
#include <unistd.h>
#include <sys/types.h>
#include <stdlib.h>

static int mygetgrouplist(const char* user, gid_t group, gid_t* groups, int* ngroups) {
	int* buf = malloc(*ngroups * sizeof(int));
	int rv = getgrouplist(user, (int) group, buf, ngroups);
	int i;
	if (rv == 0) {
		for (i = 0; i < *ngroups; i++) {
			groups[i] = (gid_t) buf[i];
		}
	}
	free(buf);
	return rv;
}
*/
