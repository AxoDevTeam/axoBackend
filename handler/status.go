package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/mcstatus-io/mcutil/v3"
	"github.com/mcstatus-io/mcutil/v3/response"
)

var Status = http.HandlerFunc(status)

func status(w http.ResponseWriter, r *http.Request) {
	ctx, canc := context.WithTimeout(context.Background(), time.Second*5)
	defer canc()
	params := r.URL.Query()
	srv := params.Get("srv")
	var resp *response.FullQuery
	var err error
	if srv == "sc" {
		resp, err = mcutil.FullQuery(ctx, "sc.mcax.cn", 25565)
	} else if srv == "mod" {
		resp, err = mcutil.FullQuery(ctx, "mod.mcax.cn", 25565)
	} else {
		resp, err = mcutil.FullQuery(ctx, "mcax.cn", 25565)
	}
	if err != nil {
		http.Error(w, "查询失败", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "JSON编码失败", http.StatusInternalServerError)
		return
	}
}
