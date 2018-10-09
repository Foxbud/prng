package source

import (
	"encoding/binary"
)

func XorShift64(state uint64) uint64 {
	nextState := state
	nextState ^= nextState << 13
	nextState ^= nextState >> 7
	nextState ^= nextState << 17
	return nextState
}

// Default XorShift64 implementation.

type XS64 struct {
	state uint64
}

func NewXS64() *XS64 {
	return &XS64{0x0000000000000001}
}

func (xs *XS64) Seed(seed []uint8) {
	xs.state = binary.LittleEndian.Uint64(seed[:8])
}

func (xs *XS64) Read(buff []uint8) (int, error) {
	buffLen := len(buff)
	for i := 0; i < buffLen; i += 8 {
		xs.state = XorShift64(xs.state)
		UnpackUint(buff[i:], xs.state)
	}
	return buffLen, nil
}

// Half XorShift64 implementation.

type HalfXS64 struct {
	state uint64
}

func NewHalfXS64() *HalfXS64 {
	return &HalfXS64{0x0000000000000001}
}

func (xs *HalfXS64) Seed(seed []uint8) {
	xs.state = binary.LittleEndian.Uint64(seed[:8])
}

func (xs *HalfXS64) Read(buff []uint8) (int, error) {
	buffLen := len(buff)
	for i := 0; i < buffLen; i += 4 {
		xs.state = XorShift64(xs.state)
		UnpackUint(buff[i:], uint32(xs.state>>32))
	}
	return buffLen, nil
}
