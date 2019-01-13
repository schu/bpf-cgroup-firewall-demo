package main

import (
	"fmt"
	"os"

	"github.com/iovisor/gobpf/elf"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <path/to/pinned/firewall> <cgroup-path>\n", os.Args[0])
		os.Exit(1)
	}

	progPath := os.Args[1]

	progFd := elf.GetProgFd(progPath)
	if progFd <= 0 {
		fmt.Fprintf(os.Stderr, "Failed to get pinned obj fd\n")
		os.Exit(1)
	}

	cgroupPath := os.Args[2]
	if err := elf.AttachCgroupProgramFromFd(progFd, cgroupPath, elf.EgressType); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to attach program to cgroup: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully attached program at %q to cgroup %q\n", progPath, cgroupPath)
}
