// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists gateways owned by an AWS account in an AWS Region specified in the
// request. The returned list is ordered by gateway Amazon Resource Name (ARN). By
// default, the operation returns a maximum of 100 gateways. This operation
// supports pagination that allows you to optionally reduce the number of gateways
// returned in a response. If you have more gateways than are returned in a
// response (that is, the response returns only a truncated list of your gateways),
// the response contains a marker that you can specify in your next request to
// fetch the next page of gateways.
func (c *Client) ListGateways(ctx context.Context, params *ListGatewaysInput, optFns ...func(*Options)) (*ListGatewaysOutput, error) {
	if params == nil {
		params = &ListGatewaysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGateways", params, optFns, addOperationListGatewaysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGatewaysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A JSON object containing zero or more of the following fields:
//
// *
// ListGatewaysInput$Limit
//
// * ListGatewaysInput$Marker
type ListGatewaysInput struct {

	// Specifies that the list of gateways returned be limited to the specified number
	// of items.
	Limit *int32

	// An opaque string that indicates the position at which to begin the returned list
	// of gateways.
	Marker *string
}

type ListGatewaysOutput struct {

	// An array of GatewayInfo objects.
	Gateways []types.GatewayInfo

	// Use the marker in your next request to fetch the next set of gateways in the
	// list. If there are no more gateways to list, this field does not appear in the
	// response.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListGatewaysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListGateways{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListGateways{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGateways(options.Region), middleware.Before); err != nil {
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

// ListGatewaysAPIClient is a client that implements the ListGateways operation.
type ListGatewaysAPIClient interface {
	ListGateways(context.Context, *ListGatewaysInput, ...func(*Options)) (*ListGatewaysOutput, error)
}

var _ ListGatewaysAPIClient = (*Client)(nil)

// ListGatewaysPaginatorOptions is the paginator options for ListGateways
type ListGatewaysPaginatorOptions struct {
	// Specifies that the list of gateways returned be limited to the specified number
	// of items.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGatewaysPaginator is a paginator for ListGateways
type ListGatewaysPaginator struct {
	options   ListGatewaysPaginatorOptions
	client    ListGatewaysAPIClient
	params    *ListGatewaysInput
	nextToken *string
	firstPage bool
}

// NewListGatewaysPaginator returns a new ListGatewaysPaginator
func NewListGatewaysPaginator(client ListGatewaysAPIClient, params *ListGatewaysInput, optFns ...func(*ListGatewaysPaginatorOptions)) *ListGatewaysPaginator {
	if params == nil {
		params = &ListGatewaysInput{}
	}

	options := ListGatewaysPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGatewaysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGatewaysPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListGateways page.
func (p *ListGatewaysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGatewaysOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListGateways(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGateways(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "ListGateways",
	}
}
