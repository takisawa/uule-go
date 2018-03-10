package uule_go

import (
	"encoding/base64"
	"fmt"
	"unicode/utf8"
)

const SPECIAL_KEY_TABLE = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func Encode(location_canonical_name string) string {
	key := []rune(SPECIAL_KEY_TABLE)[utf8.RuneCountInString(location_canonical_name)]
	return fmt.Sprintf("w+CAIQICI%s%s", string(key), base64.StdEncoding.EncodeToString([]byte(location_canonical_name)))
}

func Decode(uule string) string {
	return "Decode"
}
