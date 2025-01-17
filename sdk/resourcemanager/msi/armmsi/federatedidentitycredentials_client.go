//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package armmsi

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// FederatedIdentityCredentialsClient contains the methods for the FederatedIdentityCredentials group.
// Don't use this type directly, use NewFederatedIdentityCredentialsClient() instead.
type FederatedIdentityCredentialsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFederatedIdentityCredentialsClient creates a new instance of FederatedIdentityCredentialsClient with the specified values.
// subscriptionID - The Id of the Subscription to which the identity belongs.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFederatedIdentityCredentialsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*FederatedIdentityCredentialsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublic.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &FederatedIdentityCredentialsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// CreateOrUpdate - Create or update a federated identity credential under the specified user assigned identity.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-31-preview
// resourceGroupName - The name of the Resource Group to which the identity belongs.
// resourceName - The name of the identity resource.
// federatedIdentityCredentialResourceName - The name of the federated identity credential resource.
// parameters - Parameters to create or update the federated identity credential.
// options - FederatedIdentityCredentialsClientCreateOrUpdateOptions contains the optional parameters for the FederatedIdentityCredentialsClient.CreateOrUpdate
// method.
func (client *FederatedIdentityCredentialsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, resourceName string, federatedIdentityCredentialResourceName string, parameters FederatedIdentityCredential, options *FederatedIdentityCredentialsClientCreateOrUpdateOptions) (FederatedIdentityCredentialsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, parameters, options)
	if err != nil {
		return FederatedIdentityCredentialsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FederatedIdentityCredentialsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return FederatedIdentityCredentialsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FederatedIdentityCredentialsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, federatedIdentityCredentialResourceName string, parameters FederatedIdentityCredential, options *FederatedIdentityCredentialsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials/{federatedIdentityCredentialResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if federatedIdentityCredentialResourceName == "" {
		return nil, errors.New("parameter federatedIdentityCredentialResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{federatedIdentityCredentialResourceName}", url.PathEscape(federatedIdentityCredentialResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, runtime.MarshalAsJSON(req, parameters)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *FederatedIdentityCredentialsClient) createOrUpdateHandleResponse(resp *http.Response) (FederatedIdentityCredentialsClientCreateOrUpdateResponse, error) {
	result := FederatedIdentityCredentialsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FederatedIdentityCredential); err != nil {
		return FederatedIdentityCredentialsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Deletes the federated identity credential.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-31-preview
// resourceGroupName - The name of the Resource Group to which the identity belongs.
// resourceName - The name of the identity resource.
// federatedIdentityCredentialResourceName - The name of the federated identity credential resource.
// options - FederatedIdentityCredentialsClientDeleteOptions contains the optional parameters for the FederatedIdentityCredentialsClient.Delete
// method.
func (client *FederatedIdentityCredentialsClient) Delete(ctx context.Context, resourceGroupName string, resourceName string, federatedIdentityCredentialResourceName string, options *FederatedIdentityCredentialsClientDeleteOptions) (FederatedIdentityCredentialsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, options)
	if err != nil {
		return FederatedIdentityCredentialsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FederatedIdentityCredentialsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return FederatedIdentityCredentialsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return FederatedIdentityCredentialsClientDeleteResponse{}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FederatedIdentityCredentialsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, federatedIdentityCredentialResourceName string, options *FederatedIdentityCredentialsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials/{federatedIdentityCredentialResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if federatedIdentityCredentialResourceName == "" {
		return nil, errors.New("parameter federatedIdentityCredentialResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{federatedIdentityCredentialResourceName}", url.PathEscape(federatedIdentityCredentialResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the federated identity credential.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-31-preview
// resourceGroupName - The name of the Resource Group to which the identity belongs.
// resourceName - The name of the identity resource.
// federatedIdentityCredentialResourceName - The name of the federated identity credential resource.
// options - FederatedIdentityCredentialsClientGetOptions contains the optional parameters for the FederatedIdentityCredentialsClient.Get
// method.
func (client *FederatedIdentityCredentialsClient) Get(ctx context.Context, resourceGroupName string, resourceName string, federatedIdentityCredentialResourceName string, options *FederatedIdentityCredentialsClientGetOptions) (FederatedIdentityCredentialsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, resourceName, federatedIdentityCredentialResourceName, options)
	if err != nil {
		return FederatedIdentityCredentialsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FederatedIdentityCredentialsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FederatedIdentityCredentialsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FederatedIdentityCredentialsClient) getCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, federatedIdentityCredentialResourceName string, options *FederatedIdentityCredentialsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials/{federatedIdentityCredentialResourceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	if federatedIdentityCredentialResourceName == "" {
		return nil, errors.New("parameter federatedIdentityCredentialResourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{federatedIdentityCredentialResourceName}", url.PathEscape(federatedIdentityCredentialResourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-01-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FederatedIdentityCredentialsClient) getHandleResponse(resp *http.Response) (FederatedIdentityCredentialsClientGetResponse, error) {
	result := FederatedIdentityCredentialsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FederatedIdentityCredential); err != nil {
		return FederatedIdentityCredentialsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the federated identity credentials under the specified user assigned identity.
// If the operation fails it returns an *azcore.ResponseError type.
// Generated from API version 2022-01-31-preview
// resourceGroupName - The name of the Resource Group to which the identity belongs.
// resourceName - The name of the identity resource.
// options - FederatedIdentityCredentialsClientListOptions contains the optional parameters for the FederatedIdentityCredentialsClient.List
// method.
func (client *FederatedIdentityCredentialsClient) NewListPager(resourceGroupName string, resourceName string, options *FederatedIdentityCredentialsClientListOptions) *runtime.Pager[FederatedIdentityCredentialsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[FederatedIdentityCredentialsClientListResponse]{
		More: func(page FederatedIdentityCredentialsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FederatedIdentityCredentialsClientListResponse) (FederatedIdentityCredentialsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, resourceName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FederatedIdentityCredentialsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FederatedIdentityCredentialsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FederatedIdentityCredentialsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FederatedIdentityCredentialsClient) listCreateRequest(ctx context.Context, resourceGroupName string, resourceName string, options *FederatedIdentityCredentialsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ManagedIdentity/userAssignedIdentities/{resourceName}/federatedIdentityCredentials"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if resourceName == "" {
		return nil, errors.New("parameter resourceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceName}", url.PathEscape(resourceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	reqQP.Set("api-version", "2022-01-31-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FederatedIdentityCredentialsClient) listHandleResponse(resp *http.Response) (FederatedIdentityCredentialsClientListResponse, error) {
	result := FederatedIdentityCredentialsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FederatedIdentityCredentialsListResult); err != nil {
		return FederatedIdentityCredentialsClientListResponse{}, err
	}
	return result, nil
}
