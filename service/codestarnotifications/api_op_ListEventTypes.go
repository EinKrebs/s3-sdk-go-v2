// Code generated by smithy-go-codegen DO NOT EDIT.

package codestarnotifications

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codestarnotifications/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the event types available for configuring
// notifications.
func (c *Client) ListEventTypes(ctx context.Context, params *ListEventTypesInput, optFns ...func(*Options)) (*ListEventTypesOutput, error) {
	if params == nil {
		params = &ListEventTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListEventTypes", params, optFns, addOperationListEventTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListEventTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListEventTypesInput struct {

	// The filters to use to return information by service or resource type.
	Filters []types.ListEventTypesFilter

	// A non-negative integer used to limit the number of returned results. The default
	// number is 50. The maximum number of results that can be returned is 100.
	MaxResults int32

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string
}

type ListEventTypesOutput struct {

	// Information about each event, including service name, resource type, event ID,
	// and event name.
	EventTypes []types.EventTypeSummary

	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListEventTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListEventTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListEventTypes{}, middleware.After)
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
	if err = addOpListEventTypesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListEventTypes(options.Region), middleware.Before); err != nil {
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

// ListEventTypesAPIClient is a client that implements the ListEventTypes
// operation.
type ListEventTypesAPIClient interface {
	ListEventTypes(context.Context, *ListEventTypesInput, ...func(*Options)) (*ListEventTypesOutput, error)
}

var _ ListEventTypesAPIClient = (*Client)(nil)

// ListEventTypesPaginatorOptions is the paginator options for ListEventTypes
type ListEventTypesPaginatorOptions struct {
	// A non-negative integer used to limit the number of returned results. The default
	// number is 50. The maximum number of results that can be returned is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListEventTypesPaginator is a paginator for ListEventTypes
type ListEventTypesPaginator struct {
	options   ListEventTypesPaginatorOptions
	client    ListEventTypesAPIClient
	params    *ListEventTypesInput
	nextToken *string
	firstPage bool
}

// NewListEventTypesPaginator returns a new ListEventTypesPaginator
func NewListEventTypesPaginator(client ListEventTypesAPIClient, params *ListEventTypesInput, optFns ...func(*ListEventTypesPaginatorOptions)) *ListEventTypesPaginator {
	if params == nil {
		params = &ListEventTypesInput{}
	}

	options := ListEventTypesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListEventTypesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListEventTypesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListEventTypes page.
func (p *ListEventTypesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListEventTypesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListEventTypes(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListEventTypes(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codestar-notifications",
		OperationName: "ListEventTypes",
	}
}
