package main

import (
	"fmt"
	"log"
	"math/big"
)

func main() {
	//var output [300]byte
	//var bytes [64]uint8
	xx := [10]uint32{0x01, 0x02}
	arg1 := [10]uint32{15540839, 8128613, 35539908, 18003200, 66970125, 13484806, 53323836, 19064896, 18909259, 33172296}

	out := fiat25519CarrySquare(xx, &arg1)
	fmt.Printf("out:%v\r\n", out)
	return
}
func fiat25519CarrySquare(out1 [10]uint32, arg1 *[10]uint32) [10]uint32 {
	var (
		x1, x2, x3, x4, x5, x6, x7, x8, x9, x10,
		x11, x12, x13, x14, x15, x16, x17, x18, x19, x20,
		x21, x22, x23, x24, x25, x26, x27, x28, x29, x30,
		x31, x32, x33, x34, x35, x36, x37, x38, x39, x40,
		x41, x42, x43, x44, x45, x46, x47, x48, x49, x50,
		x51, x52, x53, x54, x55, x56, x57, x58, x59, x60,
		x61, x62, x63, x64, x65, x66, x67, x68, x69, x70,
		x71, x72, x73, x74, x75, x76, x77, x78, x79, x80,
		x81, x82, x83, x84, x85, x86, x87, x88, x89, x90,
		x91, x92, x93, x94, x95, x96, x97, x98, x99, x100,
		x101, x102, x103, x104, x105, x106, x107, x108, x109, x110,
		x111, x112, x113, x114, x115, x116, x117, x118, x119, x120 big.Int
	)

	// 计算过程
	x1.Mul(big.NewInt(int64(arg1[9])), big.NewInt(0x13))
	x2.Mul(&x1, big.NewInt(0x2))

	x3.Mul(big.NewInt(int64(arg1[9])), big.NewInt(0x2))
	x4.Mul(big.NewInt(int64(arg1[8])), big.NewInt(0x13))
	x5.Mul(&x4, big.NewInt(0x2))
	x6.Mul(big.NewInt(int64(arg1[8])), big.NewInt(0x2))
	x7.Mul(big.NewInt(int64(arg1[7])), big.NewInt(0x13))
	x8.Mul(&x7, big.NewInt(0x2))
	x9.Mul(big.NewInt(int64(arg1[7])), big.NewInt(0x2))
	x10.Mul(big.NewInt(int64(arg1[6])), big.NewInt(0x13))
	x11.Mul(&x10, big.NewInt(0x2))
	x12.Mul(big.NewInt(int64(arg1[6])), big.NewInt(0x2))
	x13.Mul(big.NewInt(int64(arg1[5])), big.NewInt(0x13))
	x14.Mul(big.NewInt(int64(arg1[5])), big.NewInt(0x2))
	x15.Mul(big.NewInt(int64(arg1[4])), big.NewInt(0x2))
	x16.Mul(big.NewInt(int64(arg1[3])), big.NewInt(0x2))
	x17.Mul(big.NewInt(int64(arg1[2])), big.NewInt(0x2))
	x18.Mul(big.NewInt(int64(arg1[1])), big.NewInt(0x2))
	//x19.Mul(big.NewInt(int64(arg1[9])), x1.Mul(&x1, big.NewInt(0x2)))
	x19.Mul(big.NewInt(int64(arg1[9])), big.NewInt(0).Mul(&x1, big.NewInt(0x2)))
	x20.Mul(big.NewInt(int64(arg1[8])), &x2)
	x21.Mul(big.NewInt(int64(arg1[8])), &x4)
	//x22.Mul(big.NewInt(int64(arg1[7])), x2.Mul(&x2, big.NewInt(0x2)))
	x22.Mul(big.NewInt(int64(arg1[7])), big.NewInt(0).Mul(&x2, big.NewInt(0x2)))
	x23.Mul(big.NewInt(int64(arg1[7])), &x5)
	//x24.Mul(big.NewInt(int64(arg1[7])), x7.Mul(&x7, big.NewInt(0x2)))
	x24.Mul(big.NewInt(int64(arg1[7])), big.NewInt(0).Mul(&x7, big.NewInt(0x2)))

	x25.Mul(big.NewInt(int64(arg1[6])), &x2)
	x26.Mul(big.NewInt(int64(arg1[6])), &x5)
	x27.Mul(big.NewInt(int64(arg1[6])), &x8)
	x28.Mul(big.NewInt(int64(arg1[6])), &x10)

	//x29.Mul(big.NewInt(int64(arg1[5])), x2.Mul(&x2, big.NewInt(0x2)))
	x29.Mul(big.NewInt(int64(arg1[5])), big.NewInt(0).Mul(&x2, big.NewInt(0x2)))

	x30.Mul(big.NewInt(int64(arg1[5])), &x5)
	//x31.Mul(big.NewInt(int64(arg1[5])), x8.Mul(&x8, big.NewInt(0x2)))

	x31.Mul(big.NewInt(int64(arg1[5])), big.NewInt(0).Mul(&x8, big.NewInt(0x2)))

	x32.Mul(big.NewInt(int64(arg1[5])), &x11)
	//x33.Mul(big.NewInt(int64(arg1[5])), x13.Mul(&x13, big.NewInt(0x2)))
	x33.Mul(big.NewInt(int64(arg1[5])), big.NewInt(0).Mul(&x13, big.NewInt(0x2)))

	x34.Mul(big.NewInt(int64(arg1[4])), &x2)
	x35.Mul(big.NewInt(int64(arg1[4])), &x5)
	x36.Mul(big.NewInt(int64(arg1[4])), &x8)
	x37.Mul(big.NewInt(int64(arg1[4])), &x11)
	x38.Mul(big.NewInt(int64(arg1[4])), &x14)
	x39.Mul(big.NewInt(int64(arg1[4])), big.NewInt(int64(arg1[4])))

	//x40.Mul(big.NewInt(int64(arg1[3])), x2.Mul(&x2, big.NewInt(0x2)))
	x40.Mul(big.NewInt(int64(arg1[3])), big.NewInt(0).Mul(&x2, big.NewInt(0x2)))

	x41.Mul(big.NewInt(int64(arg1[3])), &x5)
	//x42.Mul(big.NewInt(int64(arg1[3])), x8.Mul(&x8, big.NewInt(0x2)))
	x42.Mul(big.NewInt(int64(arg1[3])), big.NewInt(0).Mul(&x8, big.NewInt(0x2)))

	x43.Mul(big.NewInt(int64(arg1[3])), &x12)
	//x44.Mul(big.NewInt(int64(arg1[3])), x14.Mul(&x14, big.NewInt(0x2)))
	x44.Mul(big.NewInt(int64(arg1[3])), big.NewInt(0).Mul(&x14, big.NewInt(0x2)))

	x45.Mul(big.NewInt(int64(arg1[3])), &x15)
	//x46.Mul(big.NewInt(int64(arg1[3])), big.NewInt(int64(arg1[3])).Mul(big.NewInt(0x2)))
	x46.Mul(big.NewInt(int64(arg1[3])), big.NewInt(0).Mul(big.NewInt(int64(arg1[3])), big.NewInt(0x2)))

	x47.Mul(big.NewInt(int64(arg1[2])), &x2)
	x48.Mul(big.NewInt(int64(arg1[2])), &x5)
	x49.Mul(big.NewInt(int64(arg1[2])), &x9)
	x50.Mul(big.NewInt(int64(arg1[2])), &x12)
	x51.Mul(big.NewInt(int64(arg1[2])), &x14)
	x52.Mul(big.NewInt(int64(arg1[2])), &x15)
	x53.Mul(big.NewInt(int64(arg1[2])), &x16)
	x54.Mul(big.NewInt(int64(arg1[2])), big.NewInt(int64(arg1[2])))

	//x55.Mul(big.NewInt(int64(arg1[1])), x2.Mul(&x2, big.NewInt(0x2)))
	//x55 = ((arg1[1]) * ((uint64_t)x2 * 0x2));
	x55.Mul(big.NewInt(int64(arg1[1])), big.NewInt(0).Mul(&x2, big.NewInt(0x2)))
	log.Printf("x2:%d\r\n", x2)

	x56.Mul(big.NewInt(int64(arg1[1])), &x6)
	//x57.Mul(big.NewInt(int64(arg1[1])), x9.Mul(&x9, big.NewInt(0x2)))
	x57.Mul(big.NewInt(int64(arg1[1])), big.NewInt(0).Mul(&x9, big.NewInt(0x2)))

	x58.Mul(big.NewInt(int64(arg1[1])), &x12)
	//x59.Mul(big.NewInt(int64(arg1[1])), x14.Mul(&x14, big.NewInt(0x2)))
	x59.Mul(big.NewInt(int64(arg1[1])), big.NewInt(0).Mul(&x14, big.NewInt(0x2)))
	x60.Mul(big.NewInt(int64(arg1[1])), &x15)
	//x61.Mul(big.NewInt(int64(arg1[1])), x16.Mul(&x16, big.NewInt(0x2)))
	x61.Mul(big.NewInt(int64(arg1[1])), big.NewInt(0).Mul(&x16, big.NewInt(0x2)))
	x62.Mul(big.NewInt(int64(arg1[1])), &x17)
	//x63.Mul(big.NewInt(int64(arg1[1])), big.NewInt(int64(arg1[1])).Mul(big.NewInt(0x2)))

	x63.Mul(big.NewInt(int64(arg1[1])), big.NewInt(0).Mul(big.NewInt(int64(arg1[1])), big.NewInt(0x2)))

	x64.Mul(big.NewInt(int64(arg1[0])), &x3)
	x65.Mul(big.NewInt(int64(arg1[0])), &x6)
	x66.Mul(big.NewInt(int64(arg1[0])), &x9)
	x67.Mul(big.NewInt(int64(arg1[0])), &x12)
	x68.Mul(big.NewInt(int64(arg1[0])), &x14)
	x69.Mul(big.NewInt(int64(arg1[0])), &x15)
	x70.Mul(big.NewInt(int64(arg1[0])), &x16)
	x71.Mul(big.NewInt(int64(arg1[0])), &x17)
	x72.Mul(big.NewInt(int64(arg1[0])), &x18)

	x73.Mul(big.NewInt(int64(arg1[0])), big.NewInt(int64(arg1[0])))
	fmt.Printf(" x73:%d  x55:%d  x48:%d  x42:%d  x37:%d x33:%d \r\n", x73.Uint64(), x55.Uint64(), x48.Uint64(), x42.Uint64(), x37.Uint64(), x33.Uint64())
	//x74 = (x73 + (x55 + (x48 + (x42 + (x37 + x33)))))
	tempX74 := big.NewInt(0).Add(&x73, big.NewInt(0).Add(&x55, big.NewInt(0).Add(&x48, big.NewInt(0).Add(&x42, big.NewInt(0).Add(&x37, &x33)))))
	x74.Set(tempX74)

	//	x75 = (x74 >> 26)
	var x74Copy big.Int
	x74Copy.Set(&x74)
	x75.Set(x74Copy.Rsh(&x74Copy, 26))

	//x76 = uint32(x74 & 0x3ffffff)
	x76.And(&x74, big.NewInt(0x3ffffff))

	//x77 = (x64 + (x56 + (x49 + (x43 + x38))))
	tempX77 := big.NewInt(0).Add(&x64, big.NewInt(0).Add(&x56, big.NewInt(0).Add(&x49, big.NewInt(0).Add(&x43, &x38))))
	x77.Set(tempX77)

	//x78 = (x65 + (x57 + (x50 + (x44 + (x39 + x19)))))
	tempx78 := big.NewInt(0).Add(&x65, big.NewInt(0).Add(&x57, big.NewInt(0).Add(&x50, big.NewInt(0).Add(&x44, big.NewInt(0).Add(&x39, &x19)))))
	x78.Set(tempx78)

	// x79 = (x66 + (x58 + (x51 + (x45 + x20))))
	temp12 := big.NewInt(0).Add(&x45, &x20)
	temp13 := big.NewInt(0).Add(&x51, temp12)
	temp14 := big.NewInt(0).Add(&x58, temp13)
	x79.Add(&x66, temp14)

	// x80 = (x67 + (x59 + (x52 + (x46 + (x22 + x21)))))
	temp15 := big.NewInt(0).Add(&x22, &x21)
	temp16 := big.NewInt(0).Add(&x46, temp15)
	temp17 := big.NewInt(0).Add(&x52, temp16)
	temp18 := big.NewInt(0).Add(&x59, temp17)
	x80.Add(&x67, temp18)

	// x81 = (x68 + (x60 + (x53 + (x25 + x23))))
	temp19 := big.NewInt(0).Add(&x25, &x23)
	temp20 := big.NewInt(0).Add(&x53, temp19)
	temp21 := big.NewInt(0).Add(&x60, temp20)
	x81.Add(&x68, temp21)

	// x82 = (x69 + (x61 + (x54 + (x29 + (x26 + x24)))))
	temp22 := big.NewInt(0).Add(&x29, &x26)
	temp23 := big.NewInt(0).Add(&x24, temp22)
	temp24 := big.NewInt(0).Add(&x54, temp23)
	temp25 := big.NewInt(0).Add(&x61, temp24)
	x82.Add(&x69, temp25)

	// x83 = (x70 + (x62 + (x34 + (x30 + x27))))
	temp26 := big.NewInt(0).Add(&x34, &x30)
	temp27 := big.NewInt(0).Add(&x27, temp26)
	temp28 := big.NewInt(0).Add(&x62, temp27)
	x83.Add(&x70, temp28)

	// x84 = (x71 + (x63 + (x40 + (x35 + (x31 + x28)))))

	temp1 := big.NewInt(0).Add(&x28, &x31)
	temp2 := big.NewInt(0).Add(&x35, temp1)
	temp3 := big.NewInt(0).Add(&x40, temp2)
	temp4 := big.NewInt(0).Add(&x63, temp3)
	x84.Add(&x71, temp4)

	// x85 = (x72 + (x47 + (x41 + (x36 + x32))))
	temp32 := big.NewInt(0).Add(&x36, &x32)
	temp33 := big.NewInt(0).Add(&x41, temp32)
	temp34 := big.NewInt(0).Add(&x47, temp33)
	x85.Add(&x72, temp34)

	// x86 = (x75 + x85)
	x86.Add(&x75, &x85)

	// x87 = (x86 >> 25)
	var x86Copy big.Int
	x86Copy.Set(&x86)
	x87.Set(x86Copy.Rsh(&x86Copy, 25))

	// x88 = uint32(x86 & 0x1ffffff)
	x88.And(&x86, big.NewInt(0x1ffffff))
	fmt.Printf("x86:%d\r\n", x86.Uint64())

	// x89 = (x87 + x84)
	x89.Add(&x87, &x84)
	fmt.Printf(" x87:%d x84:%d\r\n", x87, x84)
	// x90 = (x89 >> 26)
	var x89Copy big.Int
	x89Copy.Set(&x89)
	x90.Set(x89Copy.Rsh(&x89Copy, 26))

	// x91 = uint32(x89 & 0x3ffffff)
	x91.And(&x89, big.NewInt(0x3ffffff))

	// x92 = (x90 + x83)
	x92.Add(&x90, &x83)

	// x93 = (x92 >> 25)
	var x92Copy big.Int
	x92Copy.Set(&x92)
	x93.Set(x92Copy.Rsh(&x92Copy, 25))

	// x94 = uint32(x92 & 0x1ffffff)
	x94.And(&x92, big.NewInt(0x1ffffff))

	// x95 = (x93 + x82)
	x95.Add(&x93, &x82)

	// x96 = (x95 >> 26)
	var x95Copy big.Int
	x95Copy.Set(&x95)
	x96.Set(x95Copy.Rsh(&x95Copy, 26))
	// x97 = uint32(x95 & 0x3ffffff)
	x97.And(&x95, big.NewInt(0x3ffffff))

	// x98 = (x96 + x81)
	x98.Add(&x96, &x81)

	// x99 = (x98 >> 25)
	var x98Copy big.Int
	x98Copy.Set(&x98)
	x99.Set(x98Copy.Rsh(&x98Copy, 25))

	// x100 = uint32(x98 & 0x1ffffff)
	x100.And(&x98, big.NewInt(0x1ffffff))

	// x101 = (x99 + x80)
	x101.Add(&x99, &x80)

	// x102 = (x101 >> 26)
	var x101Copy big.Int
	x101Copy.Set(&x101)
	x102.Set(x101Copy.Rsh(&x101Copy, 26))

	// x103 = uint32(x101 & 0x3ffffff)
	x103.And(&x101, big.NewInt(0x3ffffff))

	// x104 = (x102 + x79)
	x104.Add(&x102, &x79)

	// x105 = (x104 >> 25)
	var x104Copy big.Int
	x104Copy.Set(&x104)
	x105.Set(x104Copy.Rsh(&x104Copy, 25))

	// x106 = uint32(x104 & 0x1ffffff)
	x106.And(&x104, big.NewInt(0x1ffffff))

	// x107 = (x105 + x78)
	x107.Add(&x105, &x78)

	// x108 = (x107 >> 26)
	var x107Copy big.Int
	x107Copy.Set(&x107)
	x108.Set(x107Copy.Rsh(&x107Copy, 26))

	// x109 = uint32(x107 & 0x3ffffff)
	x109.And(&x107, big.NewInt(0x3ffffff))

	// x110 = (x108 + x77)
	x110.Add(&x108, &x77)

	// x111 = (x110 >> 25)
	var x110Copy big.Int
	x110Copy.Set(&x110)
	x111.Set(x110Copy.Rsh(&x110Copy, 25))

	// x112 = uint32(x110 & 0x1ffffff)
	x112.And(&x110, big.NewInt(0x1ffffff))

	// x113 = (x111 * 0x13)
	x113.Mul(&x111, big.NewInt(0x13))

	// x114 = (uint64(x76) + x113)
	x114.Add(&x76, &x113)

	// x115 = uint32(x114 >> 26)
	var x114Copy big.Int
	x114Copy.Set(&x114)
	x115.Set(x114Copy.Rsh(&x114Copy, 26))

	// x116 = uint32(x114 & 0x3ffffff)
	x116.And(&x114, big.NewInt(0x3ffffff))

	// x117 = uint64((x115 + x88))
	x117.Add(&x115, &x88)
	fmt.Printf("x88:%d\r\n", x88.Uint64())
	// x118 := (fiat25519Uint1)(x117 >> 25) // 假设 fiat_25519_uint1 是一个自定义类型，这里直接使用 uint64
	var x117Copy big.Int
	x117Copy.Set(&x117)
	x118.Set(x117Copy.Rsh(&x117Copy, 25))

	// x119 = uint32(x117 & 0x1ffffff)
	x119.And(&x117, big.NewInt(0x1ffffff))

	// x120 = (uint32(x118) + x91)
	x120.Add(&x118, &x91)

	fmt.Printf("x91:%d\r\n", x91)
	out1[0] = uint32(x116.Uint64())
	out1[1] = uint32(x119.Uint64())
	out1[2] = uint32(x120.Uint64())
	out1[3] = uint32(x94.Uint64())
	out1[4] = uint32(x97.Uint64())
	out1[5] = uint32(x100.Uint64())
	out1[6] = uint32(x103.Uint64())
	out1[7] = uint32(x106.Uint64())
	out1[8] = uint32(x109.Uint64())
	out1[9] = uint32(x112.Uint64())

	return out1
}
