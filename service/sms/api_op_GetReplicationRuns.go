// Code generated by smithy-go-codegen DO NOT EDIT.

package sms

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sms/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the replication runs for the specified replication job.
func (c *Client) GetReplicationRuns(ctx context.Context, params *GetReplicationRunsInput, optFns ...func(*Options)) (*GetReplicationRunsOutput, error) {
	if params == nil {
		params = &GetReplicationRunsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReplicationRuns", params, optFns, addOperationGetReplicationRunsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReplicationRunsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReplicationRunsInput struct {

	// The ID of the replication job.
	//
	// This member is required.
	ReplicationJobId *string

	// The maximum number of results to return in a single call. The default value is
	// 50. To retrieve the remaining results, make another call with the returned
	// NextToken value.
	MaxResults *int32

	// The token for the next set of results.
	NextToken *string
}

type GetReplicationRunsOutput struct {

	// The token required to retrieve the next set of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// Information about the replication job.
	ReplicationJob *types.ReplicationJob

	// Information about the replication runs.
	ReplicationRunList []types.ReplicationRun

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetReplicationRunsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetReplicationRuns{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetReplicationRuns{}, middleware.After)
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
	if err = addOpGetReplicationRunsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReplicationRuns(options.Region), middleware.Before); err != nil {
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

// GetReplicationRunsAPIClient is a client that implements the GetReplicationRuns
// operation.
type GetReplicationRunsAPIClient interface {
	GetReplicationRuns(context.Context, *GetReplicationRunsInput, ...func(*Options)) (*GetReplicationRunsOutput, error)
}

var _ GetReplicationRunsAPIClient = (*Client)(nil)

// GetReplicationRunsPaginatorOptions is the paginator options for
// GetReplicationRuns
type GetReplicationRunsPaginatorOptions struct {
	// The maximum number of results to return in a single call. The default value is
	// 50. To retrieve the remaining results, make another call with the returned
	// NextToken value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetReplicationRunsPaginator is a paginator for GetReplicationRuns
type GetReplicationRunsPaginator struct {
	options   GetReplicationRunsPaginatorOptions
	client    GetReplicationRunsAPIClient
	params    *GetReplicationRunsInput
	nextToken *string
	firstPage bool
}

// NewGetReplicationRunsPaginator returns a new GetReplicationRunsPaginator
func NewGetReplicationRunsPaginator(client GetReplicationRunsAPIClient, params *GetReplicationRunsInput, optFns ...func(*GetReplicationRunsPaginatorOptions)) *GetReplicationRunsPaginator {
	if params == nil {
		params = &GetReplicationRunsInput{}
	}

	options := GetReplicationRunsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetReplicationRunsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetReplicationRunsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetReplicationRuns page.
func (p *GetReplicationRunsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetReplicationRunsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.GetReplicationRuns(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetReplicationRuns(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms",
		OperationName: "GetReplicationRuns",
	}
}
