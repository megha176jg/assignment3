package web

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type fileSystemDownloader struct {
	SourceFileName string
	DestFileName   string
}

func (fs *fileSystemDownloader) Download(source string) (r io.Reader, err error) {

	file, err := os.Create(fs.DestFileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	input, err := ioutil.ReadFile(fs.SourceFileName)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = ioutil.WriteFile(fs.DestFileName, input, 0644)
	if err != nil {
		fmt.Println("error while creating destination file")
		fmt.Println(err)
		return nil, err
	}

	return file, nil
}
func NewFileSystemDownloader() *fileSystemDownloader {
	return &fileSystemDownloader{
		SourceFileName: "/Users/megha.gupta/Downloads/Megha_Gurgaon.JPG",
		DestFileName:   "sample2.JPG",
	}
}
