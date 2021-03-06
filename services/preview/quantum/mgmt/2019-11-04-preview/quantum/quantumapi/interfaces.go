package quantumapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/preview/quantum/mgmt/2019-11-04-preview/quantum"
)

// WorkspacesClientAPI contains the set of methods on the WorkspacesClient type.
type WorkspacesClientAPI interface {
	CreateOrUpdate(ctx context.Context, resourceGroupName string, workspaceName string, quantumWorkspace quantum.Workspace) (result quantum.WorkspacesCreateOrUpdateFuture, err error)
	Delete(ctx context.Context, resourceGroupName string, workspaceName string) (result quantum.WorkspacesDeleteFuture, err error)
	Get(ctx context.Context, resourceGroupName string, workspaceName string) (result quantum.Workspace, err error)
	ListByResourceGroup(ctx context.Context, resourceGroupName string) (result quantum.WorkspaceListResultPage, err error)
	ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result quantum.WorkspaceListResultIterator, err error)
	ListBySubscription(ctx context.Context) (result quantum.WorkspaceListResultPage, err error)
	ListBySubscriptionComplete(ctx context.Context) (result quantum.WorkspaceListResultIterator, err error)
	UpdateTags(ctx context.Context, resourceGroupName string, workspaceName string, workspaceTags quantum.TagsObject) (result quantum.Workspace, err error)
}

var _ WorkspacesClientAPI = (*quantum.WorkspacesClient)(nil)

// OfferingsClientAPI contains the set of methods on the OfferingsClient type.
type OfferingsClientAPI interface {
	List(ctx context.Context, locationName string) (result quantum.OfferingsListResultPage, err error)
	ListComplete(ctx context.Context, locationName string) (result quantum.OfferingsListResultIterator, err error)
}

var _ OfferingsClientAPI = (*quantum.OfferingsClient)(nil)

// OperationsClientAPI contains the set of methods on the OperationsClient type.
type OperationsClientAPI interface {
	List(ctx context.Context) (result quantum.OperationsListPage, err error)
	ListComplete(ctx context.Context) (result quantum.OperationsListIterator, err error)
}

var _ OperationsClientAPI = (*quantum.OperationsClient)(nil)
