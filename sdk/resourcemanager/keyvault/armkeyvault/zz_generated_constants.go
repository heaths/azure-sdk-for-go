//go:build go1.16
// +build go1.16

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armkeyvault

const (
	moduleName    = "armkeyvault"
	moduleVersion = "v0.3.0"
)

type AccessPolicyUpdateKind string

const (
	AccessPolicyUpdateKindAdd     AccessPolicyUpdateKind = "add"
	AccessPolicyUpdateKindReplace AccessPolicyUpdateKind = "replace"
	AccessPolicyUpdateKindRemove  AccessPolicyUpdateKind = "remove"
)

// PossibleAccessPolicyUpdateKindValues returns the possible values for the AccessPolicyUpdateKind const type.
func PossibleAccessPolicyUpdateKindValues() []AccessPolicyUpdateKind {
	return []AccessPolicyUpdateKind{
		AccessPolicyUpdateKindAdd,
		AccessPolicyUpdateKindReplace,
		AccessPolicyUpdateKindRemove,
	}
}

// ToPtr returns a *AccessPolicyUpdateKind pointing to the current value.
func (c AccessPolicyUpdateKind) ToPtr() *AccessPolicyUpdateKind {
	return &c
}

// ActionsRequired - A message indicating if changes on the service provider require any updates on the consumer.
type ActionsRequired string

const (
	ActionsRequiredNone ActionsRequired = "None"
)

// PossibleActionsRequiredValues returns the possible values for the ActionsRequired const type.
func PossibleActionsRequiredValues() []ActionsRequired {
	return []ActionsRequired{
		ActionsRequiredNone,
	}
}

// ToPtr returns a *ActionsRequired pointing to the current value.
func (c ActionsRequired) ToPtr() *ActionsRequired {
	return &c
}

type CertificatePermissions string

const (
	CertificatePermissionsAll            CertificatePermissions = "all"
	CertificatePermissionsBackup         CertificatePermissions = "backup"
	CertificatePermissionsCreate         CertificatePermissions = "create"
	CertificatePermissionsDelete         CertificatePermissions = "delete"
	CertificatePermissionsDeleteissuers  CertificatePermissions = "deleteissuers"
	CertificatePermissionsGet            CertificatePermissions = "get"
	CertificatePermissionsGetissuers     CertificatePermissions = "getissuers"
	CertificatePermissionsImport         CertificatePermissions = "import"
	CertificatePermissionsList           CertificatePermissions = "list"
	CertificatePermissionsListissuers    CertificatePermissions = "listissuers"
	CertificatePermissionsManagecontacts CertificatePermissions = "managecontacts"
	CertificatePermissionsManageissuers  CertificatePermissions = "manageissuers"
	CertificatePermissionsPurge          CertificatePermissions = "purge"
	CertificatePermissionsRecover        CertificatePermissions = "recover"
	CertificatePermissionsRestore        CertificatePermissions = "restore"
	CertificatePermissionsSetissuers     CertificatePermissions = "setissuers"
	CertificatePermissionsUpdate         CertificatePermissions = "update"
)

// PossibleCertificatePermissionsValues returns the possible values for the CertificatePermissions const type.
func PossibleCertificatePermissionsValues() []CertificatePermissions {
	return []CertificatePermissions{
		CertificatePermissionsAll,
		CertificatePermissionsBackup,
		CertificatePermissionsCreate,
		CertificatePermissionsDelete,
		CertificatePermissionsDeleteissuers,
		CertificatePermissionsGet,
		CertificatePermissionsGetissuers,
		CertificatePermissionsImport,
		CertificatePermissionsList,
		CertificatePermissionsListissuers,
		CertificatePermissionsManagecontacts,
		CertificatePermissionsManageissuers,
		CertificatePermissionsPurge,
		CertificatePermissionsRecover,
		CertificatePermissionsRestore,
		CertificatePermissionsSetissuers,
		CertificatePermissionsUpdate,
	}
}

// ToPtr returns a *CertificatePermissions pointing to the current value.
func (c CertificatePermissions) ToPtr() *CertificatePermissions {
	return &c
}

// CreateMode - The vault's create mode to indicate whether the vault need to be recovered or not.
type CreateMode string

const (
	CreateModeRecover CreateMode = "recover"
	CreateModeDefault CreateMode = "default"
)

// PossibleCreateModeValues returns the possible values for the CreateMode const type.
func PossibleCreateModeValues() []CreateMode {
	return []CreateMode{
		CreateModeRecover,
		CreateModeDefault,
	}
}

// ToPtr returns a *CreateMode pointing to the current value.
func (c CreateMode) ToPtr() *CreateMode {
	return &c
}

// DeletionRecoveryLevel - The deletion recovery level currently in effect for the object. If it contains 'Purgeable', then
// the object can be permanently deleted by a privileged user; otherwise, only the system can purge the
// object at the end of the retention interval.
type DeletionRecoveryLevel string

const (
	DeletionRecoveryLevelPurgeable                        DeletionRecoveryLevel = "Purgeable"
	DeletionRecoveryLevelRecoverable                      DeletionRecoveryLevel = "Recoverable"
	DeletionRecoveryLevelRecoverableProtectedSubscription DeletionRecoveryLevel = "Recoverable+ProtectedSubscription"
	DeletionRecoveryLevelRecoverablePurgeable             DeletionRecoveryLevel = "Recoverable+Purgeable"
)

// PossibleDeletionRecoveryLevelValues returns the possible values for the DeletionRecoveryLevel const type.
func PossibleDeletionRecoveryLevelValues() []DeletionRecoveryLevel {
	return []DeletionRecoveryLevel{
		DeletionRecoveryLevelPurgeable,
		DeletionRecoveryLevelRecoverable,
		DeletionRecoveryLevelRecoverableProtectedSubscription,
		DeletionRecoveryLevelRecoverablePurgeable,
	}
}

// ToPtr returns a *DeletionRecoveryLevel pointing to the current value.
func (c DeletionRecoveryLevel) ToPtr() *DeletionRecoveryLevel {
	return &c
}

// IdentityType - The type of identity.
type IdentityType string

const (
	IdentityTypeApplication     IdentityType = "Application"
	IdentityTypeKey             IdentityType = "Key"
	IdentityTypeManagedIdentity IdentityType = "ManagedIdentity"
	IdentityTypeUser            IdentityType = "User"
)

// PossibleIdentityTypeValues returns the possible values for the IdentityType const type.
func PossibleIdentityTypeValues() []IdentityType {
	return []IdentityType{
		IdentityTypeApplication,
		IdentityTypeKey,
		IdentityTypeManagedIdentity,
		IdentityTypeUser,
	}
}

// ToPtr returns a *IdentityType pointing to the current value.
func (c IdentityType) ToPtr() *IdentityType {
	return &c
}

// JSONWebKeyCurveName - The elliptic curve name. For valid values, see JsonWebKeyCurveName.
type JSONWebKeyCurveName string

const (
	JSONWebKeyCurveNameP256  JSONWebKeyCurveName = "P-256"
	JSONWebKeyCurveNameP256K JSONWebKeyCurveName = "P-256K"
	JSONWebKeyCurveNameP384  JSONWebKeyCurveName = "P-384"
	JSONWebKeyCurveNameP521  JSONWebKeyCurveName = "P-521"
)

// PossibleJSONWebKeyCurveNameValues returns the possible values for the JSONWebKeyCurveName const type.
func PossibleJSONWebKeyCurveNameValues() []JSONWebKeyCurveName {
	return []JSONWebKeyCurveName{
		JSONWebKeyCurveNameP256,
		JSONWebKeyCurveNameP256K,
		JSONWebKeyCurveNameP384,
		JSONWebKeyCurveNameP521,
	}
}

// ToPtr returns a *JSONWebKeyCurveName pointing to the current value.
func (c JSONWebKeyCurveName) ToPtr() *JSONWebKeyCurveName {
	return &c
}

// JSONWebKeyOperation - The permitted JSON web key operations of the key. For more information, see JsonWebKeyOperation.
type JSONWebKeyOperation string

const (
	JSONWebKeyOperationDecrypt   JSONWebKeyOperation = "decrypt"
	JSONWebKeyOperationEncrypt   JSONWebKeyOperation = "encrypt"
	JSONWebKeyOperationImport    JSONWebKeyOperation = "import"
	JSONWebKeyOperationRelease   JSONWebKeyOperation = "release"
	JSONWebKeyOperationSign      JSONWebKeyOperation = "sign"
	JSONWebKeyOperationUnwrapKey JSONWebKeyOperation = "unwrapKey"
	JSONWebKeyOperationVerify    JSONWebKeyOperation = "verify"
	JSONWebKeyOperationWrapKey   JSONWebKeyOperation = "wrapKey"
)

// PossibleJSONWebKeyOperationValues returns the possible values for the JSONWebKeyOperation const type.
func PossibleJSONWebKeyOperationValues() []JSONWebKeyOperation {
	return []JSONWebKeyOperation{
		JSONWebKeyOperationDecrypt,
		JSONWebKeyOperationEncrypt,
		JSONWebKeyOperationImport,
		JSONWebKeyOperationRelease,
		JSONWebKeyOperationSign,
		JSONWebKeyOperationUnwrapKey,
		JSONWebKeyOperationVerify,
		JSONWebKeyOperationWrapKey,
	}
}

// ToPtr returns a *JSONWebKeyOperation pointing to the current value.
func (c JSONWebKeyOperation) ToPtr() *JSONWebKeyOperation {
	return &c
}

// JSONWebKeyType - The type of the key. For valid values, see JsonWebKeyType.
type JSONWebKeyType string

const (
	JSONWebKeyTypeEC     JSONWebKeyType = "EC"
	JSONWebKeyTypeECHSM  JSONWebKeyType = "EC-HSM"
	JSONWebKeyTypeRSA    JSONWebKeyType = "RSA"
	JSONWebKeyTypeRSAHSM JSONWebKeyType = "RSA-HSM"
)

// PossibleJSONWebKeyTypeValues returns the possible values for the JSONWebKeyType const type.
func PossibleJSONWebKeyTypeValues() []JSONWebKeyType {
	return []JSONWebKeyType{
		JSONWebKeyTypeEC,
		JSONWebKeyTypeECHSM,
		JSONWebKeyTypeRSA,
		JSONWebKeyTypeRSAHSM,
	}
}

// ToPtr returns a *JSONWebKeyType pointing to the current value.
func (c JSONWebKeyType) ToPtr() *JSONWebKeyType {
	return &c
}

type KeyPermissions string

const (
	KeyPermissionsAll               KeyPermissions = "all"
	KeyPermissionsBackup            KeyPermissions = "backup"
	KeyPermissionsCreate            KeyPermissions = "create"
	KeyPermissionsDecrypt           KeyPermissions = "decrypt"
	KeyPermissionsDelete            KeyPermissions = "delete"
	KeyPermissionsEncrypt           KeyPermissions = "encrypt"
	KeyPermissionsGet               KeyPermissions = "get"
	KeyPermissionsGetrotationpolicy KeyPermissions = "getrotationpolicy"
	KeyPermissionsImport            KeyPermissions = "import"
	KeyPermissionsList              KeyPermissions = "list"
	KeyPermissionsPurge             KeyPermissions = "purge"
	KeyPermissionsRecover           KeyPermissions = "recover"
	KeyPermissionsRelease           KeyPermissions = "release"
	KeyPermissionsRestore           KeyPermissions = "restore"
	KeyPermissionsRotate            KeyPermissions = "rotate"
	KeyPermissionsSetrotationpolicy KeyPermissions = "setrotationpolicy"
	KeyPermissionsSign              KeyPermissions = "sign"
	KeyPermissionsUnwrapKey         KeyPermissions = "unwrapKey"
	KeyPermissionsUpdate            KeyPermissions = "update"
	KeyPermissionsVerify            KeyPermissions = "verify"
	KeyPermissionsWrapKey           KeyPermissions = "wrapKey"
)

// PossibleKeyPermissionsValues returns the possible values for the KeyPermissions const type.
func PossibleKeyPermissionsValues() []KeyPermissions {
	return []KeyPermissions{
		KeyPermissionsAll,
		KeyPermissionsBackup,
		KeyPermissionsCreate,
		KeyPermissionsDecrypt,
		KeyPermissionsDelete,
		KeyPermissionsEncrypt,
		KeyPermissionsGet,
		KeyPermissionsGetrotationpolicy,
		KeyPermissionsImport,
		KeyPermissionsList,
		KeyPermissionsPurge,
		KeyPermissionsRecover,
		KeyPermissionsRelease,
		KeyPermissionsRestore,
		KeyPermissionsRotate,
		KeyPermissionsSetrotationpolicy,
		KeyPermissionsSign,
		KeyPermissionsUnwrapKey,
		KeyPermissionsUpdate,
		KeyPermissionsVerify,
		KeyPermissionsWrapKey,
	}
}

// ToPtr returns a *KeyPermissions pointing to the current value.
func (c KeyPermissions) ToPtr() *KeyPermissions {
	return &c
}

// KeyRotationPolicyActionType - The type of action.
type KeyRotationPolicyActionType string

const (
	KeyRotationPolicyActionTypeRotate KeyRotationPolicyActionType = "rotate"
	KeyRotationPolicyActionTypeNotify KeyRotationPolicyActionType = "notify"
)

// PossibleKeyRotationPolicyActionTypeValues returns the possible values for the KeyRotationPolicyActionType const type.
func PossibleKeyRotationPolicyActionTypeValues() []KeyRotationPolicyActionType {
	return []KeyRotationPolicyActionType{
		KeyRotationPolicyActionTypeRotate,
		KeyRotationPolicyActionTypeNotify,
	}
}

// ToPtr returns a *KeyRotationPolicyActionType pointing to the current value.
func (c KeyRotationPolicyActionType) ToPtr() *KeyRotationPolicyActionType {
	return &c
}

// ManagedHsmSKUFamily - SKU Family of the managed HSM Pool
type ManagedHsmSKUFamily string

const (
	ManagedHsmSKUFamilyB ManagedHsmSKUFamily = "B"
)

// PossibleManagedHsmSKUFamilyValues returns the possible values for the ManagedHsmSKUFamily const type.
func PossibleManagedHsmSKUFamilyValues() []ManagedHsmSKUFamily {
	return []ManagedHsmSKUFamily{
		ManagedHsmSKUFamilyB,
	}
}

// ToPtr returns a *ManagedHsmSKUFamily pointing to the current value.
func (c ManagedHsmSKUFamily) ToPtr() *ManagedHsmSKUFamily {
	return &c
}

// ManagedHsmSKUName - SKU of the managed HSM Pool
type ManagedHsmSKUName string

const (
	ManagedHsmSKUNameStandardB1 ManagedHsmSKUName = "Standard_B1"
	ManagedHsmSKUNameCustomB32  ManagedHsmSKUName = "Custom_B32"
)

// PossibleManagedHsmSKUNameValues returns the possible values for the ManagedHsmSKUName const type.
func PossibleManagedHsmSKUNameValues() []ManagedHsmSKUName {
	return []ManagedHsmSKUName{
		ManagedHsmSKUNameStandardB1,
		ManagedHsmSKUNameCustomB32,
	}
}

// ToPtr returns a *ManagedHsmSKUName pointing to the current value.
func (c ManagedHsmSKUName) ToPtr() *ManagedHsmSKUName {
	return &c
}

// NetworkRuleAction - The default action when no rule from ipRules and from virtualNetworkRules match. This is only used
// after the bypass property has been evaluated.
type NetworkRuleAction string

const (
	NetworkRuleActionAllow NetworkRuleAction = "Allow"
	NetworkRuleActionDeny  NetworkRuleAction = "Deny"
)

// PossibleNetworkRuleActionValues returns the possible values for the NetworkRuleAction const type.
func PossibleNetworkRuleActionValues() []NetworkRuleAction {
	return []NetworkRuleAction{
		NetworkRuleActionAllow,
		NetworkRuleActionDeny,
	}
}

// ToPtr returns a *NetworkRuleAction pointing to the current value.
func (c NetworkRuleAction) ToPtr() *NetworkRuleAction {
	return &c
}

// NetworkRuleBypassOptions - Tells what traffic can bypass network rules. This can be 'AzureServices' or 'None'. If not specified
// the default is 'AzureServices'.
type NetworkRuleBypassOptions string

const (
	NetworkRuleBypassOptionsAzureServices NetworkRuleBypassOptions = "AzureServices"
	NetworkRuleBypassOptionsNone          NetworkRuleBypassOptions = "None"
)

// PossibleNetworkRuleBypassOptionsValues returns the possible values for the NetworkRuleBypassOptions const type.
func PossibleNetworkRuleBypassOptionsValues() []NetworkRuleBypassOptions {
	return []NetworkRuleBypassOptions{
		NetworkRuleBypassOptionsAzureServices,
		NetworkRuleBypassOptionsNone,
	}
}

// ToPtr returns a *NetworkRuleBypassOptions pointing to the current value.
func (c NetworkRuleBypassOptions) ToPtr() *NetworkRuleBypassOptions {
	return &c
}

// PrivateEndpointConnectionProvisioningState - The current provisioning state.
type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating     PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting     PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateDisconnected PrivateEndpointConnectionProvisioningState = "Disconnected"
	PrivateEndpointConnectionProvisioningStateFailed       PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded    PrivateEndpointConnectionProvisioningState = "Succeeded"
	PrivateEndpointConnectionProvisioningStateUpdating     PrivateEndpointConnectionProvisioningState = "Updating"
)

// PossiblePrivateEndpointConnectionProvisioningStateValues returns the possible values for the PrivateEndpointConnectionProvisioningState const type.
func PossiblePrivateEndpointConnectionProvisioningStateValues() []PrivateEndpointConnectionProvisioningState {
	return []PrivateEndpointConnectionProvisioningState{
		PrivateEndpointConnectionProvisioningStateCreating,
		PrivateEndpointConnectionProvisioningStateDeleting,
		PrivateEndpointConnectionProvisioningStateDisconnected,
		PrivateEndpointConnectionProvisioningStateFailed,
		PrivateEndpointConnectionProvisioningStateSucceeded,
		PrivateEndpointConnectionProvisioningStateUpdating,
	}
}

// ToPtr returns a *PrivateEndpointConnectionProvisioningState pointing to the current value.
func (c PrivateEndpointConnectionProvisioningState) ToPtr() *PrivateEndpointConnectionProvisioningState {
	return &c
}

// PrivateEndpointServiceConnectionStatus - The private endpoint connection status.
type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved     PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusDisconnected PrivateEndpointServiceConnectionStatus = "Disconnected"
	PrivateEndpointServiceConnectionStatusPending      PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected     PrivateEndpointServiceConnectionStatus = "Rejected"
)

// PossiblePrivateEndpointServiceConnectionStatusValues returns the possible values for the PrivateEndpointServiceConnectionStatus const type.
func PossiblePrivateEndpointServiceConnectionStatusValues() []PrivateEndpointServiceConnectionStatus {
	return []PrivateEndpointServiceConnectionStatus{
		PrivateEndpointServiceConnectionStatusApproved,
		PrivateEndpointServiceConnectionStatusDisconnected,
		PrivateEndpointServiceConnectionStatusPending,
		PrivateEndpointServiceConnectionStatusRejected,
	}
}

// ToPtr returns a *PrivateEndpointServiceConnectionStatus pointing to the current value.
func (c PrivateEndpointServiceConnectionStatus) ToPtr() *PrivateEndpointServiceConnectionStatus {
	return &c
}

// ProvisioningState - Provisioning state.
type ProvisioningState string

const (
	// ProvisioningStateActivated - The managed HSM pool is ready for normal use.
	ProvisioningStateActivated ProvisioningState = "Activated"
	// ProvisioningStateDeleting - The managed HSM Pool is currently being deleted.
	ProvisioningStateDeleting ProvisioningState = "Deleting"
	// ProvisioningStateFailed - Provisioning of the managed HSM Pool has failed.
	ProvisioningStateFailed ProvisioningState = "Failed"
	// ProvisioningStateProvisioning - The managed HSM Pool is currently being provisioned.
	ProvisioningStateProvisioning ProvisioningState = "Provisioning"
	// ProvisioningStateRestoring - The managed HSM pool is being restored from full HSM backup.
	ProvisioningStateRestoring ProvisioningState = "Restoring"
	// ProvisioningStateSecurityDomainRestore - The managed HSM pool is waiting for a security domain restore action.
	ProvisioningStateSecurityDomainRestore ProvisioningState = "SecurityDomainRestore"
	// ProvisioningStateSucceeded - The managed HSM Pool has been full provisioned.
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	// ProvisioningStateUpdating - The managed HSM Pool is currently being updated.
	ProvisioningStateUpdating ProvisioningState = "Updating"
)

// PossibleProvisioningStateValues returns the possible values for the ProvisioningState const type.
func PossibleProvisioningStateValues() []ProvisioningState {
	return []ProvisioningState{
		ProvisioningStateActivated,
		ProvisioningStateDeleting,
		ProvisioningStateFailed,
		ProvisioningStateProvisioning,
		ProvisioningStateRestoring,
		ProvisioningStateSecurityDomainRestore,
		ProvisioningStateSucceeded,
		ProvisioningStateUpdating,
	}
}

// ToPtr returns a *ProvisioningState pointing to the current value.
func (c ProvisioningState) ToPtr() *ProvisioningState {
	return &c
}

// PublicNetworkAccess - Control permission for data plane traffic coming from public networks while private endpoint is enabled.
type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

// PossiblePublicNetworkAccessValues returns the possible values for the PublicNetworkAccess const type.
func PossiblePublicNetworkAccessValues() []PublicNetworkAccess {
	return []PublicNetworkAccess{
		PublicNetworkAccessDisabled,
		PublicNetworkAccessEnabled,
	}
}

// ToPtr returns a *PublicNetworkAccess pointing to the current value.
func (c PublicNetworkAccess) ToPtr() *PublicNetworkAccess {
	return &c
}

// Reason - The reason that a vault name could not be used. The Reason element is only returned if NameAvailable is false.
type Reason string

const (
	ReasonAccountNameInvalid Reason = "AccountNameInvalid"
	ReasonAlreadyExists      Reason = "AlreadyExists"
)

// PossibleReasonValues returns the possible values for the Reason const type.
func PossibleReasonValues() []Reason {
	return []Reason{
		ReasonAccountNameInvalid,
		ReasonAlreadyExists,
	}
}

// ToPtr returns a *Reason pointing to the current value.
func (c Reason) ToPtr() *Reason {
	return &c
}

// SKUFamily - SKU family name
type SKUFamily string

const (
	SKUFamilyA SKUFamily = "A"
)

// PossibleSKUFamilyValues returns the possible values for the SKUFamily const type.
func PossibleSKUFamilyValues() []SKUFamily {
	return []SKUFamily{
		SKUFamilyA,
	}
}

// ToPtr returns a *SKUFamily pointing to the current value.
func (c SKUFamily) ToPtr() *SKUFamily {
	return &c
}

// SKUName - SKU name to specify whether the key vault is a standard vault or a premium vault.
type SKUName string

const (
	SKUNameStandard SKUName = "standard"
	SKUNamePremium  SKUName = "premium"
)

// PossibleSKUNameValues returns the possible values for the SKUName const type.
func PossibleSKUNameValues() []SKUName {
	return []SKUName{
		SKUNameStandard,
		SKUNamePremium,
	}
}

// ToPtr returns a *SKUName pointing to the current value.
func (c SKUName) ToPtr() *SKUName {
	return &c
}

type SecretPermissions string

const (
	SecretPermissionsAll     SecretPermissions = "all"
	SecretPermissionsBackup  SecretPermissions = "backup"
	SecretPermissionsDelete  SecretPermissions = "delete"
	SecretPermissionsGet     SecretPermissions = "get"
	SecretPermissionsList    SecretPermissions = "list"
	SecretPermissionsPurge   SecretPermissions = "purge"
	SecretPermissionsRecover SecretPermissions = "recover"
	SecretPermissionsRestore SecretPermissions = "restore"
	SecretPermissionsSet     SecretPermissions = "set"
)

// PossibleSecretPermissionsValues returns the possible values for the SecretPermissions const type.
func PossibleSecretPermissionsValues() []SecretPermissions {
	return []SecretPermissions{
		SecretPermissionsAll,
		SecretPermissionsBackup,
		SecretPermissionsDelete,
		SecretPermissionsGet,
		SecretPermissionsList,
		SecretPermissionsPurge,
		SecretPermissionsRecover,
		SecretPermissionsRestore,
		SecretPermissionsSet,
	}
}

// ToPtr returns a *SecretPermissions pointing to the current value.
func (c SecretPermissions) ToPtr() *SecretPermissions {
	return &c
}

type StoragePermissions string

const (
	StoragePermissionsAll           StoragePermissions = "all"
	StoragePermissionsBackup        StoragePermissions = "backup"
	StoragePermissionsDelete        StoragePermissions = "delete"
	StoragePermissionsDeletesas     StoragePermissions = "deletesas"
	StoragePermissionsGet           StoragePermissions = "get"
	StoragePermissionsGetsas        StoragePermissions = "getsas"
	StoragePermissionsList          StoragePermissions = "list"
	StoragePermissionsListsas       StoragePermissions = "listsas"
	StoragePermissionsPurge         StoragePermissions = "purge"
	StoragePermissionsRecover       StoragePermissions = "recover"
	StoragePermissionsRegeneratekey StoragePermissions = "regeneratekey"
	StoragePermissionsRestore       StoragePermissions = "restore"
	StoragePermissionsSet           StoragePermissions = "set"
	StoragePermissionsSetsas        StoragePermissions = "setsas"
	StoragePermissionsUpdate        StoragePermissions = "update"
)

// PossibleStoragePermissionsValues returns the possible values for the StoragePermissions const type.
func PossibleStoragePermissionsValues() []StoragePermissions {
	return []StoragePermissions{
		StoragePermissionsAll,
		StoragePermissionsBackup,
		StoragePermissionsDelete,
		StoragePermissionsDeletesas,
		StoragePermissionsGet,
		StoragePermissionsGetsas,
		StoragePermissionsList,
		StoragePermissionsListsas,
		StoragePermissionsPurge,
		StoragePermissionsRecover,
		StoragePermissionsRegeneratekey,
		StoragePermissionsRestore,
		StoragePermissionsSet,
		StoragePermissionsSetsas,
		StoragePermissionsUpdate,
	}
}

// ToPtr returns a *StoragePermissions pointing to the current value.
func (c StoragePermissions) ToPtr() *StoragePermissions {
	return &c
}

// VaultProvisioningState - Provisioning state of the vault.
type VaultProvisioningState string

const (
	VaultProvisioningStateRegisteringDNS VaultProvisioningState = "RegisteringDns"
	VaultProvisioningStateSucceeded      VaultProvisioningState = "Succeeded"
)

// PossibleVaultProvisioningStateValues returns the possible values for the VaultProvisioningState const type.
func PossibleVaultProvisioningStateValues() []VaultProvisioningState {
	return []VaultProvisioningState{
		VaultProvisioningStateRegisteringDNS,
		VaultProvisioningStateSucceeded,
	}
}

// ToPtr returns a *VaultProvisioningState pointing to the current value.
func (c VaultProvisioningState) ToPtr() *VaultProvisioningState {
	return &c
}
