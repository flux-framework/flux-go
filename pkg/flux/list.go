package flux

/*
#include <flux/core.h>
#include <flux/idset.h>
#include <flux/hostlist.h>
#include <stddef.h>
#include <jansson.h>
#include <stdlib.h>

// listjobs lists flux jobs under the handle.
//   there is no limit to number (max entries is 0)
//   an empty string result indicates an error
char * listjobs (flux_t *h) {
    flux_future_t *f;
    json_t *jobs;
	char *job_string;
	int max_entries = 0;

    if (!(f = flux_rpc_pack (h,
                             "job-manager.list",
                             FLUX_NODEID_ANY,
                             0,
                             "{s:i}",
                             "max_entries",
                             max_entries)))
        return job_string;

    if (flux_rpc_get_unpack (f, "{s:o}", "jobs", &jobs) < 0)
        return job_string;

	job_string = json_dumps(jobs, JSON_INDENT(1));
    flux_future_destroy (f);
    return job_string;
}
*/
import "C"
import (
	"fmt"
)

// ListJob will list jobs (print to the console) and return the string
func (f *Flux) ListJobs() (string, error) {
	result := C.listjobs((*C.struct_flux_handle)(f.Handle))
	listing := C.GoString(result)

	// An empty result indicates an error
	if listing == "" {
		return "", fmt.Errorf("unable to list jobs")
	}
	return listing, nil
}
