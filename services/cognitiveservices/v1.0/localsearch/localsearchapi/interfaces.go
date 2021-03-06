package localsearchapi

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/cognitiveservices/v1.0/localsearch"
)

// LocalClientAPI contains the set of methods on the LocalClient type.
type LocalClientAPI interface {
	Search(ctx context.Context, query string, acceptLanguage string, pragma string, userAgent string, clientID string, clientIP string, location string, countryCode string, market string, localCategories string, localCircularView string, localMapView string, count string, first string, responseFormat []localsearch.ResponseFormat, safeSearch localsearch.SafeSearch, setLang string) (result localsearch.SearchResponse, err error)
}

var _ LocalClientAPI = (*localsearch.LocalClient)(nil)
