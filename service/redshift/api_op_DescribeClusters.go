// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithytime "github.com/aws/smithy-go/time"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	smithywaiter "github.com/aws/smithy-go/waiter"
	"github.com/jmespath/go-jmespath"
	"time"
)

// Returns properties of provisioned clusters including general cluster properties,
// cluster database properties, maintenance and backup properties, and security and
// access properties. This operation supports pagination. For more information
// about managing clusters, go to Amazon Redshift Clusters
// (https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-clusters.html) in
// the Amazon Redshift Cluster Management Guide. If you specify both tag keys and
// tag values in the same request, Amazon Redshift returns all clusters that match
// any combination of the specified keys and values. For example, if you have owner
// and environment for tag keys, and admin and test for tag values, all clusters
// that have any combination of those values are returned. If both tag keys and
// values are omitted from the request, clusters are returned regardless of whether
// they have tag keys or values associated with them.
func (c *Client) DescribeClusters(ctx context.Context, params *DescribeClustersInput, optFns ...func(*Options)) (*DescribeClustersOutput, error) {
	if params == nil {
		params = &DescribeClustersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeClusters", params, optFns, addOperationDescribeClustersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeClustersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type DescribeClustersInput struct {

	// The unique identifier of a cluster whose properties you are requesting. This
	// parameter is case sensitive. The default is that all clusters defined for an
	// account are returned.
	ClusterIdentifier *string

	// An optional parameter that specifies the starting point to return a set of
	// response records. When the results of a DescribeClusters request exceed the
	// value specified in MaxRecords, AWS returns a value in the Marker field of the
	// response. You can retrieve the next set of response records by providing the
	// returned marker value in the Marker parameter and retrying the request.
	// Constraints: You can specify either the ClusterIdentifier parameter or the
	// Marker parameter, but not both.
	Marker *string

	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	MaxRecords *int32

	// A tag key or keys for which you want to return all matching clusters that are
	// associated with the specified key or keys. For example, suppose that you have
	// clusters that are tagged with keys called owner and environment. If you specify
	// both of these tag keys in the request, Amazon Redshift returns a response with
	// the clusters that have either or both of these tag keys associated with them.
	TagKeys []string

	// A tag value or values for which you want to return all matching clusters that
	// are associated with the specified tag value or values. For example, suppose that
	// you have clusters that are tagged with values called admin and test. If you
	// specify both of these tag values in the request, Amazon Redshift returns a
	// response with the clusters that have either or both of these tag values
	// associated with them.
	TagValues []string
}

// Contains the output from the DescribeClusters action.
type DescribeClustersOutput struct {

	// A list of Cluster objects, where each object describes one cluster.
	Clusters []types.Cluster

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the Marker
	// parameter and retrying the command. If the Marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeClustersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeClusters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeClusters{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeClusters(options.Region), middleware.Before); err != nil {
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

// DescribeClustersAPIClient is a client that implements the DescribeClusters
// operation.
type DescribeClustersAPIClient interface {
	DescribeClusters(context.Context, *DescribeClustersInput, ...func(*Options)) (*DescribeClustersOutput, error)
}

var _ DescribeClustersAPIClient = (*Client)(nil)

// DescribeClustersPaginatorOptions is the paginator options for DescribeClusters
type DescribeClustersPaginatorOptions struct {
	// The maximum number of response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value. Default: 100
	// Constraints: minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeClustersPaginator is a paginator for DescribeClusters
type DescribeClustersPaginator struct {
	options   DescribeClustersPaginatorOptions
	client    DescribeClustersAPIClient
	params    *DescribeClustersInput
	nextToken *string
	firstPage bool
}

// NewDescribeClustersPaginator returns a new DescribeClustersPaginator
func NewDescribeClustersPaginator(client DescribeClustersAPIClient, params *DescribeClustersInput, optFns ...func(*DescribeClustersPaginatorOptions)) *DescribeClustersPaginator {
	if params == nil {
		params = &DescribeClustersInput{}
	}

	options := DescribeClustersPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeClustersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeClustersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeClusters page.
func (p *DescribeClustersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeClustersOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeClusters(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

// ClusterRestoredWaiterOptions are waiter options for ClusterRestoredWaiter
type ClusterRestoredWaiterOptions struct {

	// Set of options to modify how an operation is invoked. These apply to all
	// operations invoked for this client. Use functional options on operation call to
	// modify this list for per operation behavior.
	APIOptions []func(*middleware.Stack) error

	// MinDelay is the minimum amount of time to delay between retries. If unset,
	// ClusterRestoredWaiter will use default minimum delay of 60 seconds. Note that
	// MinDelay must resolve to a value lesser than or equal to the MaxDelay.
	MinDelay time.Duration

	// MaxDelay is the maximum amount of time to delay between retries. If unset or set
	// to zero, ClusterRestoredWaiter will use default max delay of 120 seconds. Note
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
	Retryable func(context.Context, *DescribeClustersInput, *DescribeClustersOutput, error) (bool, error)
}

// ClusterRestoredWaiter defines the waiters for ClusterRestored
type ClusterRestoredWaiter struct {
	client DescribeClustersAPIClient

	options ClusterRestoredWaiterOptions
}

// NewClusterRestoredWaiter constructs a ClusterRestoredWaiter.
func NewClusterRestoredWaiter(client DescribeClustersAPIClient, optFns ...func(*ClusterRestoredWaiterOptions)) *ClusterRestoredWaiter {
	options := ClusterRestoredWaiterOptions{}
	options.MinDelay = 60 * time.Second
	options.MaxDelay = 120 * time.Second
	options.Retryable = clusterRestoredStateRetryable

	for _, fn := range optFns {
		fn(&options)
	}
	return &ClusterRestoredWaiter{
		client:  client,
		options: options,
	}
}

// Wait calls the waiter function for ClusterRestored waiter. The maxWaitDur is the
// maximum wait duration the waiter will wait. The maxWaitDur is required and must
// be greater than zero.
func (w *ClusterRestoredWaiter) Wait(ctx context.Context, params *DescribeClustersInput, maxWaitDur time.Duration, optFns ...func(*ClusterRestoredWaiterOptions)) error {
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

		out, err := w.client.DescribeClusters(ctx, params, func(o *Options) {
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
	return fmt.Errorf("exceeded max wait time for ClusterRestored waiter")
}

func clusterRestoredStateRetryable(ctx context.Context, input *DescribeClustersInput, output *DescribeClustersOutput, err error) (bool, error) {

	if err == nil {
		pathValue, err := jmespath.Search("Clusters[].RestoreStatus.Status", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "completed"
		var match = true
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		if len(listOfValues) == 0 {
			match = false
		}
		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) != expectedValue {
				match = false
			}
		}

		if match {
			return false, nil
		}
	}

	if err == nil {
		pathValue, err := jmespath.Search("Clusters[].ClusterStatus", output)
		if err != nil {
			return false, fmt.Errorf("error evaluating waiter state: %w", err)
		}

		expectedValue := "deleting"
		listOfValues, ok := pathValue.([]interface{})
		if !ok {
			return false, fmt.Errorf("waiter comparator expected list got %T", pathValue)
		}

		for _, v := range listOfValues {
			value, ok := v.(*string)
			if !ok {
				return false, fmt.Errorf("waiter comparator expected *string value, got %T", pathValue)
			}

			if string(*value) == expectedValue {
				return false, fmt.Errorf("waiter state transitioned to Failure")
			}
		}
	}

	return true, nil
}

func newServiceMetadataMiddleware_opDescribeClusters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeClusters",
	}
}
