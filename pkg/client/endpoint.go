package client

import (
	"context"
	"errors"
	"fmt"
	"time"

	core "github.com/flux-framework/flux-go/pkg/flux"
	pb "github.com/flux-framework/flux-go/pkg/flux-grpc"
)

// Submit will submit a job to flux
func (c *FluxClient) Submit(ctx context.Context, js *core.JobSpec) (*pb.SubmitResponse, error) {
	response := &pb.SubmitResponse{}
	if !c.Connected() {
		return response, fmt.Errorf("client is not connected")
	}
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	jobspec, err := js.ToJson()
	request := pb.SubmitRequest{Jobspec: jobspec}
	response, err = c.service.Submit(ctx, &request)
	if err != nil {
		fmt.Printf("[Flux] issue with submit %v\n", err)
	}
	return response, err
}

// Jobs will do a request to ListJobs
func (c *FluxClient) Jobs(ctx context.Context) (*pb.JobsResponse, error) {

	response := &pb.JobsResponse{}
	if !c.Connected() {
		return response, errors.New("client is not connected")
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()
	response, err := c.service.Jobs(ctx, &pb.JobsRequest{})
	if err != nil {
		fmt.Printf("[Flux] issue with submit %v\n", err)
	}
	return response, err
}
