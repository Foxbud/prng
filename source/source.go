package source

import (
	"encoding/binary"
	"io"
)

type Source interface {
	io.Reader
	Seed([]uint8)
}

type STDSource struct {
	src Source
}

func NewSTDSource(src Source) *STDSource {
	return &STDSource{src}
}

func (s *STDSource) Seed(seed int64) {
	buf := make([]uint8, 8)
	binary.LittleEndian.PutUint64(buf, uint64(seed))
	s.src.Seed(buf)
}

func (s *STDSource) Int63() int64 {
	return int64(s.Uint64() >> 1)
}

func (s *STDSource) Uint64() uint64 {
	buf := make([]uint8, 8)
	s.src.Read(buf)
	return binary.LittleEndian.Uint64(buf)
}
