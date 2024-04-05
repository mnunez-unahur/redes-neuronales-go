package main

import (
	// "fmt"
	"fmt"

	"github.com/mnunez-unahur/redes-neuronales-go/pkg/perceptron"
)

func main() {
	// muestras := [][]float32{
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

	// valoresEsperados := []int{1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1}

	// fmt.Println("calculando Ejemplo")
	// pEjemplo := perceptron.NewPerceptronSimple(2)
	// pEjemplo.Entrenar(muestras, valoresEsperados, 10, 1)
	// fmt.Printf("Epocas Entrenadas: %2d, Confiabilidad Máxima: %2.2f\n", pEjemplo.EpocasEntrenadas(), pEjemplo.Confiabilidad())

	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[0], valoresEsperados[0], pEjemplo.Calcular(muestras[0]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[1], valoresEsperados[1], pEjemplo.Calcular(muestras[1]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[2], valoresEsperados[2], pEjemplo.Calcular(muestras[2]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[3], valoresEsperados[3], pEjemplo.Calcular(muestras[3]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[4], valoresEsperados[4], pEjemplo.Calcular(muestras[4]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[5], valoresEsperados[5], pEjemplo.Calcular(muestras[5]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[6], valoresEsperados[6], pEjemplo.Calcular(muestras[6]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[7], valoresEsperados[7], pEjemplo.Calcular(muestras[7]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[8], valoresEsperados[8], pEjemplo.Calcular(muestras[8]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[9], valoresEsperados[9], pEjemplo.Calcular(muestras[9]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[10], valoresEsperados[10], pEjemplo.Calcular(muestras[10]))

	fmt.Println("calculando Función Lógica 'Y'")
	muestras := [][]float32{
		{1, 1}, {-1, 1}, {1, -1}, {-1, -1},
	}
	valoresEsperados := []int{1, -1, -1, -1}
	pAnd := perceptron.NewPerceptronSimple(2)
	pAnd.Entrenar(muestras, valoresEsperados, 1000, 1)
	fmt.Printf("Epocas Entrenadas: %2d, Confiabilidad Máxima: %2.2f\n", pAnd.EpocasEntrenadas(), pAnd.Confiabilidad())

	fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[0], valoresEsperados[0], pAnd.Calcular(muestras[0]))
	fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[1], valoresEsperados[1], pAnd.Calcular(muestras[1]))
	fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[2], valoresEsperados[2], pAnd.Calcular(muestras[2]))
	fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[3], valoresEsperados[3], pAnd.Calcular(muestras[3]))

	fmt.Println("calculando Función Lógica 'O'")
	valoresEsperados = []int{1, 1, 1, -1}
	pOr := perceptron.NewPerceptronSimple(2)
	pOr.Entrenar(muestras, valoresEsperados, 1000, 1)
	fmt.Printf("Epocas Entrenadas: %2d, Confiabilidad Máxima: %2.2f\n", pOr.EpocasEntrenadas(), pOr.Confiabilidad())

	fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[0], valoresEsperados[0], pOr.Calcular(muestras[0]))
	fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[1], valoresEsperados[1], pOr.Calcular(muestras[1]))
	fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[2], valoresEsperados[2], pOr.Calcular(muestras[2]))
	fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[3], valoresEsperados[3], pOr.Calcular(muestras[3]))

	// fmt.Println("calculando Función Lógica '=>'")
	// valoresEsperados = []int{1, -1, 1, 1}
	// pImplicacion := perceptron.NewPerceptronSimple(2)
	// pImplicacion.Entrenar(muestras, valoresEsperados, 100, 1)
	// fmt.Printf("Epocas Entrenadas: %2d, Confiabilidad Máxima: %2.2f\n", pImplicacion.EpocasEntrenadas(), pImplicacion.Confiabilidad())

	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[0], valoresEsperados[0], pImplicacion.Calcular(muestras[0]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[1], valoresEsperados[1], pImplicacion.Calcular(muestras[1]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[2], valoresEsperados[2], pImplicacion.Calcular(muestras[2]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[3], valoresEsperados[3], pImplicacion.Calcular(muestras[3]))

	fmt.Println("calculando Función Lógica 'XOR'")
	pXor := perceptron.NewPerceptronSimple(2)
	valoresEsperados = []int{1, 1, -1, -1}

	pXor.Entrenar(muestras, valoresEsperados, 100, 1)
	fmt.Printf("Epocas Entrenadas: %2d, Confiabilidad Máxima: %2.2f\n", pXor.EpocasEntrenadas(), pXor.Confiabilidad())

	fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[0], valoresEsperados[0], pXor.Calcular(muestras[0]))
	fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[1], valoresEsperados[1], pXor.Calcular(muestras[1]))
	fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[2], valoresEsperados[2], pXor.Calcular(muestras[2]))
	fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[3], valoresEsperados[3], pXor.Calcular(muestras[3]))

	// muestras = [][]float32{
	// 	{2, 3},
	// 	{2, 5},
	// 	{4, 4},
	// 	{4, 8},
	// 	{7, 6},
	// 	{7, 10},
	// 	{8, 4},
	// 	{9, 8},
	// 	{12, 6},
	// 	{6, 1},
	// 	{8, 1},
	// 	{9, 2},
	// 	{12, 1},
	// 	{14, 3},
	// 	{15, 2},
	// 	{18, 1},
	// 	{18, 4},
	// 	{21, 2},
	// 	{21, 4},
	// 	{21, 6},
	// 	{23, 4},
	// }

	// valoresEsperados = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}

	// fmt.Println("calculando Ejemplo")
	// pEjemplo := perceptron.NewPerceptronSimple(2)
	// pEjemplo.Entrenar(muestras, valoresEsperados, 100, 1)
	// fmt.Printf("Epocas Entrenadas: %2d, Confiabilidad Máxima: %2.2f\n", pEjemplo.EpocasEntrenadas(), pEjemplo.Confiabilidad())

	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[0], valoresEsperados[0], pEjemplo.Calcular(muestras[0]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[1], valoresEsperados[1], pEjemplo.Calcular(muestras[1]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[2], valoresEsperados[2], pEjemplo.Calcular(muestras[2]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[3], valoresEsperados[3], pEjemplo.Calcular(muestras[3]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[4], valoresEsperados[4], pEjemplo.Calcular(muestras[4]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[5], valoresEsperados[5], pEjemplo.Calcular(muestras[5]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[6], valoresEsperados[6], pEjemplo.Calcular(muestras[6]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[7], valoresEsperados[7], pEjemplo.Calcular(muestras[7]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[8], valoresEsperados[8], pEjemplo.Calcular(muestras[8]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[9], valoresEsperados[9], pEjemplo.Calcular(muestras[9]))
	// fmt.Printf("%v\tesperado: %2d\tresultado: %2d\n", muestras[10], valoresEsperados[10], pEjemplo.Calcular(muestras[10]))

}
