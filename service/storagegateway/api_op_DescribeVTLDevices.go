// Code generated by smithy-go-codegen DO NOT EDIT.

package storagegateway

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a description of virtual tape library (VTL) devices for the specified
// tape gateway. In the response, AWS Storage Gateway returns VTL device
// information. This operation is only supported in the tape gateway type.
func (c *Client) DescribeVTLDevices(ctx context.Context, params *DescribeVTLDevicesInput, optFns ...func(*Options)) (*DescribeVTLDevicesOutput, error) {
	if params == nil {
		params = &DescribeVTLDevicesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeVTLDevices", params, optFns, addOperationDescribeVTLDevicesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeVTLDevicesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeVTLDevicesInput
type DescribeVTLDevicesInput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	//
	// This member is required.
	GatewayARN *string

	// Specifies that the number of VTL devices described be limited to the specified
	// number.
	Limit *int32

	// An opaque string that indicates the position at which to begin describing the
	// VTL devices.
	Marker *string

	// An array of strings, where each string represents the Amazon Resource Name (ARN)
	// of a VTL device. All of the specified VTL devices must be from the same gateway.
	// If no VTL devices are specified, the result will contain all devices on the
	// specified gateway.
	VTLDeviceARNs []string
}

// DescribeVTLDevicesOutput
type DescribeVTLDevicesOutput struct {

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation to
	// return a list of gateways for your account and AWS Region.
	GatewayARN *string

	// An opaque string that indicates the position at which the VTL devices that were
	// fetched for description ended. Use the marker in your next request to fetch the
	// next set of VTL devices in the list. If there are no more VTL devices to
	// describe, this field does not appear in the response.
	Marker *string

	// An array of VTL device objects composed of the Amazon Resource Name (ARN) of the
	// VTL devices.
	VTLDevices []types.VTLDevice

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeVTLDevicesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeVTLDevices{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeVTLDevices{}, middleware.After)
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
	if err = addOpDescribeVTLDevicesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeVTLDevices(options.Region), middleware.Before); err != nil {
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

// DescribeVTLDevicesAPIClient is a client that implements the DescribeVTLDevices
// operation.
type DescribeVTLDevicesAPIClient interface {
	DescribeVTLDevices(context.Context, *DescribeVTLDevicesInput, ...func(*Options)) (*DescribeVTLDevicesOutput, error)
}

var _ DescribeVTLDevicesAPIClient = (*Client)(nil)

// DescribeVTLDevicesPaginatorOptions is the paginator options for
// DescribeVTLDevices
type DescribeVTLDevicesPaginatorOptions struct {
	// Specifies that the number of VTL devices described be limited to the specified
	// number.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeVTLDevicesPaginator is a paginator for DescribeVTLDevices
type DescribeVTLDevicesPaginator struct {
	options   DescribeVTLDevicesPaginatorOptions
	client    DescribeVTLDevicesAPIClient
	params    *DescribeVTLDevicesInput
	nextToken *string
	firstPage bool
}

// NewDescribeVTLDevicesPaginator returns a new DescribeVTLDevicesPaginator
func NewDescribeVTLDevicesPaginator(client DescribeVTLDevicesAPIClient, params *DescribeVTLDevicesInput, optFns ...func(*DescribeVTLDevicesPaginatorOptions)) *DescribeVTLDevicesPaginator {
	if params == nil {
		params = &DescribeVTLDevicesInput{}
	}

	options := DescribeVTLDevicesPaginatorOptions{}
	if params.Limit != nil {
		options.Limit = *params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeVTLDevicesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeVTLDevicesPaginator) HasMorePages() bool {
	return p.firstPage || p.nextToken != nil
}

// NextPage retrieves the next DescribeVTLDevices page.
func (p *DescribeVTLDevicesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeVTLDevicesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.Limit = limit

	result, err := p.client.DescribeVTLDevices(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken && prevToken != nil && p.nextToken != nil && *prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeVTLDevices(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "storagegateway",
		OperationName: "DescribeVTLDevices",
	}
}
