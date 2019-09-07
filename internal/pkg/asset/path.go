package asset

import (
	"os"
	"path"
	"strings"
)


const (
	UserHomeDirFlag = "~"
)


func ExpandUserDir(i string) string {
	if ! strings.HasPrefix(i, UserHomeDirFlag) {
		return i
	}

	userDir, err := os.UserHomeDir()
	if nil != err {
		// this error means current platform is not recognized by Golang.
		// See documents for os.UserHomeDir for detail.
		return i
	}

	i = i[1:]  // Remove prefix.
	i = path.Join(userDir, i)

	return i
}
