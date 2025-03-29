package internal

import (
	"regexp"
	"strings"
)

func ToString(arg []regexp.Regexp, separator ...string) string {
	sep := ", "

	if len(separator) > 0 {
		sep = separator[0]
	}

	var out string
	for _, el := range arg {
		out += el.String() + sep
	}

	out, _ = strings.CutSuffix(out, ", ")

	return out
}
