package auth

import (
	"github.com/leitiannet/template/internal/trace"
)

func Auth(a ...interface{}) error {
	trace.Printf("Auth: %v\n", a)
	return nil
}
