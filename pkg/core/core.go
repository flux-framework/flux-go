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
	"unsafe"
)

// Submit a job to the system.
func (f *Flux) Submit(jobspec *JobSpec) *C.flux_future_t {
	flag := C.int(16)
	spec := jobspec.Encoded()
	defer C.free(unsafe.Pointer(spec))
	return C.flux_job_submit(f.Handle, spec, flag, C.FLUX_JOB_WAITABLE)
}