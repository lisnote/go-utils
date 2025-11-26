package errorx

import (
	"log"
	"os"
	"runtime"
)

func Fatal(err error, msg string,) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d %s: %v\n", file, line, msg, err)
		os.Exit(1)
	}
}

func Warn(err error, msg string) bool {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d %s: %v\n", file, line, msg, err)
		return true
	}
	return false
}
