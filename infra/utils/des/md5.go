package des

import (
	"crypto/md5"
	"encoding/hex"
)

// string to md5
func ToMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
