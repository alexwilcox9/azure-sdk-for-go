package search

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/tracing"
	"github.com/gofrs/uuid"
	"net/http"
)

// SharedPrivateLinkResourcesClient is the client that can be used to manage Azure Cognitive Search services and API
// keys.
type SharedPrivateLinkResourcesClient struct {
	BaseClient
}

// NewSharedPrivateLinkResourcesClient creates an instance of the SharedPrivateLinkResourcesClient client.
func NewSharedPrivateLinkResourcesClient(subscriptionID string) SharedPrivateLinkResourcesClient {
	return NewSharedPrivateLinkResourcesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewSharedPrivateLinkResourcesClientWithBaseURI creates an instance of the SharedPrivateLinkResourcesClient client
// using a custom endpoint.  Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign
// clouds, Azure stack).
func NewSharedPrivateLinkResourcesClientWithBaseURI(baseURI string, subscriptionID string) SharedPrivateLinkResourcesClient {
	return SharedPrivateLinkResourcesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate initiates the creation or update of a shared private link resource managed by the search service in
// the given resource group.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service associated with the specified resource
// group.
// sharedPrivateLinkResourceName - the name of the shared private link resource managed by the Azure Cognitive
// Search service within the specified resource group.
// sharedPrivateLinkResource - the definition of the shared private link resource to create or update.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client SharedPrivateLinkResourcesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, sharedPrivateLinkResource SharedPrivateLinkResource, clientRequestID *uuid.UUID) (result SharedPrivateLinkResourcesCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharedPrivateLinkResourcesClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, searchServiceName, sharedPrivateLinkResourceName, sharedPrivateLinkResource, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.SharedPrivateLinkResourcesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.SharedPrivateLinkResourcesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client SharedPrivateLinkResourcesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, sharedPrivateLinkResource SharedPrivateLinkResource, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"searchServiceName":             autorest.Encode("path", searchServiceName),
		"sharedPrivateLinkResourceName": autorest.Encode("path", sharedPrivateLinkResourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}", pathParameters),
		autorest.WithJSON(sharedPrivateLinkResource),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client SharedPrivateLinkResourcesClient) CreateOrUpdateSender(req *http.Request) (future SharedPrivateLinkResourcesCreateOrUpdateFuture, err error) {
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
func (client SharedPrivateLinkResourcesClient) CreateOrUpdateResponder(resp *http.Response) (result SharedPrivateLinkResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete initiates the deletion of the shared private link resource from the search service.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service associated with the specified resource
// group.
// sharedPrivateLinkResourceName - the name of the shared private link resource managed by the Azure Cognitive
// Search service within the specified resource group.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client SharedPrivateLinkResourcesClient) Delete(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, clientRequestID *uuid.UUID) (result SharedPrivateLinkResourcesDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharedPrivateLinkResourcesClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, searchServiceName, sharedPrivateLinkResourceName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.SharedPrivateLinkResourcesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.SharedPrivateLinkResourcesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client SharedPrivateLinkResourcesClient) DeletePreparer(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"searchServiceName":             autorest.Encode("path", searchServiceName),
		"sharedPrivateLinkResourceName": autorest.Encode("path", sharedPrivateLinkResourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client SharedPrivateLinkResourcesClient) DeleteSender(req *http.Request) (future SharedPrivateLinkResourcesDeleteFuture, err error) {
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

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client SharedPrivateLinkResourcesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent, http.StatusNotFound),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the details of the shared private link resource managed by the search service in the given resource group.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service associated with the specified resource
// group.
// sharedPrivateLinkResourceName - the name of the shared private link resource managed by the Azure Cognitive
// Search service within the specified resource group.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client SharedPrivateLinkResourcesClient) Get(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, clientRequestID *uuid.UUID) (result SharedPrivateLinkResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharedPrivateLinkResourcesClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, searchServiceName, sharedPrivateLinkResourceName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.SharedPrivateLinkResourcesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.SharedPrivateLinkResourcesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.SharedPrivateLinkResourcesClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client SharedPrivateLinkResourcesClient) GetPreparer(ctx context.Context, resourceGroupName string, searchServiceName string, sharedPrivateLinkResourceName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName":             autorest.Encode("path", resourceGroupName),
		"searchServiceName":             autorest.Encode("path", searchServiceName),
		"sharedPrivateLinkResourceName": autorest.Encode("path", sharedPrivateLinkResourceName),
		"subscriptionId":                autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/sharedPrivateLinkResources/{sharedPrivateLinkResourceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client SharedPrivateLinkResourcesClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client SharedPrivateLinkResourcesClient) GetResponder(resp *http.Response) (result SharedPrivateLinkResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByService gets a list of all shared private link resources managed by the given service.
// Parameters:
// resourceGroupName - the name of the resource group within the current subscription. You can obtain this
// value from the Azure Resource Manager API or the portal.
// searchServiceName - the name of the Azure Cognitive Search service associated with the specified resource
// group.
// clientRequestID - a client-generated GUID value that identifies this request. If specified, this will be
// included in response information as a way to track the request.
func (client SharedPrivateLinkResourcesClient) ListByService(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result SharedPrivateLinkResourceListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharedPrivateLinkResourcesClient.ListByService")
		defer func() {
			sc := -1
			if result.splrlr.Response.Response != nil {
				sc = result.splrlr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listByServiceNextResults
	req, err := client.ListByServicePreparer(ctx, resourceGroupName, searchServiceName, clientRequestID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.SharedPrivateLinkResourcesClient", "ListByService", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.splrlr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "search.SharedPrivateLinkResourcesClient", "ListByService", resp, "Failure sending request")
		return
	}

	result.splrlr, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.SharedPrivateLinkResourcesClient", "ListByService", resp, "Failure responding to request")
		return
	}
	if result.splrlr.hasNextLink() && result.splrlr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByServicePreparer prepares the ListByService request.
func (client SharedPrivateLinkResourcesClient) ListByServicePreparer(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"searchServiceName": autorest.Encode("path", searchServiceName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2020-08-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Search/searchServices/{searchServiceName}/sharedPrivateLinkResources", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	if clientRequestID != nil {
		preparer = autorest.DecoratePreparer(preparer,
			autorest.WithHeader("x-ms-client-request-id", autorest.String(clientRequestID)))
	}
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByServiceSender sends the ListByService request. The method will close the
// http.Response Body if it receives an error.
func (client SharedPrivateLinkResourcesClient) ListByServiceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByServiceResponder handles the response to the ListByService request. The method always
// closes the http.Response Body.
func (client SharedPrivateLinkResourcesClient) ListByServiceResponder(resp *http.Response) (result SharedPrivateLinkResourceListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByServiceNextResults retrieves the next set of results, if any.
func (client SharedPrivateLinkResourcesClient) listByServiceNextResults(ctx context.Context, lastResults SharedPrivateLinkResourceListResult) (result SharedPrivateLinkResourceListResult, err error) {
	req, err := lastResults.sharedPrivateLinkResourceListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "search.SharedPrivateLinkResourcesClient", "listByServiceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByServiceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "search.SharedPrivateLinkResourcesClient", "listByServiceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByServiceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "search.SharedPrivateLinkResourcesClient", "listByServiceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByServiceComplete enumerates all values, automatically crossing page boundaries as required.
func (client SharedPrivateLinkResourcesClient) ListByServiceComplete(ctx context.Context, resourceGroupName string, searchServiceName string, clientRequestID *uuid.UUID) (result SharedPrivateLinkResourceListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/SharedPrivateLinkResourcesClient.ListByService")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByService(ctx, resourceGroupName, searchServiceName, clientRequestID)
	return
}
