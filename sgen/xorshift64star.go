package sgen

const (
	// Defaults recommended by Vigna in "An experimental exploration of
	// Marsaglia’s xorshift generators, scrambled".
	XS64StarParamA uint8  = 12
	XS64StarParamB        = 25
	XS64StarParamC        = 27
	XS64StarCoef   uint64 = 0x2545f4914f6cdd1d
)

type XS64Star struct {
	state uint64
	buf   uint64
	rem   uint8
}

func NewXS64Star() *XS64Star {
	return &XS64Star{XS64DefaultSeed, 0, 0}
}

func (xs *XS64Star) Seed(seed []uint8) {
	xs.rem = 0
	val := uint64(seed[7])
	for i := 6; i >= 0; i-- {
		val <<= 8
		val ^= uint64(seed[i])
	}
	xs.state = val
}

func (xs *XS64Star) Read(buf []uint8) (int, error) {
	lState := xs.state
	lBuf := xs.buf
	lRem := xs.rem
	for i := range buf {
		if lRem == 0 {
			lState ^= lState << XS64StarParamA
			lState ^= lState >> XS64StarParamB
			lState ^= lState << XS64StarParamC
			lBuf = lState * XS64StarCoef
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
