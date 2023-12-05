package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"
)

var validHashes = []string{
	"sha1",
	"sha256",
	"md5",
}

func filehash(filePath string, alg string) string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var h = sha256.New()
	if alg == "sha256" {
		h = sha256.New()
	} else if alg == "sha1" {
		h = sha1.New()
	} else if alg == "md5" {
		h = md5.New()
	}
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}

func stringhash(input string, alg string) string {
	bv := []byte(input)

	var hasher = sha256.New()
	if alg == "sha256" {
		hasher = sha256.New()
	} else if alg == "sha1" {
		hasher = sha1.New()
	} else if alg == "md5" {
		hasher = md5.New()
	}
	hasher.Write(bv)
	sha := hex.EncodeToString(hasher.Sum(nil))
	return sha
}
