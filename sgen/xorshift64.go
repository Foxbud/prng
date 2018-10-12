package sgen

const (
	// Defaults recommended by Marsaglia in "Xorshift RNGs".
	XS64DefaultSeed uint64 = 0x139408dcbbf7a44
	XS64ParamA      uint8  = 13
	XS64ParamB             = 7
	XS64ParamC             = 17
)

type XS64 struct {
	state uint64
	buf   uint64
	rem   uint8
}

func NewXS64() *XS64 {
	return &XS64{XS64DefaultSeed, 0, 0}
}

func (xs *XS64) Seed(seed []uint8) {
	xs.rem = 0
	val := uint64(seed[7])
	for i := 6; i >= 0; i-- {
		val <<= 8
		val ^= uint64(seed[i])
	}
	xs.state = val
}

func (xs *XS64) Read(buf []uint8) (int, error) {
	lState := xs.state
	lBuf := xs.buf
	lRem := xs.rem
	for i := range buf {
		if lRem == 0 {
			lState ^= lState << XS64ParamA
			lState ^= lState >> XS64ParamB
			lState ^= lState << XS64ParamC
			lBuf = lState
			lRem = 8
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
