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

// NotificationChannelsClient contains the methods for the NotificationChannels group.
// Don't use this type directly, use NewNotificationChannelsClient() instead.
type NotificationChannelsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewNotificationChannelsClient creates a new instance of NotificationChannelsClient with the specified values.
// subscriptionID - The subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewNotificationChannelsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *NotificationChannelsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &NotificationChannelsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// CreateOrUpdate - Create or replace an existing notification channel.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the notification channel.
// notificationChannel - A notification.
// options - NotificationChannelsClientCreateOrUpdateOptions contains the optional parameters for the NotificationChannelsClient.CreateOrUpdate
// method.
func (client *NotificationChannelsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, labName string, name string, notificationChannel NotificationChannel, options *NotificationChannelsClientCreateOrUpdateOptions) (NotificationChannelsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, labName, name, notificationChannel, options)
	if err != nil {
		return NotificationChannelsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationChannelsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return NotificationChannelsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *NotificationChannelsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, notificationChannel NotificationChannel, options *NotificationChannelsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}"
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
	return req, runtime.MarshalAsJSON(req, notificationChannel)
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *NotificationChannelsClient) createOrUpdateHandleResponse(resp *http.Response) (NotificationChannelsClientCreateOrUpdateResponse, error) {
	result := NotificationChannelsClientCreateOrUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationChannel); err != nil {
		return NotificationChannelsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// Delete - Delete notification channel.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the notification channel.
// options - NotificationChannelsClientDeleteOptions contains the optional parameters for the NotificationChannelsClient.Delete
// method.
func (client *NotificationChannelsClient) Delete(ctx context.Context, resourceGroupName string, labName string, name string, options *NotificationChannelsClientDeleteOptions) (NotificationChannelsClientDeleteResponse, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, labName, name, options)
	if err != nil {
		return NotificationChannelsClientDeleteResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationChannelsClientDeleteResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return NotificationChannelsClientDeleteResponse{}, runtime.NewResponseError(resp)
	}
	return NotificationChannelsClientDeleteResponse{RawResponse: resp}, nil
}

// deleteCreateRequest creates the Delete request.
func (client *NotificationChannelsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *NotificationChannelsClientDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}"
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

// Get - Get notification channel.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the notification channel.
// options - NotificationChannelsClientGetOptions contains the optional parameters for the NotificationChannelsClient.Get
// method.
func (client *NotificationChannelsClient) Get(ctx context.Context, resourceGroupName string, labName string, name string, options *NotificationChannelsClientGetOptions) (NotificationChannelsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, labName, name, options)
	if err != nil {
		return NotificationChannelsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationChannelsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NotificationChannelsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *NotificationChannelsClient) getCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, options *NotificationChannelsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}"
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
func (client *NotificationChannelsClient) getHandleResponse(resp *http.Response) (NotificationChannelsClientGetResponse, error) {
	result := NotificationChannelsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationChannel); err != nil {
		return NotificationChannelsClientGetResponse{}, err
	}
	return result, nil
}

// List - List notification channels in a given lab.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// options - NotificationChannelsClientListOptions contains the optional parameters for the NotificationChannelsClient.List
// method.
func (client *NotificationChannelsClient) List(resourceGroupName string, labName string, options *NotificationChannelsClientListOptions) *NotificationChannelsClientListPager {
	return &NotificationChannelsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, labName, options)
		},
		advancer: func(ctx context.Context, resp NotificationChannelsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.NotificationChannelList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *NotificationChannelsClient) listCreateRequest(ctx context.Context, resourceGroupName string, labName string, options *NotificationChannelsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels"
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
func (client *NotificationChannelsClient) listHandleResponse(resp *http.Response) (NotificationChannelsClientListResponse, error) {
	result := NotificationChannelsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationChannelList); err != nil {
		return NotificationChannelsClientListResponse{}, err
	}
	return result, nil
}

// Notify - Send notification to provided channel.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the notification channel.
// notifyParameters - Properties for generating a Notification.
// options - NotificationChannelsClientNotifyOptions contains the optional parameters for the NotificationChannelsClient.Notify
// method.
func (client *NotificationChannelsClient) Notify(ctx context.Context, resourceGroupName string, labName string, name string, notifyParameters NotifyParameters, options *NotificationChannelsClientNotifyOptions) (NotificationChannelsClientNotifyResponse, error) {
	req, err := client.notifyCreateRequest(ctx, resourceGroupName, labName, name, notifyParameters, options)
	if err != nil {
		return NotificationChannelsClientNotifyResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationChannelsClientNotifyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NotificationChannelsClientNotifyResponse{}, runtime.NewResponseError(resp)
	}
	return NotificationChannelsClientNotifyResponse{RawResponse: resp}, nil
}

// notifyCreateRequest creates the Notify request.
func (client *NotificationChannelsClient) notifyCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, notifyParameters NotifyParameters, options *NotificationChannelsClientNotifyOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}/notify"
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
	return req, runtime.MarshalAsJSON(req, notifyParameters)
}

// Update - Allows modifying tags of notification channels. All other properties will be ignored.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// labName - The name of the lab.
// name - The name of the notification channel.
// notificationChannel - A notification.
// options - NotificationChannelsClientUpdateOptions contains the optional parameters for the NotificationChannelsClient.Update
// method.
func (client *NotificationChannelsClient) Update(ctx context.Context, resourceGroupName string, labName string, name string, notificationChannel NotificationChannelFragment, options *NotificationChannelsClientUpdateOptions) (NotificationChannelsClientUpdateResponse, error) {
	req, err := client.updateCreateRequest(ctx, resourceGroupName, labName, name, notificationChannel, options)
	if err != nil {
		return NotificationChannelsClientUpdateResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return NotificationChannelsClientUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return NotificationChannelsClientUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateHandleResponse(resp)
}

// updateCreateRequest creates the Update request.
func (client *NotificationChannelsClient) updateCreateRequest(ctx context.Context, resourceGroupName string, labName string, name string, notificationChannel NotificationChannelFragment, options *NotificationChannelsClientUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DevTestLab/labs/{labName}/notificationchannels/{name}"
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
	return req, runtime.MarshalAsJSON(req, notificationChannel)
}

// updateHandleResponse handles the Update response.
func (client *NotificationChannelsClient) updateHandleResponse(resp *http.Response) (NotificationChannelsClientUpdateResponse, error) {
	result := NotificationChannelsClientUpdateResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.NotificationChannel); err != nil {
		return NotificationChannelsClientUpdateResponse{}, err
	}
	return result, nil
}
