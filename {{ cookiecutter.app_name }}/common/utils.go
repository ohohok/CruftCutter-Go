package common

import (
	"fmt"
	"os"
)

// 创建文件
func MakeFile(path, file_name string) error {
	// check
	if _, err := os.Stat(path); err == nil {
		fmt.Println("directory already exists")
	} else {
		fmt.Println("make directory success", path)
		err := os.MkdirAll(path, 0766)

		if err != nil {
			fmt.Println("failed to create directory")
			return err
		}
	}
	if _, err := os.Stat(path + file_name); err != nil {
		f, err := os.Create(path + file_name)
		if err != nil {
			fmt.Println("failed to create file")
			return err
		}
		defer f.Close()
	}
	return nil
}
