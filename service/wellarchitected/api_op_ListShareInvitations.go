// Code generated by smithy-go-codegen DO NOT EDIT.

package wellarchitected

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/wellarchitected/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// List the workload invitations.
func (c *Client) ListShareInvitations(ctx context.Context, params *ListShareInvitationsInput, optFns ...func(*Options)) (*ListShareInvitationsOutput, error) {
	if params == nil {
		params = &ListShareInvitationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListShareInvitations", params, optFns, addOperationListShareInvitationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListShareInvitationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Input for List Share Invitations
type ListShareInvitationsInput struct {

	// The maximum number of results to return for this request.
	MaxResults int32

	// The token to use to retrieve the next set of results.
	NextToken *string

	// An optional string added to the beginning of each workload name returned in the
	// results.
	WorkloadNamePrefix *string
}

// Input for List Share Invitations
type ListShareInvitationsOutput struct {

	// The token to use to retrieve the next set of results.
	NextToken *string

	// List of share invitation summaries in a workload.
	ShareInvitationSummaries []types.ShareInvitationSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListShareInvitationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListShareInvitations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListShareInvitations{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListShareInvitations(options.Region), middleware.Before); err != nil {
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

// ListShareInvitationsAPIClient is a client that implements the
// ListShareInvitations operation.
type ListShareInvitationsAPIClient interface {
	ListShareInvitations(context.Context, *ListShareInvitationsInput, ...func(*Options)) (*ListShareInvitationsOutput, error)
}

var _ ListShareInvitationsAPIClient = (*Client)(nil)

// ListShareInvitationsPaginatorOptions is the paginator options for
// ListShareInvitations
type ListShareInvitationsPaginatorOptions struct {
	// The maximum number of results to return for this request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListShareInvitationsPaginator is a paginator for ListShareInvitations
type ListShareInvitationsPaginator struct {
	options   ListShareInvitationsPaginatorOptions
	client    ListShareInvitationsAPIClient
	params    *ListShareInvitationsInput
	nextToken *string
	firstPage bool
}

// NewListShareInvitationsPaginator returns a new ListShareInvitationsPaginator
func NewListShareInvitationsPaginator(client ListShareInvitationsAPIClient, params *ListShareInvitationsInput, optFns ...func(*ListShareInvitationsPaginatorOptions)) *ListShareInvitationsPaginator {
	if params == nil {
		params = &ListShareInvitationsInput{}
	}

	options := ListShareInvitationsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListShareInvitationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListShareInvitationsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListShareInvitations page.
func (p *ListShareInvitationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListShareInvitationsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListShareInvitations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListShareInvitations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "wellarchitected",
		OperationName: "ListShareInvitations",
	}
}
