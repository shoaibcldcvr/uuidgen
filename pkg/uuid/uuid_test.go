package uuid

import (
	"testing"
)

func TestNewV1(t *testing.T) {
	u, err := NewV1()
	if err != nil {
		t.Fatalf("NewV1() failed: %v", err)
	}

	// Verify version bits (should be 0b00010000)
	if u[6]&0xf0 != 0x10 {
		t.Errorf("UUID version bits incorrect: got %02x, want 10", u[6]&0xf0)
	}

	// Verify variant bits (should be 0b10000000)
	if u[8]&0xc0 != 0x80 {
		t.Errorf("UUID variant bits incorrect: got %02x, want 80", u[8]&0xc0)
	}
}
