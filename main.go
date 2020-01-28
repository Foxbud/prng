package main

import (
	"fmt"
	"math/rand"

	"github.com/Foxbud/prng/sgen"
)

func main() {
	gen := rand.New(sgen.NewSource(sgen.NewXS32()))
	var buf [16]byte
	gen.Read(buf[:])
	fmt.Printf("%v\n", buf)
}
