package main

import (
	"fmt"
	"os"

	"github.com/iovisor/gobpf/elf"
	"github.com/iovisor/gobpf/pkg/bpffs"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <path/to/pinned/firewall>\n", os.Args[0])
		os.Exit(1)
	}

	progPath := os.Args[1]

	if err := bpffs.Mount(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to mount bpf fs: %v\n", err)
		os.Exit(1)
	}

	firewall := elf.NewModule("./cgroup-firewall.elf")
	if err := firewall.Load(nil); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load firewall: %v\n", err)
		os.Exit(1)
	}

	progFd := firewall.CgroupProgram("cgroup/skb/firewall").Fd()
	if err := elf.PinObject(progFd, progPath); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to pin firewall at %q: %v\n", progPath, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully pinned program at %q\n", progPath)
}
