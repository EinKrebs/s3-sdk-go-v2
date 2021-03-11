// Code generated by smithy-go-codegen DO NOT EDIT.

package kafka

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of the Scram Secrets associated with an Amazon MSK cluster.
func (c *Client) ListScramSecrets(ctx context.Context, params *ListScramSecretsInput, optFns ...func(*Options)) (*ListScramSecretsOutput, error) {
	if params == nil {
		params = &ListScramSecretsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListScramSecrets", params, optFns, addOperationListScramSecretsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListScramSecretsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListScramSecretsInput struct {

	// The arn of the cluster.
	//
	// This member is required.
	ClusterArn *string

	// The maxResults of the query.
	MaxResults int32

	// The nextToken of the query.
	NextToken *string
}

type ListScramSecretsOutput struct {

	// Paginated results marker.
	NextToken *string

	// The list of scram secrets associated with the cluster.
	SecretArnList []string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListScramSecretsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListScramSecrets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListScramSecrets{}, middleware.After)
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
	if err = addOpListScramSecretsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListScramSecrets(options.Region), middleware.Before); err != nil {
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

// ListScramSecretsAPIClient is a client that implements the ListScramSecrets
// operation.
type ListScramSecretsAPIClient interface {
	ListScramSecrets(context.Context, *ListScramSecretsInput, ...func(*Options)) (*ListScramSecretsOutput, error)
}

var _ ListScramSecretsAPIClient = (*Client)(nil)

// ListScramSecretsPaginatorOptions is the paginator options for ListScramSecrets
type ListScramSecretsPaginatorOptions struct {
	// The maxResults of the query.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListScramSecretsPaginator is a paginator for ListScramSecrets
type ListScramSecretsPaginator struct {
	options   ListScramSecretsPaginatorOptions
	client    ListScramSecretsAPIClient
	params    *ListScramSecretsInput
	nextToken *string
	firstPage bool
}

// NewListScramSecretsPaginator returns a new ListScramSecretsPaginator
func NewListScramSecretsPaginator(client ListScramSecretsAPIClient, params *ListScramSecretsInput, optFns ...func(*ListScramSecretsPaginatorOptions)) *ListScramSecretsPaginator {
	if params == nil {
		params = &ListScramSecretsInput{}
	}

	options := ListScramSecretsPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListScramSecretsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListScramSecretsPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListScramSecrets page.
func (p *ListScramSecretsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListScramSecretsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListScramSecrets(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListScramSecrets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kafka",
		OperationName: "ListScramSecrets",
	}
}
