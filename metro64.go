package metro

import "encoding/binary"

func rotate_right(v uint64, k uint) uint64 {
	return (v >> k) | (v << (64 - k))
}

func Hash64_1(key []byte, seed uint32) uint64 {

	const k0 uint64 = 0xC83A91E1
	const k1 uint64 = 0x8648DBDB
	const k2 uint64 = 0x7BDEC03B
	const k3 uint64 = 0x2F5870A5

	ptr := key

	hash := (uint64(seed)+k2)*k0 + uint64(len(key))

	if len(ptr) >= 32 {
		v := [4]uint64{hash, hash, hash, hash}

		for len(ptr) >= 32 {
			v[0] += binary.LittleEndian.Uint64(ptr) * k0
			ptr = ptr[8:]
			v[0] = rotate_right(v[0], 29) + v[2]
			v[1] += binary.LittleEndian.Uint64(ptr) * k1
			ptr = ptr[8:]
			v[1] = rotate_right(v[1], 29) + v[3]
			v[2] += binary.LittleEndian.Uint64(ptr) * k2
			ptr = ptr[8:]
			v[2] = rotate_right(v[2], 29) + v[0]
			v[3] += binary.LittleEndian.Uint64(ptr) * k3
			ptr = ptr[8:]
			v[3] = rotate_right(v[3], 29) + v[1]
		}

		v[2] ^= rotate_right(((v[0]+v[3])*k0)+v[1], 33) * k1
		v[3] ^= rotate_right(((v[1]+v[2])*k1)+v[0], 33) * k0
		v[0] ^= rotate_right(((v[0]+v[2])*k0)+v[3], 33) * k1
		v[1] ^= rotate_right(((v[1]+v[3])*k1)+v[2], 33) * k0
		hash += v[0] ^ v[1]
	}

	if len(ptr) >= 16 {
		v0 := hash + (binary.LittleEndian.Uint64(ptr) * k0)
		ptr = ptr[8:]
		v0 = rotate_right(v0, 33) * k1
		v1 := hash + (binary.LittleEndian.Uint64(ptr) * k1)
		ptr = ptr[8:]
		v1 = rotate_right(v1, 33) * k2
		v0 ^= rotate_right(v0*k0, 35) + v1
		v1 ^= rotate_right(v1*k3, 35) + v0
		hash += v1
	}

	if len(ptr) >= 8 {
		hash += binary.LittleEndian.Uint64(ptr) * k3
		ptr = ptr[8:]
		hash ^= rotate_right(hash, 33) * k1
	}

	if len(ptr) > 4 {
		hash += uint64(binary.LittleEndian.Uint32(ptr)) * k3
		ptr = ptr[4:]
		hash ^= rotate_right(hash, 15) * k1
	}

	if len(ptr) >= 2 {
		hash += uint64(binary.LittleEndian.Uint16(ptr)) * k3
		ptr = ptr[2:]
		hash ^= rotate_right(hash, 13) * k1
	}

	if len(ptr) >= 1 {
		hash += uint64(ptr[0]) * k3
		hash ^= rotate_right(hash, 25) * k1
	}

	hash ^= rotate_right(hash, 33)
	hash *= k0
	hash ^= rotate_right(hash, 33)

	return hash
}

func Hash64_2(key []byte, seed uint32) uint64 {
	const k0 uint64 = 0xD6D018F5
	const k1 uint64 = 0xA2AA033B
	const k2 uint64 = 0x62992FC1
	const k3 uint64 = 0x30BC5B29

	ptr := key

	hash := (uint64(seed)+k2)*k0 + uint64(len(key))

	if len(ptr) >= 32 {
		v := [4]uint64{hash, hash, hash, hash}

		for len(ptr) >= 32 {
			v[0] += binary.LittleEndian.Uint64(ptr) * k0
			ptr = ptr[8:]
			v[0] = rotate_right(v[0], 29) + v[2]
			v[1] += binary.LittleEndian.Uint64(ptr) * k1
			ptr = ptr[8:]
			v[1] = rotate_right(v[1], 29) + v[3]
			v[2] += binary.LittleEndian.Uint64(ptr) * k2
			ptr = ptr[8:]
			v[2] = rotate_right(v[2], 29) + v[0]
			v[3] += binary.LittleEndian.Uint64(ptr) * k3
			ptr = ptr[8:]
			v[3] = rotate_right(v[3], 29) + v[1]
		}

		v[2] ^= rotate_right(((v[0]+v[3])*k0)+v[1], 30) * k1
		v[3] ^= rotate_right(((v[1]+v[2])*k1)+v[0], 30) * k0
		v[0] ^= rotate_right(((v[0]+v[2])*k0)+v[3], 30) * k1
		v[1] ^= rotate_right(((v[1]+v[3])*k1)+v[2], 30) * k0
		hash += v[0] ^ v[1]
	}

	if len(ptr) >= 16 {
		v0 := hash + (binary.LittleEndian.Uint64(ptr) * k2)
		ptr = ptr[8:]
		v0 = rotate_right(v0, 29) * k3
		v1 := hash + (binary.LittleEndian.Uint64(ptr) * k2)
		ptr = ptr[8:]
		v1 = rotate_right(v1, 29) * k3
		v0 ^= rotate_right(v0*k0, 34) + v1
		v1 ^= rotate_right(v1*k3, 34) + v0
		hash += v1
	}

	if len(ptr) >= 8 {
		hash += binary.LittleEndian.Uint64(ptr) * k3
		ptr = ptr[8:]
		hash ^= rotate_right(hash, 36) * k1
	}

	if len(ptr) >= 4 {
		hash += uint64(binary.LittleEndian.Uint32(ptr)) * k3
		ptr = ptr[4:]
		hash ^= rotate_right(hash, 15) * k1
	}

	if len(ptr) >= 2 {
		hash += uint64(binary.LittleEndian.Uint16(ptr)) * k3
		ptr = ptr[2:]
		hash ^= rotate_right(hash, 15) * k1
	}

	if len(ptr) >= 1 {
		hash += uint64(ptr[0]) * k3
		hash ^= rotate_right(hash, 23) * k1
	}

	hash ^= rotate_right(hash, 28)
	hash *= k0
	hash ^= rotate_right(hash, 29)

	return hash
}
