package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {

	var buf bytes.Buffer

	tracer := New(&buf)

	if tracer == nil {
		t.Error("Return from New function should not be nil")

	} else {
		tracer.Trace("hello")
		if buf.String() != "hello\n" {
			t.Errorf("the strings doesn't match")
		}
	}

}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()

	silentTracer.Trace("something")

}
