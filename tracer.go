package trace

import (
	"fmt"
	"io"
)

// a struct type
type tracer struct {
	out io.Writer
}

// an interface type
type Tracer interface {
	Trace(...interface{})
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)

}
func New(w io.Writer) Tracer {

	return &tracer{out: w}
}
