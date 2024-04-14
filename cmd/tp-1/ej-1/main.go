package main

import (
	// "fmt"
	"fmt"

	perceptron "github.com/mnunez-unahur/redes-neuronales-go/pkg/perceptron"
)

func main() {
	fmt.Println("calculando Función Lógica 'AND'")
	muestras := [][]float64{
		{1, 1}, {-1, 1}, {1, -1}, {-1, -1},
	}
	valoresEsperados := []float64{1, -1, -1, -1}
	w, i, err := perceptron.Entrenar(muestras, valoresEsperados, len(muestras)*10, perceptron.ActivadorLineal)
	if err != nil {
		fmt.Println(err.Error())
	}
	salidas := perceptron.Clasificar(muestras, w, perceptron.ActivadorLineal)
	fmt.Printf("w: %v - i:%v\n", w, i)
	for i := range muestras {
		fmt.Printf("\t%v\t%2.0f\t%2.0f\n", muestras[i], valoresEsperados[i], salidas[i])
	}
	fmt.Println("")

	// fmt.Println("calculando Función Lógica 'OR'")
	// muestras = [][]float64{
	// 	{1, 1}, {-1, 1}, {1, -1}, {-1, -1},
	// }
	// valoresEsperados = []float64{1, 1, 1, -1}
	// w, i, err = perceptron.Entrenar(muestras, valoresEsperados, len(muestras)*10, perceptron.ActivadorLineal)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// salidas = perceptron.Clasificar(muestras, w, perceptron.ActivadorLineal)
	// fmt.Printf("w: %v - i:%v\n", w, i)
	// for i := range muestras {
	// 	fmt.Printf("\t%v\t%2.0f\t%2.0f\n", muestras[i], valoresEsperados[i], salidas[i])
	// }
	// fmt.Println("")

	// fmt.Println("calculando Función Lógica 'Implica'")
	// muestras = [][]float64{
	// 	{1, 1}, {-1, 1}, {1, -1}, {-1, -1},
	// }
	// valoresEsperados = []float64{1, 1, -1, 1}
	// w, i, err = perceptron.Entrenar(muestras, valoresEsperados, len(muestras)*10, perceptron.ActivadorLineal)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// salidas = perceptron.Clasificar(muestras, w, perceptron.ActivadorLineal)
	// fmt.Printf("w: %v - i:%v\n", w, i)
	// for i := range muestras {
	// 	fmt.Printf("\t%v\t%2.0f\t%2.0f\n", muestras[i], valoresEsperados[i], salidas[i])
	// }
	// fmt.Println("")

	// fmt.Println("calculando Función Lógica 'XOR'")
	// muestras = [][]float64{
	// 	{1, 1}, {-1, 1}, {1, -1}, {-1, -1},
	// }
	// valoresEsperados = []float64{-1, 1, 1, -1}
	// w, i, err = perceptron.Entrenar(muestras, valoresEsperados, len(muestras)*10, perceptron.ActivadorLineal)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// salidas = perceptron.Clasificar(muestras, w, perceptron.ActivadorLineal)
	// fmt.Printf("w: %v - i:%v\n", w, i)
	// for i := range muestras {
	// 	fmt.Printf("\t%v\t%2.0f\t%2.0f\n", muestras[i], valoresEsperados[i], salidas[i])
	// }
	// fmt.Println("")

	// muestras = [][]float64{
	// 	// E1    E2
	// 	{1.1946, 3.8427},
	// 	{0.8788, 1.1695},
	// 	{1.1907, 1.6117},
	// 	{1.1480, 3.8272},
	// 	{0.2032, 1.9208},
	// 	{2.7571, 1.0931},
	// 	{4.7125, 2.8166},
	// 	{3.9392, 1.1032},
	// 	{1.2072, 0.8132},
	// 	{3.4799, 1.9982},
	// 	{0.4763, 0.1020},
	// }

	// valoresEsperados = []float64{1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1}

	// w, i, err = perceptron.Entrenar(muestras, valoresEsperados, len(muestras)*100, perceptron.ActivadorLineal)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// salidas = perceptron.Clasificar(muestras, w, perceptron.ActivadorLineal)
	// fmt.Printf("w: %v - i:%v\n", w, i)
	// for i := range muestras {
	// 	fmt.Printf("\t%v\t%2.0f\t%2.0f\n", muestras[i], valoresEsperados[i], salidas[i])
	// }
	// fmt.Println("")

}
