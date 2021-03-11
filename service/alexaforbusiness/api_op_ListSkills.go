// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/alexaforbusiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists all enabled skills in a specific skill group.
func (c *Client) ListSkills(ctx context.Context, params *ListSkillsInput, optFns ...func(*Options)) (*ListSkillsOutput, error) {
	if params == nil {
		params = &ListSkillsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSkills", params, optFns, addOperationListSkillsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSkillsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSkillsInput struct {

	// Whether the skill is enabled under the user's account.
	EnablementType types.EnablementTypeFilter

	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	MaxResults *int32

	// An optional token returned from a prior request. Use this token for pagination
	// of results from this action. If this parameter is specified, the response
	// includes only results beyond the token, up to the value specified by MaxResults.
	NextToken *string

	// The ARN of the skill group for which to list enabled skills.
	SkillGroupArn *string

	// Whether the skill is publicly available or is a private skill.
	SkillType types.SkillTypeFilter
}

type ListSkillsOutput struct {

	// The token returned to indicate that there is more data available.
	NextToken *string

	// The list of enabled skills requested. Required.
	SkillSummaries []types.SkillSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSkillsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListSkills{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListSkills{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSkills(options.Region), middleware.Before); err != nil {
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

// ListSkillsAPIClient is a client that implements the ListSkills operation.
type ListSkillsAPIClient interface {
	ListSkills(context.Context, *ListSkillsInput, ...func(*Options)) (*ListSkillsOutput, error)
}

var _ ListSkillsAPIClient = (*Client)(nil)

// ListSkillsPaginatorOptions is the paginator options for ListSkills
type ListSkillsPaginatorOptions struct {
	// The maximum number of results to include in the response. If more results exist
	// than the specified MaxResults value, a token is included in the response so that
	// the remaining results can be retrieved.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSkillsPaginator is a paginator for ListSkills
type ListSkillsPaginator struct {
	options   ListSkillsPaginatorOptions
	client    ListSkillsAPIClient
	params    *ListSkillsInput
	nextToken *string
	firstPage bool
}

// NewListSkillsPaginator returns a new ListSkillsPaginator
func NewListSkillsPaginator(client ListSkillsAPIClient, params *ListSkillsInput, optFns ...func(*ListSkillsPaginatorOptions)) *ListSkillsPaginator {
	if params == nil {
		params = &ListSkillsInput{}
	}

	options := ListSkillsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSkillsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSkillsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListSkills page.
func (p *ListSkillsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSkillsOutput, error) {
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

	result, err := p.client.ListSkills(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSkills(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "ListSkills",
	}
}
