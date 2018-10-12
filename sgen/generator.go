package sgen

import (
	"encoding/binary"
	"io"
)

type Generator interface {
	io.Reader
	Seed([]uint8)
}

type Source struct {
	gen Generator
}

func NewSource(gen Generator) *Source {
	return &Source{gen}
}

func (s *Source) Seed(seed int64) {
	buf := make([]uint8, 8)
	binary.LittleEndian.PutUint64(buf, uint64(seed))
	s.gen.Seed(buf)
}

func (s *Source) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

func (s *Source) Uint64() uint64 {
	buf := make([]uint8, 8)
	s.gen.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}
