package main

import (
	"log"
)

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLoggingService(svc)

	api := NewApiServer(svc)
	log.Fatal(api.Start(":3000"))
}