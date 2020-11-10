package upload

import (
	"io"
	"mime/multipart"
	"os"
	"path"

	"github.com/closetool/blog/system/constants"
	"github.com/closetool/blog/system/reply"
	"github.com/closetool/blog/utils/fileutils"
	"github.com/sirupsen/logrus"
)

func init() {
	Services[constants.DefaultType] = DefaultUploadService{}
}

var Services = map[string]UploadService{}

type UploadService interface {
	Check(*multipart.FileHeader) bool
	SaveFile(*multipart.FileHeader, string) string
}

type DefaultUploadService struct {
}

func (_ DefaultUploadService) Check(*multipart.FileHeader) bool {
	return true
}

func (_ DefaultUploadService) SaveFile(fh *multipart.FileHeader, filePath string) string {
	fileName := fileutils.RandomFileName(fh.Filename)
	_, err := os.Stat(filePath)
	if err != nil {
		if err := os.MkdirAll(filePath, 0644); err != nil {
			logrus.Errorf("can not make directories: %v", err)
			panic(reply.Error)
		}
	}
	dst := path.Join(filePath, fileName)
	fd, err := os.OpenFile(dst, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		logrus.Errorf("can not make directories: %v", err)
		panic(reply.Error)
	}
	file, err := fh.Open()
	if err != nil {
		logrus.Errorf("can not open uploaded: %v", err)
		panic(reply.Error)
	}
	if _, err := io.Copy(fd, file); err != nil {
		logrus.Errorf("copy uploaded file to local failed: %v", err)
		panic(reply.Error)
	}

	return fileName
}
