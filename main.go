package main

import (
	"fmt"
	"math/rand"
	"time"

	"bitbucket.org/Foxbud/prng/sgen"
)

func main() {
	gen := rand.New(sgen.NewSource(sgen.NewXS64()))
	buf := make([]uint8, 1048576*100)
	start := time.Now()
	gen.Read(buf)
	end := time.Now()
	fmt.Printf("%v bytes generated in %v\n", len(buf), end.Sub(start))
}
