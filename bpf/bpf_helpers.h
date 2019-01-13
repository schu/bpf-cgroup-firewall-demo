#pragma once

/* helper macro to place programs, maps, license in
 * different sections in elf_bpf file. Section names
 * are interpreted by elf_bpf loader
 */
#define SEC(NAME) __attribute__((section(NAME), used))

#ifndef __inline
# define __inline __attribute__((always_inline))
#endif

static int (*bpf_trace_printk)(const char *fmt, int fmt_size, ...) =
	(void *) BPF_FUNC_trace_printk;

#define printt(fmt, ...)                                                   \
	({                                                                 \
		char ____fmt[] = fmt;                                      \
		bpf_trace_printk(____fmt, sizeof(____fmt), ##__VA_ARGS__); \
	})
