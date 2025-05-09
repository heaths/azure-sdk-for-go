//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armhealthcareapis

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// WorkspacesClient contains the methods for the Workspaces group.
// Don't use this type directly, use NewWorkspacesClient() instead.
type WorkspacesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewWorkspacesClient creates a new instance of WorkspacesClient with the specified values.
// subscriptionID - The subscription identifier.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewWorkspacesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *WorkspacesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &WorkspacesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates a workspace resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// workspace - The parameters for creating or updating a healthcare workspace.
// options - WorkspacesClientBeginCreateOrUpdateOptions contains the optional parameters for the WorkspacesClient.BeginCreateOrUpdate
// method.
func (client *WorkspacesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspace Workspace, options *WorkspacesClientBeginCreateOrUpdateOptions) (WorkspacesClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, workspaceName, workspace, options)
	if err != nil {
		return WorkspacesClientCreateOrUpdatePollerResponse{}, err
	}
	result := WorkspacesClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspacesClient.CreateOrUpdate", "", resp, client.pl)
	if err != nil {
		return WorkspacesClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &WorkspacesClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates or updates a workspace resource with the specified parameters.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkspacesClient) createOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspace Workspace, options *WorkspacesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, workspaceName, workspace, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WorkspacesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspace Workspace, options *WorkspacesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, workspace)
}

// BeginDelete - Deletes a specified workspace.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// options - WorkspacesClientBeginDeleteOptions contains the optional parameters for the WorkspacesClient.BeginDelete method.
func (client *WorkspacesClient) BeginDelete(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientBeginDeleteOptions) (WorkspacesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return WorkspacesClientDeletePollerResponse{}, err
	}
	result := WorkspacesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspacesClient.Delete", "", resp, client.pl)
	if err != nil {
		return WorkspacesClientDeletePollerResponse{}, err
	}
	result.Poller = &WorkspacesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a specified workspace.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkspacesClient) deleteOperation(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *WorkspacesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the properties of the specified workspace.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// options - WorkspacesClientGetOptions contains the optional parameters for the WorkspacesClient.Get method.
func (client *WorkspacesClient) Get(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientGetOptions) (WorkspacesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, workspaceName, options)
	if err != nil {
		return WorkspacesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return WorkspacesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return WorkspacesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *WorkspacesClient) getCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, options *WorkspacesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *WorkspacesClient) getHandleResponse(resp *http.Response) (WorkspacesClientGetResponse, error) {
	result := WorkspacesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Workspace); err != nil {
		return WorkspacesClientGetResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists all the available workspaces under the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// options - WorkspacesClientListByResourceGroupOptions contains the optional parameters for the WorkspacesClient.ListByResourceGroup
// method.
func (client *WorkspacesClient) ListByResourceGroup(resourceGroupName string, options *WorkspacesClientListByResourceGroupOptions) *WorkspacesClientListByResourceGroupPager {
	return &WorkspacesClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp WorkspacesClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WorkspaceList.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *WorkspacesClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *WorkspacesClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *WorkspacesClient) listByResourceGroupHandleResponse(resp *http.Response) (WorkspacesClientListByResourceGroupResponse, error) {
	result := WorkspacesClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceList); err != nil {
		return WorkspacesClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// ListBySubscription - Lists all the available workspaces under the specified subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - WorkspacesClientListBySubscriptionOptions contains the optional parameters for the WorkspacesClient.ListBySubscription
// method.
func (client *WorkspacesClient) ListBySubscription(options *WorkspacesClientListBySubscriptionOptions) *WorkspacesClientListBySubscriptionPager {
	return &WorkspacesClientListBySubscriptionPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listBySubscriptionCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp WorkspacesClientListBySubscriptionResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.WorkspaceList.NextLink)
		},
	}
}

// listBySubscriptionCreateRequest creates the ListBySubscription request.
func (client *WorkspacesClient) listBySubscriptionCreateRequest(ctx context.Context, options *WorkspacesClientListBySubscriptionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.HealthcareApis/workspaces"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listBySubscriptionHandleResponse handles the ListBySubscription response.
func (client *WorkspacesClient) listBySubscriptionHandleResponse(resp *http.Response) (WorkspacesClientListBySubscriptionResponse, error) {
	result := WorkspacesClientListBySubscriptionResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.WorkspaceList); err != nil {
		return WorkspacesClientListBySubscriptionResponse{}, err
	}
	return result, nil
}

// BeginUpdate - Patch workspace details.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group that contains the service instance.
// workspaceName - The name of workspace resource.
// workspacePatchResource - The parameters for updating a specified workspace.
// options - WorkspacesClientBeginUpdateOptions contains the optional parameters for the WorkspacesClient.BeginUpdate method.
func (client *WorkspacesClient) BeginUpdate(ctx context.Context, resourceGroupName string, workspaceName string, workspacePatchResource WorkspacePatchResource, options *WorkspacesClientBeginUpdateOptions) (WorkspacesClientUpdatePollerResponse, error) {
	resp, err := client.update(ctx, resourceGroupName, workspaceName, workspacePatchResource, options)
	if err != nil {
		return WorkspacesClientUpdatePollerResponse{}, err
	}
	result := WorkspacesClientUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("WorkspacesClient.Update", "", resp, client.pl)
	if err != nil {
		return WorkspacesClientUpdatePollerResponse{}, err
	}
	result.Poller = &WorkspacesClientUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// Update - Patch workspace details.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *WorkspacesClient) update(ctx context.Context, resourceGroupName string, workspaceName string, workspacePatchResource WorkspacePatchResource, options *WorkspacesClientBeginUpdateOptions) (*http.Response, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, workspaceName, workspacePatchResource, options)
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

// updateCreateRequest creates the Update request.
func (client *WorkspacesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, workspaceName string, workspacePatchResource WorkspacePatchResource, options *WorkspacesClientBeginUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.HealthcareApis/workspaces/{workspaceName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if workspaceName == "" {
		return nil, errors.New("parameter workspaceName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{workspaceName}", url.PathEscape(workspaceName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-06-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, workspacePatchResource)
}
