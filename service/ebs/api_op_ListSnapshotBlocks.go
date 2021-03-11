// Code generated by smithy-go-codegen DO NOT EDIT.

package ebs

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ebs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns information about the blocks in an Amazon Elastic Block Store snapshot.
func (c *Client) ListSnapshotBlocks(ctx context.Context, params *ListSnapshotBlocksInput, optFns ...func(*Options)) (*ListSnapshotBlocksOutput, error) {
	if params == nil {
		params = &ListSnapshotBlocksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSnapshotBlocks", params, optFns, addOperationListSnapshotBlocksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSnapshotBlocksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSnapshotBlocksInput struct {

	// The ID of the snapshot from which to get block indexes and block tokens.
	//
	// This member is required.
	SnapshotId *string

	// The number of results to return.
	MaxResults *int32

	// The token to request the next page of results.
	NextToken *string

	// The block index from which the list should start. The list in the response will
	// start from this block index or the next valid block index in the snapshot.
	StartingBlockIndex *int32
}

type ListSnapshotBlocksOutput struct {

	// The size of the block.
	BlockSize *int32

	// An array of objects containing information about the blocks.
	Blocks []types.Block

	// The time when the BlockToken expires.
	ExpiryTime *time.Time

	// The token to use to retrieve the next page of results. This value is null when
	// there are no more results to return.
	NextToken *string

	// The size of the volume in GB.
	VolumeSize *int64

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSnapshotBlocksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSnapshotBlocks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSnapshotBlocks{}, middleware.After)
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
	if err = addOpListSnapshotBlocksValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSnapshotBlocks(options.Region), middleware.Before); err != nil {
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

// ListSnapshotBlocksAPIClient is a client that implements the ListSnapshotBlocks
// operation.
type ListSnapshotBlocksAPIClient interface {
	ListSnapshotBlocks(context.Context, *ListSnapshotBlocksInput, ...func(*Options)) (*ListSnapshotBlocksOutput, error)
}

var _ ListSnapshotBlocksAPIClient = (*Client)(nil)

// ListSnapshotBlocksPaginatorOptions is the paginator options for
// ListSnapshotBlocks
type ListSnapshotBlocksPaginatorOptions struct {
	// The number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSnapshotBlocksPaginator is a paginator for ListSnapshotBlocks
type ListSnapshotBlocksPaginator struct {
	options   ListSnapshotBlocksPaginatorOptions
	client    ListSnapshotBlocksAPIClient
	params    *ListSnapshotBlocksInput
	nextToken *string
	firstPage bool
}

// NewListSnapshotBlocksPaginator returns a new ListSnapshotBlocksPaginator
func NewListSnapshotBlocksPaginator(client ListSnapshotBlocksAPIClient, params *ListSnapshotBlocksInput, optFns ...func(*ListSnapshotBlocksPaginatorOptions)) *ListSnapshotBlocksPaginator {
	if params == nil {
		params = &ListSnapshotBlocksInput{}
	}

	options := ListSnapshotBlocksPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSnapshotBlocksPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSnapshotBlocksPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListSnapshotBlocks page.
func (p *ListSnapshotBlocksPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSnapshotBlocksOutput, error) {
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

	result, err := p.client.ListSnapshotBlocks(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSnapshotBlocks(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ebs",
		OperationName: "ListSnapshotBlocks",
	}
}
