// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves information about a gateway.
func (c *Client) DescribeGateway(ctx context.Context, params *DescribeGatewayInput, optFns ...func(*Options)) (*DescribeGatewayOutput, error) {
	if params == nil {
		params = &DescribeGatewayInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeGateway", params, optFns, addOperationDescribeGatewayMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeGatewayOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeGatewayInput struct {

	// The ID of the gateway device.
	//
	// This member is required.
	GatewayId *string
}

type DescribeGatewayOutput struct {

	// The date the gateway was created, in Unix epoch time.
	//
	// This member is required.
	CreationDate *time.Time

	// The ARN
	// (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of
	// the gateway, which has the following format.
	// arn:${Partition}:iotsitewise:${Region}:${Account}:gateway/${GatewayId}
	//
	// This member is required.
	GatewayArn *string

	// A list of gateway capability summaries that each contain a namespace and status.
	// Each gateway capability defines data sources for the gateway. To retrieve a
	// capability configuration's definition, use
	// DescribeGatewayCapabilityConfiguration
	// (https://docs.aws.amazon.com/iot-sitewise/latest/APIReference/API_DescribeGatewayCapabilityConfiguration.html).
	//
	// This member is required.
	GatewayCapabilitySummaries []types.GatewayCapabilitySummary

	// The ID of the gateway device.
	//
	// This member is required.
	GatewayId *string

	// The name of the gateway.
	//
	// This member is required.
	GatewayName *string

	// The date the gateway was last updated, in Unix epoch time.
	//
	// This member is required.
	LastUpdateDate *time.Time

	// The gateway's platform.
	GatewayPlatform *types.GatewayPlatform

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeGatewayMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeGateway{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeGateway{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addEndpointPrefix_opDescribeGatewayMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDescribeGatewayValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeGateway(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDescribeGatewayMiddleware struct {
}

func (*endpointPrefix_opDescribeGatewayMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeGatewayMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "edge." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDescribeGatewayMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribeGatewayMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeGateway(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribeGateway",
	}
}
