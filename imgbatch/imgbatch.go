// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imgbatch

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

// 判断所给路径文件/文件夹是否存在(返回true是存在)
func isExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//getDirAllFile 获取目录下所有文件
//dirpath 目录路径
func GetDirAllImgFile(dirpath string) ([]os.FileInfo, error) {
	//读取目录信息
	dir, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return dir, err
	}

	var ret []os.FileInfo

	for _, file := range dir {
		//读取到的是目录
		if file.IsDir() { //跳过
			continue
		}

		suffix := path.Ext(file.Name())

		//如果是 jpg 或者 png 装起来
		if suffix == `.JPG` || suffix == `.JPEG` || suffix == `.jpg` || suffix == `.PNG` || suffix == `.png` {
			ret = append(ret, file)
		}

	}
	return ret, nil
}

func Copy(src, dst string) (int64, error) {

	tp := path.Dir(strings.Replace(dst, "\\", "/", -1))
	if !isExist(tp) {
		err := os.MkdirAll(tp, os.ModePerm)
		if err != nil {
			return 0, err
		}
	}

	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return 0, err
	}
	if !sourceFileStat.Mode().IsRegular() {
		return 0, fmt.Errorf("%s is not a regular file", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

func GetParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}
func GetThisDirName(dirctory string) string {
	dirctory = path.Dir(strings.Replace(dirctory, "\\", "/", -1))
	ll := strings.LastIndex(dirctory, "/") + 1
	al := len(dirctory)
	fmt.Println(ll, al-ll)
	return substr(dirctory, ll, al-ll)
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	return strings.Replace(dir, "\\", "/", -1)
}
