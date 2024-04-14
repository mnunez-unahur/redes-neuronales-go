package perceptron_v2

import (
	"errors"
	"fmt"
	"math"
)

type Signo func(float32) float32

func SignoPaso(h float32) float32 {
	var ret float32 = -1
	if h >= 0 {
		ret = 1
	}
	return ret
}

func SignoLineal(h float32) float32 {
	return h
}

func Entrenar(x [][]float32, y []float32, cota int, signo Signo) ([]float32, int, error) {
	p := len(x)
	const n float32 = 0.1

	if len(x) == 0 {
		return nil, 0, errors.New("no se proporcionaron valores de entrada")
	}
	if len(y) == 0 {
		return nil, 0, errors.New("no se proporcionaron valores de salida")
	}

	w := make([]float32, len(x[0]))
	w = append(w, -1)
	wMin := w

	var err float32 = 1.0
	var errMin float32 = float32(2 * p)

	i := 0
	for i < cota && err > 0 {
		err = 0
		for ix := 0; ix < p; ix++ {
			e := append(x[ix], 1)

			h := calcularExcitacion(e, w)
			O := signo(h)
			err += float32(math.Abs(float64(y[ix] - O)))
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

func Clasificar(entradas [][]float32, w []float32, signo Signo) []float32 {
	var salidas []float32

	for _, e := range entradas {
		x := append(e, 1)
		h := calcularExcitacion(x, w)
		O := signo(h)
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
