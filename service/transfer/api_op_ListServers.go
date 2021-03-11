// Code generated by smithy-go-codegen DO NOT EDIT.

package transfer

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the file transfer protocol-enabled servers that are associated with your
// AWS account.
func (c *Client) ListServers(ctx context.Context, params *ListServersInput, optFns ...func(*Options)) (*ListServersOutput, error) {
	if params == nil {
		params = &ListServersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListServers", params, optFns, addOperationListServersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListServersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListServersInput struct {

	// Specifies the number of servers to return as a response to the ListServers
	// query.
	MaxResults *int32

	// When additional results are obtained from the ListServers command, a NextToken
	// parameter is returned in the output. You can then pass the NextToken parameter
	// in a subsequent command to continue listing additional servers.
	NextToken *string
}

type ListServersOutput struct {

	// An array of servers that were listed.
	//
	// This member is required.
	Servers []types.ListedServer

	// When you can get additional results from the ListServers operation, a NextToken
	// parameter is returned in the output. In a following command, you can pass in the
	// NextToken parameter to continue listing additional servers.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListServersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListServers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListServers{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListServers(options.Region), middleware.Before); err != nil {
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

// ListServersAPIClient is a client that implements the ListServers operation.
type ListServersAPIClient interface {
	ListServers(context.Context, *ListServersInput, ...func(*Options)) (*ListServersOutput, error)
}

var _ ListServersAPIClient = (*Client)(nil)

// ListServersPaginatorOptions is the paginator options for ListServers
type ListServersPaginatorOptions struct {
	// Specifies the number of servers to return as a response to the ListServers
	// query.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListServersPaginator is a paginator for ListServers
type ListServersPaginator struct {
	options   ListServersPaginatorOptions
	client    ListServersAPIClient
	params    *ListServersInput
	nextToken *string
	firstPage bool
}

// NewListServersPaginator returns a new ListServersPaginator
func NewListServersPaginator(client ListServersAPIClient, params *ListServersInput, optFns ...func(*ListServersPaginatorOptions)) *ListServersPaginator {
	if params == nil {
		params = &ListServersInput{}
	}

	options := ListServersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListServersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListServersPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListServers page.
func (p *ListServersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListServersOutput, error) {
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

	result, err := p.client.ListServers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListServers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "transfer",
		OperationName: "ListServers",
	}
}
