package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var config map[string]interface{}
var db *sql.DB

func main() {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal("读取程序所在路径失败：", err)
	}
	if os.Chdir(filepath.Dir(exePath)) != nil {
		log.Fatal("更改基准目录失败：", err)
	}
	if readConf() != nil {
		log.Fatal("配置文件读取失败：", err)
	}
	if readDB() != nil {
		log.Fatal("读取数据库失败：", err)
	}
	http.Handle("/gpt", http.HandlerFunc(gpt))
	http.Handle("/prompt", http.HandlerFunc(prompt))
	http.Handle("/status", http.HandlerFunc(status))
	http.Handle("/rcon", http.HandlerFunc(listenrcon))
	http.Handle("/versions", http.HandlerFunc(versions))
	log.Fatal(http.ListenAndServe(":1314", nil))
}
