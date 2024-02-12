package core

/*
#include <flux/core.h>
#include <flux/idset.h>
#include <flux/hostlist.h>
#include <stddef.h>
#include <jansson.h>
#include <stdlib.h>
#include "../flux/cgo_helpers.h"
*/
import "C"
import (
	"fmt"
	"os"
)

// Create a new Flux Handle
func NewFlux() Flux {

	// Get any FLUX_URI in the environment
	flux_uri := os.Getenv("FLUX_URI")
	uri := C.CString(flux_uri)
	flags := C.int(0)

	// Create the handle
	handle := C.flux_open(uri, flags)

	// I don't know how to catch this if fails
	if handle == nil {
		var err C.flux_error_t
		handle = C.flux_open_ex(uri, flags, &err)
	}

	// if it's still nil, nogo
	if handle == nil {
		fmt.Println("Cannot create Flux handle. Is Flux running?")
		os.Exit(1)
	}

    // TODO thread local storage to hold a reactor_running boolean, indicating
    // when the current thread is running under a Flux reactor.
    // should be FluxHandle.tls or similar
    // tls = threading.local()
    return Flux{Handle: handle}
}

type Flux struct {
	Handle *C.flux_t
    ReactorDepth int
    Exception error

    // I don't know what this is
    auxTxn int
    // TODO this should be a worker class?
    ActiveWorkers map[string]string
}

// Return True if this thread is running the Flux reactor
func (f *Flux) reactorRunning() bool {
    return f.ReactorDepth >= 0
}

func (f *Flux) reactorEnter() {
    f.ReactorDepth += 1
}

func (f *Flux) reactorExit() {
    f.ReactorDepth -= 1
}

// TODO inReactor should instead just be adding then removing a count from the reactor (enter and exit)
func (f *Flux) inReactor() {
    f.reactorEnter()
}

func (f *Flux) setException(exception error) error {
    prev := f.Exception
    f.Exception = exception
    // Set exception on thread?
    //cls.tls.exception = exception
    return prev
}
