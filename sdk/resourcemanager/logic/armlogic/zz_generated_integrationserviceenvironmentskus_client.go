//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armlogic

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

// IntegrationServiceEnvironmentSKUsClient contains the methods for the IntegrationServiceEnvironmentSKUs group.
// Don't use this type directly, use NewIntegrationServiceEnvironmentSKUsClient() instead.
type IntegrationServiceEnvironmentSKUsClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewIntegrationServiceEnvironmentSKUsClient creates a new instance of IntegrationServiceEnvironmentSKUsClient with the specified values.
// subscriptionID - The subscription id.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewIntegrationServiceEnvironmentSKUsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *IntegrationServiceEnvironmentSKUsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &IntegrationServiceEnvironmentSKUsClient{
		subscriptionID: subscriptionID,
		host:           string(cp.Endpoint),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// List - Gets a list of integration service environment Skus.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroup - The resource group.
// integrationServiceEnvironmentName - The integration service environment name.
// options - IntegrationServiceEnvironmentSKUsClientListOptions contains the optional parameters for the IntegrationServiceEnvironmentSKUsClient.List
// method.
func (client *IntegrationServiceEnvironmentSKUsClient) List(resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentSKUsClientListOptions) *IntegrationServiceEnvironmentSKUsClientListPager {
	return &IntegrationServiceEnvironmentSKUsClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroup, integrationServiceEnvironmentName, options)
		},
		advancer: func(ctx context.Context, resp IntegrationServiceEnvironmentSKUsClientListResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.IntegrationServiceEnvironmentSKUList.NextLink)
		},
	}
}

// listCreateRequest creates the List request.
func (client *IntegrationServiceEnvironmentSKUsClient) listCreateRequest(ctx context.Context, resourceGroup string, integrationServiceEnvironmentName string, options *IntegrationServiceEnvironmentSKUsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroup}/providers/Microsoft.Logic/integrationServiceEnvironments/{integrationServiceEnvironmentName}/skus"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroup == "" {
		return nil, errors.New("parameter resourceGroup cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroup}", url.PathEscape(resourceGroup))
	if integrationServiceEnvironmentName == "" {
		return nil, errors.New("parameter integrationServiceEnvironmentName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{integrationServiceEnvironmentName}", url.PathEscape(integrationServiceEnvironmentName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-05-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *IntegrationServiceEnvironmentSKUsClient) listHandleResponse(resp *http.Response) (IntegrationServiceEnvironmentSKUsClientListResponse, error) {
	result := IntegrationServiceEnvironmentSKUsClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.IntegrationServiceEnvironmentSKUList); err != nil {
		return IntegrationServiceEnvironmentSKUsClientListResponse{}, err
	}
	return result, nil
}
