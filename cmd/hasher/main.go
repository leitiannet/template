package main

import (
	"github.com/leitiannet/template/hash"
	"github.com/leitiannet/template/internal/trace"
)

func main() {
	trace.Printf("this is hasher using hash\n")
	hash.Hash("a", "b", "c")
}
