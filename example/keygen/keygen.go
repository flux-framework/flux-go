package main

/*
#include <flux/core.h>
#include <flux/idset.h>
#include <flux/hostlist.h>
#include <stddef.h>
#include <jansson.h>
#include <stdlib.h>
*/
import "C"
import (
	"fmt"
	"log"
	"os"

	"github.com/flux-framework/flux-go/pkg/flux"
)

// getHostname gets the hostname
func getHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return hostname
}

func main() {
	fmt.Println("⭐️ Testing flux keygen in Go! ⭐️")

	// A name for the certificate - this is often the hostname
	fmt.Println("Generating to path...")
	hostname := getHostname()
	name := "curve-cert"

	// Example to save to path
	path := "./curve.cert"
	flux.KeyGen(name, hostname, path)

	// Example to return as string
	// if you don't provide a hostname, will default to this host
	fmt.Println("Generating to string...")
	curveCert, err := flux.KeyGen(name, hostname, "")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(curveCert)
}
