package options

import (
	"crypto/md5"
	"encoding/hex"
)

type Option struct {
	id    string
	value string
}

func hashOfText(val string) string {
	hash := md5.Sum([]byte(val))
	return hex.EncodeToString(hash[:])
}

func newOptionFromValue(val string) Option {
	return Option{
		id:    hashOfText(val),
		value: val,
	}
}
