// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package upki

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/os"
)

// On-disk layout of index.bin from
// https://c2sp.org/upki-revocation@v1.0.0.
//
//	HEADER (14 bytes):
//	  magic:         [8]byte    "upkiidx1"
//	  num_filenames: u16 BE
//	  num_log_ids:   u32 BE
//
//	TABLES (read eagerly):
//	  filename[num_filenames]:   [32]byte  restricted ASCII, NULL-padded
//	  log_dir[num_log_ids]:                sorted lexicographically by log_id
//	    log_id:      [32]byte
//	    offset:      u64 BE                file offset of entry section
//	    num_entries: u16 BE
//
//	ENTRY SECTIONS (read on demand via ReadAt):
//	  entry[num_entries]:
//	    filter_index:  u16 BE
//	    min_timestamp: u64 BE
//	    max_timestamp: u64 BE
//
// The legacy "upkiidx0" format encoded num_filenames and filter_index as u8;
// it is not supported.
const (
	IndexFilename = "index.bin"
)

var (
	// ErrNoData reports that the cache index is absent or inaccessible.
	ErrNoData = errors.New("upki: no revocation data")
)

// OpenIndex opens the UPKI index at path. If the path is missing, a parent
// is not a directory, or access is denied, the error wraps ErrNoData.
func OpenIndex(path string) (*os.File, error)
