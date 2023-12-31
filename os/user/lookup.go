// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package user

// Currentは現在のユーザーを返します。
//
// 最初の呼び出しは現在のユーザー情報をキャッシュします。
// その後の呼び出しはキャッシュされた値を返し、現在のユーザーへの変更は反映されません。
func Current() (*User, error)

// Lookupはユーザー名でユーザーを検索します。ユーザーが見つからない場合、
// 返されるエラーのタイプはUnknownUserErrorです。
func Lookup(username string) (*User, error)

// LookupIdはユーザーIDでユーザーを検索します。ユーザーが見つからない場合、
// 返されるエラーのタイプはUnknownUserIdErrorです。
func LookupId(uid string) (*User, error)

// LookupGroupは名前でグループを検索します。グループが見つからない場合、
// 返されるエラーのタイプはUnknownGroupErrorです。
func LookupGroup(name string) (*Group, error)

// LookupGroupIdはグループIDでグループを検索します。グループが見つからない場合、
// 返されるエラーのタイプはUnknownGroupIdErrorです。
func LookupGroupId(gid string) (*Group, error)

// GroupIdsは、ユーザーがメンバーであるグループIDのリストを返します。
func (u *User) GroupIds() ([]string, error)
