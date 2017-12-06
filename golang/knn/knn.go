package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"path"
	"path/filepath"
)

var debug bool = false
var data_dir string = "./moyan"        //文件存放目录
var limen float64 = 0.1159203888322267 //阈值

const (
	MIN_HANZI rune = 0x3400
	MAX_HANZI rune = 0x9fbb
)

var labels []rune = []rune{
	0x817f, 0x80f8, 0x4e73, 0x81c0,
	0x5c41, 0x80a1, 0x88f8, 0x6deb,
}

func errHandle(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func load(name string) (m map[rune]int, err error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	m = make(map[rune]int)
	var r rune
	for {
		r, _, err = buf.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if r >= MIN_HANZI && r <= MAX_HANZI {
			m[r] += 1
		}
	}
	return m, nil
}
func classify(m map[rune]int) (idv []float64, dis float64) {
	len_m := len(m)
	for i, v := range labels {
		if debug {
			fmt.Println(i, m[v], string(v), float64(m[v])/float64(len_m))
		}
		idv = append(idv, float64(m[v])/float64(len_m))
	}
	for _, v := range idv {
		dis += math.Pow(v, 2)
	}
	dis = math.Sqrt(dis)
	return
}
func check(fp string, dis float64) {
	switch {
	case dis >= limen:
		fmt.Println(fp, dis, "涉黄")
	case dis == 1.0:
		fmt.Println(fp, dis, "你在作弊吗")
	case dis == 0:
		fmt.Println(fp, dis, "检查一下文件字符编码是不是utf8格式吧")
	default:
		fmt.Println(fp, dis, "正常")
	}
}

func walkFunc(fp string, info os.FileInfo, err error) error {
	if path.Ext(fp) == ".txt" {
		m, err := load(fp)
		errHandle(err)
		_, dis := classify(m)
		check(fp, dis)
	}
	return err
}

var file string

func init() {
	_, err := os.Stat(data_dir)
	if err != nil {
		err = os.Mkdir(data_dir, os.ModePerm)
		errHandle(err)
	}
	flag.StringVar(&file, "file", "", "file read in,if you don't give the file read in,"+
		"it will create a data dictionary,just pust your files in it")
}

func main() {
	flag.Parse()
	if file == "" {
		filepath.Walk(data_dir, walkFunc)
		return
	}
	m, err := load(file)
	errHandle(err)
	_, dis := classify(m)
	check(file, dis)

}
