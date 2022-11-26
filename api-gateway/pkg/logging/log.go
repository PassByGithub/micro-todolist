package logging

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG level = iota
	INFO
	WARN
	ERROR
	FATAL
)

func init() {
	filepath := getLogFileFullPath()
	F = openLogFile(filepath)

	logger := log.New(F, DefaultPrefix, log.LstdFlags)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v...)
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v...)
}
func Warn(v ...interface{}) {
	setPrefix(WARN)
	logger.Println(v...)
}
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v...)
}
func Fatalf(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v...)

}

func setPrefix(level level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}
