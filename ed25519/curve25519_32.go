/*
	Package main - transpiled by c2go version: v0.26.0 Erbium 2020-03-17

	If you have found any issues, please raise an issue at:
	https://github.com/elliotchance/c2go/
*/

// Warning (GCCAsmStmt):  /mnt/d/c/spake2-c/curve25519_32.h:46 : cannot transpile asm, will be ignored

package main

import (
	"unsafe"

	"github.com/elliotchance/c2go/noarch"
)

type __u_char uint8
type __u_short uint16
type __u_int uint32
type __u_long uint32
type __int8_t int8
type __uint8_t uint8
type __int16_t int16
type __uint16_t uint16
type __int32_t int32
type __uint32_t uint32
type __int64_t int32
type __uint64_t uint32
type __int_least8_t int8
type __uint_least8_t uint8
type __int_least16_t int16
type __uint_least16_t uint16
type __int_least32_t int32
type __uint_least32_t uint32
type __int_least64_t int64
type __uint_least64_t uint64
type __quad_t int32
type __u_quad_t uint32
type __intmax_t int32
type __uintmax_t uint32
type __dev_t uint32
type __uid_t uint32
type __gid_t uint32
type __ino_t uint32
type __ino64_t uint32
type __mode_t uint32
type __nlink_t uint32
type __off_t int32
type __off64_t int32
type __pid_t int32
type __fsid_t struct {
	__val [2]int32
}
type __clock_t int32
type __rlim_t uint32
type __rlim64_t uint32
type __id_t uint32
type __time_t int32
type __useconds_t uint32
type __suseconds_t int32
type __suseconds64_t int32
type __daddr_t int32
type __key_t int32
type __clockid_t int32
type __timer_t unsafe.Pointer
type __blksize_t int32
type __blkcnt_t int32
type __blkcnt64_t int32
type __fsblkcnt_t uint32
type __fsblkcnt64_t uint32
type __fsfilcnt_t uint32
type __fsfilcnt64_t uint32
type __fsword_t int32
type __ssize_t int32
type __syscall_slong_t int32
type __syscall_ulong_t uint32
type __loff_t __off64_t
type __caddr_t *byte
type __intptr_t int32
type __socklen_t uint32
type __sig_atomic_t int32
type int8_t int8
type int16_t int16
type int32_t int32
type int64_t int64
type uint8_t uint8
type uint16_t uint16
type uint32_t uint32
type uint64_t uint64
type int_least8_t __int_least8_t
type int_least16_t __int_least16_t
type int_least32_t __int_least32_t
type int_least64_t __int_least64_t
type uint_least8_t __uint_least8_t
type uint_least16_t __uint_least16_t
type uint_least32_t __uint_least32_t
type uint_least64_t __uint_least64_t
type int_fast8_t int8
type int_fast16_t int32
type int_fast32_t int32
type int_fast64_t int32
type uint_fast8_t uint8
type uint_fast16_t uint32
type uint_fast32_t uint32
type uint_fast64_t uint32
type intptr_t int32
type uintptr_t uint32
type intmax_t __intmax_t
type uintmax_t __uintmax_t
type fiat_25519_uint1 uint8
type fiat_25519_int1 int8
type fiat_25519_loose_field_element []uint32
type fiat_25519_tight_field_element []uint32

// fiat_25519_value_barrier_u32 - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:45
// Warning (GCCAsmStmt):  /mnt/d/c/spake2-c/curve25519_32.h:46 : cannot transpile asm, will be ignored
/*
 * Copyright (c) 2015-2020 the fiat-crypto authors.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */ //
/* The type fiat_25519_loose_field_element is a field element with loose bounds. */ //
/* Bounds: [[0x0 ~> 0xc000000], [0x0 ~> 0x6000000], [0x0 ~> 0xc000000], [0x0 ~> 0x6000000], [0x0 ~> 0xc000000], [0x0 ~> 0x6000000], [0x0 ~> 0xc000000], [0x0 ~> 0x6000000], [0x0 ~> 0xc000000], [0x0 ~> 0x6000000]] */ //
/* The type fiat_25519_tight_field_element is a field element with tight bounds. */ //
/* Bounds: [[0x0 ~> 0x4000000], [0x0 ~> 0x2000000], [0x0 ~> 0x4000000], [0x0 ~> 0x2000000], [0x0 ~> 0x4000000], [0x0 ~> 0x2000000], [0x0 ~> 0x4000000], [0x0 ~> 0x2000000], [0x0 ~> 0x4000000], [0x0 ~> 0x2000000]] */ //
/* no inputs */ //
//
func fiat_25519_value_barrier_u32(a uint32) uint32 {
	return uint32(a)
}

// fiat_25519_addcarryx_u26 - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:69
/*
 * The function fiat_25519_addcarryx_u26 is an addition with carry.
 *
 * Postconditions:
 *   out1 = (arg1 + arg2 + arg3) mod 2^26
 *   out2 = ⌊(arg1 + arg2 + arg3) / 2^26⌋
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [0x0 ~> 0x3ffffff]
 *   arg3: [0x0 ~> 0x3ffffff]
 * Output Bounds:
 *   out1: [0x0 ~> 0x3ffffff]
 *   out2: [0x0 ~> 0x1]
 */ //
//
func fiat_25519_addcarryx_u26(out1 *uint32, out2 *fiat_25519_uint1, arg1 fiat_25519_uint1, arg2 uint32, arg3 uint32) {
	var x1 uint32
	var x2 uint32
	var x3 fiat_25519_uint1
	x1 = uint32((uint32((uint32(uint8((arg1))) + uint32((uint32((arg2)))) + uint32((uint32((arg3))))))))
	x2 = x1 & uint32((uint32((uint32(67108863)))))
	x3 = fiat_25519_uint1((x1 >> uint64(int32(26))))
	*out1 = x2
	*out2 = x3
}

// fiat_25519_subborrowx_u26 - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:95
/*
 * The function fiat_25519_subborrowx_u26 is a subtraction with borrow.
 *
 * Postconditions:
 *   out1 = (-arg1 + arg2 + -arg3) mod 2^26
 *   out2 = -⌊(-arg1 + arg2 + -arg3) / 2^26⌋
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [0x0 ~> 0x3ffffff]
 *   arg3: [0x0 ~> 0x3ffffff]
 * Output Bounds:
 *   out1: [0x0 ~> 0x3ffffff]
 *   out2: [0x0 ~> 0x1]
 */ //
//
func fiat_25519_subborrowx_u26(out1 *uint32, out2 *fiat_25519_uint1, arg1 fiat_25519_uint1, arg2 uint32, arg3 uint32) {
	var x1 int32
	var x2 fiat_25519_int1
	var x3 uint32
	x1 = int32((uint32((uint32((arg2 - uint32((uint32((uint32(uint8((arg1))))))))))))) - int32(uint32(arg3))
	x2 = fiat_25519_int1((x1 >> uint64(int32(26))))
	x3 = uint32((uint32((uint32(int32((int32((x1))))) & uint32(67108863)))))
	*out1 = x3
	*out2 = fiat_25519_uint1((int32(0) - int32(int8((x2)))))
}

// fiat_25519_addcarryx_u25 - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:121
/*
 * The function fiat_25519_addcarryx_u25 is an addition with carry.
 *
 * Postconditions:
 *   out1 = (arg1 + arg2 + arg3) mod 2^25
 *   out2 = ⌊(arg1 + arg2 + arg3) / 2^25⌋
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [0x0 ~> 0x1ffffff]
 *   arg3: [0x0 ~> 0x1ffffff]
 * Output Bounds:
 *   out1: [0x0 ~> 0x1ffffff]
 *   out2: [0x0 ~> 0x1]
 */ //
//
func fiat_25519_addcarryx_u25(out1 *uint32, out2 *fiat_25519_uint1, arg1 fiat_25519_uint1, arg2 uint32, arg3 uint32) {
	var x1 uint32
	var x2 uint32
	var x3 fiat_25519_uint1
	x1 = uint32((uint32((uint32(uint8((arg1))) + uint32((uint32((arg2)))) + uint32((uint32((arg3))))))))
	x2 = x1 & uint32((uint32((uint32(33554431)))))
	x3 = fiat_25519_uint1((x1 >> uint64(int32(25))))
	*out1 = x2
	*out2 = x3
}

// fiat_25519_subborrowx_u25 - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:147
/*
 * The function fiat_25519_subborrowx_u25 is a subtraction with borrow.
 *
 * Postconditions:
 *   out1 = (-arg1 + arg2 + -arg3) mod 2^25
 *   out2 = -⌊(-arg1 + arg2 + -arg3) / 2^25⌋
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [0x0 ~> 0x1ffffff]
 *   arg3: [0x0 ~> 0x1ffffff]
 * Output Bounds:
 *   out1: [0x0 ~> 0x1ffffff]
 *   out2: [0x0 ~> 0x1]
 */ //
//
func fiat_25519_subborrowx_u25(out1 *uint32, out2 *fiat_25519_uint1, arg1 fiat_25519_uint1, arg2 uint32, arg3 uint32) {
	var x1 int32
	var x2 fiat_25519_int1
	var x3 uint32
	x1 = int32((uint32((uint32((arg2 - uint32((uint32((uint32(uint8((arg1))))))))))))) - int32(uint32(arg3))
	x2 = fiat_25519_int1((x1 >> uint64(int32(25))))
	x3 = uint32((uint32((uint32(int32((int32((x1))))) & uint32(33554431)))))
	*out1 = x3
	*out2 = fiat_25519_uint1((int32(0) - int32(int8((x2)))))
}

// fiat_25519_cmovznz_u32 - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:171
/*
 * The function fiat_25519_cmovznz_u32 is a single-word conditional move.
 *
 * Postconditions:
 *   out1 = (if arg1 = 0 then arg2 else arg3)
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [0x0 ~> 0xffffffff]
 *   arg3: [0x0 ~> 0xffffffff]
 * Output Bounds:
 *   out1: [0x0 ~> 0xffffffff]
 */ //
//
func fiat_25519_cmovznz_u32(out1 *uint32, arg1 fiat_25519_uint1, arg2 uint32, arg3 uint32) {
	var x1 fiat_25519_uint1
	var x2 uint32
	var x3 uint32
	x1 = fiat_25519_uint1((uint8(noarch.NotInt32((int32(uint8((noarch.NotFiat_25519_uint1(fiat_25519_uint1(arg1))))))))))
	x2 = uint32((uint32((uint32(int8((fiat_25519_int1((int32(0) - int32(uint8((x1)))))))) & uint32(4294967295)))))
	x3 = fiat_25519_value_barrier_u32(uint32(x2))&arg3 | fiat_25519_value_barrier_u32((^uint32(x2)))&arg2
	*out1 = x3
}

// fiat_25519_carry_mul - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:188
/*
 * The function fiat_25519_carry_mul multiplies two field elements and reduces the result.
 *
 * Postconditions:
 *   eval out1 mod m = (eval arg1 * eval arg2) mod m
 *
 */ //
//
func fiat_25519_carry_mul(out1 *uint32, arg1 *uint32, arg2 *uint32) {
	var x1 uint64
	var x2 uint64
	var x3 uint64
	var x4 uint64
	var x5 uint64
	var x6 uint64
	var x7 uint64
	var x8 uint64
	var x9 uint64
	var x10 uint64
	var x11 uint64
	var x12 uint64
	var x13 uint64
	var x14 uint64
	var x15 uint64
	var x16 uint64
	var x17 uint64
	var x18 uint64
	var x19 uint64
	var x20 uint64
	var x21 uint64
	var x22 uint64
	var x23 uint64
	var x24 uint64
	var x25 uint64
	var x26 uint64
	var x27 uint64
	var x28 uint64
	var x29 uint64
	var x30 uint64
	var x31 uint64
	var x32 uint64
	var x33 uint64
	var x34 uint64
	var x35 uint64
	var x36 uint64
	var x37 uint64
	var x38 uint64
	var x39 uint64
	var x40 uint64
	var x41 uint64
	var x42 uint64
	var x43 uint64
	var x44 uint64
	var x45 uint64
	var x46 uint64
	var x47 uint64
	var x48 uint64
	var x49 uint64
	var x50 uint64
	var x51 uint64
	var x52 uint64
	var x53 uint64
	var x54 uint64
	var x55 uint64
	var x56 uint64
	var x57 uint64
	var x58 uint64
	var x59 uint64
	var x60 uint64
	var x61 uint64
	var x62 uint64
	var x63 uint64
	var x64 uint64
	var x65 uint64
	var x66 uint64
	var x67 uint64
	var x68 uint64
	var x69 uint64
	var x70 uint64
	var x71 uint64
	var x72 uint64
	var x73 uint64
	var x74 uint64
	var x75 uint64
	var x76 uint64
	var x77 uint64
	var x78 uint64
	var x79 uint64
	var x80 uint64
	var x81 uint64
	var x82 uint64
	var x83 uint64
	var x84 uint64
	var x85 uint64
	var x86 uint64
	var x87 uint64
	var x88 uint64
	var x89 uint64
	var x90 uint64
	var x91 uint64
	var x92 uint64
	var x93 uint64
	var x94 uint64
	var x95 uint64
	var x96 uint64
	var x97 uint64
	var x98 uint64
	var x99 uint64
	var x100 uint64
	var x101 uint64
	var x102 uint64
	var x103 uint32
	var x104 uint64
	var x105 uint64
	var x106 uint64
	var x107 uint64
	var x108 uint64
	var x109 uint64
	var x110 uint64
	var x111 uint64
	var x112 uint64
	var x113 uint64
	var x114 uint64
	var x115 uint32
	var x116 uint64
	var x117 uint64
	var x118 uint32
	var x119 uint64
	var x120 uint64
	var x121 uint32
	var x122 uint64
	var x123 uint64
	var x124 uint32
	var x125 uint64
	var x126 uint64
	var x127 uint32
	var x128 uint64
	var x129 uint64
	var x130 uint32
	var x131 uint64
	var x132 uint64
	var x133 uint32
	var x134 uint64
	var x135 uint64
	var x136 uint32
	var x137 uint64
	var x138 uint64
	var x139 uint32
	var x140 uint64
	var x141 uint64
	var x142 uint32
	var x143 uint32
	var x144 uint32
	var x145 fiat_25519_uint1
	var x146 uint32
	var x147 uint32
	x1 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x2 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x3 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x4 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x5 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x6 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x7 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x8 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x9 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x10 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x11 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x12 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x13 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x14 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x15 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x16 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x17 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x18 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x19 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x20 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x21 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x22 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x23 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x24 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x25 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x26 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x27 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x28 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x29 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x30 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x31 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x32 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x33 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x34 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x35 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x36 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x37 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x38 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x39 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x40 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x41 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x42 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x43 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x44 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(19))))))))))))))
	x45 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(38))))))))))))))
	x46 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*arg2))))))))
	x47 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg2))))))))))))
	x48 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*arg2))))))))
	x49 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg2))))))))))))
	x50 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(2))))))))))))))
	x51 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*arg2))))))))
	x52 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg2))))))))))))
	x53 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg2))))))))))))
	x54 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg2))))))))))))
	x55 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*arg2))))))))
	x56 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg2))))))))))))
	x57 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(2))))))))))))))
	x58 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg2))))))))))))
	x59 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(2))))))))))))))
	x60 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*arg2))))))))
	x61 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg2))))))))))))
	x62 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg2))))))))))))
	x63 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg2))))))))))))
	x64 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg2))))))))))))
	x65 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg2))))))))))))
	x66 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*arg2))))))))
	x67 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg2))))))))))))
	x68 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(2))))))))))))))
	x69 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg2))))))))))))
	x70 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(2))))))))))))))
	x71 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg2))))))))))))
	x72 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(2))))))))))))))
	x73 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*arg2))))))))
	x74 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg2))))))))))))
	x75 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg2))))))))))))
	x76 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg2))))))))))))
	x77 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg2))))))))))))
	x78 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg2))))))))))))
	x79 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg2))))))))))))
	x80 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg2))))))))))))
	x81 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*arg2))))))))
	x82 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg2))))))))))))
	x83 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(2))))))))))))))
	x84 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg2))))))))))))
	x85 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(2))))))))))))))
	x86 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg2))))))))))))
	x87 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(2))))))))))))))
	x88 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg2))))))))))))
	x89 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg2)))) * uint32((uint32((uint32(int32(2))))))))))))))
	x90 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*arg2))))))))
	x91 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg2))))))))))))
	x92 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg2))))))))))))
	x93 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg2))))))))))))
	x94 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg2))))))))))))
	x95 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg2))))))))))))
	x96 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg2))))))))))))
	x97 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg2))))))))))))
	x98 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg2))))))))))))
	x99 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg2))))))))))))
	x100 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((*arg2))))))))
	x101 = x100 + (x45 + (x44 + (x42 + (x39 + (x35 + (x30 + (x24 + (x17 + x9))))))))
	x102 = x101 >> uint64(int32(26))
	x103 = uint32((uint32((uint64((x101 & uint64((uint64((uint32(67108863)))))))))))
	x104 = x91 + (x82 + (x74 + (x67 + (x61 + (x56 + (x52 + (x49 + (x47 + x46))))))))
	x105 = x92 + (x83 + (x75 + (x68 + (x62 + (x57 + (x53 + (x50 + (x48 + x1))))))))
	x106 = x93 + (x84 + (x76 + (x69 + (x63 + (x58 + (x54 + (x51 + (x10 + x2))))))))
	x107 = x94 + (x85 + (x77 + (x70 + (x64 + (x59 + (x55 + (x18 + (x11 + x3))))))))
	x108 = x95 + (x86 + (x78 + (x71 + (x65 + (x60 + (x25 + (x19 + (x12 + x4))))))))
	x109 = x96 + (x87 + (x79 + (x72 + (x66 + (x31 + (x26 + (x20 + (x13 + x5))))))))
	x110 = x97 + (x88 + (x80 + (x73 + (x36 + (x32 + (x27 + (x21 + (x14 + x6))))))))
	x111 = x98 + (x89 + (x81 + (x40 + (x37 + (x33 + (x28 + (x22 + (x15 + x7))))))))
	x112 = x99 + (x90 + (x43 + (x41 + (x38 + (x34 + (x29 + (x23 + (x16 + x8))))))))
	x113 = x102 + x112
	x114 = x113 >> uint64(int32(25))
	x115 = uint32((uint32((uint64((x113 & uint64((uint64((uint32(33554431)))))))))))
	x116 = x114 + x111
	x117 = x116 >> uint64(int32(26))
	x118 = uint32((uint32((uint64((x116 & uint64((uint64((uint32(67108863)))))))))))
	x119 = x117 + x110
	x120 = x119 >> uint64(int32(25))
	x121 = uint32((uint32((uint64((x119 & uint64((uint64((uint32(33554431)))))))))))
	x122 = x120 + x109
	x123 = x122 >> uint64(int32(26))
	x124 = uint32((uint32((uint64((x122 & uint64((uint64((uint32(67108863)))))))))))
	x125 = x123 + x108
	x126 = x125 >> uint64(int32(25))
	x127 = uint32((uint32((uint64((x125 & uint64((uint64((uint32(33554431)))))))))))
	x128 = x126 + x107
	x129 = x128 >> uint64(int32(26))
	x130 = uint32((uint32((uint64((x128 & uint64((uint64((uint32(67108863)))))))))))
	x131 = x129 + x106
	x132 = x131 >> uint64(int32(25))
	x133 = uint32((uint32((uint64((x131 & uint64((uint64((uint32(33554431)))))))))))
	x134 = x132 + x105
	x135 = x134 >> uint64(int32(26))
	x136 = uint32((uint32((uint64((x134 & uint64((uint64((uint32(67108863)))))))))))
	x137 = x135 + x104
	x138 = x137 >> uint64(int32(25))
	x139 = uint32((uint32((uint64((x137 & uint64((uint64((uint32(33554431)))))))))))
	x140 = x138 * uint64((uint64((uint32(int32(19))))))
	x141 = uint64((uint64((uint32((uint32((x103)))) + uint32((uint64((x140))))))))
	x142 = uint32((x141 >> uint64(int32(26))))
	x143 = uint32((uint32((uint64((x141 & uint64((uint64((uint32(67108863)))))))))))
	x144 = x142 + x115
	x145 = fiat_25519_uint1((x144 >> uint64(int32(25))))
	x146 = x144 & uint32((uint32((uint32(33554431)))))
	x147 = uint32((uint32((uint32(uint8((x145))) + uint32((uint32((x118))))))))
	*out1 = x143
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(1))*unsafe.Sizeof(*out1)))) = x146
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(2))*unsafe.Sizeof(*out1)))) = x147
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(3))*unsafe.Sizeof(*out1)))) = x121
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(4))*unsafe.Sizeof(*out1)))) = x124
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(5))*unsafe.Sizeof(*out1)))) = x127
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(6))*unsafe.Sizeof(*out1)))) = x130
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(7))*unsafe.Sizeof(*out1)))) = x133
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(8))*unsafe.Sizeof(*out1)))) = x136
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(9))*unsafe.Sizeof(*out1)))) = x139
}

// fiat_25519_carry_square - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:502
/*
 * The function fiat_25519_carry_square squares a field element and reduces the result.
 *
 * Postconditions:
 *   eval out1 mod m = (eval arg1 * eval arg1) mod m
 *
 */ //
//
func fiat_25519_carry_square(out1 *uint32, arg1 *uint32) {
	var x1 uint32
	var x2 uint32
	var x3 uint32
	var x4 uint32
	var x5 uint64
	var x6 uint32
	var x7 uint32
	var x8 uint32
	var x9 uint32
	var x10 uint32
	var x11 uint64
	var x12 uint32
	var x13 uint32
	var x14 uint32
	var x15 uint32
	var x16 uint32
	var x17 uint32
	var x18 uint32
	var x19 uint64
	var x20 uint64
	var x21 uint64
	var x22 uint64
	var x23 uint64
	var x24 uint64
	var x25 uint64
	var x26 uint64
	var x27 uint64
	var x28 uint64
	var x29 uint64
	var x30 uint64
	var x31 uint64
	var x32 uint64
	var x33 uint64
	var x34 uint64
	var x35 uint64
	var x36 uint64
	var x37 uint64
	var x38 uint64
	var x39 uint64
	var x40 uint64
	var x41 uint64
	var x42 uint64
	var x43 uint64
	var x44 uint64
	var x45 uint64
	var x46 uint64
	var x47 uint64
	var x48 uint64
	var x49 uint64
	var x50 uint64
	var x51 uint64
	var x52 uint64
	var x53 uint64
	var x54 uint64
	var x55 uint64
	var x56 uint64
	var x57 uint64
	var x58 uint64
	var x59 uint64
	var x60 uint64
	var x61 uint64
	var x62 uint64
	var x63 uint64
	var x64 uint64
	var x65 uint64
	var x66 uint64
	var x67 uint64
	var x68 uint64
	var x69 uint64
	var x70 uint64
	var x71 uint64
	var x72 uint64
	var x73 uint64
	var x74 uint64
	var x75 uint64
	var x76 uint32
	var x77 uint64
	var x78 uint64
	var x79 uint64
	var x80 uint64
	var x81 uint64
	var x82 uint64
	var x83 uint64
	var x84 uint64
	var x85 uint64
	var x86 uint64
	var x87 uint64
	var x88 uint32
	var x89 uint64
	var x90 uint64
	var x91 uint32
	var x92 uint64
	var x93 uint64
	var x94 uint32
	var x95 uint64
	var x96 uint64
	var x97 uint32
	var x98 uint64
	var x99 uint64
	var x100 uint32
	var x101 uint64
	var x102 uint64
	var x103 uint32
	var x104 uint64
	var x105 uint64
	var x106 uint32
	var x107 uint64
	var x108 uint64
	var x109 uint32
	var x110 uint64
	var x111 uint64
	var x112 uint32
	var x113 uint64
	var x114 uint64
	var x115 uint32
	var x116 uint32
	var x117 uint32
	var x118 fiat_25519_uint1
	var x119 uint32
	var x120 uint32
	x1 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(19))))))
	x2 = x1 * uint32((uint32((uint32(int32(2))))))
	x3 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(2))))))
	x4 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(19))))))
	x5 = uint64(uint32(x4)) * uint64((uint64((uint32(int32(2))))))
	x6 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(2))))))
	x7 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(19))))))
	x8 = x7 * uint32((uint32((uint32(int32(2))))))
	x9 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(2))))))
	x10 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(19))))))
	x11 = uint64(uint32(x10)) * uint64((uint64((uint32(int32(2))))))
	x12 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(2))))))
	x13 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(19))))))
	x14 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(2))))))
	x15 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(2))))))
	x16 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(2))))))
	x17 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(2))))))
	x18 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(2))))))
	x19 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x1 * uint32((uint32((uint32(int32(2))))))))))))))
	x20 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x2))))))))
	x21 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x4))))))))
	x22 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((uint64(uint32(x2)) * uint64((uint64((uint32(int32(2))))))))))))))
	x23 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((x5))))))))
	x24 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x7 * uint32((uint32((uint32(int32(2))))))))))))))
	x25 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x2))))))))
	x26 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((x5))))))))
	x27 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x8))))))))
	x28 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x10))))))))
	x29 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((uint64(uint32(x2)) * uint64((uint64((uint32(int32(2))))))))))))))
	x30 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((x5))))))))
	x31 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((uint64(uint32(x8)) * uint64((uint64((uint32(int32(2))))))))))))))
	x32 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((x11))))))))
	x33 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x13 * uint32((uint32((uint32(int32(2))))))))))))))
	x34 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x2))))))))
	x35 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((x5))))))))
	x36 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x8))))))))
	x37 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((x11))))))))
	x38 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x14))))))))
	x39 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))))))))
	x40 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((uint64(uint32(x2)) * uint64((uint64((uint32(int32(2))))))))))))))
	x41 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((x5))))))))
	x42 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((uint64(uint32(x8)) * uint64((uint64((uint32(int32(2))))))))))))))
	x43 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x12))))))))
	x44 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x14 * uint32((uint32((uint32(int32(2))))))))))))))
	x45 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x15))))))))
	x46 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(2))))))))))))))
	x47 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x2))))))))
	x48 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((x5))))))))
	x49 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x9))))))))
	x50 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x12))))))))
	x51 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x14))))))))
	x52 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x15))))))))
	x53 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x16))))))))
	x54 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))))))))
	x55 = uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1)))))))) * uint32((uint64((uint64(uint32(x2)) * uint64((uint64((uint32(int32(2))))))))))))))
	x56 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x6))))))))
	x57 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x9 * uint32((uint32((uint32(int32(2))))))))))))))
	x58 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x12))))))))
	x59 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x14 * uint32((uint32((uint32(int32(2))))))))))))))
	x60 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x15))))))))
	x61 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x16 * uint32((uint32((uint32(int32(2))))))))))))))
	x62 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((x17))))))))
	x63 = uint64(uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1)))) * uint32((uint32((uint32(int32(2))))))))))))))
	x64 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((x3))))))))
	x65 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((x6))))))))
	x66 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((x9))))))))
	x67 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((x12))))))))
	x68 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((x14))))))))
	x69 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((x15))))))))
	x70 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((x16))))))))
	x71 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((x17))))))))
	x72 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((x18))))))))
	x73 = uint64(uint32((*arg1))) * uint64((uint64((uint32((uint32((*arg1))))))))
	x74 = x73 + (x55 + (x48 + (x42 + (x37 + x33))))
	x75 = x74 >> uint64(int32(26))
	x76 = uint32((uint32((uint64((x74 & uint64((uint64((uint32(67108863)))))))))))
	x77 = x64 + (x56 + (x49 + (x43 + x38)))
	x78 = x65 + (x57 + (x50 + (x44 + (x39 + x19))))
	x79 = x66 + (x58 + (x51 + (x45 + x20)))
	x80 = x67 + (x59 + (x52 + (x46 + (x22 + x21))))
	x81 = x68 + (x60 + (x53 + (x25 + x23)))
	x82 = x69 + (x61 + (x54 + (x29 + (x26 + x24))))
	x83 = x70 + (x62 + (x34 + (x30 + x27)))
	x84 = x71 + (x63 + (x40 + (x35 + (x31 + x28))))
	x85 = x72 + (x47 + (x41 + (x36 + x32)))
	x86 = x75 + x85
	x87 = x86 >> uint64(int32(25))
	x88 = uint32((uint32((uint64((x86 & uint64((uint64((uint32(33554431)))))))))))
	x89 = x87 + x84
	x90 = x89 >> uint64(int32(26))
	x91 = uint32((uint32((uint64((x89 & uint64((uint64((uint32(67108863)))))))))))
	x92 = x90 + x83
	x93 = x92 >> uint64(int32(25))
	x94 = uint32((uint32((uint64((x92 & uint64((uint64((uint32(33554431)))))))))))
	x95 = x93 + x82
	x96 = x95 >> uint64(int32(26))
	x97 = uint32((uint32((uint64((x95 & uint64((uint64((uint32(67108863)))))))))))
	x98 = x96 + x81
	x99 = x98 >> uint64(int32(25))
	x100 = uint32((uint32((uint64((x98 & uint64((uint64((uint32(33554431)))))))))))
	x101 = x99 + x80
	x102 = x101 >> uint64(int32(26))
	x103 = uint32((uint32((uint64((x101 & uint64((uint64((uint32(67108863)))))))))))
	x104 = x102 + x79
	x105 = x104 >> uint64(int32(25))
	x106 = uint32((uint32((uint64((x104 & uint64((uint64((uint32(33554431)))))))))))
	x107 = x105 + x78
	x108 = x107 >> uint64(int32(26))
	x109 = uint32((uint32((uint64((x107 & uint64((uint64((uint32(67108863)))))))))))
	x110 = x108 + x77
	x111 = x110 >> uint64(int32(25))
	x112 = uint32((uint32((uint64((x110 & uint64((uint64((uint32(33554431)))))))))))
	x113 = x111 * uint64((uint64((uint32(int32(19))))))
	x114 = uint64((uint64((uint32((uint32((x76)))) + uint32((uint64((x113))))))))
	x115 = uint32((x114 >> uint64(int32(26))))
	x116 = uint32((uint32((uint64((x114 & uint64((uint64((uint32(67108863)))))))))))
	x117 = x115 + x88
	x118 = fiat_25519_uint1((x117 >> uint64(int32(25))))
	x119 = x117 & uint32((uint32((uint32(33554431)))))
	x120 = uint32((uint32((uint32(uint8((x118))) + uint32((uint32((x91))))))))
	*out1 = x116
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(1))*unsafe.Sizeof(*out1)))) = x119
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(2))*unsafe.Sizeof(*out1)))) = x120
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(3))*unsafe.Sizeof(*out1)))) = x94
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(4))*unsafe.Sizeof(*out1)))) = x97
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(5))*unsafe.Sizeof(*out1)))) = x100
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(6))*unsafe.Sizeof(*out1)))) = x103
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(7))*unsafe.Sizeof(*out1)))) = x106
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(8))*unsafe.Sizeof(*out1)))) = x109
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(9))*unsafe.Sizeof(*out1)))) = x112
}

// fiat_25519_carry - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:762
/*
 * The function fiat_25519_carry reduces a field element.
 *
 * Postconditions:
 *   eval out1 mod m = eval arg1 mod m
 *
 */ //
//
func fiat_25519_carry(out1 *uint32, arg1 *uint32) {
	var x1 uint32
	var x2 uint32
	var x3 uint32
	var x4 uint32
	var x5 uint32
	var x6 uint32
	var x7 uint32
	var x8 uint32
	var x9 uint32
	var x10 uint32
	var x11 uint32
	var x12 uint32
	var x13 uint32
	var x14 uint32
	var x15 uint32
	var x16 uint32
	var x17 uint32
	var x18 uint32
	var x19 uint32
	var x20 uint32
	var x21 uint32
	var x22 uint32
	x1 = *arg1
	x2 = x1>>uint64(int32(26)) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))
	x3 = x2>>uint64(int32(25)) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))
	x4 = x3>>uint64(int32(26)) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))
	x5 = x4>>uint64(int32(25)) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))
	x6 = x5>>uint64(int32(26)) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))
	x7 = x6>>uint64(int32(25)) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))
	x8 = x7>>uint64(int32(26)) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))
	x9 = x8>>uint64(int32(25)) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))
	x10 = x9>>uint64(int32(26)) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))
	x11 = x1&uint32((uint32((uint32(67108863))))) + x10>>uint64(int32(25))*uint32((uint32((uint32(int32(19))))))
	x12 = uint32((uint32((uint32(uint8((fiat_25519_uint1((x11 >> uint64(int32(26))))))) + uint32((uint32((x2 & uint32((uint32((uint32(33554431)))))))))))))
	x13 = x11 & uint32((uint32((uint32(67108863)))))
	x14 = x12 & uint32((uint32((uint32(33554431)))))
	x15 = uint32((uint32((uint32(uint8((fiat_25519_uint1((x12 >> uint64(int32(25))))))) + uint32((uint32((x3 & uint32((uint32((uint32(67108863)))))))))))))
	x16 = x4 & uint32((uint32((uint32(33554431)))))
	x17 = x5 & uint32((uint32((uint32(67108863)))))
	x18 = x6 & uint32((uint32((uint32(33554431)))))
	x19 = x7 & uint32((uint32((uint32(67108863)))))
	x20 = x8 & uint32((uint32((uint32(33554431)))))
	x21 = x9 & uint32((uint32((uint32(67108863)))))
	x22 = x10 & uint32((uint32((uint32(33554431)))))
	*out1 = x13
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(1))*unsafe.Sizeof(*out1)))) = x14
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(2))*unsafe.Sizeof(*out1)))) = x15
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(3))*unsafe.Sizeof(*out1)))) = x16
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(4))*unsafe.Sizeof(*out1)))) = x17
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(5))*unsafe.Sizeof(*out1)))) = x18
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(6))*unsafe.Sizeof(*out1)))) = x19
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(7))*unsafe.Sizeof(*out1)))) = x20
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(8))*unsafe.Sizeof(*out1)))) = x21
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(9))*unsafe.Sizeof(*out1)))) = x22
}

// fiat_25519_add - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:826
/*
 * The function fiat_25519_add adds two field elements.
 *
 * Postconditions:
 *   eval out1 mod m = (eval arg1 + eval arg2) mod m
 *
 */ //
//
func fiat_25519_add(out1 *uint32, arg1 *uint32, arg2 *uint32) {
	var x1 uint32
	var x2 uint32
	var x3 uint32
	var x4 uint32
	var x5 uint32
	var x6 uint32
	var x7 uint32
	var x8 uint32
	var x9 uint32
	var x10 uint32
	x1 = *arg1 + *arg2
	x2 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1)))) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg2))))
	x3 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1)))) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg2))))
	x4 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1)))) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg2))))
	x5 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1)))) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg2))))
	x6 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1)))) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg2))))
	x7 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1)))) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg2))))
	x8 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1)))) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg2))))
	x9 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1)))) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg2))))
	x10 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1)))) + *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg2))))
	*out1 = x1
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(1))*unsafe.Sizeof(*out1)))) = x2
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(2))*unsafe.Sizeof(*out1)))) = x3
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(3))*unsafe.Sizeof(*out1)))) = x4
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(4))*unsafe.Sizeof(*out1)))) = x5
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(5))*unsafe.Sizeof(*out1)))) = x6
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(6))*unsafe.Sizeof(*out1)))) = x7
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(7))*unsafe.Sizeof(*out1)))) = x8
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(8))*unsafe.Sizeof(*out1)))) = x9
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(9))*unsafe.Sizeof(*out1)))) = x10
}

// fiat_25519_sub - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:866
/*
 * The function fiat_25519_sub subtracts two field elements.
 *
 * Postconditions:
 *   eval out1 mod m = (eval arg1 - eval arg2) mod m
 *
 */ //
//
func fiat_25519_sub(out1 *uint32, arg1 *uint32, arg2 *uint32) {
	var x1 uint32
	var x2 uint32
	var x3 uint32
	var x4 uint32
	var x5 uint32
	var x6 uint32
	var x7 uint32
	var x8 uint32
	var x9 uint32
	var x10 uint32
	x1 = uint32((uint32((uint32(134217690 + int32(uint32((uint32((*arg1))))) - int32(uint32((uint32((*arg2))))))))))
	x2 = uint32((uint32((uint32(67108862 + int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))))) - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg2))))))))))))))
	x3 = uint32((uint32((uint32(134217726 + int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))))) - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg2))))))))))))))
	x4 = uint32((uint32((uint32(67108862 + int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))))) - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg2))))))))))))))
	x5 = uint32((uint32((uint32(134217726 + int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))))) - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg2))))))))))))))
	x6 = uint32((uint32((uint32(67108862 + int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))))) - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg2))))))))))))))
	x7 = uint32((uint32((uint32(134217726 + int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))))) - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg2))))))))))))))
	x8 = uint32((uint32((uint32(67108862 + int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))))) - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg2))))))))))))))
	x9 = uint32((uint32((uint32(134217726 + int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))))) - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg2))))))))))))))
	x10 = uint32((uint32((uint32(67108862 + int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))))) - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg2))))))))))))))
	*out1 = x1
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(1))*unsafe.Sizeof(*out1)))) = x2
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(2))*unsafe.Sizeof(*out1)))) = x3
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(3))*unsafe.Sizeof(*out1)))) = x4
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(4))*unsafe.Sizeof(*out1)))) = x5
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(5))*unsafe.Sizeof(*out1)))) = x6
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(6))*unsafe.Sizeof(*out1)))) = x7
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(7))*unsafe.Sizeof(*out1)))) = x8
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(8))*unsafe.Sizeof(*out1)))) = x9
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(9))*unsafe.Sizeof(*out1)))) = x10
}

// fiat_25519_opp - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:906
/*
 * The function fiat_25519_opp negates a field element.
 *
 * Postconditions:
 *   eval out1 mod m = -eval arg1 mod m
 *
 */ //
//
func fiat_25519_opp(out1 *uint32, arg1 *uint32) {
	var x1 uint32
	var x2 uint32
	var x3 uint32
	var x4 uint32
	var x5 uint32
	var x6 uint32
	var x7 uint32
	var x8 uint32
	var x9 uint32
	var x10 uint32
	x1 = uint32((uint32((uint32(134217690 - int32(uint32((uint32((*arg1))))))))))
	x2 = uint32((uint32((uint32(67108862 - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))))))))))
	x3 = uint32((uint32((uint32(134217726 - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))))))))))
	x4 = uint32((uint32((uint32(67108862 - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))))))))))
	x5 = uint32((uint32((uint32(134217726 - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))))))))))
	x6 = uint32((uint32((uint32(67108862 - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))))))))))
	x7 = uint32((uint32((uint32(134217726 - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))))))))))
	x8 = uint32((uint32((uint32(67108862 - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))))))))))
	x9 = uint32((uint32((uint32(134217726 - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))))))))))
	x10 = uint32((uint32((uint32(67108862 - int32(uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))))))))))
	*out1 = x1
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(1))*unsafe.Sizeof(*out1)))) = x2
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(2))*unsafe.Sizeof(*out1)))) = x3
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(3))*unsafe.Sizeof(*out1)))) = x4
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(4))*unsafe.Sizeof(*out1)))) = x5
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(5))*unsafe.Sizeof(*out1)))) = x6
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(6))*unsafe.Sizeof(*out1)))) = x7
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(7))*unsafe.Sizeof(*out1)))) = x8
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(8))*unsafe.Sizeof(*out1)))) = x9
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(9))*unsafe.Sizeof(*out1)))) = x10
}

// fiat_25519_selectznz - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:952
/*
 * The function fiat_25519_selectznz is a multi-limb conditional select.
 *
 * Postconditions:
 *   eval out1 = (if arg1 = 0 then eval arg2 else eval arg3)
 *
 * Input Bounds:
 *   arg1: [0x0 ~> 0x1]
 *   arg2: [[0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff]]
 *   arg3: [[0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff]]
 * Output Bounds:
 *   out1: [[0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff], [0x0 ~> 0xffffffff]]
 */ //
//
func fiat_25519_selectznz(out1 *uint32, arg1 fiat_25519_uint1, arg2 *uint32, arg3 *uint32) {
	var x1 uint32
	var x2 uint32
	var x3 uint32
	var x4 uint32
	var x5 uint32
	var x6 uint32
	var x7 uint32
	var x8 uint32
	var x9 uint32
	var x10 uint32
	fiat_25519_cmovznz_u32(&x1, fiat_25519_uint1(arg1), uint32((*arg2)), uint32((*arg3)))
	fiat_25519_cmovznz_u32(&x2, fiat_25519_uint1(arg1), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg2)))))), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg3)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg3)))))))
	fiat_25519_cmovznz_u32(&x3, fiat_25519_uint1(arg1), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg2)))))), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg3)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg3)))))))
	fiat_25519_cmovznz_u32(&x4, fiat_25519_uint1(arg1), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg2)))))), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg3)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg3)))))))
	fiat_25519_cmovznz_u32(&x5, fiat_25519_uint1(arg1), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg2)))))), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg3)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg3)))))))
	fiat_25519_cmovznz_u32(&x6, fiat_25519_uint1(arg1), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg2)))))), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg3)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg3)))))))
	fiat_25519_cmovznz_u32(&x7, fiat_25519_uint1(arg1), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg2)))))), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg3)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg3)))))))
	fiat_25519_cmovznz_u32(&x8, fiat_25519_uint1(arg1), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg2)))))), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg3)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg3)))))))
	fiat_25519_cmovznz_u32(&x9, fiat_25519_uint1(arg1), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg2)))))), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg3)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg3)))))))
	fiat_25519_cmovznz_u32(&x10, fiat_25519_uint1(arg1), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg2)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg2)))))), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg3)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg3)))))))
	*out1 = x1
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(1))*unsafe.Sizeof(*out1)))) = x2
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(2))*unsafe.Sizeof(*out1)))) = x3
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(3))*unsafe.Sizeof(*out1)))) = x4
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(4))*unsafe.Sizeof(*out1)))) = x5
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(5))*unsafe.Sizeof(*out1)))) = x6
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(6))*unsafe.Sizeof(*out1)))) = x7
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(7))*unsafe.Sizeof(*out1)))) = x8
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(8))*unsafe.Sizeof(*out1)))) = x9
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(9))*unsafe.Sizeof(*out1)))) = x10
}

// fiat_25519_to_bytes - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:994
/*
 * The function fiat_25519_to_bytes serializes a field element to bytes in little-endian order.
 *
 * Postconditions:
 *   out1 = map (λ x, ⌊((eval arg1 mod m) mod 2^(8 * (x + 1))) / 2^(8 * x)⌋) [0..31]
 *
 * Output Bounds:
 *   out1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x7f]]
 */ //
//
func fiat_25519_to_bytes(out1 *uint8, arg1 *uint32) {
	var x1 uint32
	var x2 fiat_25519_uint1
	var x3 uint32
	var x4 fiat_25519_uint1
	var x5 uint32
	var x6 fiat_25519_uint1
	var x7 uint32
	var x8 fiat_25519_uint1
	var x9 uint32
	var x10 fiat_25519_uint1
	var x11 uint32
	var x12 fiat_25519_uint1
	var x13 uint32
	var x14 fiat_25519_uint1
	var x15 uint32
	var x16 fiat_25519_uint1
	var x17 uint32
	var x18 fiat_25519_uint1
	var x19 uint32
	var x20 fiat_25519_uint1
	var x21 uint32
	var x22 uint32
	var x23 fiat_25519_uint1
	var x24 uint32
	var x25 fiat_25519_uint1
	var x26 uint32
	var x27 fiat_25519_uint1
	var x28 uint32
	var x29 fiat_25519_uint1
	var x30 uint32
	var x31 fiat_25519_uint1
	var x32 uint32
	var x33 fiat_25519_uint1
	var x34 uint32
	var x35 fiat_25519_uint1
	var x36 uint32
	var x37 fiat_25519_uint1
	var x38 uint32
	var x39 fiat_25519_uint1
	var x40 uint32
	var x41 fiat_25519_uint1
	var x42 uint32
	var x43 uint32
	var x44 uint32
	var x45 uint32
	var x46 uint32
	var x47 uint32
	var x48 uint32
	var x49 uint32
	var x50 uint8
	var x51 uint32
	var x52 uint8
	var x53 uint32
	var x54 uint8
	var x55 uint8
	var x56 uint32
	var x57 uint8
	var x58 uint32
	var x59 uint8
	var x60 uint32
	var x61 uint8
	var x62 uint8
	var x63 uint32
	var x64 uint8
	var x65 uint32
	var x66 uint8
	var x67 uint32
	var x68 uint8
	var x69 uint8
	var x70 uint32
	var x71 uint8
	var x72 uint32
	var x73 uint8
	var x74 uint32
	var x75 uint8
	var x76 uint8
	var x77 uint32
	var x78 uint8
	var x79 uint32
	var x80 uint8
	var x81 uint32
	var x82 uint8
	var x83 uint8
	var x84 uint8
	var x85 uint32
	var x86 uint8
	var x87 uint32
	var x88 uint8
	var x89 fiat_25519_uint1
	var x90 uint32
	var x91 uint8
	var x92 uint32
	var x93 uint8
	var x94 uint32
	var x95 uint8
	var x96 uint8
	var x97 uint32
	var x98 uint8
	var x99 uint32
	var x100 uint8
	var x101 uint32
	var x102 uint8
	var x103 uint8
	var x104 uint32
	var x105 uint8
	var x106 uint32
	var x107 uint8
	var x108 uint32
	var x109 uint8
	var x110 uint8
	var x111 uint32
	var x112 uint8
	var x113 uint32
	var x114 uint8
	var x115 uint32
	var x116 uint8
	var x117 uint8
	fiat_25519_subborrowx_u26(&x1, &x2, fiat_25519_uint1(int32(0)), uint32((*arg1)), uint32((uint32((uint32(67108845))))))
	fiat_25519_subborrowx_u25(&x3, &x4, fiat_25519_uint1(x2), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1)))))), uint32((uint32((uint32(33554431))))))
	fiat_25519_subborrowx_u26(&x5, &x6, fiat_25519_uint1(x4), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1)))))), uint32((uint32((uint32(67108863))))))
	fiat_25519_subborrowx_u25(&x7, &x8, fiat_25519_uint1(x6), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1)))))), uint32((uint32((uint32(33554431))))))
	fiat_25519_subborrowx_u26(&x9, &x10, fiat_25519_uint1(x8), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1)))))), uint32((uint32((uint32(67108863))))))
	fiat_25519_subborrowx_u25(&x11, &x12, fiat_25519_uint1(x10), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1)))))), uint32((uint32((uint32(33554431))))))
	fiat_25519_subborrowx_u26(&x13, &x14, fiat_25519_uint1(x12), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1)))))), uint32((uint32((uint32(67108863))))))
	fiat_25519_subborrowx_u25(&x15, &x16, fiat_25519_uint1(x14), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1)))))), uint32((uint32((uint32(33554431))))))
	fiat_25519_subborrowx_u26(&x17, &x18, fiat_25519_uint1(x16), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1)))))), uint32((uint32((uint32(67108863))))))
	fiat_25519_subborrowx_u25(&x19, &x20, fiat_25519_uint1(x18), uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1)))))), uint32((uint32((uint32(33554431))))))
	fiat_25519_cmovznz_u32(&x21, fiat_25519_uint1(x20), uint32(int32(0)), uint32((uint32((uint32(4294967295))))))
	fiat_25519_addcarryx_u26(&x22, &x23, fiat_25519_uint1(int32(0)), uint32(x1), uint32((uint32((uint32((uint32((x21 & uint32((uint32((uint32(67108845))))))))))))))
	fiat_25519_addcarryx_u25(&x24, &x25, fiat_25519_uint1(x23), uint32(x3), uint32((uint32((uint32((uint32((x21 & uint32((uint32((uint32(33554431))))))))))))))
	fiat_25519_addcarryx_u26(&x26, &x27, fiat_25519_uint1(x25), uint32(x5), uint32((uint32((uint32((uint32((x21 & uint32((uint32((uint32(67108863))))))))))))))
	fiat_25519_addcarryx_u25(&x28, &x29, fiat_25519_uint1(x27), uint32(x7), uint32((uint32((uint32((uint32((x21 & uint32((uint32((uint32(33554431))))))))))))))
	fiat_25519_addcarryx_u26(&x30, &x31, fiat_25519_uint1(x29), uint32(x9), uint32((uint32((uint32((uint32((x21 & uint32((uint32((uint32(67108863))))))))))))))
	fiat_25519_addcarryx_u25(&x32, &x33, fiat_25519_uint1(x31), uint32(x11), uint32((uint32((uint32((uint32((x21 & uint32((uint32((uint32(33554431))))))))))))))
	fiat_25519_addcarryx_u26(&x34, &x35, fiat_25519_uint1(x33), uint32(x13), uint32((uint32((uint32((uint32((x21 & uint32((uint32((uint32(67108863))))))))))))))
	fiat_25519_addcarryx_u25(&x36, &x37, fiat_25519_uint1(x35), uint32(x15), uint32((uint32((uint32((uint32((x21 & uint32((uint32((uint32(33554431))))))))))))))
	fiat_25519_addcarryx_u26(&x38, &x39, fiat_25519_uint1(x37), uint32(x17), uint32((uint32((uint32((uint32((x21 & uint32((uint32((uint32(67108863))))))))))))))
	fiat_25519_addcarryx_u25(&x40, &x41, fiat_25519_uint1(x39), uint32(x19), uint32((uint32((uint32((uint32((x21 & uint32((uint32((uint32(33554431))))))))))))))
	x42 = x40 << uint64(int32(6))
	x43 = x38 << uint64(int32(4))
	x44 = x36 << uint64(int32(3))
	x45 = x34 * uint32(int32(2))
	x46 = x30 << uint64(int32(6))
	x47 = x28 << uint64(int32(5))
	x48 = x26 << uint64(int32(3))
	x49 = x24 << uint64(int32(2))
	x50 = uint8((uint32((uint32((x22 & uint32((uint32((uint32(int32(255))))))))))))
	x51 = x22 >> uint64(int32(8))
	x52 = uint8((uint32((uint32((x51 & uint32((uint32((uint32(int32(255))))))))))))
	x53 = x51 >> uint64(int32(8))
	x54 = uint8((uint32((uint32((x53 & uint32((uint32((uint32(int32(255))))))))))))
	x55 = uint8((x53 >> uint64(int32(8))))
	x56 = x49 + uint32(uint8(x55))
	x57 = uint8((uint32((uint32((x56 & uint32((uint32((uint32(int32(255))))))))))))
	x58 = x56 >> uint64(int32(8))
	x59 = uint8((uint32((uint32((x58 & uint32((uint32((uint32(int32(255))))))))))))
	x60 = x58 >> uint64(int32(8))
	x61 = uint8((uint32((uint32((x60 & uint32((uint32((uint32(int32(255))))))))))))
	x62 = uint8((x60 >> uint64(int32(8))))
	x63 = x48 + uint32(uint8(x62))
	x64 = uint8((uint32((uint32((x63 & uint32((uint32((uint32(int32(255))))))))))))
	x65 = x63 >> uint64(int32(8))
	x66 = uint8((uint32((uint32((x65 & uint32((uint32((uint32(int32(255))))))))))))
	x67 = x65 >> uint64(int32(8))
	x68 = uint8((uint32((uint32((x67 & uint32((uint32((uint32(int32(255))))))))))))
	x69 = uint8((x67 >> uint64(int32(8))))
	x70 = x47 + uint32(uint8(x69))
	x71 = uint8((uint32((uint32((x70 & uint32((uint32((uint32(int32(255))))))))))))
	x72 = x70 >> uint64(int32(8))
	x73 = uint8((uint32((uint32((x72 & uint32((uint32((uint32(int32(255))))))))))))
	x74 = x72 >> uint64(int32(8))
	x75 = uint8((uint32((uint32((x74 & uint32((uint32((uint32(int32(255))))))))))))
	x76 = uint8((x74 >> uint64(int32(8))))
	x77 = x46 + uint32(uint8(x76))
	x78 = uint8((uint32((uint32((x77 & uint32((uint32((uint32(int32(255))))))))))))
	x79 = x77 >> uint64(int32(8))
	x80 = uint8((uint32((uint32((x79 & uint32((uint32((uint32(int32(255))))))))))))
	x81 = x79 >> uint64(int32(8))
	x82 = uint8((uint32((uint32((x81 & uint32((uint32((uint32(int32(255))))))))))))
	x83 = uint8((x81 >> uint64(int32(8))))
	x84 = uint8((uint32((uint32((x32 & uint32((uint32((uint32(int32(255))))))))))))
	x85 = x32 >> uint64(int32(8))
	x86 = uint8((uint32((uint32((x85 & uint32((uint32((uint32(int32(255))))))))))))
	x87 = x85 >> uint64(int32(8))
	x88 = uint8((uint32((uint32((x87 & uint32((uint32((uint32(int32(255))))))))))))
	x89 = fiat_25519_uint1((x87 >> uint64(int32(8))))
	x90 = x45 + uint32(fiat_25519_uint1(x89))
	x91 = uint8((uint32((uint32((x90 & uint32((uint32((uint32(int32(255))))))))))))
	x92 = x90 >> uint64(int32(8))
	x93 = uint8((uint32((uint32((x92 & uint32((uint32((uint32(int32(255))))))))))))
	x94 = x92 >> uint64(int32(8))
	x95 = uint8((uint32((uint32((x94 & uint32((uint32((uint32(int32(255))))))))))))
	x96 = uint8((x94 >> uint64(int32(8))))
	x97 = x44 + uint32(uint8(x96))
	x98 = uint8((uint32((uint32((x97 & uint32((uint32((uint32(int32(255))))))))))))
	x99 = x97 >> uint64(int32(8))
	x100 = uint8((uint32((uint32((x99 & uint32((uint32((uint32(int32(255))))))))))))
	x101 = x99 >> uint64(int32(8))
	x102 = uint8((uint32((uint32((x101 & uint32((uint32((uint32(int32(255))))))))))))
	x103 = uint8((x101 >> uint64(int32(8))))
	x104 = x43 + uint32(uint8(x103))
	x105 = uint8((uint32((uint32((x104 & uint32((uint32((uint32(int32(255))))))))))))
	x106 = x104 >> uint64(int32(8))
	x107 = uint8((uint32((uint32((x106 & uint32((uint32((uint32(int32(255))))))))))))
	x108 = x106 >> uint64(int32(8))
	x109 = uint8((uint32((uint32((x108 & uint32((uint32((uint32(int32(255))))))))))))
	x110 = uint8((x108 >> uint64(int32(8))))
	x111 = x42 + uint32(uint8(x110))
	x112 = uint8((uint32((uint32((x111 & uint32((uint32((uint32(int32(255))))))))))))
	x113 = x111 >> uint64(int32(8))
	x114 = uint8((uint32((uint32((x113 & uint32((uint32((uint32(int32(255))))))))))))
	x115 = x113 >> uint64(int32(8))
	x116 = uint8((uint32((uint32((x115 & uint32((uint32((uint32(int32(255))))))))))))
	x117 = uint8((x115 >> uint64(int32(8))))
	*out1 = x50
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(1))*unsafe.Sizeof(*out1)))) = x52
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(2))*unsafe.Sizeof(*out1)))) = x54
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(3))*unsafe.Sizeof(*out1)))) = x57
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(4))*unsafe.Sizeof(*out1)))) = x59
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(5))*unsafe.Sizeof(*out1)))) = x61
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(6))*unsafe.Sizeof(*out1)))) = x64
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(7))*unsafe.Sizeof(*out1)))) = x66
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(8))*unsafe.Sizeof(*out1)))) = x68
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(9))*unsafe.Sizeof(*out1)))) = x71
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(10))*unsafe.Sizeof(*out1)))) = x73
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(11))*unsafe.Sizeof(*out1)))) = x75
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(12))*unsafe.Sizeof(*out1)))) = x78
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(13))*unsafe.Sizeof(*out1)))) = x80
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(14))*unsafe.Sizeof(*out1)))) = x82
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(15))*unsafe.Sizeof(*out1)))) = x83
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(16))*unsafe.Sizeof(*out1)))) = x84
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(17))*unsafe.Sizeof(*out1)))) = x86
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(18))*unsafe.Sizeof(*out1)))) = x88
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(19))*unsafe.Sizeof(*out1)))) = x91
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(20))*unsafe.Sizeof(*out1)))) = x93
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(21))*unsafe.Sizeof(*out1)))) = x95
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(22))*unsafe.Sizeof(*out1)))) = x98
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(23))*unsafe.Sizeof(*out1)))) = x100
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(24))*unsafe.Sizeof(*out1)))) = x102
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(25))*unsafe.Sizeof(*out1)))) = x105
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(26))*unsafe.Sizeof(*out1)))) = x107
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(27))*unsafe.Sizeof(*out1)))) = x109
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(28))*unsafe.Sizeof(*out1)))) = x112
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(29))*unsafe.Sizeof(*out1)))) = x114
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(30))*unsafe.Sizeof(*out1)))) = x116
	*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(31))*unsafe.Sizeof(*out1)))) = x117
}

// fiat_25519_from_bytes - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:1252
/*
 * The function fiat_25519_from_bytes deserializes a field element from bytes in little-endian order.
 *
 * Postconditions:
 *   eval out1 mod m = bytes_eval arg1 mod m
 *
 * Input Bounds:
 *   arg1: [[0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0xff], [0x0 ~> 0x7f]]
 */ //
//
func fiat_25519_from_bytes(out1 *uint32, arg1 *uint8) {
	var x1 uint32
	var x2 uint32
	var x3 uint32
	var x4 uint32
	var x5 uint32
	var x6 uint32
	var x7 uint32
	var x8 uint32
	var x9 uint32
	var x10 uint32
	var x11 uint32
	var x12 uint32
	var x13 uint32
	var x14 uint32
	var x15 uint32
	var x16 uint8
	var x17 uint32
	var x18 uint32
	var x19 uint32
	var x20 uint32
	var x21 uint32
	var x22 uint32
	var x23 uint32
	var x24 uint32
	var x25 uint32
	var x26 uint32
	var x27 uint32
	var x28 uint32
	var x29 uint32
	var x30 uint32
	var x31 uint32
	var x32 uint8
	var x33 uint32
	var x34 uint32
	var x35 uint32
	var x36 uint32
	var x37 uint8
	var x38 uint32
	var x39 uint32
	var x40 uint32
	var x41 uint32
	var x42 uint8
	var x43 uint32
	var x44 uint32
	var x45 uint32
	var x46 uint32
	var x47 uint8
	var x48 uint32
	var x49 uint32
	var x50 uint32
	var x51 uint32
	var x52 uint8
	var x53 uint32
	var x54 uint32
	var x55 uint32
	var x56 uint32
	var x57 uint32
	var x58 uint32
	var x59 uint32
	var x60 uint8
	var x61 uint32
	var x62 uint32
	var x63 uint32
	var x64 uint32
	var x65 uint8
	var x66 uint32
	var x67 uint32
	var x68 uint32
	var x69 uint32
	var x70 uint8
	var x71 uint32
	var x72 uint32
	var x73 uint32
	var x74 uint32
	var x75 uint8
	var x76 uint32
	var x77 uint32
	var x78 uint32
	x1 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(31))*unsafe.Sizeof(*arg1))))))) << uint64(int32(18))
	x2 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(30))*unsafe.Sizeof(*arg1))))))) << uint64(int32(10))
	x3 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(29))*unsafe.Sizeof(*arg1))))))) << uint64(int32(2))
	x4 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(28))*unsafe.Sizeof(*arg1))))))) << uint64(int32(20))
	x5 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(27))*unsafe.Sizeof(*arg1))))))) << uint64(int32(12))
	x6 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(26))*unsafe.Sizeof(*arg1))))))) << uint64(int32(4))
	x7 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(25))*unsafe.Sizeof(*arg1))))))) << uint64(int32(21))
	x8 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(24))*unsafe.Sizeof(*arg1))))))) << uint64(int32(13))
	x9 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(23))*unsafe.Sizeof(*arg1))))))) << uint64(int32(5))
	x10 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(22))*unsafe.Sizeof(*arg1))))))) << uint64(int32(23))
	x11 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(21))*unsafe.Sizeof(*arg1))))))) << uint64(int32(15))
	x12 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(20))*unsafe.Sizeof(*arg1))))))) << uint64(int32(7))
	x13 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(19))*unsafe.Sizeof(*arg1))))))) << uint64(int32(24))
	x14 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(18))*unsafe.Sizeof(*arg1))))))) << uint64(int32(16))
	x15 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(17))*unsafe.Sizeof(*arg1))))))) << uint64(int32(8))
	x16 = *((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(16))*unsafe.Sizeof(*arg1))))
	x17 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(15))*unsafe.Sizeof(*arg1))))))) << uint64(int32(18))
	x18 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(14))*unsafe.Sizeof(*arg1))))))) << uint64(int32(10))
	x19 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(13))*unsafe.Sizeof(*arg1))))))) << uint64(int32(2))
	x20 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(12))*unsafe.Sizeof(*arg1))))))) << uint64(int32(19))
	x21 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(11))*unsafe.Sizeof(*arg1))))))) << uint64(int32(11))
	x22 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(10))*unsafe.Sizeof(*arg1))))))) << uint64(int32(3))
	x23 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))) << uint64(int32(21))
	x24 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))) << uint64(int32(13))
	x25 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))) << uint64(int32(5))
	x26 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))) << uint64(int32(22))
	x27 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))) << uint64(int32(14))
	x28 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))) << uint64(int32(6))
	x29 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))) << uint64(int32(24))
	x30 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))) << uint64(int32(16))
	x31 = uint32(uint8((*((*uint8)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))) << uint64(int32(8))
	x32 = *arg1
	x33 = x31 + uint32(uint8(x32))
	x34 = x30 + x33
	x35 = x29 + x34
	x36 = x35 & uint32((uint32((uint32(67108863)))))
	x37 = uint8((x35 >> uint64(int32(26))))
	x38 = x28 + uint32(uint8(x37))
	x39 = x27 + x38
	x40 = x26 + x39
	x41 = x40 & uint32((uint32((uint32(33554431)))))
	x42 = uint8((x40 >> uint64(int32(25))))
	x43 = x25 + uint32(uint8(x42))
	x44 = x24 + x43
	x45 = x23 + x44
	x46 = x45 & uint32((uint32((uint32(67108863)))))
	x47 = uint8((x45 >> uint64(int32(26))))
	x48 = x22 + uint32(uint8(x47))
	x49 = x21 + x48
	x50 = x20 + x49
	x51 = x50 & uint32((uint32((uint32(33554431)))))
	x52 = uint8((x50 >> uint64(int32(25))))
	x53 = x19 + uint32(uint8(x52))
	x54 = x18 + x53
	x55 = x17 + x54
	x56 = x15 + uint32(uint8(x16))
	x57 = x14 + x56
	x58 = x13 + x57
	x59 = x58 & uint32((uint32((uint32(33554431)))))
	x60 = uint8((x58 >> uint64(int32(25))))
	x61 = x12 + uint32(uint8(x60))
	x62 = x11 + x61
	x63 = x10 + x62
	x64 = x63 & uint32((uint32((uint32(67108863)))))
	x65 = uint8((x63 >> uint64(int32(26))))
	x66 = x9 + uint32(uint8(x65))
	x67 = x8 + x66
	x68 = x7 + x67
	x69 = x68 & uint32((uint32((uint32(33554431)))))
	x70 = uint8((x68 >> uint64(int32(25))))
	x71 = x6 + uint32(uint8(x70))
	x72 = x5 + x71
	x73 = x4 + x72
	x74 = x73 & uint32((uint32((uint32(67108863)))))
	x75 = uint8((x73 >> uint64(int32(26))))
	x76 = x3 + uint32(uint8(x75))
	x77 = x2 + x76
	x78 = x1 + x77
	*out1 = x36
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(1))*unsafe.Sizeof(*out1)))) = x41
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(2))*unsafe.Sizeof(*out1)))) = x46
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(3))*unsafe.Sizeof(*out1)))) = x51
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(4))*unsafe.Sizeof(*out1)))) = x55
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(5))*unsafe.Sizeof(*out1)))) = x59
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(6))*unsafe.Sizeof(*out1)))) = x64
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(7))*unsafe.Sizeof(*out1)))) = x69
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(8))*unsafe.Sizeof(*out1)))) = x74
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(9))*unsafe.Sizeof(*out1)))) = x78
}

// fiat_25519_relax - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:1428
/*
 * The function fiat_25519_relax is the identity function converting from tight field elements to loose field elements.
 *
 * Postconditions:
 *   out1 = arg1
 *
 */ //
//
func fiat_25519_relax(out1 *uint32, arg1 *uint32) {
	var x1 uint32
	var x2 uint32
	var x3 uint32
	var x4 uint32
	var x5 uint32
	var x6 uint32
	var x7 uint32
	var x8 uint32
	var x9 uint32
	var x10 uint32
	x1 = *arg1
	x2 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))
	x3 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))
	x4 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))
	x5 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))
	x6 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))
	x7 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))
	x8 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))
	x9 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))
	x10 = *((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))
	*out1 = x1
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(1))*unsafe.Sizeof(*out1)))) = x2
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(2))*unsafe.Sizeof(*out1)))) = x3
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(3))*unsafe.Sizeof(*out1)))) = x4
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(4))*unsafe.Sizeof(*out1)))) = x5
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(5))*unsafe.Sizeof(*out1)))) = x6
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(6))*unsafe.Sizeof(*out1)))) = x7
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(7))*unsafe.Sizeof(*out1)))) = x8
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(8))*unsafe.Sizeof(*out1)))) = x9
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(9))*unsafe.Sizeof(*out1)))) = x10
}

// fiat_25519_carry_scmul_121666 - transpiled function from  /mnt/d/c/spake2-c/curve25519_32.h:1468
/*
 * The function fiat_25519_carry_scmul_121666 multiplies a field element by 121666 and reduces the result.
 *
 * Postconditions:
 *   eval out1 mod m = (121666 * eval arg1) mod m
 *
 */ //
//
func fiat_25519_carry_scmul_121666(out1 *uint32, arg1 *uint32) {
	var x1 uint64
	var x2 uint64
	var x3 uint64
	var x4 uint64
	var x5 uint64
	var x6 uint64
	var x7 uint64
	var x8 uint64
	var x9 uint64
	var x10 uint64
	var x11 uint32
	var x12 uint32
	var x13 uint64
	var x14 uint32
	var x15 uint32
	var x16 uint64
	var x17 uint32
	var x18 uint32
	var x19 uint64
	var x20 uint32
	var x21 uint32
	var x22 uint64
	var x23 uint32
	var x24 uint32
	var x25 uint64
	var x26 uint32
	var x27 uint32
	var x28 uint64
	var x29 uint32
	var x30 uint32
	var x31 uint64
	var x32 uint32
	var x33 uint32
	var x34 uint64
	var x35 uint32
	var x36 uint32
	var x37 uint64
	var x38 uint32
	var x39 uint32
	var x40 uint32
	var x41 uint32
	var x42 fiat_25519_uint1
	var x43 uint32
	var x44 uint32
	var x45 fiat_25519_uint1
	var x46 uint32
	var x47 uint32
	x1 = uint64(121666) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(9))*unsafe.Sizeof(*arg1))))))))))))
	x2 = uint64(121666) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(8))*unsafe.Sizeof(*arg1))))))))))))
	x3 = uint64(121666) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(7))*unsafe.Sizeof(*arg1))))))))))))
	x4 = uint64(121666) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(6))*unsafe.Sizeof(*arg1))))))))))))
	x5 = uint64(121666) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(5))*unsafe.Sizeof(*arg1))))))))))))
	x6 = uint64(121666) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(4))*unsafe.Sizeof(*arg1))))))))))))
	x7 = uint64(121666) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(3))*unsafe.Sizeof(*arg1))))))))))))
	x8 = uint64(121666) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(2))*unsafe.Sizeof(*arg1))))))))))))
	x9 = uint64(121666) * uint64((uint64((uint32((uint32((*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(arg1)) + (uintptr)(int32(1))*unsafe.Sizeof(*arg1))))))))))))
	x10 = uint64(121666) * uint64((uint64((uint32((uint32((*arg1))))))))
	x11 = uint32((x10 >> uint64(int32(26))))
	x12 = uint32((uint32((uint64((x10 & uint64((uint64((uint32(67108863)))))))))))
	x13 = uint64((uint64((uint32((uint32((x11)))) + uint32((uint64((x9))))))))
	x14 = uint32((x13 >> uint64(int32(25))))
	x15 = uint32((uint32((uint64((x13 & uint64((uint64((uint32(33554431)))))))))))
	x16 = uint64((uint64((uint32((uint32((x14)))) + uint32((uint64((x8))))))))
	x17 = uint32((x16 >> uint64(int32(26))))
	x18 = uint32((uint32((uint64((x16 & uint64((uint64((uint32(67108863)))))))))))
	x19 = uint64((uint64((uint32((uint32((x17)))) + uint32((uint64((x7))))))))
	x20 = uint32((x19 >> uint64(int32(25))))
	x21 = uint32((uint32((uint64((x19 & uint64((uint64((uint32(33554431)))))))))))
	x22 = uint64((uint64((uint32((uint32((x20)))) + uint32((uint64((x6))))))))
	x23 = uint32((x22 >> uint64(int32(26))))
	x24 = uint32((uint32((uint64((x22 & uint64((uint64((uint32(67108863)))))))))))
	x25 = uint64((uint64((uint32((uint32((x23)))) + uint32((uint64((x5))))))))
	x26 = uint32((x25 >> uint64(int32(25))))
	x27 = uint32((uint32((uint64((x25 & uint64((uint64((uint32(33554431)))))))))))
	x28 = uint64((uint64((uint32((uint32((x26)))) + uint32((uint64((x4))))))))
	x29 = uint32((x28 >> uint64(int32(26))))
	x30 = uint32((uint32((uint64((x28 & uint64((uint64((uint32(67108863)))))))))))
	x31 = uint64((uint64((uint32((uint32((x29)))) + uint32((uint64((x3))))))))
	x32 = uint32((x31 >> uint64(int32(25))))
	x33 = uint32((uint32((uint64((x31 & uint64((uint64((uint32(33554431)))))))))))
	x34 = uint64((uint64((uint32((uint32((x32)))) + uint32((uint64((x2))))))))
	x35 = uint32((x34 >> uint64(int32(26))))
	x36 = uint32((uint32((uint64((x34 & uint64((uint64((uint32(67108863)))))))))))
	x37 = uint64((uint64((uint32((uint32((x35)))) + uint32((uint64((x1))))))))
	x38 = uint32((x37 >> uint64(int32(25))))
	x39 = uint32((uint32((uint64((x37 & uint64((uint64((uint32(33554431)))))))))))
	x40 = x38 * uint32((uint32((uint32(int32(19))))))
	x41 = x12 + x40
	x42 = fiat_25519_uint1((x41 >> uint64(int32(26))))
	x43 = x41 & uint32((uint32((uint32(67108863)))))
	x44 = uint32((uint32((uint32(uint8((x42))) + uint32((uint32((x15))))))))
	x45 = fiat_25519_uint1((x44 >> uint64(int32(25))))
	x46 = x44 & uint32((uint32((uint32(33554431)))))
	x47 = uint32((uint32((uint32(uint8((x45))) + uint32((uint32((x18))))))))
	*out1 = x43
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(1))*unsafe.Sizeof(*out1)))) = x46
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(2))*unsafe.Sizeof(*out1)))) = x47
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(3))*unsafe.Sizeof(*out1)))) = x21
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(4))*unsafe.Sizeof(*out1)))) = x24
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(5))*unsafe.Sizeof(*out1)))) = x27
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(6))*unsafe.Sizeof(*out1)))) = x30
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(7))*unsafe.Sizeof(*out1)))) = x33
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(8))*unsafe.Sizeof(*out1)))) = x36
	*((*uint32)(unsafe.Pointer(uintptr(unsafe.Pointer(out1)) + (uintptr)(int32(9))*unsafe.Sizeof(*out1)))) = x39
}
