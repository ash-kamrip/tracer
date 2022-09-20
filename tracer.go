package trace

import (
	"fmt"
	"io"
)

// a struct type : private
type tracer struct {
	out io.Writer
}
type niltracer struct {
}

// an interface type : Exported
type Tracer interface {
	Trace(...interface{})
}

func (t *tracer) Trace(a ...interface{}) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)

}
func (t *niltracer) Trace(a ...interface{}) {

}
func New(w io.Writer) Tracer {

	return &tracer{out: w}
}

func Off() Tracer {
	return &niltracer{}
}
