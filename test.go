package main

import (
	//"bytes"
	//"encoding/binary"
	"fmt"
	//"math/rand"
	"bitbucket.org/Foxbud/prng/source"
)

func main() {
	buf := make([]uint8, 1048)
	src := source.NewXS64Star()
	for {
		src.Read(buf)
		fmt.Printf("%v\n", buf)
	}
}
