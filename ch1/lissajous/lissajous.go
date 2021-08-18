package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

// 定义一个颜色数组
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func Main() {
	lissajous(os.Stdout)
}

// 绘制图形
func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	// 创建一个固定帧数的gif动画
	anim := gif.GIF{LoopCount: nframes}
	phrase := 0.0
	// 绘制帧
	for j := 0; j < nframes; j++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// 根据规格新建画布
		img := image.NewPaletted(rect, palette)
		// 正弦函数周期
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			// 根据周期绘制正弦函数
			x := math.Sin(t)
			// 随机频率 + 移动相位
			y := math.Sin(t*freq + phrase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phrase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
