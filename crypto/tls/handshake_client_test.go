// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// Script of interaction with gnutls implementation.
// The values for this test are obtained by building and running in client mode:
//   % go test -run "TestRunClient" -connect
// and then:
//   % gnutls-serv -p 10443 --debug 100 --x509keyfile key.pem --x509certfile cert.pem -a > /tmp/log 2>&1
//   % python parse-gnutls-cli-debug-log.py < /tmp/log
//
// Where key.pem is:
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
// and cert.pem is:
// -----BEGIN CERTIFICATE-----
// MIIBoDCCAUoCAQAwDQYJKoZIhvcNAQEEBQAwYzELMAkGA1UEBhMCQVUxEzARBgNV
// BAgTClF1ZWVuc2xhbmQxGjAYBgNVBAoTEUNyeXB0U29mdCBQdHkgTHRkMSMwIQYD
// VQQDExpTZXJ2ZXIgdGVzdCBjZXJ0ICg1MTIgYml0KTAeFw05NzA5MDkwMzQxMjZa
// Fw05NzEwMDkwMzQxMjZaMF4xCzAJBgNVBAYTAkFVMRMwEQYDVQQIEwpTb21lLVN0
// YXRlMSEwHwYDVQQKExhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQxFzAVBgNVBAMT
// DkVyaWMgdGhlIFlvdW5nMFEwCQYFKw4DAgwFAANEAAJBALVEqPODnpI4rShlY8S7
// tB713JNvabvn6Gned7zylwLLiXQAo/PAT6mfdWPTyCX9RlId/Aroh1ou893BA32Q
// sggwDQYJKoZIhvcNAQEEBQADQQCU5SSgapJSdRXJoX+CpCvFy+JVh9HpSjCpSNKO
// 19raHv98hKAUJuP9HyM+SUsffO6mAIgitUaqW8/wDMePhEC3
// -----END CERTIFICATE-----
