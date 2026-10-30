// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package user

// Currentは現在のユーザーを返します。
//
// 最初の呼び出しは現在のユーザー情報をキャッシュします。
// その後の呼び出しはキャッシュされた値を返し、現在のユーザーへの変更は反映されません。
//
// 一部のシステムでは、現在のユーザーがユーザーデータベースに見つからない場合、
// Current はプロセスのユーザー ID とグループ ID、および USER や HOME などの
// 環境変数から、返すユーザーを構築します。
// その場合、返されるユーザーはユーザーデータベースのエントリーには対応しません。
func Current() (*User, error)

// Lookupはユーザー名でユーザーを検索します。ユーザーが見つからない場合、
// 返されるエラーのタイプは [UnknownUserError] です。
//
// Lookup はまず [Current] を呼び出します。それが成功し、現在のユーザーの
// 名前が username と一致する場合、Lookup は名前による個別の検索を行わずに
// 現在のユーザーを返します。
// その結果は、ユーザーデータベースではなく環境から [Current] が構築したものである可能性があります。
func Lookup(username string) (*User, error)

// LookupIdはユーザーIDでユーザーを検索します。ユーザーが見つからない場合、
// 返されるエラーのタイプは [UnknownUserIdError] です。
//
// LookupId はまず [Current] を呼び出します。それが成功し、現在のユーザーの
// ID が uid と一致する場合、LookupId は ID による個別の検索を行わずに
// 現在のユーザーを返します。
// その結果は、ユーザーデータベースではなく環境から [Current] が構築したものである可能性があります。
func LookupId(uid string) (*User, error)

// LookupGroupは名前でグループを検索します。グループが見つからない場合、
// 返されるエラーのタイプは [UnknownGroupError] です。
func LookupGroup(name string) (*Group, error)

// LookupGroupIdはグループIDでグループを検索します。グループが見つからない場合、
// 返されるエラーのタイプは [UnknownGroupIdError] です。
func LookupGroupId(gid string) (*Group, error)

// GroupIdsは、ユーザーがメンバーであるグループIDのリストを返します。
func (u *User) GroupIds() ([]string, error)
