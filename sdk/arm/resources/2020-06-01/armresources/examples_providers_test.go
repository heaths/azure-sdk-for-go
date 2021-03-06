// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package armresources_test

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/arm/resources/2020-06-01/armresources"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
)

func ExampleProvidersClient_Register() {
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatalf("failed to obtain a credential: %v", err)
	}
	client := armresources.NewProvidersClient(armcore.NewDefaultConnection(cred, nil), "<subscription ID>")
	resp, err := client.Register(context.Background(), "<resource provider namespace>", nil)
	if err != nil {
		log.Fatalf("failed to obtain a response: %v", err)
	}
	log.Printf("resource group exists: %v\n", *resp.Provider.ID)
}
