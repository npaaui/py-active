package wxhelper

import (
	"sort"
	"crypto/sha1"
	"io"
	"strings"
	"fmt"
)

/////////////////////////Signature////////////////////////////////

func MakeSignature(timestamp, nonce string) string {
	sl := []string{Token, timestamp, nonce}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}

// Signature 对加密的报文计算签名
func MakeMsgSignature(timestamp, nonce, msgEncrypt string) string {
	sl := []string{Token, timestamp, nonce, msgEncrypt}
	sort.Strings(sl)
	s := sha1.New()
	io.WriteString(s, strings.Join(sl, ""))
	return fmt.Sprintf("%x", s.Sum(nil))
}
