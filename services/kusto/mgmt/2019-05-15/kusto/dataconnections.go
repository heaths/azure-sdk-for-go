package kusto

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

// DataConnectionsClient is the the Azure Kusto management API provides a RESTful set of web services that interact
// with Azure Kusto services to manage your clusters and databases. The API enables you to create, update, and delete
// clusters and databases.
type DataConnectionsClient struct {
	BaseClient
}

// NewDataConnectionsClient creates an instance of the DataConnectionsClient client.
func NewDataConnectionsClient(subscriptionID string) DataConnectionsClient {
	return NewDataConnectionsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDataConnectionsClientWithBaseURI creates an instance of the DataConnectionsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDataConnectionsClientWithBaseURI(baseURI string, subscriptionID string) DataConnectionsClient {
	return DataConnectionsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CheckNameAvailability checks that the data connection name is valid and is not already in use.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// dataConnectionName - the name of the data connection.
func (client DataConnectionsClient) CheckNameAvailability(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName DataConnectionCheckNameRequest) (result CheckNameResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.CheckNameAvailability")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: dataConnectionName,
			Constraints: []validation.Constraint{{Target: "dataConnectionName.Name", Name: validation.Null, Rule: true, Chain: nil},
				{Target: "dataConnectionName.Type", Name: validation.Null, Rule: true, Chain: nil}}}}); err != nil {
		return result, validation.NewError("kusto.DataConnectionsClient", "CheckNameAvailability", err.Error())
	}

	req, err := client.CheckNameAvailabilityPreparer(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "CheckNameAvailability", nil, "Failure preparing request")
		return
	}

	resp, err := client.CheckNameAvailabilitySender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "CheckNameAvailability", resp, "Failure sending request")
		return
	}

	result, err = client.CheckNameAvailabilityResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "CheckNameAvailability", resp, "Failure responding to request")
		return
	}

	return
}

// CheckNameAvailabilityPreparer prepares the CheckNameAvailability request.
func (client DataConnectionsClient) CheckNameAvailabilityPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName DataConnectionCheckNameRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/checkNameAvailability", pathParameters),
		autorest.WithJSON(dataConnectionName),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CheckNameAvailabilitySender sends the CheckNameAvailability request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) CheckNameAvailabilitySender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// CheckNameAvailabilityResponder handles the response to the CheckNameAvailability request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) CheckNameAvailabilityResponder(resp *http.Response) (result CheckNameResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// CreateOrUpdate creates or updates a data connection.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// dataConnectionName - the name of the data connection.
// parameters - the data connection parameters supplied to the CreateOrUpdate operation.
func (client DataConnectionsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters BasicDataConnection) (result DataConnectionsCreateOrUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.CreateOrUpdate")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client DataConnectionsClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters BasicDataConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":        autorest.Encode("path", clusterName),
		"databaseName":       autorest.Encode("path", databaseName),
		"dataConnectionName": autorest.Encode("path", dataConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) CreateOrUpdateSender(req *http.Request) (future DataConnectionsCreateOrUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) CreateOrUpdateResponder(resp *http.Response) (result DataConnectionModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// DataConnectionValidationMethod checks that the data connection parameters are valid.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// parameters - the data connection parameters supplied to the CreateOrUpdate operation.
func (client DataConnectionsClient) DataConnectionValidationMethod(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DataConnectionValidation) (result DataConnectionValidationListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.DataConnectionValidationMethod")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DataConnectionValidationMethodPreparer(ctx, resourceGroupName, clusterName, databaseName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "DataConnectionValidationMethod", nil, "Failure preparing request")
		return
	}

	resp, err := client.DataConnectionValidationMethodSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "DataConnectionValidationMethod", resp, "Failure sending request")
		return
	}

	result, err = client.DataConnectionValidationMethodResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "DataConnectionValidationMethod", resp, "Failure responding to request")
		return
	}

	return
}

// DataConnectionValidationMethodPreparer prepares the DataConnectionValidationMethod request.
func (client DataConnectionsClient) DataConnectionValidationMethodPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, parameters DataConnectionValidation) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnectionValidation", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DataConnectionValidationMethodSender sends the DataConnectionValidationMethod request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) DataConnectionValidationMethodSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// DataConnectionValidationMethodResponder handles the response to the DataConnectionValidationMethod request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) DataConnectionValidationMethodResponder(resp *http.Response) (result DataConnectionValidationListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the data connection with the given name.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// dataConnectionName - the name of the data connection.
func (client DataConnectionsClient) Delete(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string) (result DataConnectionsDeleteFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.Delete")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.DeletePreparer(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client DataConnectionsClient) DeletePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":        autorest.Encode("path", clusterName),
		"databaseName":       autorest.Encode("path", databaseName),
		"dataConnectionName": autorest.Encode("path", dataConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) DeleteSender(req *http.Request) (future DataConnectionsDeleteFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get returns a data connection.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// dataConnectionName - the name of the data connection.
func (client DataConnectionsClient) Get(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string) (result DataConnectionModel, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DataConnectionsClient) GetPreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":        autorest.Encode("path", clusterName),
		"databaseName":       autorest.Encode("path", databaseName),
		"dataConnectionName": autorest.Encode("path", dataConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) GetResponder(resp *http.Response) (result DataConnectionModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// ListByDatabase returns the list of data connections of the given Kusto database.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
func (client DataConnectionsClient) ListByDatabase(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (result DataConnectionListResult, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.ListByDatabase")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.ListByDatabasePreparer(ctx, resourceGroupName, clusterName, databaseName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "ListByDatabase", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByDatabaseSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "ListByDatabase", resp, "Failure sending request")
		return
	}

	result, err = client.ListByDatabaseResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "ListByDatabase", resp, "Failure responding to request")
		return
	}

	return
}

// ListByDatabasePreparer prepares the ListByDatabase request.
func (client DataConnectionsClient) ListByDatabasePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":       autorest.Encode("path", clusterName),
		"databaseName":      autorest.Encode("path", databaseName),
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByDatabaseSender sends the ListByDatabase request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) ListByDatabaseSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListByDatabaseResponder handles the response to the ListByDatabase request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) ListByDatabaseResponder(resp *http.Response) (result DataConnectionListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Update updates a data connection.
// Parameters:
// resourceGroupName - the name of the resource group containing the Kusto cluster.
// clusterName - the name of the Kusto cluster.
// databaseName - the name of the database in the Kusto cluster.
// dataConnectionName - the name of the data connection.
// parameters - the data connection parameters supplied to the Update operation.
func (client DataConnectionsClient) Update(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters BasicDataConnection) (result DataConnectionsUpdateFuture, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DataConnectionsClient.Update")
		defer func() {
			sc := -1
			if result.FutureAPI != nil && result.FutureAPI.Response() != nil {
				sc = result.FutureAPI.Response().StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.UpdatePreparer(ctx, resourceGroupName, clusterName, databaseName, dataConnectionName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Update", nil, "Failure preparing request")
		return
	}

	result, err = client.UpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "kusto.DataConnectionsClient", "Update", result.Response(), "Failure sending request")
		return
	}

	return
}

// UpdatePreparer prepares the Update request.
func (client DataConnectionsClient) UpdatePreparer(ctx context.Context, resourceGroupName string, clusterName string, databaseName string, dataConnectionName string, parameters BasicDataConnection) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"clusterName":        autorest.Encode("path", clusterName),
		"databaseName":       autorest.Encode("path", databaseName),
		"dataConnectionName": autorest.Encode("path", dataConnectionName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2019-05-15"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Kusto/clusters/{clusterName}/databases/{databaseName}/dataConnections/{dataConnectionName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// UpdateSender sends the Update request. The method will close the
// http.Response Body if it receives an error.
func (client DataConnectionsClient) UpdateSender(req *http.Request) (future DataConnectionsUpdateFuture, err error) {
	var resp *http.Response
	future.FutureAPI = &azure.Future{}
	resp, err = client.Send(req, azure.DoRetryWithRegistration(client.Client))
	if err != nil {
		return
	}
	var azf azure.Future
	azf, err = azure.NewFutureFromResponse(resp)
	future.FutureAPI = &azf
	future.Result = future.result
	return
}

// UpdateResponder handles the response to the Update request. The method always
// closes the http.Response Body.
func (client DataConnectionsClient) UpdateResponder(resp *http.Response) (result DataConnectionModel, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Value),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
