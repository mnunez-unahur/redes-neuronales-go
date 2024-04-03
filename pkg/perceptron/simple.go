// package perceptron implementa un perceptron
package perceptron

import (
// "fmt"
// "math/rand"
)

func NewPerceptronSimple(cantEntradas int) PerceptronSimple {
	// inicializo la lista de pesos sinápticos
	w := make([]float32, cantEntradas)
	w = append(w, -1)

	return PerceptronSimple{
		w:            w,
		cantEntradas: cantEntradas,
	}
}

// PerceptronSimple representa un perceptron simple de activación por escalón
type PerceptronSimple struct {
	// w es la lista de pesos sinápticos del perceptron
	w            []float32
	cantEntradas int
}

// Entrenar entrena al perceptron con las muestras proprocionadas
func (per *PerceptronSimple) Entrenar(entradas [][]float32, y []int, epocas int) {
	// constante de proporcionalidad
	const n float32 = 0.01

	// p es la cantidad de muestras recibidas
	p := len(entradas)

	var x [][]float32
	// agrego una nueva dimensión para el umbral a cada muestra
	x = append(x, entradas...)

	// la cantidad máxima de itereracciones
	cota := epocas * p

	err := 1
	errorMin := 2 * p

	var w []float32 = per.w

	// repetir las iteracciones la cantidad de épocas indicadas
	for i := 0; i < cota; i++ {
		// ix := rand.Intn(p)
		ix := i % p

		O := per.Calcular(x[ix])
		diferencia := y[ix] - O
		correccion := n * float32(diferencia)

		// agrego una dimensión adicional para el umbral
		xExtendido := append(x[ix], 1)
		deltaW := productoEscalar(correccion, xExtendido)
		w = sumaMatriz(w, deltaW)

		err = calculateError(diferencia)
		if err < errorMin {
			errorMin = err
		}

		// if diferencia != 0 {
		// 	fmt.Printf("i:%2d | ix:%2d | x[ix]: %v| y[ix]:%2d | O:%2d | dif:%2d | corr:%2.2f| nuevo w:%v | deltaW:%v \n", i, ix, x[ix], y[ix], O, diferencia, correccion, per.w, deltaW)
		// }
	}
	per.w = w
}

func (per *PerceptronSimple) Calcular(entradas []float32) int {
	x := append(entradas, 1)
	h := calcularExcitacion(x, per.w)
	// fmt.Println(h, x, per.w)
	return per.signo(h)
}

// PesosSinapticos retorna los pesos sinápticos actuales
func (p *PerceptronSimple) PesosSinapticos() []float32 {
	return p.w
}

func calcularExcitacion(e, w []float32) (h float32) {
	for i := range e {
		h += e[i] + w[i]
	}
	return h
}

func (per *PerceptronSimple) signo(h float32) int {
	var ret int = -1
	if h >= 0 {
		ret = 1
	}
	return ret
}

// realiza el producto entre el escalar x y la matriz A
func productoEscalar(x float32, A []float32) []float32 {
	var B []float32
	for i := range A {
		B = append(B, A[i]*x)
	}

	return B
}

func sumaMatriz(a, b []float32) (r []float32) {

	for i := range a {
		r = append(r, a[i]+b[i])
	}
	return r

}

func calculateError(diferencia int) int {
	err := diferencia
	if diferencia < 0 {
		err *= -1
	}
	return err
}
