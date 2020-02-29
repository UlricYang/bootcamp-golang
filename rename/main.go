package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"

	"github.com/karrick/godirwalk"
	"github.com/panjf2000/ants"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()
var poolsize = 8

func initLogger() {
	logfile, err := os.OpenFile("rename.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.Out = logfile
	} else {
		logger.Info("Failed to log to file, using default stderr")
	}
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
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
		logger.WithFields(logrus.Fields{
			"dir": rootdir,
		}).Error(err)
	}
	return files
}

func calcFileMD5(filename string) string {
	f, err := os.Open(filename)
	defer f.Close()
	if nil != err {
		logger.WithFields(logrus.Fields{
			"file": filename,
		}).Error(err)
		return ""
	}

	md5Handle := md5.New()
	_, err = io.Copy(md5Handle, f)
	if nil != err {
		logger.WithFields(logrus.Fields{
			"file": filename,
		}).Error(err)
		return ""
	}
	md := md5Handle.Sum(nil)
	md5str := fmt.Sprintf("%x", md)
	return md5str
}

func findPrefix(abspath string) string {
	if filepath.IsAbs(abspath) {
		dirnames := strings.Split(filepath.Dir(abspath), string(os.PathSeparator))
		if len(dirnames) > 0 {
			return dirnames[len(dirnames)-1]
		}
	}
	return ""
}

func renameFile(oldname string, prefix string) bool {
	absname, err := filepath.Abs(oldname)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"file": oldname,
		}).Error(err)
		return false
	}

	if prefix == "" {
		logger.WithFields(logrus.Fields{
			"file": oldname,
		}).Error("prefix is empty")
		return false
	}
	if strings.HasPrefix(path.Base(oldname), prefix) {
		return true
	}

	md5sum := calcFileMD5(oldname)
	if md5sum == "" {
		logger.WithFields(logrus.Fields{
			"file": oldname,
		}).Error("fail to calculate md5")
		return false
	}
	filename := strings.Join([]string{prefix, md5sum + filepath.Ext(absname)}, "-")
	newname := filepath.Join(filepath.Dir(absname), filename)
	os.Rename(absname, newname)
	logger.WithFields(logrus.Fields{
		"oldname": oldname,
		"newname": newname,
	}).Info("successfully rename")
	return true
}

func main() {
	initLogger()

	currentdir, _ := os.Getwd()
	files := getFiles(filepath.Join(currentdir, "files"))

	var wg sync.WaitGroup

	p, _ := ants.NewPoolWithFunc(poolsize, func(f interface{}) {
		file := f.(string)
		renameFile(file, findPrefix(file))
		wg.Done()
	})
	defer p.Release()

	for _, f := range files {
		wg.Add(1)
		_ = p.Invoke(f)
	}
	wg.Wait()
}
