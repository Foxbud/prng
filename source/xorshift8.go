package source

func XorShift8(state uint8) uint8 {
	nextState := state
	nextState ^= nextState << 3
	nextState ^= nextState >> 5
	nextState ^= nextState << 4
	return nextState
}

// Default XorShift8 implementation.

type XS8 struct {
	state uint8
}

func NewXS8() *XS8 {
	return &XS8{0x01}
}

func (xs *XS8) Seed(seed []uint8) {
	xs.state = seed[0]
}

func (xs *XS8) Read(buff []uint8) (int, error) {
	buffLen := len(buff)
	for i := 0; i < buffLen; i++ {
		xs.state = XorShift8(xs.state)
		buff[i] = xs.state
	}
	return buffLen, nil
}
