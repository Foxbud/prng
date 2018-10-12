package main

import (
	"fmt"
	"time"

	"bitbucket.org/Foxbud/prng/sgen"
)

func main() {
	gen := sgen.NewXS32()
	buf := make([]uint8, 1048576*100)
	//val := uint32(0)
	//pos := 0
	//state := sgen.XS32DefaultSeed
	start := time.Now()
	gen.Read(buf)
	//for i := range buf {
	//	if pos == 0 {
	//		state ^= state << sgen.XS32ParamA
	//		state ^= state >> sgen.XS32ParamB
	//		state ^= state << sgen.XS32ParamC
	//		val = state
	//		pos = 4
	//	}
	//	buf[i] = uint8(val)
	//	val >>= 8
	//	pos--
	//}
	end := time.Now()
	fmt.Printf("%v bytes generated in %v\n", len(buf), end.Sub(start))
}
