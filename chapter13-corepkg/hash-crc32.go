package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	h := crc32.NewIEEE()
	// the crc32 implements the Writer interface, allowing us to 
	// write bytes to it like any other writer.
	h.Write([]byte("test"))
	// then we can call sum32 to return a unint32 (commonly used)
	// to compare two files.
	v := h.Sum32()
	fmt.Println(v)
}
