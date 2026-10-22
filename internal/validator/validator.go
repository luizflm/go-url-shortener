package validator

import (
	"net/url"
	"regexp"
	"strings"
)

var Base62RX = regexp.MustCompile(`^[1-9A-Za-z][0-9A-Za-z]{0,10}$`)

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func ValidURL(longUrl string) bool {
	u, err := url.Parse(longUrl)
	return err == nil &&
		(u.Scheme == "http" || u.Scheme == "https") &&
		u.Hostname() != ""
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}
