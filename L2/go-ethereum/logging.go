package main

import (
	"bytes"
	"fmt"
	"log"
	"log/slog"
	"os"
)

func main() {
	log.Println("standard logger")

	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	myLog := log.New(os.Stdout, "my", log.LstdFlags)
	myLog.Println("from mylog")
	myLog.SetPrefix("ohmy")
	myLog.Println("from mylog")

	var buf bytes.Buffer
	buflog := log.New(&buf, "buf", log.LstdFlags)
	buflog.Println("hello")
	fmt.Print("from buflog:", buf.String())

	handler := slog.NewJSONHandler(os.Stdout, nil)
	myslog := slog.New(handler)
	myslog.Info("hi there")

	myslog.Info("hello again", "key", "val", "age", 25)
}

//2025/05/18 16:32:33 standard logger
//2025/05/18 16:32:33.246027 with micro
//2025/05/18 16:32:33 main.go:18: with file/line
//my2025/05/18 16:32:33 from mylog
//ohmy2025/05/18 16:32:33 from mylog
//from buflog:buf2025/05/18 16:32:33 hello
//{"time":"2025-05-18T16:32:33.2460277+08:00","level":"INFO","msg":"hi there"}
//{"time":"2025-05-18T16:32:33.2464394+08:00","level":"INFO","msg":"hello again","key":"val","age":25}
