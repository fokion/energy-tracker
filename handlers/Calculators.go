package handlers

type Calculator interface {
	Calculate(from float64, to float64, price float64) float64
}
