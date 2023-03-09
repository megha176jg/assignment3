package zip

import "io"

type Archiver interface {
	Archive(names []string, readers ...io.Reader) (outR io.Reader, err error)
}
