.PHONY: all
all: build

.PHONY: build
build: build-load-firewall build-attach-firewall build-elf

.PHONY: build-load-firewall
build-load-firewall:
	go build -o load-firewall load-firewall.go

.PHONY: build-attach-firewall
build-attach-firewall:
	go build -o attach-firewall attach-firewall.go

.PHONY: build-elf
build-elf:
	clang \
		-D__KERNEL__ \
		-O2 -emit-llvm -c bpf/cgroup-firewall.c \
		-o - | \
		llc -march=bpf -filetype=obj -o cgroup-firewall.elf

.PHONY: clean
clean:
	rm -vf \
		cgroup-firewall.elf \
		load-firewall \
		attach-firewall
