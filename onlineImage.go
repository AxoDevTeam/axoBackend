package main

import (
	"encoding/json"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"net/http"
	"os"
	"strconv"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
)

func onlineImage(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://127.0.0.1:8000/status")
	if err != nil {
		http.Error(w, "调用在线人数API失败", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "转换请求体失败", http.StatusInternalServerError)
		return
	}

	var ol_data map[string]interface{}
	if err := json.Unmarshal(body, &ol_data); err != nil {
		http.Error(w, "读取请求体失败", http.StatusInternalServerError)
		return
	}

	ol_count := 0
	if player, ok := ol_data["players"].(map[string]interface{}); ok {
		if temp, ok := player["online"].(int); ok {
			ol_count = temp
		}
	}

	bgfile, err := os.Open("onlineBG.png")
	if err != nil {
		http.Error(w, "服务器没有图片", http.StatusInternalServerError)
		return
	}
	defer bgfile.Close()

	img, err := png.Decode(bgfile)
	if err != nil {
		http.Error(w, "读取图片失败", http.StatusInternalServerError)
		return
	}

	bounds := img.Bounds()
	newImg := image.NewRGBA(bounds)
	draw.Draw(newImg, bounds, img, bounds.Min, draw.Src)
	fonturi, ok := config["mcfont"].(string)
	if !ok {
		http.Error(w, "字体路径格式有误", http.StatusInternalServerError)
		return
	}
	fontBytes, err := os.ReadFile(fonturi)
	if err != nil {
		http.Error(w, "未找到字体", http.StatusInternalServerError)
		return
	}
	f, err := opentype.Parse(fontBytes)
	if err != nil {
		http.Error(w, "字体文件有误", http.StatusInternalServerError)
		return
	}
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    18,
		DPI:     72,
		Hinting: font.HintingNone,
	})
	if err != nil {
		http.Error(w, "字体设置失败", http.StatusInternalServerError)
		return
	}

	d := &font.Drawer{
		Dst:  newImg,
		Src:  image.NewUniform(color.RGBA{255, 255, 255, 255}),
		Face: face,
		Dot:  fixed.Point26_6{X: fixed.Int26_6(1000), Y: fixed.Int26_6(2000)},
	}

	d.DrawString("当前在线 " + strconv.Itoa(ol_count) + " 人")
	d.Dot = fixed.Point26_6{X: fixed.Int26_6(1000), Y: fixed.Int26_6(3500)}
	versions, ok := config["ver"].(map[string]interface{})
	if !ok {
		http.Error(w, "类型断言失败", http.StatusInternalServerError)
		return
	}
	mcver, ok := versions["mainbe"].(string)
	if !ok {
		http.Error(w, "类型断言失败", http.StatusInternalServerError)
		return
	}
	d.DrawString("版本：" + mcver)
	d.Dot = fixed.Point26_6{X: fixed.Int26_6(1000), Y: fixed.Int26_6(5000)}
	d.DrawString("点我启动游戏添加服务器")

	w.Header().Set("Content-type", "image/png")
	if err := png.Encode(w, newImg); err != nil {
		http.Error(w, "图片发送失败", http.StatusInternalServerError)
	}
	return
}
