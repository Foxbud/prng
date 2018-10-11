package main

import (
	"encoding/binary"

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
	nBytes := 8
	switch {
	case bound <= 0x100:
		nBytes = 1
	case bound <= 0x10000:
		nBytes = 2
	case bound <= 0x1000000:
		nBytes = 3
	case bound <= 0x100000000:
		nBytes = 4
	case bound <= 0x10000000000:
		nBytes = 5
	case bound <= 0x1000000000000:
		nBytes = 6
	case bound <= 0x100000000000000:
		nBytes = 7
	}
	sm.src.Read(sm.buf[:nBytes-1])
	for i := range sm.buf[nBytes:8] {
		sm.buf[i] = 0
	}
	val := bound
	for val >= bound {
		sm.src.Read(sm.buf[nBytes-1 : nBytes])
		val = binary.LittleEndian.Uint64(sm.buf[:8])
	}
	return val
}
