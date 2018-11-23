// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imgbatch

import (
	"testing"
)

func Test_TM(t *testing.T) {

	var str1, str2 string
	str1 = getCurrentDirectory()

	str2 = getThisDirName(str1)

	t.Fatal(str1, str2)
}

func Test_AllFile(t *testing.T) {

	var str1, str2 string
	str1 = getCurrentDirectory()

	str2 = getThisDirName(str1)

	t.Fatal(str1, str2)
}

func Test_Copy(t *testing.T) {

	var str1, str2 string
	str1 = `F:\yize\go\src\github.com\GanEasy\dev\imgbatch\test\testcopy.txt`
	str2 = `F:\yize\go\src\github.com\GanEasy\dev\imgbatch\test\to\copy.txt`

	tt, _ := copy(str1, str2)

	t.Fatal(tt)
}
