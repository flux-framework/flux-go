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
import (
	"encoding/json"
	"os"
	"strings"
	"unsafe"
)

// A jobspec defines a command and resources for it
// These are exposed for the creator to easily set
type JobSpec struct {
	Command []string `json:"command"`

	// I'm not sure how this is a list of strings
	Env          []string `json:"env"`
	Tasks        int      `json:"tasks"`
	CoresPerTask int      `json:"coresPerTask"`
	GpusPerTask  int      `json:"gpusPerTask"`
	Nodes        int      `json:"nodes"`
	Duration     float64  `json:"duration"`
}

// Create a new JobSpec and set defaults (only require a command)
// This is not technically a jobspec, it is intended for use with jobspec from command
func NewJobSpec(command string) *JobSpec {

	// Duration of 0 defaults to no limit
	jobspec := JobSpec{
		Tasks:        1,
		CoresPerTask: 1,
		GpusPerTask:  0,
		Nodes:        0,   // unset Nodes lets the scheduler choose
		Duration:     0.0, // no duration set!
	}
	jobspec.SetCommand(command)
	return &jobspec
}

// FromJson loads a jobspec from a json string
func JobFromJson(jsonString string) (*JobSpec, error) {
	js := JobSpec{}
	err := json.Unmarshal([]byte(jsonString), &js)
	if err != nil {
		return &js, err
	}
	return &js, nil
}

// FromFile loads from a json file
func FromFile(jsonFile string) (*JobSpec, error) {
	js := JobSpec{}
	file, err := os.ReadFile(jsonFile)
	if err != nil {
		return &js, err
	}
	return JobFromJson(string(file))
}

// ToJson convets back to json string
func (j *JobSpec) ToJson() (string, error) {
	out, err := json.Marshal(j)
	if err != nil {
		return "", err
	}
	return string(out), nil
}

// SetCommand is a courtesy function to set the command
// Other attributes can be set directly
func (j *JobSpec) SetCommand(command string) {
	j.Command = strings.Split(command, " ")
}

// toString converts the JobSpec to flux_jobspec1_t -> string
func (j *JobSpec) Encoded() *C.char {

	// This is the number of args the command has (length)
	argc := C.int(len(j.Command))
	commandArray := arrayToChar(j.Command)
	envArray := arrayToChar(os.Environ())

	// The command needs to be converted to char **
	jobspec := C.flux_jobspec1_from_command(
		argc,
		(**C.char)(&commandArray[0]),
		(**C.char)(&envArray[0]),
		C.int(j.Tasks),
		C.int(j.CoresPerTask),
		C.int(j.GpusPerTask),
		C.int(j.Nodes),
		C.double(j.Duration))

	for _, element := range commandArray {
		defer C.free(unsafe.Pointer(element))
	}
	for _, element := range envArray {
		defer C.free(unsafe.Pointer(element))
	}
	encoded := C.flux_jobspec1_encode(jobspec, 0)
	C.flux_jobspec1_destroy(jobspec)
	return encoded
}

// arrayToChar converts a string array to an array of char
func arrayToChar(array []string) []*C.char {

	argv := make([]*C.char, len(array))
	for i, element := range array {
		cs := C.CString(element)
		argv[i] = cs
	}
	return argv
}
