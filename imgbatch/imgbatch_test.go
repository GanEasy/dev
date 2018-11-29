// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imgbatch

import (
	"fmt"
	"path"
	"strings"
	"testing"
)

func Test_TM(t *testing.T) {

	var str1, str2 string
	str1 = GetCurrentDirectory()

	str2 = GetThisDirName(str1)

	t.Fatal(str1, str2)
}

func Test_AllFile(t *testing.T) {

	var str1, str2 string
	str1 = GetCurrentDirectory()

	str2 = GetThisDirName(str1)

	dir := `E:\web\www\93zp_com\Uploads\file\2015-12-19/`

	dirName := GetThisDirName(dir)

	dir = path.Dir(strings.Replace(dir, "\\", "/", -1))

	_, n := path.Split(dir)
	t.Fatal(str1, str2, dirName, n)
}

func Test_Copy(t *testing.T) {

	var str1, str2 string
	str1 = `F:\yize\go\src\github.com\GanEasy\dev\imgbatch\test\testcopy.txt`
	str2 = `F:\yize\go\src\github.com\GanEasy\dev\imgbatch\test\to\copy.txt`

	// filePath := path.Dir(strings.Replace(str2, "\\", "/", -1))
	// if !isExist(filePath) {
	// 	err := os.MkdirAll(filePath, os.ModePerm)
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// }

	tt, _ := Copy(str1, str2)

	t.Fatal(tt)
}

func Test_GetDirAllImgFile(t *testing.T) {

	// 输出图片目录
	var echoDirs = []string{`ADP`, `OEM`, `JXE`, `BST`, `HIP`, `DVA`, `IDC`, `APP`, `GPS`, `WXP`, `JAM`}

	//
	var from, to, dirName string
	dir := `E:\web\www\93zp_com\Uploads\file\CJ20181124`

	dirName = GetThisDirName(dir + `/`)
	tt, _ := GetDirAllImgFile(dir)
	for _, ed := range echoDirs {
		i := 1
		for _, file := range tt {
			i++
			from = fmt.Sprintf("%s/%s", dir, file.Name())
			to = fmt.Sprintf("%s/%s/%s%s-%v%s", dir, ed, dirName, ed, i, path.Ext(file.Name()))
			// copy(f,t )
			t.Fatal(ed, from, to)
		}

	}
	t.Fatal(tt)
}
