/*
	Custom implementation of a UUID

	Just for the sake of learning a little bit more about UUIDs and why they are the way they are
	And also to have a little bit of fun with it
*/

package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/pynezz/pynezzentials/fsutil"
)

const (
	version = 0   // 0 is the first version
	variant = "s" // variant s is the standard variant
)

// This is my homemade UUID variant
type UUID struct {
	// Time is the time the UUID was created
	Time int64
	// UUID is the unique identifier
	Identifier string

	// The UUID has a version
	Version int

	// The UUID has a variant
	Variant string

	// Signature is the HMAC signature of the UUID
	Signature string
}

func getEnv(key string) (value string) {
	c, err := fsutil.GetFileContent(".env")
	if err != nil {
		return ""
	}

	sStr := strings.Split(c, "\n")
	for _, line := range sStr {
		if strings.Contains(line, key) {
			return strings.Split(line, "=")[1]
		}
	}

	return ""
}

func (u UUID) String() string {
	return u.Identifier
}

// Hmmmm - I'm just winging it for now
func (u UUID) AsUint() uint {
	n := new(big.Int)
	res, ok := n.SetString(fmt.Sprintf("%x", u.String()), 16)
	if !ok {
		fmt.Printf("Error converting UUID to uint: %s\n", u.String())
		return 0
	}

	return uint(res.Uint64())
}

// Uuid generates a UUID for a given username
func Uuid(username string) UUID {
	key := getEnv("UUID_KEY")
	t := time.Now().Unix()
	tBytes := []byte(fmt.Sprintf("%d", t))

	part1 := fmt.Sprintf("%x%x", version, variant)                        // Version and variant
	part2 := hmac.New(sha256.New, []byte(key)).Sum([]byte(tBytes))[0:8]   // Time signature
	part3 := hmac.New(sha256.New, []byte(key)).Sum([]byte(username))[0:8] // username signature

	uuid := fmt.Sprintf("%s-%x-%x", part1, part2, part3) // UUID

	uuidSigned := hmac.New(sha256.New, []byte(key)).Sum([]byte(uuid))[0:8] // UUID signature

	return UUID{
		Time:       t,
		Identifier: uuid,
		Version:    version,
		Variant:    variant,
		Signature:  fmt.Sprintf("%x", uuidSigned),
	}
}

func ParseUUID(uuid string) (UUID, error) {
	// verify the integrity of the UUID
	key := getEnv("UUID_KEY")

	signature := uuid[0:8]
	part1 := hmac.New(sha256.New, []byte(key)).Sum(nil)[0:8]

	// TODO: Verify if this is correct or not later
	if signature != fmt.Sprintf("%x", part1) {
		return UUID{}, fmt.Errorf("UUID signature is invalid")
	}

	return UUID{}, nil
}
