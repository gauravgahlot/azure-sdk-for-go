// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armredis

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"time"
)

// LinkedServerCreatePoller provides polling facilities until the operation reaches a terminal state.
type LinkedServerCreatePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final LinkedServerCreateResponse will be returned.
	FinalResponse(ctx context.Context) (LinkedServerCreateResponse, error)
}

type linkedServerCreatePoller struct {
	pt *armcore.LROPoller
}

func (p *linkedServerCreatePoller) Done() bool {
	return p.pt.Done()
}

func (p *linkedServerCreatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *linkedServerCreatePoller) FinalResponse(ctx context.Context) (LinkedServerCreateResponse, error) {
	respType := LinkedServerCreateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.RedisLinkedServerWithProperties)
	if err != nil {
		return LinkedServerCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *linkedServerCreatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *linkedServerCreatePoller) pollUntilDone(ctx context.Context, freq time.Duration) (LinkedServerCreateResponse, error) {
	respType := LinkedServerCreateResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.RedisLinkedServerWithProperties)
	if err != nil {
		return LinkedServerCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// PrivateEndpointConnectionsPutPoller provides polling facilities until the operation reaches a terminal state.
type PrivateEndpointConnectionsPutPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final PrivateEndpointConnectionsPutResponse will be returned.
	FinalResponse(ctx context.Context) (PrivateEndpointConnectionsPutResponse, error)
}

type privateEndpointConnectionsPutPoller struct {
	pt *armcore.LROPoller
}

func (p *privateEndpointConnectionsPutPoller) Done() bool {
	return p.pt.Done()
}

func (p *privateEndpointConnectionsPutPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *privateEndpointConnectionsPutPoller) FinalResponse(ctx context.Context) (PrivateEndpointConnectionsPutResponse, error) {
	respType := PrivateEndpointConnectionsPutResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionsPutResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *privateEndpointConnectionsPutPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *privateEndpointConnectionsPutPoller) pollUntilDone(ctx context.Context, freq time.Duration) (PrivateEndpointConnectionsPutResponse, error) {
	respType := PrivateEndpointConnectionsPutResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.PrivateEndpointConnection)
	if err != nil {
		return PrivateEndpointConnectionsPutResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// RedisCreatePoller provides polling facilities until the operation reaches a terminal state.
type RedisCreatePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final RedisCreateResponse will be returned.
	FinalResponse(ctx context.Context) (RedisCreateResponse, error)
}

type redisCreatePoller struct {
	pt *armcore.LROPoller
}

func (p *redisCreatePoller) Done() bool {
	return p.pt.Done()
}

func (p *redisCreatePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *redisCreatePoller) FinalResponse(ctx context.Context) (RedisCreateResponse, error) {
	respType := RedisCreateResponse{}
	resp, err := p.pt.FinalResponse(ctx, &respType.RedisResource)
	if err != nil {
		return RedisCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *redisCreatePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *redisCreatePoller) pollUntilDone(ctx context.Context, freq time.Duration) (RedisCreateResponse, error) {
	respType := RedisCreateResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, &respType.RedisResource)
	if err != nil {
		return RedisCreateResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// RedisDeletePoller provides polling facilities until the operation reaches a terminal state.
type RedisDeletePoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final RedisDeleteResponse will be returned.
	FinalResponse(ctx context.Context) (RedisDeleteResponse, error)
}

type redisDeletePoller struct {
	pt *armcore.LROPoller
}

func (p *redisDeletePoller) Done() bool {
	return p.pt.Done()
}

func (p *redisDeletePoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *redisDeletePoller) FinalResponse(ctx context.Context) (RedisDeleteResponse, error) {
	respType := RedisDeleteResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return RedisDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *redisDeletePoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *redisDeletePoller) pollUntilDone(ctx context.Context, freq time.Duration) (RedisDeleteResponse, error) {
	respType := RedisDeleteResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return RedisDeleteResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// RedisExportDataPoller provides polling facilities until the operation reaches a terminal state.
type RedisExportDataPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final RedisExportDataResponse will be returned.
	FinalResponse(ctx context.Context) (RedisExportDataResponse, error)
}

type redisExportDataPoller struct {
	pt *armcore.LROPoller
}

func (p *redisExportDataPoller) Done() bool {
	return p.pt.Done()
}

func (p *redisExportDataPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *redisExportDataPoller) FinalResponse(ctx context.Context) (RedisExportDataResponse, error) {
	respType := RedisExportDataResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return RedisExportDataResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *redisExportDataPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *redisExportDataPoller) pollUntilDone(ctx context.Context, freq time.Duration) (RedisExportDataResponse, error) {
	respType := RedisExportDataResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return RedisExportDataResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

// RedisImportDataPoller provides polling facilities until the operation reaches a terminal state.
type RedisImportDataPoller interface {
	azcore.Poller
	// FinalResponse performs a final GET to the service and returns the final response
	// for the polling operation. If there is an error performing the final GET then an error is returned.
	// If the final GET succeeded then the final RedisImportDataResponse will be returned.
	FinalResponse(ctx context.Context) (RedisImportDataResponse, error)
}

type redisImportDataPoller struct {
	pt *armcore.LROPoller
}

func (p *redisImportDataPoller) Done() bool {
	return p.pt.Done()
}

func (p *redisImportDataPoller) Poll(ctx context.Context) (*http.Response, error) {
	return p.pt.Poll(ctx)
}

func (p *redisImportDataPoller) FinalResponse(ctx context.Context) (RedisImportDataResponse, error) {
	respType := RedisImportDataResponse{}
	resp, err := p.pt.FinalResponse(ctx, nil)
	if err != nil {
		return RedisImportDataResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}

func (p *redisImportDataPoller) ResumeToken() (string, error) {
	return p.pt.ResumeToken()
}

func (p *redisImportDataPoller) pollUntilDone(ctx context.Context, freq time.Duration) (RedisImportDataResponse, error) {
	respType := RedisImportDataResponse{}
	resp, err := p.pt.PollUntilDone(ctx, freq, nil)
	if err != nil {
		return RedisImportDataResponse{}, err
	}
	respType.RawResponse = resp
	return respType, nil
}