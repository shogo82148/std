// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !purego

package aes

// BlockFunction returns the function code for the block cipher.
// It is used by the GCM implementation to invoke the KMA instruction.
func BlockFunction(c *Block) int

// BlockKey returns the key for the block cipher.
// It is used by the GCM implementation to invoke the KMA instruction.
func BlockKey(c *Block) []byte
