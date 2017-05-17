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
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// NotificationChannelsClient is the the DevTest Labs Client.
type NotificationChannelsClient struct {
	ManagementClient
}

// NewNotificationChannelsClient creates an instance of the
// NotificationChannelsClient client.
func NewNotificationChannelsClient(subscriptionID string) NotificationChannelsClient {
	return NewNotificationChannelsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewNotificationChannelsClientWithBaseURI creates an instance of the
// NotificationChannelsClient client.
func NewNotificationChannelsClientWithBaseURI(baseURI string, subscriptionID string) NotificationChannelsClient {
	return NotificationChannelsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or replace an existing notificationChannel.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the notificationChannel. notificationChannel is
// a notification.
func (client NotificationChannelsClient) CreateOrUpdate(resourceGroupName string, labName string, name string, notificationChannel NotificationChannel) (result NotificationChannel, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: notificationChannel,
			Constraints: []validation.Constraint{{Target: "notificationChannel.NotificationChannelProperties", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "devtestlabs.NotificationChannelsClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(resourceGroupName, labName, name, notificationChannel)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "CreateOrUpdate", resp, "Failure responding to request")
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client NotificationChannelsClient) CreateOrUpdatePreparer(resourceGroupName string, labName string, name string, notificationChannel NotificationChannel) (*http.Request, error) {
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", pathParameters),
		autorest.WithJSON(notificationChannel),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client NotificationChannelsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client NotificationChannelsClient) CreateOrUpdateResponder(resp *http.Response) (result NotificationChannel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete notificationchannel.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the notificationChannel.
func (client NotificationChannelsClient) Delete(resourceGroupName string, labName string, name string) (result autorest.Response, err error) {
	req, err := client.DeletePreparer(resourceGroupName, labName, name)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "Delete", resp, "Failure responding to request")
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client NotificationChannelsClient) DeletePreparer(resourceGroupName string, labName string, name string) (*http.Request, error) {
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client NotificationChannelsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client NotificationChannelsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get notificationchannel.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the notificationChannel. expand is specify the
// $expand query. Example: 'properties($select=webHookUrl)'
func (client NotificationChannelsClient) Get(resourceGroupName string, labName string, name string, expand string) (result NotificationChannel, err error) {
	req, err := client.GetPreparer(resourceGroupName, labName, name, expand)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client NotificationChannelsClient) GetPreparer(resourceGroupName string, labName string, name string, expand string) (*http.Request, error) {
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client NotificationChannelsClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client NotificationChannelsClient) GetResponder(resp *http.Response) (result NotificationChannel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List list notificationchannels in a given lab.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. expand is specify the $expand query. Example:
// 'properties($select=webHookUrl)' filter is the filter to apply to the
// operation. top is the maximum number of resources to return from the
// operation. orderby is the ordering expression for the results, using OData
// notation.
func (client NotificationChannelsClient) List(resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (result ResponseWithContinuationNotificationChannel, err error) {
	req, err := client.ListPreparer(resourceGroupName, labName, expand, filter, top, orderby)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "List", resp, "Failure sending request")
		return
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client NotificationChannelsClient) ListPreparer(resourceGroupName string, labName string, expand string, filter string, top *int32, orderby string) (*http.Request, error) {
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client NotificationChannelsClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client NotificationChannelsClient) ListResponder(resp *http.Response) (result ResponseWithContinuationNotificationChannel, err error) {
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
func (client NotificationChannelsClient) ListNextResults(lastResults ResponseWithContinuationNotificationChannel) (result ResponseWithContinuationNotificationChannel, err error) {
	req, err := lastResults.ResponseWithContinuationNotificationChannelPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "List", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "List", resp, "Failure sending next results request")
	}

	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "List", resp, "Failure responding to next results request")
	}

	return
}

// Notify send notification to provided channel.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the notificationChannel. notifyParameters is
// properties for generating a Notification.
func (client NotificationChannelsClient) Notify(resourceGroupName string, labName string, name string, notifyParameters NotifyParameters) (result autorest.Response, err error) {
	req, err := client.NotifyPreparer(resourceGroupName, labName, name, notifyParameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "Notify", nil, "Failure preparing request")
		return
	}

	resp, err := client.NotifySender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "Notify", resp, "Failure sending request")
		return
	}

	result, err = client.NotifyResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "Notify", resp, "Failure responding to request")
	}

	return
}

// NotifyPreparer prepares the Notify request.
func (client NotificationChannelsClient) NotifyPreparer(resourceGroupName string, labName string, name string, notifyParameters NotifyParameters) (*http.Request, error) {
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
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}/notify", pathParameters),
		autorest.WithJSON(notifyParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// NotifySender sends the Notify request. The method will close the
// http.Response Body if it receives an error.
func (client NotificationChannelsClient) NotifySender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// NotifyResponder handles the response to the Notify request. The method always
// closes the http.Response Body.
func (client NotificationChannelsClient) NotifyResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Update modify properties of notificationchannels.
//
// resourceGroupName is the name of the resource group. labName is the name of
// the lab. name is the name of the notificationChannel. notificationChannel is
// a notification.
func (client NotificationChannelsClient) Update(resourceGroupName string, labName string, name string, notificationChannel NotificationChannelFragment) (result NotificationChannel, err error) {
	req, err := client.UpdatePreparer(resourceGroupName, labName, name, notificationChannel)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "devtestlabs.NotificationChannelsClient", "Update", resp, "Failure responding to request")
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client NotificationChannelsClient) UpdatePreparer(resourceGroupName string, labName string, name string, notificationChannel NotificationChannelFragment) (*http.Request, error) {
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
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}", pathParameters),
		autorest.WithJSON(notificationChannel),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare(&http.Request{})
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client NotificationChannelsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req)
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client NotificationChannelsClient) UpdateResponder(resp *http.Response) (result NotificationChannel, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
