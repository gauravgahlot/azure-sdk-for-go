//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsecurity

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
)

// CustomEntityStoreAssignmentsClient contains the methods for the CustomEntityStoreAssignments group.
// Don't use this type directly, use NewCustomEntityStoreAssignmentsClient() instead.
type CustomEntityStoreAssignmentsClient struct {
	ep             string
	pl             runtime.Pipeline
	subscriptionID string
}

// NewCustomEntityStoreAssignmentsClient creates a new instance of CustomEntityStoreAssignmentsClient with the specified values.
func NewCustomEntityStoreAssignmentsClient(con *arm.Connection, subscriptionID string) *CustomEntityStoreAssignmentsClient {
	return &CustomEntityStoreAssignmentsClient{ep: con.Endpoint(), pl: con.NewPipeline(module, version), subscriptionID: subscriptionID}
}

// Create - Creates a custom entity store assignment for the provided subscription, if not already exists.
// If the operation fails it returns the *CloudError error type.
func (client *CustomEntityStoreAssignmentsClient) Create(ctx context.Context, resourceGroupName string, customEntityStoreAssignmentName string, customEntityStoreAssignmentRequestBody CustomEntityStoreAssignmentRequest, options *CustomEntityStoreAssignmentsCreateOptions) (CustomEntityStoreAssignmentsCreateResponse, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, customEntityStoreAssignmentName, customEntityStoreAssignmentRequestBody, options)
	if err != nil {
		return CustomEntityStoreAssignmentsCreateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomEntityStoreAssignmentsCreateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return CustomEntityStoreAssignmentsCreateResponse{}, client.createHandleError(resp)
	}
	return client.createHandleResponse(resp)
}

// createCreateRequest creates the Create request.
func (client *CustomEntityStoreAssignmentsClient) createCreateRequest(ctx context.Context, resourceGroupName string, customEntityStoreAssignmentName string, customEntityStoreAssignmentRequestBody CustomEntityStoreAssignmentRequest, options *CustomEntityStoreAssignmentsCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customEntityStoreAssignments/{customEntityStoreAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customEntityStoreAssignmentName == "" {
		return nil, errors.New("parameter customEntityStoreAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customEntityStoreAssignmentName}", url.PathEscape(customEntityStoreAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, customEntityStoreAssignmentRequestBody)
}

// createHandleResponse handles the Create response.
func (client *CustomEntityStoreAssignmentsClient) createHandleResponse(resp *http.Response) (CustomEntityStoreAssignmentsCreateResponse, error) {
	result := CustomEntityStoreAssignmentsCreateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomEntityStoreAssignment); err != nil {
		return CustomEntityStoreAssignmentsCreateResponse{}, err
	}
	return result, nil
}

// createHandleError handles the Create error response.
func (client *CustomEntityStoreAssignmentsClient) createHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Delete - Delete a custom entity store assignment by name for a provided subscription
// If the operation fails it returns the *CloudError error type.
func (client *CustomEntityStoreAssignmentsClient) Delete(ctx context.Context, resourceGroupName string, customEntityStoreAssignmentName string, options *CustomEntityStoreAssignmentsDeleteOptions) (CustomEntityStoreAssignmentsDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, customEntityStoreAssignmentName, options)
	if err != nil {
		return CustomEntityStoreAssignmentsDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomEntityStoreAssignmentsDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return CustomEntityStoreAssignmentsDeleteResponse{}, client.deleteHandleError(resp)
	}
	return CustomEntityStoreAssignmentsDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *CustomEntityStoreAssignmentsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, customEntityStoreAssignmentName string, options *CustomEntityStoreAssignmentsDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customEntityStoreAssignments/{customEntityStoreAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customEntityStoreAssignmentName == "" {
		return nil, errors.New("parameter customEntityStoreAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customEntityStoreAssignmentName}", url.PathEscape(customEntityStoreAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// deleteHandleError handles the Delete error response.
func (client *CustomEntityStoreAssignmentsClient) deleteHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// Get - Gets a single custom entity store assignment by name for the provided subscription and resource group.
// If the operation fails it returns the *CloudError error type.
func (client *CustomEntityStoreAssignmentsClient) Get(ctx context.Context, resourceGroupName string, customEntityStoreAssignmentName string, options *CustomEntityStoreAssignmentsGetOptions) (CustomEntityStoreAssignmentsGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, customEntityStoreAssignmentName, options)
	if err != nil {
		return CustomEntityStoreAssignmentsGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return CustomEntityStoreAssignmentsGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return CustomEntityStoreAssignmentsGetResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *CustomEntityStoreAssignmentsClient) getCreateRequest(ctx context.Context, resourceGroupName string, customEntityStoreAssignmentName string, options *CustomEntityStoreAssignmentsGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customEntityStoreAssignments/{customEntityStoreAssignmentName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if customEntityStoreAssignmentName == "" {
		return nil, errors.New("parameter customEntityStoreAssignmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{customEntityStoreAssignmentName}", url.PathEscape(customEntityStoreAssignmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *CustomEntityStoreAssignmentsClient) getHandleResponse(resp *http.Response) (CustomEntityStoreAssignmentsGetResponse, error) {
	result := CustomEntityStoreAssignmentsGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomEntityStoreAssignment); err != nil {
		return CustomEntityStoreAssignmentsGetResponse{}, err
	}
	return result, nil
}

// getHandleError handles the Get error response.
func (client *CustomEntityStoreAssignmentsClient) getHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListByResourceGroup - List custom entity store assignments by a provided subscription and resource group
// If the operation fails it returns the *CloudError error type.
func (client *CustomEntityStoreAssignmentsClient) ListByResourceGroup(resourceGroupName string, options *CustomEntityStoreAssignmentsListByResourceGroupOptions) *CustomEntityStoreAssignmentsListByResourceGroupPager {
	return &CustomEntityStoreAssignmentsListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp CustomEntityStoreAssignmentsListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CustomEntityStoreAssignmentsListResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *CustomEntityStoreAssignmentsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *CustomEntityStoreAssignmentsListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/Microsoft.Security/customEntityStoreAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *CustomEntityStoreAssignmentsClient) listByResourceGroupHandleResponse(resp *http.Response) (CustomEntityStoreAssignmentsListByResourceGroupResponse, error) {
	result := CustomEntityStoreAssignmentsListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomEntityStoreAssignmentsListResult); err != nil {
		return CustomEntityStoreAssignmentsListByResourceGroupResponse{}, err
	}
	return result, nil
}

// listByResourceGroupHandleError handles the ListByResourceGroup error response.
func (client *CustomEntityStoreAssignmentsClient) listByResourceGroupHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}

// ListBySubscription - List custom entity store assignments by provided subscription
// If the operation fails it returns the *CloudError error type.
func (client *CustomEntityStoreAssignmentsClient) ListBySubscription(options *CustomEntityStoreAssignmentsListBySubscriptionOptions) *CustomEntityStoreAssignmentsListBySubscriptionPager {
	return &CustomEntityStoreAssignmentsListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp CustomEntityStoreAssignmentsListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.CustomEntityStoreAssignmentsListResult.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *CustomEntityStoreAssignmentsClient) listBySubscriptionCreateRequest(ctx context.Context, options *CustomEntityStoreAssignmentsListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Security/customEntityStoreAssignments"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.ep, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *CustomEntityStoreAssignmentsClient) listBySubscriptionHandleResponse(resp *http.Response) (CustomEntityStoreAssignmentsListBySubscriptionResponse, error) {
	result := CustomEntityStoreAssignmentsListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.CustomEntityStoreAssignmentsListResult); err != nil {
		return CustomEntityStoreAssignmentsListBySubscriptionResponse{}, err
	}
	return result, nil
}

// listBySubscriptionHandleError handles the ListBySubscription error response.
func (client *CustomEntityStoreAssignmentsClient) listBySubscriptionHandleError(resp *http.Response) error {
	body, err := runtime.Payload(resp)
	if err != nil {
		return runtime.NewResponseError(err, resp)
	}
	errType := CloudError{raw: string(body)}
	if err := runtime.UnmarshalAsJSON(resp, &errType.InnerError); err != nil {
		return runtime.NewResponseError(fmt.Errorf("%s\n%s", string(body), err), resp)
	}
	return runtime.NewResponseError(&errType, resp)
}