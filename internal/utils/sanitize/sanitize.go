package sanitize

import (
	"strconv"

	"github.com/microcosm-cc/bluemonday"
)

// sanitizes a string and safely parses it into int64
func SanitizeAndParseInt(input string) (int64, error) {
	p := bluemonday.UGCPolicy()
	safe := p.Sanitize(input)
	return strconv.ParseInt(safe, 10, 64)
}

// Sanatize the json fields
func SanitizeJsonItems(name, email string) (string, string) {
	p := bluemonday.UGCPolicy()
	name = p.Sanitize(name)
	email = p.Sanitize(email)

	return name, email
}
