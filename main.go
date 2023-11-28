package main

import (
	"flag"
	"path/filepath"
	"runtime"
	"web_tpl1/app"
	"web_tpl1/app/http"
)

func main() {

	homePath := flag.String("prjHome", "", "项目的根目录路径")
	flag.Parse()

	if *homePath == "" {
		_, f, _, ok := runtime.Caller(0)
		if !ok {
			panic("尝试获取文件路径失败")
		}
		*homePath = filepath.Dir(f)
	}

	err := app.Init(*homePath)
	if err != nil {
		panic(err)
	}

	err = http.NewServer()
	if err != nil {
		panic(err)
	}

}
