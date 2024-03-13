package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var config map[string]interface{}

func main() {
	if err := readConf(); err != nil {
		fmt.Fprintf(os.Stderr, "配置文件错误：%v\n", err)
	}
	http.HandleFunc("/gpt", gpt)
	http.HandleFunc("/onlineImage", onlineImage)
	http.HandleFunc("/status", status)
	http.HandleFunc("/rcon", listenrcon)
	log.Fatal(http.ListenAndServe(":1314", nil))
}
