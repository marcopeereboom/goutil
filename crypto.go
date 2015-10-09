package goutil

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"hash"
	"io"
	"os"
)

func Zero(in []byte) {
	if in == nil {
		return
	}
	for i := 0; i < len(in); i++ {
		in[i] ^= in[i]
	}
}

func fileDigest(filename string, h hash.Hash) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer func() { _ = f.Close() }()

	_, err = io.Copy(h, f)
	return err
}

func FileSHA256(filename string) (*[sha256.Size]byte, error) {
	h := sha256.New()
	err := fileDigest(filename, h)
	if err != nil {
		return nil, err
	}

	var d [sha256.Size]byte
	copy(d[:], h.Sum(nil)[:])
	return &d, nil
}

func HMACSHA256(blob []byte, key []byte) (*[sha256.Size]byte, error) {
	h := hmac.New(sha256.New, key)

	r := bytes.NewReader(blob)
	_, err := io.Copy(h, r)
	if err != nil {
		return nil, err
	}

	var d [sha256.Size]byte
	copy(d[:], h.Sum(nil)[:])
	return &d, nil
}
func FileHMACSHA256(filename string, key []byte) (*[sha256.Size]byte, error) {
	h := hmac.New(sha256.New, key)
	err := fileDigest(filename, h)
	if err != nil {
		return nil, err
	}

	var d [sha256.Size]byte
	copy(d[:], h.Sum(nil)[:])
	return &d, nil
}
