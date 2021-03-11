// Code generated by smithy-go-codegen DO NOT EDIT.

package appmesh

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appmesh/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of existing virtual services in a service mesh.
func (c *Client) ListVirtualServices(ctx context.Context, params *ListVirtualServicesInput, optFns ...func(*Options)) (*ListVirtualServicesOutput, error) {
	if params == nil {
		params = &ListVirtualServicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListVirtualServices", params, optFns, addOperationListVirtualServicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListVirtualServicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

//
type ListVirtualServicesInput struct {

	// The name of the service mesh to list virtual services in.
	//
	// This member is required.
	MeshName *string

	// The maximum number of results returned by ListVirtualServices in paginated
	// output. When you use this parameter, ListVirtualServices returns only limit
	// results in a single page along with a nextToken response element. You can see
	// the remaining results of the initial request by sending another
	// ListVirtualServices request with the returned nextToken value. This value can be
	// between 1 and 100. If you don't use this parameter, ListVirtualServices returns
	// up to 100 results and a nextToken value if applicable.
	Limit *int32

	// The AWS IAM account ID of the service mesh owner. If the account ID is not your
	// own, then it's the ID of the account that shared the mesh with your account. For
	// more information about mesh sharing, see Working with shared meshes
	// (https://docs.aws.amazon.com/app-mesh/latest/userguide/sharing.html).
	MeshOwner *string

	// The nextToken value returned from a previous paginated ListVirtualServices
	// request where limit was used and the results exceeded the value of that
	// parameter. Pagination continues from the end of the previous results that
	// returned the nextToken value.
	NextToken *string
}

//
type ListVirtualServicesOutput struct {

	// The list of existing virtual services for the specified service mesh.
	//
	// This member is required.
	VirtualServices []types.VirtualServiceRef

	// The nextToken value to include in a future ListVirtualServices request. When the
	// results of a ListVirtualServices request exceed limit, you can use this value to
	// retrieve the next page of results. This value is null when there are no more
	// results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListVirtualServicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListVirtualServices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListVirtualServices{}, middleware.After)
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
	if err = addOpListVirtualServicesValidationMiddleware(stack); err != nil {
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

// ListVirtualServicesAPIClient is a client that implements the ListVirtualServices
// operation.
type ListVirtualServicesAPIClient interface {
	ListVirtualServices(context.Context, *ListVirtualServicesInput, ...func(*Options)) (*ListVirtualServicesOutput, error)
}

var _ ListVirtualServicesAPIClient = (*Client)(nil)

// ListVirtualServicesPaginatorOptions is the paginator options for
// ListVirtualServices
type ListVirtualServicesPaginatorOptions struct {
	// The maximum number of results returned by ListVirtualServices in paginated
	// output. When you use this parameter, ListVirtualServices returns only limit
	// results in a single page along with a nextToken response element. You can see
	// the remaining results of the initial request by sending another
	// ListVirtualServices request with the returned nextToken value. This value can be
	// between 1 and 100. If you don't use this parameter, ListVirtualServices returns
	// up to 100 results and a nextToken value if applicable.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListVirtualServicesPaginator is a paginator for ListVirtualServices
type ListVirtualServicesPaginator struct {
	options   ListVirtualServicesPaginatorOptions
	client    ListVirtualServicesAPIClient
	params    *ListVirtualServicesInput
	nextToken *string
	firstPage bool
}

// NewListVirtualServicesPaginator returns a new ListVirtualServicesPaginator
func NewListVirtualServicesPaginator(client ListVirtualServicesAPIClient, params *ListVirtualServicesInput, optFns ...func(*ListVirtualServicesPaginatorOptions)) *ListVirtualServicesPaginator {
	if params == nil {
		params = &ListVirtualServicesInput{}
	}

	options := ListVirtualServicesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListVirtualServicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListVirtualServicesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next ListVirtualServices page.
func (p *ListVirtualServicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListVirtualServicesOutput, error) {
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

	result, err := p.client.ListVirtualServices(ctx, &params, optFns...)
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
