package main

import "fmt"

type Audio struct {
	titulo   string
	formato  string
	duracion uint32
}

func (audio *Audio) mostrar() {
	fmt.Println(audio.titulo, audio.formato, audio.duracion)
}
