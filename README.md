Demo code showing how to create and pin a `cgroup/skb` eBPF program that
can be loaded and used from a second program.

Note: the currently used gobpf version is https://github.com/iovisor/gobpf/pull/169

```
make
./load-firewall
mkdir /sys/fs/cgroup/unified/cgroup-firewall-demo
./attach-firewall /sys/fs/bpf/cgroup-firewall-demo /sys/fs/cgroup/unified/cgroup-firewall-demo/
```

Now open a new shell and add its pid to the cgroup:

```
echo PID > /sys/fs/cgroup/unified/cgroup-firewall-demo/cgroup.procs
```

Then run e.g. `ping -4 ipschwein.de` from the shell in the cgroup.

In the kernel's `trace_pipe` you should see a TODO message for each packet:

```
cat /sys/kernel/debug/tracing/trace_pipe
```
