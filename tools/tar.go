package tools

import (
	"bytes"
	"archive/tar"
	"log"
	"os"
	"io/ioutil"
)

func Tar(path string, name string) {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new tar archive.
	tw := tar.NewWriter(buf)

	// Add some files to the archive.
	allFiles := make([]File, 0)
	files := GetFileList(path)
	for _, full_file := range files {
		dat, _ := ioutil.ReadFile(full_file)
		allFiles = append(allFiles, File{full_file, string(dat)})
	}
	for _, file := range allFiles {
		hdr := &tar.Header{
			Name: file.Name,
			Mode: 0600,
			Size: int64(len(file.Body)),
		}
		if err := tw.WriteHeader(hdr); err != nil {
			log.Fatalln(err)
		}
		if _, err := tw.Write([]byte(file.Body)); err != nil {
			log.Fatalln(err)
		}
	}
	// Make sure to check the error on Close.
	if err := tw.Close(); err != nil {
		log.Fatalln(err)
	}

	// 将压缩文档内容写入文件
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	buf.WriteTo(f)
}
