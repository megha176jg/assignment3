package zip

import (
	"archive/zip"
	"io"
	"os"
)

type FileArchive struct {
	zipFileName string
}

func (fa *FileArchive) Archive(names []string, readers ...io.Reader) (outR io.Reader, err error) {
	zippedfile, err := os.Create(fa.zipFileName)
	if err != nil {
		panic(err)
	}
	zipWriter := zip.NewWriter(zippedfile)
	for i, r := range readers {
		w, err := zipWriter.Create(names[i])
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(w, r)
		if err != nil {
			panic(err)
		}
	}
	return zippedfile, nil
}

func New() *FileArchive {
	return &FileArchive{
		zipFileName: "zippedfile.zip",
	}
}
