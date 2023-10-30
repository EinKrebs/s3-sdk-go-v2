// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type ActivityStreamMode string

// Enum values for ActivityStreamMode
const (
	ActivityStreamModeSync  ActivityStreamMode = "sync"
	ActivityStreamModeAsync ActivityStreamMode = "async"
)

// Values returns all known values for ActivityStreamMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActivityStreamMode) Values() []ActivityStreamMode {
	return []ActivityStreamMode{
		"sync",
		"async",
	}
}

type ActivityStreamPolicyStatus string

// Enum values for ActivityStreamPolicyStatus
const (
	ActivityStreamPolicyStatusLocked          ActivityStreamPolicyStatus = "locked"
	ActivityStreamPolicyStatusUnlocked        ActivityStreamPolicyStatus = "unlocked"
	ActivityStreamPolicyStatusLockingPolicy   ActivityStreamPolicyStatus = "locking-policy"
	ActivityStreamPolicyStatusUnlockingPolicy ActivityStreamPolicyStatus = "unlocking-policy"
)

// Values returns all known values for ActivityStreamPolicyStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ActivityStreamPolicyStatus) Values() []ActivityStreamPolicyStatus {
	return []ActivityStreamPolicyStatus{
		"locked",
		"unlocked",
		"locking-policy",
		"unlocking-policy",
	}
}

type ActivityStreamStatus string

// Enum values for ActivityStreamStatus
const (
	ActivityStreamStatusStopped  ActivityStreamStatus = "stopped"
	ActivityStreamStatusStarting ActivityStreamStatus = "starting"
	ActivityStreamStatusStarted  ActivityStreamStatus = "started"
	ActivityStreamStatusStopping ActivityStreamStatus = "stopping"
)

// Values returns all known values for ActivityStreamStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ActivityStreamStatus) Values() []ActivityStreamStatus {
	return []ActivityStreamStatus{
		"stopped",
		"starting",
		"started",
		"stopping",
	}
}

type ApplyMethod string

// Enum values for ApplyMethod
const (
	ApplyMethodImmediate     ApplyMethod = "immediate"
	ApplyMethodPendingReboot ApplyMethod = "pending-reboot"
)

// Values returns all known values for ApplyMethod. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ApplyMethod) Values() []ApplyMethod {
	return []ApplyMethod{
		"immediate",
		"pending-reboot",
	}
}

type AuditPolicyState string

// Enum values for AuditPolicyState
const (
	AuditPolicyStateLockedPolicy   AuditPolicyState = "locked"
	AuditPolicyStateUnlockedPolicy AuditPolicyState = "unlocked"
)

// Values returns all known values for AuditPolicyState. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AuditPolicyState) Values() []AuditPolicyState {
	return []AuditPolicyState{
		"locked",
		"unlocked",
	}
}

type AuthScheme string

// Enum values for AuthScheme
const (
	AuthSchemeSecrets AuthScheme = "SECRETS"
)

// Values returns all known values for AuthScheme. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (AuthScheme) Values() []AuthScheme {
	return []AuthScheme{
		"SECRETS",
	}
}

type AutomationMode string

// Enum values for AutomationMode
const (
	AutomationModeFull      AutomationMode = "full"
	AutomationModeAllPaused AutomationMode = "all-paused"
)

// Values returns all known values for AutomationMode. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (AutomationMode) Values() []AutomationMode {
	return []AutomationMode{
		"full",
		"all-paused",
	}
}

type ClientPasswordAuthType string

// Enum values for ClientPasswordAuthType
const (
	ClientPasswordAuthTypeMysqlNativePassword     ClientPasswordAuthType = "MYSQL_NATIVE_PASSWORD"
	ClientPasswordAuthTypePostgresScramSha256     ClientPasswordAuthType = "POSTGRES_SCRAM_SHA_256"
	ClientPasswordAuthTypePostgresMd5             ClientPasswordAuthType = "POSTGRES_MD5"
	ClientPasswordAuthTypeSqlServerAuthentication ClientPasswordAuthType = "SQL_SERVER_AUTHENTICATION"
)

// Values returns all known values for ClientPasswordAuthType. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ClientPasswordAuthType) Values() []ClientPasswordAuthType {
	return []ClientPasswordAuthType{
		"MYSQL_NATIVE_PASSWORD",
		"POSTGRES_SCRAM_SHA_256",
		"POSTGRES_MD5",
		"SQL_SERVER_AUTHENTICATION",
	}
}

type CustomEngineVersionStatus string

// Enum values for CustomEngineVersionStatus
const (
	CustomEngineVersionStatusAvailable             CustomEngineVersionStatus = "available"
	CustomEngineVersionStatusInactive              CustomEngineVersionStatus = "inactive"
	CustomEngineVersionStatusInactiveExceptRestore CustomEngineVersionStatus = "inactive-except-restore"
)

// Values returns all known values for CustomEngineVersionStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (CustomEngineVersionStatus) Values() []CustomEngineVersionStatus {
	return []CustomEngineVersionStatus{
		"available",
		"inactive",
		"inactive-except-restore",
	}
}

type DBProxyEndpointStatus string

// Enum values for DBProxyEndpointStatus
const (
	DBProxyEndpointStatusAvailable                  DBProxyEndpointStatus = "available"
	DBProxyEndpointStatusModifying                  DBProxyEndpointStatus = "modifying"
	DBProxyEndpointStatusIncompatibleNetwork        DBProxyEndpointStatus = "incompatible-network"
	DBProxyEndpointStatusInsufficientResourceLimits DBProxyEndpointStatus = "insufficient-resource-limits"
	DBProxyEndpointStatusCreating                   DBProxyEndpointStatus = "creating"
	DBProxyEndpointStatusDeleting                   DBProxyEndpointStatus = "deleting"
)

// Values returns all known values for DBProxyEndpointStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DBProxyEndpointStatus) Values() []DBProxyEndpointStatus {
	return []DBProxyEndpointStatus{
		"available",
		"modifying",
		"incompatible-network",
		"insufficient-resource-limits",
		"creating",
		"deleting",
	}
}

type DBProxyEndpointTargetRole string

// Enum values for DBProxyEndpointTargetRole
const (
	DBProxyEndpointTargetRoleReadWrite DBProxyEndpointTargetRole = "READ_WRITE"
	DBProxyEndpointTargetRoleReadOnly  DBProxyEndpointTargetRole = "READ_ONLY"
)

// Values returns all known values for DBProxyEndpointTargetRole. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (DBProxyEndpointTargetRole) Values() []DBProxyEndpointTargetRole {
	return []DBProxyEndpointTargetRole{
		"READ_WRITE",
		"READ_ONLY",
	}
}

type DBProxyStatus string

// Enum values for DBProxyStatus
const (
	DBProxyStatusAvailable                  DBProxyStatus = "available"
	DBProxyStatusModifying                  DBProxyStatus = "modifying"
	DBProxyStatusIncompatibleNetwork        DBProxyStatus = "incompatible-network"
	DBProxyStatusInsufficientResourceLimits DBProxyStatus = "insufficient-resource-limits"
	DBProxyStatusCreating                   DBProxyStatus = "creating"
	DBProxyStatusDeleting                   DBProxyStatus = "deleting"
	DBProxyStatusSuspended                  DBProxyStatus = "suspended"
	DBProxyStatusSuspending                 DBProxyStatus = "suspending"
	DBProxyStatusReactivating               DBProxyStatus = "reactivating"
)

// Values returns all known values for DBProxyStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (DBProxyStatus) Values() []DBProxyStatus {
	return []DBProxyStatus{
		"available",
		"modifying",
		"incompatible-network",
		"insufficient-resource-limits",
		"creating",
		"deleting",
		"suspended",
		"suspending",
		"reactivating",
	}
}

type EngineFamily string

// Enum values for EngineFamily
const (
	EngineFamilyMysql      EngineFamily = "MYSQL"
	EngineFamilyPostgresql EngineFamily = "POSTGRESQL"
	EngineFamilySqlserver  EngineFamily = "SQLSERVER"
)

// Values returns all known values for EngineFamily. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (EngineFamily) Values() []EngineFamily {
	return []EngineFamily{
		"MYSQL",
		"POSTGRESQL",
		"SQLSERVER",
	}
}

type ExportSourceType string

// Enum values for ExportSourceType
const (
	ExportSourceTypeSnapshot ExportSourceType = "SNAPSHOT"
	ExportSourceTypeCluster  ExportSourceType = "CLUSTER"
)

// Values returns all known values for ExportSourceType. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ExportSourceType) Values() []ExportSourceType {
	return []ExportSourceType{
		"SNAPSHOT",
		"CLUSTER",
	}
}

type FailoverStatus string

// Enum values for FailoverStatus
const (
	FailoverStatusPending     FailoverStatus = "pending"
	FailoverStatusFailingOver FailoverStatus = "failing-over"
	FailoverStatusCancelling  FailoverStatus = "cancelling"
)

// Values returns all known values for FailoverStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (FailoverStatus) Values() []FailoverStatus {
	return []FailoverStatus{
		"pending",
		"failing-over",
		"cancelling",
	}
}

type GlobalClusterMemberSynchronizationStatus string

// Enum values for GlobalClusterMemberSynchronizationStatus
const (
	GlobalClusterMemberSynchronizationStatusConnected     GlobalClusterMemberSynchronizationStatus = "connected"
	GlobalClusterMemberSynchronizationStatusPendingResync GlobalClusterMemberSynchronizationStatus = "pending-resync"
)

// Values returns all known values for GlobalClusterMemberSynchronizationStatus.
// Note that this can be expanded in the future, and so it is only as up to date as
// the client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (GlobalClusterMemberSynchronizationStatus) Values() []GlobalClusterMemberSynchronizationStatus {
	return []GlobalClusterMemberSynchronizationStatus{
		"connected",
		"pending-resync",
	}
}

type IAMAuthMode string

// Enum values for IAMAuthMode
const (
	IAMAuthModeDisabled IAMAuthMode = "DISABLED"
	IAMAuthModeRequired IAMAuthMode = "REQUIRED"
	IAMAuthModeEnabled  IAMAuthMode = "ENABLED"
)

// Values returns all known values for IAMAuthMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (IAMAuthMode) Values() []IAMAuthMode {
	return []IAMAuthMode{
		"DISABLED",
		"REQUIRED",
		"ENABLED",
	}
}

type IntegrationStatus string

// Enum values for IntegrationStatus
const (
	IntegrationStatusCreating       IntegrationStatus = "creating"
	IntegrationStatusActive         IntegrationStatus = "active"
	IntegrationStatusModifying      IntegrationStatus = "modifying"
	IntegrationStatusFailed         IntegrationStatus = "failed"
	IntegrationStatusDeleting       IntegrationStatus = "deleting"
	IntegrationStatusSyncing        IntegrationStatus = "syncing"
	IntegrationStatusNeedsAttention IntegrationStatus = "needs_attention"
)

// Values returns all known values for IntegrationStatus. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (IntegrationStatus) Values() []IntegrationStatus {
	return []IntegrationStatus{
		"creating",
		"active",
		"modifying",
		"failed",
		"deleting",
		"syncing",
		"needs_attention",
	}
}

type LocalWriteForwardingStatus string

// Enum values for LocalWriteForwardingStatus
const (
	LocalWriteForwardingStatusEnabled   LocalWriteForwardingStatus = "enabled"
	LocalWriteForwardingStatusDisabled  LocalWriteForwardingStatus = "disabled"
	LocalWriteForwardingStatusEnabling  LocalWriteForwardingStatus = "enabling"
	LocalWriteForwardingStatusDisabling LocalWriteForwardingStatus = "disabling"
	LocalWriteForwardingStatusRequested LocalWriteForwardingStatus = "requested"
)

// Values returns all known values for LocalWriteForwardingStatus. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (LocalWriteForwardingStatus) Values() []LocalWriteForwardingStatus {
	return []LocalWriteForwardingStatus{
		"enabled",
		"disabled",
		"enabling",
		"disabling",
		"requested",
	}
}

type ReplicaMode string

// Enum values for ReplicaMode
const (
	ReplicaModeOpenReadOnly ReplicaMode = "open-read-only"
	ReplicaModeMounted      ReplicaMode = "mounted"
)

// Values returns all known values for ReplicaMode. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (ReplicaMode) Values() []ReplicaMode {
	return []ReplicaMode{
		"open-read-only",
		"mounted",
	}
}

type SourceType string

// Enum values for SourceType
const (
	SourceTypeDbInstance          SourceType = "db-instance"
	SourceTypeDbParameterGroup    SourceType = "db-parameter-group"
	SourceTypeDbSecurityGroup     SourceType = "db-security-group"
	SourceTypeDbSnapshot          SourceType = "db-snapshot"
	SourceTypeDbCluster           SourceType = "db-cluster"
	SourceTypeDbClusterSnapshot   SourceType = "db-cluster-snapshot"
	SourceTypeCustomEngineVersion SourceType = "custom-engine-version"
	SourceTypeDbProxy             SourceType = "db-proxy"
	SourceTypeBlueGreenDeployment SourceType = "blue-green-deployment"
)

// Values returns all known values for SourceType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (SourceType) Values() []SourceType {
	return []SourceType{
		"db-instance",
		"db-parameter-group",
		"db-security-group",
		"db-snapshot",
		"db-cluster",
		"db-cluster-snapshot",
		"custom-engine-version",
		"db-proxy",
		"blue-green-deployment",
	}
}

type TargetHealthReason string

// Enum values for TargetHealthReason
const (
	TargetHealthReasonUnreachable             TargetHealthReason = "UNREACHABLE"
	TargetHealthReasonConnectionFailed        TargetHealthReason = "CONNECTION_FAILED"
	TargetHealthReasonAuthFailure             TargetHealthReason = "AUTH_FAILURE"
	TargetHealthReasonPendingProxyCapacity    TargetHealthReason = "PENDING_PROXY_CAPACITY"
	TargetHealthReasonInvalidReplicationState TargetHealthReason = "INVALID_REPLICATION_STATE"
)

// Values returns all known values for TargetHealthReason. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (TargetHealthReason) Values() []TargetHealthReason {
	return []TargetHealthReason{
		"UNREACHABLE",
		"CONNECTION_FAILED",
		"AUTH_FAILURE",
		"PENDING_PROXY_CAPACITY",
		"INVALID_REPLICATION_STATE",
	}
}

type TargetRole string

// Enum values for TargetRole
const (
	TargetRoleReadWrite TargetRole = "READ_WRITE"
	TargetRoleReadOnly  TargetRole = "READ_ONLY"
	TargetRoleUnknown   TargetRole = "UNKNOWN"
)

// Values returns all known values for TargetRole. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TargetRole) Values() []TargetRole {
	return []TargetRole{
		"READ_WRITE",
		"READ_ONLY",
		"UNKNOWN",
	}
}

type TargetState string

// Enum values for TargetState
const (
	TargetStateRegistering TargetState = "REGISTERING"
	TargetStateAvailable   TargetState = "AVAILABLE"
	TargetStateUnavailable TargetState = "UNAVAILABLE"
)

// Values returns all known values for TargetState. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TargetState) Values() []TargetState {
	return []TargetState{
		"REGISTERING",
		"AVAILABLE",
		"UNAVAILABLE",
	}
}

type TargetType string

// Enum values for TargetType
const (
	TargetTypeRdsInstance           TargetType = "RDS_INSTANCE"
	TargetTypeRdsServerlessEndpoint TargetType = "RDS_SERVERLESS_ENDPOINT"
	TargetTypeTrackedCluster        TargetType = "TRACKED_CLUSTER"
)

// Values returns all known values for TargetType. Note that this can be expanded
// in the future, and so it is only as up to date as the client. The ordering of
// this slice is not guaranteed to be stable across updates.
func (TargetType) Values() []TargetType {
	return []TargetType{
		"RDS_INSTANCE",
		"RDS_SERVERLESS_ENDPOINT",
		"TRACKED_CLUSTER",
	}
}

type WriteForwardingStatus string

// Enum values for WriteForwardingStatus
const (
	WriteForwardingStatusEnabled   WriteForwardingStatus = "enabled"
	WriteForwardingStatusDisabled  WriteForwardingStatus = "disabled"
	WriteForwardingStatusEnabling  WriteForwardingStatus = "enabling"
	WriteForwardingStatusDisabling WriteForwardingStatus = "disabling"
	WriteForwardingStatusUnknown   WriteForwardingStatus = "unknown"
)

// Values returns all known values for WriteForwardingStatus. Note that this can
// be expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (WriteForwardingStatus) Values() []WriteForwardingStatus {
	return []WriteForwardingStatus{
		"enabled",
		"disabled",
		"enabling",
		"disabling",
		"unknown",
	}
}
