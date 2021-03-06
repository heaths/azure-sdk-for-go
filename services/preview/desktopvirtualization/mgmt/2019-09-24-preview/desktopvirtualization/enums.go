package desktopvirtualization

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// ApplicationGroupType enumerates the values for application group type.
type ApplicationGroupType string

const (
	// ApplicationGroupTypeDesktop ...
	ApplicationGroupTypeDesktop ApplicationGroupType = "Desktop"
	// ApplicationGroupTypeRemoteApp ...
	ApplicationGroupTypeRemoteApp ApplicationGroupType = "RemoteApp"
)

// PossibleApplicationGroupTypeValues returns an array of possible values for the ApplicationGroupType const type.
func PossibleApplicationGroupTypeValues() []ApplicationGroupType {
	return []ApplicationGroupType{ApplicationGroupTypeDesktop, ApplicationGroupTypeRemoteApp}
}

// ApplicationType enumerates the values for application type.
type ApplicationType string

const (
	// ApplicationTypeDesktop ...
	ApplicationTypeDesktop ApplicationType = "Desktop"
	// ApplicationTypeRemoteApp ...
	ApplicationTypeRemoteApp ApplicationType = "RemoteApp"
)

// PossibleApplicationTypeValues returns an array of possible values for the ApplicationType const type.
func PossibleApplicationTypeValues() []ApplicationType {
	return []ApplicationType{ApplicationTypeDesktop, ApplicationTypeRemoteApp}
}

// CommandLineSetting enumerates the values for command line setting.
type CommandLineSetting string

const (
	// Allow ...
	Allow CommandLineSetting = "Allow"
	// DoNotAllow ...
	DoNotAllow CommandLineSetting = "DoNotAllow"
	// Require ...
	Require CommandLineSetting = "Require"
)

// PossibleCommandLineSettingValues returns an array of possible values for the CommandLineSetting const type.
func PossibleCommandLineSettingValues() []CommandLineSetting {
	return []CommandLineSetting{Allow, DoNotAllow, Require}
}

// HostPoolType enumerates the values for host pool type.
type HostPoolType string

const (
	// Personal ...
	Personal HostPoolType = "Personal"
	// Shared ...
	Shared HostPoolType = "Shared"
)

// PossibleHostPoolTypeValues returns an array of possible values for the HostPoolType const type.
func PossibleHostPoolTypeValues() []HostPoolType {
	return []HostPoolType{Personal, Shared}
}

// LoadBalancerType enumerates the values for load balancer type.
type LoadBalancerType string

const (
	// BreadthFirst ...
	BreadthFirst LoadBalancerType = "BreadthFirst"
	// DepthFirst ...
	DepthFirst LoadBalancerType = "DepthFirst"
	// Persistent ...
	Persistent LoadBalancerType = "Persistent"
)

// PossibleLoadBalancerTypeValues returns an array of possible values for the LoadBalancerType const type.
func PossibleLoadBalancerTypeValues() []LoadBalancerType {
	return []LoadBalancerType{BreadthFirst, DepthFirst, Persistent}
}

// PersonalDesktopAssignmentType enumerates the values for personal desktop assignment type.
type PersonalDesktopAssignmentType string

const (
	// Automatic ...
	Automatic PersonalDesktopAssignmentType = "Automatic"
	// Direct ...
	Direct PersonalDesktopAssignmentType = "Direct"
)

// PossiblePersonalDesktopAssignmentTypeValues returns an array of possible values for the PersonalDesktopAssignmentType const type.
func PossiblePersonalDesktopAssignmentTypeValues() []PersonalDesktopAssignmentType {
	return []PersonalDesktopAssignmentType{Automatic, Direct}
}

// PreferredAppGroupType enumerates the values for preferred app group type.
type PreferredAppGroupType string

const (
	// PreferredAppGroupTypeDesktop ...
	PreferredAppGroupTypeDesktop PreferredAppGroupType = "Desktop"
	// PreferredAppGroupTypeNone ...
	PreferredAppGroupTypeNone PreferredAppGroupType = "None"
	// PreferredAppGroupTypeRailApplications ...
	PreferredAppGroupTypeRailApplications PreferredAppGroupType = "RailApplications"
)

// PossiblePreferredAppGroupTypeValues returns an array of possible values for the PreferredAppGroupType const type.
func PossiblePreferredAppGroupTypeValues() []PreferredAppGroupType {
	return []PreferredAppGroupType{PreferredAppGroupTypeDesktop, PreferredAppGroupTypeNone, PreferredAppGroupTypeRailApplications}
}

// RegistrationTokenOperation enumerates the values for registration token operation.
type RegistrationTokenOperation string

const (
	// Delete ...
	Delete RegistrationTokenOperation = "Delete"
	// None ...
	None RegistrationTokenOperation = "None"
	// Update ...
	Update RegistrationTokenOperation = "Update"
)

// PossibleRegistrationTokenOperationValues returns an array of possible values for the RegistrationTokenOperation const type.
func PossibleRegistrationTokenOperationValues() []RegistrationTokenOperation {
	return []RegistrationTokenOperation{Delete, None, Update}
}

// SessionState enumerates the values for session state.
type SessionState string

const (
	// Active ...
	Active SessionState = "Active"
	// Disconnected ...
	Disconnected SessionState = "Disconnected"
	// LogOff ...
	LogOff SessionState = "LogOff"
	// Pending ...
	Pending SessionState = "Pending"
	// Unknown ...
	Unknown SessionState = "Unknown"
	// UserProfileDiskMounted ...
	UserProfileDiskMounted SessionState = "UserProfileDiskMounted"
)

// PossibleSessionStateValues returns an array of possible values for the SessionState const type.
func PossibleSessionStateValues() []SessionState {
	return []SessionState{Active, Disconnected, LogOff, Pending, Unknown, UserProfileDiskMounted}
}

// Status enumerates the values for status.
type Status string

const (
	// StatusAvailable ...
	StatusAvailable Status = "Available"
	// StatusDisconnected ...
	StatusDisconnected Status = "Disconnected"
	// StatusShutdown ...
	StatusShutdown Status = "Shutdown"
	// StatusUnavailable ...
	StatusUnavailable Status = "Unavailable"
	// StatusUpgradeFailed ...
	StatusUpgradeFailed Status = "UpgradeFailed"
	// StatusUpgrading ...
	StatusUpgrading Status = "Upgrading"
)

// PossibleStatusValues returns an array of possible values for the Status const type.
func PossibleStatusValues() []Status {
	return []Status{StatusAvailable, StatusDisconnected, StatusShutdown, StatusUnavailable, StatusUpgradeFailed, StatusUpgrading}
}

// UpdateState enumerates the values for update state.
type UpdateState string

const (
	// UpdateStateFailed ...
	UpdateStateFailed UpdateState = "Failed"
	// UpdateStateInitial ...
	UpdateStateInitial UpdateState = "Initial"
	// UpdateStatePending ...
	UpdateStatePending UpdateState = "Pending"
	// UpdateStateStarted ...
	UpdateStateStarted UpdateState = "Started"
	// UpdateStateSucceeded ...
	UpdateStateSucceeded UpdateState = "Succeeded"
)

// PossibleUpdateStateValues returns an array of possible values for the UpdateState const type.
func PossibleUpdateStateValues() []UpdateState {
	return []UpdateState{UpdateStateFailed, UpdateStateInitial, UpdateStatePending, UpdateStateStarted, UpdateStateSucceeded}
}
