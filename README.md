# uule-go

Manage uule parameter for Google Local Search.

References:
  - https://moz.com/ugc/geolocation-the-ultimate-tip-to-emulate-local-search
  - https://developers.google.com/adwords/api/docs/appendix/geotargeting

# Usage

```
package main

import (
	"github.com/takisawa/uule-go"
)

func main() {
	uule.Encode("Yokohama,Kanagawa Prefecture,Japan")
	uule.Decode("w+CAIQICIiWW9rb2hhbWEsS2FuYWdhd2EgUHJlZmVjdHVyZSxKYXBhbg==")
}

```
