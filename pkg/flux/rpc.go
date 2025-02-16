package flux

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

// Send an RPC message. We assume sending to lead broker
func (f *Flux) RPC(topic, payload string, nodeid int32) *C.flux_future_t {
	t := C.CString(topic)
	s := C.CString(payload)
	// The last argument is flags, we can assume none for now
	return C.flux_rpc(f.Handle, t, s, C.uint(nodeid), 0)
}
