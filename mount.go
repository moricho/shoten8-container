package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func pivotRoot(newroot string) error {
	putold := filepath.Join(newroot, "/oldrootfs")

	// pivot_rootの条件を満たすために、新たなrootで自分自身をバインドマウント
	if err := syscall.Mount(
		newroot,
		newroot,
		"",
		syscall.MS_BIND|syscall.MS_REC,
		"",
	); err != nil {
		return err
	}

	if err := os.MkdirAll(putold, 0700); err != nil {
		return err
	}

	if err := syscall.PivotRoot(newroot, putold); err != nil {
		return err
	}

	if err := os.Chdir("/"); err != nil {
		return err
	}

	putold = "/oldrootfs"
	if err := syscall.Unmount(putold, syscall.MNT_DETACH); err != nil {
		return err
	}

	if err := os.RemoveAll(putold); err != nil {
		return err
	}

	return nil
}

func mountProc(newroot string) error {
	target := filepath.Join(newroot, "/proc")
	os.MkdirAll(target, 0755)
	if err := syscall.Mount("proc", target, "proc", uintptr(0), ""); err != nil {
		return err
	}

	return nil
}

func exitIfRootfsNotFound(rootfsPath string) {
	if _, err := os.Stat(rootfsPath); os.IsNotExist(err) {
		errorMsg := fmt.Sprint("rootfsPath not set")
		fmt.Println(errorMsg)
		os.Exit(1)
	}
}
