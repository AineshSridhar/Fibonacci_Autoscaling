package handlers

import (
	"net/http"
	"strconv"
)

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}

func FibHandler(w http.ResponseWriter, r *http.Request) {
	nStr := r.URL.Query().Get("n")
	if nStr == "" {
		nStr = "40"
	}

	n, err := strconv.Atoi(nStr)
	if err != nil || n < 1 || n > 45 {
		http.Error(w, "Invalid n", http.StatusBadRequest)
		return
	}

	result := fib(n)
	w.Write([]byte(strconv.Itoa(result)))
}
