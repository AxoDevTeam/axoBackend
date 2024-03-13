package main

import (
	"fmt"
	"net/http"

	"github.com/mcstatus-io/mcutil/v3/rcon"
)

func listenrcon(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	srv := params.Get("srv")
	cmd := params.Get("cmd")
	pwd := params.Get("pwd")

	var ip string
	var port uint16
	switch srv {
	case "sc":
		ip, port = "192.168.50.38", 25577
	case "mod":
		ip, port = "192.168.50.38", 25574
	default:
		ip, port = "192.168.50.38", 25575
	}

	client, err := rcon.Connect(ip, port)
	if err != nil {
		http.Error(w, fmt.Sprintf("执行失败: %v", err), http.StatusInternalServerError)
		return
	}
	defer func() {
		if closeErr := client.Close(); closeErr != nil {
			fmt.Printf("关闭连接时发生错误: %v\n", closeErr)
		}
	}()

	if err := client.Login(pwd); err != nil {
		http.Error(w, fmt.Sprintf("密码错误: %v", err), http.StatusBadRequest)
		return
	}

	if err = client.Run(cmd); err != nil {
		http.Error(w, fmt.Sprintf("命令执行失败: %v", err), http.StatusBadRequest)
		return
	}
}
