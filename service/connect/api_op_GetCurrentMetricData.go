// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Gets the real-time metric data from the specified Amazon Connect instance. For a
// description of each metric, see Real-time Metrics Definitions
// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html)
// in the Amazon Connect Administrator Guide.
func (c *Client) GetCurrentMetricData(ctx context.Context, params *GetCurrentMetricDataInput, optFns ...func(*Options)) (*GetCurrentMetricDataOutput, error) {
	if params == nil {
		params = &GetCurrentMetricDataInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetCurrentMetricData", params, optFns, addOperationGetCurrentMetricDataMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetCurrentMetricDataOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetCurrentMetricDataInput struct {

	// The metrics to retrieve. Specify the name and unit for each metric. The
	// following metrics are available. For a description of all the metrics, see
	// Real-time Metrics Definitions
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html)
	// in the Amazon Connect Administrator Guide. AGENTS_AFTER_CONTACT_WORK Unit: COUNT
	// Name in real-time metrics report: ACW
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html#aftercallwork-real-time)
	// AGENTS_AVAILABLE Unit: COUNT Name in real-time metrics report: Available
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html#available-real-time)
	// AGENTS_ERROR Unit: COUNT Name in real-time metrics report: Error
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html#error-real-time)
	// AGENTS_NON_PRODUCTIVE Unit: COUNT Name in real-time metrics report: NPT
	// (Non-Productive Time)
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html#non-productive-time-real-time)
	// AGENTS_ON_CALL Unit: COUNT Name in real-time metrics report: On contact
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html#on-call-real-time)
	// AGENTS_ON_CONTACT Unit: COUNT Name in real-time metrics report: On contact
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html#on-call-real-time)
	// AGENTS_ONLINE Unit: COUNT Name in real-time metrics report: Online
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html#online-real-time)
	// AGENTS_STAFFED Unit: COUNT Name in real-time metrics report: Staffed
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html#staffed-real-time)
	// CONTACTS_IN_QUEUE Unit: COUNT Name in real-time metrics report: In queue
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html#in-queue-real-time)
	// CONTACTS_SCHEDULED Unit: COUNT Name in real-time metrics report: Scheduled
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html#scheduled-real-time)
	// OLDEST_CONTACT_AGE Unit: SECONDS When you use groupings, Unit says SECONDS but
	// the Value is returned in MILLISECONDS. For example, if you get a response like
	// this: { "Metric": { "Name": "OLDEST_CONTACT_AGE", "Unit": "SECONDS" }, "Value":
	// 24113.0 } The actual OLDEST_CONTACT_AGE is 24 seconds. Name in real-time metrics
	// report: Oldest
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html#oldest-real-time)
	// SLOTS_ACTIVE Unit: COUNT Name in real-time metrics report: Active
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html#active-real-time)
	// SLOTS_AVAILABLE Unit: COUNT Name in real-time metrics report: Availability
	// (https://docs.aws.amazon.com/connect/latest/adminguide/real-time-metrics-definitions.html#availability-real-time)
	//
	// This member is required.
	CurrentMetrics []types.CurrentMetric

	// The queues, up to 100, or channels, to use to filter the metrics returned.
	// Metric data is retrieved only for the resources associated with the queues or
	// channels included in the filter. You can include both queue IDs and queue ARNs
	// in the same request. VOICE, CHAT, and TASK channels are supported.
	//
	// This member is required.
	Filters *types.Filters

	// The identifier of the Amazon Connect instance.
	//
	// This member is required.
	InstanceId *string

	// The grouping applied to the metrics returned. For example, when grouped by
	// QUEUE, the metrics returned apply to each queue rather than aggregated for all
	// queues. If you group by CHANNEL, you should include a Channels filter. VOICE,
	// CHAT, and TASK channels are supported. If no Grouping is included in the
	// request, a summary of metrics is returned.
	Groupings []types.Grouping

	// The maximimum number of results to return per page.
	MaxResults int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results. The token
	// expires after 5 minutes from the time it is created. Subsequent requests that
	// use the token must use the same request parameters as the request that generated
	// the token.
	NextToken *string
}

type GetCurrentMetricDataOutput struct {

	// The time at which the metrics were retrieved and cached for pagination.
	DataSnapshotTime *time.Time

	// Information about the real-time metrics.
	MetricResults []types.CurrentMetricResult

	// If there are additional results, this is the token for the next set of results.
	// The token expires after 5 minutes from the time it is created. Subsequent
	// requests that use the token must use the same request parameters as the request
	// that generated the token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetCurrentMetricDataMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetCurrentMetricData{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetCurrentMetricData{}, middleware.After)
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
	if err = addOpGetCurrentMetricDataValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetCurrentMetricData(options.Region), middleware.Before); err != nil {
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

// GetCurrentMetricDataAPIClient is a client that implements the
// GetCurrentMetricData operation.
type GetCurrentMetricDataAPIClient interface {
	GetCurrentMetricData(context.Context, *GetCurrentMetricDataInput, ...func(*Options)) (*GetCurrentMetricDataOutput, error)
}

var _ GetCurrentMetricDataAPIClient = (*Client)(nil)

// GetCurrentMetricDataPaginatorOptions is the paginator options for
// GetCurrentMetricData
type GetCurrentMetricDataPaginatorOptions struct {
	// The maximimum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetCurrentMetricDataPaginator is a paginator for GetCurrentMetricData
type GetCurrentMetricDataPaginator struct {
	options   GetCurrentMetricDataPaginatorOptions
	client    GetCurrentMetricDataAPIClient
	params    *GetCurrentMetricDataInput
	nextToken *string
	firstPage bool
}

// NewGetCurrentMetricDataPaginator returns a new GetCurrentMetricDataPaginator
func NewGetCurrentMetricDataPaginator(client GetCurrentMetricDataAPIClient, params *GetCurrentMetricDataInput, optFns ...func(*GetCurrentMetricDataPaginatorOptions)) *GetCurrentMetricDataPaginator {
	if params == nil {
		params = &GetCurrentMetricDataInput{}
	}

	options := GetCurrentMetricDataPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetCurrentMetricDataPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetCurrentMetricDataPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetCurrentMetricData page.
func (p *GetCurrentMetricDataPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetCurrentMetricDataOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetCurrentMetricData(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetCurrentMetricData(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "GetCurrentMetricData",
	}
}
