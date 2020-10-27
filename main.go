package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Multimedia interface {
	mostrar()
}

type ContenidoWeb struct {
	multimedia []Multimedia
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var opcStr string
	var strData string
	mediaContent := &ContenidoWeb{
		multimedia: []Multimedia{},
	}

	for opcStr != "5" {
		fmt.Println("Menu principal")
		fmt.Println("1. Agregar imagen")
		fmt.Println("2. Agregar audio")
		fmt.Println("3. Agregar video")
		fmt.Println("4. Mostrar todo")
		fmt.Println("5. Salir")
		fmt.Println("Ingrese una opcion")
		scanner.Scan()
		opcStr = scanner.Text()

		switch opcStr {
		case "1":
			imagen := &Imagen{}
			fmt.Println("Ingrese el nombre de la imagen")
			scanner.Scan()
			strData = scanner.Text()
			imagen.titulo = strData
			fmt.Println("Ingrese el formato de la imagen")
			scanner.Scan()
			strData = scanner.Text()
			imagen.formato = strData
			fmt.Println("Ingrese los canales de la imagen")
			scanner.Scan()
			strData = scanner.Text()

			intData, err := strconv.ParseUint(strData, 10, 32)
			if err != nil {
				fmt.Println("Ingrese un numero valido")
				break
			}
			imagen.canales = uint32(intData)
			mediaContent.multimedia = append(mediaContent.multimedia, imagen)
			break
		case "2":
			imagen := &Audio{}
			fmt.Println("Ingrese el nombre del audio")
			scanner.Scan()
			strData = scanner.Text()
			imagen.titulo = strData
			fmt.Println("Ingrese el formato del audio")
			scanner.Scan()
			strData = scanner.Text()
			imagen.formato = strData
			fmt.Println("Ingrese la duracion del audio")
			scanner.Scan()
			strData = scanner.Text()

			intData, err := strconv.ParseUint(strData, 10, 32)
			if err != nil {
				fmt.Println("Ingrese un numero valido")
				break
			}
			imagen.duracion = uint32(intData)
			mediaContent.multimedia = append(mediaContent.multimedia, imagen)
			break
		case "3":
			imagen := &Video{}
			fmt.Println("Ingrese el nombre del video")
			scanner.Scan()
			strData = scanner.Text()

			imagen.titulo = strData
			fmt.Println("Ingrese el formato del video")

			scanner.Scan()
			strData = scanner.Text()
			imagen.formato = strData

			fmt.Println("Ingrese los frames del video")
			scanner.Scan()
			strData = scanner.Text()

			intData, err := strconv.ParseUint(strData, 10, 32)
			if err != nil {
				fmt.Println("Ingrese un numero valido")
				break
			}
			imagen.frames = uint32(intData)

			mediaContent.multimedia = append(mediaContent.multimedia, imagen)
			break
		case "4":
			fmt.Println("Titulo Formato canales/duracion/frames")
			for _, data := range mediaContent.multimedia {
				data.mostrar()
			}
			break
		case "5":
			fmt.Println("Adios")
			break
		default:
			fmt.Println("Ingrese una opcion valida")
		}
		fmt.Println("Presione enter para continuar...")
		scanner.Scan()
	}
}
