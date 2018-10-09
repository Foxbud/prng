package source

import (
	"encoding/binary"
)

func XorShift16(state uint16) uint16 {
	nextState := state
	nextState ^= nextState << 7
	nextState ^= nextState >> 9
	nextState ^= nextState << 8
	return nextState
}

// Default XorShift16 implementation.

type XS16 struct {
	state uint16
}

func NewXS16() *XS16 {
	return &XS16{0x0001}
}

func (xs *XS16) Seed(seed []uint8) {
	xs.state = binary.LittleEndian.Uint16(seed[:2])
}

func (xs *XS16) Read(buff []uint8) (int, error) {
	buffLen := len(buff)
	for i := 0; i < buffLen; i += 2 {
		xs.state = XorShift16(xs.state)
		UnpackUint(buff[i:], xs.state)
	}
	return buffLen, nil
}

// Half XorShift16 implementation.

type HalfXS16 struct {
	state uint16
}

func NewHalfXS16() *HalfXS16 {
	return &HalfXS16{0x0001}
}

func (xs *HalfXS16) Seed(seed []uint8) {
	xs.state = binary.LittleEndian.Uint16(seed[:2])
}

func (xs *HalfXS16) Read(buff []uint8) (int, error) {
	buffLen := len(buff)
	for i := 0; i < buffLen; i++ {
		xs.state = XorShift16(xs.state)
		buff[i] = uint8(xs.state)
	}
	return buffLen, nil
}
