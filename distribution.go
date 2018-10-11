package main

import (
	"encoding/binary"
	"math/bits"

	"bitbucket.org/Foxbud/prng/source"
)

type SourceMapper struct {
	src source.Source
	buf []uint8
}

func NewSourceMapper(src source.Source) *SourceMapper {
	return &SourceMapper{src, make([]uint8, 8, 8)}
}

func (sm *SourceMapper) Seed(seed uint64) {
	binary.LittleEndian.PutUint64(sm.buf, seed)
	sm.src.Seed(sm.buf)
}

func (sm *SourceMapper) Read(ebuf []uint8) (int, error) {
	n, err := sm.src.Read(ebuf)
	return n, err
}

func (sm *SourceMapper) Uint64(bound uint64) uint64 {
	nBits := bits.Len64(bound)
	nBytes := (nBits + 7) >> 3
	for i := range sm.buf[nBytes:] {
		sm.buf[i] = 0
	}
	val := bound
	for val >= bound {
		sm.src.Read(sm.buf[:nBytes])
		val = binary.LittleEndian.Uint64(sm.buf[:8])
	}
	return val
}
