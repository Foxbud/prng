package main

import (
	"fmt"
	"time"

	"bitbucket.org/Foxbud/prng/sgen"
)

func main() {
	gen := sgen.NewXS64Star()
	//rng := rand.New(sgen.NewSource(gen))
	buf := make([]uint8, 1024*1024*100)
	gen.Read(buf)
	start := time.Now()
	for i := 0; i < 1024*1024*100/8; i++ {
		gen.Seed(buf[i : i+8])
	}
	end := time.Now()
	fmt.Printf("%v\n", end.Sub(start))
}
