package main

import (
	"context"
	"fmt"
	"log"
)

func main() {
	svc := NewCatFactService("https://catfact.ninja/fact")
	svc = NewLoggingService(svc)

	fact, err := svc.GetCatFact(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("%v+\n", fact.Fact)
}