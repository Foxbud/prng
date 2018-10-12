package main

import (
	"fmt"
	"math/rand"
	"time"

	"bitbucket.org/Foxbud/prng/sgen"
)

func main() {
	gen := sgen.NewXS64Star()
	rng := rand.New(sgen.NewSource(gen))
	buf := make([]uint8, 1024*1024*100)
	start := time.Now()
	rng.Read(buf)
	end := time.Now()
	fmt.Printf("%v\n", end.Sub(start))
}
