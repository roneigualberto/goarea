package goarea

import "math"

// Pi é uma constante PI
const Pi = 3.1416

// Circu Área do Circulo
func Circu(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect é
func Rect(base, altura float64) float64 {
	return base * altura
}

func _Triangulo(base, altura float64) float64 {
	return (base * altura) / 2
}
