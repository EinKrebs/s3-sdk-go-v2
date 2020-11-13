// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"time"
)

// The details of a DDoS attack.
type AttackDetail struct {

	// List of counters that describe the attack for the specified time period.
	AttackCounters []SummarizedCounter

	// The unique identifier (ID) of the attack.
	AttackId *string

	// The array of AttackProperty objects.
	AttackProperties []AttackProperty

	// The time the attack ended, in Unix time in seconds. For more information see
	// timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	EndTime *time.Time

	// List of mitigation actions taken for the attack.
	Mitigations []Mitigation

	// The ARN (Amazon Resource Name) of the resource that was attacked.
	ResourceArn *string

	// The time the attack started, in Unix time in seconds. For more information see
	// timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	StartTime *time.Time

	// If applicable, additional detail about the resource being attacked, for example,
	// IP address or URL.
	SubResources []SubResourceSummary
}

// Details of the described attack.
type AttackProperty struct {

	// The type of distributed denial of service (DDoS) event that was observed.
	// NETWORK indicates layer 3 and layer 4 events and APPLICATION indicates layer 7
	// events.
	AttackLayer AttackLayer

	// Defines the DDoS attack property information that is provided. The
	// WORDPRESS_PINGBACK_REFLECTOR and WORDPRESS_PINGBACK_SOURCE values are valid only
	// for WordPress reflective pingback DDoS attacks.
	AttackPropertyIdentifier AttackPropertyIdentifier

	// The array of Contributor objects that includes the top five contributors to an
	// attack.
	TopContributors []Contributor

	// The total contributions made to this attack by all contributors, not just the
	// five listed in the TopContributors list.
	Total int64

	// The unit of the Value of the contributions.
	Unit Unit
}

// Summarizes all DDoS attacks for a specified time period.
type AttackSummary struct {

	// The unique identifier (ID) of the attack.
	AttackId *string

	// The list of attacks for a specified time period.
	AttackVectors []AttackVectorDescription

	// The end time of the attack, in Unix time in seconds. For more information see
	// timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	EndTime *time.Time

	// The ARN (Amazon Resource Name) of the resource that was attacked.
	ResourceArn *string

	// The start time of the attack, in Unix time in seconds. For more information see
	// timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	StartTime *time.Time
}

// Describes the attack.
type AttackVectorDescription struct {

	// The attack type. Valid values:
	//
	// * UDP_TRAFFIC
	//
	// * UDP_FRAGMENT
	//
	// *
	// GENERIC_UDP_REFLECTION
	//
	// * DNS_REFLECTION
	//
	// * NTP_REFLECTION
	//
	// *
	// CHARGEN_REFLECTION
	//
	// * SSDP_REFLECTION
	//
	// * PORT_MAPPER
	//
	// * RIP_REFLECTION
	//
	// *
	// SNMP_REFLECTION
	//
	// * MSSQL_REFLECTION
	//
	// * NET_BIOS_REFLECTION
	//
	// * SYN_FLOOD
	//
	// *
	// ACK_FLOOD
	//
	// * REQUEST_FLOOD
	//
	// * HTTP_REFLECTION
	//
	// * UDS_REFLECTION
	//
	// *
	// MEMCACHED_REFLECTION
	//
	// This member is required.
	VectorType *string
}

// A contributor to the attack and their contribution.
type Contributor struct {

	// The name of the contributor. This is dependent on the AttackPropertyIdentifier.
	// For example, if the AttackPropertyIdentifier is SOURCE_COUNTRY, the Name could
	// be United States.
	Name *string

	// The contribution of this contributor expressed in Protection units. For example
	// 10,000.
	Value int64
}

// Contact information that the DRT can use to contact you if you have proactive
// engagement enabled, for escalations to the DRT and to initiate proactive
// customer support.
type EmergencyContact struct {

	// The email address for the contact.
	//
	// This member is required.
	EmailAddress *string

	// Additional notes regarding the contact.
	ContactNotes *string

	// The phone number for the contact.
	PhoneNumber *string
}

// Specifies how many protections of a given type you can create.
type Limit struct {

	// The maximum number of protections that can be created for the specified Type.
	Max int64

	// The type of protection.
	Type *string
}

// The mitigation applied to a DDoS attack.
type Mitigation struct {

	// The name of the mitigation taken for this attack.
	MitigationName *string
}

// An object that represents a resource that is under DDoS protection.
type Protection struct {

	// The unique identifier (ID) for the Route 53 health check that's associated with
	// the protection.
	HealthCheckIds []string

	// The unique identifier (ID) of the protection.
	Id *string

	// The friendly name of the protection. For example, My CloudFront distributions.
	Name *string

	// The ARN (Amazon Resource Name) of the AWS resource that is protected.
	ResourceArn *string
}

// The attack information for the specified SubResource.
type SubResourceSummary struct {

	// The list of attack types and associated counters.
	AttackVectors []SummarizedAttackVector

	// The counters that describe the details of the attack.
	Counters []SummarizedCounter

	// The unique identifier (ID) of the SubResource.
	Id *string

	// The SubResource type.
	Type SubResourceType
}

// Information about the AWS Shield Advanced subscription for an account.
type Subscription struct {

	// If ENABLED, the subscription will be automatically renewed at the end of the
	// existing subscription period. When you initally create a subscription, AutoRenew
	// is set to ENABLED. You can change this by submitting an UpdateSubscription
	// request. If the UpdateSubscription request does not included a value for
	// AutoRenew, the existing value for AutoRenew remains unchanged.
	AutoRenew AutoRenew

	// The date and time your subscription will end.
	EndTime *time.Time

	// Specifies how many protections of a given type you can create.
	Limits []Limit

	// If ENABLED, the DDoS Response Team (DRT) will use email and phone to notify
	// contacts about escalations to the DRT and to initiate proactive customer
	// support. If PENDING, you have requested proactive engagement and the request is
	// pending. The status changes to ENABLED when your request is fully processed. If
	// DISABLED, the DRT will not proactively notify contacts about escalations or to
	// initiate proactive customer support.
	ProactiveEngagementStatus ProactiveEngagementStatus

	// The start time of the subscription, in Unix time in seconds. For more
	// information see timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	StartTime *time.Time

	// The length, in seconds, of the AWS Shield Advanced subscription for the account.
	TimeCommitmentInSeconds int64
}

// A summary of information about the attack.
type SummarizedAttackVector struct {

	// The attack type, for example, SNMP reflection or SYN flood.
	//
	// This member is required.
	VectorType *string

	// The list of counters that describe the details of the attack.
	VectorCounters []SummarizedCounter
}

// The counter that describes a DDoS attack.
type SummarizedCounter struct {

	// The average value of the counter for a specified time period.
	Average float64

	// The maximum value of the counter for a specified time period.
	Max float64

	// The number of counters for a specified time period.
	N int32

	// The counter name.
	Name *string

	// The total of counter values for a specified time period.
	Sum float64

	// The unit of the counters.
	Unit *string
}

// The time range.
type TimeRange struct {

	// The start time, in Unix time in seconds. For more information see timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	FromInclusive *time.Time

	// The end time, in Unix time in seconds. For more information see timestamp
	// (http://docs.aws.amazon.com/cli/latest/userguide/cli-using-param.html#parameter-types).
	ToExclusive *time.Time
}
