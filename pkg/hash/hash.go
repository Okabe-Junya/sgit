package hash

import (
	"crypto/sha1"
	"encoding/hex"
)

func CalcSHA1(content string) string {
	h := sha1.New()
	h.Write([]byte(content))
	hs := h.Sum(nil)
	sha1_content := hex.EncodeToString(hs)
	return sha1_content
}
