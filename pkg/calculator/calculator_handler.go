package calculator

import (
	"encoding/json"
	"log"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result float64 `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func CalculatorHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	result, err := Calc(req.Expression)
	if err != nil {
		if err.Error() == "Недопустимое выражение" || err.Error() == "Недопустимый символ" ||
			err.Error() == "Несоответствующие скобки" || err.Error() == "Деление на ноль" ||
			err.Error() == "Неизвестный оператор" {
			http.Error(w, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity)
		} else {
			log.Println("Internal server error:", err)
			http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
		}
		return
	}

	response := Response{Result: result}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
