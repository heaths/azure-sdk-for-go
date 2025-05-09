//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

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

// EventsClient contains the methods for the Events group.
// Don't use this type directly, use NewEventsClient() instead.
type EventsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewEventsClient creates a new instance of EventsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewEventsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *EventsClient {
	cp := arm.ClientOptions{}
	if options != nil {
		cp = *options
	}
	if len(cp.Endpoint) == 0 {
		cp.Endpoint = arm.AzurePublicCloud
	}
	client := &EventsClient{
		host: string(cp.Endpoint),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, &cp),
	}
	return client
}

// ListByBillingAccount - Lists the events that decrements Azure credits or Microsoft Azure consumption commitment for a billing
// account or a billing profile for a given start and end date.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountID - BillingAccount ID
// options - EventsClientListByBillingAccountOptions contains the optional parameters for the EventsClient.ListByBillingAccount
// method.
func (client *EventsClient) ListByBillingAccount(billingAccountID string, options *EventsClientListByBillingAccountOptions) *EventsClientListByBillingAccountPager {
	return &EventsClientListByBillingAccountPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByBillingAccountCreateRequest(ctx, billingAccountID, options)
		},
		advancer: func(ctx context.Context, resp EventsClientListByBillingAccountResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.Events.NextLink)
		},
	}
}

// listByBillingAccountCreateRequest creates the ListByBillingAccount request.
func (client *EventsClient) listByBillingAccountCreateRequest(ctx context.Context, billingAccountID string, options *EventsClientListByBillingAccountOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/providers/Microsoft.Consumption/events"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	if options != nil && options.Filter != nil {
		reqQP.Set("$filter", *options.Filter)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBillingAccountHandleResponse handles the ListByBillingAccount response.
func (client *EventsClient) listByBillingAccountHandleResponse(resp *http.Response) (EventsClientListByBillingAccountResponse, error) {
	result := EventsClientListByBillingAccountResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Events); err != nil {
		return EventsClientListByBillingAccountResponse{}, err
	}
	return result, nil
}

// ListByBillingProfile - Lists the events that decrements Azure credits or Microsoft Azure consumption commitment for a billing
// account or a billing profile for a given start and end date.
// If the operation fails it returns an *azcore.ResponseError type.
// billingAccountID - BillingAccount ID
// billingProfileID - Azure Billing Profile ID.
// startDate - Start date
// endDate - End date
// options - EventsClientListByBillingProfileOptions contains the optional parameters for the EventsClient.ListByBillingProfile
// method.
func (client *EventsClient) ListByBillingProfile(billingAccountID string, billingProfileID string, startDate string, endDate string, options *EventsClientListByBillingProfileOptions) *EventsClientListByBillingProfilePager {
	return &EventsClientListByBillingProfilePager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listByBillingProfileCreateRequest(ctx, billingAccountID, billingProfileID, startDate, endDate, options)
		},
		advancer: func(ctx context.Context, resp EventsClientListByBillingProfileResponse) (*policy.Request, error) {
			return runtime.NewRequest(ctx, http.MethodGet, *resp.Events.NextLink)
		},
	}
}

// listByBillingProfileCreateRequest creates the ListByBillingProfile request.
func (client *EventsClient) listByBillingProfileCreateRequest(ctx context.Context, billingAccountID string, billingProfileID string, startDate string, endDate string, options *EventsClientListByBillingProfileOptions) (*policy.Request, error) {
	urlPath := "/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/billingProfiles/{billingProfileId}/providers/Microsoft.Consumption/events"
	if billingAccountID == "" {
		return nil, errors.New("parameter billingAccountID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingAccountId}", url.PathEscape(billingAccountID))
	if billingProfileID == "" {
		return nil, errors.New("parameter billingProfileID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingProfileId}", url.PathEscape(billingProfileID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-10-01")
	reqQP.Set("startDate", startDate)
	reqQP.Set("endDate", endDate)
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listByBillingProfileHandleResponse handles the ListByBillingProfile response.
func (client *EventsClient) listByBillingProfileHandleResponse(resp *http.Response) (EventsClientListByBillingProfileResponse, error) {
	result := EventsClientListByBillingProfileResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.Events); err != nil {
		return EventsClientListByBillingProfileResponse{}, err
	}
	return result, nil
}
