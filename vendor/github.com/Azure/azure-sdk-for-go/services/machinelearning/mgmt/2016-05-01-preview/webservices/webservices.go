package webservices

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
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// Client is the these APIs allow end users to operate on Azure Machine Learning Web Services resources. They support
// the following operations:<ul><li>Create or update a web service</li><li>Get a web service</li><li>Patch a web
// service</li><li>Delete a web service</li><li>Get All Web Services in a Resource Group </li><li>Get All Web Services
// in a Subscription</li><li>Get Web Services Keys</li></ul>
type Client struct {
	BaseClient
}

// NewClient creates an instance of the Client client.
func NewClient(subscriptionID string) Client {
	return NewClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewClientWithBaseURI creates an instance of the Client client.
func NewClientWithBaseURI(baseURI string, subscriptionID string) Client {
	return Client{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create or update a web service. This call will overwrite an existing web service. Note that there is
// no warning or confirmation. This is a nonrecoverable operation. If your intent is to create a new web service, call
// the Get operation first to verify that it does not exist.
//
// resourceGroupName is name of the resource group in which the web service is located. webServiceName is the name
// of the web service. createOrUpdatePayload is the payload that is used to create or update the web service.
func (client Client) CreateOrUpdate(ctx context.Context, resourceGroupName string, webServiceName string, createOrUpdatePayload WebService) (result CreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: createOrUpdatePayload,
			Constraints: []validation.Constraint{{Target: "createOrUpdatePayload.Properties", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.RealtimeConfiguration", Name: validation.Null, Rule: false,
					Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.RealtimeConfiguration.MaxConcurrentCalls", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.RealtimeConfiguration.MaxConcurrentCalls", Name: validation.InclusiveMaximum, Rule: 200, Chain: nil},
							{Target: "createOrUpdatePayload.Properties.RealtimeConfiguration.MaxConcurrentCalls", Name: validation.InclusiveMinimum, Rule: 4, Chain: nil},
						}},
					}},
					{Target: "createOrUpdatePayload.Properties.MachineLearningWorkspace", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.MachineLearningWorkspace.ID", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "createOrUpdatePayload.Properties.CommitmentPlan", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.CommitmentPlan.ID", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "createOrUpdatePayload.Properties.Input", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.Input.Type", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "createOrUpdatePayload.Properties.Input.Properties", Name: validation.Null, Rule: true, Chain: nil},
						}},
					{Target: "createOrUpdatePayload.Properties.Output", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "createOrUpdatePayload.Properties.Output.Type", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "createOrUpdatePayload.Properties.Output.Properties", Name: validation.Null, Rule: true, Chain: nil},
						}},
				}}}}}); err != nil {
		return result, validation.NewError("webservices.Client", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, webServiceName, createOrUpdatePayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client Client) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, webServiceName string, createOrUpdatePayload WebService) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webServiceName":    autorest.Encode("path", webServiceName),
	}

	const APIVersion = "2016-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", pathParameters),
		autorest.WithJSON(createOrUpdatePayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client Client) CreateOrUpdateSender(req *http.Request) (future CreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client Client) CreateOrUpdateResponder(resp *http.Response) (result WebService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Get gets the Web Service Definiton as specified by a subscription, resource group, and name. Note that the storage
// credentials and web service keys are not returned by this call. To get the web service access keys, call List Keys.
//
// resourceGroupName is name of the resource group in which the web service is located. webServiceName is the name
// of the web service.
func (client Client) Get(ctx context.Context, resourceGroupName string, webServiceName string) (result WebService, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, webServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "webservices.Client", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client Client) GetPreparer(ctx context.Context, resourceGroupName string, webServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webServiceName":    autorest.Encode("path", webServiceName),
	}

	const APIVersion = "2016-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client Client) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client Client) GetResponder(resp *http.Response) (result WebService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets the web services in the specified subscription.
//
// skiptoken is continuation token for pagination.
func (client Client) List(ctx context.Context, skiptoken string) (result PaginatedWebServicesListPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.pwsl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "webservices.Client", "List", resp, "Failure sending request")
		return
	}

	result.pwsl, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client Client) ListPreparer(ctx context.Context, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.MachineLearning/webServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client Client) ListResponder(resp *http.Response) (result PaginatedWebServicesList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client Client) listNextResults(lastResults PaginatedWebServicesList) (result PaginatedWebServicesList, err error) {
	req, err := lastResults.paginatedWebServicesListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "webservices.Client", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "webservices.Client", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) ListComplete(ctx context.Context, skiptoken string) (result PaginatedWebServicesListIterator, err error) {
	result.page, err = client.List(ctx, skiptoken)
	return
}

// ListByResourceGroup gets the web services in the specified resource group.
//
// resourceGroupName is name of the resource group in which the web service is located. skiptoken is continuation
// token for pagination.
func (client Client) ListByResourceGroup(ctx context.Context, resourceGroupName string, skiptoken string) (result PaginatedWebServicesListPage, err error) {
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName, skiptoken)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.pwsl.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "webservices.Client", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.pwsl, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client Client) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string, skiptoken string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2016-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(skiptoken) > 0 {
		queryParameters["$skiptoken"] = autorest.Encode("query", skiptoken)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client Client) ListByResourceGroupResponder(resp *http.Response) (result PaginatedWebServicesList, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client Client) listByResourceGroupNextResults(lastResults PaginatedWebServicesList) (result PaginatedWebServicesList, err error) {
	req, err := lastResults.paginatedWebServicesListPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "webservices.Client", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "webservices.Client", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client Client) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string, skiptoken string) (result PaginatedWebServicesListIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName, skiptoken)
	return
}

// ListKeys gets the access keys for the specified web service.
//
// resourceGroupName is name of the resource group in which the web service is located. webServiceName is the name
// of the web service.
func (client Client) ListKeys(ctx context.Context, resourceGroupName string, webServiceName string) (result Keys, err error) {
	req, err := client.ListKeysPreparer(ctx, resourceGroupName, webServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "ListKeys", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListKeysSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "webservices.Client", "ListKeys", resp, "Failure sending request")
		return
	}

	result, err = client.ListKeysResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "ListKeys", resp, "Failure responding to request")
	}

	return
}

// ListKeysPreparer prepares the ListKeys request.
func (client Client) ListKeysPreparer(ctx context.Context, resourceGroupName string, webServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webServiceName":    autorest.Encode("path", webServiceName),
	}

	const APIVersion = "2016-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}/listKeys", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListKeysSender sends the ListKeys request. The method will close the
// http.Response Body if it receives an error.
func (client Client) ListKeysSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListKeysResponder handles the response to the ListKeys request. The method always
// closes the http.Response Body.
func (client Client) ListKeysResponder(resp *http.Response) (result Keys, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Patch modifies an existing web service resource. The PATCH API call is an asynchronous operation. To determine
// whether it has completed successfully, you must perform a Get operation.
//
// resourceGroupName is name of the resource group in which the web service is located. webServiceName is the name
// of the web service. patchPayload is the payload to use to patch the web service.
func (client Client) Patch(ctx context.Context, resourceGroupName string, webServiceName string, patchPayload WebService) (result PatchFuture, err error) {
	req, err := client.PatchPreparer(ctx, resourceGroupName, webServiceName, patchPayload)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "Patch", nil, "Failure preparing request")
		return
	}

	result, err = client.PatchSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "Patch", result.Response(), "Failure sending request")
		return
	}

	return
}

// PatchPreparer prepares the Patch request.
func (client Client) PatchPreparer(ctx context.Context, resourceGroupName string, webServiceName string, patchPayload WebService) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webServiceName":    autorest.Encode("path", webServiceName),
	}

	const APIVersion = "2016-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", pathParameters),
		autorest.WithJSON(patchPayload),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PatchSender sends the Patch request. The method will close the
// http.Response Body if it receives an error.
func (client Client) PatchSender(req *http.Request) (future PatchFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK))
	return
}

// PatchResponder handles the response to the Patch request. The method always
// closes the http.Response Body.
func (client Client) PatchResponder(resp *http.Response) (result WebService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Remove deletes the specified web service.
//
// resourceGroupName is name of the resource group in which the web service is located. webServiceName is the name
// of the web service.
func (client Client) Remove(ctx context.Context, resourceGroupName string, webServiceName string) (result RemoveFuture, err error) {
	req, err := client.RemovePreparer(ctx, resourceGroupName, webServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "Remove", nil, "Failure preparing request")
		return
	}

	result, err = client.RemoveSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webservices.Client", "Remove", result.Response(), "Failure sending request")
		return
	}

	return
}

// RemovePreparer prepares the Remove request.
func (client Client) RemovePreparer(ctx context.Context, resourceGroupName string, webServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"webServiceName":    autorest.Encode("path", webServiceName),
	}

	const APIVersion = "2016-05-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.MachineLearning/webServices/{webServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// RemoveSender sends the Remove request. The method will close the
// http.Response Body if it receives an error.
func (client Client) RemoveSender(req *http.Request) (future RemoveFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// RemoveResponder handles the response to the Remove request. The method always
// closes the http.Response Body.
func (client Client) RemoveResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}
