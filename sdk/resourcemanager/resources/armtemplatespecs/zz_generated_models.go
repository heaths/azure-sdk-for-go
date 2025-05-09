//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armtemplatespecs

import (
	"encoding/json"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"reflect"
	"time"
)

// AzureResourceBase - Common properties for all Azure resources.
type AzureResourceBase struct {
	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ClientCreateOrUpdateOptions contains the optional parameters for the Client.CreateOrUpdate method.
type ClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// ClientDeleteOptions contains the optional parameters for the Client.Delete method.
type ClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// ClientGetOptions contains the optional parameters for the Client.Get method.
type ClientGetOptions struct {
	// Allows for expansion of additional Template Spec details in the response. Optional.
	Expand *TemplateSpecExpandKind
}

// ClientListByResourceGroupOptions contains the optional parameters for the Client.ListByResourceGroup method.
type ClientListByResourceGroupOptions struct {
	// Allows for expansion of additional Template Spec details in the response. Optional.
	Expand *TemplateSpecExpandKind
}

// ClientListBySubscriptionOptions contains the optional parameters for the Client.ListBySubscription method.
type ClientListBySubscriptionOptions struct {
	// Allows for expansion of additional Template Spec details in the response. Optional.
	Expand *TemplateSpecExpandKind
}

// ClientUpdateOptions contains the optional parameters for the Client.Update method.
type ClientUpdateOptions struct {
	// Template Spec resource with the tags to be updated.
	TemplateSpec *TemplateSpecUpdateModel
}

// Error - Template Specs error response.
type Error struct {
	// Common error response for all Azure Resource Manager APIs to return error details for failed operations. (This also follows
	// the OData error response format.)
	Error *ErrorResponse `json:"error,omitempty"`
}

// ErrorAdditionalInfo - The resource management error additional info.
type ErrorAdditionalInfo struct {
	// READ-ONLY; The additional info.
	Info map[string]interface{} `json:"info,omitempty" azure:"ro"`

	// READ-ONLY; The additional info type.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// ErrorResponse - Common error response for all Azure Resource Manager APIs to return error details for failed operations.
// (This also follows the OData error response format.)
type ErrorResponse struct {
	// READ-ONLY; The error additional info.
	AdditionalInfo []*ErrorAdditionalInfo `json:"additionalInfo,omitempty" azure:"ro"`

	// READ-ONLY; The error code.
	Code *string `json:"code,omitempty" azure:"ro"`

	// READ-ONLY; The error details.
	Details []*ErrorResponse `json:"details,omitempty" azure:"ro"`

	// READ-ONLY; The error message.
	Message *string `json:"message,omitempty" azure:"ro"`

	// READ-ONLY; The error target.
	Target *string `json:"target,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ErrorResponse.
func (e ErrorResponse) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "additionalInfo", e.AdditionalInfo)
	populate(objectMap, "code", e.Code)
	populate(objectMap, "details", e.Details)
	populate(objectMap, "message", e.Message)
	populate(objectMap, "target", e.Target)
	return json.Marshal(objectMap)
}

// LinkedTemplateArtifact - Represents a Template Spec artifact containing an embedded Azure Resource Manager template for
// use as a linked template.
type LinkedTemplateArtifact struct {
	// REQUIRED; A filesystem safe relative path of the artifact.
	Path *string `json:"path,omitempty"`

	// REQUIRED; The Azure Resource Manager template.
	Template map[string]interface{} `json:"template,omitempty"`
}

// ListResult - List of Template Specs.
type ListResult struct {
	// An array of Template Specs.
	Value []*TemplateSpec `json:"value,omitempty"`

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type ListResult.
func (l ListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", l.NextLink)
	populate(objectMap, "value", l.Value)
	return json.Marshal(objectMap)
}

// SystemData - Metadata pertaining to creation and last modification of the resource.
type SystemData struct {
	// The timestamp of resource creation (UTC).
	CreatedAt *time.Time `json:"createdAt,omitempty"`

	// The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// The type of identity that created the resource.
	CreatedByType *CreatedByType `json:"createdByType,omitempty"`

	// The timestamp of resource last modification (UTC)
	LastModifiedAt *time.Time `json:"lastModifiedAt,omitempty"`

	// The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// The type of identity that last modified the resource.
	LastModifiedByType *CreatedByType `json:"lastModifiedByType,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type SystemData.
func (s SystemData) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populateTimeRFC3339(objectMap, "createdAt", s.CreatedAt)
	populate(objectMap, "createdBy", s.CreatedBy)
	populate(objectMap, "createdByType", s.CreatedByType)
	populateTimeRFC3339(objectMap, "lastModifiedAt", s.LastModifiedAt)
	populate(objectMap, "lastModifiedBy", s.LastModifiedBy)
	populate(objectMap, "lastModifiedByType", s.LastModifiedByType)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type SystemData.
func (s *SystemData) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "createdAt":
			err = unpopulateTimeRFC3339(val, &s.CreatedAt)
			delete(rawMsg, key)
		case "createdBy":
			err = unpopulate(val, &s.CreatedBy)
			delete(rawMsg, key)
		case "createdByType":
			err = unpopulate(val, &s.CreatedByType)
			delete(rawMsg, key)
		case "lastModifiedAt":
			err = unpopulateTimeRFC3339(val, &s.LastModifiedAt)
			delete(rawMsg, key)
		case "lastModifiedBy":
			err = unpopulate(val, &s.LastModifiedBy)
			delete(rawMsg, key)
		case "lastModifiedByType":
			err = unpopulate(val, &s.LastModifiedByType)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// TemplateSpec - Template Spec object.
type TemplateSpec struct {
	// REQUIRED; The location of the Template Spec. It cannot be changed after Template Spec creation. It must be one of the supported
	// Azure locations.
	Location *string `json:"location,omitempty"`

	// Template Spec properties.
	Properties *TemplateSpecProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type TemplateSpec.
func (t TemplateSpec) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "properties", t.Properties)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// TemplateSpecProperties - Template Spec properties.
type TemplateSpecProperties struct {
	// Template Spec description.
	Description *string `json:"description,omitempty"`

	// Template Spec display name.
	DisplayName *string `json:"displayName,omitempty"`

	// The Template Spec metadata. Metadata is an open-ended object and is typically a collection of key-value pairs.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// READ-ONLY; High-level information about the versions within this Template Spec. The keys are the version names. Only populated
	// if the $expand query parameter is set to 'versions'.
	Versions map[string]*TemplateSpecVersionInfo `json:"versions,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type TemplateSpecProperties.
func (t TemplateSpecProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", t.Description)
	populate(objectMap, "displayName", t.DisplayName)
	populate(objectMap, "metadata", t.Metadata)
	populate(objectMap, "versions", t.Versions)
	return json.Marshal(objectMap)
}

// TemplateSpecUpdateModel - Template Spec properties to be updated (only tags are currently supported).
type TemplateSpecUpdateModel struct {
	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type TemplateSpecUpdateModel.
func (t TemplateSpecUpdateModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// TemplateSpecVersion - Template Spec Version object.
type TemplateSpecVersion struct {
	// REQUIRED; The location of the Template Spec Version. It must match the location of the parent Template Spec.
	Location *string `json:"location,omitempty"`

	// REQUIRED; Template Spec Version properties.
	Properties *TemplateSpecVersionProperties `json:"properties,omitempty"`

	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type TemplateSpecVersion.
func (t TemplateSpecVersion) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "location", t.Location)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "properties", t.Properties)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// TemplateSpecVersionInfo - High-level information about a Template Spec version.
type TemplateSpecVersionInfo struct {
	// READ-ONLY; Template Spec version description.
	Description *string `json:"description,omitempty" azure:"ro"`

	// READ-ONLY; The timestamp of when the version was created.
	TimeCreated *time.Time `json:"timeCreated,omitempty" azure:"ro"`

	// READ-ONLY; The timestamp of when the version was last modified.
	TimeModified *time.Time `json:"timeModified,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type TemplateSpecVersionInfo.
func (t TemplateSpecVersionInfo) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", t.Description)
	populateTimeRFC3339(objectMap, "timeCreated", t.TimeCreated)
	populateTimeRFC3339(objectMap, "timeModified", t.TimeModified)
	return json.Marshal(objectMap)
}

// UnmarshalJSON implements the json.Unmarshaller interface for type TemplateSpecVersionInfo.
func (t *TemplateSpecVersionInfo) UnmarshalJSON(data []byte) error {
	var rawMsg map[string]json.RawMessage
	if err := json.Unmarshal(data, &rawMsg); err != nil {
		return err
	}
	for key, val := range rawMsg {
		var err error
		switch key {
		case "description":
			err = unpopulate(val, &t.Description)
			delete(rawMsg, key)
		case "timeCreated":
			err = unpopulateTimeRFC3339(val, &t.TimeCreated)
			delete(rawMsg, key)
		case "timeModified":
			err = unpopulateTimeRFC3339(val, &t.TimeModified)
			delete(rawMsg, key)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

// TemplateSpecVersionProperties - Template Spec Version properties.
type TemplateSpecVersionProperties struct {
	// Template Spec version description.
	Description *string `json:"description,omitempty"`

	// An array of linked template artifacts.
	LinkedTemplates []*LinkedTemplateArtifact `json:"linkedTemplates,omitempty"`

	// The main Azure Resource Manager template content.
	MainTemplate map[string]interface{} `json:"mainTemplate,omitempty"`

	// The version metadata. Metadata is an open-ended object and is typically a collection of key-value pairs.
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// The Azure Resource Manager template UI definition content.
	UIFormDefinition map[string]interface{} `json:"uiFormDefinition,omitempty"`
}

// MarshalJSON implements the json.Marshaller interface for type TemplateSpecVersionProperties.
func (t TemplateSpecVersionProperties) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "description", t.Description)
	populate(objectMap, "linkedTemplates", t.LinkedTemplates)
	populate(objectMap, "mainTemplate", t.MainTemplate)
	populate(objectMap, "metadata", t.Metadata)
	populate(objectMap, "uiFormDefinition", t.UIFormDefinition)
	return json.Marshal(objectMap)
}

// TemplateSpecVersionUpdateModel - Template Spec Version properties to be updated (only tags are currently supported).
type TemplateSpecVersionUpdateModel struct {
	// Resource tags.
	Tags map[string]*string `json:"tags,omitempty"`

	// READ-ONLY; String Id used to locate any resource on Azure.
	ID *string `json:"id,omitempty" azure:"ro"`

	// READ-ONLY; Name of this resource.
	Name *string `json:"name,omitempty" azure:"ro"`

	// READ-ONLY; Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData `json:"systemData,omitempty" azure:"ro"`

	// READ-ONLY; Type of this resource.
	Type *string `json:"type,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type TemplateSpecVersionUpdateModel.
func (t TemplateSpecVersionUpdateModel) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "id", t.ID)
	populate(objectMap, "name", t.Name)
	populate(objectMap, "systemData", t.SystemData)
	populate(objectMap, "tags", t.Tags)
	populate(objectMap, "type", t.Type)
	return json.Marshal(objectMap)
}

// TemplateSpecVersionsClientCreateOrUpdateOptions contains the optional parameters for the TemplateSpecVersionsClient.CreateOrUpdate
// method.
type TemplateSpecVersionsClientCreateOrUpdateOptions struct {
	// placeholder for future optional parameters
}

// TemplateSpecVersionsClientDeleteOptions contains the optional parameters for the TemplateSpecVersionsClient.Delete method.
type TemplateSpecVersionsClientDeleteOptions struct {
	// placeholder for future optional parameters
}

// TemplateSpecVersionsClientGetOptions contains the optional parameters for the TemplateSpecVersionsClient.Get method.
type TemplateSpecVersionsClientGetOptions struct {
	// placeholder for future optional parameters
}

// TemplateSpecVersionsClientListOptions contains the optional parameters for the TemplateSpecVersionsClient.List method.
type TemplateSpecVersionsClientListOptions struct {
	// placeholder for future optional parameters
}

// TemplateSpecVersionsClientUpdateOptions contains the optional parameters for the TemplateSpecVersionsClient.Update method.
type TemplateSpecVersionsClientUpdateOptions struct {
	// Template Spec Version resource with the tags to be updated.
	TemplateSpecVersionUpdateModel *TemplateSpecVersionUpdateModel
}

// TemplateSpecVersionsListResult - List of Template Specs versions
type TemplateSpecVersionsListResult struct {
	// An array of Template Spec versions.
	Value []*TemplateSpecVersion `json:"value,omitempty"`

	// READ-ONLY; The URL to use for getting the next set of results.
	NextLink *string `json:"nextLink,omitempty" azure:"ro"`
}

// MarshalJSON implements the json.Marshaller interface for type TemplateSpecVersionsListResult.
func (t TemplateSpecVersionsListResult) MarshalJSON() ([]byte, error) {
	objectMap := make(map[string]interface{})
	populate(objectMap, "nextLink", t.NextLink)
	populate(objectMap, "value", t.Value)
	return json.Marshal(objectMap)
}

func populate(m map[string]interface{}, k string, v interface{}) {
	if v == nil {
		return
	} else if azcore.IsNullValue(v) {
		m[k] = nil
	} else if !reflect.ValueOf(v).IsNil() {
		m[k] = v
	}
}

func unpopulate(data json.RawMessage, v interface{}) error {
	if data == nil {
		return nil
	}
	return json.Unmarshal(data, v)
}
