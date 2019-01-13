package main

import (
	"fmt"
	"os"

	"github.com/iovisor/gobpf/elf"
	"github.com/iovisor/gobpf/pkg/bpffs"
)

func main() {
	if err := bpffs.Mount(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to mount bpf fs: %v\n", err)
		os.Exit(1)
	}

	firewall := elf.NewModule("./cgroup-firewall.elf")
	if err := firewall.Load(nil); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load firewall: %v\n", err)
		os.Exit(1)
	}

	pinPath := "/sys/fs/bpf/cgroup-firewall-demo"
	progFd := firewall.CgroupProgram("cgroup/skb/firewall").Fd()
	if err := elf.PinObject(progFd, pinPath); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to pin firewall at %q: %v\n", pinPath, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully pinned program at %q\n", pinPath)
}
