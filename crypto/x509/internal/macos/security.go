// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

package macOS

const (
	// various macOS error codes that can be returned from
	// SecTrustEvaluateWithError that we can map to Go cert
	// verification error types.
	ErrSecCertificateExpired = -67818
	ErrSecHostNameMismatch   = -67602
	ErrSecNotTrusted         = -67843
)

type OSStatus struct {
	call   string
	status int32
}

func (s OSStatus) Error() string

func SecTrustCreateWithCertificates(certs CFRef, policies CFRef) (CFRef, error)

func SecCertificateCreateWithData(b []byte) (CFRef, error)

func SecPolicyCreateSSL(name string) (CFRef, error)

func SecTrustSetVerifyDate(trustObj CFRef, dateRef CFRef) error

func SecTrustEvaluate(trustObj CFRef) (CFRef, error)

func SecTrustEvaluateWithError(trustObj CFRef) (int, error)

func SecCertificateCopyData(cert CFRef) ([]byte, error)

func SecTrustCopyCertificateChain(trustObj CFRef) (CFRef, error)
