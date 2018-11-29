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
	tt, _ := imgbatch.GetDirAllImgFile(dir)
	for _, ed := range echoDirs {
		i = 1
		for _, file := range tt {
			i++
			from = fmt.Sprintf("%s/%s", dir, file.Name())
			to = fmt.Sprintf("%s/%s/%s%s-1-%v%s", dir, ed, dirName, ed, i, path.Ext(file.Name()))
			imgbatch.Copy(from, to)
		}

	}

	i = 1
	for _, file := range tt {
		i++
		from = fmt.Sprintf("%s/%s", dir, file.Name())
		to = fmt.Sprintf("%s/%s/%s-1-%v%s", dir, `全部图片`, dirName, i, path.Ext(file.Name()))
		imgbatch.Copy(from, to)
	}

}
