package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/go-ffmt/ffmt"
	"github.com/karrick/godirwalk"
)

type lineResult struct {
	date string
	text string
	link string
	pswd string
}

func getFiles(rootdir string) []string {
	files := []string{}
	err := godirwalk.Walk(rootdir, &godirwalk.Options{
		Callback: func(osPathname string, de *godirwalk.Dirent) error {
			fi, _ := os.Stat(osPathname)
			if !fi.IsDir() {
				files = append(files, osPathname)
			}
			return nil
		},
		Unsorted: false,
	})
	if err != nil {
		fmt.Println(err)
	}
	return files
}

func readFileContent(filepath string) string {
	content := ""

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer f.Close()

	fd, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return ""
	}
	content = string(fd)

	return content
}

func readFileLines(filepath string) []string {
	lines := []string{}

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		b, errR := buf.ReadBytes('\n')
		if errR != nil {
			if errR == io.EOF {
				break
			}
			fmt.Println(errR.Error())
		}
		lines = append(lines, string(b))
	}
	return lines
}

func processLine(line string) lineResult {
	res := lineResult{}
	reg := regexp.MustCompile(`https://pan.baidu.com/s/[\w\-]+`)
	if len(reg.FindAllString(line, -1)) > 0 {
		res.link = reg.FindAllString(line, -1)[0]
	}
	reg = regexp.MustCompile(`提取码：(\w{4,4})`)
	if len(reg.FindAllString(line, -1)) > 0 {
		res.pswd = reg.FindSubmatch(line, -1)[0]
	}
	fmt.Println(res)
	return res
}

func main() {
	dir := "/home/ulric/workspace/codes/go/src"
	fs := getFiles(dir)
	ffmt.Puts(fs)
	// ffmt.Puts(readFileLines(fs[0]))
	// fmt.Println(readFileContent(fs[0]))
	processLine("https://pan.baidu.com/s/sd_0-1  提取码：1234")
}
