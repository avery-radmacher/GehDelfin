package dcipher

import "fmt"

// Cipher represents a stream cipher with a 128-bit seed and a very long period.
type Cipher struct {
	LFSRs  [5]uint32
	SRTaps [5]uint32
}

var tapCodes = [5][64]uint32{
	{0x10118,
		0x1011B,
		0x10122,
		0x10135,
		0x1013C,
		0x1013F,
		0x10141,
		0x10148,
		0x1015A,
		0x10163,
		0x1016F,
		0x10171,
		0x10174,
		0x1017E,
		0x10184,
		0x10188,
		0x1018B,
		0x1018D,
		0x10193,
		0x10199,
		0x1019A,
		0x101A9,
		0x101BB,
		0x101D8,
		0x101DB,
		0x101E1,
		0x101E8,
		0x101ED,
		0x101F5,
		0x10203,
		0x1020A,
		0x1020C,
		0x1020F,
		0x10217,
		0x1021E,
		0x10221,
		0x1022B,
		0x1022E,
		0x10230,
		0x10233,
		0x1023A,
		0x10242,
		0x10247,
		0x1024E,
		0x10255,
		0x1025C,
		0x10260,
		0x10266,
		0x1027B,
		0x1027D,
		0x1028D,
		0x1028E,
		0x10293,
		0x1029A,
		0x1029F,
		0x102B1,
		0x102B2,
		0x102B7,
		0x102C9,
		0x102D2,
		0x102DD,
		0x102F6,
		0x10302,
		0x1030E},
	{0x2023C,
		0x20244,
		0x20250,
		0x20290,
		0x202C9,
		0x202CF,
		0x202E2,
		0x202ED,
		0x202F0,
		0x202FF,
		0x20304,
		0x2030E,
		0x20315,
		0x2033E,
		0x20346,
		0x20394,
		0x20398,
		0x203A8,
		0x203DF,
		0x203E6,
		0x203F8,
		0x20400,
		0x2041D,
		0x20439,
		0x20442,
		0x20465,
		0x2046F,
		0x20477,
		0x2047E,
		0x20482,
		0x20493,
		0x20496,
		0x204B2,
		0x204BD,
		0x204C9,
		0x204D1,
		0x204D2,
		0x204E4,
		0x204E7,
		0x2050E,
		0x20545,
		0x2054A,
		0x20562,
		0x20567,
		0x2056B,
		0x20570,
		0x2057A,
		0x2058C,
		0x20594,
		0x2059D,
		0x205A8,
		0x205B3,
		0x205C1,
		0x205E3,
		0x205EC,
		0x205F8,
		0x20625,
		0x20638,
		0x2063D,
		0x20676,
		0x2067F,
		0x206C2,
		0x206CB,
		0x206CD},
	{0x40150,
		0x40190,
		0x401A3,
		0x401B4,
		0x401B7,
		0x401CF,
		0x401D2,
		0x401D4,
		0x401E4,
		0x401EB,
		0x401ED,
		0x401EE,
		0x401F9,
		0x4021B,
		0x40228,
		0x4023A,
		0x4023F,
		0x40244,
		0x40248,
		0x4025A,
		0x4025C,
		0x40269,
		0x4026C,
		0x4028E,
		0x40295,
		0x4029C,
		0x402A0,
		0x402A9,
		0x402AF,
		0x402B8,
		0x402BB,
		0x402D4,
		0x402DD,
		0x402E8,
		0x402EB,
		0x402EE,
		0x402F5,
		0x402FA,
		0x402FC,
		0x40304,
		0x40308,
		0x40319,
		0x40320,
		0x40325,
		0x40326,
		0x4032F,
		0x40332,
		0x40343,
		0x40345,
		0x4035E,
		0x40361,
		0x40368,
		0x40373,
		0x40376,
		0x4038A,
		0x4039D,
		0x403A2,
		0x403A4,
		0x403AB,
		0x403C2,
		0x403C8,
		0x403CE,
		0x403F7,
		0x403FE},
	{0x1001B1,
		0x1001BB,
		0x1001CF,
		0x1001D1,
		0x1001D7,
		0x1001DD,
		0x1001E2,
		0x1001EE,
		0x10020A,
		0x10020C,
		0x100211,
		0x100224,
		0x10022E,
		0x10023A,
		0x100248,
		0x10024B,
		0x100253,
		0x100272,
		0x100274,
		0x100277,
		0x10027E,
		0x100287,
		0x100295,
		0x1002AA,
		0x1002DB,
		0x1002EB,
		0x1002F0,
		0x1002F5,
		0x1002F6,
		0x100308,
		0x10031F,
		0x100329,
		0x100345,
		0x10035D,
		0x100361,
		0x10036B,
		0x100370,
		0x10037F,
		0x1003A1,
		0x1003AE,
		0x1003B5,
		0x1003BA,
		0x1003CB,
		0x1003D5,
		0x1003D9,
		0x1003E0,
		0x1003E6,
		0x1003E9,
		0x100406,
		0x100409,
		0x10040A,
		0x100427,
		0x100430,
		0x100441,
		0x100448,
		0x10044B,
		0x100455,
		0x10045C,
		0x10046F,
		0x100478,
		0x10047D,
		0x10049A,
		0x1004AC,
		0x1004B2},
	{0x40016F,
		0x400177,
		0x400181,
		0x400193,
		0x400196,
		0x4001A0,
		0x4001B2,
		0x4001D1,
		0x4001D2,
		0x4001D8,
		0x4001DE,
		0x4001F6,
		0x400206,
		0x400214,
		0x400228,
		0x400247,
		0x400260,
		0x40026F,
		0x400293,
		0x400295,
		0x4002A0,
		0x4002B7,
		0x4002D2,
		0x4002D4,
		0x4002E2,
		0x4002F0,
		0x400308,
		0x40030E,
		0x40032F,
		0x400331,
		0x400340,
		0x40035D,
		0x400362,
		0x40036D,
		0x40036E,
		0x400376,
		0x400379,
		0x400389,
		0x40038A,
		0x400391,
		0x40039B,
		0x4003BF,
		0x4003D0,
		0x4003DA,
		0x4003EF,
		0x4003F2,
		0x4003FB,
		0x4003FD,
		0x40041D,
		0x40041E,
		0x40042D,
		0x40043F,
		0x400444,
		0x40044D,
		0x400455,
		0x400456,
		0x400460,
		0x400463,
		0x40046A,
		0x400484,
		0x400490,
		0x40049C,
		0x4004AA,
		0x4004B2}}

// NewCipher creates a new cipher from a 16-byte seed.
func NewCipher(initVector [16]byte) Cipher {
	LFSRs, SRTaps := [5]uint32{}, [5]uint32{}
	var initInts [16]uint32
	for index, item := range initVector {
		initInts[index] = uint32(item)
	}
	LFSRs[0] = uint32((initInts[0] << 9) | (initInts[1] << 1) | (initInts[2] >> 7))
	LFSRs[1] = uint32(((initInts[2] & 127) << 11) | (initInts[3] << 3) | (initInts[4] >> 5))
	LFSRs[2] = uint32(((initInts[4] & 7) << 14) | (initInts[5] << 6) | (initInts[6] >> 2)) // TODO why not ... & 31 ? ignoring two bits.
	LFSRs[3] = uint32(((initInts[6] & 3) << 19) | (initInts[7] << 11) | (initInts[8] << 3) | (initInts[9] >> 5))
	LFSRs[4] = uint32(((initInts[9] & 31) << 18) | (initInts[10] << 10) | (initInts[11] << 2) | (initInts[12] >> 6))
	SRTaps[0] = tapCodes[0][initVector[12]&63]
	SRTaps[1] = tapCodes[1][initVector[13]>>2]
	SRTaps[2] = tapCodes[2][(initVector[13]&3)<<4|(initVector[14]>>4)]
	SRTaps[3] = tapCodes[3][(initVector[14]&15)<<2|(initVector[15]>>6)]
	SRTaps[4] = tapCodes[4][initVector[15]&63]

	for i := 0; i < 5; i++ {
		if LFSRs[i] == 0 {
			LFSRs[i] = 1
		}
	}

	return Cipher{LFSRs, SRTaps}
}

// GetByte computes and returns the next 8 bits in the stream.
func (c *Cipher) GetByte() (result byte) {
	for i := 0; i < 8; i++ {
		result = (result << 1) | c.tick()
	}
	return
}

// Tick returns the next bit in the stream as a byte
func (c *Cipher) tick() byte {
	num16sPlaceOnes, majorityBit := uint32(0), uint32(0)
	for i := 0; i < 5; i++ {
		num16sPlaceOnes += c.LFSRs[i] & 16
	}
	if num16sPlaceOnes > 32 {
		majorityBit = 16 // majority bit in its place (10000₂)
	}

	for i := 0; i < 5; i++ {
		if c.LFSRs[i]&16 == majorityBit {
			if c.LFSRs[i]&1 == 1 {
				c.LFSRs[i] = (c.LFSRs[i] >> 1) ^ c.SRTaps[i]
			} else {
				c.LFSRs[i] = c.LFSRs[i] >> 1
			}
		}
	}

	return byte(c.LFSRs[0]^c.LFSRs[1]^c.LFSRs[2]^c.LFSRs[3]^c.LFSRs[4]) & 1
}

// String relates Cipher data to the terminal.
func (c Cipher) String() (result string) {
	for index, item := range c.LFSRs {
		result += fmt.Sprintf("Register %d: %6X\n", index, item)
	}
	for index, item := range c.SRTaps {
		result += fmt.Sprintf("Tapcode %d:  %6X\n", index, item)
	}
	return
}
