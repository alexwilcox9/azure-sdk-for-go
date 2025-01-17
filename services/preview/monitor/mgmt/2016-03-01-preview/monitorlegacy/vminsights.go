package monitorlegacy

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
	"net/http"
)

// VMInsightsClient is the monitor Management Client
type VMInsightsClient struct {
	BaseClient
}

// NewVMInsightsClient creates an instance of the VMInsightsClient client.
func NewVMInsightsClient(subscriptionID string) VMInsightsClient {
	return NewVMInsightsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewVMInsightsClientWithBaseURI creates an instance of the VMInsightsClient client using a custom endpoint.  Use this
// when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewVMInsightsClientWithBaseURI(baseURI string, subscriptionID string) VMInsightsClient {
	return VMInsightsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// GetOnboardingStatus retrieves the VM Insights onboarding status for the specified resource or resource scope.
// Parameters:
// resourceURI - the fully qualified Azure Resource manager identifier of the resource, or scope, whose status
// to retrieve.
func (client VMInsightsClient) GetOnboardingStatus(ctx context.Context, resourceURI string) (result VMInsightsOnboardingStatus, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/VMInsightsClient.GetOnboardingStatus")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetOnboardingStatusPreparer(ctx, resourceURI)
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitorlegacy.VMInsightsClient", "GetOnboardingStatus", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetOnboardingStatusSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "monitorlegacy.VMInsightsClient", "GetOnboardingStatus", resp, "Failure sending request")
		return
	}

	result, err = client.GetOnboardingStatusResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "monitorlegacy.VMInsightsClient", "GetOnboardingStatus", resp, "Failure responding to request")
		return
	}

	return
}

// GetOnboardingStatusPreparer prepares the GetOnboardingStatus request.
func (client VMInsightsClient) GetOnboardingStatusPreparer(ctx context.Context, resourceURI string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceUri": resourceURI,
	}

	const APIVersion = "2018-11-27-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{resourceUri}/providers/Microsoft.Insights/vmInsightsOnboardingStatuses/default", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetOnboardingStatusSender sends the GetOnboardingStatus request. The method will close the
// http.Response Body if it receives an error.
func (client VMInsightsClient) GetOnboardingStatusSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetOnboardingStatusResponder handles the response to the GetOnboardingStatus request. The method always
// closes the http.Response Body.
func (client VMInsightsClient) GetOnboardingStatusResponder(resp *http.Response) (result VMInsightsOnboardingStatus, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
