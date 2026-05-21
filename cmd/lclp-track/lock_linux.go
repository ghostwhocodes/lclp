//go:build linux

package main

import (
	"os"
	"path/filepath"
	"syscall"
)

func withAuditReadLock(baseDir, slug string, fn func() error) error {
	return withAuditReadLockForInput(baseDir, true, slug, fn)
}

func withAuditReadLockForInput(baseDir string, explicit bool, slug string, fn func() error) error {
	return withAuditLockForInput(baseDir, explicit, slug, syscall.LOCK_SH, fn)
}

func withAuditWriteLock(baseDir, slug string, fn func() error) error {
	return withAuditWriteLockForInput(baseDir, true, slug, fn)
}

func withAuditWriteLockForInput(baseDir string, explicit bool, slug string, fn func() error) error {
	return withAuditLockForInput(baseDir, explicit, slug, syscall.LOCK_EX, fn)
}

func withAuditLock(baseDir, slug string, lockMode int, fn func() error) error {
	return withAuditLockForInput(baseDir, true, slug, lockMode, fn)
}

func withAuditLockForInput(baseDir string, explicit bool, slug string, lockMode int, fn func() error) error {
	lockPath, err := auditLockPathForInput(baseDir, explicit, slug)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(lockPath), 0o777); err != nil {
		return err
	}
	if err := validateAuditLeafPath(lockPath); err != nil {
		return err
	}

	lockFile, err := os.OpenFile(lockPath, os.O_CREATE|os.O_RDWR, 0o666)
	if err != nil {
		return err
	}
	defer lockFile.Close()

	if err := syscall.Flock(int(lockFile.Fd()), lockMode); err != nil {
		return err
	}
	defer syscall.Flock(int(lockFile.Fd()), syscall.LOCK_UN)

	return fn()
}
