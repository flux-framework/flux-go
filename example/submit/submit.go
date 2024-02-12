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
	"github.com/flux-framework/flux-core-go/pkg/core"
)

func main() {
	fmt.Println("⭐️ Testing flux submit in Go! ⭐️")

	flux := core.NewFlux()

	// Handle is at flux.Handle
	fmt.Printf("Submitting a Sleep Job: sleep 10\n")

	// Create and submit a jobspec
	jobspec := core.NewJobSpec("sleep 10")
	future := flux.Submit(jobspec)
	fmt.Printf("Flux Future: %s\n", future)

	C.flux_future_wait_all_create()

	// Get the id for it
	// ran into this issue https://github.com/golang/go/issues/13467
	// cannot use _cgo0 (variable of type *core._Ctype_struct_flux_future) as *_Ctype_struct_flux_future value in argument to _Cfunc_flux_job_submit_get_id
	// id := new(C.ulong)
	// fluxid := C.flux_job_submit_get_id(future, id)
	// fmt.Println("%s", fluxid)
	// flux_future_destroy (f);
}