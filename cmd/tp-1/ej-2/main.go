package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	perceptron "github.com/mnunez-unahur/redes-neuronales-go/pkg/perceptron"
)

func main() {
	fmt.Println("Cargando Entradas")
	entradas, err := cargarEntradas("./cmd/tp-1/ej-2/entradas.txt")
	if err != nil {
		fmt.Printf("Error al cargar entradas: %s\n", err)
		return
	}

	fmt.Println("Cargando Salidas Deseadas")
	salidasEsperadas, err := cargarSalidas("./cmd/tp-1/ej-2/salidas.txt")
	if err != nil {
		fmt.Printf("Error al cargar lista valores deseados: %s\n", err)
		return
	}

	setEntrenamientoEntradas := entradas[0:150]
	setEntrenamientoSalidas := salidasEsperadas[0:150]

	w, i, err := perceptron.Entrenar(setEntrenamientoEntradas, setEntrenamientoSalidas, len(entradas)*100, perceptron.SignoLineal)
	if err != nil {
		fmt.Printf("Error al entrenar: %v\n", err)
		return
	}
	fmt.Printf("w: %v - i:%v\n", w, i)

	setPruebasEntradas := entradas[150:]
	setPruebasSalidas := salidasEsperadas[150:]

	salidas := perceptron.Clasificar(setPruebasEntradas, w, perceptron.SignoLineal)
	for i := range setPruebasEntradas {
		fmt.Printf("\t%d\t%8v\t%2.4f\t%2.4f\n", i+1, setPruebasEntradas[i], setPruebasSalidas[i], salidas[i])
	}
	fmt.Println("")
}

func cargarEntradas(filename string) (lista [][]float64, err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		partes := strings.Fields(scanner.Text())
		var numeros []float64
		for _, p := range partes {
			n, err := strconv.ParseFloat(p, 32)
			if err != nil {
				n = 0
			}
			numeros = append(numeros, float64(n))
		}

		lista = append(lista, numeros)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lista, nil
}

func cargarSalidas(filename string) (lista []float64, err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		partes := strings.Fields(scanner.Text())
		n, err := strconv.ParseFloat(partes[0], 32)
		if err != nil {
			n = 0
			continue
		}
		lista = append(lista, float64(n))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lista, nil
}
