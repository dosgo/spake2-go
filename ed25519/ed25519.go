package ed25519

import (
	"bytes"
	"math/big"
	"spake2-go/curve25519"
)

const FE_NUM_LIMBS = 10

type Fe struct {
	v [FE_NUM_LIMBS]uint32
}

type Fe_loose struct {
	v [FE_NUM_LIMBS]uint32
}

var d = Fe{
	v: [FE_NUM_LIMBS]uint32{
		56195235, 13857412, 51736253, 6949390, 114729, 24766616, 60832955, 30306712,
		48412415, 21499315,
	},
}

var sqrtm1 = Fe{
	v: [FE_NUM_LIMBS]uint32{
		34513072, 25610706, 9377949, 3500415, 12389472, 33281959, 41962654,
		31548777, 326685, 11406482,
	},
}

var D2 = Fe{
	v: [FE_NUM_LIMBS]uint32{
		45281625, 27714825, 36363642, 13898781, 229458, 15978800, 54557047,
		27058993, 29715967, 9444199,
	},
}

func fe_0(h *Fe) {
	for i := range h.v {
		h.v[i] = 0
	}
}

func fe_loose_0(h *Fe_loose) {
	for i := range h.v {
		h.v[i] = 0
	}
}

func fe_1(h *Fe) {
	fe_0(h)
	h.v[0] = 1
}

func fe_loose_1(h *Fe_loose) {
	fe_loose_0(h)
	h.v[0] = 1
}

func fe_copy(h *Fe, f *Fe) {
	h.v = f.v
}

func fe_copy_lt(h *Fe_loose, f *Fe) {
	h.v = f.v
}

func Fe_frombytes_strict(h *Fe, s []byte) {
	// |fiat_25519_from_bytes| requires the top-most bit be clear.
	suint32 := convertToUint32v2(s)
	var arg1 [32]uint32
	copy(arg1[:], suint32)
	curve25519.Fiat_25519_from_bytes(&h.v, &arg1)
}
func Fe_frombytes(h *Fe, s []byte) {
	var s_copy [32]byte
	copy(s_copy[:], s[:])
	s_copy[31] &= 0x7f
	Fe_frombytes_strict(h, s_copy[:])
}

func Fe_tobytes(s *[]byte, f *Fe) {
	var out1 [32]uint32
	curve25519.Fiat_25519_to_bytes(&out1, &f.v)
	*s = convertToByte(out1)
}

func convertToByte(arr [32]uint32) []byte {
	var result []byte
	for _, v := range arr {
		result = append(result, uint8(v))
	}
	return result
}

func convertToUint32v2(bytes []byte) []uint32 {
	var result []uint32
	for i := 0; i < len(bytes); i++ {
		result = append(result, uint32(bytes[i]))
	}
	return result
}

func Fe_add(h *Fe_loose, f *Fe, g *Fe) {
	curve25519.Fiat_25519_add(&h.v, &f.v, &g.v)
}

func Fe_sub(h *Fe_loose, f *Fe, g *Fe) {
	curve25519.Fiat_25519_sub(&h.v, &f.v, &g.v)
}

func Fe_carry(h *Fe, f *Fe_loose) {
	curve25519.Fiat_25519_carry(&h.v, &f.v)
}

func Fe_mul_impl(out [10]uint32, in1 [10]uint32, in2 [10]uint32) [10]uint32 {
	curve25519.Fiat_25519_carry_mul(&out, &in1, &in2)
	return out
}

func Fe_mul_ltt(h *Fe_loose, f *Fe, g *Fe) {

	h.v = Fe_mul_impl(h.v, f.v, g.v)

}

func Fe_mul_llt(h *Fe_loose, f *Fe_loose, g *Fe) {
	h.v = Fe_mul_impl(h.v, f.v, g.v)
}

func Fe_mul_ttt(h *Fe, f *Fe, g *Fe) {

	h.v = Fe_mul_impl(h.v, f.v, g.v)
}

func Fe_mul_tlt(h *Fe, f *Fe_loose, g *Fe) {
	h.v = Fe_mul_impl(h.v, f.v, g.v)
}

func fe_mul_ttl(h *Fe, f *Fe, g *Fe_loose) {
	h.v = Fe_mul_impl(h.v, f.v, g.v)
}

func fe_mul_tll(h *Fe, f *Fe_loose, g *Fe_loose) {
	h.v = Fe_mul_impl(h.v, f.v, g.v)
}

func fe_sq_tl(h *Fe, f *Fe_loose) {
	curve25519.Fiat_25519_carry_square(&h.v, &f.v)
}

func fe_sq_tt(h *Fe, f *Fe) {
	curve25519.Fiat_25519_carry_square(&h.v, &f.v)
}

func fe_neg(h *Fe_loose, f *Fe) {
	curve25519.Fiat_25519_opp(&h.v, &f.v)
}

func Fe_cmov22(f *Fe_loose, g *Fe_loose, b uint32) {
	b = 0 - b
	for i := 0; i < FE_NUM_LIMBS; i++ {
		x := f.v[i] ^ g.v[i]
		x &= b
		f.v[i] ^= x
	}
}

func Fe_cmovii(f, g *Fe_loose, b uint32) {
	b = 0 - b
	for i := range f.v {
		x := f.v[i] ^ g.v[i]
		x &= b
		f.v[i] ^= x
	}
}

func Fe_cmov(f, g *Fe_loose, b uint32) {
	b = 0 - b
	for i := range f.v {
		fBig := big.NewInt(int64(f.v[i]))
		gBig := big.NewInt(int64(g.v[i]))
		x := new(big.Int).Xor(fBig, gBig)
		temp := new(big.Int).SetUint64(uint64(b))
		x.And(x, temp)
		f.v[i] = uint32(new(big.Int).Xor(fBig, x).Uint64())
	}
}

func fe_copy_ll(h *Fe_loose, f *Fe_loose) {
	h.v = f.v
}

func fe_loose_invert(out *Fe, z *Fe_loose) {
	var t0 Fe
	var t1 Fe
	var t2 Fe
	var t3 Fe

	var i int

	fe_sq_tl(&t0, z)

	fe_sq_tt(&t1, &t0)

	for i = 1; i < 2; i++ {
		fe_sq_tt(&t1, &t1)
	}
	Fe_mul_tlt(&t1, z, &t1)
	Fe_mul_ttt(&t0, &t0, &t1)
	fe_sq_tt(&t2, &t0)
	Fe_mul_ttt(&t1, &t1, &t2)
	fe_sq_tt(&t2, &t1)
	for i = 1; i < 5; i++ {
		fe_sq_tt(&t2, &t2)
	}
	Fe_mul_ttt(&t1, &t2, &t1)
	fe_sq_tt(&t2, &t1)
	for i = 1; i < 10; i++ {
		fe_sq_tt(&t2, &t2)
	}
	Fe_mul_ttt(&t2, &t2, &t1)
	fe_sq_tt(&t3, &t2)
	for i = 1; i < 20; i++ {
		fe_sq_tt(&t3, &t3)
	}
	Fe_mul_ttt(&t2, &t3, &t2)
	fe_sq_tt(&t2, &t2)
	for i = 1; i < 10; i++ {
		fe_sq_tt(&t2, &t2)
	}
	Fe_mul_ttt(&t1, &t2, &t1)
	fe_sq_tt(&t2, &t1)
	for i = 1; i < 50; i++ {
		fe_sq_tt(&t2, &t2)
	}
	Fe_mul_ttt(&t2, &t2, &t1)
	fe_sq_tt(&t3, &t2)
	for i = 1; i < 100; i++ {
		fe_sq_tt(&t3, &t3)
	}
	Fe_mul_ttt(&t2, &t3, &t2)
	fe_sq_tt(&t2, &t2)
	for i = 1; i < 50; i++ {
		fe_sq_tt(&t2, &t2)
	}
	Fe_mul_ttt(&t1, &t2, &t1)
	fe_sq_tt(&t1, &t1)
	for i = 1; i < 5; i++ {
		fe_sq_tt(&t1, &t1)
	}
	Fe_mul_ttt(out, &t1, &t0)
}

func fe_invert(out *Fe, z *Fe) {
	var l Fe_loose
	fe_copy_lt(&l, z)
	fe_loose_invert(out, &l)
}

/*
func fe_isnonzero(f *Fe_loose) int {
	var tight Fe
	Fe_carry(&tight, f)
	//var s [32]byte
	var s = make([]byte, 32)
	Fe_tobytes(&s, &tight)

	var zero [32]byte
	if binary.BigEndian.Uint32(s[:]) != binary.BigEndian.Uint32(zero[:]) {
		return 1
	}
	return 0
}*/

func fe_isnonzero(f *Fe_loose) bool {
	var tight Fe
	Fe_carry(&tight, f)
	var s = make([]byte, 32)
	Fe_tobytes(&s, &tight)
	zero := make([]byte, 32)
	// 比较 s 和 zero
	return bytes.Compare(s, zero) != 0
}

/*
func fe_isnonzero(f *Fe_loose) int {
	var tight Fe
	Fe_carry(&tight, f)
	s := make([]byte, 32)
	Fe_tobytes(&s, &tight)
	zero := make([]byte, 32)
	if !bytes.Equal(s, zero) {
		return 1
	}
	return 0
}*/

func fe_isnegative(f *Fe) int {
	//var s [32]byte
	var s = make([]byte, 32)
	Fe_tobytes(&s, f)
	return int(s[0] & 1)
}

func fe_sq2_tt(h *Fe, f *Fe) {
	fe_sq_tt(h, f)
	var tmp Fe_loose
	Fe_add(&tmp, h, h)
	Fe_carry(h, &tmp)
}

func fe_pow22523(out *Fe, z *Fe) {
	var t0, t1, t2 Fe
	var i int

	fe_sq_tt(&t0, z)
	fe_sq_tt(&t1, &t0)
	for i = 1; i < 2; i++ {
		fe_sq_tt(&t1, &t1)
	}
	Fe_mul_ttt(&t1, z, &t1)
	Fe_mul_ttt(&t0, &t0, &t1)
	fe_sq_tt(&t0, &t0)
	Fe_mul_ttt(&t0, &t1, &t0)
	fe_sq_tt(&t1, &t0)
	for i = 1; i < 5; i++ {
		fe_sq_tt(&t1, &t1)
	}
	Fe_mul_ttt(&t0, &t1, &t0)
	fe_sq_tt(&t1, &t0)
	for i = 1; i < 10; i++ {
		fe_sq_tt(&t1, &t1)
	}
	Fe_mul_ttt(&t1, &t1, &t0)
	fe_sq_tt(&t2, &t1)
	for i = 1; i < 20; i++ {
		fe_sq_tt(&t2, &t2)
	}
	Fe_mul_ttt(&t1, &t2, &t1)
	fe_sq_tt(&t1, &t1)
	for i = 1; i < 10; i++ {
		fe_sq_tt(&t1, &t1)
	}
	Fe_mul_ttt(&t0, &t1, &t0)
	fe_sq_tt(&t1, &t0)
	for i = 1; i < 50; i++ {
		fe_sq_tt(&t1, &t1)
	}
	Fe_mul_ttt(&t1, &t1, &t0)
	fe_sq_tt(&t2, &t1)
	for i = 1; i < 100; i++ {
		fe_sq_tt(&t2, &t2)
	}
	Fe_mul_ttt(&t1, &t2, &t1)
	fe_sq_tt(&t1, &t1)
	for i = 1; i < 50; i++ {
		fe_sq_tt(&t1, &t1)
	}
	Fe_mul_ttt(&t0, &t1, &t0)
	fe_sq_tt(&t0, &t0)
	for i = 1; i < 2; i++ {
		fe_sq_tt(&t0, &t0)
	}
	Fe_mul_ttt(out, &t0, z)
}

type Ge_p2 struct {
	X, Y, Z Fe
}

type Ge_p3 struct {
	X, Y, Z, T Fe
}

type Ge_p1p1 struct {
	X, Y, Z, T Fe_loose
}

type Ge_precomp struct {
	Yplusx, Yminusx, Xy2d Fe_loose
}

type Ge_cached struct {
	YplusX, YminusX, Z, T2d Fe_loose
}

func Ge_p2_0(h *Ge_p2) {
	fe_0(&h.X)
	fe_1(&h.Y)
	fe_1(&h.Z)
}

func Ge_p3_0(h *Ge_p3) {
	fe_0(&h.X)
	fe_1(&h.Y)
	fe_1(&h.Z)
	fe_0(&h.T)
}

func Ge_cached_0(h *Ge_cached) {
	fe_loose_1(&h.YplusX)
	fe_loose_1(&h.YminusX)
	fe_loose_1(&h.Z)
	fe_loose_0(&h.T2d)
}

func Ge_precomp_0(h *Ge_precomp) {
	fe_loose_1(&h.Yplusx)
	fe_loose_1(&h.Yminusx)
	fe_loose_0(&h.Xy2d)
}

func Ge_p3_to_p2(r *Ge_p2, p *Ge_p3) {
	fe_copy(&r.X, &p.X)
	fe_copy(&r.Y, &p.Y)
	fe_copy(&r.Z, &p.Z)
}

func X25519_ge_p3_to_cached(r *Ge_cached, p *Ge_p3) {
	Fe_add(&r.YplusX, &p.Y, &p.X)
	Fe_sub(&r.YminusX, &p.Y, &p.X)
	fe_copy_lt(&r.Z, &p.Z)
	Fe_mul_ltt(&r.T2d, &p.T, &D2)
}

func X25519_ge_p1p1_to_p2(r *Ge_p2, p *Ge_p1p1) {
	fe_mul_tll(&r.X, &p.X, &p.T)
	fe_mul_tll(&r.Y, &p.Y, &p.Z)
	fe_mul_tll(&r.Z, &p.Z, &p.T)
}

func X25519_ge_p1p1_to_p3(r *Ge_p3, p *Ge_p1p1) {
	fe_mul_tll(&r.X, &p.X, &p.T)
	fe_mul_tll(&r.Y, &p.Y, &p.Z)
	fe_mul_tll(&r.Z, &p.Z, &p.T)
	fe_mul_tll(&r.T, &p.X, &p.Y)
}

func Ge_p1p1_to_cached(r *Ge_cached, p *Ge_p1p1) {
	var t Ge_p3
	X25519_ge_p1p1_to_p3(&t, p)
	X25519_ge_p3_to_cached(r, &t)
}

func Ge_p2_dbl(r *Ge_p1p1, p *Ge_p2) {
	var trX, trZ, trT Fe
	var t0 Fe

	fe_sq_tt(&trX, &p.X)
	fe_sq_tt(&trZ, &p.Y)
	fe_sq2_tt(&trT, &p.Z)
	Fe_add(&r.Y, &p.X, &p.Y)
	fe_sq_tl(&t0, &r.Y)

	Fe_add(&r.Y, &trZ, &trX)
	Fe_sub(&r.Z, &trZ, &trX)
	Fe_carry(&trZ, &r.Y)
	Fe_sub(&r.X, &t0, &trZ)
	Fe_carry(&trZ, &r.Z)
	Fe_sub(&r.T, &trT, &trZ)
}

func X25519_ge_tobytes(s *[]byte, h *Ge_p2) {
	var recip, x, y Fe

	fe_invert(&recip, &h.Z)

	Fe_mul_ttt(&x, &h.X, &recip)
	Fe_mul_ttt(&y, &h.Y, &recip)

	Fe_tobytes(s, &y)
	(*s)[31] ^= byte(fe_isnegative(&x) << 7)
}

func ge_p3_tobytes(s *[]byte, h *Ge_p3) {
	var recip, x, y Fe

	fe_invert(&recip, &h.Z)
	Fe_mul_ttt(&x, &h.X, &recip)
	Fe_mul_ttt(&y, &h.Y, &recip)
	Fe_tobytes(s, &y)
	(*s)[31] ^= byte(fe_isnegative(&x) << 7)
}

func Ge_madd(r *Ge_p1p1, p *Ge_p3, q *Ge_precomp) {
	var trY, trZ, trT Fe

	Fe_add(&r.X, &p.Y, &p.X)
	Fe_sub(&r.Y, &p.Y, &p.X)
	fe_mul_tll(&trZ, &r.X, &q.Yplusx)
	fe_mul_tll(&trY, &r.Y, &q.Yminusx)
	Fe_mul_tlt(&trT, &q.Xy2d, &p.T)
	Fe_add(&r.T, &p.Z, &p.Z)
	Fe_sub(&r.X, &trZ, &trY)
	Fe_add(&r.Y, &trZ, &trY)
	Fe_carry(&trZ, &r.T)
	Fe_add(&r.Z, &trZ, &trT)
	Fe_sub(&r.T, &trZ, &trT)
}

func ge_msub(r *Ge_p1p1, p *Ge_p3, q *Ge_precomp) {
	var trY, trZ, trT Fe

	Fe_add(&r.X, &p.Y, &p.X)
	Fe_sub(&r.Y, &p.Y, &p.X)
	fe_mul_tll(&trZ, &r.X, &q.Yminusx)
	fe_mul_tll(&trY, &r.Y, &q.Yplusx)
	Fe_mul_tlt(&trT, &q.Xy2d, &p.T)
	Fe_add(&r.T, &p.Z, &p.Z)
	Fe_sub(&r.X, &trZ, &trY)
	Fe_add(&r.Y, &trZ, &trY)
	Fe_carry(&trZ, &r.T)
	Fe_sub(&r.Z, &trZ, &trT)
	Fe_add(&r.T, &trZ, &trT)
}

func X25519_ge_add(r *Ge_p1p1, p *Ge_p3, q *Ge_cached) {
	var trX, trY, trZ, trT Fe

	Fe_add(&r.X, &p.Y, &p.X)
	Fe_sub(&r.Y, &p.Y, &p.X)
	fe_mul_tll(&trZ, &r.X, &q.YplusX)
	fe_mul_tll(&trY, &r.Y, &q.YminusX)
	Fe_mul_tlt(&trT, &q.T2d, &p.T)
	fe_mul_ttl(&trX, &p.Z, &q.Z)
	Fe_add(&r.T, &trX, &trX)
	Fe_sub(&r.X, &trZ, &trY)
	Fe_add(&r.Y, &trZ, &trY)
	Fe_carry(&trZ, &r.T)
	Fe_add(&r.Z, &trZ, &trT)
	Fe_sub(&r.T, &trZ, &trT)
}

func X25519_ge_sub(r *Ge_p1p1, p *Ge_p3, q *Ge_cached) {
	var trX, trY, trZ, trT Fe

	Fe_add(&r.X, &p.Y, &p.X)
	Fe_sub(&r.Y, &p.Y, &p.X)
	fe_mul_tll(&trZ, &r.X, &q.YminusX)
	fe_mul_tll(&trY, &r.Y, &q.YplusX)
	Fe_mul_tlt(&trT, &q.T2d, &p.T)
	fe_mul_ttl(&trX, &p.Z, &q.Z)
	Fe_add(&r.T, &trX, &trX)
	Fe_sub(&r.X, &trZ, &trY)
	Fe_add(&r.Y, &trZ, &trY)
	Fe_carry(&trZ, &r.T)
	Fe_sub(&r.Z, &trZ, &trT)
	Fe_add(&r.T, &trZ, &trT)
}

func X25519_ge_frombytes_vartime(h *Ge_p3, s []byte) int {
	var u Fe
	var v Fe_loose
	var v3 Fe
	var vxx Fe
	var check Fe_loose

	Fe_frombytes(&h.Y, s)
	fe_1(&h.Z)
	fe_sq_tt(&v3, &h.Y)
	Fe_mul_ttt(&vxx, &v3, &d)
	Fe_sub(&v, &v3, &h.Z) // u = y^2-1
	Fe_carry(&u, &v)
	Fe_add(&v, &vxx, &h.Z) // v = dy^2+1
	fe_sq_tl(&v3, &v)
	fe_mul_ttl(&v3, &v3, &v) // v3 = v^3
	fe_sq_tt(&h.X, &v3)
	fe_mul_ttl(&h.X, &h.X, &v)
	Fe_mul_ttt(&h.X, &h.X, &u) // x = uv^7

	fe_pow22523(&h.X, &h.X) // x = (uv^7)^((q-5)/8)
	Fe_mul_ttt(&h.X, &h.X, &v3)
	Fe_mul_ttt(&h.X, &h.X, &u) // x = uv^3(uv^7)^((q-5)/8)

	fe_sq_tt(&vxx, &h.X)
	fe_mul_ttl(&vxx, &vxx, &v)
	Fe_sub(&check, &vxx, &u)

	if fe_isnonzero(&check) {
		Fe_add(&check, &vxx, &u)
		if fe_isnonzero(&check) {
			return 0
		}
		Fe_mul_ttt(&h.X, &h.X, &sqrtm1)
	}

	if fe_isnegative(&h.X) != int(s[31]>>7) {
		var t Fe_loose
		fe_neg(&t, &h.X)
		Fe_carry(&h.X, &t)
	}

	Fe_mul_ttt(&h.T, &h.X, &h.Y)
	return 1
}
