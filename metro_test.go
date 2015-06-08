package metro

import (
	"encoding/binary"
	"encoding/hex"
	"testing"
)

// From src/testvector.h

var key63 = []byte("012345678901234567890123456789012345678901234567890123456789012")

func Test64(t *testing.T) {

	tests := []struct {
		f     func([]byte, uint32) uint64
		which int
		seed  uint32
		want  string
	}{
		{Hash64_1, 1, 0, "658F044F5C730E40"},
		{Hash64_2, 2, 0, "073CAAB960623211"},
		{Hash64_1, 1, 1, "AE49EBB0A856537B"},
		{Hash64_2, 2, 1, "CF518E9CF58402C0"},
	}

	for _, tt := range tests {
		want, _ := hex.DecodeString(tt.want)
		want64 := binary.LittleEndian.Uint64(want)
		if got := tt.f(key63, tt.seed); got != want64 {
			t.Errorf("Hash64_%d(%q, %d)=%x, want %x\n", tt.which, key63, tt.seed, got, want64)
		}
	}
}

func Test128(t *testing.T) {

	tests := []struct {
		f     func([]byte, uint32) (uint64, uint64)
		which int
		seed  uint32
		want  string
	}{
		{Hash128_1, 1, 0, "ED9997ED9D0A8B0FF3F266399477788F"},
		{Hash128_2, 2, 0, "7BBA6FE119CF35D45507EDF3505359AB"},
		{Hash128_1, 1, 1, "DDA6BA67F7DE755EFDF6BEABECCFD1F4"},
		{Hash128_2, 2, 1, "2DA6AF149A5CDBC12B09DB0846D69EF0"},
	}

	for _, tt := range tests {
		want, _ := hex.DecodeString(tt.want)
		want64a := binary.LittleEndian.Uint64(want)
		want64b := binary.LittleEndian.Uint64(want[8:])
		if gota, gotb := tt.f(key63, tt.seed); gota != want64a || gotb != want64b {
			t.Errorf("Hash128_%d(%q, %d)=(%x, %x), want (%x, %x)\n", tt.which, key63, tt.seed, gota, gotb, want64a, want64b)
		}
	}
}
