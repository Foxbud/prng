package source

import (
	"encoding/binary"
)

func XorShift32(state uint32) uint32 {
	nextState := state
	nextState ^= nextState << 5
	nextState ^= nextState >> 17
	nextState ^= nextState << 13
	return nextState
}

// Default XorShift32 implementation.

type XS32 struct {
	state uint32
}

func NewXS32() *XS32 {
	return &XS32{0x00000001}
}

func (xs *XS32) Seed(seed []uint8) {
	xs.state = binary.LittleEndian.Uint32(seed[:4])
}

func (xs *XS32) Read(buff []uint8) (int, error) {
	buffLen := len(buff)
	for i := 0; i < buffLen; i += 4 {
		xs.state = XorShift32(xs.state)
		UnpackUint(buff[i:], xs.state)
	}
	return buffLen, nil
}

// Half XorShift32 implementation.

type HalfXS32 struct {
	state uint32
}

func NewHalfXS32() *HalfXS32 {
	return &HalfXS32{0x00000001}
}

func (xs *HalfXS32) Seed(seed []uint8) {
	xs.state = binary.LittleEndian.Uint32(seed[:4])
}

func (xs *HalfXS32) Read(buff []uint8) (int, error) {
	buffLen := len(buff)
	for i := 0; i < buffLen; i += 2 {
		xs.state = XorShift32(xs.state)
		UnpackUint(buff[i:], uint16(xs.state>>16))
	}
	return buffLen, nil
}
