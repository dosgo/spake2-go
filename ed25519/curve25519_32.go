package ed25519

import (
	"math/big"
)

type fiat25519Uint1 uint8
type fiat25519Int1 int8

type Fiat25519FieldElement [10]uint32

func fiat25519ValueBarrierU32(a uint32) uint32 {
	return a // No inline assembly equivalent in Go
}

func fiat25519AddCarryXU26(out1 *uint32, out2 *fiat25519Uint1, arg1 fiat25519Uint1, arg2 uint32, arg3 uint32) {
	x1 := uint32(arg1) + arg2 + arg3
	*out1 = x1 & 0x3ffffff
	*out2 = fiat25519Uint1(x1 >> 26)
}

func fiat25519SubBorrowXU26(out1 *uint32, out2 *fiat25519Uint1, arg1 fiat25519Uint1, arg2 uint32, arg3 uint32) {
	x1 := int32(arg2) - int32(arg1) - int32(arg3)
	*out1 = uint32(x1 & 0x3ffffff)
	*out2 = fiat25519Uint1(0 - (x1 >> 26))
}

func fiat25519AddCarryXU25(out1 *uint32, out2 *fiat25519Uint1, arg1 fiat25519Uint1, arg2 uint32, arg3 uint32) {
	x1 := uint32(arg1) + arg2 + arg3
	*out1 = x1 & 0x1ffffff
	*out2 = fiat25519Uint1(x1 >> 25)
}

func fiat25519SubBorrowXU25(out1 *uint32, out2 *fiat25519Uint1, arg1 fiat25519Uint1, arg2 uint32, arg3 uint32) {
	var x1 int32
	var x2 fiat25519Int1
	var x3 uint32
	x1 = (int32(arg2-uint32(arg1)) - int32(arg3))
	x2 = (fiat25519Int1)(x1 >> 25)
	x3 = uint32(uint32(x1) & uint32(0x1ffffff))
	*out1 = x3
	*out2 = (fiat25519Uint1)(0x0 - x2)
}

func fiat25519CmovznzU32(out1 *uint32, arg1 fiat25519Uint1, arg2 uint32, arg3 uint32) {
	x1 := fiat25519Uint1(0)
	if arg1 != 0 {
		x1 = 1
	}
	x2 := uint32(0 - x1)
	*out1 = (fiat25519ValueBarrierU32(x2) & arg3) | (fiat25519ValueBarrierU32(^x2) & arg2)
}

func fiat25519CarryMul(out1 Fiat25519FieldElement, arg1, arg2 Fiat25519FieldElement) Fiat25519FieldElement {
	var x1, x2, x3, x4, x5, x6, x7, x8, x9, x10, x11, x12, x13, x14, x15, x16, x17, x18, x19, x20, x21, x22, x23, x24, x25, x26, x27, x28, x29, x30, x31, x32, x33, x34, x35, x36, x37, x38, x39, x40, x41, x42, x43, x44, x45, x46, x47, x48, x49, x50, x51, x52, x53, x54, x55, x56, x57, x58, x59, x60, x61, x62, x63, x64, x65, x66, x67, x68, x69, x70, x71, x72, x73, x74, x75, x76, x77, x78, x79, x80, x81, x82, x83, x84, x85, x86, x87, x88, x89, x90, x91, x92, x93, x94, x95, x96, x97, x98, x99, x100, x101, x102, x103, x104, x105, x106, x107, x108, x109, x110, x111, x112, x113, x114, x115, x116, x117, x118, x119, x120, x121, x122, x123, x124, x125, x126, x127, x128, x129, x130, x131, x132, x133, x134, x135, x136, x137, x138, x139, x140, x141, x142, x143, x144, x145, x146, x147 big.Int

	// 计算过程
	x1.Mul(big.NewInt(int64(arg1[9])), new(big.Int).Mul(big.NewInt(int64(arg2[9])), big.NewInt(0x26)))
	x2.Mul(big.NewInt(int64(arg1[9])), new(big.Int).Mul(big.NewInt(int64(arg2[8])), big.NewInt(0x13)))
	x3.Mul(big.NewInt(int64(arg1[9])), new(big.Int).Mul(big.NewInt(int64(arg2[7])), big.NewInt(0x26)))
	x4.Mul(big.NewInt(int64(arg1[9])), new(big.Int).Mul(big.NewInt(int64(arg2[6])), big.NewInt(0x13)))
	x5.Mul(big.NewInt(int64(arg1[9])), new(big.Int).Mul(big.NewInt(int64(arg2[5])), big.NewInt(0x26)))
	x6.Mul(big.NewInt(int64(arg1[9])), new(big.Int).Mul(big.NewInt(int64(arg2[4])), big.NewInt(0x13)))
	x7.Mul(big.NewInt(int64(arg1[9])), new(big.Int).Mul(big.NewInt(int64(arg2[3])), big.NewInt(0x26)))
	x8.Mul(big.NewInt(int64(arg1[9])), new(big.Int).Mul(big.NewInt(int64(arg2[2])), big.NewInt(0x13)))
	x9.Mul(big.NewInt(int64(arg1[9])), new(big.Int).Mul(big.NewInt(int64(arg2[1])), big.NewInt(0x26)))
	x10.Mul(big.NewInt(int64(arg1[8])), new(big.Int).Mul(big.NewInt(int64(arg2[9])), big.NewInt(0x13)))
	x11.Mul(big.NewInt(int64(arg1[8])), new(big.Int).Mul(big.NewInt(int64(arg2[8])), big.NewInt(0x13)))
	x12.Mul(big.NewInt(int64(arg1[8])), new(big.Int).Mul(big.NewInt(int64(arg2[7])), big.NewInt(0x13)))
	x13.Mul(big.NewInt(int64(arg1[8])), new(big.Int).Mul(big.NewInt(int64(arg2[6])), big.NewInt(0x13)))
	x14.Mul(big.NewInt(int64(arg1[8])), new(big.Int).Mul(big.NewInt(int64(arg2[5])), big.NewInt(0x13)))
	x15.Mul(big.NewInt(int64(arg1[8])), new(big.Int).Mul(big.NewInt(int64(arg2[4])), big.NewInt(0x13)))
	x16.Mul(big.NewInt(int64(arg1[8])), new(big.Int).Mul(big.NewInt(int64(arg2[3])), big.NewInt(0x13)))
	x17.Mul(big.NewInt(int64(arg1[8])), new(big.Int).Mul(big.NewInt(int64(arg2[2])), big.NewInt(0x13)))
	x18.Mul(big.NewInt(int64(arg1[7])), new(big.Int).Mul(big.NewInt(int64(arg2[9])), big.NewInt(0x26)))
	x19.Mul(big.NewInt(int64(arg1[7])), new(big.Int).Mul(big.NewInt(int64(arg2[8])), big.NewInt(0x13)))
	x20.Mul(big.NewInt(int64(arg1[7])), new(big.Int).Mul(big.NewInt(int64(arg2[7])), big.NewInt(0x26)))
	x21.Mul(big.NewInt(int64(arg1[7])), new(big.Int).Mul(big.NewInt(int64(arg2[6])), big.NewInt(0x13)))
	x22.Mul(big.NewInt(int64(arg1[7])), new(big.Int).Mul(big.NewInt(int64(arg2[5])), big.NewInt(0x26)))
	x23.Mul(big.NewInt(int64(arg1[7])), new(big.Int).Mul(big.NewInt(int64(arg2[4])), big.NewInt(0x13)))
	x24.Mul(big.NewInt(int64(arg1[7])), new(big.Int).Mul(big.NewInt(int64(arg2[3])), big.NewInt(0x26)))
	x25.Mul(big.NewInt(int64(arg1[6])), new(big.Int).Mul(big.NewInt(int64(arg2[9])), big.NewInt(0x13)))
	x26.Mul(big.NewInt(int64(arg1[6])), new(big.Int).Mul(big.NewInt(int64(arg2[8])), big.NewInt(0x13)))
	x27.Mul(big.NewInt(int64(arg1[6])), new(big.Int).Mul(big.NewInt(int64(arg2[7])), big.NewInt(0x13)))
	x28.Mul(big.NewInt(int64(arg1[6])), new(big.Int).Mul(big.NewInt(int64(arg2[6])), big.NewInt(0x13)))
	x29.Mul(big.NewInt(int64(arg1[6])), new(big.Int).Mul(big.NewInt(int64(arg2[5])), big.NewInt(0x13)))
	x30.Mul(big.NewInt(int64(arg1[6])), new(big.Int).Mul(big.NewInt(int64(arg2[4])), big.NewInt(0x13)))
	x31.Mul(big.NewInt(int64(arg1[5])), new(big.Int).Mul(big.NewInt(int64(arg2[9])), big.NewInt(0x26)))
	x32.Mul(big.NewInt(int64(arg1[5])), new(big.Int).Mul(big.NewInt(int64(arg2[8])), big.NewInt(0x13)))
	x33.Mul(big.NewInt(int64(arg1[5])), new(big.Int).Mul(big.NewInt(int64(arg2[7])), big.NewInt(0x26)))
	x34.Mul(big.NewInt(int64(arg1[5])), new(big.Int).Mul(big.NewInt(int64(arg2[6])), big.NewInt(0x13)))
	x35.Mul(big.NewInt(int64(arg1[5])), new(big.Int).Mul(big.NewInt(int64(arg2[5])), big.NewInt(0x26)))
	x36.Mul(big.NewInt(int64(arg1[4])), new(big.Int).Mul(big.NewInt(int64(arg2[9])), big.NewInt(0x13)))
	x37.Mul(big.NewInt(int64(arg1[4])), new(big.Int).Mul(big.NewInt(int64(arg2[8])), big.NewInt(0x13)))
	x38.Mul(big.NewInt(int64(arg1[4])), new(big.Int).Mul(big.NewInt(int64(arg2[7])), big.NewInt(0x13)))
	x39.Mul(big.NewInt(int64(arg1[4])), new(big.Int).Mul(big.NewInt(int64(arg2[6])), big.NewInt(0x13)))
	x40.Mul(big.NewInt(int64(arg1[3])), new(big.Int).Mul(big.NewInt(int64(arg2[9])), big.NewInt(0x26)))
	x41.Mul(big.NewInt(int64(arg1[3])), new(big.Int).Mul(big.NewInt(int64(arg2[8])), big.NewInt(0x13)))
	x42.Mul(big.NewInt(int64(arg1[3])), new(big.Int).Mul(big.NewInt(int64(arg2[7])), big.NewInt(0x26)))
	x43.Mul(big.NewInt(int64(arg1[2])), new(big.Int).Mul(big.NewInt(int64(arg2[9])), big.NewInt(0x13)))
	x44.Mul(big.NewInt(int64(arg1[2])), new(big.Int).Mul(big.NewInt(int64(arg2[8])), big.NewInt(0x13)))
	x45.Mul(big.NewInt(int64(arg1[1])), new(big.Int).Mul(big.NewInt(int64(arg2[9])), big.NewInt(0x26)))
	x46.Mul(big.NewInt(int64(arg1[9])), big.NewInt(int64(arg2[0])))
	x47.Mul(big.NewInt(int64(arg1[8])), big.NewInt(int64(arg2[1])))
	x48.Mul(big.NewInt(int64(arg1[8])), big.NewInt(int64(arg2[0])))
	x49.Mul(big.NewInt(int64(arg1[7])), big.NewInt(int64(arg2[2])))
	x50.Mul(big.NewInt(int64(arg1[7])), new(big.Int).Mul(big.NewInt(int64(arg2[1])), big.NewInt(0x2)))
	x51.Mul(big.NewInt(int64(arg1[7])), big.NewInt(int64(arg2[0])))
	x52.Mul(big.NewInt(int64(arg1[6])), big.NewInt(int64(arg2[3])))
	x53.Mul(big.NewInt(int64(arg1[6])), big.NewInt(int64(arg2[2])))
	x54.Mul(big.NewInt(int64(arg1[6])), big.NewInt(int64(arg2[1])))
	x55.Mul(big.NewInt(int64(arg1[6])), big.NewInt(int64(arg2[0])))
	x56.Mul(big.NewInt(int64(arg1[5])), big.NewInt(int64(arg2[4])))
	x57.Mul(big.NewInt(int64(arg1[5])), new(big.Int).Mul(big.NewInt(int64(arg2[3])), big.NewInt(0x2)))
	x58.Mul(big.NewInt(int64(arg1[5])), big.NewInt(int64(arg2[2])))
	x59.Mul(big.NewInt(int64(arg1[5])), new(big.Int).Mul(big.NewInt(int64(arg2[1])), big.NewInt(0x2)))
	x60.Mul(big.NewInt(int64(arg1[5])), big.NewInt(int64(arg2[0])))
	x61.Mul(big.NewInt(int64(arg1[4])), big.NewInt(int64(arg2[5])))
	x62.Mul(big.NewInt(int64(arg1[4])), big.NewInt(int64(arg2[4])))
	x63.Mul(big.NewInt(int64(arg1[4])), big.NewInt(int64(arg2[3])))
	x64.Mul(big.NewInt(int64(arg1[4])), big.NewInt(int64(arg2[2])))
	x65.Mul(big.NewInt(int64(arg1[4])), big.NewInt(int64(arg2[1])))
	x66.Mul(big.NewInt(int64(arg1[4])), big.NewInt(int64(arg2[0])))
	x67.Mul(big.NewInt(int64(arg1[3])), big.NewInt(int64(arg2[6])))
	x68.Mul(big.NewInt(int64(arg1[3])), new(big.Int).Mul(big.NewInt(int64(arg2[5])), big.NewInt(0x2)))
	x69.Mul(big.NewInt(int64(arg1[3])), big.NewInt(int64(arg2[4])))
	x70.Mul(big.NewInt(int64(arg1[3])), new(big.Int).Mul(big.NewInt(int64(arg2[3])), big.NewInt(0x2)))
	x71.Mul(big.NewInt(int64(arg1[3])), big.NewInt(int64(arg2[2])))
	x72.Mul(big.NewInt(int64(arg1[3])), new(big.Int).Mul(big.NewInt(int64(arg2[1])), big.NewInt(0x2)))
	x73.Mul(big.NewInt(int64(arg1[3])), big.NewInt(int64(arg2[0])))
	x74.Mul(big.NewInt(int64(arg1[2])), big.NewInt(int64(arg2[7])))
	x75.Mul(big.NewInt(int64(arg1[2])), big.NewInt(int64(arg2[6])))
	x76.Mul(big.NewInt(int64(arg1[2])), big.NewInt(int64(arg2[5])))
	x77.Mul(big.NewInt(int64(arg1[2])), big.NewInt(int64(arg2[4])))
	x78.Mul(big.NewInt(int64(arg1[2])), big.NewInt(int64(arg2[3])))
	x79.Mul(big.NewInt(int64(arg1[2])), big.NewInt(int64(arg2[2])))
	x80.Mul(big.NewInt(int64(arg1[2])), big.NewInt(int64(arg2[1])))
	x81.Mul(big.NewInt(int64(arg1[2])), big.NewInt(int64(arg2[0])))
	x82.Mul(big.NewInt(int64(arg1[1])), big.NewInt(int64(arg2[8])))
	x83.Mul(big.NewInt(int64(arg1[1])), new(big.Int).Mul(big.NewInt(int64(arg2[7])), big.NewInt(0x2)))
	x84.Mul(big.NewInt(int64(arg1[1])), big.NewInt(int64(arg2[6])))
	x85.Mul(big.NewInt(int64(arg1[1])), new(big.Int).Mul(big.NewInt(int64(arg2[5])), big.NewInt(0x2)))
	x86.Mul(big.NewInt(int64(arg1[1])), big.NewInt(int64(arg2[4])))
	x87.Mul(big.NewInt(int64(arg1[1])), new(big.Int).Mul(big.NewInt(int64(arg2[3])), big.NewInt(0x2)))
	x88.Mul(big.NewInt(int64(arg1[1])), big.NewInt(int64(arg2[2])))
	x89.Mul(big.NewInt(int64(arg1[1])), new(big.Int).Mul(big.NewInt(int64(arg2[1])), big.NewInt(0x2)))
	x90.Mul(big.NewInt(int64(arg1[1])), big.NewInt(int64(arg2[0])))
	x91.Mul(big.NewInt(int64(arg1[0])), big.NewInt(int64(arg2[9])))
	x92.Mul(big.NewInt(int64(arg1[0])), big.NewInt(int64(arg2[8])))
	x93.Mul(big.NewInt(int64(arg1[0])), big.NewInt(int64(arg2[7])))
	x94.Mul(big.NewInt(int64(arg1[0])), big.NewInt(int64(arg2[6])))
	x95.Mul(big.NewInt(int64(arg1[0])), big.NewInt(int64(arg2[5])))
	x96.Mul(big.NewInt(int64(arg1[0])), big.NewInt(int64(arg2[4])))
	x97.Mul(big.NewInt(int64(arg1[0])), big.NewInt(int64(arg2[3])))
	x98.Mul(big.NewInt(int64(arg1[0])), big.NewInt(int64(arg2[2])))
	x99.Mul(big.NewInt(int64(arg1[0])), big.NewInt(int64(arg2[1])))
	x100.Mul(big.NewInt(int64(arg1[0])), big.NewInt(int64(arg2[0])))

	temp := big.NewInt(0)
	temp.Add(temp, &x9)
	temp.Add(temp, &x17)
	temp.Add(temp, &x24)
	temp.Add(temp, &x30)
	temp.Add(temp, &x35)
	temp.Add(temp, &x39)
	temp.Add(temp, &x42)
	temp.Add(temp, &x44)
	temp.Add(temp, &x45)
	temp.Add(temp, &x100)

	x101.Set(temp)

	var x101Copy big.Int
	x101Copy.Set(&x101)
	x102.Set(x101Copy.Rsh(&x101Copy, 26))

	x103.And(&x101, big.NewInt(0x3ffffff))

	tempX104 := big.NewInt(0).Add(&x91, big.NewInt(0).Add(&x82, big.NewInt(0).Add(&x74, big.NewInt(0).Add(&x67, big.NewInt(0).Add(&x61, big.NewInt(0).Add(&x56, big.NewInt(0).Add(&x52, big.NewInt(0).Add(&x49, big.NewInt(0).Add(&x47, &x46)))))))))
	x104.Set(tempX104)

	// x105 = (x92 + (x83 + (x75 + (x68 + (x62 + (x57 + (x53 + (x50 + (x48 + x1)))))))))
	tempX105 := big.NewInt(0).Add(&x92, big.NewInt(0).Add(&x83, big.NewInt(0).Add(&x75, big.NewInt(0).Add(&x68, big.NewInt(0).Add(&x62, big.NewInt(0).Add(&x57, big.NewInt(0).Add(&x53, big.NewInt(0).Add(&x50, big.NewInt(0).Add(&x48, &x1)))))))))
	x105.Set(tempX105)

	// x106 = (x93 + (x84 + (x76 + (x69 + (x63 + (x58 + (x54 + (x51 + (x10 + x2)))))))))
	tempX106 := big.NewInt(0).Add(&x93, big.NewInt(0).Add(&x84, big.NewInt(0).Add(&x76, big.NewInt(0).Add(&x69, big.NewInt(0).Add(&x63, big.NewInt(0).Add(&x58, big.NewInt(0).Add(&x54, big.NewInt(0).Add(&x51, big.NewInt(0).Add(&x10, &x2)))))))))
	x106.Set(tempX106)

	// x107 = (x94 + (x85 + (x77 + (x70 + (x64 + (x59 + (x55 + (x18 + (x11 + x3)))))))))
	tempX107 := big.NewInt(0).Add(&x94, big.NewInt(0).Add(&x85, big.NewInt(0).Add(&x77, big.NewInt(0).Add(&x70, big.NewInt(0).Add(&x64, big.NewInt(0).Add(&x59, big.NewInt(0).Add(&x55, big.NewInt(0).Add(&x18, big.NewInt(0).Add(&x11, &x3)))))))))
	x107.Set(tempX107)

	// x108 = (x95 + (x86 + (x78 + (x71 + (x65 + (x60 + (x25 + (x19 + (x12 + x4)))))))))
	tempX108 := big.NewInt(0).Add(&x95, big.NewInt(0).Add(&x86, big.NewInt(0).Add(&x78, big.NewInt(0).Add(&x71, big.NewInt(0).Add(&x65, big.NewInt(0).Add(&x60, big.NewInt(0).Add(&x25, big.NewInt(0).Add(&x19, big.NewInt(0).Add(&x12, &x4)))))))))
	x108.Set(tempX108)

	// x109 = (x96 + (x87 + (x79 + (x72 + (x66 + (x31 + (x26 + (x20 + (x13 + x5)))))))))
	tempX109 := big.NewInt(0).Add(&x96, big.NewInt(0).Add(&x87, big.NewInt(0).Add(&x79, big.NewInt(0).Add(&x72, big.NewInt(0).Add(&x66, big.NewInt(0).Add(&x31, big.NewInt(0).Add(&x26, big.NewInt(0).Add(&x20, big.NewInt(0).Add(&x13, &x5)))))))))
	x109.Set(tempX109)

	// x110 = (x97 + (x88 + (x80 + (x73 + (x36 + (x32 + (x27 + (x21 + (x14 + x6)))))))))
	tempX110 := big.NewInt(0).Add(&x97, big.NewInt(0).Add(&x88, big.NewInt(0).Add(&x80, big.NewInt(0).Add(&x73, big.NewInt(0).Add(&x36, big.NewInt(0).Add(&x32, big.NewInt(0).Add(&x27, big.NewInt(0).Add(&x21, big.NewInt(0).Add(&x14, &x6)))))))))
	x110.Set(tempX110)

	// x111 = (x98 + (x89 + (x81 + (x40 + (x37 + (x33 + (x28 + (x22 + (x15 + x7)))))))))
	tempX111 := big.NewInt(0).Add(&x98, big.NewInt(0).Add(&x89, big.NewInt(0).Add(&x81, big.NewInt(0).Add(&x40, big.NewInt(0).Add(&x37, big.NewInt(0).Add(&x33, big.NewInt(0).Add(&x28, big.NewInt(0).Add(&x22, big.NewInt(0).Add(&x15, &x7)))))))))
	x111.Set(tempX111)

	// x112 = (x99 + (x90 + (x43 + (x41 + (x38 + (x34 + (x29 + (x23 + (x16 + x8)))))))))
	tempX112 := big.NewInt(0).Add(&x99, big.NewInt(0).Add(&x90, big.NewInt(0).Add(&x43, big.NewInt(0).Add(&x41, big.NewInt(0).Add(&x38, big.NewInt(0).Add(&x34, big.NewInt(0).Add(&x29, big.NewInt(0).Add(&x23, big.NewInt(0).Add(&x16, &x8)))))))))
	x112.Set(tempX112)

	x113.Add(&x102, &x112)

	var x113Copy big.Int
	x113Copy.Set(&x113)

	// x114 = (x113 >> 25)
	x114.Set(x113Copy.Rsh(&x113Copy, 25))

	// x115 = uint32(uint32(x113) & uint32(0x1ffffff))
	x115.And(&x113, big.NewInt(0x1ffffff))

	// x116 = (x114 + x111)
	x116.Add(&x114, &x111)

	// x117 = (x116 >> 26)
	var x116Copy big.Int
	x116Copy.Set(&x116)

	x117.Set(x116Copy.Rsh(&x116Copy, 26))

	// x118 = uint32(uint32(x116) & uint32(0x3ffffff))
	x118.And(&x116, big.NewInt(0x3ffffff))

	// x119 = (x117 + x110)
	x119.Add(&x117, &x110)

	// x120 = (x119 >> 25)
	var x119Copy big.Int
	x119Copy.Set(&x119)

	x120.Set(x119Copy.Rsh(&x119Copy, 25))

	// x121 = uint32(uint32(x119) & uint32(0x1ffffff))
	x121.And(&x119, big.NewInt(0x1ffffff))
	// x122 = (x120 + x109)
	x122.Add(&x120, &x109)

	// x123 = (x122 >> 26)
	var x122Copy big.Int
	x122Copy.Set(&x122)
	x123.Set(x122Copy.Rsh(&x122Copy, 26))

	// x124 = uint32(uint32(x122) & uint32(0x3ffffff))
	x124.And(&x122, big.NewInt(0x3ffffff))
	// x125 = (x123 + x108)
	x125.Add(&x123, &x108)

	// x126 = (x125 >> 25)
	var x125Copy big.Int
	x125Copy.Set(&x125)
	x126.Set(x125Copy.Rsh(&x125Copy, 25))

	// x127 = uint32(uint32(x125) & uint32(0x1ffffff))
	x127.And(&x125, big.NewInt(0x1ffffff))
	// x128 = (x126 + x107)
	x128.Add(&x126, &x107)

	// x129 = (x128 >> 26)

	var x128Copy big.Int
	x128Copy.Set(&x128)
	x129.Set(x128Copy.Rsh(&x128Copy, 26))

	// x130 = uint32(uint32(x128) & uint32(0x3ffffff))

	x130.And(&x128, big.NewInt(0x3ffffff))
	// x131 = (x129 + x106)
	x131.Add(&x129, &x106)

	// x132 = (x131 >> 25)
	var x131Copy big.Int
	x131Copy.Set(&x131)
	x132.Set(x131Copy.Rsh(&x131Copy, 25))

	// x133 = uint32(uint32(x131) & uint32(0x1ffffff))
	x133.And(&x131, big.NewInt(0x1ffffff))

	// x134 = (x132 + x105)
	x134.Add(&x132, &x105)

	// x135 = (x134 >> 26)
	var x134Copy big.Int
	x134Copy.Set(&x134)
	x135.Set(x134Copy.Rsh(&x134Copy, 26))

	// x136 = uint32(uint32(x134) & uint32(0x3ffffff))
	x136.And(&x134, big.NewInt(0x3ffffff))
	// x137 = (x135 + x104)
	x137.Add(&x135, &x104)

	// x138 = (x137 >> 25)
	var x137Copy big.Int
	x137Copy.Set(&x137)
	x138.Set(x137Copy.Rsh(&x137Copy, 25))

	// x139 = uint32(uint32(x137) & uint32(0x1ffffff))
	x139.And(&x137, big.NewInt(0x1ffffff))
	// x140 = (x138 * (0x13))
	x140.Mul(&x138, big.NewInt(0x13))

	// x141 = uint64(x103 + uint32(x140))
	temp141 := big.NewInt(0).Add(&x103, &x140)
	x141.Set(temp141)

	// x142 = uint32(x141 >> 26)
	var x141Copy big.Int
	x141Copy.Set(&x141)
	x142.Set(x141Copy.Rsh(&x141Copy, 26))

	// x143 = uint32(x141 & 0x3ffffff)
	x143.And(&x141, big.NewInt(0x3ffffff))
	// x144 = (x142 + x115)
	x144.Add(&x142, &x115)

	// x145 = (fiat25519Uint1)(x144 >> 25)
	var x144Copy big.Int
	x144Copy.Set(&x144)
	x145.Set(x144Copy.Rsh(&x144Copy, 25))

	// x146 = (x144 & uint32(0x1ffffff))
	x146.And(&x144, big.NewInt(0x1ffffff))

	// x147 = (uint32(x145) + x118)
	x147.Add(&x145, &x118)

	out1[0] = uint32(x143.Uint64())
	out1[1] = uint32(x146.Uint64())
	out1[2] = uint32(x147.Uint64())
	out1[3] = uint32(x121.Uint64())
	out1[4] = uint32(x124.Uint64())
	out1[5] = uint32(x127.Uint64())
	out1[6] = uint32(x130.Uint64())
	out1[7] = uint32(x133.Uint64())
	out1[8] = uint32(x136.Uint64())
	out1[9] = uint32(x139.Uint64())

	return out1
}

func fiat25519CarrySquare(out1 Fiat25519FieldElement, arg1 *Fiat25519FieldElement) Fiat25519FieldElement {
	// 声明所有的局部变量
	var (
		x1, x2, x3, x4, x6, x7, x8, x9, x10, x12, x116, x13, x14, x15, x16, x17, x18, x76, x88, x91, x94, x97, x100, x103, x106, x109, x112, x115, x119, x120                                                                                                                                                                                                                                                                                                                 uint32
		x5, x11, x19, x20, x21, x22, x23, x24, x25, x26, x27, x28, x29, x30, x31, x32, x33, x34, x35, x36, x37, x38, x39, x40, x41, x42, x43, x44, x45, x46, x47, x48, x49, x50, x51, x52, x53, x54, x55, x56, x57, x58, x59, x60, x61, x62, x63, x64, x65, x66, x67, x68, x69, x70, x71, x72, x73, x74, x75, x77, x78, x79, x80, x81, x82, x83, x84, x85, x86, x87, x89, x90, x92, x93, x95, x96, x98, x99, x101, x102, x104, x105, x107, x108, x110, x111, x113, x114, x117 uint64
	)

	// 计算过程
	x1 = (arg1[9] * 0x13)
	x2 = (x1 * 0x2)
	x3 = (arg1[9] * 0x2)
	x4 = (arg1[8] * 0x13)
	x5 = (uint64(x4) * 0x2)
	x6 = (arg1[8] * 0x2)
	x7 = (arg1[7] * 0x13)
	x8 = (x7 * 0x2)
	x9 = (arg1[7] * 0x2)
	x10 = ((arg1[6]) * 0x13)
	x11 = (uint64(x10) * 0x2)
	x12 = (arg1[6] * 0x2)
	x13 = (arg1[5] * 0x13)
	x14 = (arg1[5] * 0x2)
	x15 = (arg1[4] * 0x2)
	x16 = (arg1[3] * 0x2)
	x17 = (arg1[2] * 0x2)
	x18 = (arg1[1] * 0x2)
	x19 = (uint64(arg1[9]) * (uint64(x1) * 0x2))
	x20 = (uint64(arg1[8]) * uint64(x2))
	x21 = (uint64(arg1[8]) * uint64(x4))
	x22 = (uint64(arg1[7]) * (uint64(x2) * 0x2))
	x23 = (uint64(arg1[7]) * x5)
	x24 = (uint64(arg1[7]) * (uint64(x7) * 0x2))
	x25 = (uint64(arg1[6]) * uint64(x2))
	x26 = (uint64(arg1[6]) * uint64(x5))
	x27 = (uint64(arg1[6]) * uint64(x8))
	x28 = (uint64(arg1[6]) * uint64(x10))
	x29 = (uint64(arg1[5]) * (uint64(x2) * 0x2))
	x30 = (uint64(arg1[5]) * x5)
	x31 = (uint64(arg1[5]) * (uint64(x8) * 0x2))
	x32 = (uint64(arg1[5]) * x11)
	x33 = (uint64(arg1[5]) * (uint64(x13) * 0x2))
	x34 = (uint64(arg1[4]) * uint64(x2))
	x35 = (uint64(arg1[4]) * uint64(x5))
	x36 = (uint64(arg1[4]) * uint64(x8))
	x37 = (uint64(arg1[4]) * x11)
	x38 = (uint64(arg1[4]) * uint64(x14))
	x39 = (uint64(arg1[4]) * uint64(arg1[4]))
	x40 = (uint64(arg1[3]) * (uint64(x2) * 0x2))
	x41 = (uint64(arg1[3]) * x5)
	x42 = (uint64(arg1[3]) * (uint64(x8) * 0x2))
	x43 = (uint64(arg1[3]) * uint64(x12))
	x44 = (uint64(arg1[3]) * (uint64(x14) * 0x2))
	x45 = (uint64(arg1[3]) * uint64(x15))
	x46 = (uint64(arg1[3]) * (uint64(arg1[3]) * 0x2))
	x47 = (uint64(arg1[2]) * uint64(x2))
	x48 = (uint64(arg1[2]) * x5)
	x49 = (uint64(arg1[2]) * uint64(x9))
	x50 = (uint64(arg1[2]) * uint64(x12))
	x51 = (uint64(arg1[2]) * uint64(x14))
	x52 = (uint64(arg1[2]) * uint64(x15))
	x53 = (uint64(arg1[2]) * uint64(x16))
	x54 = (uint64(arg1[2]) * uint64(arg1[2]))
	x55 = (uint64(arg1[1]) * (uint64(x2) * 0x2))
	x56 = (uint64(arg1[1]) * uint64(x6))
	x57 = (uint64(arg1[1]) * (uint64(x9) * 0x2))
	x58 = (uint64(arg1[1]) * uint64(x12))
	x59 = (uint64(arg1[1]) * (uint64(x14) * 0x2))
	x60 = (uint64(arg1[1]) * uint64(x15))
	x61 = (uint64(arg1[1]) * (uint64(x16) * 0x2))
	x62 = (uint64(arg1[1]) * uint64(x17))
	x63 = (uint64(arg1[1]) * (uint64(arg1[1]) * 0x2))
	x64 = (uint64(arg1[0]) * uint64(x3))
	x65 = (uint64(arg1[0]) * uint64(x6))
	x66 = (uint64(arg1[0]) * uint64(x9))
	x67 = (uint64(arg1[0]) * uint64(x12))
	x68 = (uint64(arg1[0]) * uint64(x14))
	x69 = (uint64(arg1[0]) * uint64(x15))
	x70 = (uint64(arg1[0]) * uint64(x16))
	x71 = (uint64(arg1[0]) * uint64(x17))
	x72 = (uint64(arg1[0]) * uint64(x18))
	x73 = (uint64(arg1[0]) * uint64(arg1[0]))
	x74 = (x73 + (x55 + (x48 + (x42 + (x37 + x33)))))
	x75 = (x74 >> 26)
	x76 = uint32(x74 & 0x3ffffff)
	x77 = (x64 + (x56 + (x49 + (x43 + x38))))
	x78 = (x65 + (x57 + (x50 + (x44 + (x39 + x19)))))
	x79 = (x66 + (x58 + (x51 + (x45 + x20))))
	x80 = (x67 + (x59 + (x52 + (x46 + (x22 + x21)))))
	x81 = (x68 + (x60 + (x53 + (x25 + x23))))
	x82 = (x69 + (x61 + (x54 + (x29 + (x26 + x24)))))
	x83 = (x70 + (x62 + (x34 + (x30 + x27))))
	x84 = (x71 + (x63 + (x40 + (x35 + (x31 + x28)))))
	x85 = (x72 + (x47 + (x41 + (x36 + x32))))
	x86 = (x75 + x85)
	x87 = (x86 >> 25)
	x88 = uint32(x86 & 0x1ffffff)
	x89 = (x87 + x84)
	x90 = (x89 >> 26)
	x91 = uint32(x89 & 0x3ffffff)
	x92 = (x90 + x83)
	x93 = (x92 >> 25)
	x94 = uint32(x92 & 0x1ffffff)
	x95 = (x93 + x82)
	x96 = (x95 >> 26)
	x97 = uint32(x95 & 0x3ffffff)
	x98 = (x96 + x81)
	x99 = (x98 >> 25)
	x100 = uint32(x98 & 0x1ffffff)
	x101 = (x99 + x80)
	x102 = (x101 >> 26)
	x103 = uint32(x101 & 0x3ffffff)
	x104 = (x102 + x79)
	x105 = (x104 >> 25)
	x106 = uint32(x104 & 0x1ffffff)
	x107 = (x105 + x78)
	x108 = (x107 >> 26)
	x109 = uint32(x107 & 0x3ffffff)
	x110 = (x108 + x77)
	x111 = (x110 >> 25)
	x112 = uint32(x110 & 0x1ffffff)
	x113 = (x111 * 0x13)
	x114 = (uint64(x76) + x113)
	x115 = uint32(x114 >> 26)
	x116 = uint32(x114 & 0x3ffffff)
	x117 = uint64((x115 + x88))
	x118 := (fiat25519Uint1)(x117 >> 25) // 假设 fiat_25519_uint1 是一个自定义类型，这里直接使用 uint64
	x119 = uint32(x117 & 0x1ffffff)
	x120 = (uint32(x118) + x91)
	out1[0] = x116
	out1[1] = x119
	out1[2] = x120
	out1[3] = x94
	out1[4] = x97
	out1[5] = x100
	out1[6] = x103
	out1[7] = x106
	out1[8] = x109
	out1[9] = x112
	return out1
}

/*
	func fiat25519CarrySquare(out1 *fiat25519TightFieldElement, arg1 *fiat25519LooseFieldElement) {
		var x [120]uint64
		x[1] = uint64(arg1[9]) * 0x13
		x[2] = x[1] * 2
		x[3] = uint64(arg1[9]) * 2
		x[4] = uint64(arg1[8]) * 0x13
		x[5] = x[4] * 2
		x[6] = uint64(arg1[8]) * 2
		x[7] = uint64(arg1[7]) * 0x13
		x[8] = x[7] * 2
		x[9] = uint64(arg1[7]) * 2
		x[10] = uint64(arg1[6]) * 0x13
		x[11] = x[10] * 2
		x[12] = uint64(arg1[6]) * 2
		x[13] = uint64(arg1[5]) * 0x13
		x[14] = uint64(arg1[5]) * 2
		x[15] = uint64(arg1[4]) * 2
		x[16] = uint64(arg1[3]) * 2
		x[17] = uint64(arg1[2]) * 2
		x[18] = uint64(arg1[1]) * 2
		x[19] = uint64(arg1[9]) * (x[1] * 2)
		x[20] = uint64(arg1[8]) * x[2]
		x[21] = uint64(arg1[8]) * x[4]
		x[22] = uint64(arg1[7]) * (x[2] * 2)
		x[23] = uint64(arg1[7]) * x[5]
		x[24] = uint64(arg1[7]) * (x[7] * 2)
		x[25] = uint64(arg1[6]) * x[2]
		x[26] = uint64(arg1[6]) * x[5]
		x[27] = uint64(arg1[6]) * x[8]
		x[28] = uint64(arg1[6]) * x[10]
		x[29] = uint64(arg1[5]) * (x[2] * 2)
		x[30] = uint64(arg1[5]) * x[5]
		x[31] = uint64(arg1[5]) * (x[8] * 2)
		x[32] = uint64(arg1[5]) * x[11]
		x[33] = uint64(arg1[5]) * (x[13] * 2)
		x[34] = uint64(arg1[4]) * x[2]
		x[35] = uint64(arg1[4]) * x[5]
		x[36] = uint64(arg1[4]) * x[8]
		x[37] = uint64(arg1[4]) * x[11]
		x[38] = uint64(arg1[4]) * x[14]
		x[39] = uint64(arg1[4]) * (arg1[4])
		x[40] = uint64(arg1[3]) * (x[2] * 2)
		x[41] = uint64(arg1[3]) * x[5]
		x[42] = uint64(arg1[3]) * (x[8] * 2)
		x[43] = uint64(arg1[3]) * x[12]
		x[44] = uint64(arg1[3]) * (x[14] * 2)
		x[45] = uint64(arg1[3]) * x[15]
		x[46] = uint64(arg1[3]) * (arg1[3] * 2)
		x[47] = uint64(arg1[2]) * x[2]
		x[48] = uint64(arg1[2]) * x[5]
		x[49] = uint64(arg1[2]) * x[9]
		x[50] = uint64(arg1[2]) * x[12]
		x[51] = uint64(arg1[2]) * x[14]
		x[52] = uint64(arg1[2]) * x[15]
		x[53] = uint64(arg1[2]) * x[16]
		x[54] = uint64(arg1[2]) * x[17]
		x[55] = uint64(arg1[2]) * (arg1[2])
		x[56] = uint64(arg1[1]) * (x[2] * 2)
		x[57] = uint64(arg1[1]) * x[6]
		x[58] = uint64(arg1[1]) * (x[9] * 2)
		x[59] = uint64(arg1[1]) * x[12]
		x[60] = uint64(arg1[1]) * (x[14] * 2)
		x[61] = uint64(arg1[1]) * x[15]
		x[62] = uint64(arg1[1]) * (x[16] * 2)
		x[63] = uint64(arg1[1]) * x[17]
		x[64] = uint64(arg1[1]) * (arg1[1] * 2)
		x[65] = uint64(arg1[0]) * x[3]
		x[66] = uint64(arg1[0]) * x[6]
		x[67] = uint64(arg1[0]) * x[9]
		x[68] = uint64(arg1[0]) * x[12]
		x[69] = uint64(arg1[0]) * x[14]
		x[70] = uint64(arg1[0]) * x[15]
		x[71] = uint64(arg1[0]) * x[16]
		x[72] = uint64(arg1[0]) * x[17]
		x[73] = uint64(arg1[0]) * x[18]
		x[74] = uint64(arg1[0]) * (arg1[0])
		x[75] = (x[74] + (x[56] + (x[48] + (x[42] + (x[37] + x[33])))))
		x[76] = (x[75] >> 26)
		x[77] = uint32(x[75] & 0x3ffffff)
		x[78] = (x[65] + (x[57] + (x[50] + (x[44] + (x[39] + x[19])))))
		x[79] = (x[66] + (x[58] + (x[51] + (x[45] + (x[20] + x[21])))))
		x[80] = (x[67] + (x[59] + (x[52] + (x[46] + (x[22] + x[21])))))
		x[81] = (x[68] + (x[60] + (x[53] + (x[25] + x[23]))))
		x[82] = (x[69] + (x[61] + (x[54] + (x[29] + (x[26] + x[24])))))
		x[83] = (x[70] + (x[62] + (x[34] + (x[30] + x[27]))))
		x[84] = (x[71] + (x[63] + (x[40] + (x[35] + (x[31] + x[28])))))
		x[85] = (x[72] + (x[47] + (x[41] + (x[36] + x[32]))))
		x[86] = (x[76] + x[85])
		x[87] = (x[86] >> 25)
		x[88] = uint32(x[86] & 0x1ffffff)
		x[89] = (x[87] + x[84])
		x[90] = (x[89] >> 26)
		x[91] = uint32(x[89] & 0x3ffffff)
		x[92] = (x[90] + x[83])
		x[93] = (x[92] >> 25)
		x[94] = uint32(x[92] & 0x1ffffff)
		x[95] = (x[93] + x[82])
		x[96] = (x[95] >> 26)
		x[97] = uint32(x[95] & 0x3ffffff)
		x[98] = (x[96] + x[81])
		x[99] = (x[98] >> 25)
		x[100] = uint32(x[98] & 0x1ffffff)
		x[101] = (x[99] + x[80])
		x[102] = (x[101] >> 26)
		x[103] = uint32(x[101] & 0x3ffffff)
		x[104] = (x[102] + x[79])
		x[105] = (x[104] >> 25)
		x[106] = uint32(x[104] & 0x1ffffff)
		x[107] = (x[105] + x[78])
		x[108] = (x[107] >> 26)
		x[109] = uint32(x[107] & 0x3ffffff)
		x[110] = (x[108] + x[77])
		x[111] = (x[110] >> 25)
		x[112] = uint32(x[110] & 0x1ffffff)
		x[113] = (x[111] * 0x13)
		x[114] = (x[77] + x[113])
		x[115] = uint32(x[114] >> 26)
		x[116] = uint32(x[114] & 0x3ffffff)
		x[117] = (x[115] + x[88])
		x[118] = fiat25519Uint1(x[117] >> 25)
		x[119] = (x[117] & 0x1ffffff)
		x[120] = (x[118] + x[91])
		out1[0] = x[116]
		out1[1] = x[119]
		out1[2] = x[120]
		out1[3] = x[94]
		out1[4] = x[97]
		out1[5] = x[100]
		out1[6] = x[103]
		out1[7] = x[106]
		out1[8] = x[109]
		out1[9] = x[112]
	}
*/
func fiat25519Carry(out1 *Fiat25519FieldElement, arg1 *Fiat25519FieldElement) {
	var x [23]uint32
	x[1] = arg1[0]
	x[2] = (x[1] >> 26) + arg1[1]
	x[3] = (x[2] >> 25) + arg1[2]
	x[4] = (x[3] >> 26) + arg1[3]
	x[5] = (x[4] >> 25) + arg1[4]
	x[6] = (x[5] >> 26) + arg1[5]
	x[7] = (x[6] >> 25) + arg1[6]
	x[8] = (x[7] >> 26) + arg1[7]
	x[9] = (x[8] >> 25) + arg1[8]
	x[10] = (x[9] >> 26) + arg1[9]
	x[11] = (x[1] & 0x3ffffff) + ((x[10] >> 25) * 0x13)
	x[12] = (x[11] >> 26) + (x[2] & 0x1ffffff)
	x[13] = (x[11] & 0x3ffffff)
	x[14] = (x[12] & 0x1ffffff)
	x[15] = (x[12] >> 25) + (x[3] & 0x3ffffff)
	x[16] = (x[4] & 0x1ffffff)
	x[17] = (x[5] & 0x3ffffff)
	x[18] = (x[6] & 0x1ffffff)
	x[19] = (x[7] & 0x3ffffff)
	x[20] = (x[8] & 0x1ffffff)
	x[21] = (x[9] & 0x3ffffff)
	x[22] = (x[10] & 0x1ffffff)
	out1[0] = x[13]
	out1[1] = x[14]
	out1[2] = x[15]
	out1[3] = x[16]
	out1[4] = x[17]
	out1[5] = x[18]
	out1[6] = x[19]
	out1[7] = x[20]
	out1[8] = x[21]
	out1[9] = x[22]
}

func fiat25519Add(out1 *Fiat25519FieldElement, arg1, arg2 *Fiat25519FieldElement) {
	for i := 0; i < 10; i++ {
		(*out1)[i] = (*arg1)[i] + (*arg2)[i]
	}
}

func fiat25519Sub(out1 *Fiat25519FieldElement, arg1 *Fiat25519FieldElement, arg2 *Fiat25519FieldElement) {
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

	x1 = ((uint32(0x7ffffda) + (arg1[0])) - (arg2[0]))
	x2 = ((uint32(0x3fffffe) + (arg1[1])) - (arg2[1]))
	x3 = ((uint32(0x7fffffe) + (arg1[2])) - (arg2[2]))
	x4 = ((uint32(0x3fffffe) + (arg1[3])) - (arg2[3]))
	x5 = ((uint32(0x7fffffe) + (arg1[4])) - (arg2[4]))
	x6 = ((uint32(0x3fffffe) + (arg1[5])) - (arg2[5]))
	x7 = ((uint32(0x7fffffe) + (arg1[6])) - (arg2[6]))
	x8 = ((uint32(0x3fffffe) + (arg1[7])) - (arg2[7]))
	x9 = ((uint32(0x7fffffe) + (arg1[8])) - (arg2[8]))
	x10 = ((uint32(0x3fffffe) + (arg1[9])) - (arg2[9]))
	out1[0] = x1
	out1[1] = x2
	out1[2] = x3
	out1[3] = x4
	out1[4] = x5
	out1[5] = x6
	out1[6] = x7
	out1[7] = x8
	out1[8] = x9
	out1[9] = x10
}

func fiat25519Opp(out1 *Fiat25519FieldElement, arg1 *Fiat25519FieldElement) {
	const mask1 = 0x7fffffe
	const mask2 = 0x3fffffe
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			(*out1)[i] = mask1 - (*arg1)[i]
		} else {
			(*out1)[i] = mask2 - (*arg1)[i]
		}
	}
}

func fiat25519Selectznz(out1 *[10]uint32, arg1 fiat25519Uint1, arg2, arg3 *[10]uint32) {
	for i := 0; i < 10; i++ {
		if arg1 == 0 {
			(*out1)[i] = (*arg2)[i]
		} else {
			(*out1)[i] = (*arg3)[i]
		}
	}
}

func fiat25519ToBytes(out1 *[]byte, arg1 *Fiat25519FieldElement) {
	var x1 uint32
	var x2 fiat25519Uint1
	var x3 uint32
	var x4 fiat25519Uint1
	var x5 uint32
	var x6 fiat25519Uint1
	var x7 uint32
	var x8 fiat25519Uint1
	var x9 uint32
	var x10 fiat25519Uint1
	var x11 uint32
	var x12 fiat25519Uint1
	var x13 uint32
	var x14 fiat25519Uint1

	var x15 uint32
	var x16 fiat25519Uint1
	var x17 uint32
	var x18 fiat25519Uint1
	var x19 uint32
	var x20 fiat25519Uint1
	var x21 uint32
	var x22 uint32

	var x23 fiat25519Uint1
	var x24 uint32
	var x25 fiat25519Uint1
	var x26 uint32
	var x27 fiat25519Uint1
	var x28 uint32
	var x29 fiat25519Uint1

	var x30 uint32
	var x31 fiat25519Uint1
	var x32 uint32
	var x33 fiat25519Uint1
	var x34 uint32
	var x35 fiat25519Uint1
	var x36 uint32
	var x37 fiat25519Uint1
	var x38 uint32
	var x39 fiat25519Uint1
	var x40 uint32
	var x41 fiat25519Uint1

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
	var x89 fiat25519Uint1

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

	fiat25519SubBorrowXU26(&x1, &x2, 0x0, (arg1[0]), uint32(0x3ffffed))
	fiat25519SubBorrowXU25(&x3, &x4, x2, (arg1[1]), uint32(0x1ffffff))
	fiat25519SubBorrowXU26(&x5, &x6, x4, (arg1[2]), uint32(0x3ffffff))
	fiat25519SubBorrowXU25(&x7, &x8, x6, (arg1[3]), uint32(0x1ffffff))
	fiat25519SubBorrowXU26(&x9, &x10, x8, (arg1[4]), uint32(0x3ffffff))
	fiat25519SubBorrowXU25(&x11, &x12, x10, (arg1[5]), uint32(0x1ffffff))
	fiat25519SubBorrowXU26(&x13, &x14, x12, (arg1[6]), uint32(0x3ffffff))
	fiat25519SubBorrowXU25(&x15, &x16, x14, (arg1[7]), uint32(0x1ffffff))
	fiat25519SubBorrowXU26(&x17, &x18, x16, (arg1[8]), uint32(0x3ffffff))
	fiat25519SubBorrowXU25(&x19, &x20, x18, (arg1[9]), uint32(0x1ffffff))
	fiat25519CmovznzU32(&x21, x20, 0x0, uint32(0xffffffff))
	fiat25519SubBorrowXU26(&x22, &x23, 0x0, x1, (x21 & uint32(0x3ffffed)))
	fiat25519AddCarryXU25(&x24, &x25, x23, x3, (x21 & uint32(0x1ffffff)))
	fiat25519SubBorrowXU26(&x26, &x27, x25, x5, (x21 & uint32(0x3ffffff)))
	fiat25519AddCarryXU25(&x28, &x29, x27, x7, (x21 & uint32(0x1ffffff)))
	fiat25519SubBorrowXU26(&x30, &x31, x29, x9, (x21 & uint32(0x3ffffff)))
	fiat25519AddCarryXU25(&x32, &x33, x31, x11, (x21 & uint32(0x1ffffff)))
	fiat25519SubBorrowXU26(&x34, &x35, x33, x13, (x21 & uint32(0x3ffffff)))
	fiat25519AddCarryXU25(&x36, &x37, x35, x15, (x21 & uint32(0x1ffffff)))
	fiat25519SubBorrowXU26(&x38, &x39, x37, x17, (x21 & uint32(0x3ffffff)))
	fiat25519AddCarryXU25(&x40, &x41, x39, x19, (x21 & uint32(0x1ffffff)))
	x42 = (x40 << 6)
	x43 = (x38 << 4)
	x44 = (x36 << 3)
	x45 = (x34 * 0x2)
	x46 = (x30 << 6)
	x47 = (x28 << 5)
	x48 = (x26 << 3)
	x49 = (x24 << 2)
	x50 = uint8(x22 & 0xff)
	x51 = (x22 >> 8)
	x52 = uint8(x51 & 0xff)
	x53 = (x51 >> 8)
	x54 = uint8(x53 & 0xff)
	x55 = uint8(x53 >> 8)
	x56 = (x49 + uint32(x55))
	x57 = uint8(x56 & 0xff)
	x58 = (x56 >> 8)
	x59 = uint8(x58 & 0xff)
	x60 = (x58 >> 8)
	x61 = uint8(x60 & 0xff)
	x62 = uint8(x60 >> 8)
	x63 = (x48 + uint32(x62))
	x64 = uint8(x63 & 0xff)
	x65 = (x63 >> 8)
	x66 = uint8(x65 & 0xff)
	x67 = (x65 >> 8)
	x68 = uint8(x67 & 0xff)
	x69 = uint8(x67 >> 8)
	x70 = (x47 + uint32(x69))
	x71 = uint8(x70 & 0xff)
	x72 = (x70 >> 8)
	x73 = uint8(x72 & (0xff))
	x74 = (x72 >> 8)
	x75 = uint8(x74 & (0xff))
	x76 = uint8(x74 >> 8)
	x77 = (x46 + uint32(x76))
	x78 = uint8(x77 & (0xff))
	x79 = (x77 >> 8)
	x80 = uint8(x79 & (0xff))
	x81 = (x79 >> 8)
	x82 = uint8(x81 & (0xff))
	x83 = uint8(x81 >> 8)
	x84 = uint8(x32 & (0xff))
	x85 = (x32 >> 8)
	x86 = uint8(x85 & (0xff))
	x87 = (x85 >> 8)
	x88 = uint8(x87 & (0xff))
	x89 = (fiat25519Uint1)(x87 >> 8)
	x90 = (x45 + uint32(x89))
	x91 = uint8(x90 & (0xff))
	x92 = (x90 >> 8)
	x93 = uint8(x92 & (0xff))
	x94 = (x92 >> 8)
	x95 = uint8(x94 & (0xff))
	x96 = uint8(x94 >> 8)
	x97 = (x44 + uint32(x96))
	x98 = uint8(x97 & (0xff))
	x99 = (x97 >> 8)
	x100 = uint8(x99 & (0xff))
	x101 = (x99 >> 8)
	x102 = uint8(x101 & (0xff))
	x103 = uint8(x101 >> 8)
	x104 = (x43 + uint32(x103))
	x105 = uint8(x104 & (0xff))
	x106 = (x104 >> 8)
	x107 = uint8(x106 & (0xff))
	x108 = (x106 >> 8)
	x109 = uint8(x108 & (0xff))
	x110 = uint8(x108 >> 8)
	x111 = (x42 + uint32(x110))
	x112 = uint8(x111 & (0xff))
	x113 = (x111 >> 8)
	x114 = uint8(x113 & 0xff)
	x115 = (x113 >> 8)
	x116 = uint8(x115 & 0xff)
	x117 = uint8(x115 >> 8)
	(*out1)[0] = x50
	(*out1)[1] = x52
	(*out1)[2] = x54
	(*out1)[3] = x57
	(*out1)[4] = x59
	(*out1)[5] = x61
	(*out1)[6] = x64
	(*out1)[7] = x66
	(*out1)[8] = x68
	(*out1)[9] = x71
	(*out1)[10] = x73
	(*out1)[11] = x75
	(*out1)[12] = x78
	(*out1)[13] = x80
	(*out1)[14] = x82
	(*out1)[15] = x83
	(*out1)[16] = x84
	(*out1)[17] = x86
	(*out1)[18] = x88
	(*out1)[19] = x91
	(*out1)[20] = x93
	(*out1)[21] = x95
	(*out1)[22] = x98
	(*out1)[23] = x100
	(*out1)[24] = x102
	(*out1)[25] = x105
	(*out1)[26] = x107
	(*out1)[27] = x109
	(*out1)[28] = x112
	(*out1)[29] = x114
	(*out1)[30] = x116
	(*out1)[31] = x117
}

func fiat25519FromBytes(out1 *Fiat25519FieldElement, arg1 []byte) {
	var x [79]uint32
	x[1] = uint32(arg1[31]) << 18
	x[2] = uint32(arg1[30]) << 10
	x[3] = uint32(arg1[29]) << 2
	x[4] = uint32(arg1[28]) << 20
	x[5] = uint32(arg1[27]) << 12
	x[6] = uint32(arg1[26]) << 4
	x[7] = uint32(arg1[25]) << 21
	x[8] = uint32(arg1[24]) << 13
	x[9] = uint32(arg1[23]) << 5
	x[10] = uint32(arg1[22]) << 23
	x[11] = uint32(arg1[21]) << 15
	x[12] = uint32(arg1[20]) << 7
	x[13] = uint32(arg1[19]) << 24
	x[14] = uint32(arg1[18]) << 16
	x[15] = uint32(arg1[17]) << 8
	x[16] = uint32(arg1[16])
	x[17] = uint32(arg1[15]) << 18
	x[18] = uint32(arg1[14]) << 10
	x[19] = uint32(arg1[13]) << 2
	x[20] = uint32(arg1[12]) << 19
	x[21] = uint32(arg1[11]) << 11
	x[22] = uint32(arg1[10]) << 3
	x[23] = uint32(arg1[9]) << 21
	x[24] = uint32(arg1[8]) << 13
	x[25] = uint32(arg1[7]) << 5
	x[26] = uint32(arg1[6]) << 22
	x[27] = uint32(arg1[5]) << 14
	x[28] = uint32(arg1[4]) << 6
	x[29] = uint32(arg1[3]) << 24
	x[30] = uint32(arg1[2]) << 16
	x[31] = uint32(arg1[1]) << 8
	x[32] = uint32(arg1[0])

	x[33] = x[31] + x[32]
	x[34] = x[30] + x[33]
	x[35] = x[29] + x[34]
	x[36] = x[35] & 0x3ffffff
	x[37] = x[35] >> 26
	x[38] = x[28] + x[37]
	x[39] = x[27] + x[38]
	x[40] = x[26] + x[39]
	x[41] = x[40] & 0x1ffffff
	x[42] = x[40] >> 25
	x[43] = x[25] + x[42]
	x[44] = x[24] + x[43]
	x[45] = x[23] + x[44]
	x[46] = x[45] & 0x3ffffff
	x[47] = x[45] >> 26
	x[48] = x[22] + x[47]
	x[49] = x[21] + x[48]
	x[50] = x[20] + x[49]
	x[51] = x[50] & 0x1ffffff
	x[52] = x[50] >> 25
	x[53] = x[19] + x[52]
	x[54] = x[18] + x[53]
	x[55] = x[17] + x[54]
	x[56] = x[15] + x[16]
	x[57] = x[14] + x[56]
	x[58] = x[13] + x[57]
	x[59] = x[58] & 0x1ffffff
	x[60] = x[58] >> 25
	x[61] = x[12] + x[60]
	x[62] = x[11] + x[61]
	x[63] = x[10] + x[62]
	x[64] = x[63] & 0x3ffffff
	x[65] = x[63] >> 26
	x[66] = x[9] + x[65]
	x[67] = x[8] + x[66]
	x[68] = x[7] + x[67]
	x[69] = x[68] & 0x1ffffff
	x[70] = x[68] >> 25
	x[71] = x[6] + x[70]
	x[72] = x[5] + x[71]
	x[73] = x[4] + x[72]
	x[74] = x[73] & 0x3ffffff
	x[75] = x[73] >> 26
	x[76] = x[3] + x[75]
	x[77] = x[2] + x[76]
	x[78] = x[1] + x[77]

	out1[0] = x[36]
	out1[1] = x[41]
	out1[2] = x[46]
	out1[3] = x[51]
	out1[4] = x[55]
	out1[5] = x[59]
	out1[6] = x[64]
	out1[7] = x[69]
	out1[8] = x[74]
	out1[9] = x[78]
}

func fiat25519Relax(out1 *Fiat25519FieldElement, arg1 *Fiat25519FieldElement) {
	for i := 0; i < 10; i++ {
		out1[i] = arg1[i]
	}
}

func fiat25519CarryScmul121666(out1 *Fiat25519FieldElement, arg1 *Fiat25519FieldElement) {

	var (
		x1, x2, x3, x4, x5, x6, x7, x8, x9, x10                                                                                                                   uint64
		x11, x12, x13, x14, x15, x16, x17, x18, x19, x20, x21, x22, x23, x24, x25, x26, x27, x28, x29, x30, x31, x32, x33, x34, x35, x36, x37, x38, x39, x40, x41 uint32
	)

	var x42 fiat25519Uint1
	var x43 uint32
	var x44 uint32
	var x45 fiat25519Uint1
	var x46 uint32
	var x47 uint32

	x1 = uint64(uint32(0x1db42) * uint32(arg1[9]))
	x2 = uint64(uint32(0x1db42) * uint32(arg1[8]))
	x3 = uint64(uint32(0x1db42) * uint32(arg1[7]))
	x4 = uint64(uint32(0x1db42) * uint32(arg1[6]))
	x5 = uint64(uint32(0x1db42) * uint32(arg1[5]))
	x6 = uint64(uint32(0x1db42) * uint32(arg1[4]))
	x7 = uint64(uint32(0x1db42) * uint32(arg1[3]))
	x8 = uint64(uint32(0x1db42) * uint32(arg1[2]))
	x9 = uint64(uint32(0x1db42) * uint32(arg1[1]))
	x10 = uint64(uint32(0x1db42) * uint32(arg1[0]))
	x11 = uint32(x10 >> 26)
	x12 = (uint32(x10) & uint32(0x3ffffff))
	x13 = (x11 + uint32(x9))
	x14 = (x13 >> 25)
	x15 = (x13 & uint32(0x1ffffff))
	x16 = (x14 + uint32(x8))
	x17 = (x16 >> 26)
	x18 = (x16 & uint32(0x3ffffff))
	x19 = (x17 + uint32(x7))
	x20 = (x19 >> 25)
	x21 = (x19 & uint32(0x1ffffff))
	x22 = (x20 + uint32(x6))
	x23 = (x22 >> 26)
	x24 = (x22 & uint32(0x3ffffff))
	x25 = (x23 + uint32(x5))
	x26 = (x25 >> 25)
	x27 = (x25 & uint32(0x1ffffff))
	x28 = (x26 + uint32(x4))
	x29 = (x28 >> 26)
	x30 = (x28 & uint32(0x3ffffff))
	x31 = (x29 + uint32(x3))
	x32 = (x31 >> 25)
	x33 = (x31 & uint32(0x1ffffff))
	x34 = (x32 + uint32(x2))
	x35 = (x34 >> 26)
	x36 = (x34 & uint32(0x3ffffff))
	x37 = (x35 + uint32(x1))
	x38 = (x37 >> 25)
	x39 = (x37 & uint32(0x1ffffff))
	x40 = (x38 * uint32(0x13))
	x41 = (x12 + x40)
	x42 = (fiat25519Uint1)(x41 >> 26)
	x43 = (x41 & uint32(0x3ffffff))
	x44 = (uint32(x42) + x15)
	x45 = (fiat25519Uint1)(x44 >> 25)
	x46 = (x44 & uint32(0x1ffffff))
	x47 = (uint32(x45) + x18)
	out1[0] = x43
	out1[1] = x46
	out1[2] = x47
	out1[3] = x21
	out1[4] = x24
	out1[5] = x27
	out1[6] = x30
	out1[7] = x33
	out1[8] = x36
	out1[9] = x39
}
