package perceptron

import (
	"errors"
	"fmt"
	"math"
	// "math/rand"
)

type Signo func(float64) float64

func SignoPaso(h float64) float64 {
	if h >= 0 {
		return 1
	}
	return -1
}

func SignoLineal(h float64) float64 {
	return h
}

func Entrenar(x [][]float64, y []float64, cota int, signo Signo) ([]float64, int, error) {
	p := len(x)
	const n float64 = 0.1

	if len(x) == 0 {
		return nil, 0, errors.New("no se proporcionaron valores de entrada")
	}
	if len(y) == 0 {
		return nil, 0, errors.New("no se proporcionaron valores de salida")
	}

	w := make([]float64, len(x[0]))
	w = append(w, 1)
	wMin := w

	var err uint64 = 1
	var errMin uint64 = 2 * uint64(p)

	fmt.Println(w, err)
	i := 0
	for i = 0; i < cota && err > 0; i++ {
		// ix := rand.Intn(p)
		ix := i % p
		e := append(x[ix], 1)

		h := calcularExcitacion(e, w)
		O := signo(h)
		dif := n * (y[ix] - O)
		deltaW := productoEscalar(dif, e)
		w = sumarVectores(w, deltaW)
		err = calcularError(x, y, w, signo)
		fmt.Println(e, y[ix], h, dif, deltaW, w, err)
		if err < errMin {
			errMin = err
			wMin = w
		}
	}

	return wMin, i, nil
}

func Clasificar(entradas [][]float64, w []float64, signo Signo) []float64 {
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

func sumarVectores(a, b []float64) (r []float64) {

	for i := range a {
		r = append(r, a[i]+b[i])
	}
	return r

}

func calcularError(entradas [][]float64, salidasEsperadas []float64, w []float64, signo Signo) uint64 {
	var err uint64 = 0
	for i, e := range entradas {
		x := append(e, 1)
		h := calcularExcitacion(x, w)
		O := signo(h)
		err += uint64(math.Abs(salidasEsperadas[i] - O))
	}
	return err
}
