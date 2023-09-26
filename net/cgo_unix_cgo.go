// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cgo && !netgo && unix && !darwin

package net

/*
#cgo CFLAGS: -fno-stack-protector
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <netdb.h>
#include <unistd.h>
#include <string.h>
#include <stdlib.h>

// If nothing else defined EAI_OVERFLOW, make sure it has a value.
#ifndef EAI_OVERFLOW
#define EAI_OVERFLOW -12
#endif
*/
