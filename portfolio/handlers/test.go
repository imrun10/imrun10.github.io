package handlers

import (
	"fmt"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Println("test")
	fmt.Println("✅ Handler called!")  // More visible log
	w.Write([]byte("Mouse entered!")) // Simple response
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	// hello("Imran").Render(r.Context(), w)
	// header("My Portfolio").Render(r.Context(), w)
}
