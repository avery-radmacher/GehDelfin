package header

import "fmt"

// CurrentVersion is the latest header version, used for all new encryptions.
const CurrentVersion = 1

// Header represents the header data at the beginning of an encrypted file,
// including methods for converting the header to and from buffers. As the
// header is updated, the structure code itself changes.
type Header struct {
	IsComplete    bool
	IsUnsupported bool
	FileSize      int32
	UseOldCipher  bool
	HeaderVersion byte
	HeaderSize    int
	byteScan      byte
}

// NewHeader returns a new header of the most current version.
func NewHeader() Header {
	return Header{
		IsComplete:    false,
		IsUnsupported: false,
		FileSize:      0,
		UseOldCipher:  false,
		HeaderVersion: CurrentVersion,
		HeaderSize:    5,
		byteScan:      0}
}

// AddByte reads a byte as if from a header buffer and updates internal state.
func (h *Header) AddByte(b byte) {
	// quit if unsupported header
	if h.IsUnsupported {
		return
	}

	// byte 0: header version
	if h.byteScan == 0 {
		// assign version and check support
		if h.HeaderVersion = b; h.HeaderVersion > CurrentVersion {
			h.IsUnsupported = true
			return
		}
		h.byteScan = 1 // that is, byteScan++
		return
	}

	if h.HeaderVersion == 0 {
		h.UseOldCipher = true // HV0 uses deprecated short-circuiting Cipherv1
	}

	if h.HeaderVersion == 0 || h.HeaderVersion == 1 {
		// bytes 1–4 (4 bytes): filesize
		if h.byteScan < 5 {
			h.FileSize <<= 8
			h.FileSize |= int32(b)
			h.byteScan++
			if h.byteScan == 5 {
				h.HeaderSize = 5
				h.IsComplete = true
			}
		}
	}

	// new header versions are handled here with new ifs
}

// ToBuffer converts the header to its buffer representation.
func (h Header) ToBuffer() (buffer []byte) {
	// HV0 & HV1:
	//  0: header version
	//	1–4: filesize (int32)
	buffer = make([]byte, 5)
	buffer[0] = h.HeaderVersion
	buffer[1] = byte(h.FileSize >> 24 & 255)
	buffer[2] = byte(h.FileSize >> 16 & 255)
	buffer[3] = byte(h.FileSize >> 8 & 255)
	buffer[4] = byte(h.FileSize & 255)
	return
}

// Test runs a unit test on header
func Test() {
	fmt.Println("header test")
	h1 := NewHeader()
	buffer := []byte{0x02, 0x00, 0x00, 0x01, 0x00}
	for i := 0; !h1.IsComplete && !h1.IsUnsupported; i++ {
		h1.AddByte(buffer[i])
	}
	fmt.Println(h1)
}
