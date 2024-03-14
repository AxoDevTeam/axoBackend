package main

import (
	"fmt"
	"net/http"

	"github.com/gorcon/rcon"
)

func listenrcon(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	srv, cmd, pwd := params.Get("srv"), params.Get("cmd"), params.Get("pwd")
	var conn *rcon.Conn
	var err error
	switch srv {
	case "sc":
		conn, err = rcon.Dial("192.168.50.38:25577", pwd)
	case "mod":
		conn, err = rcon.Dial("192.168.50.38:25574", pwd)
	default:
		conn, err = rcon.Dial("192.168.50.38:25575", pwd)
	}
	if err != nil {
		http.Error(w, "密码不正确", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	response, err := conn.Execute(cmd)
	if err != nil {
		http.Error(w, "命令执行失败", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, response)
}
