// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AutoScalingMetric string

// Enum values for AutoScalingMetric
const (
	AutoScalingMetricCpuUtilizationPercentage AutoScalingMetric = "CPU_UTILIZATION_PERCENTAGE"
)

// Values returns all known values for AutoScalingMetric. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AutoScalingMetric) Values() []AutoScalingMetric {
	return []AutoScalingMetric{
		"CPU_UTILIZATION_PERCENTAGE",
	}
}

type ChangesetStatus string

// Enum values for ChangesetStatus
const (
	ChangesetStatusPending    ChangesetStatus = "PENDING"
	ChangesetStatusProcessing ChangesetStatus = "PROCESSING"
	ChangesetStatusFailed     ChangesetStatus = "FAILED"
	ChangesetStatusCompleted  ChangesetStatus = "COMPLETED"
)

// Values returns all known values for ChangesetStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChangesetStatus) Values() []ChangesetStatus {
	return []ChangesetStatus{
		"PENDING",
		"PROCESSING",
		"FAILED",
		"COMPLETED",
	}
}

type ChangeType string

// Enum values for ChangeType
const (
	ChangeTypePut    ChangeType = "PUT"
	ChangeTypeDelete ChangeType = "DELETE"
)

// Values returns all known values for ChangeType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ChangeType) Values() []ChangeType {
	return []ChangeType{
		"PUT",
		"DELETE",
	}
}

type DnsStatus string

// Enum values for DnsStatus
const (
	DnsStatusNone                DnsStatus = "NONE"
	DnsStatusUpdateRequested     DnsStatus = "UPDATE_REQUESTED"
	DnsStatusUpdating            DnsStatus = "UPDATING"
	DnsStatusFailedUpdate        DnsStatus = "FAILED_UPDATE"
	DnsStatusSuccessfullyUpdated DnsStatus = "SUCCESSFULLY_UPDATED"
)

// Values returns all known values for DnsStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (DnsStatus) Values() []DnsStatus {
	return []DnsStatus{
		"NONE",
		"UPDATE_REQUESTED",
		"UPDATING",
		"FAILED_UPDATE",
		"SUCCESSFULLY_UPDATED",
	}
}

type EnvironmentStatus string

// Enum values for EnvironmentStatus
const (
	EnvironmentStatusCreateRequested        EnvironmentStatus = "CREATE_REQUESTED"
	EnvironmentStatusCreating               EnvironmentStatus = "CREATING"
	EnvironmentStatusCreated                EnvironmentStatus = "CREATED"
	EnvironmentStatusDeleteRequested        EnvironmentStatus = "DELETE_REQUESTED"
	EnvironmentStatusDeleting               EnvironmentStatus = "DELETING"
	EnvironmentStatusDeleted                EnvironmentStatus = "DELETED"
	EnvironmentStatusFailedCreation         EnvironmentStatus = "FAILED_CREATION"
	EnvironmentStatusRetryDeletion          EnvironmentStatus = "RETRY_DELETION"
	EnvironmentStatusFailedDeletion         EnvironmentStatus = "FAILED_DELETION"
	EnvironmentStatusUpdateNetworkRequested EnvironmentStatus = "UPDATE_NETWORK_REQUESTED"
	EnvironmentStatusUpdatingNetwork        EnvironmentStatus = "UPDATING_NETWORK"
	EnvironmentStatusFailedUpdatingNetwork  EnvironmentStatus = "FAILED_UPDATING_NETWORK"
	EnvironmentStatusSuspended              EnvironmentStatus = "SUSPENDED"
)

// Values returns all known values for EnvironmentStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EnvironmentStatus) Values() []EnvironmentStatus {
	return []EnvironmentStatus{
		"CREATE_REQUESTED",
		"CREATING",
		"CREATED",
		"DELETE_REQUESTED",
		"DELETING",
		"DELETED",
		"FAILED_CREATION",
		"RETRY_DELETION",
		"FAILED_DELETION",
		"UPDATE_NETWORK_REQUESTED",
		"UPDATING_NETWORK",
		"FAILED_UPDATING_NETWORK",
		"SUSPENDED",
	}
}

type ErrorDetails string

// Enum values for ErrorDetails
const (
	ErrorDetailsValidation               ErrorDetails = "The inputs to this request are invalid."
	ErrorDetailsServiceQuotaExceeded     ErrorDetails = "Service limits have been exceeded."
	ErrorDetailsAccessDenied             ErrorDetails = "Missing required permission to perform this request."
	ErrorDetailsResourceNotFound         ErrorDetails = "One or more inputs to this request were not found."
	ErrorDetailsThrottling               ErrorDetails = "The system temporarily lacks sufficient resources to process the request."
	ErrorDetailsInternalServiceException ErrorDetails = "An internal error has occurred."
	ErrorDetailsCancelled                ErrorDetails = "Cancelled"
	ErrorDetailsUserRecoverable          ErrorDetails = "A user recoverable error has occurred"
)

// Values returns all known values for ErrorDetails. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ErrorDetails) Values() []ErrorDetails {
	return []ErrorDetails{
		"The inputs to this request are invalid.",
		"Service limits have been exceeded.",
		"Missing required permission to perform this request.",
		"One or more inputs to this request were not found.",
		"The system temporarily lacks sufficient resources to process the request.",
		"An internal error has occurred.",
		"Cancelled",
		"A user recoverable error has occurred",
	}
}

type FederationMode string

// Enum values for FederationMode
const (
	FederationModeFederated FederationMode = "FEDERATED"
	FederationModeLocal     FederationMode = "LOCAL"
)

// Values returns all known values for FederationMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FederationMode) Values() []FederationMode {
	return []FederationMode{
		"FEDERATED",
		"LOCAL",
	}
}

type IPAddressType string

// Enum values for IPAddressType
const (
	IPAddressTypeIpV4 IPAddressType = "IP_V4"
)

// Values returns all known values for IPAddressType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IPAddressType) Values() []IPAddressType {
	return []IPAddressType{
		"IP_V4",
	}
}

type KxAzMode string

// Enum values for KxAzMode
const (
	KxAzModeSingle KxAzMode = "SINGLE"
	KxAzModeMulti  KxAzMode = "MULTI"
)

// Values returns all known values for KxAzMode. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (KxAzMode) Values() []KxAzMode {
	return []KxAzMode{
		"SINGLE",
		"MULTI",
	}
}

type KxClusterCodeDeploymentStrategy string

// Enum values for KxClusterCodeDeploymentStrategy
const (
	KxClusterCodeDeploymentStrategyRolling KxClusterCodeDeploymentStrategy = "ROLLING"
	KxClusterCodeDeploymentStrategyForce   KxClusterCodeDeploymentStrategy = "FORCE"
)

// Values returns all known values for KxClusterCodeDeploymentStrategy. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (KxClusterCodeDeploymentStrategy) Values() []KxClusterCodeDeploymentStrategy {
	return []KxClusterCodeDeploymentStrategy{
		"ROLLING",
		"FORCE",
	}
}

type KxClusterStatus string

// Enum values for KxClusterStatus
const (
	KxClusterStatusPending      KxClusterStatus = "PENDING"
	KxClusterStatusCreating     KxClusterStatus = "CREATING"
	KxClusterStatusCreateFailed KxClusterStatus = "CREATE_FAILED"
	KxClusterStatusRunning      KxClusterStatus = "RUNNING"
	KxClusterStatusUpdating     KxClusterStatus = "UPDATING"
	KxClusterStatusDeleting     KxClusterStatus = "DELETING"
	KxClusterStatusDeleted      KxClusterStatus = "DELETED"
	KxClusterStatusDeleteFailed KxClusterStatus = "DELETE_FAILED"
)

// Values returns all known values for KxClusterStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KxClusterStatus) Values() []KxClusterStatus {
	return []KxClusterStatus{
		"PENDING",
		"CREATING",
		"CREATE_FAILED",
		"RUNNING",
		"UPDATING",
		"DELETING",
		"DELETED",
		"DELETE_FAILED",
	}
}

type KxClusterType string

// Enum values for KxClusterType
const (
	KxClusterTypeHdb     KxClusterType = "HDB"
	KxClusterTypeRdb     KxClusterType = "RDB"
	KxClusterTypeGateway KxClusterType = "GATEWAY"
)

// Values returns all known values for KxClusterType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KxClusterType) Values() []KxClusterType {
	return []KxClusterType{
		"HDB",
		"RDB",
		"GATEWAY",
	}
}

type KxDeploymentStrategy string

// Enum values for KxDeploymentStrategy
const (
	KxDeploymentStrategyNoRestart KxDeploymentStrategy = "NO_RESTART"
	KxDeploymentStrategyRolling   KxDeploymentStrategy = "ROLLING"
)

// Values returns all known values for KxDeploymentStrategy. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KxDeploymentStrategy) Values() []KxDeploymentStrategy {
	return []KxDeploymentStrategy{
		"NO_RESTART",
		"ROLLING",
	}
}

type KxSavedownStorageType string

// Enum values for KxSavedownStorageType
const (
	KxSavedownStorageTypeSds01 KxSavedownStorageType = "SDS01"
)

// Values returns all known values for KxSavedownStorageType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (KxSavedownStorageType) Values() []KxSavedownStorageType {
	return []KxSavedownStorageType{
		"SDS01",
	}
}

type RuleAction string

// Enum values for RuleAction
const (
	RuleActionAllow RuleAction = "allow"
	RuleActionDeny  RuleAction = "deny"
)

// Values returns all known values for RuleAction. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (RuleAction) Values() []RuleAction {
	return []RuleAction{
		"allow",
		"deny",
	}
}

type TgwStatus string

// Enum values for TgwStatus
const (
	TgwStatusNone                TgwStatus = "NONE"
	TgwStatusUpdateRequested     TgwStatus = "UPDATE_REQUESTED"
	TgwStatusUpdating            TgwStatus = "UPDATING"
	TgwStatusFailedUpdate        TgwStatus = "FAILED_UPDATE"
	TgwStatusSuccessfullyUpdated TgwStatus = "SUCCESSFULLY_UPDATED"
)

// Values returns all known values for TgwStatus. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TgwStatus) Values() []TgwStatus {
	return []TgwStatus{
		"NONE",
		"UPDATE_REQUESTED",
		"UPDATING",
		"FAILED_UPDATE",
		"SUCCESSFULLY_UPDATED",
	}
}
