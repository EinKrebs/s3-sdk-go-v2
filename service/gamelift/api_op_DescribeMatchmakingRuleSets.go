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

// Retrieves the details for FlexMatch matchmaking rule sets. You can request all
// existing rule sets for the Region, or provide a list of one or more rule set
// names. When requesting multiple items, use the pagination parameters to retrieve
// results as a set of sequential pages. If successful, a rule set is returned for
// each requested name. Learn more
//
// * Build a Rule Set
// (https://docs.aws.amazon.com/gamelift/latest/flexmatchguide/match-rulesets.html)
//
// Related
// operations
//
// * CreateMatchmakingConfiguration
//
// *
// DescribeMatchmakingConfigurations
//
// * UpdateMatchmakingConfiguration
//
// *
// DeleteMatchmakingConfiguration
//
// * CreateMatchmakingRuleSet
//
// *
// DescribeMatchmakingRuleSets
//
// * ValidateMatchmakingRuleSet
//
// *
// DeleteMatchmakingRuleSet
func (c *Client) DescribeMatchmakingRuleSets(ctx context.Context, params *DescribeMatchmakingRuleSetsInput, optFns ...func(*Options)) (*DescribeMatchmakingRuleSetsOutput, error) {
	if params == nil {
		params = &DescribeMatchmakingRuleSetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMatchmakingRuleSets", params, optFns, addOperationDescribeMatchmakingRuleSetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMatchmakingRuleSetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request operation.
type DescribeMatchmakingRuleSetsInput struct {

	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32

	// A list of one or more matchmaking rule set names to retrieve details for. (Note:
	// The rule set name is different from the optional "name" field in the rule set
	// body.) You can use either the rule set name or ARN value.
	Names []string

	// A token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this operation. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string
}

// Represents the returned data in response to a request operation.
type DescribeMatchmakingRuleSetsOutput struct {

	// A collection of requested matchmaking rule set objects.
	//
	// This member is required.
	RuleSets []types.MatchmakingRuleSet

	// A token that indicates where to resume retrieving results on the next call to
	// this operation. If no token is returned, these results represent the end of the
	// list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMatchmakingRuleSetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMatchmakingRuleSets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMatchmakingRuleSets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMatchmakingRuleSets(options.Region), middleware.Before); err != nil {
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

// DescribeMatchmakingRuleSetsAPIClient is a client that implements the
// DescribeMatchmakingRuleSets operation.
type DescribeMatchmakingRuleSetsAPIClient interface {
	DescribeMatchmakingRuleSets(context.Context, *DescribeMatchmakingRuleSetsInput, ...func(*Options)) (*DescribeMatchmakingRuleSetsOutput, error)
}

var _ DescribeMatchmakingRuleSetsAPIClient = (*Client)(nil)

// DescribeMatchmakingRuleSetsPaginatorOptions is the paginator options for
// DescribeMatchmakingRuleSets
type DescribeMatchmakingRuleSetsPaginatorOptions struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeMatchmakingRuleSetsPaginator is a paginator for
// DescribeMatchmakingRuleSets
type DescribeMatchmakingRuleSetsPaginator struct {
	options   DescribeMatchmakingRuleSetsPaginatorOptions
	client    DescribeMatchmakingRuleSetsAPIClient
	params    *DescribeMatchmakingRuleSetsInput
	nextToken *string
	firstPage bool
}

// NewDescribeMatchmakingRuleSetsPaginator returns a new
// DescribeMatchmakingRuleSetsPaginator
func NewDescribeMatchmakingRuleSetsPaginator(client DescribeMatchmakingRuleSetsAPIClient, params *DescribeMatchmakingRuleSetsInput, optFns ...func(*DescribeMatchmakingRuleSetsPaginatorOptions)) *DescribeMatchmakingRuleSetsPaginator {
	if params == nil {
		params = &DescribeMatchmakingRuleSetsInput{}
	}

	options := DescribeMatchmakingRuleSetsPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeMatchmakingRuleSetsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeMatchmakingRuleSetsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeMatchmakingRuleSets page.
func (p *DescribeMatchmakingRuleSetsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeMatchmakingRuleSetsOutput, error) {
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

	result, err := p.client.DescribeMatchmakingRuleSets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeMatchmakingRuleSets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "DescribeMatchmakingRuleSets",
	}
}
