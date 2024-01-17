package url

import (
	"errors"
	"fmt"
	"strings"
)

type URL struct {
	Host   string
	Scheme string
	Path   string
}

// Parse parses raw url into a URL structure.
func Parse(url string) (*URL, error) {
	if !strings.Contains(url, "https://") {
		return nil, errors.New("invalid url format")
	}
	i := strings.Index(url, "://")
	fmt.Println(i)
	if i < 0 {
		return nil, errors.New("missing scheme")
	}
	scheme, hostPath := url[:i], url[i+3:]
	host, path := hostPath, ""
	if i := strings.Index(hostPath, "/"); i >= 0 {
		host, path = hostPath[:i], hostPath[i:]
	}
	fmt.Println(scheme, host, path)
	return &URL{Host: host, Scheme: scheme, Path: path}, nil
}
