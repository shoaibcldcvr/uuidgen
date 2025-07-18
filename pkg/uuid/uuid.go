package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"time"
)

// UUID formats (RFC 4122)
type UUID [16]byte

// Generate a random UUID (v4)
func NewV4() (UUID, error) {
	var uuid UUID
	_, err := rand.Read(uuid[:])
	if err != nil {
		return UUID{}, err
	}
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant RFC 4122
	return uuid, nil
}

func getTimeBytes(t time.Time) []byte {
	nanos := t.UnixNano()
	// UUID v1 uses 100-nanosecond intervals since 00:00:00.00, 15 October 1582
	uuidEpoch := time.Date(1582, 10, 15, 0, 0, 0, 0, time.UTC)
	intervals := uint64(nanos/100) - uint64(uuidEpoch.UnixNano()/100)

	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, intervals)
	return b
}

func NewV1() (UUID, error) {
	var uuid UUID
	timeBytes := getTimeBytes(time.Now())

	// Copy time (first 8 bytes)
	copy(uuid[:], timeBytes[2:]) // Skip the first 2 bytes of the 8-byte time

	// Set version (bits 4-7 of byte 6)
	uuid[6] = (uuid[6] & 0x0f) | 0x10 // Version 1

	// Set variant (bits 6-7 of byte 8)
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant RFC 4122

	return uuid, nil
}

// String representation (8-4-4-4-12 format)
func (u UUID) String() string {
	return fmt.Sprintf(
		"%08x-%04x-%04x-%04x-%012x",
		u[0:4], u[4:6], u[6:8], u[8:10], u[10:],
	)
}

func uint64ToBytes(n uint64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, n)
	return b
}
