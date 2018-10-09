package source

func UnpackUint(buff []uint8, item interface{}) {
	var accum uint64
	var cNum int
	switch val := item.(type) {
	case uint16:
		accum = uint64(val)
		cNum = 2
	case uint32:
		accum = uint64(val)
		cNum = 4
	case uint64:
		accum = uint64(val)
		cNum = 8
	}
	buffLen := len(buff)
	if cNum > buffLen {
		cNum = buffLen
	}
	for i := 0; i < cNum; i++ {
		buff[i] = uint8(accum)
		accum >>= 8
	}
}
