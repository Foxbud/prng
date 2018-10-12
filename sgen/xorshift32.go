package sgen

const (
	// Defaults recommended by Marsaglia in "Xorshift RNGs".
	XS32DefaultSeed uint32 = 0x92d68ca2
	XS32ParamA      uint8  = 5
	XS32ParamB             = 17
	XS32ParamC             = 13
)

type XS32 struct {
	state uint32
	buf   uint32
	rem   uint8
}

func NewXS32() *XS32 {
	return &XS32{XS32DefaultSeed, 0, 0}
}

func (xs *XS32) Seed(seed []uint8) {
	xs.rem = 0
	val := uint32(seed[3])
	for i := 2; i >= 0; i-- {
		val <<= 8
		val ^= uint32(seed[i])
	}
	xs.state = val
}

func (xs *XS32) Read(buf []uint8) (int, error) {
	lState := xs.state
	lBuf := xs.buf
	lRem := xs.rem
	for i := range buf {
		if lRem == 0 {
			lState ^= lState << XS32ParamA
			lState ^= lState >> XS32ParamB
			lState ^= lState << XS32ParamC
			lBuf = lState
			lRem = 4
		}
		buf[i] = uint8(lBuf)
		lBuf >>= 8
		lRem--
	}
	xs.state = lState
	xs.buf = lBuf
	xs.rem = lRem
	return len(buf), nil
}
