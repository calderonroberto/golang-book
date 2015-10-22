package main

import (
	"crypto/sha1"
	"fmt"
)

//  The main difference is that whereas crc32 computes a 32 bit hash,
//  sha1 computes a 160 bit hash. There is no native type to represent 
// a 160 bit number, so we use a slice of 20 bytes instead.

func main() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
