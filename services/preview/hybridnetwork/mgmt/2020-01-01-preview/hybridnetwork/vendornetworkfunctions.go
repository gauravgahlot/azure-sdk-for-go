package hybridnetwork

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// VendorNetworkFunctionsClient is the client for the VendorNetworkFunctions methods of the Hybridnetwork service.
type VendorNetworkFunctionsClient struct {
	BaseClient
}

// NewVendorNetworkFunctionsClient creates an instance of the VendorNetworkFunctionsClient client.
func NewVendorNetworkFunctionsClient(subscriptionID string) VendorNetworkFunctionsClient {
	return NewVendorNetworkFunctionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVendorNetworkFunctionsClientWithBaseURI creates an instance of the VendorNetworkFunctionsClient client using a
// custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds,
// Azure stack).
func NewVendorNetworkFunctionsClientWithBaseURI(baseURI string, subscriptionID string) VendorNetworkFunctionsClient {
	return VendorNetworkFunctionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a vendor network function.
// Parameters:
// locationName - the Azure region where the network function resource was created by the customer.
// vendorName - the name of the vendor.
// serviceKey - the GUID for the vendor network function.
// parameters - parameters supplied to the create or update vendor network function operation.
func (client VendorNetworkFunctionsClient) CreateOrUpdate(ctx context.Context, locationName string, vendorName string, serviceKey string, parameters VendorNetworkFunction) (result VendorNetworkFunctionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorNetworkFunctionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybridnetwork.VendorNetworkFunctionsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, locationName, vendorName, serviceKey, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorNetworkFunctionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorNetworkFunctionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VendorNetworkFunctionsClient) CreateOrUpdatePreparer(ctx context.Context, locationName string, vendorName string, serviceKey string, parameters VendorNetworkFunction) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"serviceKey":     autorest.Encode("path", serviceKey),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"vendorName":     autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VendorNetworkFunctionsClient) CreateOrUpdateSender(req *http.Request) (future VendorNetworkFunctionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VendorNetworkFunctionsClient) CreateOrUpdateResponder(resp *http.Response) (result VendorNetworkFunction, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets information about the specified vendor network function.
// Parameters:
// locationName - the Azure region where the network function resource was created by the customer.
// vendorName - the name of the vendor.
// serviceKey - the GUID for the vendor network function.
func (client VendorNetworkFunctionsClient) Get(ctx context.Context, locationName string, vendorName string, serviceKey string) (result VendorNetworkFunction, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorNetworkFunctionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybridnetwork.VendorNetworkFunctionsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, locationName, vendorName, serviceKey)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorNetworkFunctionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorNetworkFunctionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorNetworkFunctionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client VendorNetworkFunctionsClient) GetPreparer(ctx context.Context, locationName string, vendorName string, serviceKey string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"serviceKey":     autorest.Encode("path", serviceKey),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"vendorName":     autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions/{serviceKey}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VendorNetworkFunctionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VendorNetworkFunctionsClient) GetResponder(resp *http.Response) (result VendorNetworkFunction, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List lists all the vendor network function sub resources in an Azure region, filtered by skuType, skuName,
// vendorProvisioningState.
// Parameters:
// locationName - the Azure region where the network function resource was created by the customer.
// vendorName - the name of the vendor.
// filter - the filter to apply on the operation. The properties you can use for eq (equals) are: skuType,
// skuName and vendorProvisioningState.
func (client VendorNetworkFunctionsClient) List(ctx context.Context, locationName string, vendorName string, filter string) (result VendorNetworkFunctionListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorNetworkFunctionsClient.List")
		defer func() {
			sc := -1
			if result.vnflr.Response.Response != nil {
				sc = result.vnflr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("hybridnetwork.VendorNetworkFunctionsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, locationName, vendorName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorNetworkFunctionsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.vnflr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorNetworkFunctionsClient", "List", resp, "Failure sending request")
		return
	}

	result.vnflr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorNetworkFunctionsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.vnflr.hasNextLink() && result.vnflr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client VendorNetworkFunctionsClient) ListPreparer(ctx context.Context, locationName string, vendorName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"locationName":   autorest.Encode("path", locationName),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
		"vendorName":     autorest.Encode("path", vendorName),
	}

	const APIVersion = "2020-01-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.HybridNetwork/locations/{locationName}/vendors/{vendorName}/networkFunctions", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VendorNetworkFunctionsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VendorNetworkFunctionsClient) ListResponder(resp *http.Response) (result VendorNetworkFunctionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client VendorNetworkFunctionsClient) listNextResults(ctx context.Context, lastResults VendorNetworkFunctionListResult) (result VendorNetworkFunctionListResult, err error) {
	req, err := lastResults.vendorNetworkFunctionListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "hybridnetwork.VendorNetworkFunctionsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "hybridnetwork.VendorNetworkFunctionsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hybridnetwork.VendorNetworkFunctionsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client VendorNetworkFunctionsClient) ListComplete(ctx context.Context, locationName string, vendorName string, filter string) (result VendorNetworkFunctionListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VendorNetworkFunctionsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, locationName, vendorName, filter)
	return
}
