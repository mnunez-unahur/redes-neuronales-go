package perceptron

// Muestra representa una muestra
type Muestra struct {
	Entradas       []float32
	SalidaEsperada float32
}

func NewSetPruebas(cantidadEntradas int) SetPruebas {
	return SetPruebas{}
}

// SetPruebas representa un set de pruebas completo
type SetPruebas struct {
	Muetras []Muestra
}

func (s *SetPruebas) AgregarMuestra(salidaEsperada float32, entradas ...float32) {
	s.Muetras = append(s.Muetras, Muestra{
		Entradas:       entradas,
		SalidaEsperada: salidaEsperada,
	})
}
