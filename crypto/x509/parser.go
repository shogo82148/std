// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

// ParseCertificateは与えられたASN.1 DERデータから単一の証明書を解析します。
func ParseCertificate(der []byte) (*Certificate, error)

// ParseCertificates関数は、与えられたASN.1 DERデータから1つ以上の証明書を解析します。
// 証明書は、間にパディングがない形式で連結されている必要があります。
func ParseCertificates(der []byte) ([]*Certificate, error)

<<<<<<< HEAD
// ParseRevocationListは、与えられたASN.1 DERデータからX509 v2証明書失効リストをパースします。
=======
// ParseRevocationList parses a X509 v2 [Certificate] Revocation List from the given
// ASN.1 DER data.
>>>>>>> upstream/master
func ParseRevocationList(der []byte) (*RevocationList, error)
