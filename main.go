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
	if err := readConf(); err != nil {
		log.Fatal("配置文件读取失败：", err)
	}
	if err := readDB(); err != nil {
		log.Fatal("Read DB failed: ", err)
	}
	http.HandleFunc("/gpt", gpt)
	http.HandleFunc("/prompt", prompt)
	http.HandleFunc("/status", status)
	http.HandleFunc("/rcon", listenrcon)
	http.HandleFunc("/versions", versions)
	log.Fatal(http.ListenAndServe(":1314", nil))
}
