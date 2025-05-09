//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdevtestlabs

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
	"strconv"
	"strings"
)

// SchedulesClient contains the methods for the Schedules group.
// Don't use this type directly, use NewSchedulesClient() instead.
type SchedulesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSchedulesClient creates a new instance of SchedulesClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewSchedulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *SchedulesClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &SchedulesClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Create or replace an existing schedule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the schedule.
// schedule - A schedule.
// options - SchedulesClientCreateOrUpdateOptions contains the optional parameters for the SchedulesClient.CreateOrUpdate
// method.
func (client *SchedulesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, schedule Schedule, options *SchedulesClientCreateOrUpdateOptions) (SchedulesClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, name, schedule, options)
	if err != nil {
		return SchedulesClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchedulesClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return SchedulesClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *SchedulesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, schedule Schedule, options *SchedulesClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, schedule)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *SchedulesClient) createOrUpdateHandleResponse(resp *http.Response) (SchedulesClientCreateOrUpdateResponse, error) {
	result := SchedulesClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return SchedulesClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete schedule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the schedule.
// options - SchedulesClientDeleteOptions contains the optional parameters for the SchedulesClient.Delete method.
func (client *SchedulesClient) Delete(ctx context.Context, resourceGroupName string, labName string, name string, options *SchedulesClientDeleteOptions) (SchedulesClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, name, options)
	if err != nil {
		return SchedulesClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchedulesClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return SchedulesClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return SchedulesClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *SchedulesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *SchedulesClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// BeginExecute - Execute a schedule. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the schedule.
// options - SchedulesClientBeginExecuteOptions contains the optional parameters for the SchedulesClient.BeginExecute method.
func (client *SchedulesClient) BeginExecute(ctx context.Context, resourceGroupName string, labName string, name string, options *SchedulesClientBeginExecuteOptions) (SchedulesClientExecutePollerResponse, error) {
	resp, err := client.execute(ctx, resourceGroupName, labName, name, options)
	if err != nil {
		return SchedulesClientExecutePollerResponse{}, err
	}
	result := SchedulesClientExecutePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("SchedulesClient.Execute", "", resp, client.pl)
	if err != nil {
		return SchedulesClientExecutePollerResponse{}, err
	}
	result.Poller = &SchedulesClientExecutePoller{
		pt: pt,
	}
	return result, nil
}

// Execute - Execute a schedule. This operation can take a while to complete.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *SchedulesClient) execute(ctx context.Context, resourceGroupName string, labName string, name string, options *SchedulesClientBeginExecuteOptions) (*http.Response, error) {
	req, err := client.executeCreateRequest(ctx, resourceGroupName, labName, name, options)
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

// executeCreateRequest creates the Execute request.
func (client *SchedulesClient) executeCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *SchedulesClientBeginExecuteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}/execute"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Get schedule.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the schedule.
// options - SchedulesClientGetOptions contains the optional parameters for the SchedulesClient.Get method.
func (client *SchedulesClient) Get(ctx context.Context, resourceGroupName string, labName string, name string, options *SchedulesClientGetOptions) (SchedulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, name, options)
	if err != nil {
		return SchedulesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchedulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SchedulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *SchedulesClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *SchedulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *SchedulesClient) getHandleResponse(resp *http.Response) (SchedulesClientGetResponse, error) {
	result := SchedulesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return SchedulesClientGetResponse{}, err
	}
	return result, nil
}

// List - List schedules in a given lab.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// options - SchedulesClientListOptions contains the optional parameters for the SchedulesClient.List method.
func (client *SchedulesClient) List(resourceGroupName string, labName string, options *SchedulesClientListOptions) *SchedulesClientListPager {
	return &SchedulesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, labName, options)
		},
		advancer: func(ctx context.Context, resp SchedulesClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ScheduleList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *SchedulesClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *SchedulesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.Orderby != nil {
		reqQP.Set("$orderby", *options.Orderby)
	}
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *SchedulesClient) listHandleResponse(resp *http.Response) (SchedulesClientListResponse, error) {
	result := SchedulesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScheduleList); err != nil {
		return SchedulesClientListResponse{}, err
	}
	return result, nil
}

// ListApplicable - Lists all applicable schedules
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the schedule.
// options - SchedulesClientListApplicableOptions contains the optional parameters for the SchedulesClient.ListApplicable
// method.
func (client *SchedulesClient) ListApplicable(resourceGroupName string, labName string, name string, options *SchedulesClientListApplicableOptions) *SchedulesClientListApplicablePager {
	return &SchedulesClientListApplicablePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listApplicableCreateRequest(ctx, resourceGroupName, labName, name, options)
		},
		advancer: func(ctx context.Context, resp SchedulesClientListApplicableResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ScheduleList.NextLink)
		},
	}
}

// listApplicableCreateRequest creates the ListApplicable request.
func (client *SchedulesClient) listApplicableCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *SchedulesClientListApplicableOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}/listApplicable"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listApplicableHandleResponse handles the ListApplicable response.
func (client *SchedulesClient) listApplicableHandleResponse(resp *http.Response) (SchedulesClientListApplicableResponse, error) {
	result := SchedulesClientListApplicableResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ScheduleList); err != nil {
		return SchedulesClientListApplicableResponse{}, err
	}
	return result, nil
}

// Update - Allows modifying tags of schedules. All other properties will be ignored.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the schedule.
// schedule - A schedule.
// options - SchedulesClientUpdateOptions contains the optional parameters for the SchedulesClient.Update method.
func (client *SchedulesClient) Update(ctx context.Context, resourceGroupName string, labName string, name string, schedule ScheduleFragment, options *SchedulesClientUpdateOptions) (SchedulesClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, name, schedule, options)
	if err != nil {
		return SchedulesClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SchedulesClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SchedulesClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *SchedulesClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, schedule ScheduleFragment, options *SchedulesClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/schedules/{name}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if labName == "" {
		return nil, errors.New("parameter labName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{labName}", url.PathEscape(labName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2018-09-15")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, schedule)
}

// updateHandleResponse handles the Update response.
func (client *SchedulesClient) updateHandleResponse(resp *http.Response) (SchedulesClientUpdateResponse, error) {
	result := SchedulesClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Schedule); err != nil {
		return SchedulesClientUpdateResponse{}, err
	}
	return result, nil
}
