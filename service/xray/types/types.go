// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// An alias for an edge.
type Alias struct {

	// The canonical name of the alias.
	Name *string

	// A list of names for the alias, including the canonical name.
	Names []string

	// The type of the alias.
	Type *string
}

// Value of a segment annotation. Has one of three value types: Number, Boolean, or
// String.
type AnnotationValue struct {

	// Value for a Boolean annotation.
	BooleanValue *bool

	// Value for a Number annotation.
	NumberValue *float64

	// Value for a String annotation.
	StringValue *string
}

// A list of Availability Zones corresponding to the segments in a trace.
type AvailabilityZoneDetail struct {

	// The name of a corresponding Availability Zone.
	Name *string
}

//
type BackendConnectionErrors struct {

	//
	ConnectionRefusedCount *int32

	//
	HTTPCode4XXCount *int32

	//
	HTTPCode5XXCount *int32

	//
	OtherCount *int32

	//
	TimeoutCount *int32

	//
	UnknownHostCount *int32
}

// Information about a connection between two services.
type Edge struct {

	// Aliases for the edge.
	Aliases []Alias

	// The end time of the last segment on the edge.
	EndTime *time.Time

	// Identifier of the edge. Unique within a service map.
	ReferenceId *int32

	// A histogram that maps the spread of client response times on an edge.
	ResponseTimeHistogram []HistogramEntry

	// The start time of the first segment on the edge.
	StartTime *time.Time

	// Response statistics for segments on the edge.
	SummaryStatistics *EdgeStatistics
}

// Response statistics for an edge.
type EdgeStatistics struct {

	// Information about requests that failed with a 4xx Client Error status code.
	ErrorStatistics *ErrorStatistics

	// Information about requests that failed with a 5xx Server Error status code.
	FaultStatistics *FaultStatistics

	// The number of requests that completed with a 2xx Success status code.
	OkCount *int64

	// The total number of completed requests.
	TotalCount *int64

	// The aggregate response time of completed requests.
	TotalResponseTime *float64
}

// A configuration document that specifies encryption configuration settings.
type EncryptionConfig struct {

	// The ID of the customer master key (CMK) used for encryption, if applicable.
	KeyId *string

	// The encryption status. While the status is UPDATING, X-Ray may encrypt data with
	// a combination of the new and old settings.
	Status EncryptionStatus

	// The type of encryption. Set to KMS for encryption with CMKs. Set to NONE for
	// default encryption.
	Type EncryptionType
}

// The root cause of a trace summary error.
type ErrorRootCause struct {

	// A flag that denotes that the root cause impacts the trace client.
	ClientImpacting *bool

	// A list of services corresponding to an error. A service identifies a segment and
	// it contains a name, account ID, type, and inferred flag.
	Services []ErrorRootCauseService
}

// A collection of segments and corresponding subsegments associated to a trace
// summary error.
type ErrorRootCauseEntity struct {

	// The types and messages of the exceptions.
	Exceptions []RootCauseException

	// The name of the entity.
	Name *string

	// A flag that denotes a remote subsegment.
	Remote *bool
}

// A collection of fields identifying the services in a trace summary error.
type ErrorRootCauseService struct {

	// The account ID associated to the service.
	AccountId *string

	// The path of root cause entities found on the service.
	EntityPath []ErrorRootCauseEntity

	// A Boolean value indicating if the service is inferred from the trace.
	Inferred *bool

	// The service name.
	Name *string

	// A collection of associated service names.
	Names []string

	// The type associated to the service.
	Type *string
}

// Information about requests that failed with a 4xx Client Error status code.
type ErrorStatistics struct {

	// The number of requests that failed with untracked 4xx Client Error status codes.
	OtherCount *int64

	// The number of requests that failed with a 419 throttling status code.
	ThrottleCount *int64

	// The total number of requests that failed with a 4xx Client Error status code.
	TotalCount *int64
}

// The root cause information for a trace summary fault.
type FaultRootCause struct {

	// A flag that denotes that the root cause impacts the trace client.
	ClientImpacting *bool

	// A list of corresponding services. A service identifies a segment and it contains
	// a name, account ID, type, and inferred flag.
	Services []FaultRootCauseService
}

// A collection of segments and corresponding subsegments associated to a trace
// summary fault error.
type FaultRootCauseEntity struct {

	// The types and messages of the exceptions.
	Exceptions []RootCauseException

	// The name of the entity.
	Name *string

	// A flag that denotes a remote subsegment.
	Remote *bool
}

// A collection of fields identifying the services in a trace summary fault.
type FaultRootCauseService struct {

	// The account ID associated to the service.
	AccountId *string

	// The path of root cause entities found on the service.
	EntityPath []FaultRootCauseEntity

	// A Boolean value indicating if the service is inferred from the trace.
	Inferred *bool

	// The service name.
	Name *string

	// A collection of associated service names.
	Names []string

	// The type associated to the service.
	Type *string
}

// Information about requests that failed with a 5xx Server Error status code.
type FaultStatistics struct {

	// The number of requests that failed with untracked 5xx Server Error status codes.
	OtherCount *int64

	// The total number of requests that failed with a 5xx Server Error status code.
	TotalCount *int64
}

// Details and metadata for a group.
type Group struct {

	// The filter expression defining the parameters to include traces.
	FilterExpression *string

	// The Amazon Resource Name (ARN) of the group generated based on the GroupName.
	GroupARN *string

	// The unique case-sensitive name of the group.
	GroupName *string

	// The structure containing configurations related to insights.
	//
	// * The
	// InsightsEnabled boolean can be set to true to enable insights for the group or
	// false to disable insights for the group.
	//
	// * The NotifcationsEnabled boolean can
	// be set to true to enable insights notifications through Amazon EventBridge for
	// the group.
	InsightsConfiguration *InsightsConfiguration
}

// Details for a group without metadata.
type GroupSummary struct {

	// The filter expression defining the parameters to include traces.
	FilterExpression *string

	// The ARN of the group generated based on the GroupName.
	GroupARN *string

	// The unique case-sensitive name of the group.
	GroupName *string

	// The structure containing configurations related to insights.
	//
	// * The
	// InsightsEnabled boolean can be set to true to enable insights for the group or
	// false to disable insights for the group.
	//
	// * The NotificationsEnabled boolean can
	// be set to true to enable insights notifications. Notifications can only be
	// enabled on a group with InsightsEnabled set to true.
	InsightsConfiguration *InsightsConfiguration
}

// An entry in a histogram for a statistic. A histogram maps the range of observed
// values on the X axis, and the prevalence of each value on the Y axis.
type HistogramEntry struct {

	// The prevalence of the entry.
	Count int32

	// The value of the entry.
	Value float64
}

// Information about an HTTP request.
type Http struct {

	// The IP address of the requestor.
	ClientIp *string

	// The request method.
	HttpMethod *string

	// The response status.
	HttpStatus *int32

	// The request URL.
	HttpURL *string

	// The request's user agent string.
	UserAgent *string
}

// The structure containing configurations related to insights.
type InsightsConfiguration struct {

	// Set the InsightsEnabled value to true to enable insights or false to disable
	// insights.
	InsightsEnabled *bool

	// Set the NotificationsEnabled value to true to enable insights notifications.
	// Notifications can only be enabled on a group with InsightsEnabled set to true.
	NotificationsEnabled *bool
}

// A list of EC2 instance IDs corresponding to the segments in a trace.
type InstanceIdDetail struct {

	// The ID of a corresponding EC2 instance.
	Id *string
}

// A list of resources ARNs corresponding to the segments in a trace.
type ResourceARNDetail struct {

	// The ARN of a corresponding resource.
	ARN *string
}

// The root cause information for a response time warning.
type ResponseTimeRootCause struct {

	// A flag that denotes that the root cause impacts the trace client.
	ClientImpacting *bool

	// A list of corresponding services. A service identifies a segment and contains a
	// name, account ID, type, and inferred flag.
	Services []ResponseTimeRootCauseService
}

// A collection of segments and corresponding subsegments associated to a response
// time warning.
type ResponseTimeRootCauseEntity struct {

	// The type and messages of the exceptions.
	Coverage *float64

	// The name of the entity.
	Name *string

	// A flag that denotes a remote subsegment.
	Remote *bool
}

// A collection of fields identifying the service in a response time warning.
type ResponseTimeRootCauseService struct {

	// The account ID associated to the service.
	AccountId *string

	// The path of root cause entities found on the service.
	EntityPath []ResponseTimeRootCauseEntity

	// A Boolean value indicating if the service is inferred from the trace.
	Inferred *bool

	// The service name.
	Name *string

	// A collection of associated service names.
	Names []string

	// The type associated to the service.
	Type *string
}

// The exception associated with a root cause.
type RootCauseException struct {

	// The message of the exception.
	Message *string

	// The name of the exception.
	Name *string
}

// A sampling rule that services use to decide whether to instrument a request.
// Rule fields can match properties of the service, or properties of a request. The
// service can ignore rules that don't match its properties.
type SamplingRule struct {

	// The percentage of matching requests to instrument, after the reservoir is
	// exhausted.
	//
	// This member is required.
	FixedRate float64

	// Matches the HTTP method of a request.
	//
	// This member is required.
	HTTPMethod *string

	// Matches the hostname from a request URL.
	//
	// This member is required.
	Host *string

	// The priority of the sampling rule.
	//
	// This member is required.
	Priority int32

	// A fixed number of matching requests to instrument per second, prior to applying
	// the fixed rate. The reservoir is not used directly by services, but applies to
	// all services using the rule collectively.
	//
	// This member is required.
	ReservoirSize int32

	// Matches the ARN of the AWS resource on which the service runs.
	//
	// This member is required.
	ResourceARN *string

	// Matches the name that the service uses to identify itself in segments.
	//
	// This member is required.
	ServiceName *string

	// Matches the origin that the service uses to identify its type in segments.
	//
	// This member is required.
	ServiceType *string

	// Matches the path from a request URL.
	//
	// This member is required.
	URLPath *string

	// The version of the sampling rule format (1).
	//
	// This member is required.
	Version int32

	// Matches attributes derived from the request.
	Attributes map[string]string

	// The ARN of the sampling rule. Specify a rule by either name or ARN, but not
	// both.
	RuleARN *string

	// The name of the sampling rule. Specify a rule by either name or ARN, but not
	// both.
	RuleName *string
}

// A SamplingRule and its metadata.
type SamplingRuleRecord struct {

	// When the rule was created.
	CreatedAt *time.Time

	// When the rule was last modified.
	ModifiedAt *time.Time

	// The sampling rule.
	SamplingRule *SamplingRule
}

// A document specifying changes to a sampling rule's configuration.
type SamplingRuleUpdate struct {

	// Matches attributes derived from the request.
	Attributes map[string]string

	// The percentage of matching requests to instrument, after the reservoir is
	// exhausted.
	FixedRate *float64

	// Matches the HTTP method of a request.
	HTTPMethod *string

	// Matches the hostname from a request URL.
	Host *string

	// The priority of the sampling rule.
	Priority *int32

	// A fixed number of matching requests to instrument per second, prior to applying
	// the fixed rate. The reservoir is not used directly by services, but applies to
	// all services using the rule collectively.
	ReservoirSize *int32

	// Matches the ARN of the AWS resource on which the service runs.
	ResourceARN *string

	// The ARN of the sampling rule. Specify a rule by either name or ARN, but not
	// both.
	RuleARN *string

	// The name of the sampling rule. Specify a rule by either name or ARN, but not
	// both.
	RuleName *string

	// Matches the name that the service uses to identify itself in segments.
	ServiceName *string

	// Matches the origin that the service uses to identify its type in segments.
	ServiceType *string

	// Matches the path from a request URL.
	URLPath *string
}

// Request sampling results for a single rule from a service. Results are for the
// last 10 seconds unless the service has been assigned a longer reporting interval
// after a previous call to GetSamplingTargets.
type SamplingStatisticsDocument struct {

	// A unique identifier for the service in hexadecimal.
	//
	// This member is required.
	ClientID *string

	// The number of requests that matched the rule.
	//
	// This member is required.
	RequestCount int32

	// The name of the sampling rule.
	//
	// This member is required.
	RuleName *string

	// The number of requests recorded.
	//
	// This member is required.
	SampledCount int32

	// The current time.
	//
	// This member is required.
	Timestamp *time.Time

	// The number of requests recorded with borrowed reservoir quota.
	BorrowCount int32
}

// Aggregated request sampling data for a sampling rule across all services for a
// 10-second window.
type SamplingStatisticSummary struct {

	// The number of requests recorded with borrowed reservoir quota.
	BorrowCount int32

	// The number of requests that matched the rule.
	RequestCount int32

	// The name of the sampling rule.
	RuleName *string

	// The number of requests recorded.
	SampledCount int32

	// The start time of the reporting window.
	Timestamp *time.Time
}

// The name and value of a sampling rule to apply to a trace summary.
type SamplingStrategy struct {

	// The name of a sampling rule.
	Name SamplingStrategyName

	// The value of a sampling rule.
	Value *float64
}

// Temporary changes to a sampling rule configuration. To meet the global sampling
// target for a rule, X-Ray calculates a new reservoir for each service based on
// the recent sampling results of all services that called GetSamplingTargets.
type SamplingTargetDocument struct {

	// The percentage of matching requests to instrument, after the reservoir is
	// exhausted.
	FixedRate float64

	// The number of seconds for the service to wait before getting sampling targets
	// again.
	Interval *int32

	// The number of requests per second that X-Ray allocated for this service.
	ReservoirQuota *int32

	// When the reservoir quota expires.
	ReservoirQuotaTTL *time.Time

	// The name of the sampling rule.
	RuleName *string
}

// A segment from a trace that has been ingested by the X-Ray service. The segment
// can be compiled from documents uploaded with PutTraceSegments, or an inferred
// segment for a downstream service, generated from a subsegment sent by the
// service that called it. For the full segment document schema, see AWS X-Ray
// Segment Documents
// (https://docs.aws.amazon.com/xray/latest/devguide/xray-api-segmentdocuments.html)
// in the AWS X-Ray Developer Guide.
type Segment struct {

	// The segment document.
	Document *string

	// The segment's ID.
	Id *string
}

// Information about an application that processed requests, users that made
// requests, or downstream services, resources, and applications that an
// application used.
type Service struct {

	// Identifier of the AWS account in which the service runs.
	AccountId *string

	// A histogram that maps the spread of service durations.
	DurationHistogram []HistogramEntry

	// Connections to downstream services.
	Edges []Edge

	// The end time of the last segment that the service generated.
	EndTime *time.Time

	// The canonical name of the service.
	Name *string

	// A list of names for the service, including the canonical name.
	Names []string

	// Identifier for the service. Unique within the service map.
	ReferenceId *int32

	// A histogram that maps the spread of service response times.
	ResponseTimeHistogram []HistogramEntry

	// Indicates that the service was the first service to process a request.
	Root *bool

	// The start time of the first segment that the service generated.
	StartTime *time.Time

	// The service's state.
	State *string

	// Aggregated statistics for the service.
	SummaryStatistics *ServiceStatistics

	// The type of service.
	//
	// * AWS Resource - The type of an AWS resource. For example,
	// AWS::EC2::Instance for an application running on Amazon EC2 or
	// AWS::DynamoDB::Table for an Amazon DynamoDB table that the application used.
	//
	// *
	// AWS Service - The type of an AWS service. For example, AWS::DynamoDB for
	// downstream calls to Amazon DynamoDB that didn't target a specific table.
	//
	// *
	// client - Represents the clients that sent requests to a root service.
	//
	// * remote
	// - A downstream service of indeterminate type.
	Type *string
}

//
type ServiceId struct {

	//
	AccountId *string

	//
	Name *string

	//
	Names []string

	//
	Type *string
}

// Response statistics for a service.
type ServiceStatistics struct {

	// Information about requests that failed with a 4xx Client Error status code.
	ErrorStatistics *ErrorStatistics

	// Information about requests that failed with a 5xx Server Error status code.
	FaultStatistics *FaultStatistics

	// The number of requests that completed with a 2xx Success status code.
	OkCount *int64

	// The total number of completed requests.
	TotalCount *int64

	// The aggregate response time of completed requests.
	TotalResponseTime *float64
}

// A map that contains tag keys and tag values to attach to an AWS X-Ray group or
// sampling rule. For more information about ways to use tags, see Tagging AWS
// resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html) in
// the AWS General Reference. The following restrictions apply to tags:
//
// * Maximum
// number of user-applied tags per resource: 50
//
// * Tag keys and values are case
// sensitive.
//
// * Don't use aws: as a prefix for keys; it's reserved for AWS use.
// You cannot edit or delete system tags.
type Tag struct {

	// A tag key, such as Stage or Name. A tag key cannot be empty. The key can be a
	// maximum of 128 characters, and can contain only Unicode letters, numbers, or
	// separators, or the following special characters: + - = . _ : /
	//
	// This member is required.
	Key *string

	// An optional tag value, such as Production or test-only. The value can be a
	// maximum of 255 characters, and contain only Unicode letters, numbers, or
	// separators, or the following special characters: + - = . _ : /
	//
	// This member is required.
	Value *string
}

//
type TelemetryRecord struct {

	//
	//
	// This member is required.
	Timestamp *time.Time

	//
	BackendConnectionErrors *BackendConnectionErrors

	//
	SegmentsReceivedCount *int32

	//
	SegmentsRejectedCount *int32

	//
	SegmentsSentCount *int32

	//
	SegmentsSpilloverCount *int32
}

// A list of TimeSeriesStatistic structures.
type TimeSeriesServiceStatistics struct {

	// Response statistics for an edge.
	EdgeSummaryStatistics *EdgeStatistics

	// The response time histogram for the selected entities.
	ResponseTimeHistogram []HistogramEntry

	// Response statistics for a service.
	ServiceSummaryStatistics *ServiceStatistics

	// Timestamp of the window for which statistics are aggregated.
	Timestamp *time.Time
}

// A collection of segment documents with matching trace IDs.
type Trace struct {

	// The length of time in seconds between the start time of the root segment and the
	// end time of the last segment that completed.
	Duration *float64

	// The unique identifier for the request that generated the trace's segments and
	// subsegments.
	Id *string

	// LimitExceeded is set to true when the trace has exceeded one of the defined
	// quotas. For more information about quotas, see AWS X-Ray endpoints and quotas
	// (https://docs.aws.amazon.com/general/latest/gr/xray.html).
	LimitExceeded *bool

	// Segment documents for the segments and subsegments that comprise the trace.
	Segments []Segment
}

// Metadata generated from the segment documents in a trace.
type TraceSummary struct {

	// Annotations from the trace's segment documents.
	Annotations map[string][]ValueWithServiceIds

	// A list of Availability Zones for any zone corresponding to the trace segments.
	AvailabilityZones []AvailabilityZoneDetail

	// The length of time in seconds between the start time of the root segment and the
	// end time of the last segment that completed.
	Duration *float64

	// The root of a trace.
	EntryPoint *ServiceId

	// A collection of ErrorRootCause structures corresponding to the trace segments.
	ErrorRootCauses []ErrorRootCause

	// A collection of FaultRootCause structures corresponding to the trace segments.
	FaultRootCauses []FaultRootCause

	// The root segment document has a 400 series error.
	HasError *bool

	// The root segment document has a 500 series error.
	HasFault *bool

	// One or more of the segment documents has a 429 throttling error.
	HasThrottle *bool

	// Information about the HTTP request served by the trace.
	Http *Http

	// The unique identifier for the request that generated the trace's segments and
	// subsegments.
	Id *string

	// A list of EC2 instance IDs for any instance corresponding to the trace segments.
	InstanceIds []InstanceIdDetail

	// One or more of the segment documents is in progress.
	IsPartial *bool

	// The matched time stamp of a defined event.
	MatchedEventTime *time.Time

	// A list of resource ARNs for any resource corresponding to the trace segments.
	ResourceARNs []ResourceARNDetail

	// The length of time in seconds between the start and end times of the root
	// segment. If the service performs work asynchronously, the response time measures
	// the time before the response is sent to the user, while the duration measures
	// the amount of time before the last traced activity completes.
	ResponseTime *float64

	// A collection of ResponseTimeRootCause structures corresponding to the trace
	// segments.
	ResponseTimeRootCauses []ResponseTimeRootCause

	// The revision number of a trace.
	Revision int32

	// Service IDs from the trace's segment documents.
	ServiceIds []ServiceId

	// Users from the trace's segment documents.
	Users []TraceUser
}

// Information about a user recorded in segment documents.
type TraceUser struct {

	// Services that the user's request hit.
	ServiceIds []ServiceId

	// The user's name.
	UserName *string
}

// Sampling statistics from a call to GetSamplingTargets that X-Ray could not
// process.
type UnprocessedStatistics struct {

	// The error code.
	ErrorCode *string

	// The error message.
	Message *string

	// The name of the sampling rule.
	RuleName *string
}

// Information about a segment that failed processing.
type UnprocessedTraceSegment struct {

	// The error that caused processing to fail.
	ErrorCode *string

	// The segment's ID.
	Id *string

	// The error message.
	Message *string
}

// Information about a segment annotation.
type ValueWithServiceIds struct {

	// Values of the annotation.
	AnnotationValue *AnnotationValue

	// Services to which the annotation applies.
	ServiceIds []ServiceId
}
