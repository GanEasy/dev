package main

import (
	"fmt"
	"path"

	"github.com/GanEasy/dev/imgbatch"
)

func main() {

	// 当前程序执行文件夹名
	// var thisDirName = ""

	// 抓取封面文件夹
	// var coverDirName = "cover"

	// 输出图片目录
	var echoDirs = []string{`ADP`, `OEM`, `JXE`, `BST`, `HIP`, `DVA`, `IDC`, `APP`, `GPS`, `WXP`, `JAM`}

	//

	var from, to, dirName string
	// dir := `E:\web\www\93zp_com\Uploads\file\CJ20181124`
	dir := imgbatch.GetCurrentDirectory()
	i := 0
	dirName = imgbatch.GetThisDirName(dir + `/`)
	tt, _ := imgbatch.GetDirAllImgFile(dir) // 获取当前目录所有文件
	for _, ed := range echoDirs {
		i = 1
		for _, file := range tt {
			i++
			from = fmt.Sprintf("%s/%s", dir, file.Name())
			to = fmt.Sprintf("%s/%s/%s%s-%v%s", dir, ed, dirName, ed, i, path.Ext(file.Name()))
			imgbatch.Copy(from, to)
		}
	}

	i = 0

	ttc, _ := imgbatch.GetDirAllImgFile(dir + `/1`) // 获取文件夹1(首图文件夹，所有图片)
	i = 0
	for k, file := range ttc {
		// fmt.Println("echoDirs", echoDirs[k])
		i++
		from = fmt.Sprintf("%s/%s", dir+`/1`, file.Name())
		to = fmt.Sprintf("%s/%s/%s-%v%s", dir, echoDirs[k], dirName, 1, path.Ext(file.Name()))
		imgbatch.Copy(from, to)

		from = fmt.Sprintf("%s/%s", dir+`/1`, file.Name())
		to = fmt.Sprintf("%s/%s/%s-%v%s", dir, `全部图片`, dirName, i, path.Ext(file.Name()))
		imgbatch.Copy(from, to)
	}

	for _, file := range tt { // 将所有图片放入
		i++
		from = fmt.Sprintf("%s/%s", dir, file.Name())
		to = fmt.Sprintf("%s/%s/%s-%v%s", dir, `全部图片`, dirName, i, path.Ext(file.Name()))
		imgbatch.Copy(from, to)
	}

}
