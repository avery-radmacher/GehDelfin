package cipher

import "fmt"

// Cipherv2 represents a stream cipher with a 128-bit seed and a very long period.
type Cipherv2 struct {
	LFSRs  [5]uint32
	SRTaps [5]uint32
}

// NewCipherv2 creates a new cipher from a 16-byte seed.
func NewCipherv2(password string) Cipherv2 {
	initVector := vector(password, 16)
	var initInts [16]uint32
	for index, item := range initVector {
		initInts[index] = uint32(item)
	}
	LFSRs, SRTaps := [5]uint32{}, [5]uint32{}
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
	return Cipherv2{LFSRs, SRTaps}
}

// GetByte computes and returns the next 8 bits in the stream.
func (c *Cipherv2) GetByte() (result byte) {
	for i := 0; i < 8; i++ {
		result = (result << 1) | c.tick()
	}
	return
}

// tick returns the next bit in the stream as a byte
func (c *Cipherv2) tick() byte {
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

// String represents the Cipher as a string.
func (c Cipherv2) String() (result string) {
	result += "Cipherv2 object:\n"
	for index, item := range c.LFSRs {
		result += fmt.Sprintf("Register %d: %6X\n", index, item)
	}
	for index, item := range c.SRTaps {
		result += fmt.Sprintf("Tapcode %d:  %6X\n", index, item)
	}
	return
}
