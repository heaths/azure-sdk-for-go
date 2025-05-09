//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armautomation

const (
	moduleName    = "armautomation"
	moduleVersion = "v0.3.0"
)

// AgentRegistrationKeyName - Gets or sets the agent registration key name - primary or secondary.
type AgentRegistrationKeyName string

const (
	AgentRegistrationKeyNamePrimary   AgentRegistrationKeyName = "primary"
	AgentRegistrationKeyNameSecondary AgentRegistrationKeyName = "secondary"
)

// PossibleAgentRegistrationKeyNameValues returns the possible values for the AgentRegistrationKeyName const type.
func PossibleAgentRegistrationKeyNameValues() []AgentRegistrationKeyName {
	return []AgentRegistrationKeyName{
		AgentRegistrationKeyNamePrimary,
		AgentRegistrationKeyNameSecondary,
	}
}

// ToPtr returns a *AgentRegistrationKeyName pointing to the current value.
func (c AgentRegistrationKeyName) ToPtr() *AgentRegistrationKeyName {
	return &c
}

// AutomationAccountState - Gets status of account.
type AutomationAccountState string

const (
	AutomationAccountStateOk          AutomationAccountState = "Ok"
	AutomationAccountStateSuspended   AutomationAccountState = "Suspended"
	AutomationAccountStateUnavailable AutomationAccountState = "Unavailable"
)

// PossibleAutomationAccountStateValues returns the possible values for the AutomationAccountState const type.
func PossibleAutomationAccountStateValues() []AutomationAccountState {
	return []AutomationAccountState{
		AutomationAccountStateOk,
		AutomationAccountStateSuspended,
		AutomationAccountStateUnavailable,
	}
}

// ToPtr returns a *AutomationAccountState pointing to the current value.
func (c AutomationAccountState) ToPtr() *AutomationAccountState {
	return &c
}

// AutomationKeyName - Automation key name.
type AutomationKeyName string

const (
	AutomationKeyNamePrimary   AutomationKeyName = "Primary"
	AutomationKeyNameSecondary AutomationKeyName = "Secondary"
)

// PossibleAutomationKeyNameValues returns the possible values for the AutomationKeyName const type.
func PossibleAutomationKeyNameValues() []AutomationKeyName {
	return []AutomationKeyName{
		AutomationKeyNamePrimary,
		AutomationKeyNameSecondary,
	}
}

// ToPtr returns a *AutomationKeyName pointing to the current value.
func (c AutomationKeyName) ToPtr() *AutomationKeyName {
	return &c
}

// AutomationKeyPermissions - Automation key permissions.
type AutomationKeyPermissions string

const (
	AutomationKeyPermissionsFull AutomationKeyPermissions = "Full"
	AutomationKeyPermissionsRead AutomationKeyPermissions = "Read"
)

// PossibleAutomationKeyPermissionsValues returns the possible values for the AutomationKeyPermissions const type.
func PossibleAutomationKeyPermissionsValues() []AutomationKeyPermissions {
	return []AutomationKeyPermissions{
		AutomationKeyPermissionsFull,
		AutomationKeyPermissionsRead,
	}
}

// ToPtr returns a *AutomationKeyPermissions pointing to the current value.
func (c AutomationKeyPermissions) ToPtr() *AutomationKeyPermissions {
	return &c
}

// ContentSourceType - Gets or sets the content source type.
type ContentSourceType string

const (
	ContentSourceTypeEmbeddedContent ContentSourceType = "embeddedContent"
	ContentSourceTypeURI             ContentSourceType = "uri"
)

// PossibleContentSourceTypeValues returns the possible values for the ContentSourceType const type.
func PossibleContentSourceTypeValues() []ContentSourceType {
	return []ContentSourceType{
		ContentSourceTypeEmbeddedContent,
		ContentSourceTypeURI,
	}
}

// ToPtr returns a *ContentSourceType pointing to the current value.
func (c ContentSourceType) ToPtr() *ContentSourceType {
	return &c
}

type CountType string

const (
	CountTypeNodeconfiguration CountType = "nodeconfiguration"
	CountTypeStatus            CountType = "status"
)

// PossibleCountTypeValues returns the possible values for the CountType const type.
func PossibleCountTypeValues() []CountType {
	return []CountType{
		CountTypeNodeconfiguration,
		CountTypeStatus,
	}
}

// ToPtr returns a *CountType pointing to the current value.
func (c CountType) ToPtr() *CountType {
	return &c
}

// CreatedByType - The type of identity that created the resource.
type CreatedByType string

const (
	CreatedByTypeApplication     CreatedByType = "Application"
	CreatedByTypeKey             CreatedByType = "Key"
	CreatedByTypeManagedIdentity CreatedByType = "ManagedIdentity"
	CreatedByTypeUser            CreatedByType = "User"
)

// PossibleCreatedByTypeValues returns the possible values for the CreatedByType const type.
func PossibleCreatedByTypeValues() []CreatedByType {
	return []CreatedByType{
		CreatedByTypeApplication,
		CreatedByTypeKey,
		CreatedByTypeManagedIdentity,
		CreatedByTypeUser,
	}
}

// ToPtr returns a *CreatedByType pointing to the current value.
func (c CreatedByType) ToPtr() *CreatedByType {
	return &c
}

// DscConfigurationState - Gets or sets the state of the configuration.
type DscConfigurationState string

const (
	DscConfigurationStateEdit      DscConfigurationState = "Edit"
	DscConfigurationStateNew       DscConfigurationState = "New"
	DscConfigurationStatePublished DscConfigurationState = "Published"
)

// PossibleDscConfigurationStateValues returns the possible values for the DscConfigurationState const type.
func PossibleDscConfigurationStateValues() []DscConfigurationState {
	return []DscConfigurationState{
		DscConfigurationStateEdit,
		DscConfigurationStateNew,
		DscConfigurationStatePublished,
	}
}

// ToPtr returns a *DscConfigurationState pointing to the current value.
func (c DscConfigurationState) ToPtr() *DscConfigurationState {
	return &c
}

// EncryptionKeySourceType - Encryption Key Source
type EncryptionKeySourceType string

const (
	EncryptionKeySourceTypeMicrosoftAutomation EncryptionKeySourceType = "Microsoft.Automation"
	EncryptionKeySourceTypeMicrosoftKeyvault   EncryptionKeySourceType = "Microsoft.Keyvault"
)

// PossibleEncryptionKeySourceTypeValues returns the possible values for the EncryptionKeySourceType const type.
func PossibleEncryptionKeySourceTypeValues() []EncryptionKeySourceType {
	return []EncryptionKeySourceType{
		EncryptionKeySourceTypeMicrosoftAutomation,
		EncryptionKeySourceTypeMicrosoftKeyvault,
	}
}

// ToPtr returns a *EncryptionKeySourceType pointing to the current value.
func (c EncryptionKeySourceType) ToPtr() *EncryptionKeySourceType {
	return &c
}

// GraphRunbookType - Runbook Type
type GraphRunbookType string

const (
	GraphRunbookTypeGraphPowerShell         GraphRunbookType = "GraphPowerShell"
	GraphRunbookTypeGraphPowerShellWorkflow GraphRunbookType = "GraphPowerShellWorkflow"
)

// PossibleGraphRunbookTypeValues returns the possible values for the GraphRunbookType const type.
func PossibleGraphRunbookTypeValues() []GraphRunbookType {
	return []GraphRunbookType{
		GraphRunbookTypeGraphPowerShell,
		GraphRunbookTypeGraphPowerShellWorkflow,
	}
}

// ToPtr returns a *GraphRunbookType pointing to the current value.
func (c GraphRunbookType) ToPtr() *GraphRunbookType {
	return &c
}

// GroupTypeEnum - Type of the HybridWorkerGroup.
type GroupTypeEnum string

const (
	GroupTypeEnumSystem GroupTypeEnum = "System"
	GroupTypeEnumUser   GroupTypeEnum = "User"
)

// PossibleGroupTypeEnumValues returns the possible values for the GroupTypeEnum const type.
func PossibleGroupTypeEnumValues() []GroupTypeEnum {
	return []GroupTypeEnum{
		GroupTypeEnumSystem,
		GroupTypeEnumUser,
	}
}

// ToPtr returns a *GroupTypeEnum pointing to the current value.
func (c GroupTypeEnum) ToPtr() *GroupTypeEnum {
	return &c
}

type HTTPStatusCode string

const (
	HTTPStatusCodeAccepted                     HTTPStatusCode = "Accepted"
	HTTPStatusCodeAmbiguous                    HTTPStatusCode = "Ambiguous"
	HTTPStatusCodeBadGateway                   HTTPStatusCode = "BadGateway"
	HTTPStatusCodeBadRequest                   HTTPStatusCode = "BadRequest"
	HTTPStatusCodeConflict                     HTTPStatusCode = "Conflict"
	HTTPStatusCodeContinue                     HTTPStatusCode = "Continue"
	HTTPStatusCodeCreated                      HTTPStatusCode = "Created"
	HTTPStatusCodeExpectationFailed            HTTPStatusCode = "ExpectationFailed"
	HTTPStatusCodeForbidden                    HTTPStatusCode = "Forbidden"
	HTTPStatusCodeFound                        HTTPStatusCode = "Found"
	HTTPStatusCodeGatewayTimeout               HTTPStatusCode = "GatewayTimeout"
	HTTPStatusCodeGone                         HTTPStatusCode = "Gone"
	HTTPStatusCodeHTTPVersionNotSupported      HTTPStatusCode = "HttpVersionNotSupported"
	HTTPStatusCodeInternalServerError          HTTPStatusCode = "InternalServerError"
	HTTPStatusCodeLengthRequired               HTTPStatusCode = "LengthRequired"
	HTTPStatusCodeMethodNotAllowed             HTTPStatusCode = "MethodNotAllowed"
	HTTPStatusCodeMoved                        HTTPStatusCode = "Moved"
	HTTPStatusCodeMovedPermanently             HTTPStatusCode = "MovedPermanently"
	HTTPStatusCodeMultipleChoices              HTTPStatusCode = "MultipleChoices"
	HTTPStatusCodeNoContent                    HTTPStatusCode = "NoContent"
	HTTPStatusCodeNonAuthoritativeInformation  HTTPStatusCode = "NonAuthoritativeInformation"
	HTTPStatusCodeNotAcceptable                HTTPStatusCode = "NotAcceptable"
	HTTPStatusCodeNotFound                     HTTPStatusCode = "NotFound"
	HTTPStatusCodeNotImplemented               HTTPStatusCode = "NotImplemented"
	HTTPStatusCodeNotModified                  HTTPStatusCode = "NotModified"
	HTTPStatusCodeOK                           HTTPStatusCode = "OK"
	HTTPStatusCodePartialContent               HTTPStatusCode = "PartialContent"
	HTTPStatusCodePaymentRequired              HTTPStatusCode = "PaymentRequired"
	HTTPStatusCodePreconditionFailed           HTTPStatusCode = "PreconditionFailed"
	HTTPStatusCodeProxyAuthenticationRequired  HTTPStatusCode = "ProxyAuthenticationRequired"
	HTTPStatusCodeRedirect                     HTTPStatusCode = "Redirect"
	HTTPStatusCodeRedirectKeepVerb             HTTPStatusCode = "RedirectKeepVerb"
	HTTPStatusCodeRedirectMethod               HTTPStatusCode = "RedirectMethod"
	HTTPStatusCodeRequestEntityTooLarge        HTTPStatusCode = "RequestEntityTooLarge"
	HTTPStatusCodeRequestTimeout               HTTPStatusCode = "RequestTimeout"
	HTTPStatusCodeRequestURITooLong            HTTPStatusCode = "RequestUriTooLong"
	HTTPStatusCodeRequestedRangeNotSatisfiable HTTPStatusCode = "RequestedRangeNotSatisfiable"
	HTTPStatusCodeResetContent                 HTTPStatusCode = "ResetContent"
	HTTPStatusCodeSeeOther                     HTTPStatusCode = "SeeOther"
	HTTPStatusCodeServiceUnavailable           HTTPStatusCode = "ServiceUnavailable"
	HTTPStatusCodeSwitchingProtocols           HTTPStatusCode = "SwitchingProtocols"
	HTTPStatusCodeTemporaryRedirect            HTTPStatusCode = "TemporaryRedirect"
	HTTPStatusCodeUnauthorized                 HTTPStatusCode = "Unauthorized"
	HTTPStatusCodeUnsupportedMediaType         HTTPStatusCode = "UnsupportedMediaType"
	HTTPStatusCodeUnused                       HTTPStatusCode = "Unused"
	HTTPStatusCodeUpgradeRequired              HTTPStatusCode = "UpgradeRequired"
	HTTPStatusCodeUseProxy                     HTTPStatusCode = "UseProxy"
)

// PossibleHTTPStatusCodeValues returns the possible values for the HTTPStatusCode const type.
func PossibleHTTPStatusCodeValues() []HTTPStatusCode {
	return []HTTPStatusCode{
		HTTPStatusCodeAccepted,
		HTTPStatusCodeAmbiguous,
		HTTPStatusCodeBadGateway,
		HTTPStatusCodeBadRequest,
		HTTPStatusCodeConflict,
		HTTPStatusCodeContinue,
		HTTPStatusCodeCreated,
		HTTPStatusCodeExpectationFailed,
		HTTPStatusCodeForbidden,
		HTTPStatusCodeFound,
		HTTPStatusCodeGatewayTimeout,
		HTTPStatusCodeGone,
		HTTPStatusCodeHTTPVersionNotSupported,
		HTTPStatusCodeInternalServerError,
		HTTPStatusCodeLengthRequired,
		HTTPStatusCodeMethodNotAllowed,
		HTTPStatusCodeMoved,
		HTTPStatusCodeMovedPermanently,
		HTTPStatusCodeMultipleChoices,
		HTTPStatusCodeNoContent,
		HTTPStatusCodeNonAuthoritativeInformation,
		HTTPStatusCodeNotAcceptable,
		HTTPStatusCodeNotFound,
		HTTPStatusCodeNotImplemented,
		HTTPStatusCodeNotModified,
		HTTPStatusCodeOK,
		HTTPStatusCodePartialContent,
		HTTPStatusCodePaymentRequired,
		HTTPStatusCodePreconditionFailed,
		HTTPStatusCodeProxyAuthenticationRequired,
		HTTPStatusCodeRedirect,
		HTTPStatusCodeRedirectKeepVerb,
		HTTPStatusCodeRedirectMethod,
		HTTPStatusCodeRequestEntityTooLarge,
		HTTPStatusCodeRequestTimeout,
		HTTPStatusCodeRequestURITooLong,
		HTTPStatusCodeRequestedRangeNotSatisfiable,
		HTTPStatusCodeResetContent,
		HTTPStatusCodeSeeOther,
		HTTPStatusCodeServiceUnavailable,
		HTTPStatusCodeSwitchingProtocols,
		HTTPStatusCodeTemporaryRedirect,
		HTTPStatusCodeUnauthorized,
		HTTPStatusCodeUnsupportedMediaType,
		HTTPStatusCodeUnused,
		HTTPStatusCodeUpgradeRequired,
		HTTPStatusCodeUseProxy,
	}
}

// ToPtr returns a *HTTPStatusCode pointing to the current value.
func (c HTTPStatusCode) ToPtr() *HTTPStatusCode {
	return &c
}

// JobProvisioningState - The provisioning state of the resource.
type JobProvisioningState string

const (
	JobProvisioningStateFailed     JobProvisioningState = "Failed"
	JobProvisioningStateProcessing JobProvisioningState = "Processing"
	JobProvisioningStateSucceeded  JobProvisioningState = "Succeeded"
	JobProvisioningStateSuspended  JobProvisioningState = "Suspended"
)

// PossibleJobProvisioningStateValues returns the possible values for the JobProvisioningState const type.
func PossibleJobProvisioningStateValues() []JobProvisioningState {
	return []JobProvisioningState{
		JobProvisioningStateFailed,
		JobProvisioningStateProcessing,
		JobProvisioningStateSucceeded,
		JobProvisioningStateSuspended,
	}
}

// ToPtr returns a *JobProvisioningState pointing to the current value.
func (c JobProvisioningState) ToPtr() *JobProvisioningState {
	return &c
}

// JobStatus - Gets or sets the status of the job.
type JobStatus string

const (
	JobStatusActivating   JobStatus = "Activating"
	JobStatusBlocked      JobStatus = "Blocked"
	JobStatusCompleted    JobStatus = "Completed"
	JobStatusDisconnected JobStatus = "Disconnected"
	JobStatusFailed       JobStatus = "Failed"
	JobStatusNew          JobStatus = "New"
	JobStatusRemoving     JobStatus = "Removing"
	JobStatusResuming     JobStatus = "Resuming"
	JobStatusRunning      JobStatus = "Running"
	JobStatusStopped      JobStatus = "Stopped"
	JobStatusStopping     JobStatus = "Stopping"
	JobStatusSuspended    JobStatus = "Suspended"
	JobStatusSuspending   JobStatus = "Suspending"
)

// PossibleJobStatusValues returns the possible values for the JobStatus const type.
func PossibleJobStatusValues() []JobStatus {
	return []JobStatus{
		JobStatusActivating,
		JobStatusBlocked,
		JobStatusCompleted,
		JobStatusDisconnected,
		JobStatusFailed,
		JobStatusNew,
		JobStatusRemoving,
		JobStatusResuming,
		JobStatusRunning,
		JobStatusStopped,
		JobStatusStopping,
		JobStatusSuspended,
		JobStatusSuspending,
	}
}

// ToPtr returns a *JobStatus pointing to the current value.
func (c JobStatus) ToPtr() *JobStatus {
	return &c
}

// JobStreamType - Gets or sets the stream type.
type JobStreamType string

const (
	JobStreamTypeAny      JobStreamType = "Any"
	JobStreamTypeDebug    JobStreamType = "Debug"
	JobStreamTypeError    JobStreamType = "Error"
	JobStreamTypeOutput   JobStreamType = "Output"
	JobStreamTypeProgress JobStreamType = "Progress"
	JobStreamTypeVerbose  JobStreamType = "Verbose"
	JobStreamTypeWarning  JobStreamType = "Warning"
)

// PossibleJobStreamTypeValues returns the possible values for the JobStreamType const type.
func PossibleJobStreamTypeValues() []JobStreamType {
	return []JobStreamType{
		JobStreamTypeAny,
		JobStreamTypeDebug,
		JobStreamTypeError,
		JobStreamTypeOutput,
		JobStreamTypeProgress,
		JobStreamTypeVerbose,
		JobStreamTypeWarning,
	}
}

// ToPtr returns a *JobStreamType pointing to the current value.
func (c JobStreamType) ToPtr() *JobStreamType {
	return &c
}

// LinuxUpdateClasses - Update classifications included in the software update configuration.
type LinuxUpdateClasses string

const (
	LinuxUpdateClassesCritical     LinuxUpdateClasses = "Critical"
	LinuxUpdateClassesOther        LinuxUpdateClasses = "Other"
	LinuxUpdateClassesSecurity     LinuxUpdateClasses = "Security"
	LinuxUpdateClassesUnclassified LinuxUpdateClasses = "Unclassified"
)

// PossibleLinuxUpdateClassesValues returns the possible values for the LinuxUpdateClasses const type.
func PossibleLinuxUpdateClassesValues() []LinuxUpdateClasses {
	return []LinuxUpdateClasses{
		LinuxUpdateClassesCritical,
		LinuxUpdateClassesOther,
		LinuxUpdateClassesSecurity,
		LinuxUpdateClassesUnclassified,
	}
}

// ToPtr returns a *LinuxUpdateClasses pointing to the current value.
func (c LinuxUpdateClasses) ToPtr() *LinuxUpdateClasses {
	return &c
}

// ModuleProvisioningState - Gets or sets the provisioning state of the module.
type ModuleProvisioningState string

const (
	ModuleProvisioningStateCreated                     ModuleProvisioningState = "Created"
	ModuleProvisioningStateCreating                    ModuleProvisioningState = "Creating"
	ModuleProvisioningStateStartingImportModuleRunbook ModuleProvisioningState = "StartingImportModuleRunbook"
	ModuleProvisioningStateRunningImportModuleRunbook  ModuleProvisioningState = "RunningImportModuleRunbook"
	ModuleProvisioningStateContentRetrieved            ModuleProvisioningState = "ContentRetrieved"
	ModuleProvisioningStateContentDownloaded           ModuleProvisioningState = "ContentDownloaded"
	ModuleProvisioningStateContentValidated            ModuleProvisioningState = "ContentValidated"
	ModuleProvisioningStateConnectionTypeImported      ModuleProvisioningState = "ConnectionTypeImported"
	ModuleProvisioningStateContentStored               ModuleProvisioningState = "ContentStored"
	ModuleProvisioningStateModuleDataStored            ModuleProvisioningState = "ModuleDataStored"
	ModuleProvisioningStateActivitiesStored            ModuleProvisioningState = "ActivitiesStored"
	ModuleProvisioningStateModuleImportRunbookComplete ModuleProvisioningState = "ModuleImportRunbookComplete"
	ModuleProvisioningStateSucceeded                   ModuleProvisioningState = "Succeeded"
	ModuleProvisioningStateFailed                      ModuleProvisioningState = "Failed"
	ModuleProvisioningStateCancelled                   ModuleProvisioningState = "Cancelled"
	ModuleProvisioningStateUpdating                    ModuleProvisioningState = "Updating"
)

// PossibleModuleProvisioningStateValues returns the possible values for the ModuleProvisioningState const type.
func PossibleModuleProvisioningStateValues() []ModuleProvisioningState {
	return []ModuleProvisioningState{
		ModuleProvisioningStateCreated,
		ModuleProvisioningStateCreating,
		ModuleProvisioningStateStartingImportModuleRunbook,
		ModuleProvisioningStateRunningImportModuleRunbook,
		ModuleProvisioningStateContentRetrieved,
		ModuleProvisioningStateContentDownloaded,
		ModuleProvisioningStateContentValidated,
		ModuleProvisioningStateConnectionTypeImported,
		ModuleProvisioningStateContentStored,
		ModuleProvisioningStateModuleDataStored,
		ModuleProvisioningStateActivitiesStored,
		ModuleProvisioningStateModuleImportRunbookComplete,
		ModuleProvisioningStateSucceeded,
		ModuleProvisioningStateFailed,
		ModuleProvisioningStateCancelled,
		ModuleProvisioningStateUpdating,
	}
}

// ToPtr returns a *ModuleProvisioningState pointing to the current value.
func (c ModuleProvisioningState) ToPtr() *ModuleProvisioningState {
	return &c
}

// OperatingSystemType - Target operating system for the software update configuration.
type OperatingSystemType string

const (
	OperatingSystemTypeWindows OperatingSystemType = "Windows"
	OperatingSystemTypeLinux   OperatingSystemType = "Linux"
)

// PossibleOperatingSystemTypeValues returns the possible values for the OperatingSystemType const type.
func PossibleOperatingSystemTypeValues() []OperatingSystemType {
	return []OperatingSystemType{
		OperatingSystemTypeWindows,
		OperatingSystemTypeLinux,
	}
}

// ToPtr returns a *OperatingSystemType pointing to the current value.
func (c OperatingSystemType) ToPtr() *OperatingSystemType {
	return &c
}

// ProvisioningState - The provisioning state of the job.
type ProvisioningState string

const (
	ProvisioningStateCompleted ProvisioningState = "Completed"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateRunning   ProvisioningState = "Running"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateCompleted,
		ProvisioningStateFailed,
		ProvisioningStateRunning,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// ResourceIdentityType - The identity type.
type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned             ResourceIdentityType = "SystemAssigned"
	ResourceIdentityTypeUserAssigned               ResourceIdentityType = "UserAssigned"
	ResourceIdentityTypeSystemAssignedUserAssigned ResourceIdentityType = "SystemAssigned, UserAssigned"
	ResourceIdentityTypeNone                       ResourceIdentityType = "None"
)

// PossibleResourceIdentityTypeValues returns the possible values for the ResourceIdentityType const type.
func PossibleResourceIdentityTypeValues() []ResourceIdentityType {
	return []ResourceIdentityType{
		ResourceIdentityTypeSystemAssigned,
		ResourceIdentityTypeUserAssigned,
		ResourceIdentityTypeSystemAssignedUserAssigned,
		ResourceIdentityTypeNone,
	}
}

// ToPtr returns a *ResourceIdentityType pointing to the current value.
func (c ResourceIdentityType) ToPtr() *ResourceIdentityType {
	return &c
}

// RunbookState - Gets or sets the state of the runbook.
type RunbookState string

const (
	RunbookStateEdit      RunbookState = "Edit"
	RunbookStateNew       RunbookState = "New"
	RunbookStatePublished RunbookState = "Published"
)

// PossibleRunbookStateValues returns the possible values for the RunbookState const type.
func PossibleRunbookStateValues() []RunbookState {
	return []RunbookState{
		RunbookStateEdit,
		RunbookStateNew,
		RunbookStatePublished,
	}
}

// ToPtr returns a *RunbookState pointing to the current value.
func (c RunbookState) ToPtr() *RunbookState {
	return &c
}

// RunbookTypeEnum - Gets or sets the type of the runbook.
type RunbookTypeEnum string

const (
	RunbookTypeEnumGraph                   RunbookTypeEnum = "Graph"
	RunbookTypeEnumGraphPowerShell         RunbookTypeEnum = "GraphPowerShell"
	RunbookTypeEnumGraphPowerShellWorkflow RunbookTypeEnum = "GraphPowerShellWorkflow"
	RunbookTypeEnumPowerShell              RunbookTypeEnum = "PowerShell"
	RunbookTypeEnumPowerShellWorkflow      RunbookTypeEnum = "PowerShellWorkflow"
	RunbookTypeEnumScript                  RunbookTypeEnum = "Script"
)

// PossibleRunbookTypeEnumValues returns the possible values for the RunbookTypeEnum const type.
func PossibleRunbookTypeEnumValues() []RunbookTypeEnum {
	return []RunbookTypeEnum{
		RunbookTypeEnumGraph,
		RunbookTypeEnumGraphPowerShell,
		RunbookTypeEnumGraphPowerShellWorkflow,
		RunbookTypeEnumPowerShell,
		RunbookTypeEnumPowerShellWorkflow,
		RunbookTypeEnumScript,
	}
}

// ToPtr returns a *RunbookTypeEnum pointing to the current value.
func (c RunbookTypeEnum) ToPtr() *RunbookTypeEnum {
	return &c
}

// SKUNameEnum - Gets or sets the SKU name of the account.
type SKUNameEnum string

const (
	SKUNameEnumBasic SKUNameEnum = "Basic"
	SKUNameEnumFree  SKUNameEnum = "Free"
)

// PossibleSKUNameEnumValues returns the possible values for the SKUNameEnum const type.
func PossibleSKUNameEnumValues() []SKUNameEnum {
	return []SKUNameEnum{
		SKUNameEnumBasic,
		SKUNameEnumFree,
	}
}

// ToPtr returns a *SKUNameEnum pointing to the current value.
func (c SKUNameEnum) ToPtr() *SKUNameEnum {
	return &c
}

// ScheduleDay - Day of the occurrence. Must be one of monday, tuesday, wednesday, thursday, friday, saturday, sunday.
type ScheduleDay string

const (
	ScheduleDayFriday    ScheduleDay = "Friday"
	ScheduleDayMonday    ScheduleDay = "Monday"
	ScheduleDaySaturday  ScheduleDay = "Saturday"
	ScheduleDaySunday    ScheduleDay = "Sunday"
	ScheduleDayThursday  ScheduleDay = "Thursday"
	ScheduleDayTuesday   ScheduleDay = "Tuesday"
	ScheduleDayWednesday ScheduleDay = "Wednesday"
)

// PossibleScheduleDayValues returns the possible values for the ScheduleDay const type.
func PossibleScheduleDayValues() []ScheduleDay {
	return []ScheduleDay{
		ScheduleDayFriday,
		ScheduleDayMonday,
		ScheduleDaySaturday,
		ScheduleDaySunday,
		ScheduleDayThursday,
		ScheduleDayTuesday,
		ScheduleDayWednesday,
	}
}

// ToPtr returns a *ScheduleDay pointing to the current value.
func (c ScheduleDay) ToPtr() *ScheduleDay {
	return &c
}

// ScheduleFrequency - Gets or sets the frequency of the schedule.
type ScheduleFrequency string

const (
	ScheduleFrequencyDay  ScheduleFrequency = "Day"
	ScheduleFrequencyHour ScheduleFrequency = "Hour"
	// ScheduleFrequencyMinute - The minimum allowed interval for Minute schedules is 15 minutes.
	ScheduleFrequencyMinute  ScheduleFrequency = "Minute"
	ScheduleFrequencyMonth   ScheduleFrequency = "Month"
	ScheduleFrequencyOneTime ScheduleFrequency = "OneTime"
	ScheduleFrequencyWeek    ScheduleFrequency = "Week"
)

// PossibleScheduleFrequencyValues returns the possible values for the ScheduleFrequency const type.
func PossibleScheduleFrequencyValues() []ScheduleFrequency {
	return []ScheduleFrequency{
		ScheduleFrequencyDay,
		ScheduleFrequencyHour,
		ScheduleFrequencyMinute,
		ScheduleFrequencyMonth,
		ScheduleFrequencyOneTime,
		ScheduleFrequencyWeek,
	}
}

// ToPtr returns a *ScheduleFrequency pointing to the current value.
func (c ScheduleFrequency) ToPtr() *ScheduleFrequency {
	return &c
}

// SourceType - The source type. Must be one of VsoGit, VsoTfvc, GitHub.
type SourceType string

const (
	SourceTypeGitHub  SourceType = "GitHub"
	SourceTypeVsoGit  SourceType = "VsoGit"
	SourceTypeVsoTfvc SourceType = "VsoTfvc"
)

// PossibleSourceTypeValues returns the possible values for the SourceType const type.
func PossibleSourceTypeValues() []SourceType {
	return []SourceType{
		SourceTypeGitHub,
		SourceTypeVsoGit,
		SourceTypeVsoTfvc,
	}
}

// ToPtr returns a *SourceType pointing to the current value.
func (c SourceType) ToPtr() *SourceType {
	return &c
}

// StreamType - The type of the sync job stream.
type StreamType string

const (
	StreamTypeError  StreamType = "Error"
	StreamTypeOutput StreamType = "Output"
)

// PossibleStreamTypeValues returns the possible values for the StreamType const type.
func PossibleStreamTypeValues() []StreamType {
	return []StreamType{
		StreamTypeError,
		StreamTypeOutput,
	}
}

// ToPtr returns a *StreamType pointing to the current value.
func (c StreamType) ToPtr() *StreamType {
	return &c
}

// SyncType - The sync type.
type SyncType string

const (
	SyncTypeFullSync    SyncType = "FullSync"
	SyncTypePartialSync SyncType = "PartialSync"
)

// PossibleSyncTypeValues returns the possible values for the SyncType const type.
func PossibleSyncTypeValues() []SyncType {
	return []SyncType{
		SyncTypeFullSync,
		SyncTypePartialSync,
	}
}

// ToPtr returns a *SyncType pointing to the current value.
func (c SyncType) ToPtr() *SyncType {
	return &c
}

// TagOperators - Filter VMs by Any or All specified tags.
type TagOperators string

const (
	TagOperatorsAll TagOperators = "All"
	TagOperatorsAny TagOperators = "Any"
)

// PossibleTagOperatorsValues returns the possible values for the TagOperators const type.
func PossibleTagOperatorsValues() []TagOperators {
	return []TagOperators{
		TagOperatorsAll,
		TagOperatorsAny,
	}
}

// ToPtr returns a *TagOperators pointing to the current value.
func (c TagOperators) ToPtr() *TagOperators {
	return &c
}

// TokenType - The token type. Must be either PersonalAccessToken or Oauth.
type TokenType string

const (
	TokenTypeOauth               TokenType = "Oauth"
	TokenTypePersonalAccessToken TokenType = "PersonalAccessToken"
)

// PossibleTokenTypeValues returns the possible values for the TokenType const type.
func PossibleTokenTypeValues() []TokenType {
	return []TokenType{
		TokenTypeOauth,
		TokenTypePersonalAccessToken,
	}
}

// ToPtr returns a *TokenType pointing to the current value.
func (c TokenType) ToPtr() *TokenType {
	return &c
}

// WindowsUpdateClasses - Update classification included in the software update configuration. A comma separated string with
// required values
type WindowsUpdateClasses string

const (
	WindowsUpdateClassesCritical     WindowsUpdateClasses = "Critical"
	WindowsUpdateClassesDefinition   WindowsUpdateClasses = "Definition"
	WindowsUpdateClassesFeaturePack  WindowsUpdateClasses = "FeaturePack"
	WindowsUpdateClassesSecurity     WindowsUpdateClasses = "Security"
	WindowsUpdateClassesServicePack  WindowsUpdateClasses = "ServicePack"
	WindowsUpdateClassesTools        WindowsUpdateClasses = "Tools"
	WindowsUpdateClassesUnclassified WindowsUpdateClasses = "Unclassified"
	WindowsUpdateClassesUpdateRollup WindowsUpdateClasses = "UpdateRollup"
	WindowsUpdateClassesUpdates      WindowsUpdateClasses = "Updates"
)

// PossibleWindowsUpdateClassesValues returns the possible values for the WindowsUpdateClasses const type.
func PossibleWindowsUpdateClassesValues() []WindowsUpdateClasses {
	return []WindowsUpdateClasses{
		WindowsUpdateClassesCritical,
		WindowsUpdateClassesDefinition,
		WindowsUpdateClassesFeaturePack,
		WindowsUpdateClassesSecurity,
		WindowsUpdateClassesServicePack,
		WindowsUpdateClassesTools,
		WindowsUpdateClassesUnclassified,
		WindowsUpdateClassesUpdateRollup,
		WindowsUpdateClassesUpdates,
	}
}

// ToPtr returns a *WindowsUpdateClasses pointing to the current value.
func (c WindowsUpdateClasses) ToPtr() *WindowsUpdateClasses {
	return &c
}

// WorkerType - Type of the HybridWorker.
type WorkerType string

const (
	WorkerTypeHybridV1 WorkerType = "HybridV1"
	WorkerTypeHybridV2 WorkerType = "HybridV2"
)

// PossibleWorkerTypeValues returns the possible values for the WorkerType const type.
func PossibleWorkerTypeValues() []WorkerType {
	return []WorkerType{
		WorkerTypeHybridV1,
		WorkerTypeHybridV2,
	}
}

// ToPtr returns a *WorkerType pointing to the current value.
func (c WorkerType) ToPtr() *WorkerType {
	return &c
}
