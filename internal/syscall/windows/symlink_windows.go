// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

import "github.com/shogo82148/std/syscall"

const (
	ERROR_INVALID_PARAMETER syscall.Errno = 87

	FILE_SUPPORTS_OBJECT_IDS      = 0x00010000
	FILE_SUPPORTS_OPEN_BY_FILE_ID = 0x01000000

	// symlink support for CreateSymbolicLink() starting with Windows 10 (1703, v10.0.14972)
	SYMBOLIC_LINK_FLAG_ALLOW_UNPRIVILEGED_CREATE = 0x2

	// FileInformationClass values
	FileBasicInfo                  = 0
	FileStandardInfo               = 1
	FileNameInfo                   = 2
	FileDispositionInfo            = 4
	FileStreamInfo                 = 7
	FileCompressionInfo            = 8
	FileAttributeTagInfo           = 9
	FileIdBothDirectoryInfo        = 0xa
	FileIdBothDirectoryRestartInfo = 0xb
	FileRemoteProtocolInfo         = 0xd
	FileFullDirectoryInfo          = 0xe
	FileFullDirectoryRestartInfo   = 0xf
	FileStorageInfo                = 0x10
	FileAlignmentInfo              = 0x11
	FileIdInfo                     = 0x12
	FileIdExtdDirectoryInfo        = 0x13
	FileIdExtdDirectoryRestartInfo = 0x14
)

type FILE_ATTRIBUTE_TAG_INFO struct {
	FileAttributes uint32
	ReparseTag     uint32
}
