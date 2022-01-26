package utils

import "os"

// PathExist 判断目录是否存在
//
// @param path string 目录
func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}
