package internal

import (
	"encoding/base64"
)

// ByteSlice is a wrapper of []byte to implement StringCodec.
// NOTE: we're using base64.StdEncoding here, not base64.URLEncoding.
type ByteSlice []byte

func (bs ByteSlice) ToString() (string, error) {
	return base64.StdEncoding.EncodeToString(bs), nil
}

func (bs *ByteSlice) FromString(s string) error {
	v, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	*bs = ByteSlice(v)
	return nil
}
