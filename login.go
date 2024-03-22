package main

import "net/http"

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "只接受POST请求", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "读取表单数据失败", http.StatusInternalServerError)
		return
	}
	username, password := r.FormValue("username"), r.FormValue("password")
}
