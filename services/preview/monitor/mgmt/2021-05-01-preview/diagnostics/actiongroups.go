package diagnostics

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

// ActionGroupsClient is the monitor Management Client
type ActionGroupsClient struct {
	BaseClient
}

// NewActionGroupsClient creates an instance of the ActionGroupsClient client.
func NewActionGroupsClient(subscriptionID string) ActionGroupsClient {
	return NewActionGroupsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewActionGroupsClientWithBaseURI creates an instance of the ActionGroupsClient client using a custom endpoint.  Use
// this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewActionGroupsClientWithBaseURI(baseURI string, subscriptionID string) ActionGroupsClient {
	return ActionGroupsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate create a new action group or update an existing one.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// actionGroupName - the name of the action group.
// actionGroup - the action group to create or use for the update.
func (client ActionGroupsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroup ActionGroupResource) (result ActionGroupResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionGroupsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: actionGroup,
			Constraints: []validation.Constraint{{Target: "actionGroup.ActionGroup", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "actionGroup.ActionGroup.GroupShortName", Name: validation.Null, Rule: true,
					Chain: []validation.Constraint{{Target: "actionGroup.ActionGroup.GroupShortName", Name: validation.MaxLength, Rule: 12, Chain: nil}}},
					{Target: "actionGroup.ActionGroup.Enabled", Name: validation.Null, Rule: true, Chain: nil},
				}}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("diagnostics.ActionGroupsClient", "CreateOrUpdate", err.Error())
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, actionGroupName, actionGroup)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	resp, err := client.CreateOrUpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "CreateOrUpdate", resp, "Failure sending request")
		return
	}

	result, err = client.CreateOrUpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "CreateOrUpdate", resp, "Failure responding to request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ActionGroupsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroup ActionGroupResource) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionGroupName":   autorest.Encode("path", actionGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/actionGroups/{actionGroupName}", pathParameters),
		autorest.WithJSON(actionGroup),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ActionGroupsClient) CreateOrUpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ActionGroupsClient) CreateOrUpdateResponder(resp *http.Response) (result ActionGroupResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete delete an action group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// actionGroupName - the name of the action group.
func (client ActionGroupsClient) Delete(ctx context.Context, resourceGroupName string, actionGroupName string) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionGroupsClient.Delete")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("diagnostics.ActionGroupsClient", "Delete", err.Error())
	}

	req, err := client.DeletePreparer(ctx, resourceGroupName, actionGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "Delete", nil, "Failure preparing request")
		return
	}

	resp, err := client.DeleteSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "Delete", resp, "Failure sending request")
		return
	}

	result, err = client.DeleteResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "Delete", resp, "Failure responding to request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ActionGroupsClient) DeletePreparer(ctx context.Context, resourceGroupName string, actionGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionGroupName":   autorest.Encode("path", actionGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/actionGroups/{actionGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ActionGroupsClient) DeleteSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ActionGroupsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// EnableReceiver enable a receiver in an action group. This changes the receiver's status from Disabled to Enabled.
// This operation is only supported for Email or SMS receivers.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// actionGroupName - the name of the action group.
// enableRequest - the receiver to re-enable.
func (client ActionGroupsClient) EnableReceiver(ctx context.Context, resourceGroupName string, actionGroupName string, enableRequest EnableRequest) (result autorest.Response, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionGroupsClient.EnableReceiver")
		defer func() {
			sc := -1
			if result.Response != nil {
				sc = result.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: enableRequest,
			Constraints: []validation.Constraint{{Target: "enableRequest.ReceiverName", Name: validation.Null, Rule: true, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("diagnostics.ActionGroupsClient", "EnableReceiver", err.Error())
	}

	req, err := client.EnableReceiverPreparer(ctx, resourceGroupName, actionGroupName, enableRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "EnableReceiver", nil, "Failure preparing request")
		return
	}

	resp, err := client.EnableReceiverSender(req)
	if err != nil {
		result.Response = resp
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "EnableReceiver", resp, "Failure sending request")
		return
	}

	result, err = client.EnableReceiverResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "EnableReceiver", resp, "Failure responding to request")
		return
	}

	return
}

// EnableReceiverPreparer prepares the EnableReceiver request.
func (client ActionGroupsClient) EnableReceiverPreparer(ctx context.Context, resourceGroupName string, actionGroupName string, enableRequest EnableRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionGroupName":   autorest.Encode("path", actionGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/actionGroups/{actionGroupName}/subscribe", pathParameters),
		autorest.WithJSON(enableRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// EnableReceiverSender sends the EnableReceiver request. The method will close the
// http.Response Body if it receives an error.
func (client ActionGroupsClient) EnableReceiverSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// EnableReceiverResponder handles the response to the EnableReceiver request. The method always
// closes the http.Response Body.
func (client ActionGroupsClient) EnableReceiverResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get get an action group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// actionGroupName - the name of the action group.
func (client ActionGroupsClient) Get(ctx context.Context, resourceGroupName string, actionGroupName string) (result ActionGroupResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionGroupsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("diagnostics.ActionGroupsClient", "Get", err.Error())
	}

	req, err := client.GetPreparer(ctx, resourceGroupName, actionGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client ActionGroupsClient) GetPreparer(ctx context.Context, resourceGroupName string, actionGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionGroupName":   autorest.Encode("path", actionGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/actionGroups/{actionGroupName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ActionGroupsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ActionGroupsClient) GetResponder(resp *http.Response) (result ActionGroupResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetTestNotifications get the test notifications by the notification id
// Parameters:
// notificationID - the notification id
func (client ActionGroupsClient) GetTestNotifications(ctx context.Context, notificationID string) (result TestNotificationDetailsResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionGroupsClient.GetTestNotifications")
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
		return result, validation.NewError("diagnostics.ActionGroupsClient", "GetTestNotifications", err.Error())
	}

	req, err := client.GetTestNotificationsPreparer(ctx, notificationID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "GetTestNotifications", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetTestNotificationsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "GetTestNotifications", resp, "Failure sending request")
		return
	}

	result, err = client.GetTestNotificationsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "GetTestNotifications", resp, "Failure responding to request")
		return
	}

	return
}

// GetTestNotificationsPreparer prepares the GetTestNotifications request.
func (client ActionGroupsClient) GetTestNotificationsPreparer(ctx context.Context, notificationID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"notificationId": autorest.Encode("path", notificationID),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Insights/notificationStatus/{notificationId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetTestNotificationsSender sends the GetTestNotifications request. The method will close the
// http.Response Body if it receives an error.
func (client ActionGroupsClient) GetTestNotificationsSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetTestNotificationsResponder handles the response to the GetTestNotifications request. The method always
// closes the http.Response Body.
func (client ActionGroupsClient) GetTestNotificationsResponder(resp *http.Response) (result TestNotificationDetailsResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByResourceGroup get a list of all action groups in a resource group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
func (client ActionGroupsClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ActionGroupList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionGroupsClient.ListByResourceGroup")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("diagnostics.ActionGroupsClient", "ListByResourceGroup", err.Error())
	}

	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "ListByResourceGroup", resp, "Failure responding to request")
		return
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ActionGroupsClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/actionGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ActionGroupsClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ActionGroupsClient) ListByResourceGroupResponder(resp *http.Response) (result ActionGroupList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListBySubscriptionID get a list of all action groups in a subscription.
func (client ActionGroupsClient) ListBySubscriptionID(ctx context.Context) (result ActionGroupList, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionGroupsClient.ListBySubscriptionID")
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
		return result, validation.NewError("diagnostics.ActionGroupsClient", "ListBySubscriptionID", err.Error())
	}

	req, err := client.ListBySubscriptionIDPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "ListBySubscriptionID", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListBySubscriptionIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "ListBySubscriptionID", resp, "Failure sending request")
		return
	}

	result, err = client.ListBySubscriptionIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "ListBySubscriptionID", resp, "Failure responding to request")
		return
	}

	return
}

// ListBySubscriptionIDPreparer prepares the ListBySubscriptionID request.
func (client ActionGroupsClient) ListBySubscriptionIDPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Insights/actionGroups", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListBySubscriptionIDSender sends the ListBySubscriptionID request. The method will close the
// http.Response Body if it receives an error.
func (client ActionGroupsClient) ListBySubscriptionIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListBySubscriptionIDResponder handles the response to the ListBySubscriptionID request. The method always
// closes the http.Response Body.
func (client ActionGroupsClient) ListBySubscriptionIDResponder(resp *http.Response) (result ActionGroupList, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// PostTestNotifications send test notifications to a set of provided receivers
// Parameters:
// notificationRequest - the notification request body which includes the contact details
func (client ActionGroupsClient) PostTestNotifications(ctx context.Context, notificationRequest NotificationRequestBody) (result ActionGroupsPostTestNotificationsFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionGroupsClient.PostTestNotifications")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: notificationRequest,
			Constraints: []validation.Constraint{{Target: "notificationRequest.AlertType", Name: validation.Null, Rule: true,
				Chain: []validation.Constraint{{Target: "notificationRequest.AlertType", Name: validation.MaxLength, Rule: 30, Chain: nil}}}}}}); err != nil {
		return result, validation.NewError("diagnostics.ActionGroupsClient", "PostTestNotifications", err.Error())
	}

	req, err := client.PostTestNotificationsPreparer(ctx, notificationRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "PostTestNotifications", nil, "Failure preparing request")
		return
	}

	result, err = client.PostTestNotificationsSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "PostTestNotifications", result.Response(), "Failure sending request")
		return
	}

	return
}

// PostTestNotificationsPreparer prepares the PostTestNotifications request.
func (client ActionGroupsClient) PostTestNotificationsPreparer(ctx context.Context, notificationRequest NotificationRequestBody) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Insights/createNotifications", pathParameters),
		autorest.WithJSON(notificationRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// PostTestNotificationsSender sends the PostTestNotifications request. The method will close the
// http.Response Body if it receives an error.
func (client ActionGroupsClient) PostTestNotificationsSender(req *http.Request) (future ActionGroupsPostTestNotificationsFuture, err error) {
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

// PostTestNotificationsResponder handles the response to the PostTestNotifications request. The method always
// closes the http.Response Body.
func (client ActionGroupsClient) PostTestNotificationsResponder(resp *http.Response) (result TestNotificationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates an existing action group's tags. To update other fields use the CreateOrUpdate method.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// actionGroupName - the name of the action group.
// actionGroupPatch - parameters supplied to the operation.
func (client ActionGroupsClient) Update(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroupPatch ActionGroupPatchBody) (result ActionGroupResource, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/ActionGroupsClient.Update")
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
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("diagnostics.ActionGroupsClient", "Update", err.Error())
	}

	req, err := client.UpdatePreparer(ctx, resourceGroupName, actionGroupName, actionGroupPatch)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "Update", nil, "Failure preparing request")
		return
	}

	resp, err := client.UpdateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "Update", resp, "Failure sending request")
		return
	}

	result, err = client.UpdateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "diagnostics.ActionGroupsClient", "Update", resp, "Failure responding to request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client ActionGroupsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, actionGroupName string, actionGroupPatch ActionGroupPatchBody) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"actionGroupName":   autorest.Encode("path", actionGroupName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-09-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Insights/actionGroups/{actionGroupName}", pathParameters),
		autorest.WithJSON(actionGroupPatch),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client ActionGroupsClient) UpdateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client ActionGroupsClient) UpdateResponder(resp *http.Response) (result ActionGroupResource, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
