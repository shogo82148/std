// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// Script of interaction with gnutls implementation.
// The values for this test are obtained by building and running in client mode:
//   % go test -test.run "TestRunClient" -connect
// The recorded bytes are written to stdout.
//
// The server private key is:
// -----BEGIN RSA PRIVATE KEY-----
// MIIBPAIBAAJBAJ+zw4Qnlf8SMVIPFe9GEcStgOY2Ww/dgNdhjeD8ckUJNP5VZkVD
// TGiXav6ooKXfX3j/7tdkuD8Ey2//Kv7+ue0CAwEAAQJAN6W31vDEP2DjdqhzCDDu
// OA4NACqoiFqyblo7yc2tM4h4xMbC3Yx5UKMN9ZkCtX0gzrz6DyF47bdKcWBzNWCj
// gQIhANEoojVt7hq+SQ6MCN6FTAysGgQf56Q3TYoJMoWvdiXVAiEAw3e3rc+VJpOz
// rHuDo6bgpjUAAXM+v3fcpsfZSNO6V7kCIQCtbVjanpUwvZkMI9by02oUk9taki3b
// PzPfAfNPYAbCJQIhAJXNQDWyqwn/lGmR11cqY2y9nZ1+5w3yHGatLrcDnQHxAiEA
// vnlEGo8K85u+KwIOimM48ZG8oTk7iFdkqLJR1utT3aU=
// -----END RSA PRIVATE KEY-----
//
// and certificate is:
// -----BEGIN CERTIFICATE-----
// MIICKzCCAdWgAwIBAgIJALE1E2URIMWSMA0GCSqGSIb3DQEBBQUAMEUxCzAJBgNV
// BAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBX
// aWRnaXRzIFB0eSBMdGQwHhcNMTIwNDA2MTcxMDEzWhcNMTUwNDA2MTcxMDEzWjBF
// MQswCQYDVQQGEwJBVTETMBEGA1UECBMKU29tZS1TdGF0ZTEhMB8GA1UEChMYSW50
// ZXJuZXQgV2lkZ2l0cyBQdHkgTHRkMFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAJ+z
// w4Qnlf8SMVIPFe9GEcStgOY2Ww/dgNdhjeD8ckUJNP5VZkVDTGiXav6ooKXfX3j/
// 7tdkuD8Ey2//Kv7+ue0CAwEAAaOBpzCBpDAdBgNVHQ4EFgQUeKaXmmO1xaGlM7oi
// fCNuWxt6zCswdQYDVR0jBG4wbIAUeKaXmmO1xaGlM7oifCNuWxt6zCuhSaRHMEUx
// CzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0YXRlMSEwHwYDVQQKExhJbnRl
// cm5ldCBXaWRnaXRzIFB0eSBMdGSCCQCxNRNlESDFkjAMBgNVHRMEBTADAQH/MA0G
// CSqGSIb3DQEBBQUAA0EAhTZAc8G7GtrUWZ8tonAxRnTsg26oyDxRrzms7EC86CJG
// HZnWRiok1IsFCEv7NRFukrt3uuQSu/TIXpyBqJdgTA==
// -----END CERTIFICATE-----
