package hash

import (
	"github.com/leitiannet/template/internal/trace"
)

func Hash(a ...interface{}) string {
	trace.Printf("Hash: %v\n", a)
	return ""
}
