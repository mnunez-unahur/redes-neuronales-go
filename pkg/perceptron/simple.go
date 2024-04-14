package perceptron

import (
	"math"
	"math/rand"
)

func Entrenar(x [][]float32, y []float32, cota int) ([]float32, int) {
	p := len(x)
	const n float32 = 0.1

	w := make([]float32, len(x[0]))
	w = append(w, -1)
	wMin := w

	err := 1
	errMin := 2 * p

	i := 0
	for i = 0; i < cota && err > 0; i++ {
		ix := rand.Intn(p)
		// ix := i % p
		e := append(x[ix], 1)

		h := calcularExcitacion(e, w)
		O := signo(h)
		deltaW := productoEscalar(n*(float32(y[ix])-float32(O)), e)
		w = sumaMatriz(w, deltaW)
		err = calcularError(x, y, w)
		if err < errMin {
			errMin = err
			wMin = w
		}

	}

	return wMin, i
}

func Clasificar(e [][]float32, w []float32) []float32 {
	p := len(e)

	var salidas []float32

	for i := 0; i < p; i++ {
		x := append(e[i], 1)
		h := calcularExcitacion(x, w)
		O := float32(signo(h))
		salidas = append(salidas, O)
	}

	return salidas
}

func calcularExcitacion(e, w []float32) (h float32) {
	for i := range w {
		h += e[i] * w[i]
	}
	return h
}

func signo(h float32) int {
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

func calcularError(entradas [][]float32, salidasEsperadas []float32, w []float32) int {
	err := 0
	p := len(entradas)
	for i := 0; i < p; i++ {
		x := append(entradas[i], 1)
		h := calcularExcitacion(x, w)
		O := signo(h)
		err += int(math.Abs(float64(salidasEsperadas[i]) - float64(O)))
	}
	return err
}
