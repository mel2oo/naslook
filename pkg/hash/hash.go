package hash

import (
	"crypto/md5"
	"encoding/hex"
	"os"
)

func GetFileMD5(file string) (string, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}

	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil)), nil
}
