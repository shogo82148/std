// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

<<<<<<< HEAD
// ParseCertificateは与えられたASN.1 DERデータから単一の証明書を解析します。
=======
// ParseCertificate parses a single certificate from the given ASN.1 DER data.
//
// Before Go 1.23, ParseCertificate accepted certificates with negative serial
// numbers. This behavior can be restored by including "x509negativeserial=1" in
// the GODEBUG environment variable.
//
// Before Go 1.23, ParseCertificate accepted certificates with serial numbers
// longer than 20 octets. This behavior can be restored by including
// "x509seriallength=1" in the GODEBUG environment variable.
>>>>>>> d32e3230aa4d4baa9384e050abcdef2da31fe8ae
func ParseCertificate(der []byte) (*Certificate, error)

// ParseCertificates関数は、与えられたASN.1 DERデータから1つ以上の証明書を解析します。
// 証明書は、間にパディングがない形式で連結されている必要があります。
func ParseCertificates(der []byte) ([]*Certificate, error)

// ParseRevocationListは、与えられたASN.1 DERデータからX509 v2 [Certificate] 失効リストをパースします。
func ParseRevocationList(der []byte) (*RevocationList, error)
