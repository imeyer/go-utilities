package hostname

import (
	"testing"
)

func TestReverse(t *testing.T) {
	const hostname_in, hostname_out = "test.a.is.this", "this.is.a.test"

	if x := Reverse(hostname_in); x != hostname_out {
		t.Errorf("Reverse(%v) = %v, want %v", hostname_in, x, hostname_out)
	}
}

func TestReverseOffset(t *testing.T) {
	const hostname_in, hostname_out = "test.a.is.this", "a.test"

	if x := ReverseOffset(hostname_in, 2); x != hostname_out {
		t.Errorf("Reverse(%v) = %v, want %v", hostname_in, x, hostname_out)
	}
}
