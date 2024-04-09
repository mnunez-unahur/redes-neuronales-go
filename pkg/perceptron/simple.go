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
	w                []float32
	cantEntradas     int
	confiabilidad    float32
	epocasEntranadas int
}

// Entrenar entrena al perceptron con las muestras proprocionadas
func (per *PerceptronSimple) Entrenar(entradas [][]float32, y []int, maxEpocas int, minConfiabilidad int) {
	// constante de proporcionalidad
	const n float32 = 0.01

	// p es la cantidad de muestras recibidas
	p := len(entradas)

	var x [][]float32
	// agrego una nueva dimensión para el umbral a cada muestra
	x = append(x, entradas...)

	var w []float32 = per.w

	// repetir las iteracciones la cantidad de épocas indicadas
	for e := 0; e < maxEpocas; e++ {
		cantOk := 0

		for i := 0; i < p; i++ {
			// ix := rand.Intn(p)
			ix := i % p

			O := per.CalcularSalida(x[ix])
			diferencia := y[ix] - O
			correccion := n * float32(diferencia)
			if diferencia == 0 {
				cantOk++
			}

			// agrego una dimensión adicional para el umbral
			xExtendido := append(x[ix], 1)
			deltaW := productoEscalar(correccion, xExtendido)
			w = sumaMatriz(w, deltaW)
			per.w = w
		}
		per.epocasEntranadas++
		confiabilidadEpoca := float32(cantOk) / float32(p)

		if confiabilidadEpoca > per.confiabilidad {
			per.confiabilidad = confiabilidadEpoca
			// actualizo el peso sináptico
		}

		if per.confiabilidad >= float32(minConfiabilidad) {
			break
		}

	}
}

func (per *PerceptronSimple) CalcularSalida(entradas []float32) int {
	x := append(entradas, 1)
	h := calcularExcitacion(x, per.w)
	return per.signo(h)
}

// PesosSinapticos retorna los pesos sinápticos actuales
func (p *PerceptronSimple) PesosSinapticos() []float32 {
	return p.w
}

// PesosSinapticos retorna los pesos sinápticos actuales
func (p *PerceptronSimple) EpocasEntrenadas() int {
	return p.epocasEntranadas
}

// PesosSinapticos retorna los pesos sinápticos actuales
func (p *PerceptronSimple) Confiabilidad() float32 {
	return p.confiabilidad
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
