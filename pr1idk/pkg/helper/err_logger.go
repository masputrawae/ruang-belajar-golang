package helper

import (
	"fmt"
	"log"
	"runtime"
	"strings"
)

func FileLine(err error) string {
	sprt := "------------------------------"
	msg := ""
	for i := 1; i < 10; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		if !strings.Contains(file, "/helper/") {
			msg = fmt.Sprintf("\n%s\nFILE: %s\nLINE: %d\nERROR: %s\n%s\n", sprt, file, line, err, sprt)
			return msg
		}
	}
	return msg
}

func ErrFatal(err error) {
	log.Fatal(FileLine(err))
}

func ErrLog(err error) {
	log.Println(FileLine(err))
}
