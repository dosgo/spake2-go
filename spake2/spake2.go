package spake2

import (
	"crypto/sha512"
	"encoding/binary"
	"errors"
	"hash"

	"github.com/dosgo/spake2-go/ed25519"

	"unsafe"

	"crypto/rand"
)

// Low-level intrinsic operations

func load3(in []byte) uint64 {
	return uint64(in[0]) | uint64(in[1])<<8 | uint64(in[2])<<16
}

func load4(in []byte) uint64 {
	return uint64(in[0]) | uint64(in[1])<<8 | uint64(in[2])<<16 | uint64(in[3])<<24
}

// int64Lshift21 returns |a << 21| but is defined when shifting bits into the
// sign bit. This works around a language flaw in C.
func int64Lshift21(a int64) int64 {
	return int64(uint64(a) << 21)
}

func equal(b, c int8) uint8 {
	ub := uint8(b)
	uc := uint8(c)
	x := ub ^ uc   // 0: yes; 1..255: no
	y := uint32(x) // 0: yes; 1..255: no
	y -= 1         // 4294967295: yes; 0..254: no
	y >>= 31       // 1: yes; 0: no
	return uint8(y)
}

func cmov(t, u *ed25519.Ge_precomp, b uint8) {
	ed25519.Fe_cmov(&t.Yplusx, &u.Yplusx, uint32(b))
	ed25519.Fe_cmov(&t.Yminusx, &u.Yminusx, uint32(b))
	ed25519.Fe_cmov(&t.Xy2d, &u.Xy2d, uint32(b))
}

func cmovCached(t, u *ed25519.Ge_cached, b uint8) {
	ed25519.Fe_cmov(&t.YplusX, &u.YplusX, uint32(b))
	ed25519.Fe_cmov(&t.YminusX, &u.YminusX, uint32(b))
	ed25519.Fe_cmov(&t.Z, &u.Z, uint32(b))
	ed25519.Fe_cmov(&t.T2d, &u.T2d, uint32(b))
}

// This block of code replaces the standard base-point table with a much smaller
// one. The standard table is 30,720 bytes while this one is just 960.
//
// This table contains 15 pairs of group elements, (x, y), where each field
// element is serialized with |feToBytes|. If |i| is the index of the group
// element then consider i+1 as a four-bit number: (i₀, i₁, i₂, i₃) (where i₀
// is the most significant bit). The value of the group element is then:
// (i₀×2^192 + i₁×2^128 + i₂×2^64 + i₃)G, where G is the generator.
var k25519SmallPrecomp = []byte{
	0x1a, 0xd5, 0x25, 0x8f, 0x60, 0x2d, 0x56, 0xc9, 0xb2, 0xa7, 0x25, 0x95,
	0x60, 0xc7, 0x2c, 0x69, 0x5c, 0xdc, 0xd6, 0xfd, 0x31, 0xe2, 0xa4, 0xc0,
	0xfe, 0x53, 0x6e, 0xcd, 0xd3, 0x36, 0x69, 0x21, 0x58, 0x66, 0x66, 0x66,
	0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
	0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66, 0x66,
	0x66, 0x66, 0x66, 0x66, 0x02, 0xa2, 0xed, 0xf4, 0x8f, 0x6b, 0x0b, 0x3e,
	0xeb, 0x35, 0x1a, 0xd5, 0x7e, 0xdb, 0x78, 0x00, 0x96, 0x8a, 0xa0, 0xb4,
	0xcf, 0x60, 0x4b, 0xd4, 0xd5, 0xf9, 0x2d, 0xbf, 0x88, 0xbd, 0x22, 0x62,
	0x13, 0x53, 0xe4, 0x82, 0x57, 0xfa, 0x1e, 0x8f, 0x06, 0x2b, 0x90, 0xba,
	0x08, 0xb6, 0x10, 0x54, 0x4f, 0x7c, 0x1b, 0x26, 0xed, 0xda, 0x6b, 0xdd,
	0x25, 0xd0, 0x4e, 0xea, 0x42, 0xbb, 0x25, 0x03, 0xa2, 0xfb, 0xcc, 0x61,
	0x67, 0x06, 0x70, 0x1a, 0xc4, 0x78, 0x3a, 0xff, 0x32, 0x62, 0xdd, 0x2c,
	0xab, 0x50, 0x19, 0x3b, 0xf2, 0x9b, 0x7d, 0xb8, 0xfd, 0x4f, 0x29, 0x9c,
	0xa7, 0x91, 0xba, 0x0e, 0x46, 0x5e, 0x51, 0xfe, 0x1d, 0xbf, 0xe5, 0xe5,
	0x9b, 0x95, 0x0d, 0x67, 0xf8, 0xd1, 0xb5, 0x5a, 0xa1, 0x93, 0x2c, 0xc3,
	0xde, 0x0e, 0x97, 0x85, 0x2d, 0x7f, 0xea, 0xab, 0x3e, 0x47, 0x30, 0x18,
	0x24, 0xe8, 0xb7, 0x60, 0xae, 0x47, 0x80, 0xfc, 0xe5, 0x23, 0xe7, 0xc2,
	0xc9, 0x85, 0xe6, 0x98, 0xa0, 0x29, 0x4e, 0xe1, 0x84, 0x39, 0x2d, 0x95,
	0x2c, 0xf3, 0x45, 0x3c, 0xff, 0xaf, 0x27, 0x4c, 0x6b, 0xa6, 0xf5, 0x4b,
	0x11, 0xbd, 0xba, 0x5b, 0x9e, 0xc4, 0xa4, 0x51, 0x1e, 0xbe, 0xd0, 0x90,
	0x3a, 0x9c, 0xc2, 0x26, 0xb6, 0x1e, 0xf1, 0x95, 0x7d, 0xc8, 0x6d, 0x52,
	0xe6, 0x99, 0x2c, 0x5f, 0x9a, 0x96, 0x0c, 0x68, 0x29, 0xfd, 0xe2, 0xfb,
	0xe6, 0xbc, 0xec, 0x31, 0x08, 0xec, 0xe6, 0xb0, 0x53, 0x60, 0xc3, 0x8c,
	0xbe, 0xc1, 0xb3, 0x8a, 0x8f, 0xe4, 0x88, 0x2b, 0x55, 0xe5, 0x64, 0x6e,
	0x9b, 0xd0, 0xaf, 0x7b, 0x64, 0x2a, 0x35, 0x25, 0x10, 0x52, 0xc5, 0x9e,
	0x58, 0x11, 0x39, 0x36, 0x45, 0x51, 0xb8, 0x39, 0x93, 0xfc, 0x9d, 0x6a,
	0xbe, 0x58, 0xcb, 0xa4, 0x0f, 0x51, 0x3c, 0x38, 0x05, 0xca, 0xab, 0x43,
	0x63, 0x0e, 0xf3, 0x8b, 0x41, 0xa6, 0xf8, 0x9b, 0x53, 0x70, 0x80, 0x53,
	0x86, 0x5e, 0x8f, 0xe3, 0xc3, 0x0d, 0x18, 0xc8, 0x4b, 0x34, 0x1f, 0xd8,
	0x1d, 0xbc, 0xf2, 0x6d, 0x34, 0x3a, 0xbe, 0xdf, 0xd9, 0xf6, 0xf3, 0x89,
	0xa1, 0xe1, 0x94, 0x9f, 0x5d, 0x4c, 0x5d, 0xe9, 0xa1, 0x49, 0x92, 0xef,
	0x0e, 0x53, 0x81, 0x89, 0x58, 0x87, 0xa6, 0x37, 0xf1, 0xdd, 0x62, 0x60,
	0x63, 0x5a, 0x9d, 0x1b, 0x8c, 0xc6, 0x7d, 0x52, 0xea, 0x70, 0x09, 0x6a,
	0xe1, 0x32, 0xf3, 0x73, 0x21, 0x1f, 0x07, 0x7b, 0x7c, 0x9b, 0x49, 0xd8,
	0xc0, 0xf3, 0x25, 0x72, 0x6f, 0x9d, 0xed, 0x31, 0x67, 0x36, 0x36, 0x54,
	0x40, 0x92, 0x71, 0xe6, 0x11, 0x28, 0x11, 0xad, 0x93, 0x32, 0x85, 0x7b,
	0x3e, 0xb7, 0x3b, 0x49, 0x13, 0x1c, 0x07, 0xb0, 0x2e, 0x93, 0xaa, 0xfd,
	0xfd, 0x28, 0x47, 0x3d, 0x8d, 0xd2, 0xda, 0xc7, 0x44, 0xd6, 0x7a, 0xdb,
	0x26, 0x7d, 0x1d, 0xb8, 0xe1, 0xde, 0x9d, 0x7a, 0x7d, 0x17, 0x7e, 0x1c,
	0x37, 0x04, 0x8d, 0x2d, 0x7c, 0x5e, 0x18, 0x38, 0x1e, 0xaf, 0xc7, 0x1b,
	0x33, 0x48, 0x31, 0x00, 0x59, 0xf6, 0xf2, 0xca, 0x0f, 0x27, 0x1b, 0x63,
	0x12, 0x7e, 0x02, 0x1d, 0x49, 0xc0, 0x5d, 0x79, 0x87, 0xef, 0x5e, 0x7a,
	0x2f, 0x1f, 0x66, 0x55, 0xd8, 0x09, 0xd9, 0x61, 0x38, 0x68, 0xb0, 0x07,
	0xa3, 0xfc, 0xcc, 0x85, 0x10, 0x7f, 0x4c, 0x65, 0x65, 0xb3, 0xfa, 0xfa,
	0xa5, 0x53, 0x6f, 0xdb, 0x74, 0x4c, 0x56, 0x46, 0x03, 0xe2, 0xd5, 0x7a,
	0x29, 0x1c, 0xc6, 0x02, 0xbc, 0x59, 0xf2, 0x04, 0x75, 0x63, 0xc0, 0x84,
	0x2f, 0x60, 0x1c, 0x67, 0x76, 0xfd, 0x63, 0x86, 0xf3, 0xfa, 0xbf, 0xdc,
	0xd2, 0x2d, 0x90, 0x91, 0xbd, 0x33, 0xa9, 0xe5, 0x66, 0x0c, 0xda, 0x42,
	0x27, 0xca, 0xf4, 0x66, 0xc2, 0xec, 0x92, 0x14, 0x57, 0x06, 0x63, 0xd0,
	0x4d, 0x15, 0x06, 0xeb, 0x69, 0x58, 0x4f, 0x77, 0xc5, 0x8b, 0xc7, 0xf0,
	0x8e, 0xed, 0x64, 0xa0, 0xb3, 0x3c, 0x66, 0x71, 0xc6, 0x2d, 0xda, 0x0a,
	0x0d, 0xfe, 0x70, 0x27, 0x64, 0xf8, 0x27, 0xfa, 0xf6, 0x5f, 0x30, 0xa5,
	0x0d, 0x6c, 0xda, 0xf2, 0x62, 0x5e, 0x78, 0x47, 0xd3, 0x66, 0x00, 0x1c,
	0xfd, 0x56, 0x1f, 0x5d, 0x3f, 0x6f, 0xf4, 0x4c, 0xd8, 0xfd, 0x0e, 0x27,
	0xc9, 0x5c, 0x2b, 0xbc, 0xc0, 0xa4, 0xe7, 0x23, 0x29, 0x02, 0x9f, 0x31,
	0xd6, 0xe9, 0xd7, 0x96, 0xf4, 0xe0, 0x5e, 0x0b, 0x0e, 0x13, 0xee, 0x3c,
	0x09, 0xed, 0xf2, 0x3d, 0x76, 0x91, 0xc3, 0xa4, 0x97, 0xae, 0xd4, 0x87,
	0xd0, 0x5d, 0xf6, 0x18, 0x47, 0x1f, 0x1d, 0x67, 0xf2, 0xcf, 0x63, 0xa0,
	0x91, 0x27, 0xf8, 0x93, 0x45, 0x75, 0x23, 0x3f, 0xd1, 0xf1, 0xad, 0x23,
	0xdd, 0x64, 0x93, 0x96, 0x41, 0x70, 0x7f, 0xf7, 0xf5, 0xa9, 0x89, 0xa2,
	0x34, 0xb0, 0x8d, 0x1b, 0xae, 0x19, 0x15, 0x49, 0x58, 0x23, 0x6d, 0x87,
	0x15, 0x4f, 0x81, 0x76, 0xfb, 0x23, 0xb5, 0xea, 0xcf, 0xac, 0x54, 0x8d,
	0x4e, 0x42, 0x2f, 0xeb, 0x0f, 0x63, 0xdb, 0x68, 0x37, 0xa8, 0xcf, 0x8b,
	0xab, 0xf5, 0xa4, 0x6e, 0x96, 0x2a, 0xb2, 0xd6, 0xbe, 0x9e, 0xbd, 0x0d,
	0xb4, 0x42, 0xa9, 0xcf, 0x01, 0x83, 0x8a, 0x17, 0x47, 0x76, 0xc4, 0xc6,
	0x83, 0x04, 0x95, 0x0b, 0xfc, 0x11, 0xc9, 0x62, 0xb8, 0x0c, 0x76, 0x84,
	0xd9, 0xb9, 0x37, 0xfa, 0xfc, 0x7c, 0xc2, 0x6d, 0x58, 0x3e, 0xb3, 0x04,
	0xbb, 0x8c, 0x8f, 0x48, 0xbc, 0x91, 0x27, 0xcc, 0xf9, 0xb7, 0x22, 0x19,
	0x83, 0x2e, 0x09, 0xb5, 0x72, 0xd9, 0x54, 0x1c, 0x4d, 0xa1, 0xea, 0x0b,
	0xf1, 0xc6, 0x08, 0x72, 0x46, 0x87, 0x7a, 0x6e, 0x80, 0x56, 0x0a, 0x8a,
	0xc0, 0xdd, 0x11, 0x6b, 0xd6, 0xdd, 0x47, 0xdf, 0x10, 0xd9, 0xd8, 0xea,
	0x7c, 0xb0, 0x8f, 0x03, 0x00, 0x2e, 0xc1, 0x8f, 0x44, 0xa8, 0xd3, 0x30,
	0x06, 0x89, 0xa2, 0xf9, 0x34, 0xad, 0xdc, 0x03, 0x85, 0xed, 0x51, 0xa7,
	0x82, 0x9c, 0xe7, 0x5d, 0x52, 0x93, 0x0c, 0x32, 0x9a, 0x5b, 0xe1, 0xaa,
	0xca, 0xb8, 0x02, 0x6d, 0x3a, 0xd4, 0xb1, 0x3a, 0xf0, 0x5f, 0xbe, 0xb5,
	0x0d, 0x10, 0x6b, 0x38, 0x32, 0xac, 0x76, 0x80, 0xbd, 0xca, 0x94, 0x71,
	0x7a, 0xf2, 0xc9, 0x35, 0x2a, 0xde, 0x9f, 0x42, 0x49, 0x18, 0x01, 0xab,
	0xbc, 0xef, 0x7c, 0x64, 0x3f, 0x58, 0x3d, 0x92, 0x59, 0xdb, 0x13, 0xdb,
	0x58, 0x6e, 0x0a, 0xe0, 0xb7, 0x91, 0x4a, 0x08, 0x20, 0xd6, 0x2e, 0x3c,
	0x45, 0xc9, 0x8b, 0x17, 0x79, 0xe7, 0xc7, 0x90, 0x99, 0x3a, 0x18, 0x25,
}

func x25519GeScalarmultSmallPrecomp(
	h *ed25519.Ge_p3, a []byte, precompTable []byte) {
	// precomp_table is first expanded into matching |ge_precomp|
	// elements.
	multiples := make([]ed25519.Ge_precomp, 15)

	var i uint32 = 0
	for i = 0; i < 15; i++ {
		// The precomputed table is assumed to already clear the top bit, so
		// |fe_frombytes_strict| may be used directly.
		//const uint8_t *bytes = &precomp_table[i*(2*32)]
		var x, y ed25519.Fe
		ed25519.Fe_frombytes_strict(&x, precompTable[i*(2*32):])
		ed25519.Fe_frombytes_strict(&y, precompTable[i*(2*32)+32:])

		ed25519.Fe_add(&multiples[i].Yplusx, &y, &x)
		ed25519.Fe_sub(&multiples[i].Yminusx, &y, &x)

		ed25519.Fe_mul_ltt(&multiples[i].Xy2d, &x, &y)

		ed25519.Fe_mul_llt(&multiples[i].Xy2d, &multiples[i].Xy2d, &ed25519.D2)

	}

	// See the comment above |k25519SmallPrecomp| about the structure of the
	// precomputed elements. This loop does 64 additions and 64 doublings to
	// calculate the result.
	ed25519.Ge_p3_0(h)
	for i = 63; i < 64; i-- {
		var j uint32
		var index int8 = 0
		var bit uint8
		for j = 0; j < 4; j++ {
			bit = 1 & (a[(8*j)+(i/8)] >> (i & 7))
			//fmt.Printf("bit:%d a:%d\r\n", bit, a[(8*j)+(i/8)])
			index |= int8(bit << j)
		}
		//	fmt.Printf("index:%d\r\n", index)
		var e ed25519.Ge_precomp
		ed25519.Ge_precomp_0(&e)

		for j = 1; j < 16; j++ {
			cmov(&e, &multiples[j-1], equal(index, int8(j)))
		}

		var cached ed25519.Ge_cached
		var r ed25519.Ge_p1p1

		ed25519.X25519_ge_p3_to_cached(&cached, h)

		ed25519.X25519_ge_add(&r, h, &cached)

		ed25519.X25519_ge_p1p1_to_p3(h, &r)

		ed25519.Ge_madd(&r, h, &e)

		ed25519.X25519_ge_p1p1_to_p3(h, &r)

	}
}

func x25519GeScalarmultBase(h *ed25519.Ge_p3, a []byte) {
	x25519GeScalarmultSmallPrecomp(h, a, k25519SmallPrecomp)
}

func x25519GeScalarmult(r *ed25519.Ge_p2, scalar []byte, A *ed25519.Ge_p3) {
	var Ai_p2 [8]ed25519.Ge_p2
	var Ai [16]ed25519.Ge_cached
	var t ed25519.Ge_p1p1

	ed25519.Ge_cached_0(&Ai[0])
	ed25519.X25519_ge_p3_to_cached(&Ai[1], A)
	ed25519.Ge_p3_to_p2(&Ai_p2[1], A)

	var i uint32 = 0
	for i = 2; i < 16; i += 2 {
		ed25519.Ge_p2_dbl(&t, &Ai_p2[i/2])
		ed25519.Ge_p1p1_to_cached(&Ai[i], &t)
		if i < 8 {
			ed25519.X25519_ge_p1p1_to_p2(&Ai_p2[i], &t)
		}
		ed25519.X25519_ge_add(&t, A, &Ai[i])
		ed25519.Ge_p1p1_to_cached(&Ai[i+1], &t)
		if i < 7 {
			ed25519.X25519_ge_p1p1_to_p2(&Ai_p2[i+1], &t)
		}
	}

	ed25519.Ge_p2_0(r)
	var u ed25519.Ge_p3

	for i = 0; i < 256; i += 4 {
		ed25519.Ge_p2_dbl(&t, r)
		ed25519.X25519_ge_p1p1_to_p2(r, &t)
		ed25519.Ge_p2_dbl(&t, r)
		ed25519.X25519_ge_p1p1_to_p2(r, &t)
		ed25519.Ge_p2_dbl(&t, r)
		ed25519.X25519_ge_p1p1_to_p2(r, &t)
		ed25519.Ge_p2_dbl(&t, r)
		ed25519.X25519_ge_p1p1_to_p3(&u, &t)

		var index uint8 = scalar[31-i/8]
		index >>= 4 - (i & 4)
		index &= 0xf

		var j uint32
		var selected ed25519.Ge_cached
		ed25519.Ge_cached_0(&selected)
		for j = 0; j < 16; j++ {
			cmovCached(&selected, &Ai[j], equal(int8(j), int8(index)))
		}

		ed25519.X25519_ge_add(&t, &u, &selected)
		ed25519.X25519_ge_p1p1_to_p2(r, &t)
	}
}

// The set of scalars is Z/l
// where l = 2^252 + 27742317777372353535851937790883648493.

// Input:
//
//	s[0]+256*s[1]+...+256^63*s[63] = s
//
// Output:
//
//	s[0]+256*s[1]+...+256^31*s[31] = s mod l
//	where l = 2^252 + 27742317777372353535851937790883648493.
//	Overwrites s in place.
func x25519ScReduce(s []byte) []byte {
	s0 := int64(2097151 & load3(s))
	s1 := int64(2097151 & (load4(s[2:]) >> 5))
	s2 := int64(2097151 & (load3(s[5:]) >> 2))
	s3 := int64(2097151 & (load4(s[7:]) >> 7))
	s4 := int64(2097151 & (load4(s[10:]) >> 4))
	s5 := int64(2097151 & (load3(s[13:]) >> 1))
	s6 := int64(2097151 & (load4(s[15:]) >> 6))
	s7 := int64(2097151 & (load3(s[18:]) >> 3))
	s8 := int64(2097151 & load3(s[21:]))
	s9 := int64(2097151 & (load4(s[23:]) >> 5))
	s10 := int64(2097151 & (load3(s[26:]) >> 2))
	s11 := int64(2097151 & (load4(s[28:]) >> 7))
	s12 := int64(2097151 & (load4(s[31:]) >> 4))
	s13 := int64(2097151 & (load3(s[34:]) >> 1))
	s14 := int64(2097151 & (load4(s[36:]) >> 6))
	s15 := int64(2097151 & (load3(s[39:]) >> 3))
	s16 := int64(2097151 & load3(s[42:]))
	s17 := int64(2097151 & (load4(s[44:]) >> 5))
	s18 := int64(2097151 & (load3(s[47:]) >> 2))
	s19 := int64(2097151 & (load4(s[49:]) >> 7))
	s20 := int64(2097151 & (load4(s[52:]) >> 4))
	s21 := int64(2097151 & (load3(s[55:]) >> 1))
	s22 := int64(2097151 & (load4(s[57:]) >> 6))
	s23 := int64(load4(s[60:]) >> 3)

	var carry0, carry1, carry2, carry3, carry4, carry5, carry6, carry7, carry8, carry9, carry10, carry11, carry12, carry13, carry14, carry15, carry16 int64

	s11 += s23 * 666643
	s12 += s23 * 470296
	s13 += s23 * 654183
	s14 -= s23 * 997805
	s15 += s23 * 136657
	s16 -= s23 * 683901
	s23 = 0

	s10 += s22 * 666643
	s11 += s22 * 470296
	s12 += s22 * 654183
	s13 -= s22 * 997805
	s14 += s22 * 136657
	s15 -= s22 * 683901
	s22 = 0

	s9 += s21 * 666643
	s10 += s21 * 470296
	s11 += s21 * 654183
	s12 -= s21 * 997805
	s13 += s21 * 136657
	s14 -= s21 * 683901
	s21 = 0

	s8 += s20 * 666643
	s9 += s20 * 470296
	s10 += s20 * 654183
	s11 -= s20 * 997805
	s12 += s20 * 136657
	s13 -= s20 * 683901
	s20 = 0

	s7 += s19 * 666643
	s8 += s19 * 470296
	s9 += s19 * 654183
	s10 -= s19 * 997805
	s11 += s19 * 136657
	s12 -= s19 * 683901
	s19 = 0

	s6 += s18 * 666643
	s7 += s18 * 470296
	s8 += s18 * 654183
	s9 -= s18 * 997805
	s10 += s18 * 136657
	s11 -= s18 * 683901
	s18 = 0

	carry6 = (s6 + (1 << 20)) >> 21
	s7 += carry6
	s6 -= int64Lshift21(carry6)
	carry8 = (s8 + (1 << 20)) >> 21
	s9 += carry8
	s8 -= int64Lshift21(carry8)
	carry10 = (s10 + (1 << 20)) >> 21
	s11 += carry10
	s10 -= int64Lshift21(carry10)
	carry12 = (s12 + (1 << 20)) >> 21
	s13 += carry12
	s12 -= int64Lshift21(carry12)
	carry14 = (s14 + (1 << 20)) >> 21
	s15 += carry14
	s14 -= int64Lshift21(carry14)
	carry16 = (s16 + (1 << 20)) >> 21
	s17 += carry16
	s16 -= int64Lshift21(carry16)

	carry7 = (s7 + (1 << 20)) >> 21
	s8 += carry7
	s7 -= int64Lshift21(carry7)
	carry9 = (s9 + (1 << 20)) >> 21
	s10 += carry9
	s9 -= int64Lshift21(carry9)
	carry11 = (s11 + (1 << 20)) >> 21
	s12 += carry11
	s11 -= int64Lshift21(carry11)
	carry13 = (s13 + (1 << 20)) >> 21
	s14 += carry13
	s13 -= int64Lshift21(carry13)
	carry15 = (s15 + (1 << 20)) >> 21
	s16 += carry15
	s15 -= int64Lshift21(carry15)

	s5 += s17 * 666643
	s6 += s17 * 470296
	s7 += s17 * 654183
	s8 -= s17 * 997805
	s9 += s17 * 136657
	s10 -= s17 * 683901
	s17 = 0

	s4 += s16 * 666643
	s5 += s16 * 470296
	s6 += s16 * 654183
	s7 -= s16 * 997805
	s8 += s16 * 136657
	s9 -= s16 * 683901
	s16 = 0

	s3 += s15 * 666643
	s4 += s15 * 470296
	s5 += s15 * 654183
	s6 -= s15 * 997805
	s7 += s15 * 136657
	s8 -= s15 * 683901
	s15 = 0

	s2 += s14 * 666643
	s3 += s14 * 470296
	s4 += s14 * 654183
	s5 -= s14 * 997805
	s6 += s14 * 136657
	s7 -= s14 * 683901
	s14 = 0

	s1 += s13 * 666643
	s2 += s13 * 470296
	s3 += s13 * 654183
	s4 -= s13 * 997805
	s5 += s13 * 136657
	s6 -= s13 * 683901
	s13 = 0

	s0 += s12 * 666643
	s1 += s12 * 470296
	s2 += s12 * 654183
	s3 -= s12 * 997805
	s4 += s12 * 136657
	s5 -= s12 * 683901
	s12 = 0

	carry0 = (s0 + (1 << 20)) >> 21
	s1 += carry0
	s0 -= int64Lshift21(carry0)
	carry2 = (s2 + (1 << 20)) >> 21
	s3 += carry2
	s2 -= int64Lshift21(carry2)
	carry4 = (s4 + (1 << 20)) >> 21
	s5 += carry4
	s4 -= int64Lshift21(carry4)
	carry6 = (s6 + (1 << 20)) >> 21
	s7 += carry6
	s6 -= int64Lshift21(carry6)
	carry8 = (s8 + (1 << 20)) >> 21
	s9 += carry8
	s8 -= int64Lshift21(carry8)
	carry10 = (s10 + (1 << 20)) >> 21
	s11 += carry10
	s10 -= int64Lshift21(carry10)

	carry1 = (s1 + (1 << 20)) >> 21
	s2 += carry1
	s1 -= int64Lshift21(carry1)
	carry3 = (s3 + (1 << 20)) >> 21
	s4 += carry3
	s3 -= int64Lshift21(carry3)
	carry5 = (s5 + (1 << 20)) >> 21
	s6 += carry5
	s5 -= int64Lshift21(carry5)
	carry7 = (s7 + (1 << 20)) >> 21
	s8 += carry7
	s7 -= int64Lshift21(carry7)
	carry9 = (s9 + (1 << 20)) >> 21
	s10 += carry9
	s9 -= int64Lshift21(carry9)
	carry11 = (s11 + (1 << 20)) >> 21
	s12 += carry11
	s11 -= int64Lshift21(carry11)

	s0 += s12 * 666643
	s1 += s12 * 470296
	s2 += s12 * 654183
	s3 -= s12 * 997805
	s4 += s12 * 136657
	s5 -= s12 * 683901
	s12 = 0

	carry0 = s0 >> 21
	s1 += carry0
	s0 -= int64Lshift21(carry0)
	carry1 = s1 >> 21
	s2 += carry1
	s1 -= int64Lshift21(carry1)
	carry2 = s2 >> 21
	s3 += carry2
	s2 -= int64Lshift21(carry2)
	carry3 = s3 >> 21
	s4 += carry3
	s3 -= int64Lshift21(carry3)
	carry4 = s4 >> 21
	s5 += carry4
	s4 -= int64Lshift21(carry4)
	carry5 = s5 >> 21
	s6 += carry5
	s5 -= int64Lshift21(carry5)
	carry6 = s6 >> 21
	s7 += carry6
	s6 -= int64Lshift21(carry6)
	carry7 = s7 >> 21
	s8 += carry7
	s7 -= int64Lshift21(carry7)
	carry8 = s8 >> 21
	s9 += carry8
	s8 -= int64Lshift21(carry8)
	carry9 = s9 >> 21
	s10 += carry9
	s9 -= int64Lshift21(carry9)
	carry10 = s10 >> 21
	s11 += carry10
	s10 -= int64Lshift21(carry10)
	carry11 = s11 >> 21
	s12 += carry11
	s11 -= int64Lshift21(carry11)

	s0 += s12 * 666643
	s1 += s12 * 470296
	s2 += s12 * 654183
	s3 -= s12 * 997805
	s4 += s12 * 136657
	s5 -= s12 * 683901
	s12 = 0

	carry0 = s0 >> 21
	s1 += carry0
	s0 -= int64Lshift21(carry0)
	carry1 = s1 >> 21
	s2 += carry1
	s1 -= int64Lshift21(carry1)
	carry2 = s2 >> 21
	s3 += carry2
	s2 -= int64Lshift21(carry2)
	carry3 = s3 >> 21
	s4 += carry3
	s3 -= int64Lshift21(carry3)
	carry4 = s4 >> 21
	s5 += carry4
	s4 -= int64Lshift21(carry4)
	carry5 = s5 >> 21
	s6 += carry5
	s5 -= int64Lshift21(carry5)
	carry6 = s6 >> 21
	s7 += carry6
	s6 -= int64Lshift21(carry6)
	carry7 = s7 >> 21
	s8 += carry7
	s7 -= int64Lshift21(carry7)
	carry8 = s8 >> 21
	s9 += carry8
	s8 -= int64Lshift21(carry8)
	carry9 = s9 >> 21
	s10 += carry9
	s9 -= int64Lshift21(carry9)
	carry10 = s10 >> 21
	s11 += carry10
	s10 -= int64Lshift21(carry10)

	s[0] = byte(s0 >> 0)
	s[1] = byte(s0 >> 8)
	s[2] = byte((s0 >> 16) | (s1 << 5))
	s[3] = byte(s1 >> 3)
	s[4] = byte(s1 >> 11)
	s[5] = byte((s1 >> 19) | (s2 << 2))
	s[6] = byte(s2 >> 6)
	s[7] = byte((s2 >> 14) | (s3 << 7))
	s[8] = byte(s3 >> 1)
	s[9] = byte(s3 >> 9)
	s[10] = byte((s3 >> 17) | (s4 << 4))
	s[11] = byte(s4 >> 4)
	s[12] = byte(s4 >> 12)
	s[13] = byte((s4 >> 20) | (s5 << 1))
	s[14] = byte(s5 >> 7)
	s[15] = byte((s5 >> 15) | (s6 << 6))
	s[16] = byte(s6 >> 2)
	s[17] = byte(s6 >> 10)
	s[18] = byte((s6 >> 18) | (s7 << 3))
	s[19] = byte(s7 >> 5)
	s[20] = byte(s7 >> 13)
	s[21] = byte(s8 >> 0)
	s[22] = byte(s8 >> 8)
	s[23] = byte((s8 >> 16) | (s9 << 5))
	s[24] = byte(s9 >> 3)
	s[25] = byte(s9 >> 11)
	s[26] = byte((s9 >> 19) | (s10 << 2))
	s[27] = byte(s10 >> 6)
	s[28] = byte((s10 >> 14) | (s11 << 7))
	s[29] = byte(s11 >> 1)
	s[30] = byte(s11 >> 9)
	s[31] = byte(s11 >> 17)
	return s
}

var kSpakeNSmallPrecomp = []byte{

	0x20, 0x1b, 0xc5, 0xb3, 0x43, 0x17, 0x71, 0x10, 0x44, 0x1e, 0x73, 0xb3,
	0xae, 0x3f, 0xbf, 0x9f, 0xf5, 0x44, 0xc8, 0x13, 0x8f, 0xd1, 0x01, 0xc2,
	0x8a, 0x1a, 0x6d, 0xea, 0x4d, 0x00, 0x5d, 0x6e, 0x10, 0xe3, 0xdf, 0x0a,
	0xe3, 0x7d, 0x8e, 0x7a, 0x99, 0xb5, 0xfe, 0x74, 0xb4, 0x46, 0x72, 0x10,
	0x3d, 0xbd, 0xdc, 0xbd, 0x06, 0xaf, 0x68, 0x0d, 0x71, 0x32, 0x9a, 0x11,
	0x69, 0x3b, 0xc7, 0x78, 0x93, 0xf1, 0x57, 0x97, 0x6e, 0xf0, 0x6e, 0x45,
	0x37, 0x4a, 0xf4, 0x0b, 0x18, 0x51, 0xf5, 0x4f, 0x67, 0x3c, 0xdc, 0xec,
	0x84, 0xed, 0xd0, 0xeb, 0xca, 0xfb, 0xdb, 0xff, 0x7f, 0xeb, 0xa8, 0x23,
	0x68, 0x87, 0x13, 0x64, 0x6a, 0x10, 0xf7, 0x45, 0xe0, 0x0f, 0x32, 0x21,
	0x59, 0x7c, 0x0e, 0x50, 0xad, 0x56, 0xd7, 0x12, 0x69, 0x7b, 0x58, 0xf8,
	0xb9, 0x3b, 0xa5, 0xbb, 0x4d, 0x1b, 0x87, 0x1c, 0x46, 0xa7, 0x17, 0x9d,
	0x6d, 0x84, 0x45, 0xbe, 0x7f, 0x95, 0xd2, 0x34, 0xcd, 0x89, 0x95, 0xc0,
	0xf0, 0xd3, 0xdf, 0x6e, 0x10, 0x4a, 0xe3, 0x7b, 0xce, 0x7f, 0x40, 0x27,
	0xc7, 0x2b, 0xab, 0x66, 0x03, 0x59, 0xb4, 0x7b, 0xc7, 0xc7, 0xf0, 0x39,
	0x9a, 0x33, 0x35, 0xbf, 0xcc, 0x2f, 0xf3, 0x2e, 0x68, 0x9d, 0x53, 0x5c,
	0x88, 0x52, 0xe3, 0x77, 0x90, 0xa1, 0x27, 0x85, 0xc5, 0x74, 0x7f, 0x23,
	0x0e, 0x93, 0x01, 0x3e, 0xe7, 0x2e, 0x2e, 0x95, 0xf3, 0x0d, 0xc2, 0x25,
	0x25, 0x39, 0x39, 0x3d, 0x6e, 0x8e, 0x89, 0xbd, 0xe8, 0xbb, 0x67, 0x5e,
	0x8c, 0x66, 0x8b, 0x63, 0x28, 0x1e, 0x4e, 0x74, 0x85, 0xa8, 0xaf, 0x0f,
	0x12, 0x5d, 0xb6, 0x8a, 0x83, 0x1a, 0x77, 0x76, 0x5e, 0x62, 0x8a, 0xa7,
	0x3c, 0xb8, 0x05, 0x57, 0x2b, 0xaf, 0x36, 0x2e, 0x10, 0x90, 0xb2, 0x39,
	0xb4, 0x3e, 0x75, 0x6d, 0x3a, 0xa8, 0x31, 0x35, 0xc2, 0x1e, 0x8f, 0xc2,
	0x79, 0x89, 0x35, 0x16, 0x26, 0xd1, 0xc7, 0x0b, 0x04, 0x1f, 0x1d, 0xf9,
	0x9c, 0x05, 0xa6, 0x6b, 0xb5, 0x19, 0x5a, 0x24, 0x6d, 0x91, 0xc5, 0x31,
	0xfd, 0xc5, 0xfa, 0xe7, 0xa6, 0xcb, 0x0e, 0x4b, 0x18, 0x0d, 0x94, 0xc7,
	0xee, 0x1d, 0x46, 0x1f, 0x92, 0xb1, 0xb2, 0x4a, 0x2b, 0x43, 0x37, 0xfe,
	0xc2, 0x15, 0x11, 0x89, 0xef, 0x59, 0x73, 0x3c, 0x06, 0x76, 0x78, 0xcb,
	0xa6, 0x0d, 0x79, 0x5f, 0x28, 0x0b, 0x5b, 0x8c, 0x9e, 0xe4, 0xaa, 0x51,
	0x9a, 0x42, 0x6f, 0x11, 0x50, 0x3d, 0x01, 0xd6, 0x21, 0xc0, 0x99, 0x5e,
	0x1a, 0xe8, 0x81, 0x25, 0x80, 0xeb, 0xed, 0x5d, 0x37, 0x47, 0x30, 0x70,
	0xa0, 0x4e, 0x0b, 0x43, 0x17, 0xbe, 0xb6, 0x47, 0xe7, 0x2a, 0x62, 0x9d,
	0x5d, 0xa6, 0xc5, 0x33, 0x62, 0x9d, 0x56, 0x24, 0x9d, 0x1d, 0xb2, 0x13,
	0xbc, 0x17, 0x66, 0x43, 0xd1, 0x68, 0xd5, 0x3b, 0x17, 0x69, 0x17, 0xa6,
	0x06, 0x9e, 0x12, 0xb8, 0x7c, 0xd5, 0xaf, 0x3e, 0x21, 0x1b, 0x31, 0xeb,
	0x0b, 0xa4, 0x98, 0x1c, 0xf2, 0x6a, 0x5e, 0x7c, 0x9b, 0x45, 0x8f, 0xb2,
	0x12, 0x06, 0xd5, 0x8c, 0x1d, 0xb2, 0xa7, 0x57, 0x5f, 0x2f, 0x4f, 0xdb,
	0x52, 0x99, 0x7c, 0x58, 0x01, 0x5f, 0xf2, 0xa5, 0xf6, 0x51, 0x86, 0x21,
	0x2f, 0x5b, 0x8d, 0x6a, 0xae, 0x83, 0x34, 0x6d, 0x58, 0x4b, 0xef, 0xfe,
	0xbf, 0x73, 0x5d, 0xdb, 0xc4, 0x97, 0x2a, 0x85, 0xf3, 0x6c, 0x46, 0x42,
	0xb3, 0x90, 0xc1, 0x57, 0x97, 0x50, 0x35, 0xb1, 0x9d, 0xb7, 0xc7, 0x3c,
	0x85, 0x6d, 0x6c, 0xfd, 0xce, 0xb0, 0xc9, 0xa2, 0x77, 0xee, 0xc3, 0x6b,
	0x0c, 0x37, 0xfa, 0x30, 0x91, 0xd1, 0x2c, 0xb8, 0x5e, 0x7f, 0x81, 0x5f,
	0x87, 0xfd, 0x18, 0x02, 0x5a, 0x30, 0x4e, 0x62, 0xbc, 0x65, 0xc6, 0xce,
	0x1a, 0xcf, 0x2b, 0xaa, 0x56, 0x3e, 0x4d, 0xcf, 0xba, 0x62, 0x5f, 0x9a,
	0xd0, 0x72, 0xff, 0xef, 0x28, 0xbd, 0xbe, 0xd8, 0x57, 0x3d, 0xf5, 0x57,
	0x7d, 0xe9, 0x71, 0x31, 0xec, 0x98, 0x90, 0x94, 0xd9, 0x54, 0xbf, 0x84,
	0x0b, 0xe3, 0x06, 0x47, 0x19, 0x9a, 0x13, 0x1d, 0xef, 0x9d, 0x13, 0xf3,
	0xdb, 0xc3, 0x5c, 0x72, 0x9e, 0xed, 0x24, 0xaa, 0x64, 0xed, 0xe7, 0x0d,
	0xa0, 0x7c, 0x73, 0xba, 0x9b, 0x86, 0xa7, 0x3b, 0x55, 0xab, 0x58, 0x30,
	0xf1, 0x15, 0x81, 0x83, 0x2f, 0xf9, 0x62, 0x84, 0x98, 0x66, 0xf6, 0x55,
	0x21, 0xd8, 0xf2, 0x25, 0x64, 0x71, 0x4b, 0x12, 0x76, 0x59, 0xc5, 0xaa,
	0x93, 0x67, 0xc3, 0x86, 0x25, 0xab, 0x4e, 0x4b, 0xf6, 0xd8, 0x3f, 0x44,
	0x2e, 0x11, 0xe0, 0xbd, 0x6a, 0xf2, 0x5d, 0xf5, 0xf9, 0x53, 0xea, 0xa4,
	0xc8, 0xd9, 0x50, 0x33, 0x81, 0xd9, 0xa8, 0x2d, 0x91, 0x7d, 0x13, 0x2a,
	0x11, 0xcf, 0xde, 0x3f, 0x0a, 0xd2, 0xbc, 0x33, 0xb2, 0x62, 0x53, 0xea,
	0x77, 0x88, 0x43, 0x66, 0x27, 0x43, 0x85, 0xe9, 0x5f, 0x55, 0xf5, 0x2a,
	0x8a, 0xac, 0xdf, 0xff, 0x9b, 0x4c, 0x96, 0x9c, 0xa5, 0x7a, 0xce, 0xd5,
	0x79, 0x18, 0xf1, 0x0b, 0x58, 0x95, 0x7a, 0xe7, 0xd3, 0x74, 0x65, 0x0b,
	0xa4, 0x64, 0x30, 0xe8, 0x5c, 0xfc, 0x55, 0x56, 0xee, 0x14, 0x14, 0xd3,
	0x45, 0x3b, 0xf8, 0xde, 0x05, 0x3e, 0xb9, 0x3c, 0xd7, 0x6a, 0x52, 0x72,
	0x5b, 0x39, 0x09, 0xbe, 0x82, 0x23, 0x10, 0x4a, 0xb7, 0xc3, 0xdc, 0x4c,
	0x5d, 0xc9, 0xf1, 0x14, 0x83, 0xf9, 0x0b, 0x9b, 0xe9, 0x23, 0x84, 0x6a,
	0xc4, 0x08, 0x3d, 0xda, 0x3d, 0x12, 0x95, 0x87, 0x18, 0xa4, 0x7d, 0x3f,
	0x23, 0xde, 0xd4, 0x1e, 0xa8, 0x47, 0xc3, 0x71, 0xdb, 0xf5, 0x03, 0x6c,
	0x57, 0xe7, 0xa4, 0x43, 0x82, 0x33, 0x7b, 0x62, 0x46, 0x7d, 0xf7, 0x10,
	0x69, 0x18, 0x38, 0x27, 0x9a, 0x6f, 0x38, 0xac, 0xfa, 0x92, 0xc5, 0xae,
	0x66, 0xa6, 0x73, 0x95, 0x15, 0x0e, 0x4c, 0x04, 0xb6, 0xfc, 0xf5, 0xc7,
	0x21, 0x3a, 0x99, 0xdb, 0x0e, 0x36, 0xf0, 0x56, 0xbc, 0x75, 0xf9, 0x87,
	0x9b, 0x11, 0x18, 0x92, 0x64, 0x1a, 0xe7, 0xc7, 0xab, 0x5a, 0xc7, 0x26,
	0x7f, 0x13, 0x98, 0x42, 0x52, 0x43, 0xdb, 0xc8, 0x6d, 0x0b, 0xb7, 0x31,
	0x93, 0x24, 0xd6, 0xe8, 0x24, 0x1f, 0x6f, 0x21, 0xa7, 0x8c, 0xeb, 0xdb,
	0x83, 0xb8, 0x89, 0xe3, 0xc1, 0xd7, 0x69, 0x3b, 0x02, 0x6b, 0x54, 0x0f,
	0x84, 0x2f, 0xb5, 0x5c, 0x17, 0x77, 0xbe, 0xe5, 0x61, 0x0d, 0xc5, 0xdf,
	0x3b, 0xcf, 0x3e, 0x93, 0x4f, 0xf5, 0x89, 0xb9, 0x5a, 0xc5, 0x29, 0x31,
	0xc0, 0xc2, 0xff, 0xe5, 0x3f, 0xa6, 0xac, 0x03, 0xca, 0xf5, 0xff, 0xe0,
	0x36, 0xce, 0xf3, 0xe2, 0xb7, 0x9c, 0x02, 0xe9, 0x9e, 0xd2, 0xbc, 0x87,
	0x2f, 0x3d, 0x9a, 0x1d, 0x8f, 0xc5, 0x72, 0xb8, 0xa2, 0x01, 0xd4, 0x68,
	0xb1, 0x84, 0x16, 0x10, 0xf6, 0xf3, 0x52, 0x25, 0xd9, 0xdc, 0x4c, 0xdd,
	0x0f, 0xd6, 0x4a, 0xcf, 0x60, 0x96, 0x7e, 0xcc, 0x42, 0x0f, 0x64, 0x9d,
	0x72, 0x46, 0x04, 0x07, 0xf2, 0x5b, 0xf4, 0x07, 0xd1, 0xf4, 0x59, 0x71,
}
var kSpakeMSmallPrecomp = []byte{

	0xc8, 0xa6, 0x63, 0xc5, 0x97, 0xf1, 0xee, 0x40, 0xab, 0x62, 0x42, 0xee,
	0x25, 0x6f, 0x32, 0x6c, 0x75, 0x2c, 0xa7, 0xd3, 0xbd, 0x32, 0x3b, 0x1e,
	0x11, 0x9c, 0xbd, 0x04, 0xa9, 0x78, 0x6f, 0x45, 0x5a, 0xda, 0x7e, 0x4b,
	0xf6, 0xdd, 0xd9, 0xad, 0xb6, 0x62, 0x6d, 0x32, 0x13, 0x1c, 0x6b, 0x5c,
	0x51, 0xa1, 0xe3, 0x47, 0xa3, 0x47, 0x8f, 0x53, 0xcf, 0xcf, 0x44, 0x1b,
	0x88, 0xee, 0xd1, 0x2e, 0x03, 0x89, 0xaf, 0xc0, 0x61, 0x2d, 0x9e, 0x35,
	0xeb, 0x0e, 0x03, 0xe0, 0xb7, 0xfb, 0xa5, 0xbc, 0x44, 0xbe, 0x0c, 0x89,
	0x0a, 0x0f, 0xd6, 0x59, 0x47, 0x9e, 0xe6, 0x3d, 0x36, 0x9d, 0xff, 0x44,
	0x5e, 0xac, 0xab, 0xe5, 0x3a, 0xd5, 0xb0, 0x35, 0x9f, 0x6d, 0x7f, 0xba,
	0xc0, 0x85, 0x0e, 0xf4, 0x70, 0x3f, 0x13, 0x90, 0x4c, 0x50, 0x1a, 0xee,
	0xc5, 0xeb, 0x69, 0xfe, 0x98, 0x42, 0x87, 0x1d, 0xce, 0x6c, 0x29, 0xaa,
	0x2b, 0x31, 0xc2, 0x38, 0x7b, 0x6b, 0xee, 0x88, 0x0b, 0xba, 0xce, 0xa8,
	0xca, 0x19, 0x60, 0x1b, 0x16, 0xf1, 0x25, 0x1e, 0xcf, 0x63, 0x66, 0x1e,
	0xbb, 0x63, 0xeb, 0x7d, 0xca, 0xd2, 0xb4, 0x23, 0x5a, 0x01, 0x6f, 0x05,
	0xd1, 0xdc, 0x41, 0x73, 0x75, 0xc0, 0xfd, 0x30, 0x91, 0x52, 0x68, 0x96,
	0x45, 0xb3, 0x66, 0x01, 0x3b, 0x53, 0x89, 0x3c, 0x69, 0xbc, 0x6c, 0x69,
	0xe3, 0x51, 0x8f, 0xe3, 0xd2, 0x84, 0xd5, 0x28, 0x66, 0xb5, 0xe6, 0x06,
	0x09, 0xfe, 0x6d, 0xb0, 0x72, 0x16, 0xe0, 0x8a, 0xce, 0x61, 0x65, 0xa9,
	0x21, 0x32, 0x48, 0xdc, 0x7a, 0x1d, 0xe1, 0x38, 0x7f, 0x8c, 0x75, 0x88,
	0x3d, 0x08, 0xa9, 0x4a, 0x6f, 0x3d, 0x9f, 0x7f, 0x3f, 0xbd, 0x57, 0x6b,
	0x19, 0xce, 0x3f, 0x4a, 0xc9, 0xd3, 0xf9, 0x6e, 0x72, 0x7b, 0x5b, 0x74,
	0xea, 0xbe, 0x9c, 0x7a, 0x6d, 0x9c, 0x40, 0x49, 0xe6, 0xfb, 0x2a, 0x1a,
	0x75, 0x70, 0xe5, 0x4e, 0xed, 0x74, 0xe0, 0x75, 0xac, 0xc0, 0xb1, 0x11,
	0x3e, 0xf2, 0xaf, 0x88, 0x4d, 0x66, 0xb6, 0xf6, 0x15, 0x4f, 0x3c, 0x6c,
	0x77, 0xae, 0x47, 0x51, 0x63, 0x9a, 0xfe, 0xe1, 0xb4, 0x1a, 0x12, 0xdf,
	0xe9, 0x54, 0x8d, 0x3b, 0x30, 0x2a, 0x75, 0xe3, 0xe5, 0x29, 0xb1, 0x4c,
	0xb0, 0x7c, 0x6d, 0xb5, 0xae, 0x85, 0xdb, 0x1e, 0x38, 0x55, 0x96, 0xa5,
	0x5b, 0x9f, 0x15, 0x23, 0x28, 0x36, 0xb8, 0xa2, 0x41, 0xb4, 0xd7, 0x19,
	0x91, 0x8d, 0x26, 0x3e, 0xca, 0x9c, 0x05, 0x7a, 0x2b, 0x60, 0x45, 0x86,
	0x8b, 0xee, 0x64, 0x6f, 0x5c, 0x09, 0x4d, 0x4b, 0x5a, 0x7f, 0xb0, 0xc3,
	0x26, 0x9d, 0x8b, 0xb8, 0x83, 0x69, 0xcf, 0x16, 0x72, 0x62, 0x3e, 0x5e,
	0x53, 0x4f, 0x9c, 0x73, 0x76, 0xfc, 0x19, 0xef, 0xa0, 0x74, 0x3a, 0x11,
	0x1e, 0xd0, 0x4d, 0xb7, 0x87, 0xa1, 0xd6, 0x87, 0x6c, 0x0e, 0x6c, 0x8c,
	0xe9, 0xa0, 0x44, 0xc4, 0x72, 0x3e, 0x73, 0x17, 0x13, 0xd1, 0x4e, 0x3d,
	0x8e, 0x1d, 0x5a, 0x8b, 0x75, 0xcb, 0x59, 0x2c, 0x47, 0x87, 0x15, 0x41,
	0xfe, 0x08, 0xe9, 0xa6, 0x97, 0x17, 0x08, 0x26, 0x6a, 0xb5, 0xbb, 0x73,
	0xaa, 0xb8, 0x5b, 0x65, 0x65, 0x5b, 0x30, 0x9e, 0x62, 0x59, 0x02, 0xf8,
	0xb8, 0x0f, 0x32, 0x10, 0xc1, 0x36, 0x08, 0x52, 0x98, 0x4a, 0x1e, 0xf0,
	0xab, 0x21, 0x5e, 0xde, 0x16, 0x0c, 0xda, 0x09, 0x99, 0x6b, 0x9e, 0xc0,
	0x90, 0xa5, 0x5a, 0xcc, 0xb0, 0xb7, 0xbb, 0xd2, 0x8b, 0x5f, 0xd3, 0x3b,
	0x3e, 0x8c, 0xa5, 0x71, 0x66, 0x06, 0xe3, 0x28, 0xd4, 0xf8, 0x3f, 0xe5,
	0x27, 0xdf, 0xfe, 0x0f, 0x09, 0xb2, 0x8a, 0x09, 0x5a, 0x23, 0x61, 0x0d,
	0x2d, 0xf5, 0x44, 0xf1, 0x5c, 0xf8, 0x82, 0x4e, 0xdc, 0x78, 0x7a, 0xab,
	0xc3, 0x57, 0x91, 0xaf, 0x65, 0x6e, 0x71, 0xf1, 0x44, 0xbf, 0xed, 0x43,
	0x50, 0xb4, 0x67, 0x48, 0xef, 0x5a, 0x10, 0x46, 0x81, 0xb4, 0x0c, 0xc8,
	0x48, 0xed, 0x99, 0x7a, 0x45, 0xa5, 0x92, 0xc3, 0x69, 0xd6, 0xd7, 0x8a,
	0x20, 0x1b, 0xeb, 0x8f, 0xb2, 0xff, 0xec, 0x6d, 0x76, 0x04, 0xf8, 0xc2,
	0x58, 0x9b, 0xf2, 0x20, 0x53, 0xc4, 0x74, 0x91, 0x19, 0xdd, 0x2d, 0x12,
	0x53, 0xc7, 0x6e, 0xd0, 0x02, 0x51, 0x3c, 0xa6, 0x7d, 0x80, 0x75, 0x6b,
	0x1d, 0xdf, 0xf8, 0x6a, 0x52, 0xbb, 0x81, 0xf8, 0x30, 0x45, 0xef, 0x51,
	0x85, 0x36, 0xbe, 0x8e, 0xcf, 0x0b, 0x9a, 0x46, 0xe8, 0x3f, 0x99, 0xfd,
	0xf7, 0xd9, 0x3e, 0x84, 0xe5, 0xe3, 0x37, 0xcf, 0x98, 0x7f, 0xeb, 0x5e,
	0x5a, 0x53, 0x77, 0x1c, 0x20, 0xdc, 0xf1, 0x20, 0x99, 0xec, 0x60, 0x40,
	0x93, 0xef, 0x5c, 0x1c, 0x81, 0xe2, 0xa5, 0xad, 0x2a, 0xc2, 0xdb, 0x6b,
	0xc1, 0x7e, 0x8f, 0xa9, 0x23, 0x5b, 0xd9, 0x0d, 0xfe, 0xa0, 0xac, 0x11,
	0x28, 0xba, 0x8e, 0x92, 0x07, 0x2d, 0x07, 0x40, 0x83, 0x14, 0x4c, 0x35,
	0x8d, 0xd0, 0x11, 0xff, 0x98, 0xdb, 0x00, 0x30, 0x6f, 0x65, 0xb6, 0xa0,
	0x7f, 0x9c, 0x08, 0xb8, 0xce, 0xb3, 0xa8, 0x42, 0xd3, 0x84, 0x45, 0xe1,
	0xe3, 0x8f, 0xa6, 0x89, 0x21, 0xd7, 0x74, 0x02, 0x4d, 0x64, 0xdf, 0x54,
	0x15, 0x9e, 0xba, 0x12, 0x49, 0x09, 0x41, 0xf6, 0x10, 0x24, 0xa1, 0x84,
	0x15, 0xfd, 0x68, 0x6a, 0x57, 0x66, 0xb3, 0x6d, 0x4c, 0xea, 0xbf, 0xbc,
	0x60, 0x3f, 0x52, 0x1c, 0x44, 0x1b, 0xc0, 0x4a, 0x25, 0xe3, 0xd9, 0x4c,
	0x9a, 0x74, 0xad, 0xfc, 0x9e, 0x8d, 0x0b, 0x18, 0x66, 0x24, 0xd1, 0x06,
	0xac, 0x68, 0xc1, 0xae, 0x14, 0xce, 0xb1, 0xf3, 0x86, 0x9f, 0x87, 0x11,
	0xd7, 0x9f, 0x30, 0x92, 0xdb, 0xec, 0x0b, 0x4a, 0xe8, 0xf6, 0x53, 0x36,
	0x68, 0x12, 0x11, 0x5e, 0xe0, 0x34, 0xa4, 0xff, 0x00, 0x0a, 0x26, 0xb8,
	0x62, 0x79, 0x9c, 0x0c, 0xd5, 0xe5, 0xf5, 0x1c, 0x1a, 0x16, 0x84, 0x4d,
	0x8e, 0x5d, 0x31, 0x7e, 0xf7, 0xe2, 0xd3, 0xa1, 0x41, 0x90, 0x61, 0x5d,
	0x04, 0xb2, 0x9a, 0x18, 0x9e, 0x54, 0xfb, 0xd1, 0x61, 0x95, 0x1b, 0x08,
	0xca, 0x7c, 0x49, 0x44, 0x74, 0x1d, 0x2f, 0xca, 0xc4, 0x7a, 0xe1, 0x8b,
	0x2f, 0xbb, 0x96, 0xee, 0x19, 0x8a, 0x5d, 0xfb, 0x3e, 0x82, 0xe7, 0x15,
	0xdb, 0x29, 0x14, 0xee, 0xc9, 0x4d, 0x9a, 0xfb, 0x9f, 0x8a, 0xbb, 0x17,
	0x37, 0x1b, 0x6e, 0x28, 0x6c, 0xf9, 0xff, 0xb5, 0xb5, 0x8b, 0x9d, 0x88,
	0x20, 0x08, 0x10, 0xd7, 0xca, 0x58, 0xf6, 0xe1, 0x32, 0x91, 0x6f, 0x36,
	0xc0, 0xad, 0xc1, 0x57, 0x5d, 0x76, 0x31, 0x43, 0xf3, 0xdd, 0xec, 0xf1,
	0xa9, 0x79, 0xe9, 0xe9, 0x85, 0xd7, 0x91, 0xc7, 0x31, 0x62, 0x3c, 0xd2,
	0x90, 0x2c, 0x9c, 0xa4, 0x56, 0x37, 0x7b, 0xbe, 0x40, 0x58, 0xc0, 0x81,
	0x83, 0x22, 0xe8, 0x13, 0x79, 0x18, 0xdb, 0x3a, 0x1b, 0x31, 0x0d, 0x00,
	0x6c, 0x22, 0x62, 0x75, 0x70, 0xd8, 0x96, 0x59, 0x99, 0x44, 0x79, 0x71,
	0xa6, 0x76, 0x81, 0x28, 0xb2, 0x65, 0xe8, 0x47, 0x14, 0xc6, 0x39, 0x06,
}

type scalar struct {
	bytes [32]uint8
	//words [8]uint32
}

var kOrder = scalar{
	bytes: [32]uint8{0xed, 0xd3, 0xf5, 0x5c, 0x1a, 0x63, 0x12, 0x58,
		0xd6, 0x9c, 0xf7, 0xa2, 0xde, 0xf9, 0xde, 0x14,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x10},
}

func leftShift3(n *[]byte) {
	var carry uint8
	for i := 0; i < 32; i++ {
		nextCarry := (*n)[i] >> 5
		(*n)[i] = ((*n)[i] << 3) | carry
		carry = nextCarry
	}
}

func constantTimeMSB(a uint32) uint32 {
	return 0 - (a >> (uint32(unsafe.Sizeof(a))*8 - 1))
}

func constantTimeIsZero(a uint32) uint32 {
	return constantTimeMSB(^a & (a - 1))
}

func constantTimeEq(a, b uint32) uint32 {
	return constantTimeIsZero(a ^ b)
}

func valueBarrier(a uint32) uint32 {
	return a
}

func constantTimeSelect(mask, a, b uint32) uint32 {
	return (valueBarrier(mask) & a) | (valueBarrier(^mask) & b)
}

func scalarCmov(dest, src *scalar, mask uint32) {
	for i := 0; i < 8; i++ {
		data32 := constantTimeSelect(mask, bytesToUint32(src.bytes[i*4:i*4+4]), bytesToUint32(dest.bytes[i*4:i*4+4]))
		data32Byte := uint32ToBytes(data32)
		copy(dest.bytes[i*4:], data32Byte)
	}
}

func scalarDouble(s *scalar) {
	var carry uint32
	for i := 0; i < 8; i++ {

		sword := bytesToUint32(s.bytes[i*4 : i*4+4])

		carryOut := sword >> 31
		tmp := (sword << 1) | carry
		data32Byte := uint32ToBytes(uint32(tmp))
		copy(s.bytes[i*4:], data32Byte)
		carry = carryOut
	}
}

func uint32ToBytes(num uint32) []byte {
	return []byte{
		byte(num),
		byte(num >> 8),
		byte(num >> 16),
		byte(num >> 24),
	}
}

func bytesToUint32(bytes []byte) uint32 {
	if len(bytes) < 4 {
		return 0
	}
	return uint32(bytes[0]) | uint32(bytes[1])<<8 | uint32(bytes[2])<<16 | uint32(bytes[3])<<24
}

func scalarAdd(dest, src *scalar) {
	var carry uint32
	for i := 0; i < 8; i++ {

		destInt := bytesToUint32(dest.bytes[i*4 : i*4+4])
		srcInt := bytesToUint32(src.bytes[i*4 : i*4+4])

		tmp := uint64(destInt) + uint64(srcInt) + uint64(carry)

		data32Byte := uint32ToBytes(uint32(tmp))
		//dest.words[i] = uint32(tmp)
		copy(dest.bytes[i*4:], data32Byte)
		carry = uint32(uint64(tmp) >> 32)
	}
}

type Spake2Ctx struct {
	myRole int
	myName []uint8
	//myNameLen                 uint32
	theirName []uint8
	//theirNameLen              uint32
	privateKey                [32]uint8
	passwordHash              [64]uint8
	passwordScalar            [32]uint8
	myMsg                     []byte
	state                     int
	disablePasswordScalarHack bool
}

func SPAKE2_CTX_new(myRole int, myName []uint8, theirName []uint8) (*Spake2Ctx, error) {
	ctx := &Spake2Ctx{}
	ctx.myRole = myRole
	//ctx.myNameLen = uint32(len(myName))
	ctx.myName = myName
	//ctx.theirNameLen = uint32(len(theirName))
	ctx.theirName = theirName
	ctx.myMsg = make([]byte, 32)
	return ctx, nil
}

func (ctx *Spake2Ctx) SPAKE2_CTX_free() {
	ctx.myName = nil
	ctx.theirName = nil
}

func (ctx *Spake2Ctx) SPAKE2_generate_msg(out []byte, maxOutLen uint32, password []byte) uint32 {
	if ctx.state != 0 {
		return 0
	}

	if int(maxOutLen) < len(ctx.myMsg) {
		return 0
	}

	privateTmp := make([]byte, 64)

	rand.Read(privateTmp)

	privateTmp = x25519ScReduce(privateTmp)
	// Multiply by the cofactor (eight) so that we'll clear it when operating on
	// the peer's point later in the protocol.

	leftShift3(&privateTmp)

	copy(ctx.privateKey[:], privateTmp[:])

	var p ed25519.Ge_p3
	x25519GeScalarmultBase(&p, ctx.privateKey[:])

	// 创建 SHA-512 哈希对象
	hasher := sha512.New()

	// 更新哈希对象
	hasher.Write(password)

	// 获取哈希值
	passwordTmp := hasher.Sum(nil)

	copy(ctx.passwordHash[:], passwordTmp[:])

	passwordTmp = x25519ScReduce(passwordTmp[:])

	var passwordScalar scalar
	copy(passwordScalar.bytes[:], passwordTmp[:])

	// |password_scalar| is the result of |x25519_sc_reduce| and thus is, at
	// most, $l-1$ (where $l$ is |kOrder|, the order of the prime-order subgroup
	// of Ed25519). In the following, we may add $l + 2×l + 4×l$ for a max value
	// of $8×l-1$. That is < 2**256, as required.

	if !ctx.disablePasswordScalarHack {
		var order scalar = kOrder
		var tmp scalar
		scalarCmov(&tmp, &order, constantTimeEq(uint32(passwordScalar.bytes[0]&1), 1))
		scalarAdd(&passwordScalar, &tmp)
		scalarDouble(&order)
		var tmp1 scalar
		scalarCmov(&tmp1, &order, constantTimeEq(uint32(passwordScalar.bytes[0]&2), 2))

		scalarAdd(&passwordScalar, &tmp1)
		scalarDouble(&order)
		var tmp2 scalar
		scalarCmov(&tmp2, &order, constantTimeEq(uint32(passwordScalar.bytes[0]&4), 4))

		//log.Printf("password_scalar:%+v\r\n", passwordScalar)
		scalarAdd(&passwordScalar, &tmp2)
	}

	copy(ctx.passwordScalar[:], passwordScalar.bytes[:])

	var mask ed25519.Ge_p3

	var precompTable []byte
	if ctx.myRole == 0 {
		precompTable = kSpakeMSmallPrecomp
	} else {
		precompTable = kSpakeNSmallPrecomp
	}

	x25519GeScalarmultSmallPrecomp(&mask, ctx.passwordScalar[:], precompTable)

	// P* = P + mask.
	var mask_cached ed25519.Ge_cached
	ed25519.X25519_ge_p3_to_cached(&mask_cached, &mask)

	var Pstar ed25519.Ge_p1p1
	ed25519.X25519_ge_add(&Pstar, &p, &mask_cached)

	// Encode P*
	var Pstar_proj ed25519.Ge_p2
	ed25519.X25519_ge_p1p1_to_p2(&Pstar_proj, &Pstar)
	ed25519.X25519_ge_tobytes(&ctx.myMsg, &Pstar_proj)
	copy(out, ctx.myMsg)
	ctx.state = 1

	return 1
}

func updateWithLengthPrefix(sha hash.Hash, data []uint8) {
	lenLE := make([]uint8, 8)
	binary.LittleEndian.PutUint64(lenLE, uint64(len(data)))
	sha.Write(lenLE)
	sha.Write(data)
}

func (ctx *Spake2Ctx) SPAKE2_process_msg(outKey []uint8, maxOutKeyLen uint32, theirMsg []uint8) (uint32, error) {
	if ctx.state != 1 || len(theirMsg) != 32 { // Assuming 1 is spake2_state_msg_generated
		return 0, errors.New("invalid state or message length")
	}

	var Qstar ed25519.Ge_p3
	if ed25519.X25519_ge_frombytes_vartime(&Qstar, theirMsg) == 0 {
		return 0, errors.New("point received from peer was not on the curve")
	}

	var peersMask ed25519.Ge_p3

	var precompTable []byte
	if ctx.myRole == 0 {
		precompTable = kSpakeNSmallPrecomp
	} else {
		precompTable = kSpakeMSmallPrecomp
	}

	x25519GeScalarmultSmallPrecomp(&peersMask, ctx.passwordScalar[:], precompTable)

	var peersMaskCached ed25519.Ge_cached
	ed25519.X25519_ge_p3_to_cached(&peersMaskCached, &peersMask)

	var QCompl ed25519.Ge_p1p1
	var QExt ed25519.Ge_p3

	ed25519.X25519_ge_sub(&QCompl, &Qstar, &peersMaskCached)
	ed25519.X25519_ge_p1p1_to_p3(&QExt, &QCompl)

	var dhShared ed25519.Ge_p2
	x25519GeScalarmult(&dhShared, ctx.privateKey[:], &QExt)

	dhSharedEncoded := make([]byte, 32)
	ed25519.X25519_ge_tobytes(&dhSharedEncoded, &dhShared)
	hasher := sha512.New()
	if ctx.myRole == 0 { // Assuming 0 is spake2_role_alice

		updateWithLengthPrefix(hasher, ctx.myName)
		updateWithLengthPrefix(hasher, ctx.theirName)
		updateWithLengthPrefix(hasher, ctx.myMsg)
		updateWithLengthPrefix(hasher, theirMsg)

	} else {
		//log.Printf("bob:%s|%s|%+v|%+v\r\n", ctx.theirName, ctx.myName, theirMsg, ctx.myMsg)
		updateWithLengthPrefix(hasher, ctx.theirName)
		updateWithLengthPrefix(hasher, ctx.myName)
		updateWithLengthPrefix(hasher, theirMsg)
		updateWithLengthPrefix(hasher, ctx.myMsg)

	}

	updateWithLengthPrefix(hasher, dhSharedEncoded)
	updateWithLengthPrefix(hasher, ctx.passwordHash[:])

	key := hasher.Sum(nil)
	toCopy := maxOutKeyLen
	if toCopy > uint32(len(key)) {
		toCopy = uint32(len(key))
	}
	copy(outKey, key[:toCopy])
	ctx.state = 2 // Assuming 2 is spake2_state_key_generated

	return toCopy, nil
}
