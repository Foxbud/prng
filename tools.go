package main

import (
	"encoding/binary"
	"math/bits"

	"github.com/Foxbud/prng/sgen"
)

type GenTools struct {
	gen sgen.Generator
	buf []uint8
}

func NewGenTools(gen sgen.Generator) *GenTools {
	return &GenTools{gen, make([]uint8, 8, 8)}
}

func (gt *GenTools) Seed(seed uint64) {
	binary.LittleEndian.PutUint64(gt.buf, seed)
	gt.gen.Seed(gt.buf)
}

func (gt *GenTools) Read(ebuf []uint8) (int, error) {
	n, err := gt.gen.Read(ebuf)
	return n, err
}

func (gt *GenTools) Uint64(bound uint64) uint64 {
	nBits := bits.Len64(bound)
	nBytes := (nBits + 7) >> 3
	for i := range gt.buf[nBytes:] {
		gt.buf[i] = 0
	}
	val := bound
	for val >= bound {
		gt.gen.Read(gt.buf[:nBytes])
		val = binary.LittleEndian.Uint64(gt.buf[:8])
	}
	return val
}
