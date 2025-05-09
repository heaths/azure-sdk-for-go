//go:build go1.16
// +build go1.16

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
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// VPNServerConfigurationsClient contains the methods for the VPNServerConfigurations group.
// Don't use this type directly, use NewVPNServerConfigurationsClient() instead.
type VPNServerConfigurationsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewVPNServerConfigurationsClient creates a new instance of VPNServerConfigurationsClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewVPNServerConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *VPNServerConfigurationsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &VPNServerConfigurationsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// BeginCreateOrUpdate - Creates a VpnServerConfiguration resource if it doesn't exist else updates the existing VpnServerConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - The name of the VpnServerConfiguration being created or updated.
// vpnServerConfigurationParameters - Parameters supplied to create or update VpnServerConfiguration.
// options - VPNServerConfigurationsClientBeginCreateOrUpdateOptions contains the optional parameters for the VPNServerConfigurationsClient.BeginCreateOrUpdate
// method.
func (client *VPNServerConfigurationsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VPNServerConfiguration, options *VPNServerConfigurationsClientBeginCreateOrUpdateOptions) (VPNServerConfigurationsClientCreateOrUpdatePollerResponse, error) {
	resp, err := client.createOrUpdate(ctx, resourceGroupName, vpnServerConfigurationName, vpnServerConfigurationParameters, options)
	if err != nil {
		return VPNServerConfigurationsClientCreateOrUpdatePollerResponse{}, err
	}
	result := VPNServerConfigurationsClientCreateOrUpdatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VPNServerConfigurationsClient.CreateOrUpdate", "azure-async-operation", resp, client.pl)
	if err != nil {
		return VPNServerConfigurationsClientCreateOrUpdatePollerResponse{}, err
	}
	result.Poller = &VPNServerConfigurationsClientCreateOrUpdatePoller{
		pt: pt,
	}
	return result, nil
}

// CreateOrUpdate - Creates a VpnServerConfiguration resource if it doesn't exist else updates the existing VpnServerConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VPNServerConfigurationsClient) createOrUpdate(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VPNServerConfiguration, options *VPNServerConfigurationsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, vpnServerConfigurationParameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *VPNServerConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters VPNServerConfiguration, options *VPNServerConfigurationsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnServerConfigurationName == "" {
		return nil, errors.New("parameter vpnServerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, vpnServerConfigurationParameters)
}

// BeginDelete - Deletes a VpnServerConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - The name of the VpnServerConfiguration being deleted.
// options - VPNServerConfigurationsClientBeginDeleteOptions contains the optional parameters for the VPNServerConfigurationsClient.BeginDelete
// method.
func (client *VPNServerConfigurationsClient) BeginDelete(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsClientBeginDeleteOptions) (VPNServerConfigurationsClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, vpnServerConfigurationName, options)
	if err != nil {
		return VPNServerConfigurationsClientDeletePollerResponse{}, err
	}
	result := VPNServerConfigurationsClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("VPNServerConfigurationsClient.Delete", "location", resp, client.pl)
	if err != nil {
		return VPNServerConfigurationsClientDeletePollerResponse{}, err
	}
	result.Poller = &VPNServerConfigurationsClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes a VpnServerConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *VPNServerConfigurationsClient) deleteOperation(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, options)
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
func (client *VPNServerConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnServerConfigurationName == "" {
		return nil, errors.New("parameter vpnServerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Retrieves the details of a VpnServerConfiguration.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - The name of the VpnServerConfiguration being retrieved.
// options - VPNServerConfigurationsClientGetOptions contains the optional parameters for the VPNServerConfigurationsClient.Get
// method.
func (client *VPNServerConfigurationsClient) Get(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsClientGetOptions) (VPNServerConfigurationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, options)
	if err != nil {
		return VPNServerConfigurationsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VPNServerConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNServerConfigurationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *VPNServerConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, options *VPNServerConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnServerConfigurationName == "" {
		return nil, errors.New("parameter vpnServerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
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

// getHandleResponse handles the Get response.
func (client *VPNServerConfigurationsClient) getHandleResponse(resp *http.Response) (VPNServerConfigurationsClientGetResponse, error) {
	result := VPNServerConfigurationsClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNServerConfiguration); err != nil {
		return VPNServerConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all the VpnServerConfigurations in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - VPNServerConfigurationsClientListOptions contains the optional parameters for the VPNServerConfigurationsClient.List
// method.
func (client *VPNServerConfigurationsClient) List(options *VPNServerConfigurationsClientListOptions) *VPNServerConfigurationsClientListPager {
	return &VPNServerConfigurationsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, options)
		},
		advancer: func(ctx context.Context, resp VPNServerConfigurationsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVPNServerConfigurationsResult.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *VPNServerConfigurationsClient) listCreateRequest(ctx context.Context, options *VPNServerConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/vpnServerConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
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
func (client *VPNServerConfigurationsClient) listHandleResponse(resp *http.Response) (VPNServerConfigurationsClientListResponse, error) {
	result := VPNServerConfigurationsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNServerConfigurationsResult); err != nil {
		return VPNServerConfigurationsClientListResponse{}, err
	}
	return result, nil
}

// ListByResourceGroup - Lists all the vpnServerConfigurations in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnServerConfiguration.
// options - VPNServerConfigurationsClientListByResourceGroupOptions contains the optional parameters for the VPNServerConfigurationsClient.ListByResourceGroup
// method.
func (client *VPNServerConfigurationsClient) ListByResourceGroup(resourceGroupName string, options *VPNServerConfigurationsClientListByResourceGroupOptions) *VPNServerConfigurationsClientListByResourceGroupPager {
	return &VPNServerConfigurationsClientListByResourceGroupPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
		},
		advancer: func(ctx context.Context, resp VPNServerConfigurationsClientListByResourceGroupResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.ListVPNServerConfigurationsResult.NextLink)
		},
	}
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *VPNServerConfigurationsClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *VPNServerConfigurationsClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations"
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
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *VPNServerConfigurationsClient) listByResourceGroupHandleResponse(resp *http.Response) (VPNServerConfigurationsClientListByResourceGroupResponse, error) {
	result := VPNServerConfigurationsClientListByResourceGroupResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNServerConfigurationsResult); err != nil {
		return VPNServerConfigurationsClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates VpnServerConfiguration tags.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The resource group name of the VpnServerConfiguration.
// vpnServerConfigurationName - The name of the VpnServerConfiguration being updated.
// vpnServerConfigurationParameters - Parameters supplied to update VpnServerConfiguration tags.
// options - VPNServerConfigurationsClientUpdateTagsOptions contains the optional parameters for the VPNServerConfigurationsClient.UpdateTags
// method.
func (client *VPNServerConfigurationsClient) UpdateTags(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters TagsObject, options *VPNServerConfigurationsClientUpdateTagsOptions) (VPNServerConfigurationsClientUpdateTagsResponse, error) {
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, vpnServerConfigurationName, vpnServerConfigurationParameters, options)
	if err != nil {
		return VPNServerConfigurationsClientUpdateTagsResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return VPNServerConfigurationsClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return VPNServerConfigurationsClientUpdateTagsResponse{}, runtime.NewResponseError(resp)
	}
	return client.updateTagsHandleResponse(resp)
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *VPNServerConfigurationsClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, vpnServerConfigurationName string, vpnServerConfigurationParameters TagsObject, options *VPNServerConfigurationsClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnServerConfigurations/{vpnServerConfigurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if vpnServerConfigurationName == "" {
		return nil, errors.New("parameter vpnServerConfigurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnServerConfigurationName}", url.PathEscape(vpnServerConfigurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, vpnServerConfigurationParameters)
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *VPNServerConfigurationsClient) updateTagsHandleResponse(resp *http.Response) (VPNServerConfigurationsClientUpdateTagsResponse, error) {
	result := VPNServerConfigurationsClientUpdateTagsResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNServerConfiguration); err != nil {
		return VPNServerConfigurationsClientUpdateTagsResponse{}, err
	}
	return result, nil
}
