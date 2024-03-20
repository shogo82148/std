// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
パッケージuserは、名前またはIDによるユーザーアカウントの検索を可能にします。

ほとんどのUnixシステムでは、このパッケージにはユーザーとグループのIDを名前に解決し、
補足的なグループIDをリストアップするための2つの内部実装があります。
一つは純粋なGoで書かれており、/etc/passwdと/etc/groupを解析します。
もう一つはcgoベースで、getpwuid_r、getgrnam_r、getgrouplistなどの
標準Cライブラリ(libc)のルーチンに依存しています。

cgoが利用可能で、特定のプラットフォームのlibcに必要なルーチンが実装されている場合、
cgoベース（libcバックエンド）のコードが使用されます。
これは、純粋なGoの実装を強制するosusergoビルドタグを使用することで上書きすることができます。
*/
package user

// Userはユーザーアカウントを表します。
type User struct {
	// UidはユーザーIDです。
	// POSIXシステムでは、これはuidを表す10進数です。
	// Windowsでは、これは文字列形式のセキュリティ識別子（SID）です。
	// Plan 9では、これは/dev/userの内容です。
	Uid string
	// GidはプライマリグループIDです。
	// POSIXシステムでは、これはgidを表す10進数です。
	// Windowsでは、これは文字列形式のセキュリティ識別子（SID）です。
	// Plan 9では、これは/dev/userの内容です。
	Gid string
	// Usernameはログイン名です。
	Username string
	// Nameはユーザーの実名または表示名です。
	// 空である可能性があります。
	// POSIXシステムでは、これはGECOSフィールドリストの最初（または唯一）のエントリです。
	// Windowsでは、これはユーザーの表示名です。
	// Plan 9では、これは/dev/userの内容です。
	Name string
	// HomeDirはユーザーのホームディレクトリへのパスです（もし存在する場合）。
	HomeDir string
}

// Groupはユーザーのグループを表します。
//
// POSIXシステムでは、GidはグループIDを表す10進数を含みます。
type Group struct {
	Gid  string
	Name string
}

// UnknownUserIdErrorは、ユーザーが見つからない場合に [LookupId] によって返されるエラーです。
type UnknownUserIdError int

func (e UnknownUserIdError) Error() string

// UnknownUserErrorは、ユーザーが見つからない場合に [Lookup] によって返されるエラーです。
type UnknownUserError string

func (e UnknownUserError) Error() string

// UnknownGroupIdErrorは、グループが見つからない場合に [LookupGroupId] によって返されるエラーです。
type UnknownGroupIdError string

func (e UnknownGroupIdError) Error() string

// UnknownGroupErrorは、グループが見つからない場合に [LookupGroup] によって返されるエラーです。
type UnknownGroupError string

func (e UnknownGroupError) Error() string
