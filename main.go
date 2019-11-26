package main

import (
	"fmt"
	"laki/controllers"
	"net/http"
)

func main() {
	router := controllers.InitRouter()

	s := &http.Server{
		Addr:    fmt.Sprintf(":8080"),
		Handler: router,
		// ReadTimeout:    setting.ReadTimeout,
		// WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
