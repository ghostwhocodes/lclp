//go:build !linux

package main

import "errors"

var errUnsupportedPlatform = errors.New("lclp-track is supported on Linux only")

func withAuditReadLock(baseDir, slug string, fn func() error) error {
	return withAuditReadLockForInput(baseDir, true, slug, fn)
}

func withAuditReadLockForInput(baseDir string, explicit bool, slug string, fn func() error) error {
	return errUnsupportedPlatform
}

func withAuditWriteLock(baseDir, slug string, fn func() error) error {
	return withAuditWriteLockForInput(baseDir, true, slug, fn)
}

func withAuditWriteLockForInput(baseDir string, explicit bool, slug string, fn func() error) error {
	return errUnsupportedPlatform
}

func withAuditLock(baseDir, slug string, lockMode int, fn func() error) error {
	return withAuditLockForInput(baseDir, true, slug, lockMode, fn)
}

func withAuditLockForInput(baseDir string, explicit bool, slug string, lockMode int, fn func() error) error {
	return errUnsupportedPlatform
}
