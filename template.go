package template

import (
	"github.com/leitiannet/template/internal/trace"
)

func Add(a, b int) int {
	trace.Printf("Add: %v, %v\n", a, b)
	return a + b
}
