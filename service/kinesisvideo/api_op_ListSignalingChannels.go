// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesisvideo

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/kinesisvideo/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an array of ChannelInfo objects. Each object describes a signaling
// channel. To retrieve only those channels that satisfy a specific condition, you
// can specify a ChannelNameCondition.
func (c *Client) ListSignalingChannels(ctx context.Context, params *ListSignalingChannelsInput, optFns ...func(*Options)) (*ListSignalingChannelsOutput, error) {
	if params == nil {
		params = &ListSignalingChannelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSignalingChannels", params, optFns, addOperationListSignalingChannelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSignalingChannelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSignalingChannelsInput struct {

	// Optional: Returns only the channels that satisfy a specific condition.
	ChannelNameCondition *types.ChannelNameCondition

	// The maximum number of channels to return in the response. The default is 500.
	MaxResults *int32

	// If you specify this parameter, when the result of a ListSignalingChannels
	// operation is truncated, the call returns the NextToken in the response. To get
	// another batch of channels, provide this token in your next request.
	NextToken *string
}

type ListSignalingChannelsOutput struct {

	// An array of ChannelInfo objects.
	ChannelInfoList []types.ChannelInfo

	// If the response is truncated, the call returns this element with a token. To get
	// the next batch of streams, use this token in your next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListSignalingChannelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSignalingChannels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSignalingChannels{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSignalingChannels(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListSignalingChannels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "ListSignalingChannels",
	}
}
