package sha512Test

import (
	"encoding/binary"
)

const (
	blockSize = 128
)

var fillbuf = [128]byte{0x80}

var K = [80]uint64{
	0x428a2f98d728ae22, 0x7137449123ef65cd,
	0xb5c0fbcfec4d3b2f, 0xe9b5dba58189dbbc,
	0x3956c25bf348b538, 0x59f111f1b605d019,
	0x923f82a4af194f9b, 0xab1c5ed5da6d8118,
	0xd807aa98a3030242, 0x12835b0145706fbe,
	0x243185be4ee4b28c, 0x550c7dc3d5ffb4e2,
	0x72be5d74f27b896f, 0x80deb1fe3b1696b1,
	0x9bdc06a725c71235, 0xc19bf174cf692694,
	0xe49b69c19ef14ad2, 0xefbe4786384f25e3,
	0x0fc19dc68b8cd5b5, 0x240ca1cc77ac9c65,
	0x2de92c6f592b0275, 0x4a7484aa6ea6e483,
	0x5cb0a9dcbd41fbd4, 0x76f988da831153b5,
	0x983e5152ee66dfab, 0xa831c66d2db43210,
	0xb00327c898fb213f, 0xbf597fc7beef0ee4,
	0xc6e00bf33da88fc2, 0xd5a79147930aa725,
	0x06ca6351e003826f, 0x142929670a0e6e70,
	0x27b70a8546d22ffc, 0x2e1b21385c26c926,
	0x4d2c6dfc5ac42aed, 0x53380d139d95b3df,
	0x650a73548baf63de, 0x766a0abb3c77b2a8,
	0x81c2c92e47edaee6, 0x92722c851482353b,
	0xa2bfe8a14cf10364, 0xa81a664bbc423001,
	0xc24b8b70d0f89791, 0xc76c51a30654be30,
	0xd192e819d6ef5218, 0xd69906245565a910,
	0xf40e35855771202a, 0x106aa07032bbd1b8,
	0x19a4c116b8d2d0c8, 0x1e376c085141ab53,
	0x2748774cdf8eeb99, 0x34b0bcb5e19b48a8,
	0x391c0cb3c5c95a63, 0x4ed8aa4ae3418acb,
	0x5b9cca4f7763e373, 0x682e6ff3d6b2b8a3,
	0x748f82ee5defb2fc, 0x78a5636f43172f60,
	0x84c87814a1f0ab72, 0x8cc702081a6439ec,
	0x90befffa23631e28, 0xa4506cebde82bde9,
	0xbef9a3f7b2c67915, 0xc67178f2e372532b,
	0xca273eceea26619c, 0xd186b8c721c0c207,
	0xeada7dd6cde0eb1e, 0xf57d4f7fee6ed178,
	0x06f067aa72176fba, 0x0a637dc5a2c898a6,
	0x113f9804bef90dae, 0x1b710b35131c471b,
	0x28db77f523047d84, 0x32caab7b40c72493,
	0x3c9ebe0a15c9bebc, 0x431d67c49c100d4c,
	0x4cc5d4becb3e42b6, 0x597f299cfc657e2a,
	0x5fcb6fab3ad6faec, 0x6c44198c4a475817,
}

type Sha512Ctx struct {
	H      [8]uint64
	total  [2]uint64
	buflen uint64
	buffer [blockSize]byte
}

func Sha512InitCtx(ctx *Sha512Ctx) {
	ctx.H[0] = 0x6a09e667f3bcc908
	ctx.H[1] = 0xbb67ae8584caa73b
	ctx.H[2] = 0x3c6ef372fe94f82b
	ctx.H[3] = 0xa54ff53a5f1d36f1
	ctx.H[4] = 0x510e527fade682d1
	ctx.H[5] = 0x9b05688c2b3e6c1f
	ctx.H[6] = 0x1f83d9abfb41bd6b
	ctx.H[7] = 0x5be0cd19137e2179

	ctx.total[0] = 0
	ctx.total[1] = 0
	ctx.buflen = 0
}

func swapBytes(n uint64) uint64 {
	return (n<<56 | ((n>>8)&0xff00)<<40 | ((n>>24)&0xff0000)<<24 | ((n>>40)&0xff000000)<<8 |
		((n << 8) & 0xff000000) | ((n << 24) & 0xff0000) | ((n << 40) & 0xff00) | (n >> 56))
}

func Sha512FinishCtx(ctx *Sha512Ctx, resbuf []byte) []byte {
	bytes := ctx.buflen
	var pad uint64

	if bytes >= 112 {
		pad = 128 + 112 - bytes
	} else {
		pad = 112 - bytes
	}
	copy(ctx.buffer[bytes:], fillbuf[:pad])

	ctx.buffer[bytes+pad+8/8] = uint8(swapBytes(ctx.total[0] << 3))
	ctx.buffer[bytes+pad/8] = uint8(swapBytes((ctx.total[1] << 3) | (ctx.total[0] >> 61)))

	sha512ProcessBlock(ctx.buffer[:], bytes+pad+16, ctx)

	for i := 0; i < 8; i++ {
		binary.BigEndian.PutUint64(resbuf[i*8:], swap(ctx.H[i]))
	}

	return resbuf
}

func Sha512ProcessBytes(buffer []byte, ctx *Sha512Ctx) {
	if ctx.buflen != 0 {
		leftOver := ctx.buflen
		add := min(256-leftOver, uint64(len(buffer)))

		copy(ctx.buffer[leftOver:], buffer[:add])
		ctx.buflen += add

		if ctx.buflen > 128 {
			sha512ProcessBlock(ctx.buffer[:], ctx.buflen&^127, ctx)

			ctx.buflen &= 127
			copy(ctx.buffer[:], ctx.buffer[(leftOver+add)&^127:][:ctx.buflen])
		}

		buffer = buffer[add:]
	}

	if len(buffer) >= 128 {
		sha512ProcessBlock(buffer, uint64(len(buffer)&^127), ctx)
		buffer = buffer[len(buffer)&^127:]
	}

	if len(buffer) > 0 {
		leftOver := ctx.buflen

		copy(ctx.buffer[leftOver:], buffer)
		leftOver += uint64(len(buffer))
		if leftOver >= 128 {
			sha512ProcessBlock(ctx.buffer[:], 128, ctx)
			leftOver -= 128
			copy(ctx.buffer[:], ctx.buffer[128:][:leftOver])
		}
		ctx.buflen = leftOver
	}
}

func sha512ProcessBlock(buffer []byte, len uint64, ctx *Sha512Ctx) {
	words := make([]uint64, len/8)
	for i := range words {
		words[i] = swap(binary.BigEndian.Uint64(buffer[i*8:]))
	}
	nwords := len / 8
	a, b, c, d, e, f, g, h := ctx.H[0], ctx.H[1], ctx.H[2], ctx.H[3], ctx.H[4], ctx.H[5], ctx.H[6], ctx.H[7]

	ctx.total[0] += len
	if ctx.total[0] < len {
		ctx.total[1]++
	}

	for nwords > 0 {
		W := make([]uint64, 80)
		aSave, bSave, cSave, dSave, eSave, fSave, gSave, hSave := a, b, c, d, e, f, g, h

		for t := 0; t < 16; t++ {
			W[t] = words[t]
		}
		for t := 16; t < 80; t++ {
			W[t] = R1(W[t-2]) + W[t-7] + R0(W[t-15]) + W[t-16]
		}

		for t := 0; t < 80; t++ {
			T1 := h + S1(e) + Ch(e, f, g) + K[t] + W[t]
			T2 := S0(a) + Maj(a, b, c)
			h = g
			g = f
			f = e
			e = d + T1
			d = c
			c = b
			b = a
			a = T1 + T2
		}

		a += aSave
		b += bSave
		c += cSave
		d += dSave
		e += eSave
		f += fSave
		g += gSave
		h += hSave

		nwords -= 16
	}

	ctx.H[0] = a
	ctx.H[1] = b
	ctx.H[2] = c
	ctx.H[3] = d
	ctx.H[4] = e
	ctx.H[5] = f
	ctx.H[6] = g
	ctx.H[7] = h
}

func swap(n uint64) uint64 {
	return (n << 56) | ((n & 0xff00) << 40) | ((n & 0xff0000) << 24) | ((n & 0xff000000) << 8) | ((n >> 8) & 0xff000000) | ((n >> 24) & 0xff0000) | ((n >> 40) & 0xff00) | (n >> 56)
}

func Ch(x, y, z uint64) uint64 {
	return (x & y) ^ (^x & z)
}

func Maj(x, y, z uint64) uint64 {
	return (x & y) ^ (x & z) ^ (y & z)
}

func S0(x uint64) uint64 {
	return CYCLIC(x, 28) ^ CYCLIC(x, 34) ^ CYCLIC(x, 39)
}

func S1(x uint64) uint64 {
	return CYCLIC(x, 14) ^ CYCLIC(x, 18) ^ CYCLIC(x, 41)
}

func R0(x uint64) uint64 {
	return CYCLIC(x, 1) ^ CYCLIC(x, 8) ^ (x >> 7)
}

func R1(x uint64) uint64 {
	return CYCLIC(x, 19) ^ CYCLIC(x, 61) ^ (x >> 6)
}

func CYCLIC(w, s uint64) uint64 {
	return (w >> s) | (w << (64 - s))
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}
	return b
}
