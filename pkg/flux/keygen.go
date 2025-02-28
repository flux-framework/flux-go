package flux

/*
#include <unistd.h>
#include <flux/core.h>
#include <czmq.h>
#include <sodium.h>
#include <stddef.h>
#include <stdlib.h>
#include "../flux/cgo_helpers.h"

void flux_zcert_set_meta (zcert_t *cert, const char *field, const char *name) {
    zcert_set_meta (cert, field, "%s", name);
}
*/
import "C"
import (
	"fmt"
	"os"
	"unsafe"
)

// Keygen uses zeromq, and was originally part of the flux/cmd directory
// so we put it under flux implying that, since "cmd" has special meaning
// in Go!

var (
	template = `#   ****  Generated on 2023-04-26 22:54:42 by CZMQ  ****
#   ZeroMQ CURVE **Secret** Certificate
#   DO NOT PROVIDE THIS FILE TO OTHER USERS nor change its permissions.

metadata
    name = "%s"
    keygen.hostname = "%s"
curve
    public-key = "%s"
    secret-key = "%s"`
)

// getHostname gets the hostname
func getHostname() (string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return "", err
	}
	return hostname, nil
}

// KeyGen generates a curve certificate
// This does not require a flux handle
func KeyGen(name string, hostname string, path string) (string, error) {

	// Create the new certificate (likely want to check for error here)
	cert := C.zcert_new()

	// If no hostname provided, assume this host
	var err error
	if hostname == "" {
		hostname, err = getHostname()
		if err != nil {
			return "", err
		}
	}
	// Use wrapper to set the cert metadata
	// Name (typically the hostname but doesn't need to be)
	// used in overlay logging
	nameValue := C.CString(name)
	nameField := C.CString("name")
	defer C.free(unsafe.Pointer(nameValue))
	defer C.free(unsafe.Pointer(nameField))
	C.flux_zcert_set_meta(cert, nameField, nameValue)

	// Hostname
	hostnameValue := C.CString(hostname)
	hostnameField := C.CString("keygen.hostname")
	defer C.free(unsafe.Pointer(hostnameValue))
	defer C.free(unsafe.Pointer(hostnameField))
	C.flux_zcert_set_meta(cert, hostnameField, hostnameValue)

	// If we don't have a path, save to string
	var curveCert string
	if path == "" {

		publicKey := C.zcert_public_txt(cert)
		secretKey := C.zcert_secret_txt(cert)
		public := C.GoString((*C.char)(unsafe.Pointer(publicKey)))
		secret := C.GoString((*C.char)(unsafe.Pointer(secretKey)))
		curveCert = fmt.Sprintf(template, name, hostname, public, secret)

	} else {
		// Note that we can also generate keygen.time, keygen.userid,
		// And other version metadata. See
		cpath := C.CString(path)
		defer C.free(unsafe.Pointer(cpath))
		fmt.Printf("Saving to %s\n", path)
		C.zcert_save_secret(cert, cpath)
	}

	C.zcert_destroy(&cert)
	return curveCert, nil
}
