// Code generated by smithy-go-codegen DO NOT EDIT.

package networkfirewall

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/networkfirewall/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the metadata for the firewall policies that you have defined.
// Depending on your setting for max results and the number of firewall policies, a
// single call might not return the full list.
func (c *Client) ListFirewallPolicies(ctx context.Context, params *ListFirewallPoliciesInput, optFns ...func(*Options)) (*ListFirewallPoliciesOutput, error) {
	if params == nil {
		params = &ListFirewallPoliciesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListFirewallPolicies", params, optFns, addOperationListFirewallPoliciesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListFirewallPoliciesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListFirewallPoliciesInput struct {

	// The maximum number of objects that you want Network Firewall to return for this
	// request. If more objects are available, in the response, Network Firewall
	// provides a NextToken value that you can use in a subsequent call to get the next
	// batch of objects.
	MaxResults *int32

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Network Firewall returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string
}

type ListFirewallPoliciesOutput struct {

	// The metadata for the firewall policies. Depending on your setting for max
	// results and the number of firewall policies that you have, this might not be the
	// full list.
	FirewallPolicies []types.FirewallPolicyMetadata

	// When you request a list of objects with a MaxResults setting, if the number of
	// objects that are still available for retrieval exceeds the maximum you
	// requested, Network Firewall returns a NextToken value in the response. To
	// retrieve the next batch of objects, use the token returned from the prior
	// request in your next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListFirewallPoliciesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListFirewallPolicies{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListFirewallPolicies{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListFirewallPolicies(options.Region), middleware.Before); err != nil {
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

// ListFirewallPoliciesAPIClient is a client that implements the
// ListFirewallPolicies operation.
type ListFirewallPoliciesAPIClient interface {
	ListFirewallPolicies(context.Context, *ListFirewallPoliciesInput, ...func(*Options)) (*ListFirewallPoliciesOutput, error)
}

var _ ListFirewallPoliciesAPIClient = (*Client)(nil)

// ListFirewallPoliciesPaginatorOptions is the paginator options for
// ListFirewallPolicies
type ListFirewallPoliciesPaginatorOptions struct {
	// The maximum number of objects that you want Network Firewall to return for this
	// request. If more objects are available, in the response, Network Firewall
	// provides a NextToken value that you can use in a subsequent call to get the next
	// batch of objects.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListFirewallPoliciesPaginator is a paginator for ListFirewallPolicies
type ListFirewallPoliciesPaginator struct {
	options   ListFirewallPoliciesPaginatorOptions
	client    ListFirewallPoliciesAPIClient
	params    *ListFirewallPoliciesInput
	nextToken *string
	firstPage bool
}

// NewListFirewallPoliciesPaginator returns a new ListFirewallPoliciesPaginator
func NewListFirewallPoliciesPaginator(client ListFirewallPoliciesAPIClient, params *ListFirewallPoliciesInput, optFns ...func(*ListFirewallPoliciesPaginatorOptions)) *ListFirewallPoliciesPaginator {
	if params == nil {
		params = &ListFirewallPoliciesInput{}
	}

	options := ListFirewallPoliciesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListFirewallPoliciesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListFirewallPoliciesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListFirewallPolicies page.
func (p *ListFirewallPoliciesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListFirewallPoliciesOutput, error) {
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

	result, err := p.client.ListFirewallPolicies(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListFirewallPolicies(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "network-firewall",
		OperationName: "ListFirewallPolicies",
	}
}
