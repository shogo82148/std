// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

// An ErrCode is an unsigned 32-bit error code as defined in the HTTP/2 spec.
type ErrCode uint32

const (
	ErrCodeNo                 ErrCode = 0x0
	ErrCodeProtocol           ErrCode = 0x1
	ErrCodeInternal           ErrCode = 0x2
	ErrCodeFlowControl        ErrCode = 0x3
	ErrCodeSettingsTimeout    ErrCode = 0x4
	ErrCodeStreamClosed       ErrCode = 0x5
	ErrCodeFrameSize          ErrCode = 0x6
	ErrCodeRefusedStream      ErrCode = 0x7
	ErrCodeCancel             ErrCode = 0x8
	ErrCodeCompression        ErrCode = 0x9
	ErrCodeConnect            ErrCode = 0xa
	ErrCodeEnhanceYourCalm    ErrCode = 0xb
	ErrCodeInadequateSecurity ErrCode = 0xc
	ErrCodeHTTP11Required     ErrCode = 0xd
)

func (e ErrCode) String() string

// ConnectionError is an error that results in the termination of the
// entire connection.
type ConnectionError ErrCode

func (e ConnectionError) Error() string

// StreamError is an error that only affects one stream within an
// HTTP/2 connection.
type StreamError struct {
	StreamID uint32
	Code     ErrCode
	Cause    error
}

func (e StreamError) Error() string

// This As function permits converting a StreamError into a x/net/http2.StreamError.
func (e StreamError) As(target any) bool
