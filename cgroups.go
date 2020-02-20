package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func cgroup() error {
	if err := os.MkdirAll("/sys/fs/cgroup/cpu/shoten", 0700); err != nil {
		return fmt.Errorf("Cgroups namespace shoten create failed: %w", err)
	}

	if err := ioutil.WriteFile(
		"/sys/fs/cgroup/cpu/shoten/tasks",
		[]byte(fmt.Sprintf("%d\n", os.Getpid())),
		0644,
	); err != nil {
		return fmt.Errorf("Cgroups register tasks to shoten namespace failed: %w", err)
	}

	if err := ioutil.WriteFile(
		"/sys/fs/cgroup/cpu/shoten/cpu.cfs_quota_us",
		[]byte("5000\n"),
		0644,
	); err != nil {
		return fmt.Errorf("Cgroups add limit cpu.cfs_quota_us to 5000 failed: %w", err)
	}

	return nil
}
