package uule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	location_name string = "Yokohama,Kanagawa Prefecture,Japan"
	uule          string = "w+CAIQICIiWW9rb2hhbWEsS2FuYWdhd2EgUHJlZmVjdHVyZSxKYXBhbg=="
)

func TestEncode(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(uule, Encode(location_name), "Invalid encoded string.")
}

func TestDecode(t *testing.T) {
	assert := assert.New(t)

	lname, err := Decode(uule)

	assert.Nil(err)
	assert.Equal(location_name, lname, "Invalid decoded string.")
}
