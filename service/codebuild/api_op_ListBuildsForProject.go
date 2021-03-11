// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of build IDs for the specified build project, with each build ID
// representing a single build.
func (c *Client) ListBuildsForProject(ctx context.Context, params *ListBuildsForProjectInput, optFns ...func(*Options)) (*ListBuildsForProjectOutput, error) {
	if params == nil {
		params = &ListBuildsForProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBuildsForProject", params, optFns, addOperationListBuildsForProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBuildsForProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBuildsForProjectInput struct {

	// The name of the AWS CodeBuild project.
	//
	// This member is required.
	ProjectName *string

	// During a previous call, if there are more than 100 items in the list, only the
	// first 100 items are returned, along with a unique string called a nextToken. To
	// get the next batch of items in the list, call this operation again, adding the
	// next token to the call. To get all of the items in the list, keep calling this
	// operation with each subsequent next token that is returned, until no more next
	// tokens are returned.
	NextToken *string

	// The order to list build IDs. Valid values include:
	//
	// * ASCENDING: List the build
	// IDs in ascending order by build ID.
	//
	// * DESCENDING: List the build IDs in
	// descending order by build ID.
	SortOrder types.SortOrderType
}

type ListBuildsForProjectOutput struct {

	// A list of build IDs for the specified build project, with each build ID
	// representing a single build.
	Ids []string

	// If there are more than 100 items in the list, only the first 100 items are
	// returned, along with a unique string called a nextToken. To get the next batch
	// of items in the list, call this operation again, adding the next token to the
	// call.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListBuildsForProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBuildsForProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBuildsForProject{}, middleware.After)
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
	if err = addOpListBuildsForProjectValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBuildsForProject(options.Region), middleware.Before); err != nil {
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

// ListBuildsForProjectAPIClient is a client that implements the
// ListBuildsForProject operation.
type ListBuildsForProjectAPIClient interface {
	ListBuildsForProject(context.Context, *ListBuildsForProjectInput, ...func(*Options)) (*ListBuildsForProjectOutput, error)
}

var _ ListBuildsForProjectAPIClient = (*Client)(nil)

// ListBuildsForProjectPaginatorOptions is the paginator options for
// ListBuildsForProject
type ListBuildsForProjectPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBuildsForProjectPaginator is a paginator for ListBuildsForProject
type ListBuildsForProjectPaginator struct {
	options   ListBuildsForProjectPaginatorOptions
	client    ListBuildsForProjectAPIClient
	params    *ListBuildsForProjectInput
	nextToken *string
	firstPage bool
}

// NewListBuildsForProjectPaginator returns a new ListBuildsForProjectPaginator
func NewListBuildsForProjectPaginator(client ListBuildsForProjectAPIClient, params *ListBuildsForProjectInput, optFns ...func(*ListBuildsForProjectPaginatorOptions)) *ListBuildsForProjectPaginator {
	if params == nil {
		params = &ListBuildsForProjectInput{}
	}

	options := ListBuildsForProjectPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBuildsForProjectPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBuildsForProjectPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListBuildsForProject page.
func (p *ListBuildsForProjectPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBuildsForProjectOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListBuildsForProject(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBuildsForProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "ListBuildsForProject",
	}
}
