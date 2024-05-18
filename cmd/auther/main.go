package main

import (
	"github.com/leitiannet/template/auth"
	"github.com/leitiannet/template/internal/trace"
)

func main() {
	trace.Printf("this is auther using auth\n")
	auth.Auth("a", "b", "c")
}
