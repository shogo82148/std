// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

package macOS

import (
	"github.com/shogo82148/std/errors"
)

type SecTrustSettingsResult int32

const (
	SecTrustSettingsResultInvalid SecTrustSettingsResult = iota
	SecTrustSettingsResultTrustRoot
	SecTrustSettingsResultTrustAsRoot
	SecTrustSettingsResultDeny
	SecTrustSettingsResultUnspecified
)

type SecTrustResultType int32

const (
	SecTrustResultInvalid SecTrustResultType = iota
	SecTrustResultProceed
	SecTrustResultConfirm
	SecTrustResultDeny
	SecTrustResultUnspecified
	SecTrustResultRecoverableTrustFailure
	SecTrustResultFatalTrustFailure
	SecTrustResultOtherError
)

type SecTrustSettingsDomain int32

const (
	SecTrustSettingsDomainUser SecTrustSettingsDomain = iota
	SecTrustSettingsDomainAdmin
	SecTrustSettingsDomainSystem
)

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

var SecTrustSettingsResultKey = StringToCFString("kSecTrustSettingsResult")
var SecTrustSettingsPolicy = StringToCFString("kSecTrustSettingsPolicy")
var SecTrustSettingsPolicyString = StringToCFString("kSecTrustSettingsPolicyString")
var SecPolicyOid = StringToCFString("SecPolicyOid")
var SecPolicyAppleSSL = StringToCFString("1.2.840.113635.100.1.3")

var ErrNoTrustSettings = errors.New("no trust settings found")

func SecTrustSettingsCopyCertificates(domain SecTrustSettingsDomain) (certArray CFRef, err error)

func SecTrustSettingsCopyTrustSettings(cert CFRef, domain SecTrustSettingsDomain) (trustSettings CFRef, err error)

func SecTrustCreateWithCertificates(certs CFRef, policies CFRef) (CFRef, error)

func SecCertificateCreateWithData(b []byte) (CFRef, error)

func SecPolicyCreateSSL(name string) (CFRef, error)

func SecTrustSetVerifyDate(trustObj CFRef, dateRef CFRef) error

func SecTrustEvaluate(trustObj CFRef) (CFRef, error)

func SecTrustGetResult(trustObj CFRef, result CFRef) (CFRef, CFRef, error)

func SecTrustEvaluateWithError(trustObj CFRef) (int, error)

func SecTrustGetCertificateCount(trustObj CFRef) int

func SecTrustGetCertificateAtIndex(trustObj CFRef, i int) (CFRef, error)

func SecCertificateCopyData(cert CFRef) ([]byte, error)
