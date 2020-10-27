package main

import "fmt"

type Video struct {
	titulo  string
	formato string
	frames  uint32
}

func (video *Video) mostrar() {
	fmt.Println(video.titulo, video.formato, video.frames)
}
