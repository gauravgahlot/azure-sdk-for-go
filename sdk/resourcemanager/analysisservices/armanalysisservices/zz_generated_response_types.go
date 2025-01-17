//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armanalysisservices

// OperationsClientListResponse contains the response from method OperationsClient.List.
type OperationsClientListResponse struct {
	OperationListResult
}

// ServersClientCheckNameAvailabilityResponse contains the response from method ServersClient.CheckNameAvailability.
type ServersClientCheckNameAvailabilityResponse struct {
	CheckServerNameAvailabilityResult
}

// ServersClientCreateResponse contains the response from method ServersClient.Create.
type ServersClientCreateResponse struct {
	Server
}

// ServersClientDeleteResponse contains the response from method ServersClient.Delete.
type ServersClientDeleteResponse struct {
	// placeholder for future response values
}

// ServersClientDissociateGatewayResponse contains the response from method ServersClient.DissociateGateway.
type ServersClientDissociateGatewayResponse struct {
	// placeholder for future response values
}

// ServersClientGetDetailsResponse contains the response from method ServersClient.GetDetails.
type ServersClientGetDetailsResponse struct {
	Server
}

// ServersClientListByResourceGroupResponse contains the response from method ServersClient.ListByResourceGroup.
type ServersClientListByResourceGroupResponse struct {
	Servers
}

// ServersClientListGatewayStatusResponse contains the response from method ServersClient.ListGatewayStatus.
type ServersClientListGatewayStatusResponse struct {
	GatewayListStatusLive
}

// ServersClientListOperationResultsResponse contains the response from method ServersClient.ListOperationResults.
type ServersClientListOperationResultsResponse struct {
	// placeholder for future response values
}

// ServersClientListOperationStatusesResponse contains the response from method ServersClient.ListOperationStatuses.
type ServersClientListOperationStatusesResponse struct {
	OperationStatus
}

// ServersClientListResponse contains the response from method ServersClient.List.
type ServersClientListResponse struct {
	Servers
}

// ServersClientListSKUsForExistingResponse contains the response from method ServersClient.ListSKUsForExisting.
type ServersClientListSKUsForExistingResponse struct {
	SKUEnumerationForExistingResourceResult
}

// ServersClientListSKUsForNewResponse contains the response from method ServersClient.ListSKUsForNew.
type ServersClientListSKUsForNewResponse struct {
	SKUEnumerationForNewResourceResult
}

// ServersClientResumeResponse contains the response from method ServersClient.Resume.
type ServersClientResumeResponse struct {
	// placeholder for future response values
}

// ServersClientSuspendResponse contains the response from method ServersClient.Suspend.
type ServersClientSuspendResponse struct {
	// placeholder for future response values
}

// ServersClientUpdateResponse contains the response from method ServersClient.Update.
type ServersClientUpdateResponse struct {
	Server
}
