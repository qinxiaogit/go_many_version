package parse

import "net/url"

type Result struct {
	URL  *url.URL
	M3u8 *M3u8
	Keys map[int]string
}

