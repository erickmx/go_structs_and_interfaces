package main

import "fmt"

type Imagen struct {
	titulo  string
	formato string
	canales uint32
}

func (image *Imagen) mostrar() {
	fmt.Println(image.titulo, image.formato, image.canales)
}
