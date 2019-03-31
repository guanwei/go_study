// lissajous 产生随机利萨茹图形的GIF动画
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{0, 0, 128, 255}, color.RGBA{30, 144, 255, 255}, color.RGBA{72, 61, 139, 255}, color.RGBA{255, 0, 255, 255}, color.RGBA{124, 252, 0, 255}}

const (
	whiteIndex = 0 // 画板中的第一种颜色
	blackIndex = 1 // 画板中的下一种颜色
	greenIndex = 2 // 定义新的颜色
)

type lconfig struct {
	cycles  int     // 完整的x振荡器变化的个数
	res     float64 // 角度分辨率
	freq    float64 // y振荡器的相对频率
	size    int     // 角度分辨率
	nframes int     // 图像画布包含[-size..+size]
	delay   int     // 以10ms为单位的帧间延迟
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lconf := lconfig{
		cycles:  5,
		res:     0.001,
		freq:    rand.Float64() * 3.0,
		size:    100,
		nframes: 64,
		delay:   8,
	}

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				log.Print(err)
			}
			for k, v := range r.Form {
				switch k {
				case "cycles":
					lconf.cycles, _ = strconv.Atoi(v[0])
				case "res":
					lconf.res, _ = strconv.ParseFloat(v[0], 64)
				case "freq":
					lconf.freq, _ = strconv.ParseFloat(v[0], 64)
				case "size":
					lconf.size, _ = strconv.Atoi(v[0])
				case "nframes":
					lconf.nframes, _ = strconv.Atoi(v[0])
				case "delay":
					lconf.delay, _ = strconv.Atoi(v[0])
				}
			}
			lissajous(w, lconf)
		}
		http.HandleFunc("/", handler)
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout, lconf)
}

func lissajous(out io.Writer, lconf lconfig) {
	anim := gif.GIF{LoopCount: lconf.nframes}
	phase := 0.0 // phase difference
	for i, c := 0, 1; i < lconf.nframes; i++ {
		rect := image.Rect(0, 0, 2*lconf.size+1, 2*lconf.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(lconf.cycles)*2*math.Pi; t += lconf.res {
			x := math.Sin(t)
			y := math.Sin(t*lconf.freq + phase)
			img.SetColorIndex(lconf.size+int(x*float64(lconf.size)+0.5), lconf.size+int(y*float64(lconf.size)+0.5),
				uint8(c))
		}
		c++
		if c >= len(palette) {
			c = 1
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, lconf.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // 注意：忽略编码错误
}
