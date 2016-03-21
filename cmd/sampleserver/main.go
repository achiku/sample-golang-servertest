package main

import (
	"log"

	"github.com/achiku/sample-golang-servertest"
)

func main() {
	s := sampleserver.CreateServer()
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
