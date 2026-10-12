// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

const (
	STANDARD_RIGHTS_REQUIRED = 0xf0000
	STANDARD_RIGHTS_READ     = 0x20000
	STANDARD_RIGHTS_WRITE    = 0x20000
	STANDARD_RIGHTS_EXECUTE  = 0x20000
	STANDARD_RIGHTS_ALL      = 0x1F0000
)

const (
	NameUnknown          = 0
	NameFullyQualifiedDN = 1
	NameSamCompatible    = 2
	NameDisplay          = 3
	NameUniqueId         = 6
	NameCanonical        = 7
	NameUserPrincipal    = 8
	NameCanonicalEx      = 9
	NameServicePrincipal = 10
	NameDnsDomain        = 12
)

// TranslateAccountNameは、ディレクトリーサービスのオブジェクト名を
// ある形式から別の形式に変換します。
func TranslateAccountName(username string, from, to uint32, initSize int) (string, error)

const (
	// do not reorder
	NetSetupUnknownStatus = iota
	NetSetupUnjoined
	NetSetupWorkgroupName
	NetSetupDomainName
)

type UserInfo10 struct {
	Name       *uint16
	Comment    *uint16
	UsrComment *uint16
	FullName   *uint16
}

const (
	// do not reorder
	SidTypeUser = 1 + iota
	SidTypeGroup
	SidTypeDomain
	SidTypeAlias
	SidTypeWellKnownGroup
	SidTypeDeletedAccount
	SidTypeInvalid
	SidTypeUnknown
	SidTypeComputer
	SidTypeLabel
)

// セキュリティー識別子（SID）構造体は、ユーザーまたはグループを一意に識別するために
// 使用される可変長の構造体です。
type SID struct{}

// StringToSidは、文字列形式のセキュリティー識別子sidを、
// 有効で機能するsidに変換します。
func StringToSid(s string) (*SID, error)

// LookupSIDは、アカウントのセキュリティー識別子sidと、
// そのアカウントが見つかったドメインの名前を取得します。
// systemには検索対象のコンピューターを指定します。
func LookupSID(system, account string) (sid *SID, domain string, accType uint32, err error)

// Stringは、sidを表示、保存、または送信に適した文字列形式に変換します。
func (sid *SID) String() (string, error)

// Lenは、有効なセキュリティー識別子sidの長さをバイト単位で返します。
func (sid *SID) Len() int

// Copyは、セキュリティー識別子sidの複製を作成します。
func (sid *SID) Copy() (*SID, error)

// LookupAccountは、このsidのアカウント名と、
// このsidが見つかった最初のドメインの名前を取得します。
// systemには検索対象のコンピューターを指定します。
func (sid *SID) LookupAccount(system string) (account, domain string, accType uint32, err error)

const (
	// do not reorder
	TOKEN_ASSIGN_PRIMARY = 1 << iota
	TOKEN_DUPLICATE
	TOKEN_IMPERSONATE
	TOKEN_QUERY
	TOKEN_QUERY_SOURCE
	TOKEN_ADJUST_PRIVILEGES
	TOKEN_ADJUST_GROUPS
	TOKEN_ADJUST_DEFAULT
	TOKEN_ADJUST_SESSIONID

	TOKEN_ALL_ACCESS = STANDARD_RIGHTS_REQUIRED |
		TOKEN_ASSIGN_PRIMARY |
		TOKEN_DUPLICATE |
		TOKEN_IMPERSONATE |
		TOKEN_QUERY |
		TOKEN_QUERY_SOURCE |
		TOKEN_ADJUST_PRIVILEGES |
		TOKEN_ADJUST_GROUPS |
		TOKEN_ADJUST_DEFAULT |
		TOKEN_ADJUST_SESSIONID
	TOKEN_READ  = STANDARD_RIGHTS_READ | TOKEN_QUERY
	TOKEN_WRITE = STANDARD_RIGHTS_WRITE |
		TOKEN_ADJUST_PRIVILEGES |
		TOKEN_ADJUST_GROUPS |
		TOKEN_ADJUST_DEFAULT
	TOKEN_EXECUTE = STANDARD_RIGHTS_EXECUTE
)

const (
	// do not reorder
	TokenUser = 1 + iota
	TokenGroups
	TokenPrivileges
	TokenOwner
	TokenPrimaryGroup
	TokenDefaultDacl
	TokenSource
	TokenType
	TokenImpersonationLevel
	TokenStatistics
	TokenRestrictedSids
	TokenSessionId
	TokenGroupsAndPrivileges
	TokenSessionReference
	TokenSandBoxInert
	TokenAuditPolicy
	TokenOrigin
	TokenElevationType
	TokenLinkedToken
	TokenElevation
	TokenHasRestrictions
	TokenAccessInformation
	TokenVirtualizationAllowed
	TokenVirtualizationEnabled
	TokenIntegrityLevel
	TokenUIAccess
	TokenMandatoryPolicy
	TokenLogonSid
	MaxTokenInfoClass
)

type SIDAndAttributes struct {
	Sid        *SID
	Attributes uint32
}

type Tokenuser struct {
	User SIDAndAttributes
}

type Tokenprimarygroup struct {
	PrimaryGroup *SID
}

// アクセストークンには、ログオンセッションのセキュリティー情報が含まれます。
// システムはユーザーがログオンするとアクセストークンを作成し、
// ユーザーの代わりに実行されるすべてのプロセスはトークンのコピーを持ちます。
// トークンは、ユーザー、ユーザーのグループ、およびユーザーの特権を識別します。
// システムはトークンを使用して、セキュリティー保護可能なオブジェクトへのアクセスを制御し、
// ローカルコンピューター上でさまざまなシステム関連の操作を実行するユーザーの能力を制御します。
type Token Handle

// OpenCurrentProcessTokenは、現在のプロセスに関連付けられた
// アクセストークンを開きます。
func OpenCurrentProcessToken() (Token, error)

// Closeは、アクセストークンへのアクセスを解放します。
func (t Token) Close() error

// GetTokenUserは、アクセストークンtのユーザーアカウント情報を取得します。
func (t Token) GetTokenUser() (*Tokenuser, error)

// GetTokenPrimaryGroupは、アクセストークンtのプライマリグループ情報を取得します。
// これは、このアクセストークンを使用するプロセスによって作成されたオブジェクトの
// プライマリグループとなるグループを表すSID構造体へのポインターです。
func (t Token) GetTokenPrimaryGroup() (*Tokenprimarygroup, error)

// GetUserProfileDirectoryは、アクセストークンtのユーザーのプロファイルの
// ルートディレクトリーへのパスを取得します。
func (t Token) GetUserProfileDirectory() (string, error)
