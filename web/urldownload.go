package web

import (
	"errors"
	"io"
	"net/http"
	"os"
)

type UrlDownloader struct {
	URL      string
	fileName string
}

func (data *UrlDownloader) Download(url string) (r io.Reader, err error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("non 200 statuscode")
	}
	file, err := os.Create(data.fileName)

	if err != nil {
		return nil, err
	}

	defer file.Close()
	_, err = io.Copy(file, response.Body)

	if err != nil {
		return nil, err
	}
	return file, nil
}

func NewUrlDownloader() *UrlDownloader {
	return &UrlDownloader{
		URL:      "https://www.africau.edu/images/default/sample.pdf",
		fileName: "sample1.pdf",
	}
}
