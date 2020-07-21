package srand

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"os"
	"strconv"
	"time"
)

// Fast returns value that can be used to initialize random seed.
// That value is not cryptographically secure.
func Fast() int64 {
	return time.Now().UTC().UnixNano()
}

// Overridable will check if the "key" var is configured, else
// it will return a Fast random seed.
//
// If $SRAND is not parseable, panic is raised.
func Overridable(key string) int64 {
	if env := os.Getenv(key); env != "" {
		n, err := strconv.ParseInt(env, 10, 64)
		if err != nil {
			panic(err)
		}
		return n
	}
	return Fast()
}

// Secure returns a cryptographically secure random seed.
//
// Based on https://stackoverflow.com/a/54491783
func Secure() int64 {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("cannot seed math/rand package with cryptographically secure random number generator")
	}
	return int64(binary.LittleEndian.Uint64(b[:]))
}
