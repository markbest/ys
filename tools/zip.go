package tools

import (
	"bytes"
	"archive/zip"
	"log"
	"os"
	"io/ioutil"
)

type File struct {
	Name string
	Body string
}

func Zip(path string, name string) {
	// 创建一个缓冲区用来保存压缩文件内容
	buf := new(bytes.Buffer)

	// 创建一个压缩文档
	w := zip.NewWriter(buf)

	// 将文件加入压缩文档
	all_files := make([]File, 0)
	files := GetFileList(path)
	for _, full_file := range files {
		dat, _ := ioutil.ReadFile(full_file)
		all_files = append(all_files, File{full_file, string(dat)})
	}

	for _, file := range all_files {
		f, err := w.Create(file.Name)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte(file.Body))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 关闭压缩文档
	err := w.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 将压缩文档内容写入文件
	f, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	buf.WriteTo(f)
}
