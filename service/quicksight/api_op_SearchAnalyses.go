// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Searches for analyses that belong to the user specified in the filter.
func (c *Client) SearchAnalyses(ctx context.Context, params *SearchAnalysesInput, optFns ...func(*Options)) (*SearchAnalysesOutput, error) {
	if params == nil {
		params = &SearchAnalysesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchAnalyses", params, optFns, addOperationSearchAnalysesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchAnalysesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchAnalysesInput struct {

	// The ID of the AWS account that contains the analyses that you're searching for.
	//
	// This member is required.
	AwsAccountId *string

	// The structure for the search filters that you want to apply to your search.
	//
	// This member is required.
	Filters []types.AnalysisSearchFilter

	// The maximum number of results to return.
	MaxResults int32

	// A pagination token that can be used in a subsequent request.
	NextToken *string
}

type SearchAnalysesOutput struct {

	// Metadata describing the analyses that you searched for.
	AnalysisSummaryList []types.AnalysisSummary

	// A pagination token that can be used in a subsequent request.
	NextToken *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationSearchAnalysesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchAnalyses{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchAnalyses{}, middleware.After)
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
	if err = addOpSearchAnalysesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchAnalyses(options.Region), middleware.Before); err != nil {
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

// SearchAnalysesAPIClient is a client that implements the SearchAnalyses
// operation.
type SearchAnalysesAPIClient interface {
	SearchAnalyses(context.Context, *SearchAnalysesInput, ...func(*Options)) (*SearchAnalysesOutput, error)
}

var _ SearchAnalysesAPIClient = (*Client)(nil)

// SearchAnalysesPaginatorOptions is the paginator options for SearchAnalyses
type SearchAnalysesPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchAnalysesPaginator is a paginator for SearchAnalyses
type SearchAnalysesPaginator struct {
	options   SearchAnalysesPaginatorOptions
	client    SearchAnalysesAPIClient
	params    *SearchAnalysesInput
	nextToken *string
	firstPage bool
}

// NewSearchAnalysesPaginator returns a new SearchAnalysesPaginator
func NewSearchAnalysesPaginator(client SearchAnalysesAPIClient, params *SearchAnalysesInput, optFns ...func(*SearchAnalysesPaginatorOptions)) *SearchAnalysesPaginator {
	if params == nil {
		params = &SearchAnalysesInput{}
	}

	options := SearchAnalysesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchAnalysesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchAnalysesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next SearchAnalyses page.
func (p *SearchAnalysesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchAnalysesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.SearchAnalyses(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchAnalyses(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "SearchAnalyses",
	}
}
