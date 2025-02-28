package service

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
	"context"
	"reflect"

	core "github.com/flux-framework/flux-go/pkg/flux"
	pb "github.com/flux-framework/flux-go/pkg/flux-grpc"
)

type Flux struct {
	pb.UnimplementedFluxServiceServer
	flux core.Flux
}

// Create a new Flux Handle
func NewFlux() *Flux {

	flux := core.NewFlux()
	return &Flux{flux: flux}
}

// Submit submits a job to flux. We can't validate / check the future
func (f *Flux) Submit(ctx context.Context, request *pb.SubmitRequest) (*pb.SubmitResponse, error) {
	response := &pb.SubmitResponse{Status: pb.SubmitResponse_SUBMIT_SUCCESS}

	// If we have a string, it is a direct jobspec (ideal)
	if reflect.TypeOf(request.Jobspec).Kind() == reflect.String {
		f.flux.SubmitJobspec(request.Jobspec)

	} else {
		js, err := core.JobFromJson(request.Jobspec)
		if err != nil {
			response.Status = pb.SubmitResponse_SUBMIT_ERROR
			return response, err
		}
		f.flux.Submit(js)
	}

	return response, nil
}

// Jobs list jobs running under the instance
func (f *Flux) Jobs(ctx context.Context, request *pb.JobsRequest) (*pb.JobsResponse, error) {
	response := &pb.JobsResponse{Status: pb.JobsResponse_JOBS_SUCCESS}
	listing, err := f.flux.ListJobs()
	if err != nil {
		response.Status = pb.JobsResponse_JOBS_ERROR
	}
	response.Jobs = listing
	return response, err
}
