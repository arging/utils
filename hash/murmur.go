// Copyright 2014 li. All rights reserved.
// Use of this source code is governed by a MIT/X11
// license that can be found in the LICENSE file.

package hash

// See http://smhasher.googlecode.com/svn/trunk/MurmurHash3.cpp
// Little Endian
func get4ByteChunk(key []byte, start int) uint32 {
	s0 := 4 * start
	return uint32(key[s0]) | uint32(key[s0+1])<<8 |
		uint32(key[s0+2])<<16 | uint32(key[s0+3])<<24
}

// Little Endian
func Murmur3_32(key []byte, seed uint32) uint32 {
	const (
		c1 uint32 = 0xcc9e2d51
		c2 uint32 = 0x1b873593
		r1        = 15
		r2        = 13
		m         = 5
		n         = 0xe6546b64
	)

	var (
		keyLen  = len(key)
		nblocks = keyLen / 4
		hash    = seed
	)

	for i := 0; i < nblocks; i++ {
		k := get4ByteChunk(key, i)

		k *= c1
		k = (k << r1) | (k >> (32 - r1))
		k *= c2

		hash ^= k
		hash = (hash << r2) | (hash >> (32 - r2))
		hash = hash*m + n
	}

	k := uint32(0)
	index := nblocks * 4

	switch keyLen & 3 {
	case 3:
		k ^= uint32(key[index+2]) << 16
		fallthrough
	case 2:
		k ^= uint32(key[index+1]) << 8
		fallthrough
	case 1:
		k ^= uint32(key[index])

		k *= c1
		k = (k << r1) | (k >> (32 - r1))
		k *= c2
		hash ^= k
	}

	hash ^= uint32(keyLen)

	hash ^= hash >> 16
	hash *= 0x85ebca6b
	hash ^= hash >> 13
	hash *= 0xc2b2ae35
	hash ^= hash >> 16

	return hash
}

// Little Endian
func get8ByteChunk(key []byte, start int) uint64 {
	s0 := 8 * start
	return uint64(key[s0]) | uint64(key[s0+1])<<8 |
		uint64(key[s0+2])<<16 | uint64(key[s0+3])<<24 |
		uint64(key[s0+4])<<32 | uint64(key[s0+5])<<40 |
		uint64(key[s0+6])<<48 | uint64(key[s0+7])<<56
}

// Little Endian
func Murmur3_128(key []byte, seed uint32) (uint64, uint64) {
	const (
		c1 uint64 = 0x87c37b91114253d5
		c2 uint64 = 0x4cf5ad432745937f
	)

	var (
		h1      = uint64(seed)
		h2      = uint64(seed)
		keyLen  = len(key)
		nblocks = keyLen / 16
	)

	for i := 0; i < nblocks; i++ {
		k1 := get8ByteChunk(key, i*2+0)
		k2 := get8ByteChunk(key, i*2+1)

		k1 *= c1
		k1 = (k1 << 31) | (k1 >> (64 - 31))
		k1 *= c2
		h1 ^= k1

		h1 = (h1 << 27) | (h1 >> (64 - 27))
		h1 += h2
		h1 = h1*5 + 0x52dce729

		k2 *= c2
		k2 = (k2 << 33) | (k2 >> (64 - 33))
		k2 *= c1
		h2 ^= k2

		h2 = (h2 << 31) | (h2 >> (64 - 31))
		h2 += h1
		h2 = h2*5 + 0x38495ab5
	}

	var (
		k1    = uint64(0)
		k2    = uint64(0)
		index = nblocks * 16
	)

	switch keyLen & 15 {
	case 15:
		k2 ^= uint64(key[index+14]) << 48
		fallthrough
	case 14:
		k2 ^= uint64(key[index+13]) << 40
		fallthrough
	case 13:
		k2 ^= uint64(key[index+12]) << 32
		fallthrough
	case 12:
		k2 ^= uint64(key[index+11]) << 24
		fallthrough
	case 11:
		k2 ^= uint64(key[index+10]) << 16
		fallthrough
	case 10:
		k2 ^= uint64(key[index+9]) << 8
		fallthrough
	case 9:
		k2 ^= uint64(key[index+8])

		k2 *= c2
		k2 = (k2 << 33) | (k2 >> (64 - 33))
		k2 *= c1
		h2 ^= k2
		fallthrough

	case 8:
		k1 ^= uint64(key[index+7]) << 56
		fallthrough
	case 7:
		k1 ^= uint64(key[index+6]) << 48
		fallthrough
	case 6:
		k1 ^= uint64(key[index+5]) << 40
		fallthrough
	case 5:
		k1 ^= uint64(key[index+4]) << 32
		fallthrough
	case 4:
		k1 ^= uint64(key[index+3]) << 24
		fallthrough
	case 3:
		k1 ^= uint64(key[index+2]) << 16
		fallthrough
	case 2:
		k1 ^= uint64(key[index+1]) << 8
		fallthrough
	case 1:
		k1 ^= uint64(key[index+0])

		k1 *= c1
		k1 = (k1 << 31) | (k1 >> (64 - 31))
		k1 *= c2
		h1 ^= k1
	}

	h1 ^= uint64(keyLen)
	h2 ^= uint64(keyLen)

	h1 += h2
	h2 += h1

	h1 ^= h1 >> 33
	h1 *= uint64(0xff51afd7ed558ccd)
	h1 ^= h1 >> 33
	h1 *= uint64(0xc4ceb9fe1a85ec53)
	h1 ^= h1 >> 33

	h2 ^= h2 >> 33
	h2 *= uint64(0xff51afd7ed558ccd)
	h2 ^= h2 >> 33
	h2 *= uint64(0xc4ceb9fe1a85ec53)
	h2 ^= h2 >> 33

	h1 += h2
	h2 += h1

	return h1, h2
}
