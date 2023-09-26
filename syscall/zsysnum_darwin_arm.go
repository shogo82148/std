// mksysnum_darwin.pl /usr/include/sys/syscall.h
// Code generated by the command above; DO NOT EDIT.

//go:build arm && darwin
// +build arm,darwin

package syscall

const (
	SYS_SYSCALL                   = 0
	SYS_EXIT                      = 1
	SYS_FORK                      = 2
	SYS_READ                      = 3
	SYS_WRITE                     = 4
	SYS_OPEN                      = 5
	SYS_CLOSE                     = 6
	SYS_WAIT4                     = 7
	SYS_LINK                      = 9
	SYS_UNLINK                    = 10
	SYS_CHDIR                     = 12
	SYS_FCHDIR                    = 13
	SYS_MKNOD                     = 14
	SYS_CHMOD                     = 15
	SYS_CHOWN                     = 16
	SYS_OBREAK                    = 17
	SYS_OGETFSSTAT                = 18
	SYS_GETFSSTAT                 = 18
	SYS_GETPID                    = 20
	SYS_SETUID                    = 23
	SYS_GETUID                    = 24
	SYS_GETEUID                   = 25
	SYS_PTRACE                    = 26
	SYS_RECVMSG                   = 27
	SYS_SENDMSG                   = 28
	SYS_RECVFROM                  = 29
	SYS_ACCEPT                    = 30
	SYS_GETPEERNAME               = 31
	SYS_GETSOCKNAME               = 32
	SYS_ACCESS                    = 33
	SYS_CHFLAGS                   = 34
	SYS_FCHFLAGS                  = 35
	SYS_SYNC                      = 36
	SYS_KILL                      = 37
	SYS_GETPPID                   = 39
	SYS_DUP                       = 41
	SYS_PIPE                      = 42
	SYS_GETEGID                   = 43
	SYS_PROFIL                    = 44
	SYS_SIGACTION                 = 46
	SYS_GETGID                    = 47
	SYS_SIGPROCMASK               = 48
	SYS_GETLOGIN                  = 49
	SYS_SETLOGIN                  = 50
	SYS_ACCT                      = 51
	SYS_SIGPENDING                = 52
	SYS_SIGALTSTACK               = 53
	SYS_IOCTL                     = 54
	SYS_REBOOT                    = 55
	SYS_REVOKE                    = 56
	SYS_SYMLINK                   = 57
	SYS_READLINK                  = 58
	SYS_EXECVE                    = 59
	SYS_UMASK                     = 60
	SYS_CHROOT                    = 61
	SYS_MSYNC                     = 65
	SYS_VFORK                     = 66
	SYS_SBRK                      = 69
	SYS_SSTK                      = 70
	SYS_OVADVISE                  = 72
	SYS_MUNMAP                    = 73
	SYS_MPROTECT                  = 74
	SYS_MADVISE                   = 75
	SYS_MINCORE                   = 78
	SYS_GETGROUPS                 = 79
	SYS_SETGROUPS                 = 80
	SYS_GETPGRP                   = 81
	SYS_SETPGID                   = 82
	SYS_SETITIMER                 = 83
	SYS_SWAPON                    = 85
	SYS_GETITIMER                 = 86
	SYS_GETDTABLESIZE             = 89
	SYS_DUP2                      = 90
	SYS_FCNTL                     = 92
	SYS_SELECT                    = 93
	SYS_FSYNC                     = 95
	SYS_SETPRIORITY               = 96
	SYS_SOCKET                    = 97
	SYS_CONNECT                   = 98
	SYS_GETPRIORITY               = 100
	SYS_BIND                      = 104
	SYS_SETSOCKOPT                = 105
	SYS_LISTEN                    = 106
	SYS_SIGSUSPEND                = 111
	SYS_GETTIMEOFDAY              = 116
	SYS_GETRUSAGE                 = 117
	SYS_GETSOCKOPT                = 118
	SYS_READV                     = 120
	SYS_WRITEV                    = 121
	SYS_SETTIMEOFDAY              = 122
	SYS_FCHOWN                    = 123
	SYS_FCHMOD                    = 124
	SYS_SETREUID                  = 126
	SYS_SETREGID                  = 127
	SYS_RENAME                    = 128
	SYS_FLOCK                     = 131
	SYS_MKFIFO                    = 132
	SYS_SENDTO                    = 133
	SYS_SHUTDOWN                  = 134
	SYS_SOCKETPAIR                = 135
	SYS_MKDIR                     = 136
	SYS_RMDIR                     = 137
	SYS_UTIMES                    = 138
	SYS_FUTIMES                   = 139
	SYS_ADJTIME                   = 140
	SYS_GETHOSTUUID               = 142
	SYS_SETSID                    = 147
	SYS_GETPGID                   = 151
	SYS_SETPRIVEXEC               = 152
	SYS_PREAD                     = 153
	SYS_PWRITE                    = 154
	SYS_NFSSVC                    = 155
	SYS_STATFS                    = 157
	SYS_FSTATFS                   = 158
	SYS_UNMOUNT                   = 159
	SYS_GETFH                     = 161
	SYS_QUOTACTL                  = 165
	SYS_MOUNT                     = 167
	SYS_CSOPS                     = 169
	SYS_TABLE                     = 170
	SYS_WAITID                    = 173
	SYS_ADD_PROFIL                = 176
	SYS_KDEBUG_TRACE              = 180
	SYS_SETGID                    = 181
	SYS_SETEGID                   = 182
	SYS_SETEUID                   = 183
	SYS_SIGRETURN                 = 184
	SYS_CHUD                      = 185
	SYS_STAT                      = 188
	SYS_FSTAT                     = 189
	SYS_LSTAT                     = 190
	SYS_PATHCONF                  = 191
	SYS_FPATHCONF                 = 192
	SYS_GETRLIMIT                 = 194
	SYS_SETRLIMIT                 = 195
	SYS_GETDIRENTRIES             = 196
	SYS_MMAP                      = 197
	SYS_LSEEK                     = 199
	SYS_TRUNCATE                  = 200
	SYS_FTRUNCATE                 = 201
	SYS___SYSCTL                  = 202
	SYS_MLOCK                     = 203
	SYS_MUNLOCK                   = 204
	SYS_UNDELETE                  = 205
	SYS_ATSOCKET                  = 206
	SYS_ATGETMSG                  = 207
	SYS_ATPUTMSG                  = 208
	SYS_ATPSNDREQ                 = 209
	SYS_ATPSNDRSP                 = 210
	SYS_ATPGETREQ                 = 211
	SYS_ATPGETRSP                 = 212
	SYS_KQUEUE_FROM_PORTSET_NP    = 214
	SYS_KQUEUE_PORTSET_NP         = 215
	SYS_MKCOMPLEX                 = 216
	SYS_STATV                     = 217
	SYS_LSTATV                    = 218
	SYS_FSTATV                    = 219
	SYS_GETATTRLIST               = 220
	SYS_SETATTRLIST               = 221
	SYS_GETDIRENTRIESATTR         = 222
	SYS_EXCHANGEDATA              = 223
	SYS_SEARCHFS                  = 225
	SYS_DELETE                    = 226
	SYS_COPYFILE                  = 227
	SYS_POLL                      = 230
	SYS_WATCHEVENT                = 231
	SYS_WAITEVENT                 = 232
	SYS_MODWATCH                  = 233
	SYS_GETXATTR                  = 234
	SYS_FGETXATTR                 = 235
	SYS_SETXATTR                  = 236
	SYS_FSETXATTR                 = 237
	SYS_REMOVEXATTR               = 238
	SYS_FREMOVEXATTR              = 239
	SYS_LISTXATTR                 = 240
	SYS_FLISTXATTR                = 241
	SYS_FSCTL                     = 242
	SYS_INITGROUPS                = 243
	SYS_POSIX_SPAWN               = 244
	SYS_NFSCLNT                   = 247
	SYS_FHOPEN                    = 248
	SYS_MINHERIT                  = 250
	SYS_SEMSYS                    = 251
	SYS_MSGSYS                    = 252
	SYS_SHMSYS                    = 253
	SYS_SEMCTL                    = 254
	SYS_SEMGET                    = 255
	SYS_SEMOP                     = 256
	SYS_MSGCTL                    = 258
	SYS_MSGGET                    = 259
	SYS_MSGSND                    = 260
	SYS_MSGRCV                    = 261
	SYS_SHMAT                     = 262
	SYS_SHMCTL                    = 263
	SYS_SHMDT                     = 264
	SYS_SHMGET                    = 265
	SYS_SHM_OPEN                  = 266
	SYS_SHM_UNLINK                = 267
	SYS_SEM_OPEN                  = 268
	SYS_SEM_CLOSE                 = 269
	SYS_SEM_UNLINK                = 270
	SYS_SEM_WAIT                  = 271
	SYS_SEM_TRYWAIT               = 272
	SYS_SEM_POST                  = 273
	SYS_SEM_GETVALUE              = 274
	SYS_SEM_INIT                  = 275
	SYS_SEM_DESTROY               = 276
	SYS_OPEN_EXTENDED             = 277
	SYS_UMASK_EXTENDED            = 278
	SYS_STAT_EXTENDED             = 279
	SYS_LSTAT_EXTENDED            = 280
	SYS_FSTAT_EXTENDED            = 281
	SYS_CHMOD_EXTENDED            = 282
	SYS_FCHMOD_EXTENDED           = 283
	SYS_ACCESS_EXTENDED           = 284
	SYS_SETTID                    = 285
	SYS_GETTID                    = 286
	SYS_SETSGROUPS                = 287
	SYS_GETSGROUPS                = 288
	SYS_SETWGROUPS                = 289
	SYS_GETWGROUPS                = 290
	SYS_MKFIFO_EXTENDED           = 291
	SYS_MKDIR_EXTENDED            = 292
	SYS_IDENTITYSVC               = 293
	SYS_SHARED_REGION_CHECK_NP    = 294
	SYS_SHARED_REGION_MAP_NP      = 295
	SYS___PTHREAD_MUTEX_DESTROY   = 301
	SYS___PTHREAD_MUTEX_INIT      = 302
	SYS___PTHREAD_MUTEX_LOCK      = 303
	SYS___PTHREAD_MUTEX_TRYLOCK   = 304
	SYS___PTHREAD_MUTEX_UNLOCK    = 305
	SYS___PTHREAD_COND_INIT       = 306
	SYS___PTHREAD_COND_DESTROY    = 307
	SYS___PTHREAD_COND_BROADCAST  = 308
	SYS___PTHREAD_COND_SIGNAL     = 309
	SYS_GETSID                    = 310
	SYS_SETTID_WITH_PID           = 311
	SYS___PTHREAD_COND_TIMEDWAIT  = 312
	SYS_AIO_FSYNC                 = 313
	SYS_AIO_RETURN                = 314
	SYS_AIO_SUSPEND               = 315
	SYS_AIO_CANCEL                = 316
	SYS_AIO_ERROR                 = 317
	SYS_AIO_READ                  = 318
	SYS_AIO_WRITE                 = 319
	SYS_LIO_LISTIO                = 320
	SYS___PTHREAD_COND_WAIT       = 321
	SYS_IOPOLICYSYS               = 322
	SYS_MLOCKALL                  = 324
	SYS_MUNLOCKALL                = 325
	SYS_ISSETUGID                 = 327
	SYS___PTHREAD_KILL            = 328
	SYS___PTHREAD_SIGMASK         = 329
	SYS___SIGWAIT                 = 330
	SYS___DISABLE_THREADSIGNAL    = 331
	SYS___PTHREAD_MARKCANCEL      = 332
	SYS___PTHREAD_CANCELED        = 333
	SYS___SEMWAIT_SIGNAL          = 334
	SYS_PROC_INFO                 = 336
	SYS_SENDFILE                  = 337
	SYS_STAT64                    = 338
	SYS_FSTAT64                   = 339
	SYS_LSTAT64                   = 340
	SYS_STAT64_EXTENDED           = 341
	SYS_LSTAT64_EXTENDED          = 342
	SYS_FSTAT64_EXTENDED          = 343
	SYS_GETDIRENTRIES64           = 344
	SYS_STATFS64                  = 345
	SYS_FSTATFS64                 = 346
	SYS_GETFSSTAT64               = 347
	SYS___PTHREAD_CHDIR           = 348
	SYS___PTHREAD_FCHDIR          = 349
	SYS_AUDIT                     = 350
	SYS_AUDITON                   = 351
	SYS_GETAUID                   = 353
	SYS_SETAUID                   = 354
	SYS_GETAUDIT                  = 355
	SYS_SETAUDIT                  = 356
	SYS_GETAUDIT_ADDR             = 357
	SYS_SETAUDIT_ADDR             = 358
	SYS_AUDITCTL                  = 359
	SYS_BSDTHREAD_CREATE          = 360
	SYS_BSDTHREAD_TERMINATE       = 361
	SYS_KQUEUE                    = 362
	SYS_KEVENT                    = 363
	SYS_LCHOWN                    = 364
	SYS_STACK_SNAPSHOT            = 365
	SYS_BSDTHREAD_REGISTER        = 366
	SYS_WORKQ_OPEN                = 367
	SYS_WORKQ_OPS                 = 368
	SYS___MAC_EXECVE              = 380
	SYS___MAC_SYSCALL             = 381
	SYS___MAC_GET_FILE            = 382
	SYS___MAC_SET_FILE            = 383
	SYS___MAC_GET_LINK            = 384
	SYS___MAC_SET_LINK            = 385
	SYS___MAC_GET_PROC            = 386
	SYS___MAC_SET_PROC            = 387
	SYS___MAC_GET_FD              = 388
	SYS___MAC_SET_FD              = 389
	SYS___MAC_GET_PID             = 390
	SYS___MAC_GET_LCID            = 391
	SYS___MAC_GET_LCTX            = 392
	SYS___MAC_SET_LCTX            = 393
	SYS_SETLCID                   = 394
	SYS_GETLCID                   = 395
	SYS_READ_NOCANCEL             = 396
	SYS_WRITE_NOCANCEL            = 397
	SYS_OPEN_NOCANCEL             = 398
	SYS_CLOSE_NOCANCEL            = 399
	SYS_WAIT4_NOCANCEL            = 400
	SYS_RECVMSG_NOCANCEL          = 401
	SYS_SENDMSG_NOCANCEL          = 402
	SYS_RECVFROM_NOCANCEL         = 403
	SYS_ACCEPT_NOCANCEL           = 404
	SYS_MSYNC_NOCANCEL            = 405
	SYS_FCNTL_NOCANCEL            = 406
	SYS_SELECT_NOCANCEL           = 407
	SYS_FSYNC_NOCANCEL            = 408
	SYS_CONNECT_NOCANCEL          = 409
	SYS_SIGSUSPEND_NOCANCEL       = 410
	SYS_READV_NOCANCEL            = 411
	SYS_WRITEV_NOCANCEL           = 412
	SYS_SENDTO_NOCANCEL           = 413
	SYS_PREAD_NOCANCEL            = 414
	SYS_PWRITE_NOCANCEL           = 415
	SYS_WAITID_NOCANCEL           = 416
	SYS_POLL_NOCANCEL             = 417
	SYS_MSGSND_NOCANCEL           = 418
	SYS_MSGRCV_NOCANCEL           = 419
	SYS_SEM_WAIT_NOCANCEL         = 420
	SYS_AIO_SUSPEND_NOCANCEL      = 421
	SYS___SIGWAIT_NOCANCEL        = 422
	SYS___SEMWAIT_SIGNAL_NOCANCEL = 423
	SYS___MAC_MOUNT               = 424
	SYS___MAC_GET_MOUNT           = 425
	SYS___MAC_GETFSSTAT           = 426
	SYS_MAXSYSCALL                = 427
)
