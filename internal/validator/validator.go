package validator

import (
	"net/url"
	"strings"
)

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func ValidURL(longUrl string) bool {
	u, err := url.Parse(longUrl)
	return err == nil &&
		(u.Scheme == "http" || u.Scheme == "https") &&
		u.Hostname() != ""
}
