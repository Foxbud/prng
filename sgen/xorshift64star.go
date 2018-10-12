package sgen

import (
	"encoding/binary"
)

const (
	// Defaults recommended by Vigna in "An experimental exploration of
	// Marsaglia’s xorshift generators, scrambled".
	XS64StarParamA uint8  = 12
	XS64StarParamB        = 25
	XS64StarParamC        = 27
	XS64StarCoef   uint64 = 0x2545f4914f6cdd1d
)

type XS64Star struct {
	state uint64
	buf   []uint8
}

func NewXS64Star() *XS64Star {
	return &XS64Star{XS64DefaultSeed, make([]uint8, 0, 8)}
}

func (xs *XS64Star) Seed(seed []uint8) {
	xs.state = binary.LittleEndian.Uint64(seed[:8])
	xs.buf = xs.buf[:0]
}

func (xs *XS64Star) Read(ebuf []uint8) (int, error) {
	for i := range ebuf {
		if len(xs.buf) == 0 {
			xs.state ^= xs.state << XS64ParamA
			xs.state ^= xs.state >> XS64ParamB
			xs.state ^= xs.state << XS64ParamC
			xs.buf = xs.buf[:cap(xs.buf)]
			binary.BigEndian.PutUint64(xs.buf, xs.state*XS64StarCoef)
		}
		end := len(xs.buf) - 1
		ebuf[i] = xs.buf[end]
		xs.buf = xs.buf[:end]
	}
	return len(ebuf), nil
}
