package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func gpt(w http.ResponseWriter, r *http.Request) {
	//获取查询字符串参数
	params := r.URL.Query()
	text := params.Get("text")
	if text == "" {
		http.Error(w, "Require parameter.", http.StatusBadRequest)
		return
	}
	model := "gpt-3.5-turbo"
	if params.Get("model") == "4" {
		model = "gpt-4"
	}
	temperature := 0.7
	if inputed := params.Get("temperature"); inputed != "" {
		temp, err := strconv.ParseFloat(inputed, 64)
		if err != nil {
			http.Error(w, "Temperature invalid.", http.StatusBadRequest)
		}
		temperature = temp
	}

	//制作向GPT发送的请求
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{"role": "user", "content": text},
		},
		"temperature": temperature,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	gpt_url := "https://api.openai.com/v1/chat/completions"
	gpt_request := strings.NewReader(string(requestBody))
	req, err := http.NewRequest("POST", gpt_url, gpt_request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	gptkey, ok := config["gptkey"].(string)
	if !ok {
		log.Fatal("Gptkey is not a string.")
	}
	req.Header.Set("Authorization", "Bearer "+gptkey)

	//向GPT发送请求之后向用户发送请求
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
