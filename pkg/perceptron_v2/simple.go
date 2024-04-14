package perceptron_v2

import (
	"errors"
	"fmt"
	"math"
)

type Activador func(float64) float64

func ActivadorPaso(h float64) float64 {
	var ret float64 = -1
	if h >= 0 {
		ret = 1
	}
	return ret
}

func ActivadorLineal(h float64) float64 {
	return h
}

func Entrenar(x [][]float64, y []float64, cota int, signo Activador) ([]float64, int, error) {
	p := len(x)
	const n float64 = 0.1

	if len(x) == 0 {
		return nil, 0, errors.New("no se proporcionaron valores de entrada")
	}
	if len(y) == 0 {
		return nil, 0, errors.New("no se proporcionaron valores de salida")
	}

	w := make([]float64, len(x[0]))
	w = append(w, -1)
	wMin := w

	var err float64 = 1.0
	var errMin float64 = float64(2 * p)

	i := 0
	for i < cota && err > 0 {
		err = 0
		for ix := 0; ix < p; ix++ {
			e := append(x[ix], 1)

			h := calcularExcitacion(e, w)
			O := signo(h)
			err += float64(math.Abs(float64(y[ix] - O)))
			deltaW := productoEscalar(n*(y[ix]-O), e)
			w = sumaMatriz(w, deltaW)
			fmt.Println(e, y[ix], h, O, deltaW, w)
			i++
		}
		if err < errMin {
			errMin = err
			wMin = w
		}

	}

	return wMin, i, nil
}

func Clasificar(entradas [][]float64, w []float64, signo Activador) []float64 {
	var salidas []float64

	for _, e := range entradas {
		x := append(e, 1)
		h := calcularExcitacion(x, w)
		O := signo(h)
		salidas = append(salidas, O)
	}

	return salidas
}

func calcularExcitacion(e, w []float64) (h float64) {
	for i := range w {
		h += e[i] * w[i]
	}
	return h
}

// realiza el producto entre el escalar x y la matriz A
func productoEscalar(x float64, A []float64) []float64 {
	var B []float64
	for i := range A {
		B = append(B, A[i]*x)
	}

	return B
}

func sumaMatriz(a, b []float64) (r []float64) {

	for i := range a {
		r = append(r, a[i]+b[i])
	}
	return r

}
