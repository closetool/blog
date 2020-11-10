package fileutils

import (
	"fmt"
	"path"

	uuid "github.com/satori/go.uuid"
)

func RandomFileName(oldName string) string {
	ext := path.Ext(oldName)
	id := uuid.NewV4()
	return fmt.Sprintf("%s%s", id.String(), ext)
}
