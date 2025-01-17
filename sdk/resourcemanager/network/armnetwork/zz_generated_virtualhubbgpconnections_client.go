//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

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
	"strings"
)

// VirtualHubBgpConnectionsClient contains the methods for the VirtualHubBgpConnections group.
// Don't use this type directly, use NewVirtualHubBgpConnectionsClient() instead.
type VirtualHubBgpConnectionsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVirtualHubBgpConnectionsClient creates a new instance of VirtualHubBgpConnectionsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVirtualHubBgpConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VirtualHubBgpConnectionsClient, error) {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := cloud.AzurePublicCloud.Services[cloud.ResourceManager].Endpoint
	if c, ok := options.Cloud.Services[cloud.ResourceManager]; ok {
		ep = c.Endpoint
	}
	pl, err := armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options)
	if err != nil {
		return nil, err
	}
	client := &VirtualHubBgpConnectionsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// NewListPager - Retrieves the details of all VirtualHubBgpConnections.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VirtualHub.
// virtualHubName - The name of the VirtualHub.
// options - VirtualHubBgpConnectionsClientListOptions contains the optional parameters for the VirtualHubBgpConnectionsClient.List
// method.
func (client *VirtualHubBgpConnectionsClient) NewListPager(resourceGroupName string, virtualHubName string, options *VirtualHubBgpConnectionsClientListOptions) *runtime.Pager[VirtualHubBgpConnectionsClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[VirtualHubBgpConnectionsClientListResponse]{
		More: func(page VirtualHubBgpConnectionsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VirtualHubBgpConnectionsClientListResponse) (VirtualHubBgpConnectionsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, virtualHubName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return VirtualHubBgpConnectionsClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return VirtualHubBgpConnectionsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return VirtualHubBgpConnectionsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *VirtualHubBgpConnectionsClient) listCreateRequest(ctx context.Context, resourceGroupName string, virtualHubName string, options *VirtualHubBgpConnectionsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{virtualHubName}/bgpConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if virtualHubName == "" {
		return nil, errors.New("parameter virtualHubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{virtualHubName}", url.PathEscape(virtualHubName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualHubBgpConnectionsClient) listHandleResponse(resp *http.Response) (VirtualHubBgpConnectionsClientListResponse, error) {
	result := VirtualHubBgpConnectionsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVirtualHubBgpConnectionResults); err != nil {
		return VirtualHubBgpConnectionsClientListResponse{}, err
	}
	return result, nil
}

// BeginListAdvertisedRoutes - Retrieves a list of routes the virtual hub bgp connection is advertising to the specified peer.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hubName - The name of the virtual hub.
// connectionName - The name of the virtual hub bgp connection.
// options - VirtualHubBgpConnectionsClientBeginListAdvertisedRoutesOptions contains the optional parameters for the VirtualHubBgpConnectionsClient.BeginListAdvertisedRoutes
// method.
func (client *VirtualHubBgpConnectionsClient) BeginListAdvertisedRoutes(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *VirtualHubBgpConnectionsClientBeginListAdvertisedRoutesOptions) (*armruntime.Poller[VirtualHubBgpConnectionsClientListAdvertisedRoutesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listAdvertisedRoutes(ctx, resourceGroupName, hubName, connectionName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[VirtualHubBgpConnectionsClientListAdvertisedRoutesResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualHubBgpConnectionsClientListAdvertisedRoutesResponse](options.ResumeToken, client.pl, nil)
	}
}

// ListAdvertisedRoutes - Retrieves a list of routes the virtual hub bgp connection is advertising to the specified peer.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualHubBgpConnectionsClient) listAdvertisedRoutes(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *VirtualHubBgpConnectionsClientBeginListAdvertisedRoutesOptions) (*http.Response, error) {
	req, err := client.listAdvertisedRoutesCreateRequest(ctx, resourceGroupName, hubName, connectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// listAdvertisedRoutesCreateRequest creates the ListAdvertisedRoutes request.
func (client *VirtualHubBgpConnectionsClient) listAdvertisedRoutesCreateRequest(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *VirtualHubBgpConnectionsClientBeginListAdvertisedRoutesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{hubName}/bgpConnections/{connectionName}/advertisedRoutes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginListLearnedRoutes - Retrieves a list of routes the virtual hub bgp connection has learned.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// hubName - The name of the virtual hub.
// connectionName - The name of the virtual hub bgp connection.
// options - VirtualHubBgpConnectionsClientBeginListLearnedRoutesOptions contains the optional parameters for the VirtualHubBgpConnectionsClient.BeginListLearnedRoutes
// method.
func (client *VirtualHubBgpConnectionsClient) BeginListLearnedRoutes(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *VirtualHubBgpConnectionsClientBeginListLearnedRoutesOptions) (*armruntime.Poller[VirtualHubBgpConnectionsClientListLearnedRoutesResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.listLearnedRoutes(ctx, resourceGroupName, hubName, connectionName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[VirtualHubBgpConnectionsClientListLearnedRoutesResponse]{
			FinalStateVia: armruntime.FinalStateViaLocation,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[VirtualHubBgpConnectionsClientListLearnedRoutesResponse](options.ResumeToken, client.pl, nil)
	}
}

// ListLearnedRoutes - Retrieves a list of routes the virtual hub bgp connection has learned.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VirtualHubBgpConnectionsClient) listLearnedRoutes(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *VirtualHubBgpConnectionsClientBeginListLearnedRoutesOptions) (*http.Response, error) {
	req, err := client.listLearnedRoutesCreateRequest(ctx, resourceGroupName, hubName, connectionName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// listLearnedRoutesCreateRequest creates the ListLearnedRoutes request.
func (client *VirtualHubBgpConnectionsClient) listLearnedRoutesCreateRequest(ctx context.Context, resourceGroupName string, hubName string, connectionName string, options *VirtualHubBgpConnectionsClientBeginListLearnedRoutesOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/virtualHubs/{hubName}/bgpConnections/{connectionName}/learnedRoutes"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if hubName == "" {
		return nil, errors.New("parameter hubName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{hubName}", url.PathEscape(hubName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
