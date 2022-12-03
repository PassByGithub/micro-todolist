package logging

import (
	"fmt"
	"log"
	"os"
	"time"
)

var (
	LogSavePath = "runtime/logs/"
	LogSaveNmae = "log"
	LogFileExt  = "log"
	Timeformat  = "20060102"
)

func getLogFilePath() string {
	return fmt.Sprintf("%s", LogSavePath)
}

func getLogFileFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveNmae, time.Now().Format(Timeformat), LogFileExt)

	//return string be like : runtime/logs/log20221127.log
	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(path string) *os.File {
	_, err := os.Stat(path)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission is denied:%v", err)
	}

	handle, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to open file:%v", err)
	}

	return handle

}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
