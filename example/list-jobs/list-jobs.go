package main

import (
	"fmt"
	"log"

	core "github.com/flux-framework/flux-go/pkg/flux"
)

func main() {
	fmt.Println("⭐️ Testing flux list jobs rpc in Go! ⭐️")
	flux := core.NewFlux()
	listing, err := flux.ListJobs()
	if err != nil {
		log.Fatalf("womp womp")
	}
	fmt.Println(listing)
}
