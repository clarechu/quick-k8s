package service

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"
)

type File struct {
	// Mode 权限设置 644
	Mode os.FileMode
	// Path 下载路径
	Path string
	// FileName 文件名称
	FileName string
}

func NewFile(path, name string) *File {
	return &File{
		Mode:     644,
		Path:     path,
		FileName: name,
	}
}

func (f *File) Wget(uri string) error {
	fmt.Println("downloading file...")

	_, err := url.Parse(uri)

	if err != nil {
		return err
	}
	filename := filepath.Join(f.Path, f.FileName)
	if _, err := os.Stat(filename); err != nil {
		err = os.MkdirAll(f.Path, 644)
		if err != nil {
			return err
		}
	} else {
		err = os.Remove(filename)
		if err != nil {
			return err
		}

	}
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = file.Chmod(f.Mode)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer file.Close()

	checkStatus := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}

	response, err := checkStatus.Get(uri)

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer response.Body.Close()
	fmt.Printf("Request Status: %s\n\n", response.Status)

	filesize := response.ContentLength

	go func() {
		n, err := io.Copy(file, response.Body)
		if n != filesize {
			fmt.Println("Truncated")
		}
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
	}()

	countSize := int(filesize / 1000)
	bar := pb.StartNew(countSize)
	var fi os.FileInfo
	for fi == nil || fi.Size() < filesize {
		fi, _ = file.Stat()
		bar.Set(int(fi.Size() / 1000))
		time.Sleep(time.Millisecond)
	}
	finishMessage := fmt.Sprintf("\n%s with %v bytes downloaded",
		f.FileName, filesize)
	bar.FinishPrint(finishMessage)
	return err
}
