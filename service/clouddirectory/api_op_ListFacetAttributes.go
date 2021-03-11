// Code generated by smithy-go-codegen DO NOT EDIT.

package clouddirectory

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/clouddirectory/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves attributes attached to the facet.
func (c *Client) ListFacetAttributes(ctx context.Context, params *ListFacetAttributesInput, optFns ...func(*Options)) (*ListFacetAttributesOutput, error) {
	if params == nil {
		params = &ListFacetAttributesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFacetAttributes", params, optFns, addOperationListFacetAttributesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFacetAttributesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFacetAttributesInput struct {

	// The name of the facet whose attributes will be retrieved.
	//
	// This member is required.
	Name *string

	// The ARN of the schema where the facet resides.
	//
	// This member is required.
	SchemaArn *string

	// The maximum number of results to retrieve.
	MaxResults *int32

	// The pagination token.
	NextToken *string
}

type ListFacetAttributesOutput struct {

	// The attributes attached to the facet.
	Attributes []types.FacetAttribute

	// The pagination token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFacetAttributesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListFacetAttributes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListFacetAttributes{}, middleware.After)
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
	if err = addOpListFacetAttributesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFacetAttributes(options.Region), middleware.Before); err != nil {
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

// ListFacetAttributesAPIClient is a client that implements the ListFacetAttributes
// operation.
type ListFacetAttributesAPIClient interface {
	ListFacetAttributes(context.Context, *ListFacetAttributesInput, ...func(*Options)) (*ListFacetAttributesOutput, error)
}

var _ ListFacetAttributesAPIClient = (*Client)(nil)

// ListFacetAttributesPaginatorOptions is the paginator options for
// ListFacetAttributes
type ListFacetAttributesPaginatorOptions struct {
	// The maximum number of results to retrieve.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFacetAttributesPaginator is a paginator for ListFacetAttributes
type ListFacetAttributesPaginator struct {
	options   ListFacetAttributesPaginatorOptions
	client    ListFacetAttributesAPIClient
	params    *ListFacetAttributesInput
	nextToken *string
	firstPage bool
}

// NewListFacetAttributesPaginator returns a new ListFacetAttributesPaginator
func NewListFacetAttributesPaginator(client ListFacetAttributesAPIClient, params *ListFacetAttributesInput, optFns ...func(*ListFacetAttributesPaginatorOptions)) *ListFacetAttributesPaginator {
	if params == nil {
		params = &ListFacetAttributesInput{}
	}

	options := ListFacetAttributesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFacetAttributesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFacetAttributesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListFacetAttributes page.
func (p *ListFacetAttributesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFacetAttributesOutput, error) {
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

	result, err := p.client.ListFacetAttributes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFacetAttributes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "clouddirectory",
		OperationName: "ListFacetAttributes",
	}
}
