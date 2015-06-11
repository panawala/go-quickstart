package helper

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

// Creates a new file upload http request with optional extra params
func generate(params map[string]string, paramName, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.SetBoundary("__X_PAW_BOUNDARY__")
	part, err := writer.CreateFormFile(paramName, filepath.Base(path))
	if err != nil {
		return err
	}
	_, err = io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return err
	}

	outputFile := "post.txt"
	err = ioutil.WriteFile(outputFile, body.Bytes(), 0x664)
	return err
}

func generate_post_txt() {
	ms := map[string]string{
		"title":       "My Document",
		"author":      "Matt Aimonetti",
		"description": "A document with all the Go programming language secrets"}
	generate(ms, "avatar", "/Users/william/Documents/1.png")
}
