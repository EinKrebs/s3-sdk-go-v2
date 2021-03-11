// Code generated by smithy-go-codegen DO NOT EDIT.

package appconfig

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appconfig/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the deployments for an environment.
func (c *Client) ListDeployments(ctx context.Context, params *ListDeploymentsInput, optFns ...func(*Options)) (*ListDeploymentsOutput, error) {
	if params == nil {
		params = &ListDeploymentsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDeployments", params, optFns, addOperationListDeploymentsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDeploymentsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDeploymentsInput struct {

	// The application ID.
	//
	// This member is required.
	ApplicationId *string

	// The environment ID.
	//
	// This member is required.
	EnvironmentId *string

	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	MaxResults int32

	// A token to start the list. Use this token to get the next set of results.
	NextToken *string
}

type ListDeploymentsOutput struct {

	// The elements from this collection.
	Items []types.DeploymentSummary

	// The token for the next set of items to return. Use this token to get the next
	// set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDeploymentsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDeployments{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDeployments{}, middleware.After)
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
	if err = addOpListDeploymentsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDeployments(options.Region), middleware.Before); err != nil {
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

// ListDeploymentsAPIClient is a client that implements the ListDeployments
// operation.
type ListDeploymentsAPIClient interface {
	ListDeployments(context.Context, *ListDeploymentsInput, ...func(*Options)) (*ListDeploymentsOutput, error)
}

var _ ListDeploymentsAPIClient = (*Client)(nil)

// ListDeploymentsPaginatorOptions is the paginator options for ListDeployments
type ListDeploymentsPaginatorOptions struct {
	// The maximum number of items to return for this call. The call also returns a
	// token that you can specify in a subsequent call to get the next set of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDeploymentsPaginator is a paginator for ListDeployments
type ListDeploymentsPaginator struct {
	options   ListDeploymentsPaginatorOptions
	client    ListDeploymentsAPIClient
	params    *ListDeploymentsInput
	nextToken *string
	firstPage bool
}

// NewListDeploymentsPaginator returns a new ListDeploymentsPaginator
func NewListDeploymentsPaginator(client ListDeploymentsAPIClient, params *ListDeploymentsInput, optFns ...func(*ListDeploymentsPaginatorOptions)) *ListDeploymentsPaginator {
	if params == nil {
		params = &ListDeploymentsInput{}
	}

	options := ListDeploymentsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDeploymentsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDeploymentsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListDeployments page.
func (p *ListDeploymentsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDeploymentsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListDeployments(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDeployments(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appconfig",
		OperationName: "ListDeployments",
	}
}
