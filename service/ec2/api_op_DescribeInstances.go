// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Describes the specified instances or all instances. If you specify instance IDs,
// the output includes information for only the specified instances. If you specify
// filters, the output includes information for only those instances that meet the
// filter criteria. If you do not specify instance IDs or filters, the output
// includes information for all instances, which can affect performance. We
// recommend that you use pagination to ensure that the operation returns quickly
// and successfully. If you specify an instance ID that is not valid, an error is
// returned. If you specify an instance that you do not own, it is not included in
// the output. Recently terminated instances might appear in the returned results.
// This interval is usually less than one hour. If you describe instances in the
// rare case where an Availability Zone is experiencing a service disruption and
// you specify instance IDs that are in the affected zone, or do not specify any
// instance IDs at all, the call fails. If you describe instances and specify only
// instance IDs that are in an unaffected zone, the call works normally.
func (c *Client) DescribeInstances(ctx context.Context, params *DescribeInstancesInput, optFns ...func(*Options)) (*DescribeInstancesOutput, error) {
	if params == nil {
		params = &DescribeInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeInstances", params, optFns, addOperationDescribeInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeInstancesInput struct {

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun bool

	// The filters.
	//
	// * affinity - The affinity setting for an instance running on a
	// Dedicated Host (default | host).
	//
	// * architecture - The instance architecture
	// (i386 | x86_64 | arm64).
	//
	// * availability-zone - The Availability Zone of the
	// instance.
	//
	// * block-device-mapping.attach-time - The attach time for an EBS
	// volume mapped to the instance, for example, 2010-09-15T17:15:20.000Z.
	//
	// *
	// block-device-mapping.delete-on-termination - A Boolean that indicates whether
	// the EBS volume is deleted on instance termination.
	//
	// *
	// block-device-mapping.device-name - The device name specified in the block device
	// mapping (for example, /dev/sdh or xvdh).
	//
	// * block-device-mapping.status - The
	// status for the EBS volume (attaching | attached | detaching | detached).
	//
	// *
	// block-device-mapping.volume-id - The volume ID of the EBS volume.
	//
	// *
	// client-token - The idempotency token you provided when you launched the
	// instance.
	//
	// * dns-name - The public DNS name of the instance.
	//
	// * group-id - The
	// ID of the security group for the instance. EC2-Classic only.
	//
	// * group-name - The
	// name of the security group for the instance. EC2-Classic only.
	//
	// *
	// hibernation-options.configured - A Boolean that indicates whether the instance
	// is enabled for hibernation. A value of true means that the instance is enabled
	// for hibernation.
	//
	// * host-id - The ID of the Dedicated Host on which the instance
	// is running, if applicable.
	//
	// * hypervisor - The hypervisor type of the instance
	// (ovm | xen). The value xen is used for both Xen and Nitro hypervisors.
	//
	// *
	// iam-instance-profile.arn - The instance profile associated with the instance.
	// Specified as an ARN.
	//
	// * image-id - The ID of the image used to launch the
	// instance.
	//
	// * instance-id - The ID of the instance.
	//
	// * instance-lifecycle -
	// Indicates whether this is a Spot Instance or a Scheduled Instance (spot |
	// scheduled).
	//
	// * instance-state-code - The state of the instance, as a 16-bit
	// unsigned integer. The high byte is used for internal purposes and should be
	// ignored. The low byte is set based on the state represented. The valid values
	// are: 0 (pending), 16 (running), 32 (shutting-down), 48 (terminated), 64
	// (stopping), and 80 (stopped).
	//
	// * instance-state-name - The state of the instance
	// (pending | running | shutting-down | terminated | stopping | stopped).
	//
	// *
	// instance-type - The type of instance (for example, t2.micro).
	//
	// *
	// instance.group-id - The ID of the security group for the instance.
	//
	// *
	// instance.group-name - The name of the security group for the instance.
	//
	// *
	// ip-address - The public IPv4 address of the instance.
	//
	// * kernel-id - The kernel
	// ID.
	//
	// * key-name - The name of the key pair used when the instance was
	// launched.
	//
	// * launch-index - When launching multiple instances, this is the index
	// for the instance in the launch group (for example, 0, 1, 2, and so on).
	//
	// *
	// launch-time - The time when the instance was launched.
	//
	// *
	// metadata-options.http-tokens - The metadata request authorization state
	// (optional | required)
	//
	// * metadata-options.http-put-response-hop-limit - The http
	// metadata request put response hop limit (integer, possible values 1 to 64)
	//
	// *
	// metadata-options.http-endpoint - Enable or disable metadata access on http
	// endpoint (enabled | disabled)
	//
	// * monitoring-state - Indicates whether detailed
	// monitoring is enabled (disabled | enabled).
	//
	// *
	// network-interface.addresses.private-ip-address - The private IPv4 address
	// associated with the network interface.
	//
	// * network-interface.addresses.primary -
	// Specifies whether the IPv4 address of the network interface is the primary
	// private IPv4 address.
	//
	// * network-interface.addresses.association.public-ip - The
	// ID of the association of an Elastic IP address (IPv4) with a network
	// interface.
	//
	// * network-interface.addresses.association.ip-owner-id - The owner ID
	// of the private IPv4 address associated with the network interface.
	//
	// *
	// network-interface.association.public-ip - The address of the Elastic IP address
	// (IPv4) bound to the network interface.
	//
	// *
	// network-interface.association.ip-owner-id - The owner of the Elastic IP address
	// (IPv4) associated with the network interface.
	//
	// *
	// network-interface.association.allocation-id - The allocation ID returned when
	// you allocated the Elastic IP address (IPv4) for your network interface.
	//
	// *
	// network-interface.association.association-id - The association ID returned when
	// the network interface was associated with an IPv4 address.
	//
	// *
	// network-interface.attachment.attachment-id - The ID of the interface
	// attachment.
	//
	// * network-interface.attachment.instance-id - The ID of the instance
	// to which the network interface is attached.
	//
	// *
	// network-interface.attachment.instance-owner-id - The owner ID of the instance to
	// which the network interface is attached.
	//
	// *
	// network-interface.attachment.device-index - The device index to which the
	// network interface is attached.
	//
	// * network-interface.attachment.status - The
	// status of the attachment (attaching | attached | detaching | detached).
	//
	// *
	// network-interface.attachment.attach-time - The time that the network interface
	// was attached to an instance.
	//
	// *
	// network-interface.attachment.delete-on-termination - Specifies whether the
	// attachment is deleted when an instance is terminated.
	//
	// *
	// network-interface.availability-zone - The Availability Zone for the network
	// interface.
	//
	// * network-interface.description - The description of the network
	// interface.
	//
	// * network-interface.group-id - The ID of a security group associated
	// with the network interface.
	//
	// * network-interface.group-name - The name of a
	// security group associated with the network interface.
	//
	// *
	// network-interface.ipv6-addresses.ipv6-address - The IPv6 address associated with
	// the network interface.
	//
	// * network-interface.mac-address - The MAC address of the
	// network interface.
	//
	// * network-interface.network-interface-id - The ID of the
	// network interface.
	//
	// * network-interface.owner-id - The ID of the owner of the
	// network interface.
	//
	// * network-interface.private-dns-name - The private DNS name
	// of the network interface.
	//
	// * network-interface.requester-id - The requester ID
	// for the network interface.
	//
	// * network-interface.requester-managed - Indicates
	// whether the network interface is being managed by AWS.
	//
	// *
	// network-interface.status - The status of the network interface (available) |
	// in-use).
	//
	// * network-interface.source-dest-check - Whether the network interface
	// performs source/destination checking. A value of true means that checking is
	// enabled, and false means that checking is disabled. The value must be false for
	// the network interface to perform network address translation (NAT) in your
	// VPC.
	//
	// * network-interface.subnet-id - The ID of the subnet for the network
	// interface.
	//
	// * network-interface.vpc-id - The ID of the VPC for the network
	// interface.
	//
	// * owner-id - The AWS account ID of the instance owner.
	//
	// *
	// placement-group-name - The name of the placement group for the instance.
	//
	// *
	// placement-partition-number - The partition in which the instance is located.
	//
	// *
	// platform - The platform. To list only Windows instances, use windows.
	//
	// *
	// private-dns-name - The private IPv4 DNS name of the instance.
	//
	// *
	// private-ip-address - The private IPv4 address of the instance.
	//
	// * product-code -
	// The product code associated with the AMI used to launch the instance.
	//
	// *
	// product-code.type - The type of product code (devpay | marketplace).
	//
	// *
	// ramdisk-id - The RAM disk ID.
	//
	// * reason - The reason for the current state of
	// the instance (for example, shows "User Initiated [date]" when you stop or
	// terminate the instance). Similar to the state-reason-code filter.
	//
	// *
	// requester-id - The ID of the entity that launched the instance on your behalf
	// (for example, AWS Management Console, Auto Scaling, and so on).
	//
	// *
	// reservation-id - The ID of the instance's reservation. A reservation ID is
	// created any time you launch an instance. A reservation ID has a one-to-one
	// relationship with an instance launch request, but can be associated with more
	// than one instance if you launch multiple instances using the same launch
	// request. For example, if you launch one instance, you get one reservation ID. If
	// you launch ten instances using the same launch request, you also get one
	// reservation ID.
	//
	// * root-device-name - The device name of the root device volume
	// (for example, /dev/sda1).
	//
	// * root-device-type - The type of the root device
	// volume (ebs | instance-store).
	//
	// * source-dest-check - Indicates whether the
	// instance performs source/destination checking. A value of true means that
	// checking is enabled, and false means that checking is disabled. The value must
	// be false for the instance to perform network address translation (NAT) in your
	// VPC.
	//
	// * spot-instance-request-id - The ID of the Spot Instance request.
	//
	// *
	// state-reason-code - The reason code for the state change.
	//
	// *
	// state-reason-message - A message that describes the state change.
	//
	// * subnet-id -
	// The ID of the subnet for the instance.
	//
	// * tag: - The key/value combination of a
	// tag assigned to the resource. Use the tag key in the filter name and the tag
	// value as the filter value. For example, to find all resources that have a tag
	// with the key Owner and the value TeamA, specify tag:Owner for the filter name
	// and TeamA for the filter value.
	//
	// * tag-key - The key of a tag assigned to the
	// resource. Use this filter to find all resources that have a tag with a specific
	// key, regardless of the tag value.
	//
	// * tenancy - The tenancy of an instance
	// (dedicated | default | host).
	//
	// * virtualization-type - The virtualization type
	// of the instance (paravirtual | hvm).
	//
	// * vpc-id - The ID of the VPC that the
	// instance is running in.
	Filters []types.Filter

	// The instance IDs. Default: Describes all your instances.
	InstanceIds []string

	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 5 and 1000. You cannot specify this parameter and the
	// instance IDs parameter in the same call.
	MaxResults int32

	// The token to request the next page of results.
	NextToken *string
}

type DescribeInstancesOutput struct {

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the reservations.
	Reservations []types.Reservation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeInstances{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeInstances(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// DescribeInstancesAPIClient is a client that implements the DescribeInstances
// operation.
type DescribeInstancesAPIClient interface {
	DescribeInstances(context.Context, *DescribeInstancesInput, ...func(*Options)) (*DescribeInstancesOutput, error)
}

var _ DescribeInstancesAPIClient = (*Client)(nil)

// DescribeInstancesPaginatorOptions is the paginator options for DescribeInstances
type DescribeInstancesPaginatorOptions struct {
	// The maximum number of results to return in a single call. To retrieve the
	// remaining results, make another call with the returned NextToken value. This
	// value can be between 5 and 1000. You cannot specify this parameter and the
	// instance IDs parameter in the same call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeInstancesPaginator is a paginator for DescribeInstances
type DescribeInstancesPaginator struct {
	options   DescribeInstancesPaginatorOptions
	client    DescribeInstancesAPIClient
	params    *DescribeInstancesInput
	nextToken *string
	firstPage bool
}

// NewDescribeInstancesPaginator returns a new DescribeInstancesPaginator
func NewDescribeInstancesPaginator(client DescribeInstancesAPIClient, params *DescribeInstancesInput, optFns ...func(*DescribeInstancesPaginatorOptions)) *DescribeInstancesPaginator {
	if params == nil {
		params = &DescribeInstancesInput{}
	}

	options := DescribeInstancesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeInstancesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeInstances page.
func (p *DescribeInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.DescribeInstances(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// InstanceStoppedWaiterOptions are waiter options for InstanceStoppedWaiter
type InstanceStoppedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// InstanceStoppedWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, InstanceStoppedWaiter will use default max delay of 120 seconds. Note
	// that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeInstancesInput, *DescribeInstancesOutput, error) (bool, error)
}

// InstanceStoppedWaiter defines the waiters for InstanceStopped
type InstanceStoppedWaiter struct {
	client DescribeInstancesAPIClient

	options InstanceStoppedWaiterOptions
}

// NewInstanceStoppedWaiter constructs a InstanceStoppedWaiter.
func NewInstanceStoppedWaiter(client DescribeInstancesAPIClient, optFns ...func(*InstanceStoppedWaiterOptions)) *InstanceStoppedWaiter {
	options := InstanceStoppedWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = instanceStoppedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &InstanceStoppedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for InstanceStopped waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *InstanceStoppedWaiter) Wait(ctx context.Context, params *DescribeInstancesInput, maxWaitDur time.Duration, optFns ...func(*InstanceStoppedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeInstances(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for InstanceStopped waiter")
}

func instanceStoppedStateRetryable(ctx context.Context, input *DescribeInstancesInput, output *DescribeInstancesOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Reservations[].Instances[].State.Name", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "stopped"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.InstanceStateName)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.InstanceStateName value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Reservations[].Instances[].State.Name", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "pending"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(types.InstanceStateName)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.InstanceStateName value, got %T", pathValue)
			}

			if string(value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Reservations[].Instances[].State.Name", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "terminated"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(types.InstanceStateName)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.InstanceStateName value, got %T", pathValue)
			}

			if string(value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

// InstanceTerminatedWaiterOptions are waiter options for InstanceTerminatedWaiter
type InstanceTerminatedWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// InstanceTerminatedWaiter will use default minimum delay of 15 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, InstanceTerminatedWaiter will use default max delay of 120 seconds.
	// Note that MaxDelay must resolve to value greater than or equal to the MinDelay.
	MaxDelay time.Duration

	// LogWaitAttempts is used to enable logging for waiter retry attempts
	LogWaitAttempts bool

	// Retryable is function that can be used to override the service defined
	// waiter-behavior based on operation output, or returned error. This function is
	// used by the waiter to decide if a state is retryable or a terminal state. By
	// default service-modeled logic will populate this option. This option can thus be
	// used to define a custom waiter state with fall-back to service-modeled waiter
	// state mutators.The function returns an error in case of a failure state. In case
	// of retry state, this function returns a bool value of true and nil error, while
	// in case of success it returns a bool value of false and nil error.
	Retryable func(context.Context, *DescribeInstancesInput, *DescribeInstancesOutput, error) (bool, error)
}

// InstanceTerminatedWaiter defines the waiters for InstanceTerminated
type InstanceTerminatedWaiter struct {
	client DescribeInstancesAPIClient

	options InstanceTerminatedWaiterOptions
}

// NewInstanceTerminatedWaiter constructs a InstanceTerminatedWaiter.
func NewInstanceTerminatedWaiter(client DescribeInstancesAPIClient, optFns ...func(*InstanceTerminatedWaiterOptions)) *InstanceTerminatedWaiter {
	options := InstanceTerminatedWaiterOptions{}
	options.MinDelay = 15 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = instanceTerminatedStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &InstanceTerminatedWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for InstanceTerminated waiter. The maxWaitDur is
// the maximum wait duration the waiter will wait. The maxWaitDur is required and
// must be greater than zero.
func (w *InstanceTerminatedWaiter) Wait(ctx context.Context, params *DescribeInstancesInput, maxWaitDur time.Duration, optFns ...func(*InstanceTerminatedWaiterOptions)) error {
	if maxWaitDur <= 0 {
		return fmt.Errorf("maximum wait time for waiter must be greater than zero")
	}

	options := w.options
	for _, fn := range optFns {
		fn(&options)
	}

	if options.MaxDelay <= 0 {
		options.MaxDelay = 120 * time.Second
	}

	if options.MinDelay > options.MaxDelay {
		return fmt.Errorf("minimum waiter delay %v must be lesser than or equal to maximum waiter delay of %v.", options.MinDelay, options.MaxDelay)
	}

	ctx, cancelFn := context.WithTimeout(ctx, maxWaitDur)
	defer cancelFn()

	logger := smithywaiter.Logger{}
	remainingTime := maxWaitDur

	var attempt int64
	for {

		attempt++
		apiOptions := options.APIOptions
		start := time.Now()

		if options.LogWaitAttempts {
			logger.Attempt = attempt
			apiOptions = append([]func(*middleware.Stack) error{}, options.APIOptions...)
			apiOptions = append(apiOptions, logger.AddLogger)
		}

		out, err := w.client.DescribeInstances(ctx, params, func(o *Options) {
			o.APIOptions = append(o.APIOptions, apiOptions...)
		})

		retryable, err := options.Retryable(ctx, params, out, err)
		if err != nil {
			return err
		}
		if !retryable {
			return nil
		}

		remainingTime -= time.Since(start)
		if remainingTime < options.MinDelay || remainingTime <= 0 {
			break
		}

		// compute exponential backoff between waiter retries
		delay, err := smithywaiter.ComputeDelay(
			attempt, options.MinDelay, options.MaxDelay, remainingTime,
		)
		if err != nil {
			return fmt.Errorf("error computing waiter delay, %w", err)
		}

		remainingTime -= delay
		// sleep for the delay amount before invoking a request
		if err := smithytime.SleepWithContext(ctx, delay); err != nil {
			return fmt.Errorf("request cancelled while waiting, %w", err)
		}
	}
	return fmt.Errorf("exceeded max wait time for InstanceTerminated waiter")
}

func instanceTerminatedStateRetryable(ctx context.Context, input *DescribeInstancesInput, output *DescribeInstancesOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Reservations[].Instances[].State.Name", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "terminated"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(types.InstanceStateName)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.InstanceStateName value, got %T", pathValue)
			}

			if string(value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Reservations[].Instances[].State.Name", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "pending"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(types.InstanceStateName)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.InstanceStateName value, got %T", pathValue)
			}

			if string(value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Reservations[].Instances[].State.Name", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "stopping"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(types.InstanceStateName)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected types.InstanceStateName value, got %T", pathValue)
			}

			if string(value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeInstances",
	}
}
