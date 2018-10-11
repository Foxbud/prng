package source

import (
	"encoding/binary"
)

const (
	// Defaults recommended by Marsaglia in "Xorshift RNGs".
	XS32DefaultSeed uint32 = 0x92d68ca2
	XS32ParamA      uint8  = 5
	XS32ParamB             = 17
	XS32ParamC             = 13
)

type XS32 struct {
	state uint32
	buf   []uint8
}

func NewXS32() *XS32 {
	return &XS32{XS32DefaultSeed, make([]uint8, 0, 4)}
}

func (xs *XS32) Seed(seed []uint8) {
	xs.state = binary.LittleEndian.Uint32(seed[:4])
	xs.buf = xs.buf[:0]
}

func (xs *XS32) Read(ebuf []uint8) (int, error) {
	for i := range ebuf {
		if len(xs.buf) == 0 {
			xs.state ^= xs.state << XS32ParamA
			xs.state ^= xs.state >> XS32ParamB
			xs.state ^= xs.state << XS32ParamC
			xs.buf = xs.buf[:cap(xs.buf)]
			binary.BigEndian.PutUint32(xs.buf, xs.state)
		}
		end := len(xs.buf) - 1
		ebuf[i] = xs.buf[end]
		xs.buf = xs.buf[:end]
	}
	return len(ebuf), nil
}
