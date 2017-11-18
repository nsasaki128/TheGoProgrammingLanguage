package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"io"
	"image/gif"
	"image"
	"math"
	"math/rand"
	"image/color"
	"strconv"
	"net/url"
)

var palette = []color.Color{color.White, color.RGBA{0x00, 0xff, 0x00, 0xff}}
const (
	blackIndex = 0 //パレットの最初の色
	greenIndex = 1 //パレットの次の色
)


func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err!=nil {
			log.Print(err)
		}
		lissajous(w, r.Form)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}



func lissajous(out io.Writer, form url.Values) {
	rand.Seed(time.Now().UTC().UnixNano())

	cycles  := 5     // 発信器xが完了する周回の回数
	res     := 0.001 // 回転の分解能
	size    := 100   // 画像キャンパスは [-size..+size] の範囲で扱う
	nframes := 64    // アニメーションフレーム数
	delay   := 8     // 10ms単位でのフレーム間の遅延

	for k, v := range form {
		if i, err := strconv.Atoi(v[0]); err == nil {
			switch k {
			case "cycles" : cycles = i
			case "size" : size = i
			case "nframes" : nframes = i
			case "delay" : delay = i
			}
		}

		if f64, err := strconv.ParseFloat(v[0], 64); err == nil {
			switch k {
			case "res": res = f64
			}
		}
	}

	freq := rand.Float64() * 3.0 // 発信器yの相対周波数
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 //位相差

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img  := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*float64(size)+0.5), size+int(y*float64(size)+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	gif.EncodeAll(out, &anim) //注意: エンコードエラーを無視
}