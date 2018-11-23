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
	"path/filepath"
	"strings"
	"time"
)

func substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length
	if l > len(runes) {
		l = len(runes)
	}
	return string(runes[pos:l])
}

//删除目录下的文件信息
//dirpath 目录路径
func delDirFile(dirpath string) {
	//读取目录信息
	dir, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return
	}

	//当前系统的时间
	ct := int32(time.Now().Unix())
	//72小时
	spt := int32(72 * 3600)

	for _, file := range dir {
		//读取到的是目录
		if file.IsDir() {
			// subDir := dirpath + `/` + file.Name()
			subDir := fmt.Sprintf("%s/%s", dirpath, file.Name())
			// dir2, _ := ioutil.ReadDir(subDir)
			// fmt.Println(subDir, dir2)
			delDirFile(subDir)
			// continue
		}
		//文件的最后修改时间
		tdate := file.ModTime()
		ft := int32(tdate.Unix())

		//3天前的文件就删除
		if ft < ct-spt {
			os.Remove(dirpath + "/" + file.Name())
			fmt.Println("del file:", dirpath+"/"+file.Name())
		}
	}
}

//getDirAllFile 获取目录下所有文件
//dirpath 目录路径
func getDirAllFile(dirpath string) {
	//读取目录信息
	dir, err := ioutil.ReadDir(dirpath)
	if err != nil {
		return
	}

	//当前系统的时间
	ct := int32(time.Now().Unix())
	//72小时
	spt := int32(72 * 3600)

	for _, file := range dir {
		//读取到的是目录
		if file.IsDir() {
			continue
		}
		//文件的最后修改时间
		tdate := file.ModTime()
		ft := int32(tdate.Unix())

		//3天前的文件就删除
		if ft < ct-spt {
			os.Remove(dirpath + "/" + file.Name())
		}
		fmt.Println("del file:", dirpath+"/"+file.Name())
	}
}

func copy(src, dst string) (int64, error) {
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

func getParentDirectory(dirctory string) string {
	return substr(dirctory, 0, strings.LastIndex(dirctory, "/"))
}
func getThisDirName(dirctory string) string {
	ll := strings.LastIndex(dirctory, "/") + 1
	al := len(dirctory)
	return substr(dirctory, ll, al-ll)
}

func getCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)
	return strings.Replace(dir, "\\", "/", -1)
}

// 当前程序执行文件夹名
var thisDirName = ""

// 抓取封面文件夹
var coverDirName = "cover"

// 输出图片目录
var echoDirs = []string{`a`, `b`, `c`}
