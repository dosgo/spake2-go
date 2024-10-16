package ed25519

import (
	"encoding/binary"
)

type fiat25519Uint1 uint8
type fiat25519Int1 int8

type fiat25519LooseFieldElement [10]uint32
type fiat25519TightFieldElement [10]uint32

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
	x1 := int32(arg2) - int32(arg1) - int32(arg3)
	*out1 = uint32(x1 & 0x1ffffff)
	*out2 = fiat25519Uint1(0 - (x1 >> 25))
}

func fiat25519CmovznzU32(out1 *uint32, arg1 fiat25519Uint1, arg2 uint32, arg3 uint32) {
	x1 := fiat25519Uint1(0)
	if arg1 != 0 {
		x1 = 1
	}
	x2 := uint32(0 - x1)
	*out1 = (fiat25519ValueBarrierU32(x2) & arg3) | (fiat25519ValueBarrierU32(^x2) & arg2)
}

func fiat25519CarryMul(out1 fiat25519TightFieldElement, arg1 fiat25519LooseFieldElement, arg2 fiat25519LooseFieldElement) {
	var x [84]uint64
	var x103 uint32
	var x145 fiat25519Uint1

	x[0] = uint64(arg1[9]) * (uint64(arg2[9]) * 0x26)
	x[1] = uint64(arg1[9]) * (uint64(arg2[8]) * 0x13)
	x[2] = uint64(arg1[9]) * (uint64(arg2[7]) * 0x26)
	x[3] = uint64(arg1[9]) * (uint64(arg2[6]) * 0x13)
	x[4] = uint64(arg1[9]) * (uint64(arg2[5]) * 0x26)
	x[5] = uint64(arg1[9]) * (uint64(arg2[4]) * 0x13)
	x[6] = uint64(arg1[9]) * (uint64(arg2[3]) * 0x26)
	x[7] = uint64(arg1[9]) * (uint64(arg2[2]) * 0x13)
	x[8] = uint64(arg1[9]) * (uint64(arg2[1]) * 0x26)
	x[9] = uint64(arg1[8]) * (uint64(arg2[9]) * 0x13)
	x[10] = uint64(arg1[8]) * (uint64(arg2[8]) * 0x13)
	x[11] = uint64(arg1[8]) * (uint64(arg2[7]) * 0x13)
	x[12] = uint64(arg1[8]) * (uint64(arg2[6]) * 0x13)
	x[13] = uint64(arg1[8]) * (uint64(arg2[5]) * 0x13)
	x[14] = uint64(arg1[8]) * (uint64(arg2[4]) * 0x13)
	x[15] = uint64(arg1[8]) * (uint64(arg2[3]) * 0x13)
	x[16] = uint64(arg1[8]) * (uint64(arg2[2]) * 0x13)
	x[17] = uint64(arg1[7]) * (uint64(arg2[9]) * 0x26)
	x[18] = uint64(arg1[7]) * (uint64(arg2[8]) * 0x13)
	x[19] = uint64(arg1[7]) * (uint64(arg2[7]) * 0x26)
	x[20] = uint64(arg1[7]) * (uint64(arg2[6]) * 0x13)
	x[21] = uint64(arg1[7]) * (uint64(arg2[5]) * 0x26)
	x[22] = uint64(arg1[7]) * (uint64(arg2[4]) * 0x13)
	x[23] = uint64(arg1[7]) * (uint64(arg2[3]) * 0x26)
	x[24] = uint64(arg1[6]) * (uint64(arg2[9]) * 0x13)
	x[25] = uint64(arg1[6]) * (uint64(arg2[8]) * 0x13)
	x[26] = uint64(arg1[6]) * (uint64(arg2[7]) * 0x13)
	x[27] = uint64(arg1[6]) * (uint64(arg2[6]) * 0x13)
	x[28] = uint64(arg1[6]) * (uint64(arg2[5]) * 0x13)
	x[29] = uint64(arg1[6]) * (uint64(arg2[4]) * 0x13)
	x[30] = uint64(arg1[5]) * (uint64(arg2[9]) * 0x26)
	x[31] = uint64(arg1[5]) * (uint64(arg2[8]) * 0x13)
	x[32] = uint64(arg1[5]) * (uint64(arg2[7]) * 0x26)
	x[33] = uint64(arg1[5]) * (uint64(arg2[6]) * 0x13)
	x[34] = uint64(arg1[5]) * (uint64(arg2[5]) * 0x26)
	x[35] = uint64(arg1[4]) * (uint64(arg2[9]) * 0x13)
	x[36] = uint64(arg1[4]) * (uint64(arg2[8]) * 0x13)
	x[37] = uint64(arg1[4]) * (uint64(arg2[7]) * 0x13)
	x[38] = uint64(arg1[4]) * (uint64(arg2[6]) * 0x13)
	x[39] = uint64(arg1[3]) * (uint64(arg2[9]) * 0x26)
	x[40] = uint64(arg1[3]) * (uint64(arg2[8]) * 0x13)
	x[41] = uint64(arg1[3]) * (uint64(arg2[7]) * 0x26)
	x[42] = uint64(arg1[2]) * (uint64(arg2[9]) * 0x13)
	x[43] = uint64(arg1[2]) * (uint64(arg2[8]) * 0x13)
	x[44] = uint64(arg1[1]) * (uint64(arg2[9]) * 0x26)
	x[45] = uint64(arg1[9]) * uint64(arg2[0])
	x[46] = uint64(arg1[8]) * uint64(arg2[1])
	x[47] = uint64(arg1[8]) * uint64(arg2[0])
	x[48] = uint64(arg1[7]) * uint64(arg2[2])
	x[49] = uint64(arg1[7]) * (uint64(arg2[1]) * 2)
	x[50] = uint64(arg1[7]) * uint64(arg2[0])
	x[51] = uint64(arg1[6]) * uint64(arg2[3])
	x[52] = uint64(arg1[6]) * uint64(arg2[2])
	x[53] = uint64(arg1[6]) * uint64(arg2[1])
	x[54] = uint64(arg1[6]) * uint64(arg2[0])
	x[55] = uint64(arg1[5]) * uint64(arg2[4])
	x[56] = uint64(arg1[5]) * (uint64(arg2[3]) * 2)
	x[57] = uint64(arg1[5]) * uint64(arg2[2])
	x[58] = uint64(arg1[5]) * (uint64(arg2[1]) * 2)
	x[59] = uint64(arg1[5]) * uint64(arg2[0])
	x[60] = uint64(arg1[4]) * uint64(arg2[5])
	x[61] = uint64(arg1[4]) * uint64(arg2[4])
	x[62] = uint64(arg1[4]) * uint64(arg2[3])
	x[63] = uint64(arg1[4]) * uint64(arg2[2])
	x[64] = uint64(arg1[4]) * uint64(arg2[1])
	x[65] = uint64(arg1[4]) * uint64(arg2[0])
	x[66] = uint64(arg1[3]) * uint64(arg2[6])
	x[67] = uint64(arg1[3]) * (uint64(arg2[5]) * 2)
	x[68] = uint64(arg1[3]) * uint64(arg2[4])
	x[69] = uint64(arg1[3]) * (uint64(arg2[3]) * 2)
	x[70] = uint64(arg1[3]) * uint64(arg2[2])
	x[71] = uint64(arg1[3]) * (uint64(arg2[1]) * 2)
	x[72] = uint64(arg1[3]) * uint64(arg2[0])
	x[73] = uint64(arg1[2]) * uint64(arg2[7])
	x[74] = uint64(arg1[2]) * uint64(arg2[6])
	x[75] = uint64(arg1[2]) * uint64(arg2[5])
	x[76] = uint64(arg1[2]) * uint64(arg2[4])
	x[77] = uint64(arg1[2]) * uint64(arg2[3])
	x[78] = uint64(arg1[2]) * uint64(arg2[2])
	x[79] = uint64(arg1[2]) * uint64(arg2[1])
	x[80] = uint64(arg1[2]) * uint64(arg2[0])
	x[81] = uint64(arg1[1]) * uint64(arg2[8])
	x[82] = uint64(arg1[1]) * (uint64(arg2[7]) * 2)
	x[83] = uint64(arg1[1]) * uint64(arg2[6])
	x[84] = uint64(arg1[1]) * (uint64(arg2[5]) * 2)
	x[85] = uint64(arg1[1]) * uint64(arg2[4])
	x[86] = uint64(arg1[1]) * (uint64(arg2[3]) * 2)
	x[87] = uint64(arg1[1]) * uint64(arg2[2])
	x[88] = uint64(arg1[1]) * (uint64(arg2[1]) * 2)
	x[89] = uint64(arg1[1]) * uint64(arg2[0])
	x[90] = uint64(arg1[0]) * uint64(arg2[9])
	x[91] = uint64(arg1[0]) * uint64(arg2[8])
	x[92] = uint64(arg1[0]) * uint64(arg2[7])
	x[93] = uint64(arg1[0]) * uint64(arg2[6])
	x[94] = uint64(arg1[0]) * uint64(arg2[5])
	x[95] = uint64(arg1[0]) * uint64(arg2[4])
	x[96] = uint64(arg1[0]) * uint64(arg2[3])
	x[97] = uint64(arg1[0]) * uint64(arg2[2])
	x[98] = uint64(arg1[0]) * uint64(arg2[1])
	x[99] = uint64(arg1[0]) * uint64(arg2[0])

	x100 := x[99] + (x[45] + (x[44] + (x[42] + (x[39] + (x[35] + (x[30] + (x[24] + (x[17] + x[9]))))))))
	x102 := x100 >> 26
	x103 = uint32(x100 & 0x3ffffff)

	x104 := x[91] + (x[82] + (x[74] + (x[67] + (x[61] + (x[56] + (x[52] + (x[49] + (x[47] + x[46]))))))))

	x105 := x[92] + (x[83] + (x[75] + (x[68] + (x[62] + (x[57] + (x[53] + (x[50] + (x[48] + x[1]))))))))
	x106 := x[93] + (x[84] + (x[76] + (x[69] + (x[63] + (x[58] + (x[54] + (x[51] + (x[10] + x[2]))))))))
	x107 := x[94] + (x[85] + (x[77] + (x[70] + (x[64] + (x[59] + (x[55] + (x[18] + (x[11] + x[3]))))))))
	x108 := x[95] + (x[86] + (x[78] + (x[71] + (x[65] + (x[60] + (x[25] + (x[19] + (x[12] + x[4]))))))))
	x109 := x[96] + (x[87] + (x[79] + (x[72] + (x[66] + (x[31] + (x[26] + (x[20] + (x[13] + x[5]))))))))
	x110 := x[97] + (x[88] + (x[80] + (x[73] + (x[36] + (x[32] + (x[27] + (x[21] + (x[14] + x[6]))))))))
	x111 := x[98] + (x[89] + (x[81] + (x[40] + (x[37] + (x[33] + (x[28] + (x[22] + (x[15] + x[7]))))))))
	x112 := x[99] + (x[90] + (x[43] + (x[41] + (x[38] + (x[34] + (x[29] + (x[23] + (x[16] + x[8]))))))))

	x113 := x102 + x112
	x114 := x113 >> 25
	x115 := uint32(x113 & 0x1ffffff)

	x116 := (x114 + x111)
	x117 := x116 >> 26
	x118 := uint32(x116 & 0x3ffffff)

	x119 := (x117 + x110)
	x120 := x119 >> 25
	x121 := uint32(x119 & 0x1ffffff)

	x122 := (x120 + x109)
	x123 := x122 >> 26
	x124 := uint32(x122 & 0x3ffffff)

	x125 := (x123 + x108)
	x126 := x125 >> 25
	x127 := uint32(x125 & 0x1ffffff)

	x128 := (x126 + x107)
	x129 := x128 >> 26
	x130 := uint32(x128 & 0x3ffffff)

	x131 := (x129 + x106)
	x132 := x131 >> 25
	x133 := uint32(x131 & 0x1ffffff)

	x134 := (x132 + x105)
	x135 := x134 >> 26
	x136 := uint32(x134 & 0x3ffffff)

	x137 := (x135 + x104)
	x138 := x137 >> 25
	x139 := uint32(x137 & 0x1ffffff)

	x140 := (x138 * 0x13)
	x141 := (x103 + x140)
	x142 := uint32(x141 >> 26)
	x143 := uint32(x141 & 0x3ffffff)

	x144 := (x142 + x115)
	x145 = fiat25519Uint1(x144 >> 25)
	x146 := (x144 & 0x1ffffff)
	x147 := (x145 + x118)

	out1[0] = x143
	out1[1] = x146
	out1[2] = x147
	out1[3] = x121
	out1[4] = x124
	out1[5] = x127
	out1[6] = x130
	out1[7] = x133
	out1[8] = x136
	out1[9] = x139
}

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

func fiat25519Carry(out1 *fiat25519TightFieldElement, arg1 *fiat25519LooseFieldElement) {
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
	x[12] = (fiat25519Uint1(x[11]>>26) + (x[2] & 0x1ffffff))
	x[13] = (x[11] & 0x3ffffff)
	x[14] = (x[12] & 0x1ffffff)
	x[15] = (fiat25519Uint1(x[12]>>25) + (x[3] & 0x3ffffff))
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

func fiat25519Add(out1 *fiat25519LooseFieldElement, arg1, arg2 *fiat25519TightFieldElement) {
	for i := 0; i < 10; i++ {
		(*out1)[i] = (*arg1)[i] + (*arg2)[i]
	}
}

func fiat25519Sub(out1 *fiat25519LooseFieldElement, arg1, arg2 *fiat25519TightFieldElement) {
	const mask1 = 0x7fffffe
	const mask2 = 0x3fffffe
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			(*out1)[i] = mask1 + (*arg1)[i] - (*arg2)[i]
		} else {
			(*out1)[i] = mask2 + (*arg1)[i] - (*arg2)[i]
		}
	}
}

func fiat25519Opp(out1 *fiat25519LooseFieldElement, arg1 *fiat25519TightFieldElement) {
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

func fiat25519ToBytes(out1 *[32]byte, arg1 *fiat25519TightFieldElement) {
	var x [10]uint32
	var carry [10]fiat25519Uint1

	for i := 0; i < 10; i++ {
		x[i] = (*arg1)[i]
	}

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			x[i] = (x[i] - 0x3ffffed) & 0xffffffff
		} else {
			x[i] = (x[i] - 0x1ffffff) & 0xffffffff
		}
	}

	for i := 0; i < 10; i++ {
		binary.LittleEndian.PutUint32(out1[i*4:], x[i])
	}
}

func fiat25519FromBytes(out1 *fiat25519TightFieldElement, arg1 [32]uint8) {
	var x [78]uint32
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

func fiat25519Relax(out1 *fiat25519LooseFieldElement, arg1 *fiat25519TightFieldElement) {
	for i := 0; i < 10; i++ {
		out1[i] = arg1[i]
	}
}

func fiat25519CarryScmul121666(out1 *fiat25519TightFieldElement, arg1 *fiat25519LooseFieldElement) {
	var x [78]uint64
	for i := 0; i < 10; i++ {
		x[i+1] = uint64(0x1db42) * uint64(arg1[i])
	}

	x[11] = uint32(x[10] >> 26)
	x[12] = uint32(x[10] & 0x3ffffff)
	x[13] = x[11] + x[9]
	x[14] = uint32(x[13] >> 25)
	x[15] = uint32(x[13] & 0x1ffffff)
	x[16] = x[14] + x[8]
	x[17] = uint32(x[16] >> 26)
	x[18] = uint32(x[16] & 0x3ffffff)
	x[19] = x[17] + x[7]
	x[20] = uint32(x[19] >> 25)
	x[21] = uint32(x[19] & 0x1ffffff)
	x[22] = x[20] + x[6]
	x[23] = uint32(x[22] >> 26)
	x[24] = uint32(x[22] & 0x3ffffff)
	x[25] = x[23] + x[5]
	x[26] = uint32(x[25] >> 25)
	x[27] = uint32(x[25] & 0x1ffffff)
	x[28] = x[26] + x[4]
	x[29] = uint32(x[28] >> 26)
	x[30] = uint32(x[28] & 0x3ffffff)
	x[31] = x[29] + x[3]
	x[32] = uint32(x[31] >> 25)
	x[33] = uint32(x[31] & 0x1ffffff)
	x[34] = x[32] + x[2]
	x[35] = uint32(x[34] >> 26)
	x[36] = uint32(x[34] & 0x3ffffff)
	x[37] = x[35] + x[1]
	x[38] = uint32(x[37] >> 25)
	x[39] = uint32(x[37] & 0x1ffffff)
	x[40] = x[12] + (x[38] * 0x13)
	x[41] = uint32(x[40] >> 26)
	x[42] = uint32(x[40] & 0x3ffffff)
	x[43] = x[41] + x[15]
	x[44] = uint32(x[43] >> 25)
	x[45] = uint32(x[43] & 0x1ffffff)
	x[46] = x[44] + x[18]

	out1[0] = x[39]
	out1[1] = x[42]
	out1[2] = x[45]
	out1[3] = x[21]
	out1[4] = x[24]
	out1[5] = x[27]
	out1[6] = x[30]
	out1[7] = x[33]
	out1[8] = x[36]
	out1[9] = x[39]
}
