package metro

import "encoding/binary"

func Hash128_1(key []byte, seed uint32) (uint64, uint64) {

	const k0 uint64 = 0xC83A91E1
	const k1 uint64 = 0x8648DBDB
	const k2 uint64 = 0x7BDEC03B
	const k3 uint64 = 0x2F5870A5

	ptr := key

	var v [4]uint64

	v[0] = (uint64(seed)-k0)*k3 + uint64(len(ptr))
	v[1] = (uint64(seed)+k1)*k2 + uint64(len(ptr))

	if len(ptr) >= 32 {
		v[2] = (uint64(seed)+k0)*k2 + uint64(len(ptr))
		v[3] = (uint64(seed)-k1)*k3 + uint64(len(ptr))

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

		v[2] ^= rotate_right(((v[0]+v[3])*k0)+v[1], 26) * k1
		v[3] ^= rotate_right(((v[1]+v[2])*k1)+v[0], 26) * k0
		v[0] ^= rotate_right(((v[0]+v[2])*k0)+v[3], 26) * k1
		v[1] ^= rotate_right(((v[1]+v[3])*k1)+v[2], 30) * k0
	}

	if len(ptr) >= 16 {
		v[0] += binary.LittleEndian.Uint64(ptr) * k2
		ptr = ptr[8:]
		v[0] = rotate_right(v[0], 33) * k3
		v[1] += binary.LittleEndian.Uint64(ptr) * k2
		ptr = ptr[8:]
		v[1] = rotate_right(v[1], 33) * k3
		v[0] ^= rotate_right((v[0]*k2)+v[1], 17) * k1
		v[1] ^= rotate_right((v[1]*k3)+v[0], 17) * k0
	}

	if len(ptr) >= 8 {
		v[0] += binary.LittleEndian.Uint64(ptr) * k2
		ptr = ptr[8:]
		v[0] = rotate_right(v[0], 33) * k3
		v[0] ^= rotate_right((v[0]*k2)+v[1], 20) * k1
	}

	if len(ptr) >= 4 {
		v[1] += uint64(binary.LittleEndian.Uint32(ptr)) * k2
		ptr = ptr[4:]
		v[1] = rotate_right(v[1], 33) * k3
		v[1] ^= rotate_right((v[1]*k3)+v[0], 18) * k0
	}

	if len(ptr) >= 2 {
		v[0] += uint64(binary.LittleEndian.Uint16(ptr)) * k2
		ptr = ptr[2:]
		v[0] = rotate_right(v[0], 33) * k3
		v[0] ^= rotate_right((v[0]*k2)+v[1], 24) * k1
	}

	if len(ptr) >= 1 {
		v[1] += uint64(ptr[0]) * k2
		v[1] = rotate_right(v[1], 33) * k3
		v[1] ^= rotate_right((v[1]*k3)+v[0], 24) * k0
	}

	v[0] += rotate_right((v[0]*k0)+v[1], 13)
	v[1] += rotate_right((v[1]*k1)+v[0], 37)
	v[0] += rotate_right((v[0]*k2)+v[1], 13)
	v[1] += rotate_right((v[1]*k3)+v[0], 37)

	return v[0], v[1]
}

func Hash128_2(key []byte, seed uint32) (uint64, uint64) {
	const k0 uint64 = 0xD6D018F5
	const k1 uint64 = 0xA2AA033B
	const k2 uint64 = 0x62992FC1
	const k3 uint64 = 0x30BC5B29

	ptr := key

	var v [4]uint64

	v[0] = (uint64(seed)-k0)*k3 + uint64(len(ptr))
	v[1] = (uint64(seed)+k1)*k2 + uint64(len(ptr))

	if len(ptr) >= 32 {
		v[2] = (uint64(seed)+k0)*k2 + uint64(len(ptr))
		v[3] = (uint64(seed)-k1)*k3 + uint64(len(ptr))

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
	}

	if len(ptr) >= 16 {
		v[0] += binary.LittleEndian.Uint64(ptr) * k2
		ptr = ptr[8:]
		v[0] = rotate_right(v[0], 29) * k3
		v[1] += binary.LittleEndian.Uint64(ptr) * k2
		ptr = ptr[8:]
		v[1] = rotate_right(v[1], 29) * k3
		v[0] ^= rotate_right((v[0]*k2)+v[1], 29) * k1
		v[1] ^= rotate_right((v[1]*k3)+v[0], 29) * k0
	}

	if len(ptr) >= 8 {
		v[0] += binary.LittleEndian.Uint64(ptr) * k2
		ptr = ptr[8:]
		v[0] = rotate_right(v[0], 29) * k3
		v[0] ^= rotate_right((v[0]*k2)+v[1], 29) * k1
	}

	if len(ptr) >= 4 {
		v[1] += uint64(binary.LittleEndian.Uint32(ptr)) * k2
		ptr = ptr[4:]
		v[1] = rotate_right(v[1], 29) * k3
		v[1] ^= rotate_right((v[1]*k3)+v[0], 25) * k0
	}

	if len(ptr) >= 2 {
		v[0] += uint64(binary.LittleEndian.Uint16(ptr)) * k2
		ptr = ptr[2:]
		v[0] = rotate_right(v[0], 29) * k3
		v[0] ^= rotate_right((v[0]*k2)+v[1], 30) * k1
	}

	if len(ptr) >= 1 {
		v[1] += uint64(ptr[0]) * k2
		v[1] = rotate_right(v[1], 29) * k3
		v[1] ^= rotate_right((v[1]*k3)+v[0], 18) * k0
	}

	v[0] += rotate_right((v[0]*k0)+v[1], 33)
	v[1] += rotate_right((v[1]*k1)+v[0], 33)
	v[0] += rotate_right((v[0]*k2)+v[1], 33)
	v[1] += rotate_right((v[1]*k3)+v[0], 33)

	return v[0], v[1]
}
