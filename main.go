package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/AxoDevTeam/axoBackend/config"
	"github.com/AxoDevTeam/axoBackend/handler"
)

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("读取程序所在路径失败：", err)
	}
	if os.Chdir(filepath.Dir(exePath)) != nil {
		log.Fatal("更改基准目录失败：", err)
	}
	if config.ReadConf() != nil {
		log.Fatal("配置文件读取失败：", err)
	}
	if config.ReadDB() != nil {
		log.Fatal("读取数据库失败：", err)
	}
	http.Handle("/gpt", handler.Gpt)
	http.Handle("/prompt", handler.Prompt)
	http.Handle("/status", handler.Status)
	http.Handle("/rcon", handler.Listenrcon)
	http.Handle("/versions", handler.Versions)
	log.Fatal(http.ListenAndServe(":1314", nil))
}
