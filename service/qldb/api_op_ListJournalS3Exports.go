// Code generated by smithy-go-codegen DO NOT EDIT.

package qldb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/qldb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns an array of journal export job descriptions for all ledgers that are
// associated with the current AWS account and Region. This action returns a
// maximum of MaxResults items, and is paginated so that you can retrieve all the
// items by calling ListJournalS3Exports multiple times. This action does not
// return any expired export jobs. For more information, see Export Job Expiration
// (https://docs.aws.amazon.com/qldb/latest/developerguide/export-journal.request.html#export-journal.request.expiration)
// in the Amazon QLDB Developer Guide.
func (c *Client) ListJournalS3Exports(ctx context.Context, params *ListJournalS3ExportsInput, optFns ...func(*Options)) (*ListJournalS3ExportsOutput, error) {
	if params == nil {
		params = &ListJournalS3ExportsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListJournalS3Exports", params, optFns, addOperationListJournalS3ExportsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListJournalS3ExportsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListJournalS3ExportsInput struct {

	// The maximum number of results to return in a single ListJournalS3Exports
	// request. (The actual number of results returned might be fewer.)
	MaxResults *int32

	// A pagination token, indicating that you want to retrieve the next page of
	// results. If you received a value for NextToken in the response from a previous
	// ListJournalS3Exports call, then you should use that value as input here.
	NextToken *string
}

type ListJournalS3ExportsOutput struct {

	// The array of journal export job descriptions for all ledgers that are associated
	// with the current AWS account and Region.
	JournalS3Exports []types.JournalS3ExportDescription

	// * If NextToken is empty, then the last page of results has been processed and
	// there are no more results to be retrieved.
	//
	// * If NextToken is not empty, then
	// there are more results available. To retrieve the next page of results, use the
	// value of NextToken in a subsequent ListJournalS3Exports call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListJournalS3ExportsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListJournalS3Exports{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListJournalS3Exports{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListJournalS3Exports(options.Region), middleware.Before); err != nil {
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

// ListJournalS3ExportsAPIClient is a client that implements the
// ListJournalS3Exports operation.
type ListJournalS3ExportsAPIClient interface {
	ListJournalS3Exports(context.Context, *ListJournalS3ExportsInput, ...func(*Options)) (*ListJournalS3ExportsOutput, error)
}

var _ ListJournalS3ExportsAPIClient = (*Client)(nil)

// ListJournalS3ExportsPaginatorOptions is the paginator options for
// ListJournalS3Exports
type ListJournalS3ExportsPaginatorOptions struct {
	// The maximum number of results to return in a single ListJournalS3Exports
	// request. (The actual number of results returned might be fewer.)
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListJournalS3ExportsPaginator is a paginator for ListJournalS3Exports
type ListJournalS3ExportsPaginator struct {
	options   ListJournalS3ExportsPaginatorOptions
	client    ListJournalS3ExportsAPIClient
	params    *ListJournalS3ExportsInput
	nextToken *string
	firstPage bool
}

// NewListJournalS3ExportsPaginator returns a new ListJournalS3ExportsPaginator
func NewListJournalS3ExportsPaginator(client ListJournalS3ExportsAPIClient, params *ListJournalS3ExportsInput, optFns ...func(*ListJournalS3ExportsPaginatorOptions)) *ListJournalS3ExportsPaginator {
	if params == nil {
		params = &ListJournalS3ExportsInput{}
	}

	options := ListJournalS3ExportsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListJournalS3ExportsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListJournalS3ExportsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListJournalS3Exports page.
func (p *ListJournalS3ExportsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListJournalS3ExportsOutput, error) {
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

	result, err := p.client.ListJournalS3Exports(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListJournalS3Exports(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "qldb",
		OperationName: "ListJournalS3Exports",
	}
}
