// Code generated by smithy-go-codegen DO NOT EDIT.

package resourcegroupstaggingapi

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all tag keys in the specified Region for the AWS account.
func (c *Client) GetTagKeys(ctx context.Context, params *GetTagKeysInput, optFns ...func(*Options)) (*GetTagKeysOutput, error) {
	if params == nil {
		params = &GetTagKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetTagKeys", params, optFns, addOperationGetTagKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetTagKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetTagKeysInput struct {
	MaxResults *int32

	// A string that indicates that additional data is available. Leave this value
	// empty for your initial request. If the response includes a PaginationToken, use
	// that string for this value to request an additional page of data.
	PaginationToken *string
}

type GetTagKeysOutput struct {

	// A string that indicates that the response contains more data than can be
	// returned in a single response. To receive additional data, specify this string
	// for the PaginationToken value in a subsequent request.
	PaginationToken *string

	// A list of all tag keys in the AWS account.
	TagKeys []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetTagKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetTagKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetTagKeys{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetTagKeys(options.Region), middleware.Before); err != nil {
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

// GetTagKeysAPIClient is a client that implements the GetTagKeys operation.
type GetTagKeysAPIClient interface {
	GetTagKeys(context.Context, *GetTagKeysInput, ...func(*Options)) (*GetTagKeysOutput, error)
}

var _ GetTagKeysAPIClient = (*Client)(nil)

// GetTagKeysPaginatorOptions is the paginator options for GetTagKeys
type GetTagKeysPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetTagKeysPaginator is a paginator for GetTagKeys
type GetTagKeysPaginator struct {
	options   GetTagKeysPaginatorOptions
	client    GetTagKeysAPIClient
	params    *GetTagKeysInput
	nextToken *string
	firstPage bool
}

// NewGetTagKeysPaginator returns a new GetTagKeysPaginator
func NewGetTagKeysPaginator(client GetTagKeysAPIClient, params *GetTagKeysInput, optFns ...func(*GetTagKeysPaginatorOptions)) *GetTagKeysPaginator {
	if params == nil {
		params = &GetTagKeysInput{}
	}

	options := GetTagKeysPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetTagKeysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetTagKeysPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next GetTagKeys page.
func (p *GetTagKeysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetTagKeysOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PaginationToken = p.nextToken

	result, err := p.client.GetTagKeys(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.PaginationToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opGetTagKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tagging",
		OperationName: "GetTagKeys",
	}
}
