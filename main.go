package main

import "fmt"

func main() {
	x := uint64(0)
	m := uint64(11 * 19)
	buf := uint8(0)
	for i := 0; i < 1; i++ {
		buf = 0
		for j := 0; j < 8; j++ {
			x = blumblumshub(x, m)
			fmt.Printf("%v\n", x)
			buf ^= uint8(x) & 1
			buf <<= 1
		}
		//fmt.Printf("%v\n", buf)
	}
}

func blumblumshub(x, m uint64) uint64 {
	return (x * x) % m
}
