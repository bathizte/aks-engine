package devtestlabs

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator 1.0.1.0
// Changes may cause incorrect behavior and will be lost if the code is
// regenerated.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// VirtualNetworksClient is the the DevTest Labs Client.
type VirtualNetworksClient struct {
	ManagementClient
}

// NewVirtualNetworksClient creates an instance of the VirtualNetworksClient
// client.
func NewVirtualNetworksClient(subscriptionID string) VirtualNetworksClient {
	return NewVirtualNetworksClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVirtualNetworksClientWithBaseURI creates an instance of the
// VirtualNetworksClient client.
func NewVirtualNetworksClientWithBaseURI(baseURI string, subscriptionID string) VirtualNetworksClient {
	return VirtualNetworksClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or replace an existing virtual network. This operation
// can take a while to complete. This method may poll for completion. Polling
// can be canceled by passing the cancel channel argument. The channel will be
// used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the virtual network. virtualNetwork is a
// virtual network.
func (client VirtualNetworksClient) CreateOrUpdate(resourceGroupName string, labName string, name string, virtualNetwork VirtualNetwork, cancel <-chan struct{}) (<-chan VirtualNetwork, <-chan error) {
	resultChan := make(chan VirtualNetwork, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result VirtualNetwork
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.CreateOrUpdatePreparer(resourceGroupName, labName, name, virtualNetwork, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "CreateOrUpdate", nil, "Failure preparing request")
			return
		}

		resp, err := client.CreateOrUpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "CreateOrUpdate", resp, "Failure sending request")
			return
		}

		result, err = client.CreateOrUpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "CreateOrUpdate", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client VirtualNetworksClient) CreateOrUpdatePreparer(resourceGroupName string, labName string, name string, virtualNetwork VirtualNetwork, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", pathParameters),
		autorest.WithJSON(virtualNetwork),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) CreateOrUpdateResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete virtual network. This operation can take a while to complete.
// This method may poll for completion. Polling can be canceled by passing the
// cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the virtual network.
func (client VirtualNetworksClient) Delete(resourceGroupName string, labName string, name string, cancel <-chan struct{}) (<-chan autorest.Response, <-chan error) {
	resultChan := make(chan autorest.Response, 1)
	errChan := make(chan error, 1)
	go func() {
		var err error
		var result autorest.Response
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.DeletePreparer(resourceGroupName, labName, name, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "Delete", nil, "Failure preparing request")
			return
		}

		resp, err := client.DeleteSender(req)
		if err != nil {
			result.Response = resp
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "Delete", resp, "Failure sending request")
			return
		}

		result, err = client.DeleteResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "Delete", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// DeletePreparer prepares the Delete request.
func (client VirtualNetworksClient) DeletePreparer(resourceGroupName string, labName string, name string, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get virtual network.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the virtual network. expand is specify the
// $expand query. Example: 'properties($expand=externalSubnets)'
func (client VirtualNetworksClient) Get(resourceGroupName string, labName string, name string, expand string) (result VirtualNetwork, err error) {
	req, err := client.GetPreparer(resourceGroupName, labName, name, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client VirtualNetworksClient) GetPreparer(resourceGroupName string, labName string, name string, expand string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) GetResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list virtual networks in a given lab.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. expand is specify the $expand query. Example:
// 'properties($expand=externalSubnets)' filter is the filter to apply to the
// operation. top is the maximum number of resources to return from the
// operation. orderby is the ordering expression for the results, using OData
// notation.
func (client VirtualNetworksClient) List(resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationVirtualNetwork, err error) {
	req, err := client.ListPreparer(resourceGroupName, labName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client VirtualNetworksClient) ListPreparer(resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(expand) > 0 {
		queryParameters["$expand"] = autorest.Encode("query", expand)
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}
	if top != nil {
		queryParameters["$top"] = autorest.Encode("query", *top)
	}
	if len(orderby) > 0 {
		queryParameters["$orderby"] = autorest.Encode("query", orderby)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) ListResponder(resp *http.Response) (result ResponseWithContinuationVirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListNextResults retrieves the next set of results, if any.
func (client VirtualNetworksClient) ListNextResults(lastResults ResponseWithContinuationVirtualNetwork) (result ResponseWithContinuationVirtualNetwork, err error) {
	req, err := lastResults.ResponseWithContinuationVirtualNetworkPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// Update modify properties of virtual networks.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the virtual network. virtualNetwork is a
// virtual network.
func (client VirtualNetworksClient) Update(resourceGroupName string, labName string, name string, virtualNetwork VirtualNetworkFragment) (result VirtualNetwork, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, labName, name, virtualNetwork)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.VirtualNetworksClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client VirtualNetworksClient) UpdatePreparer(resourceGroupName string, labName string, name string, virtualNetwork VirtualNetworkFragment) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"labName":           autorest.Encode("path", labName),
		"name":              autorest.Encode("path", name),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/virtualnetworks/{name}", pathParameters),
		autorest.WithJSON(virtualNetwork),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client VirtualNetworksClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client VirtualNetworksClient) UpdateResponder(resp *http.Response) (result VirtualNetwork, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
