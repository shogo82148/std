// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

// SetFallbackRootsは、カスタムのルートが指定されておらず、プラットフォームの検証者またはシステム証明書プールが利用できない場合（たとえばルート証明書バンドルが存在しないコンテナ内部）に、証明書の検証中に使用するルートを設定します。rootsがnilの場合、SetFallbackRootsはパニックを引き起こします。
// SetFallbackRootsは1回しか呼び出すことができず、複数回呼び出した場合はパニックを引き起こします。
// GODEBUG=x509usefallbackroots=1を設定することで、システム証明書プールが存在していても、すべてのプラットフォームでフォールバックの動作を強制することができます（ただし、WindowsとmacOSでは、これによりプラットフォーム検証APIの使用が無効になり、純粋なGoの検証者が使用されます）。SetFallbackRootsを呼び出さずにx509usefallbackroots=1を設定しても効果はありません。
func SetFallbackRoots(roots *CertPool)
