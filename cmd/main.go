package main

import (
	"github.com/BorislavTimoshin/YandexAcademyOfGolang/pkg/calculator"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/calculate", CalculatorHandler)
	http.ListenAndServe(":8080", nil)
}
