package main

// #include <string.h>
// #include <stdlib.h>
// #include "sbci-core.h"
import "C"

import (
	//"bytes"
	"encoding/binary"
	"fmt"
	//"math/rand"
	"unsafe"
	//"bitbucket.org/Foxbud/prng/source"
)

func main() {
	buf := make([]uint8, 1048)
	t := NewTester(len(buf))
	end := 1048576 * 10
	for i := 0; i < end; i++ {
		t.Read(buf)
		fmt.Printf("%s", buf)
	}
}

type Tester struct {
	cMask     *C.uint8_t
	cPrevious *C.uint8_t
	cCurrent  *C.uint8_t
	cBuf      unsafe.Pointer
}

func NewTester(bufLen int) *Tester {
	t := Tester{}
	seed := uint64(0x73)
	buf := make([]uint8, 8)
	t.cBuf = C.CBytes(make([]uint8, bufLen))
	t.cCurrent = (*C.uint8_t)(C.CBytes(buf))

	binary.LittleEndian.PutUint64(buf, seed)
	t.cMask = (*C.uint8_t)(C.CBytes(buf))
	t.cPrevious = (*C.uint8_t)(C.CBytes(buf))
	C.ISXEngine(t.cMask, t.cPrevious, t.cCurrent)

	return &t
}

func (t *Tester) Read(buf []uint8) (int, error) {
	C.ISXEngine(t.cMask, t.cCurrent, (*C.uint8_t)(t.cBuf))
	for i := 8; i < len(buf); i += 8 {
		//C.memcpy(unsafe.Pointer(t.cPrevious), unsafe.Pointer(t.cCurrent), C.size_t(8))
		//C.ISXEngine(t.cMask, t.cPrevious, t.cCurrent)
		//copy(buf[i:], C.GoBytes(unsafe.Pointer(t.cCurrent), C.int(8)))
		C.ISXEngine(
			t.cMask,
			(*C.uint8_t)(unsafe.Pointer(uintptr(t.cBuf)+uintptr(i-8))),
			(*C.uint8_t)(unsafe.Pointer(uintptr(t.cBuf)+uintptr(i))),
		)
	}
	C.memcpy(unsafe.Pointer(t.cCurrent), unsafe.Pointer(uintptr(t.cBuf)+uintptr(len(buf)-8)), C.size_t(8))
	copy(buf, C.GoBytes(unsafe.Pointer(t.cBuf), C.int(len(buf))))
	return len(buf), nil
}
