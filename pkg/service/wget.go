package service

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

func Wget(uri string) error {
	fmt.Println("downloading file...")

	fileUrl, err := url.Parse(uri)

	if err != nil {
		return err
	}

	filePath := fileUrl.Path
	segments := strings.Split(filePath, "/")
	fileName := segments[len(segments)-1]
	file, err := os.Create(fileName)

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
		fileName, filesize)
	bar.FinishPrint(finishMessage)
	return err
}
