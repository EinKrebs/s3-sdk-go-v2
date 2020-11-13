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

// Returns an array of StreamInfo objects. Each object describes a stream. To
// retrieve only streams that satisfy a specific condition, you can specify a
// StreamNameCondition.
func (c *Client) ListStreams(ctx context.Context, params *ListStreamsInput, optFns ...func(*Options)) (*ListStreamsOutput, error) {
	if params == nil {
		params = &ListStreamsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListStreams", params, optFns, addOperationListStreamsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListStreamsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListStreamsInput struct {

	// The maximum number of streams to return in the response. The default is 10,000.
	MaxResults *int32

	// If you specify this parameter, when the result of a ListStreams operation is
	// truncated, the call returns the NextToken in the response. To get another batch
	// of streams, provide this token in your next request.
	NextToken *string

	// Optional: Returns only streams that satisfy a specific condition. Currently, you
	// can specify only the prefix of a stream name as a condition.
	StreamNameCondition *types.StreamNameCondition
}

type ListStreamsOutput struct {

	// If the response is truncated, the call returns this element with a token. To get
	// the next batch of streams, use this token in your next request.
	NextToken *string

	// An array of StreamInfo objects.
	StreamInfoList []types.StreamInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListStreamsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListStreams{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListStreams{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListStreams(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListStreams(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesisvideo",
		OperationName: "ListStreams",
	}
}
