package consumption

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

// TagsClient is the consumption management client provides access to consumption resources for Azure Enterprise
// Subscriptions.
type TagsClient struct {
	BaseClient
}

// NewTagsClient creates an instance of the TagsClient client.
func NewTagsClient(subscriptionID string) TagsClient {
	return NewTagsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewTagsClientWithBaseURI creates an instance of the TagsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewTagsClientWithBaseURI(baseURI string, subscriptionID string) TagsClient {
	return TagsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get all available tag keys for a billing account.
// Parameters:
// billingAccountID - billingAccount ID
func (client TagsClient) Get(ctx context.Context, billingAccountID string) (result TagsResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/TagsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, billingAccountID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.TagsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "consumption.TagsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumption.TagsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client TagsClient) GetPreparer(ctx context.Context, billingAccountID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"billingAccountId": autorest.Encode("path", billingAccountID),
	}

	const APIVersion = "2018-08-31"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/providers/Microsoft.CostManagement/billingAccounts/{billingAccountId}/providers/Microsoft.Consumption/tags", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client TagsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client TagsClient) GetResponder(resp *http.Response) (result TagsResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
