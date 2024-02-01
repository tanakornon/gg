package gg

import (
	"encoding/base64"
	"regexp"
)

type File struct {
	Bytes       []byte
	Name        string
	ContentType string
	Extension   string
}

const (
	ErrorInvalidBase64Encoding = ConstError("invalid base64 encoding")
)

var MimeExtensions = map[string]string{
	"image/png":                     "png",
	"image/jpeg":                    "jpg",
	"image/jpg":                     "jpg",
	"image/gif":                     "gif",
	"image/bmp":                     "bmp",
	"image/svg+xml":                 "svg",
	"audio/mpeg":                    "mp3",
	"audio/wav":                     "wav",
	"audio/ogg":                     "ogg",
	"video/mp4":                     "mp4",
	"video/webm":                    "webm",
	"video/ogg":                     "ogg",
	"application/pdf":               "pdf",
	"application/msword":            "doc",
	"application/vnd.ms-excel":      "xls",
	"application/vnd.ms-powerpoint": "ppt",
	"text/plain":                    "txt",
	"text/html":                     "html",
	"text/css":                      "css",
	"text/javascript":               "js",
}

func NewFileFromBase64(input string) (file File, err error) {
	pattern := `data:([^;]+);base64,([^"]*)`
	re := regexp.MustCompile(pattern)

	matches := re.FindStringSubmatch(input)

	if len(matches) < 3 {
		return File{}, ErrorInvalidBase64Encoding
	}

	file.ContentType = matches[1]
	file.Extension = MimeExtensions[file.ContentType]
	file.Bytes, err = base64.StdEncoding.DecodeString(matches[2])

	if err != nil {
		return File{}, ErrorInvalidBase64Encoding
	}

	return file, nil
}

func (file File) EncodeBase64() string {
	return "data:" + file.ContentType + ";base64," + base64.StdEncoding.EncodeToString(file.Bytes)
}
