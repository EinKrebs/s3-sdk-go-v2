// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves build resources for all builds associated with the AWS account in use.
// You can limit results to builds that are in a specific status by using the
// Status parameter. Use the pagination parameters to retrieve results in a set of
// sequential pages. Build resources are not listed in any particular order. Learn
// more  Upload a Custom Server Build
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-intro.html)
// Related operations
//
// * CreateBuild
//
// * ListBuilds
//
// * DescribeBuild
//
// *
// UpdateBuild
//
// * DeleteBuild
func (c *Client) ListBuilds(ctx context.Context, params *ListBuildsInput, optFns ...func(*Options)) (*ListBuildsOutput, error) {
	if params == nil {
		params = &ListBuildsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBuilds", params, optFns, addOperationListBuildsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBuildsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type ListBuildsInput struct {

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string

	// Build status to filter results by. To retrieve all builds, leave this parameter
	// empty. Possible build statuses include the following:
	//
	// * INITIALIZED -- A new
	// build has been defined, but no files have been uploaded. You cannot create
	// fleets for builds that are in this status. When a build is successfully created,
	// the build status is set to this value.
	//
	// * READY -- The game build has been
	// successfully uploaded. You can now create new fleets for this build.
	//
	// * FAILED
	// -- The game build upload failed. You cannot create new fleets for this build.
	Status types.BuildStatus
}

// Represents the returned data in response to a request operation.
type ListBuildsOutput struct {

	// A collection of build resources that match the request.
	Builds []types.Build

	// Token that indicates where to resume retrieving results on the next call to this
	// operation. If no token is returned, these results represent the end of the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBuildsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBuilds{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBuilds{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBuilds(options.Region), middleware.Before); err != nil {
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

// ListBuildsAPIClient is a client that implements the ListBuilds operation.
type ListBuildsAPIClient interface {
	ListBuilds(context.Context, *ListBuildsInput, ...func(*Options)) (*ListBuildsOutput, error)
}

var _ ListBuildsAPIClient = (*Client)(nil)

// ListBuildsPaginatorOptions is the paginator options for ListBuilds
type ListBuildsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBuildsPaginator is a paginator for ListBuilds
type ListBuildsPaginator struct {
	options   ListBuildsPaginatorOptions
	client    ListBuildsAPIClient
	params    *ListBuildsInput
	nextToken *string
	firstPage bool
}

// NewListBuildsPaginator returns a new ListBuildsPaginator
func NewListBuildsPaginator(client ListBuildsAPIClient, params *ListBuildsInput, optFns ...func(*ListBuildsPaginatorOptions)) *ListBuildsPaginator {
	if params == nil {
		params = &ListBuildsInput{}
	}

	options := ListBuildsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBuildsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBuildsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListBuilds page.
func (p *ListBuildsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBuildsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.ListBuilds(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBuilds(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ListBuilds",
	}
}
