# Release History

## 0.2.0 (2022-01-13)
### Breaking Changes

- Function `*ApplyUpdatesClient.CreateOrUpdateParent` parameter(s) have been changed from `(context.Context, string, string, string, string, string, string, *ApplyUpdatesCreateOrUpdateParentOptions)` to `(context.Context, string, string, string, string, string, string, *ApplyUpdatesClientCreateOrUpdateParentOptions)`
- Function `*ApplyUpdatesClient.CreateOrUpdateParent` return value(s) have been changed from `(ApplyUpdatesCreateOrUpdateParentResponse, error)` to `(ApplyUpdatesClientCreateOrUpdateParentResponse, error)`
- Function `*ApplyUpdatesClient.GetParent` parameter(s) have been changed from `(context.Context, string, string, string, string, string, string, string, *ApplyUpdatesGetParentOptions)` to `(context.Context, string, string, string, string, string, string, string, *ApplyUpdatesClientGetParentOptions)`
- Function `*ApplyUpdatesClient.GetParent` return value(s) have been changed from `(ApplyUpdatesGetParentResponse, error)` to `(ApplyUpdatesClientGetParentResponse, error)`
- Function `*ConfigurationAssignmentsClient.Delete` parameter(s) have been changed from `(context.Context, string, string, string, string, string, *ConfigurationAssignmentsDeleteOptions)` to `(context.Context, string, string, string, string, string, *ConfigurationAssignmentsClientDeleteOptions)`
- Function `*ConfigurationAssignmentsClient.Delete` return value(s) have been changed from `(ConfigurationAssignmentsDeleteResponse, error)` to `(ConfigurationAssignmentsClientDeleteResponse, error)`
- Function `*PublicMaintenanceConfigurationsClient.Get` parameter(s) have been changed from `(context.Context, string, *PublicMaintenanceConfigurationsGetOptions)` to `(context.Context, string, *PublicMaintenanceConfigurationsClientGetOptions)`
- Function `*PublicMaintenanceConfigurationsClient.Get` return value(s) have been changed from `(PublicMaintenanceConfigurationsGetResponse, error)` to `(PublicMaintenanceConfigurationsClientGetResponse, error)`
- Function `*ConfigurationAssignmentsClient.DeleteParent` parameter(s) have been changed from `(context.Context, string, string, string, string, string, string, string, *ConfigurationAssignmentsDeleteParentOptions)` to `(context.Context, string, string, string, string, string, string, string, *ConfigurationAssignmentsClientDeleteParentOptions)`
- Function `*ConfigurationAssignmentsClient.DeleteParent` return value(s) have been changed from `(ConfigurationAssignmentsDeleteParentResponse, error)` to `(ConfigurationAssignmentsClientDeleteParentResponse, error)`
- Function `*UpdatesClient.ListParent` parameter(s) have been changed from `(context.Context, string, string, string, string, string, string, *UpdatesListParentOptions)` to `(context.Context, string, string, string, string, string, string, *UpdatesClientListParentOptions)`
- Function `*UpdatesClient.ListParent` return value(s) have been changed from `(UpdatesListParentResponse, error)` to `(UpdatesClientListParentResponse, error)`
- Function `*PublicMaintenanceConfigurationsClient.List` parameter(s) have been changed from `(context.Context, *PublicMaintenanceConfigurationsListOptions)` to `(context.Context, *PublicMaintenanceConfigurationsClientListOptions)`
- Function `*PublicMaintenanceConfigurationsClient.List` return value(s) have been changed from `(PublicMaintenanceConfigurationsListResponse, error)` to `(PublicMaintenanceConfigurationsClientListResponse, error)`
- Function `*ApplyUpdatesClient.List` parameter(s) have been changed from `(context.Context, *ApplyUpdatesListOptions)` to `(context.Context, *ApplyUpdatesClientListOptions)`
- Function `*ApplyUpdatesClient.List` return value(s) have been changed from `(ApplyUpdatesListResponse, error)` to `(ApplyUpdatesClientListResponse, error)`
- Function `*ConfigurationAssignmentsClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, string, string, *ConfigurationAssignmentsGetOptions)` to `(context.Context, string, string, string, string, string, *ConfigurationAssignmentsClientGetOptions)`
- Function `*ConfigurationAssignmentsClient.Get` return value(s) have been changed from `(ConfigurationAssignmentsGetResponse, error)` to `(ConfigurationAssignmentsClientGetResponse, error)`
- Function `*ApplyUpdateForResourceGroupClient.List` parameter(s) have been changed from `(context.Context, string, *ApplyUpdateForResourceGroupListOptions)` to `(context.Context, string, *ApplyUpdateForResourceGroupClientListOptions)`
- Function `*ApplyUpdateForResourceGroupClient.List` return value(s) have been changed from `(ApplyUpdateForResourceGroupListResponse, error)` to `(ApplyUpdateForResourceGroupClientListResponse, error)`
- Function `*ApplyUpdatesClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, *ApplyUpdatesCreateOrUpdateOptions)` to `(context.Context, string, string, string, string, *ApplyUpdatesClientCreateOrUpdateOptions)`
- Function `*ApplyUpdatesClient.CreateOrUpdate` return value(s) have been changed from `(ApplyUpdatesCreateOrUpdateResponse, error)` to `(ApplyUpdatesClientCreateOrUpdateResponse, error)`
- Function `*ConfigurationAssignmentsWithinSubscriptionClient.List` parameter(s) have been changed from `(context.Context, *ConfigurationAssignmentsWithinSubscriptionListOptions)` to `(context.Context, *ConfigurationAssignmentsWithinSubscriptionClientListOptions)`
- Function `*ConfigurationAssignmentsWithinSubscriptionClient.List` return value(s) have been changed from `(ConfigurationAssignmentsWithinSubscriptionListResponse, error)` to `(ConfigurationAssignmentsWithinSubscriptionClientListResponse, error)`
- Function `*OperationsClient.List` parameter(s) have been changed from `(context.Context, *OperationsListOptions)` to `(context.Context, *OperationsClientListOptions)`
- Function `*OperationsClient.List` return value(s) have been changed from `(OperationsListResponse, error)` to `(OperationsClientListResponse, error)`
- Function `*ConfigurationAssignmentsClient.CreateOrUpdateParent` parameter(s) have been changed from `(context.Context, string, string, string, string, string, string, string, ConfigurationAssignment, *ConfigurationAssignmentsCreateOrUpdateParentOptions)` to `(context.Context, string, string, string, string, string, string, string, ConfigurationAssignment, *ConfigurationAssignmentsClientCreateOrUpdateParentOptions)`
- Function `*ConfigurationAssignmentsClient.CreateOrUpdateParent` return value(s) have been changed from `(ConfigurationAssignmentsCreateOrUpdateParentResponse, error)` to `(ConfigurationAssignmentsClientCreateOrUpdateParentResponse, error)`
- Function `*ConfigurationAssignmentsClient.GetParent` parameter(s) have been changed from `(context.Context, string, string, string, string, string, string, string, *ConfigurationAssignmentsGetParentOptions)` to `(context.Context, string, string, string, string, string, string, string, *ConfigurationAssignmentsClientGetParentOptions)`
- Function `*ConfigurationAssignmentsClient.GetParent` return value(s) have been changed from `(ConfigurationAssignmentsGetParentResponse, error)` to `(ConfigurationAssignmentsClientGetParentResponse, error)`
- Function `*ConfigurationAssignmentsClient.List` parameter(s) have been changed from `(context.Context, string, string, string, string, *ConfigurationAssignmentsListOptions)` to `(context.Context, string, string, string, string, *ConfigurationAssignmentsClientListOptions)`
- Function `*ConfigurationAssignmentsClient.List` return value(s) have been changed from `(ConfigurationAssignmentsListResponse, error)` to `(ConfigurationAssignmentsClientListResponse, error)`
- Function `*ApplyUpdatesClient.Get` parameter(s) have been changed from `(context.Context, string, string, string, string, string, *ApplyUpdatesGetOptions)` to `(context.Context, string, string, string, string, string, *ApplyUpdatesClientGetOptions)`
- Function `*ApplyUpdatesClient.Get` return value(s) have been changed from `(ApplyUpdatesGetResponse, error)` to `(ApplyUpdatesClientGetResponse, error)`
- Function `*ConfigurationAssignmentsClient.ListParent` parameter(s) have been changed from `(context.Context, string, string, string, string, string, string, *ConfigurationAssignmentsListParentOptions)` to `(context.Context, string, string, string, string, string, string, *ConfigurationAssignmentsClientListParentOptions)`
- Function `*ConfigurationAssignmentsClient.ListParent` return value(s) have been changed from `(ConfigurationAssignmentsListParentResponse, error)` to `(ConfigurationAssignmentsClientListParentResponse, error)`
- Function `*ConfigurationAssignmentsClient.CreateOrUpdate` parameter(s) have been changed from `(context.Context, string, string, string, string, string, ConfigurationAssignment, *ConfigurationAssignmentsCreateOrUpdateOptions)` to `(context.Context, string, string, string, string, string, ConfigurationAssignment, *ConfigurationAssignmentsClientCreateOrUpdateOptions)`
- Function `*ConfigurationAssignmentsClient.CreateOrUpdate` return value(s) have been changed from `(ConfigurationAssignmentsCreateOrUpdateResponse, error)` to `(ConfigurationAssignmentsClientCreateOrUpdateResponse, error)`
- Function `*UpdatesClient.List` parameter(s) have been changed from `(context.Context, string, string, string, string, *UpdatesListOptions)` to `(context.Context, string, string, string, string, *UpdatesClientListOptions)`
- Function `*UpdatesClient.List` return value(s) have been changed from `(UpdatesListResponse, error)` to `(UpdatesClientListResponse, error)`
- Type of `ListMaintenanceConfigurationsResult.Value` has been changed from `[]*MaintenanceConfiguration` to `[]*Configuration`
- Function `NewMaintenanceConfigurationsForResourceGroupClient` has been removed
- Function `Resource.MarshalJSON` has been removed
- Function `ApplyUpdate.MarshalJSON` has been removed
- Function `ConfigurationAssignment.MarshalJSON` has been removed
- Function `MaintenanceConfigurationProperties.MarshalJSON` has been removed
- Function `MaintenanceConfiguration.MarshalJSON` has been removed
- Function `*MaintenanceConfigurationsClient.Update` has been removed
- Function `*MaintenanceConfigurationsForResourceGroupClient.List` has been removed
- Function `MaintenanceError.Error` has been removed
- Function `*MaintenanceConfigurationsClient.CreateOrUpdate` has been removed
- Function `*MaintenanceConfigurationsClient.Delete` has been removed
- Function `NewMaintenanceConfigurationsClient` has been removed
- Function `*MaintenanceConfigurationsClient.Get` has been removed
- Function `*MaintenanceConfigurationsClient.List` has been removed
- Struct `ApplyUpdateForResourceGroupListOptions` has been removed
- Struct `ApplyUpdateForResourceGroupListResponse` has been removed
- Struct `ApplyUpdateForResourceGroupListResult` has been removed
- Struct `ApplyUpdatesCreateOrUpdateOptions` has been removed
- Struct `ApplyUpdatesCreateOrUpdateParentOptions` has been removed
- Struct `ApplyUpdatesCreateOrUpdateParentResponse` has been removed
- Struct `ApplyUpdatesCreateOrUpdateParentResult` has been removed
- Struct `ApplyUpdatesCreateOrUpdateResponse` has been removed
- Struct `ApplyUpdatesCreateOrUpdateResult` has been removed
- Struct `ApplyUpdatesGetOptions` has been removed
- Struct `ApplyUpdatesGetParentOptions` has been removed
- Struct `ApplyUpdatesGetParentResponse` has been removed
- Struct `ApplyUpdatesGetParentResult` has been removed
- Struct `ApplyUpdatesGetResponse` has been removed
- Struct `ApplyUpdatesGetResult` has been removed
- Struct `ApplyUpdatesListOptions` has been removed
- Struct `ApplyUpdatesListResponse` has been removed
- Struct `ApplyUpdatesListResult` has been removed
- Struct `ConfigurationAssignmentsCreateOrUpdateOptions` has been removed
- Struct `ConfigurationAssignmentsCreateOrUpdateParentOptions` has been removed
- Struct `ConfigurationAssignmentsCreateOrUpdateParentResponse` has been removed
- Struct `ConfigurationAssignmentsCreateOrUpdateParentResult` has been removed
- Struct `ConfigurationAssignmentsCreateOrUpdateResponse` has been removed
- Struct `ConfigurationAssignmentsCreateOrUpdateResult` has been removed
- Struct `ConfigurationAssignmentsDeleteOptions` has been removed
- Struct `ConfigurationAssignmentsDeleteParentOptions` has been removed
- Struct `ConfigurationAssignmentsDeleteParentResponse` has been removed
- Struct `ConfigurationAssignmentsDeleteParentResult` has been removed
- Struct `ConfigurationAssignmentsDeleteResponse` has been removed
- Struct `ConfigurationAssignmentsDeleteResult` has been removed
- Struct `ConfigurationAssignmentsGetOptions` has been removed
- Struct `ConfigurationAssignmentsGetParentOptions` has been removed
- Struct `ConfigurationAssignmentsGetParentResponse` has been removed
- Struct `ConfigurationAssignmentsGetParentResult` has been removed
- Struct `ConfigurationAssignmentsGetResponse` has been removed
- Struct `ConfigurationAssignmentsGetResult` has been removed
- Struct `ConfigurationAssignmentsListOptions` has been removed
- Struct `ConfigurationAssignmentsListParentOptions` has been removed
- Struct `ConfigurationAssignmentsListParentResponse` has been removed
- Struct `ConfigurationAssignmentsListParentResult` has been removed
- Struct `ConfigurationAssignmentsListResponse` has been removed
- Struct `ConfigurationAssignmentsListResult` has been removed
- Struct `ConfigurationAssignmentsWithinSubscriptionListOptions` has been removed
- Struct `ConfigurationAssignmentsWithinSubscriptionListResponse` has been removed
- Struct `ConfigurationAssignmentsWithinSubscriptionListResult` has been removed
- Struct `MaintenanceConfiguration` has been removed
- Struct `MaintenanceConfigurationProperties` has been removed
- Struct `MaintenanceConfigurationsClient` has been removed
- Struct `MaintenanceConfigurationsCreateOrUpdateOptions` has been removed
- Struct `MaintenanceConfigurationsCreateOrUpdateResponse` has been removed
- Struct `MaintenanceConfigurationsCreateOrUpdateResult` has been removed
- Struct `MaintenanceConfigurationsDeleteOptions` has been removed
- Struct `MaintenanceConfigurationsDeleteResponse` has been removed
- Struct `MaintenanceConfigurationsDeleteResult` has been removed
- Struct `MaintenanceConfigurationsForResourceGroupClient` has been removed
- Struct `MaintenanceConfigurationsForResourceGroupListOptions` has been removed
- Struct `MaintenanceConfigurationsForResourceGroupListResponse` has been removed
- Struct `MaintenanceConfigurationsForResourceGroupListResult` has been removed
- Struct `MaintenanceConfigurationsGetOptions` has been removed
- Struct `MaintenanceConfigurationsGetResponse` has been removed
- Struct `MaintenanceConfigurationsGetResult` has been removed
- Struct `MaintenanceConfigurationsListOptions` has been removed
- Struct `MaintenanceConfigurationsListResponse` has been removed
- Struct `MaintenanceConfigurationsListResult` has been removed
- Struct `MaintenanceConfigurationsUpdateOptions` has been removed
- Struct `MaintenanceConfigurationsUpdateResponse` has been removed
- Struct `MaintenanceConfigurationsUpdateResult` has been removed
- Struct `MaintenanceError` has been removed
- Struct `MaintenanceWindow` has been removed
- Struct `OperationsListOptions` has been removed
- Struct `OperationsListResponse` has been removed
- Struct `OperationsListResultEnvelope` has been removed
- Struct `PublicMaintenanceConfigurationsGetOptions` has been removed
- Struct `PublicMaintenanceConfigurationsGetResponse` has been removed
- Struct `PublicMaintenanceConfigurationsGetResult` has been removed
- Struct `PublicMaintenanceConfigurationsListOptions` has been removed
- Struct `PublicMaintenanceConfigurationsListResponse` has been removed
- Struct `PublicMaintenanceConfigurationsListResult` has been removed
- Struct `UpdatesListOptions` has been removed
- Struct `UpdatesListParentOptions` has been removed
- Struct `UpdatesListParentResponse` has been removed
- Struct `UpdatesListParentResult` has been removed
- Struct `UpdatesListResponse` has been removed
- Struct `UpdatesListResult` has been removed
- Field `Resource` of struct `ConfigurationAssignment` has been removed
- Field `Resource` of struct `ApplyUpdate` has been removed

### Features Added

- New function `*ConfigurationsClient.CreateOrUpdate(context.Context, string, string, Configuration, *ConfigurationsClientCreateOrUpdateOptions) (ConfigurationsClientCreateOrUpdateResponse, error)`
- New function `*ConfigurationsClient.List(context.Context, *ConfigurationsClientListOptions) (ConfigurationsClientListResponse, error)`
- New function `*ConfigurationsClient.Get(context.Context, string, string, *ConfigurationsClientGetOptions) (ConfigurationsClientGetResponse, error)`
- New function `*ConfigurationsForResourceGroupClient.List(context.Context, string, *ConfigurationsForResourceGroupClientListOptions) (ConfigurationsForResourceGroupClientListResponse, error)`
- New function `Configuration.MarshalJSON() ([]byte, error)`
- New function `NewConfigurationsClient(string, azcore.TokenCredential, *arm.ClientOptions) *ConfigurationsClient`
- New function `*ConfigurationsClient.Delete(context.Context, string, string, *ConfigurationsClientDeleteOptions) (ConfigurationsClientDeleteResponse, error)`
- New function `*ConfigurationsClient.Update(context.Context, string, string, Configuration, *ConfigurationsClientUpdateOptions) (ConfigurationsClientUpdateResponse, error)`
- New function `ConfigurationProperties.MarshalJSON() ([]byte, error)`
- New function `NewConfigurationsForResourceGroupClient(string, azcore.TokenCredential, *arm.ClientOptions) *ConfigurationsForResourceGroupClient`
- New struct `ApplyUpdateForResourceGroupClientListOptions`
- New struct `ApplyUpdateForResourceGroupClientListResponse`
- New struct `ApplyUpdateForResourceGroupClientListResult`
- New struct `ApplyUpdatesClientCreateOrUpdateOptions`
- New struct `ApplyUpdatesClientCreateOrUpdateParentOptions`
- New struct `ApplyUpdatesClientCreateOrUpdateParentResponse`
- New struct `ApplyUpdatesClientCreateOrUpdateParentResult`
- New struct `ApplyUpdatesClientCreateOrUpdateResponse`
- New struct `ApplyUpdatesClientCreateOrUpdateResult`
- New struct `ApplyUpdatesClientGetOptions`
- New struct `ApplyUpdatesClientGetParentOptions`
- New struct `ApplyUpdatesClientGetParentResponse`
- New struct `ApplyUpdatesClientGetParentResult`
- New struct `ApplyUpdatesClientGetResponse`
- New struct `ApplyUpdatesClientGetResult`
- New struct `ApplyUpdatesClientListOptions`
- New struct `ApplyUpdatesClientListResponse`
- New struct `ApplyUpdatesClientListResult`
- New struct `Configuration`
- New struct `ConfigurationAssignmentsClientCreateOrUpdateOptions`
- New struct `ConfigurationAssignmentsClientCreateOrUpdateParentOptions`
- New struct `ConfigurationAssignmentsClientCreateOrUpdateParentResponse`
- New struct `ConfigurationAssignmentsClientCreateOrUpdateParentResult`
- New struct `ConfigurationAssignmentsClientCreateOrUpdateResponse`
- New struct `ConfigurationAssignmentsClientCreateOrUpdateResult`
- New struct `ConfigurationAssignmentsClientDeleteOptions`
- New struct `ConfigurationAssignmentsClientDeleteParentOptions`
- New struct `ConfigurationAssignmentsClientDeleteParentResponse`
- New struct `ConfigurationAssignmentsClientDeleteParentResult`
- New struct `ConfigurationAssignmentsClientDeleteResponse`
- New struct `ConfigurationAssignmentsClientDeleteResult`
- New struct `ConfigurationAssignmentsClientGetOptions`
- New struct `ConfigurationAssignmentsClientGetParentOptions`
- New struct `ConfigurationAssignmentsClientGetParentResponse`
- New struct `ConfigurationAssignmentsClientGetParentResult`
- New struct `ConfigurationAssignmentsClientGetResponse`
- New struct `ConfigurationAssignmentsClientGetResult`
- New struct `ConfigurationAssignmentsClientListOptions`
- New struct `ConfigurationAssignmentsClientListParentOptions`
- New struct `ConfigurationAssignmentsClientListParentResponse`
- New struct `ConfigurationAssignmentsClientListParentResult`
- New struct `ConfigurationAssignmentsClientListResponse`
- New struct `ConfigurationAssignmentsClientListResult`
- New struct `ConfigurationAssignmentsWithinSubscriptionClientListOptions`
- New struct `ConfigurationAssignmentsWithinSubscriptionClientListResponse`
- New struct `ConfigurationAssignmentsWithinSubscriptionClientListResult`
- New struct `ConfigurationProperties`
- New struct `ConfigurationsClient`
- New struct `ConfigurationsClientCreateOrUpdateOptions`
- New struct `ConfigurationsClientCreateOrUpdateResponse`
- New struct `ConfigurationsClientCreateOrUpdateResult`
- New struct `ConfigurationsClientDeleteOptions`
- New struct `ConfigurationsClientDeleteResponse`
- New struct `ConfigurationsClientDeleteResult`
- New struct `ConfigurationsClientGetOptions`
- New struct `ConfigurationsClientGetResponse`
- New struct `ConfigurationsClientGetResult`
- New struct `ConfigurationsClientListOptions`
- New struct `ConfigurationsClientListResponse`
- New struct `ConfigurationsClientListResult`
- New struct `ConfigurationsClientUpdateOptions`
- New struct `ConfigurationsClientUpdateResponse`
- New struct `ConfigurationsClientUpdateResult`
- New struct `ConfigurationsForResourceGroupClient`
- New struct `ConfigurationsForResourceGroupClientListOptions`
- New struct `ConfigurationsForResourceGroupClientListResponse`
- New struct `ConfigurationsForResourceGroupClientListResult`
- New struct `Error`
- New struct `OperationsClientListOptions`
- New struct `OperationsClientListResponse`
- New struct `OperationsClientListResult`
- New struct `PublicMaintenanceConfigurationsClientGetOptions`
- New struct `PublicMaintenanceConfigurationsClientGetResponse`
- New struct `PublicMaintenanceConfigurationsClientGetResult`
- New struct `PublicMaintenanceConfigurationsClientListOptions`
- New struct `PublicMaintenanceConfigurationsClientListResponse`
- New struct `PublicMaintenanceConfigurationsClientListResult`
- New struct `UpdatesClientListOptions`
- New struct `UpdatesClientListParentOptions`
- New struct `UpdatesClientListParentResponse`
- New struct `UpdatesClientListParentResult`
- New struct `UpdatesClientListResponse`
- New struct `UpdatesClientListResult`
- New struct `Window`
- New field `SystemData` in struct `ApplyUpdate`
- New field `Type` in struct `ApplyUpdate`
- New field `ID` in struct `ApplyUpdate`
- New field `Name` in struct `ApplyUpdate`
- New field `ID` in struct `ConfigurationAssignment`
- New field `Name` in struct `ConfigurationAssignment`
- New field `SystemData` in struct `ConfigurationAssignment`
- New field `Type` in struct `ConfigurationAssignment`


## 0.1.0 (2021-12-07)

- Init release.