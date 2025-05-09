package labservices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// UsagesClient is the client for the Usages methods of the Labservices service.
type UsagesClient struct {
	BaseClient
}

// NewUsagesClient creates an instance of the UsagesClient client.
func NewUsagesClient(subscriptionID string) UsagesClient {
	return NewUsagesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewUsagesClientWithBaseURI creates an instance of the UsagesClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewUsagesClientWithBaseURI(baseURI string, subscriptionID string) UsagesClient {
	return UsagesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// ListByLocation returns list of usage per SKU family for the specified subscription in the specified region.
// Parameters:
// location - the location name.
// filter - the filter to apply to the operation.
func (client UsagesClient) ListByLocation(ctx context.Context, location string, filter string) (result ListUsagesResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsagesClient.ListByLocation")
		defer func() {
			sc := -1
			if result.lur.Response.Response != nil {
				sc = result.lur.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: location,
			Constraints: []validation.Constraint{{Target: "location", Name: validation.MaxLength, Rule: 100, Chain: nil},
				{Target: "location", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "location", Name: validation.Pattern, Rule: `^[-\w\._]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("labservices.UsagesClient", "ListByLocation", err.Error())
	}

	result.fn = client.listByLocationNextResults
	req, err := client.ListByLocationPreparer(ctx, location, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsagesClient", "ListByLocation", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.lur.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "labservices.UsagesClient", "ListByLocation", resp, "Failure sending request")
		return
	}

	result.lur, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsagesClient", "ListByLocation", resp, "Failure responding to request")
		return
	}
	if result.lur.hasNextLink() && result.lur.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListByLocationPreparer prepares the ListByLocation request.
func (client UsagesClient) ListByLocationPreparer(ctx context.Context, location string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2021-11-15-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.LabServices/locations/{location}/usages", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByLocationSender sends the ListByLocation request. The method will close the
// http.Response Body if it receives an error.
func (client UsagesClient) ListByLocationSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByLocationResponder handles the response to the ListByLocation request. The method always
// closes the http.Response Body.
func (client UsagesClient) ListByLocationResponder(resp *http.Response) (result ListUsagesResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByLocationNextResults retrieves the next set of results, if any.
func (client UsagesClient) listByLocationNextResults(ctx context.Context, lastResults ListUsagesResult) (result ListUsagesResult, err error) {
	req, err := lastResults.listUsagesResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "labservices.UsagesClient", "listByLocationNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByLocationSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "labservices.UsagesClient", "listByLocationNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByLocationResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "labservices.UsagesClient", "listByLocationNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByLocationComplete enumerates all values, automatically crossing page boundaries as required.
func (client UsagesClient) ListByLocationComplete(ctx context.Context, location string, filter string) (result ListUsagesResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/UsagesClient.ListByLocation")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListByLocation(ctx, location, filter)
	return
}
