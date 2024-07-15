/*
Adler32, fnv32, and crc32 are checksum algorithms, not cryptographic hash functions.
They are not meant to be secure, only for us to have a reasonable guarantee that
the id always will be unique and the same for the same input.

Which means we can quickly query the database for the id, and we can be sure that
the correct record is returned.
-------------------------------------------------------------------------------------
This benchmark shows that the adler32 is consisently about 4 times faster than crc32.
*/
package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"hash/adler32"
	"hash/crc32"
	"hash/fnv"

	"time"
)

// A million loops
const loops = 1_000_000

// const str = "The quick brown fox jumps over the lazy dog."

// I made this up - it's a pangram in Norwegian. It contains 47 characters.
const _ = "Svær hårete bauxitt koder webfjell og syr emnequizcup på wc"

// actual real life example that I used in a project
const str = "some-sample-slug-1234"

func Bench() {
	fmt.Println("Benchmarking...")

	// crc32
	tss := time.Now()
	var x uint32
	for i := 0; i < loops; i++ {
		x = crc32.ChecksumIEEE([]byte(str))
	}
	tse := time.Now()
	fmt.Printf("crc32:\t\t%.5f\t%d\n", tse.Sub(tss).Seconds(), x)

	// md5
	tss = time.Now()
	var md5Sum [16]byte
	for i := 0; i < loops; i++ {
		md5Sum = md5.Sum([]byte(str))
	}
	tse = time.Now()
	fmt.Printf("md5:\t\t%.5f\t%x\n", tse.Sub(tss).Seconds(), md5Sum)

	// sha1
	tss = time.Now()
	var sha1Sum [20]byte
	for i := 0; i < loops; i++ {
		sha1Sum = sha1.Sum([]byte(str))
	}
	tse = time.Now()
	fmt.Printf("sha1:\t\t%.5f\t%x\n", tse.Sub(tss).Seconds(), sha1Sum)

	// xor
	tss = time.Now()
	var xorRes byte
	for i := 0; i < loops; i++ {
		xorRes = 0x42
		for j := 0; j < len(str); j++ {
			xorRes ^= str[j]
		}
	}
	tse = time.Now()
	fmt.Printf("xor:\t\t%.5f\t%d\n", tse.Sub(tss).Seconds(), xorRes)

	// add
	tss = time.Now()
	var addRes int
	for i := 0; i < loops; i++ {
		addRes = 0
		for j := 0; j < len(str); j++ {
			addRes += int(str[j])
		}
	}
	tse = time.Now()
	fmt.Printf("add:\t\t%.5f\t%d\n", tse.Sub(tss).Seconds(), addRes)

	// fnv32
	tss = time.Now()
	var fnv32 = fnv.New32()
	var fnv32Res int
	for i := 0; i < loops; i++ {
		fnv32.Reset()
		fnv32.Write([]byte(str))
		fnv32Res = int(fnv32.Sum32())
	}
	tse = time.Now()
	fmt.Printf("fnv32:\t%.5f\t%d\n", tse.Sub(tss).Seconds(), fnv32Res)

	// adler32
	tss = time.Now()
	var adler32GoRes uint32
	for i := 0; i < loops; i++ {
		adler32GoRes = adler32.Checksum([]byte(str))
	}
	tse = time.Now()
	fmt.Printf("adler32:\t%.5f\t%d\n", tse.Sub(tss).Seconds(), adler32GoRes)
}

/* Output:

❯ go run hash_bench.go | sort -k 2 | column -t
Benchmarking...
add:             0.00488  1858
adler32:         0.00643  1474889539
xor:             0.00650  116
fnv32:           0.01096  2944183525
crc32:           0.02697  2117753032
md5:             0.08152  675d052e8275dce6efd9895872612370
sha1:            0.09141  42a6b210ec75ccd25d31075bce2365360d418c58

*/
