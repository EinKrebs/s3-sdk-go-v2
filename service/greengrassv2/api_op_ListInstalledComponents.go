// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of the components that a AWS IoT Greengrass core
// device runs.
func (c *Client) ListInstalledComponents(ctx context.Context, params *ListInstalledComponentsInput, optFns ...func(*Options)) (*ListInstalledComponentsOutput, error) {
	if params == nil {
		params = &ListInstalledComponentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInstalledComponents", params, optFns, addOperationListInstalledComponentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInstalledComponentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInstalledComponentsInput struct {

	// The name of the core device. This is also the name of the AWS IoT thing.
	//
	// This member is required.
	CoreDeviceThingName *string

	// The maximum number of results to be returned per paginated request.
	MaxResults int32

	// The token to be used for the next set of paginated results.
	NextToken *string
}

type ListInstalledComponentsOutput struct {

	// A list that summarizes each component on the core device.
	InstalledComponents []types.InstalledComponent

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListInstalledComponentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListInstalledComponents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListInstalledComponents{}, middleware.After)
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
	if err = addOpListInstalledComponentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInstalledComponents(options.Region), middleware.Before); err != nil {
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

// ListInstalledComponentsAPIClient is a client that implements the
// ListInstalledComponents operation.
type ListInstalledComponentsAPIClient interface {
	ListInstalledComponents(context.Context, *ListInstalledComponentsInput, ...func(*Options)) (*ListInstalledComponentsOutput, error)
}

var _ ListInstalledComponentsAPIClient = (*Client)(nil)

// ListInstalledComponentsPaginatorOptions is the paginator options for
// ListInstalledComponents
type ListInstalledComponentsPaginatorOptions struct {
	// The maximum number of results to be returned per paginated request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInstalledComponentsPaginator is a paginator for ListInstalledComponents
type ListInstalledComponentsPaginator struct {
	options   ListInstalledComponentsPaginatorOptions
	client    ListInstalledComponentsAPIClient
	params    *ListInstalledComponentsInput
	nextToken *string
	firstPage bool
}

// NewListInstalledComponentsPaginator returns a new
// ListInstalledComponentsPaginator
func NewListInstalledComponentsPaginator(client ListInstalledComponentsAPIClient, params *ListInstalledComponentsInput, optFns ...func(*ListInstalledComponentsPaginatorOptions)) *ListInstalledComponentsPaginator {
	if params == nil {
		params = &ListInstalledComponentsInput{}
	}

	options := ListInstalledComponentsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInstalledComponentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInstalledComponentsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListInstalledComponents page.
func (p *ListInstalledComponentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInstalledComponentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListInstalledComponents(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListInstalledComponents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ListInstalledComponents",
	}
}
