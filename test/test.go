package main

import (
	"fmt"
	"log"

	"github.com/disintegration/imaging"
)

func main() {
	// 打开ICO文件
	img, err := imaging.Open("D:/H3C/jd_analysis/public/img/oem/H3C/favicon.ico")
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// 获取ICO文件的宽高
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()

	fmt.Printf("ICO文件的宽度：%d, 高度：%d\n", width, height)
}
