package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const (
	//待删除文件的目录信息
	CONFFILEPATH string = "./conf.txt"
)

func main() {
	for {
		//读取目录配置文件信息
		dirinfo, err := ioutil.ReadFile(CONFFILEPATH)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		//解析数据：一行一个目录形式
		dirstr := string(dirinfo)
		dirarr := strings.Split(dirstr, "\r\n")

		//遍历删除目录下的文件
		for _, val := range dirarr {
			if strings.TrimSpace(val) == "" {
				continue
			}
			delDirFile(val)
		}

		//休眠1小时
		time.Sleep(time.Hour * 24)
		// time.Sleep(3600e9)
	}
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
