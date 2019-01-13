#include <linux/bpf.h>

#include "bpf_helpers.h"

SEC("cgroup/skb/firewall")
int firewall(struct __sk_buff *skb)
{
	printt("TODO implement firewall\n");

	return 1; // ALLOW
}

char _license[] SEC("license") = "GPL";
__u32 _version SEC("version") = 0xFFFFFFFE;
