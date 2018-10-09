package main

import (
	//"fmt"
	//"math/rand"

	"bitbucket.org/Foxbud/prng/source"
)

func main() {
	src := source.NewXS32()
	//buff := [1024]uint8{}
	//for {
	//	src.Read(buff[:])
	//	fmt.Printf("%s", buff)
	//}

	ImageProfile(src, "profile.png", 1024, 1024)
}
