package uule_go

import (
	"testing"
)

var (
	location_name string = "Yokohama,Kanagawa Prefecture,Japan"
	uule          string = "w+CAIQICIiWW9rb2hhbWEsS2FuYWdhd2EgUHJlZmVjdHVyZSxKYXBhbg=="
)

func TestEncode(t *testing.T) {
	if Encode(location_name) != uule {
		t.Errorf("Encode(%q) = %q", location_name, uule)
	}
}

func TestDecode(t *testing.T) {
	lname, _ := Decode(uule)
	if lname != location_name {
		t.Errorf("Dncode(%q) = %q", uule, location_name)
	}
}
