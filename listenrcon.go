package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/gorcon/rcon"
)

const chalVolume = 100

var challenges [chalVolume]string
var chalIte = 0

func listenrcon(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	//对GET请求返回挑战随机数
	case http.MethodGet:
		randBytes := make([]byte, 16)
		if _, err := rand.Read(randBytes); err != nil {
			http.Error(w, "随机数生成失败", http.StatusInternalServerError)
			return
		}
		randStr := hex.EncodeToString(randBytes)
		challenges[chalIte] = randStr
		chalIte = (chalIte + 1) % chalVolume
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, randStr)

	//对POST请求验证并执行命令
	case http.MethodPost:

		//解析表单数据
		if err := r.ParseForm(); err != nil {
			http.Error(w, "解析表单数据失败", http.StatusBadRequest)
			return
		}
		srv, cmd := r.FormValue("server"), r.FormValue("command")
		chal, hash := r.FormValue("challenge"), r.FormValue("hash")

		//验证挑战值是否存在
		chalExist := false
		for _, tempChal := range challenges {
			if tempChal == chal {
				chalExist = true
				break
			}
		}
		if chalExist == false {
			http.Error(w, "挑战值不存在", http.StatusBadRequest)
			return
		}

		//验证哈希值是否正确
		var RCONpwd string
		if temp, ok := config["RCONpwd"].(string); ok {
			RCONpwd = temp
		}
		hashBytes := sha256.Sum256([]byte(RCONpwd + chal))
		realhash := hex.EncodeToString(hashBytes[:])
		if realhash != hash {
			http.Error(w, "密码错误", http.StatusBadRequest)
			return
		}

		//执行RCON命令
		var conn *rcon.Conn
		var err error
		switch srv {
		case "sc":
			conn, err = rcon.Dial("192.168.50.38:25577", RCONpwd)
		case "mod":
			conn, err = rcon.Dial("192.168.50.38:25574", RCONpwd)
		default:
			conn, err = rcon.Dial("192.168.50.38:25575", RCONpwd)
		}
		if err != nil {
			http.Error(w, "连接RCON服务器失败", http.StatusInternalServerError)
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
}
