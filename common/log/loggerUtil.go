package clog

import (
	"log"
	"os"
)

func Info(v ...interface{}) {
	log.New(os.Stdout,"Info:",log.Ldate | log.Ltime | log.Lshortfile).Println(v)
}


func Warning(v ...interface{}) {
	log.New(os.Stdout,"Warning:",log.Ldate | log.Ltime | log.Lshortfile).Println(v)
}


// 取消读取文件IO
func Error(v ...interface{}) {
	log.New(os.Stdout,"Error:",log.Ldate | log.Ltime | log.Lshortfile).Println(v)
}

