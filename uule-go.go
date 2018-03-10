package uule_go

import (
	"encoding/base64"
	"fmt"
)

const SPECIAL_KEY_TABLE = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func Encode(location_canonical_name string) string {
	len_name := len(location_canonical_name)
	key := SPECIAL_KEY_TABLE[len_name : len_name+1]
	return fmt.Sprintf("w+CAIQICI%s%s", key, base64.StdEncoding.EncodeToString([]byte(location_canonical_name)))
}

func Decode(uule string) (string, error) {
	bytes, err := base64.StdEncoding.DecodeString(uule[10:])

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
