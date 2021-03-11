// Code generated by smithy-go-codegen DO NOT EDIT.

package elastictranscoder

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elastictranscoder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// The ListPresets operation gets a list of the default presets included with
// Elastic Transcoder and the presets that you've added in an AWS region.
func (c *Client) ListPresets(ctx context.Context, params *ListPresetsInput, optFns ...func(*Options)) (*ListPresetsOutput, error) {
	if params == nil {
		params = &ListPresetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPresets", params, optFns, addOperationListPresetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPresetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The ListPresetsRequest structure.
type ListPresetsInput struct {

	// To list presets in chronological order by the date and time that they were
	// created, enter true. To list presets in reverse chronological order, enter
	// false.
	Ascending *string

	// When Elastic Transcoder returns more than one page of results, use pageToken in
	// subsequent GET requests to get each successive page of results.
	PageToken *string
}

// The ListPresetsResponse structure.
type ListPresetsOutput struct {

	// A value that you use to access the second and subsequent pages of results, if
	// any. When the presets fit on one page or when you've reached the last page of
	// results, the value of NextPageToken is null.
	NextPageToken *string

	// An array of Preset objects.
	Presets []types.Preset

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListPresetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListPresets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListPresets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPresets(options.Region), middleware.Before); err != nil {
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

// ListPresetsAPIClient is a client that implements the ListPresets operation.
type ListPresetsAPIClient interface {
	ListPresets(context.Context, *ListPresetsInput, ...func(*Options)) (*ListPresetsOutput, error)
}

var _ ListPresetsAPIClient = (*Client)(nil)

// ListPresetsPaginatorOptions is the paginator options for ListPresets
type ListPresetsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPresetsPaginator is a paginator for ListPresets
type ListPresetsPaginator struct {
	options   ListPresetsPaginatorOptions
	client    ListPresetsAPIClient
	params    *ListPresetsInput
	nextToken *string
	firstPage bool
}

// NewListPresetsPaginator returns a new ListPresetsPaginator
func NewListPresetsPaginator(client ListPresetsAPIClient, params *ListPresetsInput, optFns ...func(*ListPresetsPaginatorOptions)) *ListPresetsPaginator {
	if params == nil {
		params = &ListPresetsInput{}
	}

	options := ListPresetsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPresetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPresetsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListPresets page.
func (p *ListPresetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPresetsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.PageToken = p.nextToken

	result, err := p.client.ListPresets(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextPageToken

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListPresets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elastictranscoder",
		OperationName: "ListPresets",
	}
}
