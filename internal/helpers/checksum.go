package helpers

import "hash/adler32"

func Adler32(s string) uint32 {
	return adler32.Checksum([]byte(s))
}

func cadler32(s string) uint32 {
	const mod = 65521
	var a, b uint32 = 1, 0
	for i := 0; i < len(s); i++ {
		a = (a + uint32(s[i])) % mod
		b = (b + a) % mod
	}
	return (b << 16) | a
}
