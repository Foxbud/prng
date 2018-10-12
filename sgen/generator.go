package sgen

import (
	"io"
)

type Generator interface {
	io.Reader
	Seed([]uint8)
}

type Source struct {
	gen Generator
	buf []uint8
}

func NewSource(gen Generator) *Source {
	return &Source{gen, make([]uint8, 8)}
}

func (s *Source) Seed(seed int64) {
	lBuf := s.buf
	val := uint64(seed)
	for i := range lBuf {
		lBuf[i] = uint8(val)
		val >>= 8
	}
}

func (s *Source) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s *Source) Uint64() uint64 {
	lBuf := s.buf
	s.gen.Read(lBuf)
	val := uint64(lBuf[7])
	for i := 6; i >= 0; i-- {
		val <<= 8
		val ^= uint64(lBuf[i])
	}
	return val
}
