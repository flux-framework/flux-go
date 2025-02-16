package flux

import (
	"fmt"
	"os"
	"testing"
)

// PathExists checks if the given path exists.
func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func TestJobs(t *testing.T) {

	flux := NewFlux()
	jobspec := NewJobSpec("sleep 10")
	flux.Submit(jobspec)

	jobs, err := flux.ListJobs()
	if err != nil {
		t.Log("Error listing flux jobs")
	}
	t.Log(jobs)
}

func TestSubmit(t *testing.T) {

	flux := NewFlux()
	jobspec := NewJobSpec("sleep 10")
	future := flux.Submit(jobspec)
	fmt.Printf("Flux Future: %s\n", future)
}

func TestKeygen(t *testing.T) {

	t.Log("Testing generation to file")

	// A name for the certificate - this is often the hostname
	fmt.Println("Generating to path...")
	hostname, err := getHostname()
	if err != nil {
		t.Errorf("hostname: %s", err)
	}
	name := "curve-cert"

	// Example to save to path
	path := "/tmp/curve.cert"
	KeyGen(name, hostname, path)
	if !pathExists(path) {
		t.Errorf("Path %s does not exist\n", path)
	}

	t.Log("Testing generation to string")
	curveCert, err := KeyGen(name, hostname, "")
	if err != nil {
		t.Errorf("keygen generation to string: %s", err)
	}
	fmt.Println(curveCert)
}
