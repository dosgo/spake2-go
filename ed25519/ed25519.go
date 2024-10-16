package main

import (
	"encoding/binary"
	"errors"
)

const FE_NUM_LIMBS = 10

type fe struct {
	v [FE_NUM_LIMBS]uint32
}

type fe_loose struct {
	v [FE_NUM_LIMBS]uint32
}

var d = fe{
	v: [FE_NUM_LIMBS]uint32{
		56195235, 13857412, 51736253, 6949390, 114729, 24766616, 60832955, 30306712,
		48412415, 21499315,
	},
}

var sqrtm1 = fe{
	v: [FE_NUM_LIMBS]uint32{
		34513072, 25610706, 9377949, 3500415, 12389472, 33281959, 41962654,
		31548777, 326685, 11406482,
	},
}

var d2 = fe{
	v: [FE_NUM_LIMBS]uint32{
		45281625, 27714825, 36363642, 13898781, 229458, 15978800, 54557047,
		27058993, 29715967, 9444199,
	},
}

func fe_0(h *fe) {
	for i := range h.v {
		h.v[i] = 0
	}
}

func fe_loose_0(h *fe_loose) {
	for i := range h.v {
		h.v[i] = 0
	}
}

func fe_1(h *fe) {
	fe_0(h)
	h.v[0] = 1
}

func fe_loose_1(h *fe_loose) {
	fe_loose_0(h)
	h.v[0] = 1
}

func fe_copy(h *fe, f *fe) {
	h.v = f.v
}

func fe_copy_lt(h *fe_loose, f *fe) {
	h.v = f.v
}

func fe_frombytes_strict(h *fe, s [32]byte) error {
	if s[31]&0x80 != 0 {
		return errors.New("top-most bit must be clear")
	}
	// fiat_25519_from_bytes(h.v[:], s[:])
	return nil // Placeholder for actual implementation
}

func fe_frombytes(h *fe, s [32]byte) {
	var s_copy [32]byte
	copy(s_copy[:], s[:])
	s_copy[31] &= 0x7f
	_ = fe_frombytes_strict(h, s_copy)
}

func fe_tobytes(s *[32]byte, f *fe) {
	// fiat_25519_to_bytes(s[:], f.v[:])
}

func fe_add(h *fe_loose, f *fe, g *fe) {
	// fiat_25519_add(h.v[:], f.v[:], g.v[:])
}

func fe_sub(h *fe_loose, f *fe, g *fe) {
	// fiat_25519_sub(h.v[:], f.v[:], g.v[:])
}

func fe_carry(h *fe, f *fe_loose) {
	// fiat_25519_carry(h.v[:], f.v[:])
}

func fe_mul_impl(out *[FE_NUM_LIMBS]uint32, in1 *[FE_NUM_LIMBS]uint32, in2 *[FE_NUM_LIMBS]uint32) {
	// fiat_25519_carry_mul(out[:], in1[:], in2[:])
}

func fe_mul_ltt(h *fe_loose, f *fe, g *fe) {
	fe_mul_impl(&h.v, &f.v, &g.v)
}

func fe_mul_llt(h *fe_loose, f *fe_loose, g *fe) {
	fe_mul_impl(&h.v, &f.v, &g.v)
}

func fe_mul_ttt(h *fe, f *fe, g *fe) {
	fe_mul_impl(&h.v, &f.v, &g.v)
}

func fe_mul_tlt(h *fe, f *fe_loose, g *fe) {
	fe_mul_impl(&h.v, &f.v, &g.v)
}

func fe_mul_ttl(h *fe, f *fe, g *fe_loose) {
	fe_mul_impl(&h.v, &f.v, &g.v)
}

func fe_mul_tll(h *fe, f *fe_loose, g *fe_loose) {
	fe_mul_impl(&h.v, &f.v, &g.v)
}

func fe_sq_tl(h *fe, f *fe_loose) {
	// fiat_25519_carry_square(h.v[:], f.v[:])
}

func fe_sq_tt(h *fe, f *fe) {
	// fiat_25519_carry_square(h.v[:], f.v[:])
}

func fe_neg(h *fe_loose, f *fe) {
	// fiat_25519_opp(h.v[:], f.v[:])
}

func fe_cmov(f *fe_loose, g *fe_loose, b uint32) {
	b = 0 - b
	for i := 0; i < FE_NUM_LIMBS; i++ {
		x := f.v[i] ^ g.v[i]
		x &= b
		f.v[i] ^= x
	}
}

func fe_copy_ll(h *fe_loose, f *fe_loose) {
	h.v = f.v
}

func fe_loose_invert(out *fe, z *fe_loose) {
	var t0, t1, t2, t3 fe
	for i := 1; i < 2; i++ {
		fe_sq_tt(&t1, &t1)
	}
	fe_mul_tlt(&t1, z, &t1)
	fe_mul_ttt(&t0, &t0, &t1)
	fe_sq_tt(&t2, &t0)
	fe_mul_ttt(&t1, &t1, &t2)
	fe_sq_tt(&t2, &t1)
	for i := 1; i < 5; i++ {
		fe_sq_tt(&t2, &t2)
	}
	fe_mul_ttt(&t1, &t2, &t1)
	fe_sq_tt(&t2, &t1)
	for i := 1; i < 10; i++ {
		fe_sq_tt(&t2, &t2)
	}
	fe_mul_ttt(&t2, &t2, &t1)
	fe_sq_tt(&t3, &t2)
	for i := 1; i < 20; i++ {
		fe_sq_tt(&t3, &t3)
	}
	fe_mul_ttt(&t2, &t3, &t2)
	fe_sq_tt(&t2, &t2)
	for i := 1; i < 10; i++ {
		fe_sq_tt(&t2, &t2)
	}
	fe_mul_ttt(&t1, &t2, &t1)
	fe_sq_tt(&t2, &t1)
	for i := 1; i < 50; i++ {
		fe_sq_tt(&t2, &t2)
	}
	fe_mul_ttt(&t2, &t2, &t1)
	fe_sq_tt(&t3, &t2)
	for i := 1; i < 100; i++ {
		fe_sq_tt(&t3, &t3)
	}
	fe_mul_ttt(&t2, &t3, &t2)
	fe_sq_tt(&t2, &t2)
	for i := 1; i < 50; i++ {
		fe_sq_tt(&t2, &t2)
	}
	fe_mul_ttt(&t1, &t2, &t1)
	fe_sq_tt(&t1, &t1)
	for i := 1; i < 5; i++ {
		fe_sq_tt(&t1, &t1)
	}
	fe_mul_ttt(out, &t1, &t0)
}

func fe_invert(out *fe, z *fe) {
	var l fe_loose
	fe_copy_lt(&l, z)
	fe_loose_invert(out, &l)
}

func fe_isnonzero(f *fe_loose) int {
	var tight fe
	fe_carry(&tight, f)
	var s [32]byte
	fe_tobytes(&s, &tight)

	var zero [32]byte
	return binary.BigEndian.Uint32(s[:]) != binary.BigEndian.Uint32(zero[:])
}

func fe_isnegative(f *fe) int {
	var s [32]byte
	fe_tobytes(&s, f)
	return int(s[0] & 1)
}

func fe_sq2_tt(h *fe, f *fe) {
	fe_sq_tt(h, f)
	var tmp fe_loose
	fe_add(&tmp, h, h)
	fe_carry(h, &tmp)
}

func fe_pow22523(out *fe, z *fe) {
	var t0, t1, t2 fe
	for i := 1; i < 2; i++ {
		fe_sq_tt(&t1, &t1)
	}
	fe_mul_ttt(&t1, z, &t1)
	fe_mul_ttt(&t0, &t0, &t1)
	fe_sq_tt(&t0, &t0)
	fe_mul_ttt(&t0, &t1, &t0)
	fe_sq_tt(&t1, &t0)
	for i := 1; i < 5; i++ {
		fe_sq_tt(&t1, &t1)
	}
	fe_mul_ttt(&t0, &t1, &t0)
	fe_sq_tt(&t1, &t0)
	for i := 1; i < 10; i++ {
		fe_sq_tt(&t1, &t1)
	}
	fe_mul_ttt(&t1, &t1, &t0)
	fe_sq_tt(&t2, &t1)
	for i := 1; i < 20; i++ {
		fe_sq_tt(&t2, &t2)
	}
	fe_mul_ttt(&t1, &t2, &t1)
	fe_sq_tt(&t1, &t1)
	for i := 1; i < 10; i++ {
		fe_sq_tt(&t1, &t1)
	}
	fe_mul_ttt(&t0, &t1, &t0)
	fe_sq_tt(&t1, &t0)
	for i := 1; i < 50; i++ {
		fe_sq_tt(&t1, &t1)
	}
	fe_mul_ttt(&t1, &t1, &t0)
	fe_sq_tt(&t2, &t1)
	for i := 1; i < 100; i++ {
		fe_sq_tt(&t2, &t2)
	}
	fe_mul_ttt(&t1, &t2, &t1)
	fe_sq_tt(&t1, &t1)
	for i := 1; i < 50; i++ {
		fe_sq_tt(&t1, &t1)
	}
	fe_mul_ttt(&t0, &t1, &t0)
	fe_sq_tt(&t0, &t0)
	for i := 1; i < 2; i++ {
		fe_sq_tt(&t0, &t0)
	}
	fe_mul_ttt(out, &t0, z)
}

type ge_p2 struct {
	X, Y, Z fe
}

type ge_p3 struct {
	X, Y, Z, T fe
}

type ge_p1p1 struct {
	X, Y, Z, T fe_loose
}

type ge_precomp struct {
	yplusx, yminusx, xy2d fe_loose
}

type ge_cached struct {
	YplusX, YminusX, Z, T2d fe_loose
}

func ge_p2_0(h *ge_p2) {
	fe_0(&h.X)
	fe_1(&h.Y)
	fe_1(&h.Z)
}

func ge_p3_0(h *ge_p3) {
	fe_0(&h.X)
	fe_1(&h.Y)
	fe_1(&h.Z)
	fe_0(&h.T)
}

func ge_cached_0(h *ge_cached) {
	fe_loose_1(&h.YplusX)
	fe_loose_1(&h.YminusX)
	fe_loose_1(&h.Z)
	fe_loose_0(&h.T2d)
}

func ge_precomp_0(h *ge_precomp) {
	fe_loose_1(&h.yplusx)
	fe_loose_1(&h.yminusx)
	fe_loose_0(&h.xy2d)
}

func ge_p3_to_p2(r *ge_p2, p *ge_p3) {
	fe_copy(&r.X, &p.X)
	fe_copy(&r.Y, &p.Y)
	fe_copy(&r.Z, &p.Z)
}

func x25519_ge_p3_to_cached(r *ge_cached, p *ge_p3) {
	fe_add(&r.YplusX, &p.Y, &p.X)
	fe_sub(&r.YminusX, &p.Y, &p.X)
	fe_copy_lt(&r.Z, &p.Z)
	fe_mul_ltt(&r.T2d, &p.T, &d2)
}

func x25519_ge_p1p1_to_p2(r *ge_p2, p *ge_p1p1) {
	fe_mul_tll(&r.X, &p.X, &p.T)
	fe_mul_tll(&r.Y, &p.Y, &p.Z)
	fe_mul_tll(&r.Z, &p.Z, &p.T)
}

func x25519_ge_p1p1_to_p3(r *ge_p3, p *ge_p1p1) {
	fe_mul_tll(&r.X, &p.X, &p.T)
	fe_mul_tll(&r.Y, &p.Y, &p.Z)
	fe_mul_tll(&r.Z, &p.Z, &p.T)
	fe_mul_tll(&r.T, &p.X, &p.Y)
}

func ge_p1p1_to_cached(r *ge_cached, p *ge_p1p1) {
	var t ge_p3
	x25519_ge_p1p1_to_p3(&t, p)
	x25519_ge_p3_to_cached(r, &t)
}

func ge_p2_dbl(r *ge_p1p1, p *ge_p2) {
	var trX, trZ, trT fe
	var t0 fe

	fe_sq_tt(&trX, &p.X)
	fe_sq_tt(&trZ, &p.Y)
	fe_sq2_tt(&trT, &p.Z)
	fe_add(&r.Y, &p.X, &p.Y)
	fe_sq_tl(&t0, &r.Y)

	fe_add(&r.Y, &trZ, &trX)
	fe_sub(&r.Z, &trZ, &trX)
	fe_carry(&trZ, &r.Y)
	fe_sub(&r.X, &t0, &trZ)
	fe_carry(&trZ, &r.Z)
	fe_sub(&r.T, &trT, &trZ)
}

func x25519_ge_tobytes(s *[32]byte, h *ge_p2) {
	var recip, x, y fe

	fe_invert(&recip, &h.Z)
	fe_mul_ttt(&x, &h.X, &recip)
	fe_mul_ttt(&y, &h.Y, &recip)
	fe_tobytes(s, &y)
	s[31] ^= byte(fe_isnegative(&x) << 7)
}

func ge_p3_tobytes(s *[32]byte, h *ge_p3) {
	var recip, x, y fe

	fe_invert(&recip, &h.Z)
	fe_mul_ttt(&x, &h.X, &recip)
	fe_mul_ttt(&y, &h.Y, &recip)
	fe_tobytes(s, &y)
	s[31] ^= byte(fe_isnegative(&x) << 7)
}

func ge_madd(r *ge_p1p1, p *ge_p3, q *ge_precomp) {
	var trY, trZ, trT fe

	fe_add(&r.X, &p.Y, &p.X)
	fe_sub(&r.Y, &p.Y, &p.X)
	fe_mul_tll(&trZ, &r.X, &q.yplusx)
	fe_mul_tll(&trY, &r.Y, &q.yminusx)
	fe_mul_tlt(&trT, &q.xy2d, &p.T)
	fe_add(&r.T, &p.Z, &p.Z)
	fe_sub(&r.X, &trZ, &trY)
	fe_add(&r.Y, &trZ, &trY)
	fe_carry(&trZ, &r.T)
	fe_add(&r.Z, &trZ, &trT)
	fe_sub(&r.T, &trZ, &trT)
}

func ge_msub(r *ge_p1p1, p *ge_p3, q *ge_precomp) {
	var trY, trZ, trT fe

	fe_add(&r.X, &p.Y, &p.X)
	fe_sub(&r.Y, &p.Y, &p.X)
	fe_mul_tll(&trZ, &r.X, &q.yminusx)
	fe_mul_tll(&trY, &r.Y, &q.yplusx)
	fe_mul_tlt(&trT, &q.xy2d, &p.T)
	fe_add(&r.T, &p.Z, &p.Z)
	fe_sub(&r.X, &trZ, &trY)
	fe_add(&r.Y, &trZ, &trY)
	fe_carry(&trZ, &r.T)
	fe_sub(&r.Z, &trZ, &trT)
	fe_add(&r.T, &trZ, &trT)
}

func x25519_ge_add(r *ge_p1p1, p *ge_p3, q *ge_cached) {
	var trX, trY, trZ, trT fe

	fe_add(&r.X, &p.Y, &p.X)
	fe_sub(&r.Y, &p.Y, &p.X)
	fe_mul_tll(&trZ, &r.X, &q.YplusX)
	fe_mul_tll(&trY, &r.Y, &q.YminusX)
	fe_mul_tlt(&trT, &q.T2d, &p.T)
	fe_mul_ttl(&trX, &p.Z, &q.Z)
	fe_add(&r.T, &trX, &trX)
	fe_sub(&r.X, &trZ, &trY)
	fe_add(&r.Y, &trZ, &trY)
	fe_carry(&trZ, &r.T)
	fe_add(&r.Z, &trZ, &trT)
	fe_sub(&r.T, &trZ, &trT)
}

func x25519_ge_sub(r *ge_p1p1, p *ge_p3, q *ge_cached) {
	var trX, trY, trZ, trT fe

	fe_add(&r.X, &p.Y, &p.X)
	fe_sub(&r.Y, &p.Y, &p.X)
	fe_mul_tll(&trZ, &r.X, &q.YminusX)
	fe_mul_tll(&trY, &r.Y, &q.YplusX)
	fe_mul_tlt(&trT, &q.T2d, &p.T)
	fe_mul_ttl(&trX, &p.Z, &q.Z)
	fe_add(&r.T, &trX, &trX)
	fe_sub(&r.X, &trZ, &trY)
	fe_add(&r.Y, &trZ, &trY)
	fe_carry(&trZ, &r.T)
	fe_sub(&r.Z, &trZ, &trT)
	fe_add(&r.T, &trZ, &trT)
}

func x25519_ge_frombytes_vartime(h *ge_p3, s [32]byte) int {
	var u, v fe
	var v3, vxx fe
	var check fe_loose

	fe_frombytes(&h.Y, s)
	fe_1(&h.Z)
	fe_sq_tt(&v3, &h.Y)
	fe_mul_ttt(&vxx, &v3, &d)
	fe_sub(&u, &v3, &h.Z) // u = y^2-1
	fe_carry(&u, &v)
	fe_add(&v, &vxx, &h.Z) // v = dy^2+1

	fe_sq_tl(&v3, &v)
	fe_mul_ttl(&v3, &v3, &v) // v3 = v^3
	fe_sq_tt(&h.X, &v3)
	fe_mul_ttl(&h.X, &h.X, &v)
	fe_mul_ttt(&h.X, &h.X, &u) // x = uv^7

	fe_pow22523(&h.X, &h.X) // x = (uv^7)^((q-5)/8)
	fe_mul_ttt(&h.X, &h.X, &v3)
	fe_mul_ttt(&h.X, &h.X, &u) // x = uv^3(uv^7)^((q-5)/8)

	fe_sq_tt(&vxx, &h.X)
	fe_mul_ttl(&vxx, &vxx, &v)
	fe_sub(&check, &vxx, &u)
	if fe_isnonzero(&check) != 0 {
		fe_add(&check, &vxx, &u)
		if fe_isnonzero(&check) != 0 {
			return 0
		}
		fe_mul_ttt(&h.X, &h.X, sqrtm1)
	}

	if fe_isnegative(&h.X) != int(s[31]>>7) {
		var t fe_loose
		fe_neg(&t, &h.X)
		fe_carry(&h.X, &t)
	}

	fe_mul_ttt(&h.T, &h.X, &h.Y)
	return 1
}
