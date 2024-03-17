// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package user

// Currentは現在のユーザーを返します。
//
// 最初の呼び出しは現在のユーザー情報をキャッシュします。
// その後の呼び出しはキャッシュされた値を返し、現在のユーザーへの変更は反映されません。
func Current() (*User, error)

<<<<<<< HEAD
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
=======
// Lookup looks up a user by username. If the user cannot be found, the
// returned error is of type [UnknownUserError].
func Lookup(username string) (*User, error)

// LookupId looks up a user by userid. If the user cannot be found, the
// returned error is of type [UnknownUserIdError].
func LookupId(uid string) (*User, error)

// LookupGroup looks up a group by name. If the group cannot be found, the
// returned error is of type [UnknownGroupError].
func LookupGroup(name string) (*Group, error)

// LookupGroupId looks up a group by groupid. If the group cannot be found, the
// returned error is of type [UnknownGroupIdError].
>>>>>>> upstream/master
func LookupGroupId(gid string) (*Group, error)

// GroupIdsは、ユーザーがメンバーであるグループIDのリストを返します。
func (u *User) GroupIds() ([]string, error)
