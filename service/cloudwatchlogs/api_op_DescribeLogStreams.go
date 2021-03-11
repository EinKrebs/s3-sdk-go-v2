// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudwatchlogs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the log streams for the specified log group. You can list all the log
// streams or filter the results by prefix. You can also control how the results
// are ordered. This operation has a limit of five transactions per second, after
// which transactions are throttled.
func (c *Client) DescribeLogStreams(ctx context.Context, params *DescribeLogStreamsInput, optFns ...func(*Options)) (*DescribeLogStreamsOutput, error) {
	if params == nil {
		params = &DescribeLogStreamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLogStreams", params, optFns, addOperationDescribeLogStreamsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLogStreamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLogStreamsInput struct {

	// The name of the log group.
	//
	// This member is required.
	LogGroupName *string

	// If the value is true, results are returned in descending order. If the value is
	// to false, results are returned in ascending order. The default value is false.
	Descending *bool

	// The maximum number of items returned. If you don't specify a value, the default
	// is up to 50 items.
	Limit *int32

	// The prefix to match. If orderBy is LastEventTime, you cannot specify this
	// parameter.
	LogStreamNamePrefix *string

	// The token for the next set of items to return. (You received this token from a
	// previous call.)
	NextToken *string

	// If the value is LogStreamName, the results are ordered by log stream name. If
	// the value is LastEventTime, the results are ordered by the event time. The
	// default value is LogStreamName. If you order the results by event time, you
	// cannot specify the logStreamNamePrefix parameter. lastEventTimeStamp represents
	// the time of the most recent log event in the log stream in CloudWatch Logs. This
	// number is expressed as the number of milliseconds after Jan 1, 1970 00:00:00
	// UTC. lastEventTimeStamp updates on an eventual consistency basis. It typically
	// updates in less than an hour from ingestion, but in rare situations might take
	// longer.
	OrderBy types.OrderBy
}

type DescribeLogStreamsOutput struct {

	// The log streams.
	LogStreams []types.LogStream

	// The token for the next set of items to return. The token expires after 24 hours.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLogStreamsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeLogStreams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeLogStreams{}, middleware.After)
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
	if err = addOpDescribeLogStreamsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLogStreams(options.Region), middleware.Before); err != nil {
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

// DescribeLogStreamsAPIClient is a client that implements the DescribeLogStreams
// operation.
type DescribeLogStreamsAPIClient interface {
	DescribeLogStreams(context.Context, *DescribeLogStreamsInput, ...func(*Options)) (*DescribeLogStreamsOutput, error)
}

var _ DescribeLogStreamsAPIClient = (*Client)(nil)

// DescribeLogStreamsPaginatorOptions is the paginator options for
// DescribeLogStreams
type DescribeLogStreamsPaginatorOptions struct {
	// The maximum number of items returned. If you don't specify a value, the default
	// is up to 50 items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeLogStreamsPaginator is a paginator for DescribeLogStreams
type DescribeLogStreamsPaginator struct {
	options   DescribeLogStreamsPaginatorOptions
	client    DescribeLogStreamsAPIClient
	params    *DescribeLogStreamsInput
	nextToken *string
	firstPage bool
}

// NewDescribeLogStreamsPaginator returns a new DescribeLogStreamsPaginator
func NewDescribeLogStreamsPaginator(client DescribeLogStreamsAPIClient, params *DescribeLogStreamsInput, optFns ...func(*DescribeLogStreamsPaginatorOptions)) *DescribeLogStreamsPaginator {
	if params == nil {
		params = &DescribeLogStreamsInput{}
	}

	options := DescribeLogStreamsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeLogStreamsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeLogStreamsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeLogStreams page.
func (p *DescribeLogStreamsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeLogStreamsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeLogStreams(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeLogStreams(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "logs",
		OperationName: "DescribeLogStreams",
	}
}
