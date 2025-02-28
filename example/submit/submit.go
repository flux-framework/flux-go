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

	core "github.com/flux-framework/flux-go/pkg/flux"
)

func main() {
	fmt.Println("⭐️ Testing flux submit in Go! ⭐️")

	flux := core.NewFlux()

	// Handle is at flux.Handle
	fmt.Printf("Submitting a Sleep Job: sleep 10\n")

	// Create and submit a jobspec via a submit request
	// TODO need to better formalize this, this is a command.
	//	core.NewJobSpec("sleep 10")
	jobspec := core.NewJobSpec("sleep 10")
	future := flux.Submit(jobspec)
	fmt.Printf("Flux Future: %s\n", future)

	// Get the id for it
	// ran into this issue https://github.com/golang/go/issues/13467
	// cannot use _cgo0 (variable of type *core._Ctype_struct_flux_future) as *_Ctype_struct_flux_future value in argument to _Cfunc_flux_job_submit_get_id
	// id := new(C.ulong)
	// fluxid := C.flux_job_submit_get_id(future, id)
	// fmt.Println("%s", fluxid)
	// flux_future_destroy (f);

	fmt.Printf("Submitting a job from a Jobspec")
	specJson := `{"resources": [{"type": "slot", "count": 1, "with": [{"type": "core", "count": 1}], "label": "task"}], "tasks": [{"command": ["hostname"], "slot": "task", "count": {"per_slot": 1}}], "attributes": {"system": {"duration": 0}}, "version": 1}`
	future = flux.SubmitJobspec(specJson)
	fmt.Printf("Flux Future: %s\n", future)
}
