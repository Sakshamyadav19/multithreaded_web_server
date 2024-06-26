package utils

import (
	"strings"
)

func ParseUrl(req string) string {
	parts := strings.Split(req, " ")

	if len(parts) < 2 {
		return ""
	}

	path := parts[1]

	url := strings.TrimPrefix(path, "/")

	url = strings.TrimSuffix(url, "/")

	if url == "" {
		return ""
	}

	return url

}