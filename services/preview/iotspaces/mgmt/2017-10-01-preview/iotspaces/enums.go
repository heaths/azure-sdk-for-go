package iotspaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// NameUnavailabilityReason enumerates the values for name unavailability reason.
type NameUnavailabilityReason string

const (
	// AlreadyExists ...
	AlreadyExists NameUnavailabilityReason = "AlreadyExists"
	// Invalid ...
	Invalid NameUnavailabilityReason = "Invalid"
)

// PossibleNameUnavailabilityReasonValues returns an array of possible values for the NameUnavailabilityReason const type.
func PossibleNameUnavailabilityReasonValues() []NameUnavailabilityReason {
	return []NameUnavailabilityReason{AlreadyExists, Invalid}
}

// ProvisioningState enumerates the values for provisioning state.
type ProvisioningState string

const (
	// Canceled ...
	Canceled ProvisioningState = "Canceled"
	// Deleting ...
	Deleting ProvisioningState = "Deleting"
	// Failed ...
	Failed ProvisioningState = "Failed"
	// Provisioning ...
	Provisioning ProvisioningState = "Provisioning"
	// Succeeded ...
	Succeeded ProvisioningState = "Succeeded"
)

// PossibleProvisioningStateValues returns an array of possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{Canceled, Deleting, Failed, Provisioning, Succeeded}
}

// Sku enumerates the values for sku.
type Sku string

const (
	// F1 ...
	F1 Sku = "F1"
	// S1 ...
	S1 Sku = "S1"
	// S2 ...
	S2 Sku = "S2"
	// S3 ...
	S3 Sku = "S3"
)

// PossibleSkuValues returns an array of possible values for the Sku const type.
func PossibleSkuValues() []Sku {
	return []Sku{F1, S1, S2, S3}
}
