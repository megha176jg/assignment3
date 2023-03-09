package web

import "io"

type Downloader interface {
	Download(s string) (r io.Reader, err error)
}
