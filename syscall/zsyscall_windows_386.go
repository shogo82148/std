// mksyscall_windows.pl -l32 syscall_windows.go security_windows.go syscall_windows_386.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package syscall

func GetLastError() (lasterr error)

func LoadLibrary(libname string) (handle Handle, err error)

func FreeLibrary(handle Handle) (err error)

func GetProcAddress(module Handle, procname string) (proc uintptr, err error)

func GetVersion() (ver uint32, err error)

func FormatMessage(flags uint32, msgsrc uint32, msgid uint32, langid uint32, buf []uint16, args *byte) (n uint32, err error)

func ExitProcess(exitcode uint32)

func CreateFile(name *uint16, access uint32, mode uint32, sa *SecurityAttributes, createmode uint32, attrs uint32, templatefile int32) (handle Handle, err error)

func ReadFile(handle Handle, buf []byte, done *uint32, overlapped *Overlapped) (err error)

func WriteFile(handle Handle, buf []byte, done *uint32, overlapped *Overlapped) (err error)

func SetFilePointer(handle Handle, lowoffset int32, highoffsetptr *int32, whence uint32) (newlowoffset uint32, err error)

func CloseHandle(handle Handle) (err error)

func GetStdHandle(stdhandle int) (handle Handle, err error)

func FindFirstFile(name *uint16, data *Win32finddata) (handle Handle, err error)

func FindNextFile(handle Handle, data *Win32finddata) (err error)

func FindClose(handle Handle) (err error)

func GetFileInformationByHandle(handle Handle, data *ByHandleFileInformation) (err error)

func GetCurrentDirectory(buflen uint32, buf *uint16) (n uint32, err error)

func SetCurrentDirectory(path *uint16) (err error)

func CreateDirectory(path *uint16, sa *SecurityAttributes) (err error)

func RemoveDirectory(path *uint16) (err error)

func DeleteFile(path *uint16) (err error)

func MoveFile(from *uint16, to *uint16) (err error)

func GetComputerName(buf *uint16, n *uint32) (err error)

func SetEndOfFile(handle Handle) (err error)

func GetSystemTimeAsFileTime(time *Filetime)

func GetTimeZoneInformation(tzi *Timezoneinformation) (rc uint32, err error)

func CreateIoCompletionPort(filehandle Handle, cphandle Handle, key uint32, threadcnt uint32) (handle Handle, err error)

func GetQueuedCompletionStatus(cphandle Handle, qty *uint32, key *uint32, overlapped **Overlapped, timeout uint32) (err error)

func PostQueuedCompletionStatus(cphandle Handle, qty uint32, key uint32, overlapped *Overlapped) (err error)

func CancelIo(s Handle) (err error)

func CreateProcess(appName *uint16, commandLine *uint16, procSecurity *SecurityAttributes, threadSecurity *SecurityAttributes, inheritHandles bool, creationFlags uint32, env *uint16, currentDir *uint16, startupInfo *StartupInfo, outProcInfo *ProcessInformation) (err error)

func OpenProcess(da uint32, inheritHandle bool, pid uint32) (handle Handle, err error)

func TerminateProcess(handle Handle, exitcode uint32) (err error)

func GetExitCodeProcess(handle Handle, exitcode *uint32) (err error)

func GetStartupInfo(startupInfo *StartupInfo) (err error)

func GetCurrentProcess() (pseudoHandle Handle, err error)

func GetProcessTimes(handle Handle, creationTime *Filetime, exitTime *Filetime, kernelTime *Filetime, userTime *Filetime) (err error)

func DuplicateHandle(hSourceProcessHandle Handle, hSourceHandle Handle, hTargetProcessHandle Handle, lpTargetHandle *Handle, dwDesiredAccess uint32, bInheritHandle bool, dwOptions uint32) (err error)

func WaitForSingleObject(handle Handle, waitMilliseconds uint32) (event uint32, err error)

func GetTempPath(buflen uint32, buf *uint16) (n uint32, err error)

func CreatePipe(readhandle *Handle, writehandle *Handle, sa *SecurityAttributes, size uint32) (err error)

func GetFileType(filehandle Handle) (n uint32, err error)

func CryptAcquireContext(provhandle *Handle, container *uint16, provider *uint16, provtype uint32, flags uint32) (err error)

func CryptReleaseContext(provhandle Handle, flags uint32) (err error)

func CryptGenRandom(provhandle Handle, buflen uint32, buf *byte) (err error)

func GetEnvironmentStrings() (envs *uint16, err error)

func FreeEnvironmentStrings(envs *uint16) (err error)

func GetEnvironmentVariable(name *uint16, buffer *uint16, size uint32) (n uint32, err error)

func SetEnvironmentVariable(name *uint16, value *uint16) (err error)

func SetFileTime(handle Handle, ctime *Filetime, atime *Filetime, wtime *Filetime) (err error)

func GetFileAttributes(name *uint16) (attrs uint32, err error)

func SetFileAttributes(name *uint16, attrs uint32) (err error)

func GetFileAttributesEx(name *uint16, level uint32, info *byte) (err error)

func GetCommandLine() (cmd *uint16)

func CommandLineToArgv(cmd *uint16, argc *int32) (argv *[8192]*[8192]uint16, err error)

func LocalFree(hmem Handle) (handle Handle, err error)

func SetHandleInformation(handle Handle, mask uint32, flags uint32) (err error)

func FlushFileBuffers(handle Handle) (err error)

func GetFullPathName(path *uint16, buflen uint32, buf *uint16, fname **uint16) (n uint32, err error)

func GetLongPathName(path *uint16, buf *uint16, buflen uint32) (n uint32, err error)

func GetShortPathName(longpath *uint16, shortpath *uint16, buflen uint32) (n uint32, err error)

func CreateFileMapping(fhandle Handle, sa *SecurityAttributes, prot uint32, maxSizeHigh uint32, maxSizeLow uint32, name *uint16) (handle Handle, err error)

func MapViewOfFile(handle Handle, access uint32, offsetHigh uint32, offsetLow uint32, length uintptr) (addr uintptr, err error)

func UnmapViewOfFile(addr uintptr) (err error)

func FlushViewOfFile(addr uintptr, length uintptr) (err error)

func VirtualLock(addr uintptr, length uintptr) (err error)

func VirtualUnlock(addr uintptr, length uintptr) (err error)

func TransmitFile(s Handle, handle Handle, bytesToWrite uint32, bytsPerSend uint32, overlapped *Overlapped, transmitFileBuf *TransmitFileBuffers, flags uint32) (err error)

func ReadDirectoryChanges(handle Handle, buf *byte, buflen uint32, watchSubTree bool, mask uint32, retlen *uint32, overlapped *Overlapped, completionRoutine uintptr) (err error)

func CertOpenSystemStore(hprov Handle, name *uint16) (store Handle, err error)

func CertOpenStore(storeProvider uintptr, msgAndCertEncodingType uint32, cryptProv uintptr, flags uint32, para uintptr) (handle Handle, err error)

func CertEnumCertificatesInStore(store Handle, prevContext *CertContext) (context *CertContext, err error)

func CertAddCertificateContextToStore(store Handle, certContext *CertContext, addDisposition uint32, storeContext **CertContext) (err error)

func CertCloseStore(store Handle, flags uint32) (err error)

func CertGetCertificateChain(engine Handle, leaf *CertContext, time *Filetime, additionalStore Handle, para *CertChainPara, flags uint32, reserved uintptr, chainCtx **CertChainContext) (err error)

func CertFreeCertificateChain(ctx *CertChainContext)

func CertCreateCertificateContext(certEncodingType uint32, certEncoded *byte, encodedLen uint32) (context *CertContext, err error)

func CertFreeCertificateContext(ctx *CertContext) (err error)

func CertVerifyCertificateChainPolicy(policyOID uintptr, chain *CertChainContext, para *CertChainPolicyPara, status *CertChainPolicyStatus) (err error)

func RegOpenKeyEx(key Handle, subkey *uint16, options uint32, desiredAccess uint32, result *Handle) (regerrno error)

func RegCloseKey(key Handle) (regerrno error)

func RegQueryInfoKey(key Handle, class *uint16, classLen *uint32, reserved *uint32, subkeysLen *uint32, maxSubkeyLen *uint32, maxClassLen *uint32, valuesLen *uint32, maxValueNameLen *uint32, maxValueLen *uint32, saLen *uint32, lastWriteTime *Filetime) (regerrno error)

func RegEnumKeyEx(key Handle, index uint32, name *uint16, nameLen *uint32, reserved *uint32, class *uint16, classLen *uint32, lastWriteTime *Filetime) (regerrno error)

func RegQueryValueEx(key Handle, name *uint16, reserved *uint32, valtype *uint32, buf *byte, buflen *uint32) (regerrno error)

func WSAStartup(verreq uint32, data *WSAData) (sockerr error)

func WSACleanup() (err error)

func WSAIoctl(s Handle, iocc uint32, inbuf *byte, cbif uint32, outbuf *byte, cbob uint32, cbbr *uint32, overlapped *Overlapped, completionRoutine uintptr) (err error)

func Setsockopt(s Handle, level int32, optname int32, optval *byte, optlen int32) (err error)

func Closesocket(s Handle) (err error)

func AcceptEx(ls Handle, as Handle, buf *byte, rxdatalen uint32, laddrlen uint32, raddrlen uint32, recvd *uint32, overlapped *Overlapped) (err error)

func GetAcceptExSockaddrs(buf *byte, rxdatalen uint32, laddrlen uint32, raddrlen uint32, lrsa **RawSockaddrAny, lrsalen *int32, rrsa **RawSockaddrAny, rrsalen *int32)

func WSARecv(s Handle, bufs *WSABuf, bufcnt uint32, recvd *uint32, flags *uint32, overlapped *Overlapped, croutine *byte) (err error)

func WSASend(s Handle, bufs *WSABuf, bufcnt uint32, sent *uint32, flags uint32, overlapped *Overlapped, croutine *byte) (err error)

func WSARecvFrom(s Handle, bufs *WSABuf, bufcnt uint32, recvd *uint32, flags *uint32, from *RawSockaddrAny, fromlen *int32, overlapped *Overlapped, croutine *byte) (err error)

func WSASendTo(s Handle, bufs *WSABuf, bufcnt uint32, sent *uint32, flags uint32, to *RawSockaddrAny, tolen int32, overlapped *Overlapped, croutine *byte) (err error)

func GetHostByName(name string) (h *Hostent, err error)

func GetServByName(name string, proto string) (s *Servent, err error)

func Ntohs(netshort uint16) (u uint16)

func GetProtoByName(name string) (p *Protoent, err error)

func DnsQuery(name string, qtype uint16, options uint32, extra *byte, qrs **DNSRecord, pr *byte) (status error)

func DnsRecordListFree(rl *DNSRecord, freetype uint32)

func GetIfEntry(pIfRow *MibIfRow) (errcode error)

func GetAdaptersInfo(ai *IpAdapterInfo, ol *uint32) (errcode error)

func TranslateName(accName *uint16, accNameFormat uint32, desiredNameFormat uint32, translatedName *uint16, nSize *uint32) (err error)

func GetUserNameEx(nameFormat uint32, nameBuffre *uint16, nSize *uint32) (err error)

func NetUserGetInfo(serverName *uint16, userName *uint16, level uint32, buf **byte) (neterr error)

func NetApiBufferFree(buf *byte) (neterr error)

func LookupAccountSid(systemName *uint16, sid *SID, name *uint16, nameLen *uint32, refdDomainName *uint16, refdDomainNameLen *uint32, use *uint32) (err error)

func LookupAccountName(systemName *uint16, accountName *uint16, sid *SID, sidLen *uint32, refdDomainName *uint16, refdDomainNameLen *uint32, use *uint32) (err error)

func ConvertSidToStringSid(sid *SID, stringSid **uint16) (err error)

func ConvertStringSidToSid(stringSid *uint16, sid **SID) (err error)

func GetLengthSid(sid *SID) (len uint32)

func CopySid(destSidLen uint32, destSid *SID, srcSid *SID) (err error)

func OpenProcessToken(h Handle, access uint32, token *Token) (err error)

func GetTokenInformation(t Token, infoClass uint32, info *byte, infoLen uint32, returnedLen *uint32) (err error)

func GetUserProfileDirectory(t Token, dir *uint16, dirLen *uint32) (err error)
