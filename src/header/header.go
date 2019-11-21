package header

import "fmt"

// CurrentVersion is the latest header version, used for all new encryptions.
const CurrentVersion = 1

// Header represents the header data at the beginning of an encrypted file,
// including methods for converting the header to and from buffers.
type Header interface {
	AddByte(b byte)
	ToBuffer() []byte
}

type header struct {
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
	return &header{
		IsComplete:    false,
		IsUnsupported: false,
		UseOldCipher:  false,
		HeaderVersion: CurrentVersion,
		byteScan:      0,
		HeaderSize:    5}
}

func (h *header) AddByte(b byte) {
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

func (h header) ToBuffer() (buffer []byte) {
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

func Test() {
	fmt.Println("header test")
	h1, h2 := NewHeader(), NewHeader()
	h1.(*header).IsComplete = true
	fmt.Printf("%t : %t\n", h1.(*header).IsComplete, h2.(*header).IsComplete)
}
