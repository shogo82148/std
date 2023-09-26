// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows && !plan9
// +build !windows,!plan9

// Package syslog provides a simple interface to the system log service. It
// can send messages to the syslog daemon using UNIX domain sockets, UDP, or
// TCP connections.
package syslog

import (
	"github.com/shogo82148/std/log"
)

type Priority int

const (
	// From /usr/include/sys/syslog.h.
	// These are the same on Linux, BSD, and OS X.
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

// A Writer is a connection to a syslog server.
type Writer struct {
	priority Priority
	prefix   string
	conn     serverConn
}

// New establishes a new connection to the system log daemon.
// Each write to the returned writer sends a log message with
// the given priority and prefix.
func New(priority Priority, prefix string) (w *Writer, err error)

// Dial establishes a connection to a log daemon by connecting
// to address raddr on the network net.
// Each write to the returned writer sends a log message with
// the given priority and prefix.
func Dial(network, raddr string, priority Priority, prefix string) (w *Writer, err error)

// Write sends a log message to the syslog daemon.
func (w *Writer) Write(b []byte) (int, error)

func (w *Writer) Close() error

// Emerg logs a message using the LOG_EMERG priority.
func (w *Writer) Emerg(m string) (err error)

// Alert logs a message using the LOG_ALERT priority.
func (w *Writer) Alert(m string) (err error)

// Crit logs a message using the LOG_CRIT priority.
func (w *Writer) Crit(m string) (err error)

// Err logs a message using the LOG_ERR priority.
func (w *Writer) Err(m string) (err error)

// Warning logs a message using the LOG_WARNING priority.
func (w *Writer) Warning(m string) (err error)

// Notice logs a message using the LOG_NOTICE priority.
func (w *Writer) Notice(m string) (err error)

// Info logs a message using the LOG_INFO priority.
func (w *Writer) Info(m string) (err error)

// Debug logs a message using the LOG_DEBUG priority.
func (w *Writer) Debug(m string) (err error)

// NewLogger creates a log.Logger whose output is written to
// the system log service with the specified priority. The logFlag
// argument is the flag set passed through to log.New to create
// the Logger.
func NewLogger(p Priority, logFlag int) (*log.Logger, error)
