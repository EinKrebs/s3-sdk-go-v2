// Code generated by smithy-go-codegen DO NOT EDIT.

package backup

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns detailed information about recovery points of the type specified by a
// resource Amazon Resource Name (ARN).
func (c *Client) ListRecoveryPointsByResource(ctx context.Context, params *ListRecoveryPointsByResourceInput, optFns ...func(*Options)) (*ListRecoveryPointsByResourceOutput, error) {
	if params == nil {
		params = &ListRecoveryPointsByResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListRecoveryPointsByResource", params, optFns, addOperationListRecoveryPointsByResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListRecoveryPointsByResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListRecoveryPointsByResourceInput struct {

	// An ARN that uniquely identifies a resource. The format of the ARN depends on the
	// resource type.
	//
	// This member is required.
	ResourceArn *string

	// The maximum number of items to be returned.
	MaxResults *int32

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string
}

type ListRecoveryPointsByResourceOutput struct {

	// The next item following a partial list of returned items. For example, if a
	// request is made to return maxResults number of items, NextToken allows you to
	// return more items in your list starting at the location pointed to by the next
	// token.
	NextToken *string

	// An array of objects that contain detailed information about recovery points of
	// the specified resource type.
	RecoveryPoints []types.RecoveryPointByResource

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListRecoveryPointsByResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListRecoveryPointsByResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListRecoveryPointsByResource{}, middleware.After)
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
	if err = addOpListRecoveryPointsByResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListRecoveryPointsByResource(options.Region), middleware.Before); err != nil {
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

// ListRecoveryPointsByResourceAPIClient is a client that implements the
// ListRecoveryPointsByResource operation.
type ListRecoveryPointsByResourceAPIClient interface {
	ListRecoveryPointsByResource(context.Context, *ListRecoveryPointsByResourceInput, ...func(*Options)) (*ListRecoveryPointsByResourceOutput, error)
}

var _ ListRecoveryPointsByResourceAPIClient = (*Client)(nil)

// ListRecoveryPointsByResourcePaginatorOptions is the paginator options for
// ListRecoveryPointsByResource
type ListRecoveryPointsByResourcePaginatorOptions struct {
	// The maximum number of items to be returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListRecoveryPointsByResourcePaginator is a paginator for
// ListRecoveryPointsByResource
type ListRecoveryPointsByResourcePaginator struct {
	options   ListRecoveryPointsByResourcePaginatorOptions
	client    ListRecoveryPointsByResourceAPIClient
	params    *ListRecoveryPointsByResourceInput
	nextToken *string
	firstPage bool
}

// NewListRecoveryPointsByResourcePaginator returns a new
// ListRecoveryPointsByResourcePaginator
func NewListRecoveryPointsByResourcePaginator(client ListRecoveryPointsByResourceAPIClient, params *ListRecoveryPointsByResourceInput, optFns ...func(*ListRecoveryPointsByResourcePaginatorOptions)) *ListRecoveryPointsByResourcePaginator {
	if params == nil {
		params = &ListRecoveryPointsByResourceInput{}
	}

	options := ListRecoveryPointsByResourcePaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListRecoveryPointsByResourcePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListRecoveryPointsByResourcePaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListRecoveryPointsByResource page.
func (p *ListRecoveryPointsByResourcePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListRecoveryPointsByResourceOutput, error) {
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

	result, err := p.client.ListRecoveryPointsByResource(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListRecoveryPointsByResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "backup",
		OperationName: "ListRecoveryPointsByResource",
	}
}
