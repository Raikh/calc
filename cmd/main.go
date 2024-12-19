package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"internal/calc"
	"log"
	"net/http"
)

type ExpressionRequest struct {
	Expression string `json:"expression"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/calculate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			panic(errors.New("wrong method"))
		}
		decoder := json.NewDecoder(r.Body)
		var expression ExpressionRequest
		err := decoder.Decode(&expression)
		if err != nil {
			panic(err)
		}

		result, err := calc.Calc(expression.Expression)
		if err != nil {
			log.Printf("Error: %v", err)
			http.Error(w, `{"error": "Expression is not valid"}`, http.StatusUnprocessableEntity)
			return
		}

		fmt.Fprintf(w, `{"result": %f}`, result)
	})
	handler := PanicMiddleware(mux)
	log.Fatal(http.ListenAndServe(":80", handler))
}

func PanicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic: %v", err)
				http.Error(w, `{"error": "Internal server error"}`, http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
