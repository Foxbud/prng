package source

import (
	"encoding/binary"
)

const (
	// Defaults recommended by Marsaglia in "Xorshift RNGs".
	XS64DefaultSeed uint64 = 0x139408dcbbf7a44
	XS64ParamA      uint8  = 13
	XS64ParamB             = 7
	XS64ParamC             = 17
)

type XS64 struct {
	state uint64
	buf   []uint8
}

func NewXS64() *XS64 {
	return &XS64{XS64DefaultSeed, make([]uint8, 0, 8)}
}

func (xs *XS64) Seed(seed []uint8) {
	xs.state = binary.LittleEndian.Uint64(seed[:8])
	xs.buf = xs.buf[:0]
}

func (xs *XS64) Read(ebuf []uint8) (int, error) {
	for i := range ebuf {
		if len(xs.buf) == 0 {
			xs.state ^= xs.state << XS64ParamA
			xs.state ^= xs.state >> XS64ParamB
			xs.state ^= xs.state << XS64ParamC
			xs.buf = xs.buf[:cap(xs.buf)]
			binary.BigEndian.PutUint64(xs.buf, xs.state)
		}
		end := len(xs.buf) - 1
		ebuf[i] = xs.buf[end]
		xs.buf = xs.buf[:end]
	}
	return len(ebuf), nil
}
