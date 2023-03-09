package main

import (
	"download/web"
	"download/zip"
	"log"
	"os"
)

func main() {
	d := web.NewUrlDownloader()
	_, err := d.Download(d.URL)
	if err != nil {
		panic(err)
	}
	f := web.NewFileSystemDownloader()
	_, err = f.Download(f.SourceFileName)
	if err != nil {
		panic(err)
	}
	f1, err := os.OpenFile("sample1.pdf", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}
	f2, err := os.OpenFile("sample2.JPG", os.O_RDONLY, 0400)
	if err != nil {
		log.Fatal(err)
	}

	zipper := zip.New()
	_, err = zipper.Archive([]string{"f1.pdf", "f2.JPG"}, f1, f2)
	if err != nil {
		panic(err)
	}
	f1.Close()
	f2.Close()
}
