// Code generated by 'go generate'; DO NOT EDIT.

package windows

import (
	"github.com/shogo82148/std/syscall"
	"github.com/shogo82148/std/unsafe"
)

var _ unsafe.Pointer

func DuplicateTokenEx(hExistingToken syscall.Token, dwDesiredAccess uint32, lpTokenAttributes *syscall.SecurityAttributes, impersonationLevel uint32, tokenType TokenType, phNewToken *syscall.Token) (err error)

func GetSidIdentifierAuthority(sid *syscall.SID) (idauth *SID_IDENTIFIER_AUTHORITY)

func GetSidSubAuthority(sid *syscall.SID, subAuthorityIdx uint32) (subAuth *uint32)

func GetSidSubAuthorityCount(sid *syscall.SID) (count *uint8)

func ImpersonateLoggedOnUser(token syscall.Token) (err error)

func ImpersonateSelf(impersonationlevel uint32) (err error)

func IsValidSid(sid *syscall.SID) (valid bool)

func LogonUser(username *uint16, domain *uint16, password *uint16, logonType uint32, logonProvider uint32, token *syscall.Token) (err error)

func LookupPrivilegeValue(systemname *uint16, name *uint16, luid *LUID) (err error)

func OpenSCManager(machineName *uint16, databaseName *uint16, access uint32) (handle syscall.Handle, err error)

func OpenService(mgr syscall.Handle, serviceName *uint16, access uint32) (handle syscall.Handle, err error)

func OpenThreadToken(h syscall.Handle, access uint32, openasself bool, token *syscall.Token) (err error)

func QueryServiceStatus(hService syscall.Handle, lpServiceStatus *SERVICE_STATUS) (err error)

func RevertToSelf() (err error)

func SetTokenInformation(tokenHandle syscall.Token, tokenInformationClass uint32, tokenInformation uintptr, tokenInformationLength uint32) (err error)

func ProcessPrng(buf []byte) (err error)

func GetAdaptersAddresses(family uint32, flags uint32, reserved uintptr, adapterAddresses *IpAdapterAddresses, sizePointer *uint32) (errcode error)

func CreateEvent(eventAttrs *SecurityAttributes, manualReset uint32, initialState uint32, name *uint16) (handle syscall.Handle, err error)

func GetACP() (acp uint32)

func GetComputerNameEx(nameformat uint32, buf *uint16, n *uint32) (err error)

func GetConsoleCP() (ccp uint32)

func GetCurrentThread() (pseudoHandle syscall.Handle, err error)

func GetFileInformationByHandleEx(handle syscall.Handle, class uint32, info *byte, bufsize uint32) (err error)

func GetFinalPathNameByHandle(file syscall.Handle, filePath *uint16, filePathSize uint32, flags uint32) (n uint32, err error)

func GetModuleFileName(module syscall.Handle, fn *uint16, len uint32) (n uint32, err error)

func GetModuleHandle(modulename *uint16) (handle syscall.Handle, err error)

func GetTempPath2(buflen uint32, buf *uint16) (n uint32, err error)

func GetVolumeInformationByHandle(file syscall.Handle, volumeNameBuffer *uint16, volumeNameSize uint32, volumeNameSerialNumber *uint32, maximumComponentLength *uint32, fileSystemFlags *uint32, fileSystemNameBuffer *uint16, fileSystemNameSize uint32) (err error)

func GetVolumeNameForVolumeMountPoint(volumeMountPoint *uint16, volumeName *uint16, bufferlength uint32) (err error)

func LockFileEx(file syscall.Handle, flags uint32, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *syscall.Overlapped) (err error)

func Module32First(snapshot syscall.Handle, moduleEntry *ModuleEntry32) (err error)

func Module32Next(snapshot syscall.Handle, moduleEntry *ModuleEntry32) (err error)

func MoveFileEx(from *uint16, to *uint16, flags uint32) (err error)

func MultiByteToWideChar(codePage uint32, dwFlags uint32, str *byte, nstr int32, wchar *uint16, nwchar int32) (nwrite int32, err error)

func RtlLookupFunctionEntry(pc uintptr, baseAddress *uintptr, table *byte) (ret uintptr)

func RtlVirtualUnwind(handlerType uint32, baseAddress uintptr, pc uintptr, entry uintptr, ctxt uintptr, data *uintptr, frame *uintptr, ctxptrs *byte) (ret uintptr)

func SetFileInformationByHandle(handle syscall.Handle, fileInformationClass uint32, buf unsafe.Pointer, bufsize uint32) (err error)

func UnlockFileEx(file syscall.Handle, reserved uint32, bytesLow uint32, bytesHigh uint32, overlapped *syscall.Overlapped) (err error)

func VirtualQuery(address uintptr, buffer *MemoryBasicInformation, length uintptr) (err error)

func NetShareAdd(serverName *uint16, level uint32, buf *byte, parmErr *uint16) (neterr error)

func NetShareDel(serverName *uint16, netName *uint16, reserved uint32) (neterr error)

func NetUserAdd(serverName *uint16, level uint32, buf *byte, parmErr *uint32) (neterr error)

func NetUserDel(serverName *uint16, userName *uint16) (neterr error)

func NetUserGetLocalGroups(serverName *uint16, userName *uint16, level uint32, flags uint32, buf **byte, prefMaxLen uint32, entriesRead *uint32, totalEntries *uint32) (neterr error)

func NtCreateFile(handle *syscall.Handle, access uint32, oa *OBJECT_ATTRIBUTES, iosb *IO_STATUS_BLOCK, allocationSize *int64, attributes uint32, share uint32, disposition uint32, options uint32, eabuffer uintptr, ealength uint32) (ntstatus error)

func GetProcessMemoryInfo(handle syscall.Handle, memCounters *PROCESS_MEMORY_COUNTERS, cb uint32) (err error)

func CreateEnvironmentBlock(block **uint16, token syscall.Token, inheritExisting bool) (err error)

func DestroyEnvironmentBlock(block *uint16) (err error)

func GetProfilesDirectory(dir *uint16, dirLen *uint32) (err error)

func WSAGetOverlappedResult(h syscall.Handle, o *syscall.Overlapped, bytes *uint32, wait bool, flags *uint32) (err error)

func WSASocket(af int32, typ int32, protocol int32, protinfo *syscall.WSAProtocolInfo, group uint32, flags uint32) (handle syscall.Handle, err error)
