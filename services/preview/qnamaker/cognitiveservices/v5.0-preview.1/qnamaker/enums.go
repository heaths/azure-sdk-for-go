package qnamaker

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// EnvironmentType enumerates the values for environment type.
type EnvironmentType string

const (
	// Prod ...
	Prod EnvironmentType = "Prod"
	// Test ...
	Test EnvironmentType = "Test"
)

// PossibleEnvironmentTypeValues returns an array of possible values for the EnvironmentType const type.
func PossibleEnvironmentTypeValues() []EnvironmentType {
	return []EnvironmentType{Prod, Test}
}

// ErrorCodeType enumerates the values for error code type.
type ErrorCodeType string

const (
	// BadArgument ...
	BadArgument ErrorCodeType = "BadArgument"
	// EndpointKeysError ...
	EndpointKeysError ErrorCodeType = "EndpointKeysError"
	// ExtractionFailure ...
	ExtractionFailure ErrorCodeType = "ExtractionFailure"
	// Forbidden ...
	Forbidden ErrorCodeType = "Forbidden"
	// KbNotFound ...
	KbNotFound ErrorCodeType = "KbNotFound"
	// NotFound ...
	NotFound ErrorCodeType = "NotFound"
	// OperationNotFound ...
	OperationNotFound ErrorCodeType = "OperationNotFound"
	// QnaRuntimeError ...
	QnaRuntimeError ErrorCodeType = "QnaRuntimeError"
	// QuotaExceeded ...
	QuotaExceeded ErrorCodeType = "QuotaExceeded"
	// ServiceError ...
	ServiceError ErrorCodeType = "ServiceError"
	// SKULimitExceeded ...
	SKULimitExceeded ErrorCodeType = "SKULimitExceeded"
	// Unauthorized ...
	Unauthorized ErrorCodeType = "Unauthorized"
	// Unspecified ...
	Unspecified ErrorCodeType = "Unspecified"
	// ValidationFailure ...
	ValidationFailure ErrorCodeType = "ValidationFailure"
)

// PossibleErrorCodeTypeValues returns an array of possible values for the ErrorCodeType const type.
func PossibleErrorCodeTypeValues() []ErrorCodeType {
	return []ErrorCodeType{BadArgument, EndpointKeysError, ExtractionFailure, Forbidden, KbNotFound, NotFound, OperationNotFound, QnaRuntimeError, QuotaExceeded, ServiceError, SKULimitExceeded, Unauthorized, Unspecified, ValidationFailure}
}

// OperationStateType enumerates the values for operation state type.
type OperationStateType string

const (
	// Failed ...
	Failed OperationStateType = "Failed"
	// NotStarted ...
	NotStarted OperationStateType = "NotStarted"
	// Running ...
	Running OperationStateType = "Running"
	// Succeeded ...
	Succeeded OperationStateType = "Succeeded"
)

// PossibleOperationStateTypeValues returns an array of possible values for the OperationStateType const type.
func PossibleOperationStateTypeValues() []OperationStateType {
	return []OperationStateType{Failed, NotStarted, Running, Succeeded}
}

// StrictFiltersCompoundOperationType enumerates the values for strict filters compound operation type.
type StrictFiltersCompoundOperationType string

const (
	// AND ...
	AND StrictFiltersCompoundOperationType = "AND"
	// OR ...
	OR StrictFiltersCompoundOperationType = "OR"
)

// PossibleStrictFiltersCompoundOperationTypeValues returns an array of possible values for the StrictFiltersCompoundOperationType const type.
func PossibleStrictFiltersCompoundOperationTypeValues() []StrictFiltersCompoundOperationType {
	return []StrictFiltersCompoundOperationType{AND, OR}
}
