package metro

import (
	"encoding/binary"
	"math/bits"
	"math/rand/v2"
	"strconv"
	"testing"
)

func refHash64(buffer []byte, seed uint64) uint64 {

	const (
		k0 = 0xD6D018F5
		k1 = 0xA2AA033B
		k2 = 0x62992FC1
		k3 = 0x30BC5B29
	)

	ptr := buffer

	hash := (seed + k2) * k0

	if len(ptr) >= 32 {
		v0, v1, v2, v3 := hash, hash, hash, hash

		for len(ptr) >= 32 {
			v0 += binary.LittleEndian.Uint64(ptr[:8]) * k0
			v0 = bits.RotateLeft64(v0, -29) + v2
			v1 += binary.LittleEndian.Uint64(ptr[8:16]) * k1
			v1 = bits.RotateLeft64(v1, -29) + v3
			v2 += binary.LittleEndian.Uint64(ptr[16:24]) * k2
			v2 = bits.RotateLeft64(v2, -29) + v0
			v3 += binary.LittleEndian.Uint64(ptr[24:32]) * k3
			v3 = bits.RotateLeft64(v3, -29) + v1
			ptr = ptr[32:]
		}

		v2 ^= bits.RotateLeft64(((v0+v3)*k0)+v1, -37) * k1
		v3 ^= bits.RotateLeft64(((v1+v2)*k1)+v0, -37) * k0
		v0 ^= bits.RotateLeft64(((v0+v2)*k0)+v3, -37) * k1
		v1 ^= bits.RotateLeft64(((v1+v3)*k1)+v2, -37) * k0
		hash += v0 ^ v1
	}

	if len(ptr) >= 16 {
		v0 := hash + (binary.LittleEndian.Uint64(ptr[:8]) * k2)
		v0 = bits.RotateLeft64(v0, -29) * k3
		v1 := hash + (binary.LittleEndian.Uint64(ptr[8:16]) * k2)
		v1 = bits.RotateLeft64(v1, -29) * k3
		v0 ^= bits.RotateLeft64(v0*k0, -21) + v1
		v1 ^= bits.RotateLeft64(v1*k3, -21) + v0
		hash += v1
		ptr = ptr[16:]
	}

	if len(ptr) >= 8 {
		hash += binary.LittleEndian.Uint64(ptr[:8]) * k3
		ptr = ptr[8:]
		hash ^= bits.RotateLeft64(hash, -55) * k1
	}

	if len(ptr) >= 4 {
		hash += uint64(binary.LittleEndian.Uint32(ptr[:4])) * k3
		hash ^= bits.RotateLeft64(hash, -26) * k1
		ptr = ptr[4:]
	}

	if len(ptr) >= 2 {
		hash += uint64(binary.LittleEndian.Uint16(ptr[:2])) * k3
		ptr = ptr[2:]
		hash ^= bits.RotateLeft64(hash, -48) * k1
	}

	if len(ptr) >= 1 {
		hash += uint64(ptr[0]) * k3
		hash ^= bits.RotateLeft64(hash, -37) * k1
	}

	hash ^= bits.RotateLeft64(hash, -28)
	hash *= k0
	hash ^= bits.RotateLeft64(hash, -29)

	return hash
}

func TestHash64MatchesReference(t *testing.T) {
	rng := rand.New(rand.NewPCG(1, 2))
	buf := make([]byte, 4096+8)
	for i := range buf {
		buf[i] = byte(rng.Uint64())
	}
	seeds := []uint64{0, 1, 0xffffffffffffffff, 0x0123456789abcdef}
	for n := 0; n <= 300; n++ {
		for off := 0; off < 8; off++ { // unaligned starts
			b := buf[off : off+n]
			for _, seed := range seeds {
				want := refHash64(b, seed)
				if got := Hash64(b, seed); got != want {
					t.Fatalf("Hash64(len=%d off=%d seed=%#x) = %#x, want %#x", n, off, seed, got, want)
				}
				if got := Hash64Str(string(b), seed); got != want {
					t.Fatalf("Hash64Str(len=%d off=%d seed=%#x) = %#x, want %#x", n, off, seed, got, want)
				}
			}
		}
	}
	for range 20000 {
		n := rng.IntN(4096)
		b := buf[:n]
		seed := rng.Uint64()
		if got, want := Hash64(b, seed), refHash64(b, seed); got != want {
			t.Fatalf("Hash64(len=%d seed=%#x) = %#x, want %#x", n, seed, got, want)
		}
	}
}

func BenchmarkHash64Ref(b *testing.B) {
	for _, n := range []int{8, 31, 64, 1024} {
		buf := make([]byte, n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			b.SetBytes(int64(n))
			for b.Loop() {
				refHash64(buf, 0)
			}
		})
	}
}

func BenchmarkHash64Asm(b *testing.B) {
	for _, n := range []int{8, 31, 64, 1024} {
		buf := make([]byte, n)
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			b.SetBytes(int64(n))
			for b.Loop() {
				Hash64(buf, 0)
			}
		})
	}
}
