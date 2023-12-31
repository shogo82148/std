// mksysnum_freebsd.pl
// Code generated by the command above; DO NOT EDIT.

//go:build arm64 && freebsd

package syscall

const (
	// SYS_NOSYS = 0;  // { int nosys(void); } syscall nosys_args int
	SYS_EXIT                     = 1
	SYS_FORK                     = 2
	SYS_READ                     = 3
	SYS_WRITE                    = 4
	SYS_OPEN                     = 5
	SYS_CLOSE                    = 6
	SYS_WAIT4                    = 7
	SYS_LINK                     = 9
	SYS_UNLINK                   = 10
	SYS_CHDIR                    = 12
	SYS_FCHDIR                   = 13
	SYS_MKNOD                    = 14
	SYS_CHMOD                    = 15
	SYS_CHOWN                    = 16
	SYS_OBREAK                   = 17
	SYS_GETPID                   = 20
	SYS_MOUNT                    = 21
	SYS_UNMOUNT                  = 22
	SYS_SETUID                   = 23
	SYS_GETUID                   = 24
	SYS_GETEUID                  = 25
	SYS_PTRACE                   = 26
	SYS_RECVMSG                  = 27
	SYS_SENDMSG                  = 28
	SYS_RECVFROM                 = 29
	SYS_ACCEPT                   = 30
	SYS_GETPEERNAME              = 31
	SYS_GETSOCKNAME              = 32
	SYS_ACCESS                   = 33
	SYS_CHFLAGS                  = 34
	SYS_FCHFLAGS                 = 35
	SYS_SYNC                     = 36
	SYS_KILL                     = 37
	SYS_GETPPID                  = 39
	SYS_DUP                      = 41
	SYS_PIPE                     = 42
	SYS_GETEGID                  = 43
	SYS_PROFIL                   = 44
	SYS_KTRACE                   = 45
	SYS_GETGID                   = 47
	SYS_GETLOGIN                 = 49
	SYS_SETLOGIN                 = 50
	SYS_ACCT                     = 51
	SYS_SIGALTSTACK              = 53
	SYS_IOCTL                    = 54
	SYS_REBOOT                   = 55
	SYS_REVOKE                   = 56
	SYS_SYMLINK                  = 57
	SYS_READLINK                 = 58
	SYS_EXECVE                   = 59
	SYS_UMASK                    = 60
	SYS_CHROOT                   = 61
	SYS_MSYNC                    = 65
	SYS_VFORK                    = 66
	SYS_SBRK                     = 69
	SYS_SSTK                     = 70
	SYS_OVADVISE                 = 72
	SYS_MUNMAP                   = 73
	SYS_MPROTECT                 = 74
	SYS_MADVISE                  = 75
	SYS_MINCORE                  = 78
	SYS_GETGROUPS                = 79
	SYS_SETGROUPS                = 80
	SYS_GETPGRP                  = 81
	SYS_SETPGID                  = 82
	SYS_SETITIMER                = 83
	SYS_SWAPON                   = 85
	SYS_GETITIMER                = 86
	SYS_GETDTABLESIZE            = 89
	SYS_DUP2                     = 90
	SYS_FCNTL                    = 92
	SYS_SELECT                   = 93
	SYS_FSYNC                    = 95
	SYS_SETPRIORITY              = 96
	SYS_SOCKET                   = 97
	SYS_CONNECT                  = 98
	SYS_GETPRIORITY              = 100
	SYS_BIND                     = 104
	SYS_SETSOCKOPT               = 105
	SYS_LISTEN                   = 106
	SYS_GETTIMEOFDAY             = 116
	SYS_GETRUSAGE                = 117
	SYS_GETSOCKOPT               = 118
	SYS_READV                    = 120
	SYS_WRITEV                   = 121
	SYS_SETTIMEOFDAY             = 122
	SYS_FCHOWN                   = 123
	SYS_FCHMOD                   = 124
	SYS_SETREUID                 = 126
	SYS_SETREGID                 = 127
	SYS_RENAME                   = 128
	SYS_FLOCK                    = 131
	SYS_MKFIFO                   = 132
	SYS_SENDTO                   = 133
	SYS_SHUTDOWN                 = 134
	SYS_SOCKETPAIR               = 135
	SYS_MKDIR                    = 136
	SYS_RMDIR                    = 137
	SYS_UTIMES                   = 138
	SYS_ADJTIME                  = 140
	SYS_SETSID                   = 147
	SYS_QUOTACTL                 = 148
	SYS_NLM_SYSCALL              = 154
	SYS_NFSSVC                   = 155
	SYS_LGETFH                   = 160
	SYS_GETFH                    = 161
	SYS_SYSARCH                  = 165
	SYS_RTPRIO                   = 166
	SYS_SEMSYS                   = 169
	SYS_MSGSYS                   = 170
	SYS_SHMSYS                   = 171
	SYS_SETFIB                   = 175
	SYS_NTP_ADJTIME              = 176
	SYS_SETGID                   = 181
	SYS_SETEGID                  = 182
	SYS_SETEUID                  = 183
	SYS_PATHCONF                 = 191
	SYS_FPATHCONF                = 192
	SYS_GETRLIMIT                = 194
	SYS_SETRLIMIT                = 195
	SYS___SYSCTL                 = 202
	SYS_MLOCK                    = 203
	SYS_MUNLOCK                  = 204
	SYS_UNDELETE                 = 205
	SYS_FUTIMES                  = 206
	SYS_GETPGID                  = 207
	SYS_POLL                     = 209
	SYS_SEMGET                   = 221
	SYS_SEMOP                    = 222
	SYS_MSGGET                   = 225
	SYS_MSGSND                   = 226
	SYS_MSGRCV                   = 227
	SYS_SHMAT                    = 228
	SYS_SHMDT                    = 230
	SYS_SHMGET                   = 231
	SYS_CLOCK_GETTIME            = 232
	SYS_CLOCK_SETTIME            = 233
	SYS_CLOCK_GETRES             = 234
	SYS_KTIMER_CREATE            = 235
	SYS_KTIMER_DELETE            = 236
	SYS_KTIMER_SETTIME           = 237
	SYS_KTIMER_GETTIME           = 238
	SYS_KTIMER_GETOVERRUN        = 239
	SYS_NANOSLEEP                = 240
	SYS_FFCLOCK_GETCOUNTER       = 241
	SYS_FFCLOCK_SETESTIMATE      = 242
	SYS_FFCLOCK_GETESTIMATE      = 243
	SYS_CLOCK_NANOSLEEP          = 244
	SYS_CLOCK_GETCPUCLOCKID2     = 247
	SYS_NTP_GETTIME              = 248
	SYS_MINHERIT                 = 250
	SYS_RFORK                    = 251
	SYS_OPENBSD_POLL             = 252
	SYS_ISSETUGID                = 253
	SYS_LCHOWN                   = 254
	SYS_AIO_READ                 = 255
	SYS_AIO_WRITE                = 256
	SYS_LIO_LISTIO               = 257
	SYS_GETDENTS                 = 272
	SYS_LCHMOD                   = 274
	SYS_LUTIMES                  = 276
	SYS_NSTAT                    = 278
	SYS_NFSTAT                   = 279
	SYS_NLSTAT                   = 280
	SYS_PREADV                   = 289
	SYS_PWRITEV                  = 290
	SYS_FHOPEN                   = 298
	SYS_FHSTAT                   = 299
	SYS_MODNEXT                  = 300
	SYS_MODSTAT                  = 301
	SYS_MODFNEXT                 = 302
	SYS_MODFIND                  = 303
	SYS_KLDLOAD                  = 304
	SYS_KLDUNLOAD                = 305
	SYS_KLDFIND                  = 306
	SYS_KLDNEXT                  = 307
	SYS_KLDSTAT                  = 308
	SYS_KLDFIRSTMOD              = 309
	SYS_GETSID                   = 310
	SYS_SETRESUID                = 311
	SYS_SETRESGID                = 312
	SYS_AIO_RETURN               = 314
	SYS_AIO_SUSPEND              = 315
	SYS_AIO_CANCEL               = 316
	SYS_AIO_ERROR                = 317
	SYS_YIELD                    = 321
	SYS_MLOCKALL                 = 324
	SYS_MUNLOCKALL               = 325
	SYS___GETCWD                 = 326
	SYS_SCHED_SETPARAM           = 327
	SYS_SCHED_GETPARAM           = 328
	SYS_SCHED_SETSCHEDULER       = 329
	SYS_SCHED_GETSCHEDULER       = 330
	SYS_SCHED_YIELD              = 331
	SYS_SCHED_GET_PRIORITY_MAX   = 332
	SYS_SCHED_GET_PRIORITY_MIN   = 333
	SYS_SCHED_RR_GET_INTERVAL    = 334
	SYS_UTRACE                   = 335
	SYS_KLDSYM                   = 337
	SYS_JAIL                     = 338
	SYS_SIGPROCMASK              = 340
	SYS_SIGSUSPEND               = 341
	SYS_SIGPENDING               = 343
	SYS_SIGTIMEDWAIT             = 345
	SYS_SIGWAITINFO              = 346
	SYS___ACL_GET_FILE           = 347
	SYS___ACL_SET_FILE           = 348
	SYS___ACL_GET_FD             = 349
	SYS___ACL_SET_FD             = 350
	SYS___ACL_DELETE_FILE        = 351
	SYS___ACL_DELETE_FD          = 352
	SYS___ACL_ACLCHECK_FILE      = 353
	SYS___ACL_ACLCHECK_FD        = 354
	SYS_EXTATTRCTL               = 355
	SYS_EXTATTR_SET_FILE         = 356
	SYS_EXTATTR_GET_FILE         = 357
	SYS_EXTATTR_DELETE_FILE      = 358
	SYS_AIO_WAITCOMPLETE         = 359
	SYS_GETRESUID                = 360
	SYS_GETRESGID                = 361
	SYS_KQUEUE                   = 362
	SYS_KEVENT                   = 363
	SYS_EXTATTR_SET_FD           = 371
	SYS_EXTATTR_GET_FD           = 372
	SYS_EXTATTR_DELETE_FD        = 373
	SYS___SETUGID                = 374
	SYS_EACCESS                  = 376
	SYS_NMOUNT                   = 378
	SYS___MAC_GET_PROC           = 384
	SYS___MAC_SET_PROC           = 385
	SYS___MAC_GET_FD             = 386
	SYS___MAC_GET_FILE           = 387
	SYS___MAC_SET_FD             = 388
	SYS___MAC_SET_FILE           = 389
	SYS_KENV                     = 390
	SYS_LCHFLAGS                 = 391
	SYS_UUIDGEN                  = 392
	SYS_SENDFILE                 = 393
	SYS_MAC_SYSCALL              = 394
	SYS_FHSTATFS                 = 398
	SYS_KSEM_CLOSE               = 400
	SYS_KSEM_POST                = 401
	SYS_KSEM_WAIT                = 402
	SYS_KSEM_TRYWAIT             = 403
	SYS_KSEM_INIT                = 404
	SYS_KSEM_OPEN                = 405
	SYS_KSEM_UNLINK              = 406
	SYS_KSEM_GETVALUE            = 407
	SYS_KSEM_DESTROY             = 408
	SYS___MAC_GET_PID            = 409
	SYS___MAC_GET_LINK           = 410
	SYS___MAC_SET_LINK           = 411
	SYS_EXTATTR_SET_LINK         = 412
	SYS_EXTATTR_GET_LINK         = 413
	SYS_EXTATTR_DELETE_LINK      = 414
	SYS___MAC_EXECVE             = 415
	SYS_SIGACTION                = 416
	SYS_SIGRETURN                = 417
	SYS_GETCONTEXT               = 421
	SYS_SETCONTEXT               = 422
	SYS_SWAPCONTEXT              = 423
	SYS_SWAPOFF                  = 424
	SYS___ACL_GET_LINK           = 425
	SYS___ACL_SET_LINK           = 426
	SYS___ACL_DELETE_LINK        = 427
	SYS___ACL_ACLCHECK_LINK      = 428
	SYS_SIGWAIT                  = 429
	SYS_THR_CREATE               = 430
	SYS_THR_EXIT                 = 431
	SYS_THR_SELF                 = 432
	SYS_THR_KILL                 = 433
	SYS_JAIL_ATTACH              = 436
	SYS_EXTATTR_LIST_FD          = 437
	SYS_EXTATTR_LIST_FILE        = 438
	SYS_EXTATTR_LIST_LINK        = 439
	SYS_KSEM_TIMEDWAIT           = 441
	SYS_THR_SUSPEND              = 442
	SYS_THR_WAKE                 = 443
	SYS_KLDUNLOADF               = 444
	SYS_AUDIT                    = 445
	SYS_AUDITON                  = 446
	SYS_GETAUID                  = 447
	SYS_SETAUID                  = 448
	SYS_GETAUDIT                 = 449
	SYS_SETAUDIT                 = 450
	SYS_GETAUDIT_ADDR            = 451
	SYS_SETAUDIT_ADDR            = 452
	SYS_AUDITCTL                 = 453
	SYS__UMTX_OP                 = 454
	SYS_THR_NEW                  = 455
	SYS_SIGQUEUE                 = 456
	SYS_KMQ_OPEN                 = 457
	SYS_KMQ_SETATTR              = 458
	SYS_KMQ_TIMEDRECEIVE         = 459
	SYS_KMQ_TIMEDSEND            = 460
	SYS_KMQ_NOTIFY               = 461
	SYS_KMQ_UNLINK               = 462
	SYS_ABORT2                   = 463
	SYS_THR_SET_NAME             = 464
	SYS_AIO_FSYNC                = 465
	SYS_RTPRIO_THREAD            = 466
	SYS_SCTP_PEELOFF             = 471
	SYS_SCTP_GENERIC_SENDMSG     = 472
	SYS_SCTP_GENERIC_SENDMSG_IOV = 473
	SYS_SCTP_GENERIC_RECVMSG     = 474
	SYS_PREAD                    = 475
	SYS_PWRITE                   = 476
	SYS_MMAP                     = 477
	SYS_LSEEK                    = 478
	SYS_TRUNCATE                 = 479
	SYS_FTRUNCATE                = 480
	SYS_THR_KILL2                = 481
	SYS_SHM_OPEN                 = 482
	SYS_SHM_UNLINK               = 483
	SYS_CPUSET                   = 484
	SYS_CPUSET_SETID             = 485
	SYS_CPUSET_GETID             = 486
	SYS_CPUSET_GETAFFINITY       = 487
	SYS_CPUSET_SETAFFINITY       = 488
	SYS_FACCESSAT                = 489
	SYS_FCHMODAT                 = 490
	SYS_FCHOWNAT                 = 491
	SYS_FEXECVE                  = 492
	SYS_FUTIMESAT                = 494
	SYS_LINKAT                   = 495
	SYS_MKDIRAT                  = 496
	SYS_MKFIFOAT                 = 497
	SYS_OPENAT                   = 499
	SYS_READLINKAT               = 500
	SYS_RENAMEAT                 = 501
	SYS_SYMLINKAT                = 502
	SYS_UNLINKAT                 = 503
	SYS_POSIX_OPENPT             = 504
	SYS_GSSD_SYSCALL             = 505
	SYS_JAIL_GET                 = 506
	SYS_JAIL_SET                 = 507
	SYS_JAIL_REMOVE              = 508
	SYS_CLOSEFROM                = 509
	SYS___SEMCTL                 = 510
	SYS_MSGCTL                   = 511
	SYS_SHMCTL                   = 512
	SYS_LPATHCONF                = 513
	SYS___CAP_RIGHTS_GET         = 515
	SYS_CAP_ENTER                = 516
	SYS_CAP_GETMODE              = 517
	SYS_PDFORK                   = 518
	SYS_PDKILL                   = 519
	SYS_PDGETPID                 = 520
	SYS_PSELECT                  = 522
	SYS_GETLOGINCLASS            = 523
	SYS_SETLOGINCLASS            = 524
	SYS_RCTL_GET_RACCT           = 525
	SYS_RCTL_GET_RULES           = 526
	SYS_RCTL_GET_LIMITS          = 527
	SYS_RCTL_ADD_RULE            = 528
	SYS_RCTL_REMOVE_RULE         = 529
	SYS_POSIX_FALLOCATE          = 530
	SYS_POSIX_FADVISE            = 531
	SYS_WAIT6                    = 532
	SYS_CAP_RIGHTS_LIMIT         = 533
	SYS_CAP_IOCTLS_LIMIT         = 534
	SYS_CAP_IOCTLS_GET           = 535
	SYS_CAP_FCNTLS_LIMIT         = 536
	SYS_CAP_FCNTLS_GET           = 537
	SYS_BINDAT                   = 538
	SYS_CONNECTAT                = 539
	SYS_CHFLAGSAT                = 540
	SYS_ACCEPT4                  = 541
	SYS_PIPE2                    = 542
	SYS_AIO_MLOCK                = 543
	SYS_PROCCTL                  = 544
	SYS_PPOLL                    = 545
	SYS_FUTIMENS                 = 546
	SYS_UTIMENSAT                = 547
	SYS_NUMA_GETAFFINITY         = 548
	SYS_NUMA_SETAFFINITY         = 549
	SYS_FDATASYNC                = 550
	SYS_FSTAT                    = 551
	SYS_FSTATAT                  = 552
	SYS_GETDIRENTRIES            = 554
	SYS_STATFS                   = 555
	SYS_FSTATFS                  = 556
	SYS_GETFSSTAT                = 557
	SYS_MKNODAT                  = 559
)
