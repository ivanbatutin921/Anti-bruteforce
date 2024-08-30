package main

import (
	// sr "root/internal/services"
	"root/internal/app"
)

func main() {
	app.Run()

	// tb := sr.NewTokenbucket(3,1)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	ip := r.RemoteAddr
	// 	if !tb.Take(ip, 1) {
	// 		http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
	// 		return
	// 	}
	// 	fmt.Fprint(w, "Hello, World!")
	// })

	// http.ListenAndServe(":8080", nil)
}
