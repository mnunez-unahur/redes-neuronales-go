package perceptron

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

type Activador func(float64) float64

func ActivadorPaso(h float64) float64 {
	if h >= 0 {
		return 1
	}
	return -1
}

func ActivadorLineal(h float64) float64 {
	return h
}

func Entrenar(x [][]float64, y []float64, cota int, activador Activador) ([]float64, int, error) {
	p := len(x)
	const n float64 = 0.001

	if len(x) == 0 {
		return nil, 0, errors.New("no se proporcionaron valores de entrada")
	}
	if len(y) == 0 {
		return nil, 0, errors.New("no se proporcionaron valores de salida")
	}

	w := make([]float64, len(x[0]))
	w = append(w, 1)
	wMin := w

	var err float64 = 1
	var errMin float64 = 2 * float64(p)

	i := 0
	fmt.Println("************Entrenamiento*********")
	fmt.Printf("%10v %2v %5v %5v %5v %22v %22v %v\n",
		"e    ",
		"y ",
		"h ",
		"O ",
		"dif ",
		"dW          ",
		"w           ",
		"err ",
	)

	for i = 0; i < cota && err > 0; i++ {
		ix := rand.Intn(p)
		// ix := i % p
		e := append(x[ix], 1)

		h := calcularExcitacion(e, w)
		O := activador(h)
		dif := n * (y[ix] - O)
		deltaW := productoEscalar(dif, e)
		w = sumarVectores(w, deltaW)
		err = calcularError(x, y, w, activador)
		fmt.Printf("%7.4f %7.4f %5.1f %5.1f %5.1f %6.2f %6.2f %v\n",
			e,
			y[ix],
			h,
			O,
			dif,
			deltaW,
			w,
			err,
		)
		if err < errMin {
			errMin = err
			wMin = w
		}
	}

	return wMin, i, nil
}

func Clasificar(entradas [][]float64, w []float64, activador Activador) []float64 {
	var salidas []float64

	for _, e := range entradas {
		x := append(e, 1)
		h := calcularExcitacion(x, w)
		O := activador(h)
		salidas = append(salidas, O)
	}

	return salidas
}

func calcularError(entradas [][]float64, salidasEsperadas []float64, w []float64, activador Activador) float64 {
	var err float64 = 0
	// fmt.Println("calcularError")
	for i, e := range entradas {
		x := append(e, 1)
		h := calcularExcitacion(x, w)
		O := activador(h)
		err += math.Abs(salidasEsperadas[i] - O)
		// fmt.Printf("%6.2f %6.2f %6.2f %v\n",
		// 	salidasEsperadas[i],
		// 	O,
		// 	salidasEsperadas[i]-O,
		// 	err,
		// )
	}
	return err
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
