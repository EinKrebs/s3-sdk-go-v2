// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of simulation applications. You can optionally provide filters to
// retrieve specific simulation applications.
func (c *Client) ListSimulationApplications(ctx context.Context, params *ListSimulationApplicationsInput, optFns ...func(*Options)) (*ListSimulationApplicationsOutput, error) {
	if params == nil {
		params = &ListSimulationApplicationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSimulationApplications", params, optFns, addOperationListSimulationApplicationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSimulationApplicationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSimulationApplicationsInput struct {

	// Optional list of filters to limit results. The filter name name is supported.
	// When filtering, you must use the complete value of the filtered item. You can
	// use up to three filters.
	Filters []types.Filter

	// When this parameter is used, ListSimulationApplications only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another
	// ListSimulationApplications request with the returned nextToken value. This value
	// can be between 1 and 100. If this parameter is not used, then
	// ListSimulationApplications returns up to 100 results and a nextToken value if
	// applicable.
	MaxResults *int32

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListSimulationApplications again and assign that
	// token to the request object's nextToken parameter. If there are no remaining
	// results, the previous response object's NextToken parameter is set to null.
	NextToken *string

	// The version qualifier of the simulation application.
	VersionQualifier *string
}

type ListSimulationApplicationsOutput struct {

	// If the previous paginated request did not return all of the remaining results,
	// the response object's nextToken parameter value is set to a token. To retrieve
	// the next set of results, call ListSimulationApplications again and assign that
	// token to the request object's nextToken parameter. If there are no remaining
	// results, the previous response object's NextToken parameter is set to null.
	NextToken *string

	// A list of simulation application summaries that meet the criteria of the
	// request.
	SimulationApplicationSummaries []types.SimulationApplicationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSimulationApplicationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSimulationApplications{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSimulationApplications{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSimulationApplications(options.Region), middleware.Before); err != nil {
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

// ListSimulationApplicationsAPIClient is a client that implements the
// ListSimulationApplications operation.
type ListSimulationApplicationsAPIClient interface {
	ListSimulationApplications(context.Context, *ListSimulationApplicationsInput, ...func(*Options)) (*ListSimulationApplicationsOutput, error)
}

var _ ListSimulationApplicationsAPIClient = (*Client)(nil)

// ListSimulationApplicationsPaginatorOptions is the paginator options for
// ListSimulationApplications
type ListSimulationApplicationsPaginatorOptions struct {
	// When this parameter is used, ListSimulationApplications only returns maxResults
	// results in a single page along with a nextToken response element. The remaining
	// results of the initial request can be seen by sending another
	// ListSimulationApplications request with the returned nextToken value. This value
	// can be between 1 and 100. If this parameter is not used, then
	// ListSimulationApplications returns up to 100 results and a nextToken value if
	// applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSimulationApplicationsPaginator is a paginator for
// ListSimulationApplications
type ListSimulationApplicationsPaginator struct {
	options   ListSimulationApplicationsPaginatorOptions
	client    ListSimulationApplicationsAPIClient
	params    *ListSimulationApplicationsInput
	nextToken *string
	firstPage bool
}

// NewListSimulationApplicationsPaginator returns a new
// ListSimulationApplicationsPaginator
func NewListSimulationApplicationsPaginator(client ListSimulationApplicationsAPIClient, params *ListSimulationApplicationsInput, optFns ...func(*ListSimulationApplicationsPaginatorOptions)) *ListSimulationApplicationsPaginator {
	if params == nil {
		params = &ListSimulationApplicationsInput{}
	}

	options := ListSimulationApplicationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSimulationApplicationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSimulationApplicationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListSimulationApplications page.
func (p *ListSimulationApplicationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSimulationApplicationsOutput, error) {
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

	result, err := p.client.ListSimulationApplications(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSimulationApplications(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "ListSimulationApplications",
	}
}
