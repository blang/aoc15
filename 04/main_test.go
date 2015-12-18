package main

import (
	"crypto/md5"
	"encoding/hex"
	"testing"
)

func TestMain(t *testing.T) {
	h := md5.Sum([]byte("abcdef609043"))
	t.Logf("%s", hex.EncodeToString(h[:]))
}
