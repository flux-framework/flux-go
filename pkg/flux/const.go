package flux

/*
#include <flux/idset.h>
#include <flux/hostlist.h>
#include <stddef.h>
#include <jansson.h>
#include <stdint.h>
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// IoPrioDeltaMax as defined in bits/local_lim.h:78
	IoPrioDeltaMax = 20
	// Ccessperms as defined in sys/stat.h:195
	Ccessperms = 0x5ff140
	// Llperms as defined in sys/stat.h:196
	Llperms = 0x5ff140
)
