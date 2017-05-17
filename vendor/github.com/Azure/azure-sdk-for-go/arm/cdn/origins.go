package cdn

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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// OriginsClient is the use these APIs to manage Azure CDN resources through
// the Azure Resource Manager. You must make sure that requests made to these
// resources are secure.
type OriginsClient struct {
	ManagementClient
}

// NewOriginsClient creates an instance of the OriginsClient client.
func NewOriginsClient(subscriptionID string) OriginsClient {
	return NewOriginsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewOriginsClientWithBaseURI creates an instance of the OriginsClient client.
func NewOriginsClientWithBaseURI(baseURI string, subscriptionID string) OriginsClient {
	return OriginsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get gets an existing origin within an endpoint.
//
// resourceGroupName is name of the Resource group within the Azure
// subscription. profileName is name of the CDN profile which is unique within
// the resource group. endpointName is name of the endpoint under the profile
// which is unique globally. originName is name of the origin which is unique
// within the endpoint.
func (client OriginsClient) Get(resourceGroupName string, profileName string, endpointName string, originName string) (result Origin, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "cdn.OriginsClient", "Get")
	}

	req, err := client.GetPreparer(resourceGroupName, profileName, endpointName, originName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client OriginsClient) GetPreparer(resourceGroupName string, profileName string, endpointName string, originName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"originName":        autorest.Encode("path", originName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client OriginsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client OriginsClient) GetResponder(resp *http.Response) (result Origin, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEndpoint lists all of the existing origins within an endpoint.
//
// resourceGroupName is name of the Resource group within the Azure
// subscription. profileName is name of the CDN profile which is unique within
// the resource group. endpointName is name of the endpoint under the profile
// which is unique globally.
func (client OriginsClient) ListByEndpoint(resourceGroupName string, profileName string, endpointName string) (result OriginListResult, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "cdn.OriginsClient", "ListByEndpoint")
	}

	req, err := client.ListByEndpointPreparer(resourceGroupName, profileName, endpointName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "ListByEndpoint", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByEndpointSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "ListByEndpoint", resp, "Failure sending request")
		return
	}

	result, err = client.ListByEndpointResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "ListByEndpoint", resp, "Failure responding to request")
	}

	return
}

// ListByEndpointPreparer prepares the ListByEndpoint request.
func (client OriginsClient) ListByEndpointPreparer(resourceGroupName string, profileName string, endpointName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListByEndpointSender sends the ListByEndpoint request. The method will close the
// http.Response Body if it receives an error.
func (client OriginsClient) ListByEndpointSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListByEndpointResponder handles the response to the ListByEndpoint request. The method always
// closes the http.Response Body.
func (client OriginsClient) ListByEndpointResponder(resp *http.Response) (result OriginListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByEndpointNextResults retrieves the next set of results, if any.
func (client OriginsClient) ListByEndpointNextResults(lastResults OriginListResult) (result OriginListResult, err error) {
	req, err := lastResults.OriginListResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "cdn.OriginsClient", "ListByEndpoint", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListByEndpointSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "cdn.OriginsClient", "ListByEndpoint", resp, "Failure sending next results request")
	}

	result, err = client.ListByEndpointResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "ListByEndpoint", resp, "Failure responding to next results request")
	}

	return
}

// Update updates an existing origin within an endpoint. This method may poll
// for completion. Polling can be canceled by passing the cancel channel
// argument. The channel will be used to cancel polling and any outstanding
// HTTP requests.
//
// resourceGroupName is name of the Resource group within the Azure
// subscription. profileName is name of the CDN profile which is unique within
// the resource group. endpointName is name of the endpoint under the profile
// which is unique globally. originName is name of the origin which is unique
// within the endpoint. originUpdateProperties is origin properties
func (client OriginsClient) Update(resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties OriginUpdateParameters, cancel <-chan struct{}) (<-chan Origin, <-chan error) {
	resultChan := make(chan Origin, 1)
	errChan := make(chan error, 1)
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		errChan <- validation.NewErrorWithValidationError(err, "cdn.OriginsClient", "Update")
		close(errChan)
		close(resultChan)
		return resultChan, errChan
	}

	go func() {
		var err error
		var result Origin
		defer func() {
			resultChan <- result
			errChan <- err
			close(resultChan)
			close(errChan)
		}()
		req, err := client.UpdatePreparer(resourceGroupName, profileName, endpointName, originName, originUpdateProperties, cancel)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Update", nil, "Failure preparing request")
			return
		}

		resp, err := client.UpdateSender(req)
		if err != nil {
			result.Response = autorest.Response{Response: resp}
			err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Update", resp, "Failure sending request")
			return
		}

		result, err = client.UpdateResponder(resp)
		if err != nil {
			err = autorest.NewErrorWithError(err, "cdn.OriginsClient", "Update", resp, "Failure responding to request")
		}
	}()
	return resultChan, errChan
}

// UpdatePreparer prepares the Update request.
func (client OriginsClient) UpdatePreparer(resourceGroupName string, profileName string, endpointName string, originName string, originUpdateProperties OriginUpdateParameters, cancel <-chan struct{}) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"endpointName":      autorest.Encode("path", endpointName),
		"originName":        autorest.Encode("path", originName),
		"profileName":       autorest.Encode("path", profileName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-10-02"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cdn/profiles/{profileName}/endpoints/{endpointName}/origins/{originName}", pathParameters),
		autorest.WithJSON(originUpdateProperties),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{Cancel: cancel})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client OriginsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client,
		req,
		azure.DoPollForAsynchronous(client.PollingDelay))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client OriginsClient) UpdateResponder(resp *http.Response) (result Origin, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
