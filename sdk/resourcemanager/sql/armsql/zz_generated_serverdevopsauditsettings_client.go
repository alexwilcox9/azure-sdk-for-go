//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armsql

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

// ServerDevOpsAuditSettingsClient contains the methods for the ServerDevOpsAuditSettings group.
// Don't use this type directly, use NewServerDevOpsAuditSettingsClient() instead.
type ServerDevOpsAuditSettingsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewServerDevOpsAuditSettingsClient creates a new instance of ServerDevOpsAuditSettingsClient with the specified values.
// subscriptionID - The subscription ID that identifies an Azure subscription.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewServerDevOpsAuditSettingsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ServerDevOpsAuditSettingsClient, error) {
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
	client := &ServerDevOpsAuditSettingsClient{
		subscriptionID: subscriptionID,
		host:           ep,
		pl:             pl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a server's DevOps audit settings.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// devOpsAuditingSettingsName - The name of the devops audit settings. This should always be 'default'.
// parameters - Properties of DevOps audit settings
// options - ServerDevOpsAuditSettingsClientBeginCreateOrUpdateOptions contains the optional parameters for the ServerDevOpsAuditSettingsClient.BeginCreateOrUpdate
// method.
func (client *ServerDevOpsAuditSettingsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, parameters ServerDevOpsAuditingSettings, options *ServerDevOpsAuditSettingsClientBeginCreateOrUpdateOptions) (*armruntime.Poller[ServerDevOpsAuditSettingsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, serverName, devOpsAuditingSettingsName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller(resp, client.pl, &armruntime.NewPollerOptions[ServerDevOpsAuditSettingsClientCreateOrUpdateResponse]{
			FinalStateVia: armruntime.FinalStateViaAzureAsyncOp,
		})
	} else {
		return armruntime.NewPollerFromResumeToken[ServerDevOpsAuditSettingsClientCreateOrUpdateResponse](options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates a server's DevOps audit settings.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *ServerDevOpsAuditSettingsClient) createOrUpdate(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, parameters ServerDevOpsAuditingSettings, options *ServerDevOpsAuditSettingsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, serverName, devOpsAuditingSettingsName, parameters, options)
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

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ServerDevOpsAuditSettingsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, parameters ServerDevOpsAuditingSettings, options *ServerDevOpsAuditSettingsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/devOpsAuditingSettings/{devOpsAuditingSettingsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if devOpsAuditingSettingsName == "" {
		return nil, errors.New("parameter devOpsAuditingSettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devOpsAuditingSettingsName}", url.PathEscape(devOpsAuditingSettingsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// Get - Gets a server's DevOps audit settings.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// devOpsAuditingSettingsName - The name of the devops audit settings. This should always be 'default'.
// options - ServerDevOpsAuditSettingsClientGetOptions contains the optional parameters for the ServerDevOpsAuditSettingsClient.Get
// method.
func (client *ServerDevOpsAuditSettingsClient) Get(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, options *ServerDevOpsAuditSettingsClientGetOptions) (ServerDevOpsAuditSettingsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, serverName, devOpsAuditingSettingsName, options)
	if err != nil {
		return ServerDevOpsAuditSettingsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return ServerDevOpsAuditSettingsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ServerDevOpsAuditSettingsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ServerDevOpsAuditSettingsClient) getCreateRequest(ctx context.Context, resourceGroupName string, serverName string, devOpsAuditingSettingsName string, options *ServerDevOpsAuditSettingsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/devOpsAuditingSettings/{devOpsAuditingSettingsName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if devOpsAuditingSettingsName == "" {
		return nil, errors.New("parameter devOpsAuditingSettingsName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{devOpsAuditingSettingsName}", url.PathEscape(devOpsAuditingSettingsName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *ServerDevOpsAuditSettingsClient) getHandleResponse(resp *http.Response) (ServerDevOpsAuditSettingsClientGetResponse, error) {
	result := ServerDevOpsAuditSettingsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerDevOpsAuditingSettings); err != nil {
		return ServerDevOpsAuditSettingsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByServerPager - Lists DevOps audit settings of a server.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the resource. You can obtain this value from the Azure
// Resource Manager API or the portal.
// serverName - The name of the server.
// options - ServerDevOpsAuditSettingsClientListByServerOptions contains the optional parameters for the ServerDevOpsAuditSettingsClient.ListByServer
// method.
func (client *ServerDevOpsAuditSettingsClient) NewListByServerPager(resourceGroupName string, serverName string, options *ServerDevOpsAuditSettingsClientListByServerOptions) *runtime.Pager[ServerDevOpsAuditSettingsClientListByServerResponse] {
	return runtime.NewPager(runtime.PageProcessor[ServerDevOpsAuditSettingsClientListByServerResponse]{
		More: func(page ServerDevOpsAuditSettingsClientListByServerResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ServerDevOpsAuditSettingsClientListByServerResponse) (ServerDevOpsAuditSettingsClientListByServerResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByServerCreateRequest(ctx, resourceGroupName, serverName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ServerDevOpsAuditSettingsClientListByServerResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return ServerDevOpsAuditSettingsClientListByServerResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ServerDevOpsAuditSettingsClientListByServerResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByServerHandleResponse(resp)
		},
	})
}

// listByServerCreateRequest creates the ListByServer request.
func (client *ServerDevOpsAuditSettingsClient) listByServerCreateRequest(ctx context.Context, resourceGroupName string, serverName string, options *ServerDevOpsAuditSettingsClientListByServerOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/devOpsAuditingSettings"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if serverName == "" {
		return nil, errors.New("parameter serverName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{serverName}", url.PathEscape(serverName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-11-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByServerHandleResponse handles the ListByServer response.
func (client *ServerDevOpsAuditSettingsClient) listByServerHandleResponse(resp *http.Response) (ServerDevOpsAuditSettingsClientListByServerResponse, error) {
	result := ServerDevOpsAuditSettingsClientListByServerResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ServerDevOpsAuditSettingsListResult); err != nil {
		return ServerDevOpsAuditSettingsClientListByServerResponse{}, err
	}
	return result, nil
}
